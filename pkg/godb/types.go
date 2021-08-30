package godb

type KeyType = string
type ValType = string

type SetRequest struct {
	Key KeyType
	Val ValType
}

type GetRequest struct {
	Key KeyType
}

type DelRequest struct {
	Key KeyType
}

type SetResult struct {
	InsertedCount int
}

type GetResult struct {
	Val ValType
}

type DelResult struct {
	DeletedCount int
}
