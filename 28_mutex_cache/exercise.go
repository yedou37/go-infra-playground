package mutexcache

import "sync"

// Cache stores string values by key.
//
// This exercise is about the most common lock pattern in Go:
// use a mutex to protect a map and make the zero value usable.
type Cache struct {
	mu sync.Mutex
	m  map[string]string
}

// Set stores key/value in the cache.
//
// TODO:
// - Lock around map mutation.
// - Allocate the map on first use.
func (c *Cache) Set(key, value string) {
	panic("TODO: implement Set")
}

// Get returns the value for key.
//
// TODO:
// - Lock around map access.
// - Return ("", false) when the key is missing.
func (c *Cache) Get(key string) (string, bool) {
	panic("TODO: implement Get")
}

// Snapshot returns a copy of the internal map.
//
// TODO:
// - Lock around reading the map.
// - Preserve nil if the cache has never been written.
// - Return a copy so callers cannot mutate internal state.
func (c *Cache) Snapshot() map[string]string {
	panic("TODO: implement Snapshot")
}
