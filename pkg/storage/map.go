package storage

type Map struct {
	store map[string]string
}

func NewMap() *Map {
	store := make(map[string]string)
	return &Map{store}
}

func (m *Map) Set(k string, v string) int {
	m.store[k] = v
	return 1
}

func (m *Map) Get(k string) (string, bool) {
	v, ok := m.store[k]
	return v, ok
}

func (m *Map) Del(k string) int {
	if _, ok := m.store[k]; ok {
		delete(m.store, k)
		return 1
	}
	return 0
}