package server

import (
	"bufio"
	"github.com/MattLaidlaw/GoDB/pkg/assert"
	"github.com/MattLaidlaw/GoDB/pkg/storage"
	"log"
	"net"
	"testing"
)

func StartServer() {
	StorageEngine := storage.NewMap()
	server := NewServer(":8080", StorageEngine)
	server.Listen()
}

func RunClient(input []string, expected []string, t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for i, line := range input {
		writer.WriteString(line)
		writer.Flush()
		response, _ := reader.ReadString('\n')
		assert.ExpectEq(response, expected[i], t)
	}

	conn.Close()
}

func TestServer(t *testing.T) {
}

func TestServer_NoInput(t *testing.T) {
	go StartServer()
	var input []string
	var expected []string
	RunClient(input, expected, t)
}

func TestServer_EmptyLine(t *testing.T) {
	go StartServer()
	input := []string{"\n"}
	expected := []string{"no input\n"}
	go RunClient(input, expected, t)
}

func TestServer_Quit(t *testing.T) {
	go StartServer()
	input := []string{"q\n"}
	var expected []string
	go RunClient(input, expected, t)
}

func TestServer_Input(t *testing.T) {
	go StartServer()
	input := []string{"SET key field value\n", "GET key field\n", "DEL key field\n", "GET key field\nq\n"}
	expected := []string{"insertedCount: 1\n", "value\n", "deletedCount: 1\n", "item not found\n"}
	go RunClient(input, expected, t)
}
