package fakeclockretry

import "context"

// Clock abstracts waiting so tests can control time deterministically.
type Clock interface {
	After() <-chan struct{}
}

// RetryWithClock retries fn after waiting on clock.After().
//
// This exercise focuses on a core industrial testing technique:
// inject time as a dependency so retry logic can be tested without real sleeps.
//
// TODO:
// - Treat maxAttempts <= 0 as 1.
// - Call fn at most maxAttempts times.
// - Return nil immediately when fn succeeds.
// - Between failed attempts, wait on clock.After().
// - Respect ctx cancellation while waiting.
// - Return the last error after the final failed attempt.
func RetryWithClock(
	ctx context.Context,
	clock Clock,
	maxAttempts int,
	fn func(context.Context) error,
) error {
	panic("TODO: implement RetryWithClock")
}
