package mock

import (
	"bufio"
	"net"
)

type Server struct {
	Port string
}

func NewServer(port string) *Server {
	return &Server{port}
}

func (server *Server) Listen() {
	listener, _ := net.Listen("tcp", server.Port)
	for {
		connection, _ := listener.Accept()
		go server.Handle(connection)
	}
}

func (server *Server) Handle(connection net.Conn) {
	rd := bufio.NewReader(connection)
	wr := bufio.NewWriter(connection)
	buffer := bufio.NewReadWriter(rd, wr)

	for {
		line, _ := buffer.ReadString('\r')
		_, _ = buffer.WriteString(line)
		_ = buffer.Flush()
	}
}