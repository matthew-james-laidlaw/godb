package main

import (
	"GoDB/pkg/godb"
	"log"
)

func main() {

	s := godb.NewServer()
	err := s.Listen("localhost:6532")
	if err != nil {
		log.Fatalln(err)
	}

}
