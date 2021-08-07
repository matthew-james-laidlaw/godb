package interaction

import (
	"bufio"
	"github.com/MattLaidlaw/GoDB/pkg/assert"
	"github.com/MattLaidlaw/GoDB/pkg/storage"
	"strings"
	"testing"
)

func TestEventLoop_NoInput(t *testing.T) {
	m := storage.NewMap()
	reader := bufio.NewReader(strings.NewReader(""))
	output := &strings.Builder{}
	writer := bufio.NewWriter(output)

	EventLoop(m, reader, writer)

	assert.ExpectEq(output.String(), "> ", t)
}

func TestEventLoop_EmptyLine(t *testing.T) {
	m := storage.NewMap()
	reader := bufio.NewReader(strings.NewReader("\n"))
	output := &strings.Builder{}
	writer := bufio.NewWriter(output)

	EventLoop(m, reader, writer)

	assert.ExpectEq(output.String(), "> empty input\n> ", t)
}

func TestEventLoop_Quit(t *testing.T) {
	m := storage.NewMap()
	reader := bufio.NewReader(strings.NewReader("q\n"))
	output := &strings.Builder{}
	writer := bufio.NewWriter(output)

	EventLoop(m, reader, writer)

	assert.ExpectEq(output.String(), "> ", t)
}

func TestEventLoop(t *testing.T) {
	m := storage.NewMap()
	reader := bufio.NewReader(strings.NewReader("SET key field value\nGET key field\nDEL key field\nq\n"))
	output := &strings.Builder{}
	writer := bufio.NewWriter(output)

	EventLoop(m, reader, writer)

	assert.ExpectEq(output.String(), "> insertedCount: 1\n> value\n> deletedCount: 1\n> ", t)
}
