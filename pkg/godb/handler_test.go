package godb

import (
	"github.com/MattLaidlaw/GoDB/pkg/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	storage := NewBasicMap()
	handler := NewHandler(storage)

	setReq := &SetRequest{"key", "val"}
	setRes := new(SetResult)
	err := handler.Set(setReq, setRes)
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(setRes, &SetResult{1}, t)

	getReq := &GetRequest{"key"}
	getRes := new(GetResult)
	err = handler.Get(getReq, getRes)
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(getRes, &GetResult{"val"}, t)

	delReq := &DelRequest{"key"}
	delRes := new(DelResult)
	err = handler.Del(delReq, delRes)
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(delRes, &DelResult{1}, t)

	getReq = &GetRequest{"key"}
	getRes = new(GetResult)
	err = handler.Get(getReq, getRes)
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(getRes, &GetResult{""}, t)
}
