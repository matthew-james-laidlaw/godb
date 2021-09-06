package godb

import (
	"github.com/MattLaidlaw/go-jsonrpc2"
	"log"
)

const (
	Port = "6342"
	Address = "localhost:" + Port
)

// The Server type wraps an RPC method handler and provides a light abstraction on top of a JSON RPC server for handling
// concurrent connections.
type Server struct {
	rpc *jsonrpc2.Server
}

// The NewServer method creates a Server object, encapsulating a StorageEngine implementation. Before returning a
// Server, this method registers an RPC method handler with the RPC DefaultServer.
func NewServer() *Server {
	s := &Server {
		rpc: jsonrpc2.NewServer(),
	}
	s.rpc.Register(Handler{})
	return s
}

func (s *Server) Listen() {
	err := s.rpc.Listen(Port)
	if err != nil {
		log.Fatalln(err)
	}
}


