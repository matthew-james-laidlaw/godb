package interaction

import (
	"bufio"
	"fmt"
	"github.com/MattLaidlaw/GoDB/pkg/parse"
	"github.com/MattLaidlaw/GoDB/pkg/storage"
	"io"
)

func EventLoop(store storage.ObjectStore, reader *bufio.Reader, writer *bufio.Writer) {
	for {
		writer.WriteString("> ")
		writer.Flush()

		line, err := reader.ReadString('\n')
		if err == io.EOF || line == "q\n" || line[0] == 4 {
			break
		}
		if err != nil {
			fmt.Println(err)
		}

		stmt, err := parse.Parse(line)
		if err != nil {
			writer.WriteString(err.Error() + "\n")
			writer.Flush()
			continue
		}

		res := stmt.Execute(store)
		writer.WriteString(res + "\n")
		writer.Flush()
	}
}
