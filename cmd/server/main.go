package main

import (
	"log"

	"github.com/matthew-james-laidlaw/godb/pkg/godb"
)

func main() {

	s := godb.NewServer()
	err := s.Listen(":8000")
	if err != nil {
		log.Fatalln(err)
	}

}
