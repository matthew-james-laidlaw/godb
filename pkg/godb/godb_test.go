package godb

import (
	"github.com/MattLaidlaw/GoDB/pkg/assert"
	"net/rpc/jsonrpc"
	"testing"
)

func TestServer(t *testing.T) {
	storage := NewBasicMap()
	srv, err := NewServer(storage)
	assert.ExpectEq(err, nil, t)
	go srv.Listen()

	client, err := jsonrpc.Dial(Network, Address)

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
