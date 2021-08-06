// Package storage defines an interface for the data structure that backs the GoDB storage engine.
package storage

// The ObjectStore interface defines the supported methods on the GoDB storage engine.
type ObjectStore interface {

	// The Set method accepts a key and field to address the storage engine and a value to set at this location.
	// Set returns the count of inserted or updated records, which will always be 1.
	Set(key string, field string, value string) int

	// The Get method accepts a key and field to address the storage engine. Get returns the value at this location, if
	// it exists. If it does not exist, Get returns an empty string and a false boolean value.
	Get(key string, field string) (string, bool)

	// The Del method accepts a key and field to address the storage engine. Del attempts to erase the field-value pair
	// from the given key. If the address is valid, the item is removed, and the count of deleted items (1) is returned.
	// If the address is not valid, nothing is deleted, and the method returns a count of 0.
	Del(key string, field string) int

}