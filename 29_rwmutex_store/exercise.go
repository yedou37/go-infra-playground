package rwmutexstore

import "sync"

// Store is a read-mostly map protected by RWMutex.
type Store struct {
	mu sync.RWMutex
	m  map[string]int
}

// Set stores one key/value pair.
//
// TODO:
// - Use the write lock.
// - Allocate the map on first use.
func (s *Store) Set(key string, value int) {
	panic("TODO: implement Set")
}

// Get reads one key/value pair.
//
// TODO:
// - Use the read lock.
// - Return (0, false) for missing keys.
func (s *Store) Get(key string) (int, bool) {
	panic("TODO: implement Get")
}

// Len returns the number of entries.
//
// TODO:
// - Use the read lock.
func (s *Store) Len() int {
	panic("TODO: implement Len")
}
