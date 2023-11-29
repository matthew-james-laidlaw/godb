package godb

import (
	"encoding/json"
	"log"
	"net"
)

type Server struct {
	storage *StorageEngine
}

func NewServer() *Server {
	return &Server{storage: NewStorageEngine()}
}

func (s *Server) Listen(addr string) error {
	listener, err := net.Listen("tcp", ":8000")

	if err != nil {
		return err
	}

	log.Println("== listening for connections at", listener.Addr())

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println("== accepted connection from", connection.RemoteAddr())

		err = s.HandleConnection(connection)
		if err != nil {
			log.Println(err)
			continue
		}

		err = connection.Close()
		if err != nil {
			log.Println(err)
			continue
		}
	}

	return nil
}

func (s *Server) HandleConnection(conn net.Conn) error {
	reader := json.NewDecoder(conn)
	writer := json.NewEncoder(conn)

	for {
		request := &Request{}
		err := reader.Decode(&request)
		if err != nil {
			return err
		}

		response, err := s.storage.Execute(request)
		if err != nil {
			return err
		}

		err = writer.Encode(response)
		if err != nil {
			return err
		}
	}
}
