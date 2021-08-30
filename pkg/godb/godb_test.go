package godb

import (
	"github.com/MattLaidlaw/GoDB/pkg/assert"
	"net/rpc/jsonrpc"
	"testing"
)

func TestMap_SetKeyValue(t *testing.T) {
	m := NewBasicMap()
	insertedCount := m.Set("key", "value")
	assert.ExpectEq(insertedCount, 1, t)
}

func TestMap_GetKey(t *testing.T) {
	m := NewBasicMap()
	_ = m.Set("key", "value")
	value := m.Get("key")
	assert.ExpectEq(value, "value", t)
}

func TestMap_GetNonExistentKey(t *testing.T) {
	m := NewBasicMap()
	value := m.Get("key")
	assert.ExpectEq(value, "", t)
}

func TestMap_ReplaceKeyValue(t *testing.T) {
	m := NewBasicMap()
	_ = m.Set("key", "value")
	_ = m.Set("key", "diff")
	value := m.Get("key")
	assert.ExpectEq(value, "diff", t)
}

func TestMap_DelKey(t *testing.T) {
	m := NewBasicMap()
	_ = m.Set("key", "value")
	deletedCount := m.Del("key")
	value := m.Get("key")
	assert.ExpectEq(deletedCount, 1, t)
	assert.ExpectEq(value, "", t)
}

func TestMap_DelNonexistentKey(t *testing.T) {
	m := NewBasicMap()
	deletedCount := m.Del("key")
	assert.ExpectEq(deletedCount, 0, t)
}

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

func TestServer(t *testing.T) {
	storage := NewBasicMap()
	srv, err := NewServer(storage)
	assert.ExpectEq(err, nil, t)
	go srv.Listen()

	client, err := jsonrpc.Dial(Network, Address)
	assert.ExpectEq(err, nil, t)

	setReq := &SetRequest{"key", "val"}
	setRes := new(SetResult)
	err = client.Call("Handler.Set", setReq, setRes)
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(setRes, &SetResult{1}, t)

	getReq := &GetRequest{"key"}
	getRes := new(GetResult)
	err = client.Call("Handler.Get", getReq, getRes)
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(getRes, &GetResult{"val"}, t)

	delReq := &DelRequest{"key"}
	delRes := new(DelResult)
	err = client.Call("Handler.Del", delReq, delRes)
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(delRes, &DelResult{1}, t)

	getReq = &GetRequest{"key"}
	getRes = new(GetResult)
	err = client.Call("Handler.Get", getReq, getRes)
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(getRes, &GetResult{""}, t)
}
