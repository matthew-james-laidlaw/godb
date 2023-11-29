package main

import (
	"GoDB/pkg/godb"
	"log"
)

func main() {

	s := godb.NewServer()
	err := s.Listen(":8000")
	if err != nil {
		log.Fatalln(err)
	}

}
