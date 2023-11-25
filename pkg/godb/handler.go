package godb

var gStorage Engine = NewBasicMap()

// The Handler type abstracts StorageEngine methods for use with the JSON RPC protocol.
type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Set(key string, val string) float64 {
	return gStorage.Set(key, val)
}

func (h *Handler) Get(key string) string {
	return gStorage.Get(key)
}

func (h *Handler) Del(key string) float64 {
	return gStorage.Del(key)
}
