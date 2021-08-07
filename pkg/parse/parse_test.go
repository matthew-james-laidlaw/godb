package parse

import (
	"fmt"
	"github.com/MattLaidlaw/GoDB/pkg/assert"
	"github.com/MattLaidlaw/GoDB/pkg/storage"
	"testing"
)

func TestParse_Empty(t *testing.T) {
	stmt, err := Parse("")
	assert.ExpectEq(stmt, nil, t)
	assert.ExpectEq(err, fmt.Errorf("empty input"), t)
}

func TestParse_UnexpectedStatement(t *testing.T) {
	stmt, err := Parse("UNEXPECTED key field value")
	assert.ExpectEq(stmt, nil, t)
	assert.ExpectEq(err, fmt.Errorf("unexpected statement: UNEXPECTED"), t)
}

func TestParse_Set(t *testing.T) {
	stmt, err := Parse("SET key field value")
	assert.ExpectEq(stmt, &Set{"key", "field", "value"}, t)
	assert.ExpectEq(err, nil, t)
}

func TestParse_SetInsufficientArgs(t *testing.T) {
	stmt, err := Parse("SET key field")
	assert.ExpectEq(stmt, nil, t)
	assert.ExpectEq(err, fmt.Errorf("expected 3 arguments for SET statement, got 2"), t)
}

func TestParse_SetExtraArgs(t *testing.T) {
	stmt, err := Parse("SET key field value extra")
	assert.ExpectEq(stmt, &Set{"key", "field", "value"}, t)
	assert.ExpectEq(err, nil, t)
}

func TestParse_Get(t *testing.T) {
	stmt, err := Parse("GET key field")
	assert.ExpectEq(stmt, &Get{"key", "field"}, t)
	assert.ExpectEq(err, nil, t)
}

func TestParse_GetInsufficientArgs(t *testing.T) {
	stmt, err := Parse("GET key")
	assert.ExpectEq(stmt, nil, t)
	assert.ExpectEq(err, fmt.Errorf("expected 2 arguments for GET statement, got 1"), t)
}

func TestParse_GetExtraArgs(t *testing.T) {
	stmt, err := Parse("GET key field extra")
	assert.ExpectEq(stmt, &Get{"key", "field"}, t)
	assert.ExpectEq(err, nil, t)
}

func TestParse_Del(t *testing.T) {
	stmt, err := Parse("DEL key field")
	assert.ExpectEq(stmt, &Del{"key", "field"}, t)
	assert.ExpectEq(err, nil, t)
}

func TestParse_DelInsufficientArgs(t *testing.T) {
	stmt, err := Parse("DEL key")
	assert.ExpectEq(stmt, nil, t)
	assert.ExpectEq(err, fmt.Errorf("expected 2 arguments for DEL statement, got 1"), t)
}

func TestParse_DelExtraArgs(t *testing.T) {
	stmt, err := Parse("DEL key field extra")
	assert.ExpectEq(stmt, &Del{"key", "field"}, t)
	assert.ExpectEq(err, nil, t)
}

func TestStatement_Execute(t *testing.T) {
	m := storage.NewMap()
	stmt, _ := Parse("SET key field value")
	res := stmt.Execute(m)
	assert.ExpectEq(res, "insertedCount: 1", t)
	stmt, _ = Parse("GET key field")
	res = stmt.Execute(m)
	assert.ExpectEq(res, "value", t)
	stmt, _ = Parse("DEL key field")
	res = stmt.Execute(m)
	assert.ExpectEq(res, "deletedCount: 1", t)
}
