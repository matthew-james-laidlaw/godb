package server

import (
	"bufio"
	"github.com/MattLaidlaw/GoDB/pkg/parse"
	"github.com/MattLaidlaw/GoDB/pkg/storage"
	"io"
	"log"
	"net"
)

type Server struct {
	port string
	engine storage.Engine
}

func NewServer(port string, engine storage.Engine) *Server {
	return &Server{port, engine}
}

func (s *Server) Listen() {
	listener, err := net.Listen("tcp", s.port)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("== server listening on port", s.port[1:])
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("== accepted connection from", conn.RemoteAddr().String())
		go s.Handle(conn)
	}
}

func (s *Server) Handle(conn net.Conn) {
	rd := bufio.NewReader(conn)
	wr := bufio.NewWriter(conn)
	rw := bufio.NewReadWriter(rd, wr)

	for {
		line, err := rw.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		stmt, _ := parse.Parse(line)
		line = stmt.Execute(s.engine) + "\n"

		_, err = rw.WriteString(line)
		if err != nil {
			log.Fatalln(err)
		}
		err = rw.Flush()
		if err != nil {
			log.Fatalln(err)
		}
	}

	_ = conn.Close()
}