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
	Port string
	StorageEngine storage.ObjectStore
}

func NewServer(port string, store storage.ObjectStore) *Server {
	return &Server{port, store}
}

func (s *Server) Listen() {
	// begin listening for TCP connections
	listener, err := net.Listen("tcp", s.Port)
	if err != nil {
		log.Fatalln("[FATAL] " + err.Error())
	}
	log.Println("[INFO] Server listening on port " + s.Port[1:])

	// spin up goroutines for each new connection
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("[WARN] " + err.Error())
		}
		log.Println("[INFO] Server accepted connection from " + conn.RemoteAddr().String())
		go s.Handle(conn)
	}
}

func (s *Server) Handle(conn net.Conn) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	// read-eval-print loop
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF || line == "q\n" || line[0] == 4 {
			log.Println("[INFO] EOF received from " + conn.RemoteAddr().String())
			break
		}
		if err != nil {
			log.Println("[ERROR] " + err.Error())
			break
		}

		stmt, err := parse.Parse(line)
		if err != nil {
			writer.WriteString(err.Error() + "\n")
			writer.Flush()
			continue
		}

		result := stmt.Execute(s.StorageEngine)
		writer.WriteString(result + "\n")
		writer.Flush()
	}

	conn.Close()
}