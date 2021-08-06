package interaction

import (
	"GoDB/internal/pkg/parse"
	"GoDB/internal/pkg/storage"
	"bufio"
	"fmt"
	"io"
)

func EventLoop(store storage.ObjectStore, reader *bufio.Reader, writer *bufio.Writer) {
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err == io.EOF || line == "q\n" || line[0] == 4 {
			break
		}
		if err != nil {
			fmt.Println(err)
		}

		stmt, err := parse.Parse(line)
		if err != nil {
			_, err = writer.WriteString(err.Error() + "\n")
			if err != nil {

			}
			err = writer.Flush()
			if err != nil {

			}
			continue
		}

		res := stmt.Execute(store)
		_, err = writer.WriteString(res + "\n")
		if err != nil {
			fmt.Println(err)
		}
		err = writer.Flush()
		if err != nil {
			fmt.Println(err)
		}
	}
}
