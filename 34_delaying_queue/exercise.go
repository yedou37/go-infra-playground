// Package delayingqueue is a simplified version of client-go's
// workqueue.DelayingInterface.
//
// Design notes:
//
// In Kubernetes controllers, you very often want to say "process this key,
// but not before t". For example, after a transient API error, you want to
// requeue the same object after a backoff. The data structure that powers
// this in client-go is a delaying work queue backed by a min-heap of
// (readyAt, item) pairs plus a regular FIFO of "ready now" items.
//
// To keep this exercise deterministic and easy to test, the API takes an
// explicit `now time.Time` instead of reading the wall clock. This is the
// same pattern client-go uses with its `clock.Clock` interface, and is the
// reason fake clocks exist at all.
//
// Semantics to implement:
//
//   - AddAfter(item, delay, now): schedule item to become ready at now+delay.
//     If delay <= 0 the item should be ready immediately.
//   - Get(now): pop the next item whose readyAt <= now. If none is ready,
//     return ok=false and a `wait` duration equal to the time until the
//     earliest pending item (or 0 if the queue is empty / shut down).
//   - Items with the same key may appear multiple times; this exercise does
//     NOT require dedup. (Real workqueues do dedup; we keep it simple.)
//   - After Shutdown, Get must return ok=false and wait=0 forever, even if
//     items are still scheduled.
package delayingqueue

import "time"

// DelayingQueue is a min-heap-ordered queue of items that become ready at
// some point in the future.
type DelayingQueue struct {
	// TODO: choose a representation. A min-heap by readyAt is the canonical
	// choice; a sorted slice is also acceptable for an exercise.
}

// New returns an empty queue.
func New() *DelayingQueue {
	panic("TODO: implement New")
}

// AddAfter schedules item to become ready at now.Add(delay).
// If delay is <= 0 the item is ready immediately.
// AddAfter on a shut-down queue must be a no-op.
func (q *DelayingQueue) AddAfter(item string, delay time.Duration, now time.Time) {
	panic("TODO: implement AddAfter")
}

// Get returns the next item whose readyAt <= now.
//
// Contract:
//   - If an item is ready: returns (item, true, 0).
//   - If the queue is empty or shut down: returns ("", false, 0).
//   - Otherwise: returns ("", false, durationUntilNextReady).
//     The caller is expected to sleep for `wait` (or until something else
//     wakes it up) and try Get again.
func (q *DelayingQueue) Get(now time.Time) (item string, ok bool, wait time.Duration) {
	panic("TODO: implement Get")
}

// Len returns the number of items currently scheduled, ready or not.
func (q *DelayingQueue) Len() int {
	panic("TODO: implement Len")
}

// Shutdown marks the queue as closed. Subsequent AddAfter calls are
// no-ops and Get must return ok=false, wait=0.
func (q *DelayingQueue) Shutdown() {
	panic("TODO: implement Shutdown")
}

// ShuttingDown reports whether Shutdown was called.
func (q *DelayingQueue) ShuttingDown() bool {
	panic("TODO: implement ShuttingDown")
}
