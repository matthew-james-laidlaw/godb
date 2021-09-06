package godb

// The Engine interface defines the supported methods on the GoDB storage engine.
type Engine interface {

	// The Set method inserts a key-value pair into the underlying storage engine and returns the count of inserted
	// elements.
	Set(key KeyType, value ValType) float64

	// The Get method queries the underlying storage engine for a key-value pair and returns one if found. Otherwise,
	// an empty value is returned.
	Get(key KeyType) ValType

	// The Del method attempts to delete a key-value pair in the underlying storage engine. If that pair exists, it is
	// deleted and a deleted-count of 1 is returned. Otherwise, nothing is deleted and a count of 0 is returned.
	Del(key string) float64

}

// The BasicMap type implements the Engine interface backed by a standard Golang map[KeyType]ValType
type BasicMap struct {
	store map[KeyType]ValType
}

func NewBasicMap() *BasicMap {
	store := make(map[KeyType]ValType)
	return &BasicMap{store}
}

func (m *BasicMap) Set(key KeyType, val ValType) float64 {
	m.store[key] = val
	return 1
}

func (m *BasicMap) Get(key KeyType) ValType {
	return m.store[key]
}

func (m *BasicMap) Del(key KeyType) float64 {
	if _, ok := m.store[key]; ok {
		delete(m.store, key)
		return 1
	}
	return 0
}
