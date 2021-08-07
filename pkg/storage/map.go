package storage

type Map struct {
	Data map[string]map[string]string
}

func NewMap() *Map {
	return &Map{map[string]map[string]string{}}
}

func (m *Map) Set(key string, field string, value string) int {
	if m.Data == nil {
		m.Data = map[string]map[string]string{}
	}
	if m.Data[key] == nil {
		m.Data[key] = map[string]string{}
	}
	m.Data[key][field] = value
	return 1
}

func (m *Map) Get(key string, field string) (string, bool) {
	value, ok := m.Data[key][field]
	return value, ok
}

func (m *Map) Del(key string, field string) int {
	if _, ok := m.Data[key][field]; ok {
		delete(m.Data[key], field)
		return 1
	}
	return 0
}