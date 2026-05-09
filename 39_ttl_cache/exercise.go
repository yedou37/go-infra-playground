// Package ttlcache is a TTL (time-to-live) cache, the same shape as
// client-go's `cache.NewExpirationStore` and the per-entry expiry caches
// in many controllers (e.g. token caches, DNS caches, leader-info caches).
//
// Design notes:
//
// A TTL cache stores (key -> value) along with an absolute expiry time.
// Reads must skip entries whose expiry has passed at `now`. Two design
// decisions matter:
//
//  1. Lazy vs eager expiration.
//     The simplest and most idiomatic choice is LAZY: never delete an
//     entry on a timer; just skip it on Get. This is what we do here.
//     Provide an explicit GC(now) method for callers that want to reclaim
//     memory deterministically (and for tests).
//
//  2. Time injection.
//     Just like the other exercises in this repo, time is passed in as
//     `now time.Time`. This makes the cache trivially testable without
//     spinning wall-clock timers.
//
// Concurrency:
//
//   The cache is intended to be safe for concurrent use. Use a sync.Mutex
//   (or RWMutex) internally. Do NOT return references to internal state
//   from Get; values of type V are returned by value, which is fine for
//   primitives, structs, slices/maps. (Caller-side aliasing of slice/map
//   values is the caller's problem; that's a separate ownership exercise.)
package ttlcache

import "time"

// Cache is a generic TTL cache.
type Cache[K comparable, V any] struct {
	// TODO: store entries map[K]entry[V] and a sync.Mutex.
}

// New returns an empty cache.
func New[K comparable, V any]() *Cache[K, V] {
	panic("TODO: implement New")
}

// Set inserts or overwrites (key -> val) with the given absolute expiry.
// expiresAt being in the past means the entry is born already expired; it
// should NOT be observable by Get and should be reaped by GC.
func (c *Cache[K, V]) Set(key K, val V, expiresAt time.Time) {
	panic("TODO: implement Set")
}

// Get returns the value for key, or (zero, false) if missing or expired
// at `now`.
func (c *Cache[K, V]) Get(key K, now time.Time) (V, bool) {
	panic("TODO: implement Get")
}

// Delete removes key. No-op if key is absent.
func (c *Cache[K, V]) Delete(key K) {
	panic("TODO: implement Delete")
}

// Len returns the number of entries that are NOT yet expired at `now`.
// (Already-expired entries are invisible even before GC runs.)
func (c *Cache[K, V]) Len(now time.Time) int {
	panic("TODO: implement Len")
}

// GC reaps every entry whose expiry has passed at `now` and returns the
// number of entries removed.
func (c *Cache[K, V]) GC(now time.Time) int {
	panic("TODO: implement GC")
}
