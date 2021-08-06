package parse

import (
	"GoDB/internal/pkg/storage"
	"fmt"
	"strconv"
	"strings"
)

func Parse(input string) (Statement, error) {
	fields := strings.Fields(input)

	if len(fields) < 1 {
		return nil, fmt.Errorf("empty input")
	}

	switch fields[0] {
	case "SET":
		return ParseSet(fields[1:])
	case "GET":
		return ParseGet(fields[1:])
	case "DEL":
		return ParseDel(fields[1:])
	default:
		return nil, fmt.Errorf("unexpected statement: %v", fields[0])
	}
}

func ParseSet(args []string) (Statement, error) {
	if len(args) < 3 {
		return nil, fmt.Errorf("expected 3 arguments for SET statement, got %d", len(args))
	}
	return &Set {
		Key: args[0],
		Field: args[1],
		Value: args[2],
	}, nil
}

func ParseGet(args []string) (Statement, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("expected 2 arguments for GET statement, got %d", len(args))
	}
	return &Get {
		Key: args[0],
		Field: args[1],
	}, nil
}

func ParseDel(args []string) (Statement, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("expected 2 arguments for DEL statement, got %d", len(args))
	}
	return &Del {
		Key: args[0],
		Field: args[1],
	}, nil
}

type Statement interface {
	Execute(store storage.ObjectStore) string
}

type Set struct {
	Key string
	Field string
	Value string
}

func (s *Set) Execute(store storage.ObjectStore) string {
	insertedCount := store.Set(s.Key, s.Field, s.Value)
	payload := "insertedCount: " + strconv.Itoa(insertedCount)
	return payload
}

type Get struct {
	Key string
	Field string
}

func (g *Get) Execute(store storage.ObjectStore) string {
	if value, ok := store.Get(g.Key, g.Field); ok {
		return value
	}
	return "item not found"
}

type Del struct {
	Key string
	Field string
}

func (d *Del) Execute(store storage.ObjectStore) string {
	if deletedCount := store.Del(d.Key, d.Field); deletedCount > 0 {
		return "deletedCount: " + strconv.Itoa(deletedCount)
	}
	return "item not found"
}
