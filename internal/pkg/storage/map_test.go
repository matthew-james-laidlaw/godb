package storage

import (
	"GoDB/internal/pkg/assert"
	"testing"
)

func TestMap_Set(t *testing.T) {
	m := NewMap()
	insertedCount := m.Set("key", "field", "value")
	assert.ExpectEq(insertedCount, 1, t)
}

func TestMap_Get(t *testing.T) {
	m := NewMap()
	_ = m.Set("key", "field", "value")
	value, ok := m.Get("key", "field")
	assert.ExpectEq(value, "value", t)
	assert.ExpectEq(ok, true, t)
}

func TestMap_GetNonExistentKey(t *testing.T) {
	m := NewMap()
	value, ok := m.Get("key", "field")
	assert.ExpectEq(value, "", t)
	assert.ExpectEq(ok, false, t)
}

func TestMap_GetNonExistentField(t *testing.T) {
	m := NewMap()
	_ = m.Set("key", "field", "value")
	value, ok := m.Get("key", "diff")
	assert.ExpectEq(value, "", t)
	assert.ExpectEq(ok, false, t)
}

func TestMap_ReplaceValue(t *testing.T) {
	m := NewMap()
	_ = m.Set("key", "field", "value")
	_ = m.Set("key", "field", "diff")
	value, ok := m.Get("key", "field")
	assert.ExpectEq(value, "diff", t)
	assert.ExpectEq(ok, true, t)
}

func TestMap_Del(t *testing.T) {
	m := NewMap()
	_ = m.Set("key", "field", "value")
	deletedCount := m.Del("key", "field")
	value, ok := m.Get("key", "field")
	assert.ExpectEq(deletedCount, 1, t)
	assert.ExpectEq(value, "", t)
	assert.ExpectEq(ok, false, t)
}

func TestMap_DelNonexistentKey(t *testing.T) {
	m := NewMap()
	deletedCount := m.Del("key", "field")
	assert.ExpectEq(deletedCount, 0, t)
}

func TestMap_DelNonExistentField(t *testing.T) {
	m := NewMap()
	_ = m.Set("key", "field", "value")
	deletedCount := m.Del("key", "diff")
	assert.ExpectEq(deletedCount, 0, t)
}