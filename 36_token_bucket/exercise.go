// Package tokenbucket implements a deterministic token bucket rate limiter.
//
// Design notes:
//
// The token bucket is the rate limiter used (in various forms) by:
//   - golang.org/x/time/rate
//   - client-go's rest.Config QPS / Burst
//   - many ingress / API-gateway components
//
// The intuition:
//   - the bucket holds at most `burst` tokens
//   - tokens refill at `rate` tokens per second
//   - each call consumes N tokens; if there aren't enough, it is denied
//
// Implementation tip: do NOT spin a background goroutine to add tokens.
// The standard, idiomatic implementation refills *lazily* on each call,
// computing how many tokens accumulated since the last update. That makes
// the limiter cheap, lock-friendly, and trivially testable with an
// explicit `now`.
//
// Semantics to implement:
//
//   - New(rate float64, burst int): rate is tokens per second; burst is the
//     maximum bucket size. The bucket starts FULL (= burst tokens).
//   - Allow(now): equivalent to AllowN(now, 1).
//   - AllowN(now, n): if at least n tokens are available at `now`, consume
//     them and return true; otherwise return false (consume nothing).
//   - Tokens never exceed burst, even if many seconds elapse without calls.
//   - n <= 0 is treated as a successful no-op (return true, consume nothing).
//   - n > burst is impossible to satisfy: return false.
package tokenbucket

import "time"

// Limiter is a token-bucket rate limiter.
type Limiter struct {
	// TODO: store rate, burst, current tokens, last refill timestamp.
}

// New returns a Limiter that starts FULL (burst tokens available).
func New(rate float64, burst int) *Limiter {
	panic("TODO: implement New")
}

// Allow is shorthand for AllowN(now, 1).
func (l *Limiter) Allow(now time.Time) bool {
	panic("TODO: implement Allow")
}

// AllowN reports whether n tokens can be consumed at `now`.
// On success, consume them. On failure, do not consume anything.
func (l *Limiter) AllowN(now time.Time, n int) bool {
	panic("TODO: implement AllowN")
}
