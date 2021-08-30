package godb

// The Handler type abstracts StorageEngine methods for use with the JSON RPC protocol.
type Handler struct {
	storage Engine
}

func NewHandler(storage Engine) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) Set(req *SetRequest, res *SetResult) error {
	*res = SetResult {
		InsertedCount: h.storage.Set(req.Key, req.Val),
	}
	return nil
}

func (h *Handler) Get(req *GetRequest, res *GetResult) error {
	*res = GetResult {
		Val: h.storage.Get(req.Key),
	}
	return nil
}

func (h *Handler) Del( req *DelRequest, res *DelResult) error {
	*res = DelResult {
		DeletedCount: h.storage.Del(req.Key),
	}
	return nil
}
