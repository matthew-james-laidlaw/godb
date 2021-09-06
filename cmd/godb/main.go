package main

import (
	"godb/pkg/godb"
)

func main() {
	//storage := godb.NewBasicMap()
	srv := godb.NewServer()
	srv.Listen()
}
