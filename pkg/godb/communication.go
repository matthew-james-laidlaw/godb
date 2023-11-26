package godb

type Request struct {
	Method string   `json:"method"`
	Params []string `json:"params,omitempty"`
}

type Response struct {
	Result string `json:"result,omitempty"`
}
