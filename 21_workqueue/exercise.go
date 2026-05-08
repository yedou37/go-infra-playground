package workqueue

import (
	"context"
	"errors"
	"sync"
)

// ErrClosed reports that the queue has been shut down and drained.
var ErrClosed = errors.New("queue closed")

// Queue is a tiny deduplicating key queue.
type Queue struct {
	ch     chan string
	mu     sync.Mutex
	in     map[string]bool
	closed bool
}

// New returns a queue with the provided channel capacity.
//
// TODO:
// - Treat capacity <= 0 as 1.
// - Initialize the dedup map.
func New(capacity int) *Queue {
	panic("TODO: implement New")
}

// Add enqueues key unless it is already pending or the queue is shut down.
//
// TODO:
// - Return false if the queue is shut down.
// - Deduplicate pending keys.
// - Send the key once into the channel when it becomes pending.
func (q *Queue) Add(key string) bool {
	panic("TODO: implement Add")
}

// Get returns the next key, or an error on shutdown/cancellation.
//
// TODO:
// - Return key, nil when one is available.
// - Return ctx.Err() if ctx is done first.
// - Return ErrClosed after shutdown once the queue channel is drained.
func (q *Queue) Get(ctx context.Context) (string, error) {
	panic("TODO: implement Get")
}

// Done marks key as no longer pending so it can be re-enqueued.
func (q *Queue) Done(key string) {
	panic("TODO: implement Done")
}

// ShutDown closes the queue for new work.
//
// TODO:
// - Make repeated calls safe.
// - Close the channel exactly once.
func (q *Queue) ShutDown() {
	panic("TODO: implement ShutDown")
}
