package main

import (
	"bufio"
	"github.com/MattLaidlaw/GoDB/pkg/interaction"
	"github.com/MattLaidlaw/GoDB/pkg/storage"
	"os"
)

func main() {

	m := storage.NewMap()
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	interaction.EventLoop(m, reader, writer)

}
