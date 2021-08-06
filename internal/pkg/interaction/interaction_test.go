package interaction

import (
	"GoDB/internal/pkg/assert"
	"GoDB/internal/pkg/storage"
	"bufio"
	"strings"
	"testing"
)

func TestEventLoop_NoInput(t *testing.T) {
	m := storage.NewMap()
	reader := bufio.NewReader(strings.NewReader(""))
	output := &strings.Builder{}
	writer := bufio.NewWriter(output)

	EventLoop(m, reader, writer)

	assert.ExpectEq(output.String(), "", t)
}

func TestEventLoop_EmptyLine(t *testing.T) {
	m := storage.NewMap()
	reader := bufio.NewReader(strings.NewReader("\n"))
	output := &strings.Builder{}
	writer := bufio.NewWriter(output)

	EventLoop(m, reader, writer)

	assert.ExpectEq(output.String(), "empty input", t)
}

func TestEventLoop_Quit(t *testing.T) {
	m := storage.NewMap()
	reader := bufio.NewReader(strings.NewReader("q\n"))
	output := &strings.Builder{}
	writer := bufio.NewWriter(output)

	EventLoop(m, reader, writer)

	assert.ExpectEq(output.String(), "", t)
}

func TestEventLoop(t *testing.T) {
	m := storage.NewMap()
	reader := bufio.NewReader(strings.NewReader("SET key field value\nGET key field\nDEL key field\nq\n"))
	output := &strings.Builder{}
	writer := bufio.NewWriter(output)

	EventLoop(m, reader, writer)

	assert.ExpectEq(output.String(), "insertedCount: 1valuedeletedCount: 1", t)
}
