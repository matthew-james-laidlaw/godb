package main

import (
	"github.com/MattLaidlaw/GoDB/pkg/server"
	"github.com/MattLaidlaw/GoDB/pkg/storage"
)

func main() {

	StorageEngine := storage.NewMap()
	Server := server.NewServer(":8080", StorageEngine)
	Server.Listen()

}
