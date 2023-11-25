package godb

import (
	"testing"

	"github.com/MattLaidlaw/go-assert"
	"github.com/MattLaidlaw/go-jsonrpc2"
)

func TestMap_SetKeyValue(t *testing.T) {
	m := NewBasicMap()
	insertedCount := m.Set("key", "value")
	assert.ExpectEq(insertedCount, float64(1), t)
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
	assert.ExpectEq(deletedCount, float64(1), t)
	assert.ExpectEq(value, "", t)
}

func TestMap_DelNonexistentKey(t *testing.T) {
	m := NewBasicMap()
	deletedCount := m.Del("key")
	assert.ExpectEq(deletedCount, float64(0), t)
}

func TestHandler(t *testing.T) {
	handler := NewHandler()

	value := handler.Get("key")
	assert.ExpectEq(value, "", t)

	insertedCount := handler.Set("key", "val")
	assert.ExpectEq(insertedCount, float64(1), t)

	value = handler.Get("key")
	assert.ExpectEq(value, "val", t)

	insertedCount = handler.Set("key", "otherval")
	assert.ExpectEq(insertedCount, float64(1), t)

	value = handler.Get("key")
	assert.ExpectEq(value, "otherval", t)

	deletedCount := handler.Del("key")
	assert.ExpectEq(deletedCount, float64(1), t)

	deletedCount = handler.Del("key")
	assert.ExpectEq(deletedCount, float64(0), t)
}

func TestServer(t *testing.T) {
	server := NewServer()
	go server.Listen()

	rpc, err := jsonrpc2.Dial(Address)
	assert.ExpectEq(err, nil, t)

	res, err := rpc.Call("Handler.Get", "key")
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(res.Result, "", t)

	res, err = rpc.Call("Handler.Set", "key", "val")
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(res.Result, float64(1), t)

	res, err = rpc.Call("Handler.Get", "key")
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(res.Result, "val", t)

	res, err = rpc.Call("Handler.Set", "key", "otherval")
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(res.Result, float64(1), t)

	res, err = rpc.Call("Handler.Get", "key")
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(res.Result, "otherval", t)

	res, err = rpc.Call("Handler.Del", "key")
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(res.Result, float64(1), t)

	res, err = rpc.Call("Handler.Del", "key")
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(res.Result, float64(0), t)
}
