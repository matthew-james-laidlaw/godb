package godb

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const (
	Network = "tcp"
	Port = "6342"
	Address = ":" + Port
)

// The Server type wraps an RPC method handler and provides a light abstraction on top of a JSON RPC server for handling
// concurrent connections.
type Server struct {
	handler *Handler
}

// The NewServer method creates a Server object, encapsulating a StorageEngine implementation. Before returning a
// Server, this method registers an RPC method handler with the RPC DefaultServer.
func NewServer(storage Engine) (*Server, error) {
	handler := NewHandler(storage)
	err := rpc.Register(handler)
	return &Server {
		handler: handler,
	}, err
}

// The Listen method accepts connections from GoDB clients and serves each connection on its own goroutine.
func (s *Server) Listen() {
	var listener net.Listener
	var err error

	if listener, err = net.Listen(Network, Address); err != nil {
		log.Fatalln(err)
	}
	log.Println("== listening on port", Port)

	for {
		var conn net.Conn
		var err error

		if conn, err = listener.Accept(); err != nil {
			log.Println(err)
			continue
		}
		log.Println("== accepted connection from", conn.RemoteAddr().String())

		go jsonrpc.ServeConn(conn)
	}
}
