package godb

import (
	"github.com/MattLaidlaw/GoDB/pkg/assert"
	"testing"
)

func TestMap_SetKeyValue(t *testing.T) {
	m := NewBasicMap()
	insertedCount := m.Set("key", "value")
	assert.ExpectEq(insertedCount, 1, t)
}

func TestMap_GetKey(t *testing.T) {
	m := NewBasicMap()
	_ = m.Set("key", "value")
	value := m.Get("key")
	assert.ExpectEq(value, "value", t)
}

func TestMap_GetNonExistentKey(t *testing.T) {
	m := NewBasicMap()
	value := m.Get("key")
	assert.ExpectEq(value, "", t)
}

func TestMap_ReplaceKeyValue(t *testing.T) {
	m := NewBasicMap()
	_ = m.Set("key", "value")
	_ = m.Set("key", "diff")
	value := m.Get("key")
	assert.ExpectEq(value, "diff", t)
}

func TestMap_DelKey(t *testing.T) {
	m := NewBasicMap()
	_ = m.Set("key", "value")
	deletedCount := m.Del("key")
	value := m.Get("key")
	assert.ExpectEq(deletedCount, 1, t)
	assert.ExpectEq(value, "", t)
}

func TestMap_DelNonexistentKey(t *testing.T) {
	m := NewBasicMap()
	deletedCount := m.Del("key")
	assert.ExpectEq(deletedCount, 0, t)
}
