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
// This exercise is a good place to memorize the different jobs of channel and
// context:
//
// Channel:
//   - A channel is the data or signal path between goroutines.
//   - `<-ch` means receive one value from ch.
//   - `v, ok := <-ch` means receive one value and also learn whether the channel
//     is still open.
//   - If `ok == false`, the channel was closed before a value could be received.
//
// Context:
//   - A context does not carry the business data itself.
//   - `ctx.Done()` is a channel that is closed when cancellation or timeout
//     happens, so it is used inside `select`.
//   - `ctx.Err()` returns the reason after the context is done, usually
//     `context.Canceled` or `context.DeadlineExceeded`.
//   - A common mental split is:
//     wait with `Done`, explain with `Err`.
//
// The standard "wait for result or cancellation" template is:
//
//	if err := ctx.Err(); err != nil {
//		return zero, err
//	}
//	select {
//	case <-ctx.Done():
//		return zero, ctx.Err()
//	case v, ok := <-ch:
//		if !ok {
//			return zero, ErrChannelClosed
//		}
//		return v, nil
//	}
//
// This shape is extremely common in industrial Go:
// - wait for RPC result or timeout
// - wait for worker result or shutdown
// - wait for queue/event data or stop signal
//
// The key design idea is that we are modeling two independent exit paths:
// - success path: a value arrives from ch
// - stop path: the caller no longer wants to wait
//
// The extra `ok` check matters because "channel closed" is not the same thing
// as "context canceled". In real systems, distinguishing those two cases often
// makes debugging much easier.
//
// TODO:
//   - If ctx is done first, return the zero value of T and ctx.Err().
//   - If ch is closed before a value arrives, return the zero value of T and
//     ErrChannelClosed.
//   - Otherwise return the received value and nil.
func Wait[T any](ctx context.Context, ch <-chan T) (T, error) {
	if err := ctx.Err(); err != nil {
		var zero T
		return zero, err
	}
	select {
	case <-ctx.Done():
		var zero T
		return zero, ctx.Err()
	case v, ok := <-ch:
		if !ok {
			var zero T
			return zero, ErrChannelClosed
		}
		return v, nil
	}
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
	if err := ctx.Err(); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	case ch <- v:
		return nil
	}

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
	if err := ctx.Err(); err != nil {
		return
	}
	for {
		select {
		case <-ctx.Done():
			return
		case _, ok := <-ticks:
			if !ok {
				return
			}
			fn()
		}
	}
}
