package storage

import (
	"github.com/MattLaidlaw/GoDB/pkg/assert"
	"testing"
)

func TestMap_Set(t *testing.T) {
	m := NewMap()
	insertedCount := m.Set("key", "value")
	assert.ExpectEq(insertedCount, 1, t)
}

func TestMap_Get(t *testing.T) {
	m := NewMap()
	_ = m.Set("key", "value")
	value, ok := m.Get("key")
	assert.ExpectEq(value, "value", t)
	assert.ExpectEq(ok, true, t)
}

func TestMap_GetNonExistentKey(t *testing.T) {
	m := NewMap()
	value, ok := m.Get("key")
	assert.ExpectEq(value, "", t)
	assert.ExpectEq(ok, false, t)
}

func TestMap_ReplaceValue(t *testing.T) {
	m := NewMap()
	_ = m.Set("key", "value")
	_ = m.Set("key", "diff")
	value, ok := m.Get("key")
	assert.ExpectEq(value, "diff", t)
	assert.ExpectEq(ok, true, t)
}

func TestMap_Del(t *testing.T) {
	m := NewMap()
	_ = m.Set("key", "value")
	deletedCount := m.Del("key")
	value, ok := m.Get("key")
	assert.ExpectEq(deletedCount, 1, t)
	assert.ExpectEq(value, "", t)
	assert.ExpectEq(ok, false, t)
}

func TestMap_DelNonexistentKey(t *testing.T) {
	m := NewMap()
	deletedCount := m.Del("key")
	assert.ExpectEq(deletedCount, 0, t)
}
