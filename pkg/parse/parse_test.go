package parse

import (
	"fmt"
	"github.com/MattLaidlaw/GoDB/pkg/assert"
	"testing"
)

func TestParse_Empty(t *testing.T) {
	stmt, err := Parse("\n")
	assert.ExpectEq(stmt, nil, t)
	assert.ExpectEq(err, fmt.Errorf("empty input"), t)
}

func TestParse_UnexpectedStatement(t *testing.T) {
	stmt, err := Parse("UNEXPECTED key value\n")
	assert.ExpectEq(stmt, nil, t)
	assert.ExpectEq(err, fmt.Errorf("invalid input: UNEXPECTED key value\n"), t)
}

func TestParse_Set(t *testing.T) {
	stmt, err := Parse("SET$key$value\n")
	assert.ExpectEq(stmt, &Set{"key", "value"}, t)
	assert.ExpectEq(err, nil, t)
}

func TestParse_SetInsufficientArgs(t *testing.T) {
	stmt, err := Parse("SET$key\n")
	assert.ExpectEq(stmt, nil, t)
	assert.ExpectEq(err, fmt.Errorf("expected 2 arguments, got 1"), t)
}

func TestParse_SetExtraArgs(t *testing.T) {
	stmt, err := Parse("SET$key$value$extra\n")
	assert.ExpectEq(stmt, &Set{"key", "value"}, t)
	assert.ExpectEq(err, nil, t)
}

func TestParse_Get(t *testing.T) {
	stmt, err := Parse("GET$key\n")
	assert.ExpectEq(stmt, &Get{"key"}, t)
	assert.ExpectEq(err, nil, t)
}

func TestParse_GetInsufficientArgs(t *testing.T) {
	stmt, err := Parse("GET\n")
	assert.ExpectEq(stmt, nil, t)
	assert.ExpectEq(err, fmt.Errorf("expected 1 argument, got 0"), t)
}

func TestParse_GetExtraArgs(t *testing.T) {
	stmt, err := Parse("GET$key$extra\n")
	assert.ExpectEq(stmt, &Get{"key"}, t)
	assert.ExpectEq(err, nil, t)
}

//
func TestParse_Del(t *testing.T) {
	stmt, err := Parse("DEL$key\n")
	assert.ExpectEq(stmt, &Del{"key"}, t)
	assert.ExpectEq(err, nil, t)
}

func TestParse_DelInsufficientArgs(t *testing.T) {
	stmt, err := Parse("DEL\n")
	assert.ExpectEq(stmt, nil, t)
	assert.ExpectEq(err, fmt.Errorf("expected 1 argument, got 0"), t)
}

func TestParse_DelExtraArgs(t *testing.T) {
	stmt, err := Parse("DEL$key$extra\n")
	assert.ExpectEq(stmt, &Del{"key"}, t)
	assert.ExpectEq(err, nil, t)
}
