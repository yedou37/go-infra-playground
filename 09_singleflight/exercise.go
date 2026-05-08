package singleflight

import "sync"

// Group deduplicates concurrent calls with the same key.
//
// This is useful when many goroutines may ask for the same expensive object:
// only one should load it, and the others should wait for and share the result.
type Group[T any] struct {
	mu sync.Mutex
	m  map[string]*call[T]
}

type call[T any] struct {
	wg  sync.WaitGroup
	val T
	err error
}

// Do runs fn for key unless another goroutine is already running it.
//
// TODO:
// - If no call is in flight for key, create one, run fn, store the result, wake
//   waiters, and remove the entry from the map before returning.
// - If a call is already in flight for key, wait for it and return the shared
//   result.
// - Return shared=true only for callers that did not execute fn themselves.
// - Make the zero value of Group usable.
func (g *Group[T]) Do(key string, fn func() (T, error)) (v T, err error, shared bool) {
	panic("TODO: implement Do")
}
