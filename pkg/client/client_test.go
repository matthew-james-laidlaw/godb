package client

import (
	"github.com/MattLaidlaw/GoDB/pkg/assert"
	"github.com/MattLaidlaw/GoDB/pkg/parse"
	"github.com/MattLaidlaw/GoDB/pkg/server"
	"github.com/MattLaidlaw/GoDB/pkg/storage"
	"testing"
)

func StartServer() {
	store := storage.NewMap()
	srv := server.NewServer(":8080", store)
	srv.Listen()
}

func TestClient(t *testing.T) {
	go StartServer()
	client := NewClient("localhost:8080")

	setChan := client.Set("key", "val")
	assert.ExpectEq(<- setChan, parse.SetResult{InsertedCount: 1}, t)

	getChan1 := client.Get("key")
	assert.ExpectEq(<- getChan1, parse.GetResult{Found: true, Value: "val"}, t)

	delChan := client.Del("key")
	assert.ExpectEq(<- delChan, parse.DelResult{DeletedCount: 1}, t)

	getChan2 := client.Get("key")
	assert.ExpectEq(<- getChan2, parse.GetResult{Found: false, Value: ""}, t)
}
