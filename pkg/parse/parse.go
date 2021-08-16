package parse

import (
	"encoding/json"
	"fmt"
	"github.com/MattLaidlaw/GoDB/pkg/storage"
	"strings"
)

type Statement interface {
	Execute(engine storage.Engine) string
}

type Set struct {
	k, v string
}

type SetResult struct {
	InsertedCount int `json:"inserted_count"`
}

type Get struct {
	k string
}

type GetResult struct {
	Found bool `json:"found"`
	Value string `json:"value"`
}

type Del struct {
	k string
}

type DelResult struct {
	DeletedCount int `json:"deleted_count"`
}

type Invalid struct {
	err error
}

func (cmd *Set) Execute(engine storage.Engine) string {
	insertedCount := engine.Set(cmd.k, cmd.v)
	res := &SetResult{insertedCount}
	str, _ := json.Marshal(&res)
	return string(str)
}

func (cmd *Get) Execute(engine storage.Engine) string {
	v, ok := engine.Get(cmd.k)
	res := &GetResult{ok, v}
	str, _ := json.Marshal(&res)
	return string(str)
}

func (cmd *Del) Execute(engine storage.Engine) string {
	deletedCount := engine.Del(cmd.k)
	res := &DelResult{deletedCount}
	str, _ := json.Marshal(&res)
	return string(str)
}

func (cmd *Invalid) Execute(engine storage.Engine) string {
	return cmd.err.Error()
}

func Parse(s string) (Statement, error) {
	if len(s) <= 1 {
		return nil, fmt.Errorf("empty input")
	}

	args := strings.FieldsFunc(s[:len(s)-1], func(r rune) bool {
		return r == '$'
	})

	switch args[0] {
	case "SET":
		if len(args) < 3 {
			return nil, fmt.Errorf("expected 2 arguments, got %d", len(args)-1)
		}
		return &Set{args[1], args[2]}, nil
	case "GET":
		if len(args) < 2 {
			return nil, fmt.Errorf("expected 1 argument, got %d", len(args)-1)
		}
		return &Get{args[1]}, nil
	case "DEL":
		if len(args) < 2 {
			return nil, fmt.Errorf("expected 1 argument, got %d", len(args)-1)
		}
		return &Del{args[1]}, nil
	default:
	}
	return nil, fmt.Errorf("invalid input: %s", s)
}
