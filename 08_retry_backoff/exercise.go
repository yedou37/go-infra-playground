package retrybackoff

import (
	"context"
	"errors"
	"time"
)

// ErrExhausted reports that all retry attempts were used.
var ErrExhausted = errors.New("retry attempts exhausted")

// Sleeper abstracts sleeping so tests can verify backoff decisions without
// waiting on real time.
type Sleeper func(context.Context, time.Duration) error

// Retry calls fn until it succeeds, ctx is done, or maxAttempts is exhausted.
//
// This mirrors a common infrastructure pattern:
// retry transient work with bounded attempts and backoff, but stop promptly on
// cancellation.
//
// TODO:
// - Treat maxAttempts <= 0 as 1.
// - Call fn at most maxAttempts times.
// - Return nil immediately if fn succeeds.
// - Between failed attempts, sleep using exponential backoff:
//   baseDelay, 2*baseDelay, 4*baseDelay, ...
// - Do not sleep after the last attempt.
// - If ctx is done, return ctx.Err().
// - If sleeping returns an error, return that error.
// - If all attempts fail, return an error that wraps both ErrExhausted and the
//   last error from fn.
func Retry(
	ctx context.Context,
	maxAttempts int,
	baseDelay time.Duration,
	sleep Sleeper,
	fn func(context.Context) error,
) error {
	panic("TODO: implement Retry")
}
