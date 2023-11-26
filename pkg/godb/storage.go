package godb

import (
	"errors"
)

type StorageEngine struct {
	storage map[string]string
}

func NewStorageEngine() *StorageEngine {
	return &StorageEngine{make(map[string]string)}
}

func (s *StorageEngine) Execute(request *Request) (*Response, error) {
	if request.Method == "get" {
		return s.Get(request)
	} else if request.Method == "set" {
		return s.Set(request)
	} else if request.Method == "del" {
		return s.Del(request)
	} else {
		return nil, errors.New("invalid method")
	}
}

func (s *StorageEngine) Get(request *Request) (*Response, error) {
	if len(request.Params) < 1 {
		return nil, errors.New("invalid params")
	}

	_, ok := s.storage[request.Params[0]]

	if ok {
		response := &Response{
			Result: s.storage[request.Params[0]],
		}
		return response, nil
	} else {
		response := &Response{
			Result: "",
		}
		return response, nil
	}
}

func (s *StorageEngine) Set(request *Request) (*Response, error) {
	if len(request.Params) < 2 {
		return nil, errors.New("invalid params")
	}

	s.storage[request.Params[0]] = request.Params[1]

	response := &Response{
		Result: "1",
	}

	return response, nil
}

func (s *StorageEngine) Del(request *Request) (*Response, error) {
	if len(request.Params) < 1 {
		return nil, errors.New("invalid params")
	}

	_, ok := s.storage[request.Params[0]]

	if ok {
		delete(s.storage, request.Params[0])
		response := &Response{
			Result: "1",
		}
		return response, nil
	} else {
		response := &Response{
			Result: "0",
		}
		return response, nil
	}
}
