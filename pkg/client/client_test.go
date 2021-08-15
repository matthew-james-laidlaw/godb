package client

import (
	"fmt"
	"github.com/MattLaidlaw/GoDB/pkg/mock"
	"testing"
)

func StartServer() {
	server := mock.NewServer(":8080")
	server.Listen()
}

func TestClient(t *testing.T) {
	go StartServer()
	client := NewClient("localhost:8080")
	c1 := client.Set("key", "val")
	fmt.Println(<- c1)
	c2 := client.Get("key")
	fmt.Println(<- c2)
	c3 := client.Del("key")
	fmt.Println(<- c3)
	c4 := client.Exit()
	fmt.Println(<- c4)
	c5 := client.Shutdown()
	fmt.Println(<- c5)
}
