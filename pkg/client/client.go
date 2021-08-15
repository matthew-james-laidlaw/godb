package client

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	Connection net.Conn
	Buffer *bufio.ReadWriter
}

func NewClient(connString string) *Client {
	conn, _ := net.Dial("tcp", connString)
	rd := bufio.NewReader(conn)
	wr := bufio.NewWriter(conn)
	buf := bufio.NewReadWriter(rd, wr)
	return &Client{conn, buf}
}

func (client *Client) Send(s string) string {
	_, _ = client.Buffer.WriteString(s)
	_ = client.Buffer.Flush()
	ret, _ := client.Buffer.ReadString('\r')
	return ret
}

func (client *Client) Set(k string, v string) chan string {
	ret := make(chan string)

	go func(k string, v string) {
		s := fmt.Sprintf("SET\n%d\n%s\n%d\n%s\r", len(k), k, len(v), v)
		ret <- client.Send(s)
	}(k, v)

	return ret
}

func (client *Client) Get(k string) chan string {
	ret := make(chan string)

	go func(k string) {
		s := fmt.Sprintf("GET\n%d\n%s\r", len(k), k)
		ret <- client.Send(s)
	}(k)

	return ret
}

func (client *Client) Del(k string) chan string {
	ret := make(chan string)

	go func(k string) {
		s := fmt.Sprintf("DEL\n%d\n%s\r", len(k), k)
		ret <- client.Send(s)
	}(k)

	return ret
}

func (client *Client) Exit() chan string {
	ret := make(chan string)

	go func() {
		s := "EXIT\r"
		ret <- client.Send(s)
	}()

	return ret
}

func (client *Client) Shutdown() chan string {
	ret := make(chan string)

	go func() {
		s := "SHUTDOWN\r"
		ret <- client.Send(s)
	}()

	return ret
}
