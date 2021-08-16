// Package storage defines an interface for the data structure that backs the GoDB storage engine.
package storage

// The Engine interface defines the supported methods on the GoDB storage engine.
type Engine interface {

	// The Set method inserts a key-value pair into the underlying storage engine and returns the count of inserted
	// elements.
	Set(key string, value string) int

	// The Get method queries the underlying storage engine for a key-value pair and returns one if found. Otherwise,
	// an empty string and a false boolean value are returned
	Get(key string) (string, bool)

	// The Del method attempts to delete a key-value pair in the underlying storage engine. If that pair exists, it is
	// deleted and a deleted-count of 1 is returned. Otherwise, nothing is deleted and a count of 0 is returned.
	Del(key string) int

}