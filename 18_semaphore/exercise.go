package semaphore

import "context"

// Limiter is a small counting semaphore backed by a buffered channel.
type Limiter struct {
	tokens chan struct{}
}

// New returns a limiter with capacity max.
//
// TODO:
// - Treat max <= 0 as 1.
// - Fill the buffered channel with max initial permits.
func New(max int) *Limiter {
	panic("TODO: implement New")
}

// Acquire waits for one permit or ctx cancellation.
//
// TODO:
// - Return nil after acquiring one permit.
// - Return ctx.Err() if ctx is done first.
func (l *Limiter) Acquire(ctx context.Context) error {
	panic("TODO: implement Acquire")
}

// Release returns one permit to the limiter.
//
// TODO:
// - Return one token to the buffered channel.
// - Assume callers release only after a successful Acquire.
func (l *Limiter) Release() {
	panic("TODO: implement Release")
}
