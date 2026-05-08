package queue

// Queue is a small generic FIFO queue.
//
// This exercise is intentionally "infrastructure flavored":
// you need to think about ownership, zero values, and whether callers can
// accidentally observe internal mutations.
type Queue[T any] struct {
	items []T
	head  int
}

// New returns an empty queue with optional initial capacity.
//
// TODO:
// - If capacity is negative, treat it as zero.
func New[T any](capacity int) *Queue[T] {
	panic("TODO: implement New")
}

// Len returns the number of readable items.
//
// TODO:
// - Do not count items that have already been popped.
func (q *Queue[T]) Len() int {
	panic("TODO: implement Len")
}

// Push appends one value to the tail.
func (q *Queue[T]) Push(v T) {
	panic("TODO: implement Push")
}

// Pop removes and returns the oldest value.
//
// TODO:
// - Return the zero value of T and false if the queue is empty.
// - Clear the popped slot so large referenced objects can be garbage collected.
func (q *Queue[T]) Pop() (T, bool) {
	panic("TODO: implement Pop")
}

// Snapshot returns a copy of the readable queue contents in FIFO order.
//
// TODO:
// - Return a copy, not an alias into q.items.
func (q *Queue[T]) Snapshot() []T {
	panic("TODO: implement Snapshot")
}
