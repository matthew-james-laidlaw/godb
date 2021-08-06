package main

import (
	"GoDB/internal/pkg/interaction"
	"GoDB/internal/pkg/storage"
	"bufio"
	"os"
)

func main() {

	m := storage.NewMap()
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	interaction.EventLoop(m, reader, writer)

}
