package contextloop

import (
	"context"
	"errors"
	"time"
)

// ErrChannelClosed reports that a receive-only channel was closed before a
// value arrived.
var ErrChannelClosed = errors.New("channel closed")

// Wait blocks until either ctx is done or ch produces one value.
//
// This is a common building block in request-scoped code and controller code:
// the caller wants one value, but it also wants a cancellation path so the
// goroutine does not wait forever.
//
// TODO:
// - If ctx is done first, return the zero value of T and ctx.Err().
// - If ch is closed before a value arrives, return the zero value of T and
//   ErrChannelClosed.
// - Otherwise return the received value and nil.
func Wait[T any](ctx context.Context, ch <-chan T) (T, error) {
	panic("TODO: implement Wait")
}

// Send sends v to ch unless ctx is done first.
//
// This pattern is useful when a producer should stop promptly if the request or
// background task has already been canceled.
//
// TODO:
// - Return nil on success.
// - Return ctx.Err() if ctx is done before the send completes.
func Send[T any](ctx context.Context, ch chan<- T, v T) error {
	panic("TODO: implement Send")
}

// RunOnTicks calls fn once for every value received from ticks until ctx is
// done.
//
// This mirrors the shape of many Kubernetes background loops:
// a goroutine waits on both "stop" and "periodic trigger" signals.
//
// TODO:
// - Loop until ctx.Done() is closed.
// - Call fn once for each received tick.
// - Return promptly after cancellation.
func RunOnTicks(ctx context.Context, ticks <-chan time.Time, fn func()) {
	panic("TODO: implement RunOnTicks")
}
