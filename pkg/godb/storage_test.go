package godb

import (
	"testing"

	"github.com/matthew-james-laidlaw/godb/pkg/assert"
)

func TestStorageEngine_GetNonExistentKey(t *testing.T) {
	s := NewStorageEngine()

	request := &Request{
		Method: "get",
		Params: []string{"key"},
	}

	response, err := s.Get(request)
	assert.ExpectEqual(err, nil, t)
	assert.ExpectEqual(response.Result, "", t)
}

func TestStorageEngine_GetExistingKey(t *testing.T) {
	s := NewStorageEngine()

	request := &Request{
		Method: "set",
		Params: []string{"key", "value"},
	}

	response, err := s.Set(request)
	assert.ExpectEqual(err, nil, t)
	assert.ExpectEqual(response.Result, "1", t)

	request = &Request{
		Method: "get",
		Params: []string{"key"},
	}

	response, err = s.Get(request)
	assert.ExpectEqual(err, nil, t)
	assert.ExpectEqual(response.Result, "value", t)
}

func TestStorageEngine_DelNonExistentKey(t *testing.T) {
	s := NewStorageEngine()

	request := &Request{
		Method: "del",
		Params: []string{"key"},
	}

	response, err := s.Del(request)
	assert.ExpectEqual(err, nil, t)
	assert.ExpectEqual(response.Result, "0", t)
}

func TestStorageEngine_DelExistingKey(t *testing.T) {
	s := NewStorageEngine()

	request := &Request{
		Method: "set",
		Params: []string{"key", "value"},
	}

	response, err := s.Set(request)
	assert.ExpectEqual(err, nil, t)
	assert.ExpectEqual(response.Result, "1", t)

	request = &Request{
		Method: "del",
		Params: []string{"key"},
	}

	response, err = s.Del(request)
	assert.ExpectEqual(err, nil, t)
	assert.ExpectEqual(response.Result, "1", t)

	request = &Request{
		Method: "get",
		Params: []string{"key"},
	}

	response, err = s.Get(request)
	assert.ExpectEqual(err, nil, t)
	assert.ExpectEqual(response.Result, "", t)
}

func TestStorageEngine_SetOverwritesKey(t *testing.T) {
	s := NewStorageEngine()

	request := &Request{
		Method: "set",
		Params: []string{"key", "value1"},
	}

	response, err := s.Set(request)
	assert.ExpectEqual(err, nil, t)
	assert.ExpectEqual(response.Result, "1", t)

	request = &Request{
		Method: "set",
		Params: []string{"key", "value2"},
	}

	response, err = s.Set(request)
	assert.ExpectEqual(err, nil, t)
	assert.ExpectEqual(response.Result, "1", t)

	request = &Request{
		Method: "get",
		Params: []string{"key"},
	}

	response, err = s.Get(request)
	assert.ExpectEqual(err, nil, t)
	assert.ExpectEqual(response.Result, "value2", t)
}
