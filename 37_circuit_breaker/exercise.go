// Package circuitbreaker is a tiny three-state circuit breaker, the same
// shape used by client-side libraries to avoid hammering an unhealthy
// dependency (etcd clients, kube-apiserver clients, sidecars, etc.).
//
// Design notes:
//
// The breaker is a small state machine with three states:
//
//   - Closed: calls go through. Consecutive failures are counted.
//     When the count reaches `failureThreshold`, the breaker opens.
//   - Open: calls are short-circuited and immediately fail with ErrOpen
//     without invoking fn. After `openDuration` has elapsed since the
//     breaker opened, it transitions to HalfOpen on the next Call.
//   - HalfOpen: a single trial call is allowed through. If it succeeds,
//     the breaker closes and the failure counter resets. If it fails,
//     the breaker re-opens and the cooldown timer restarts.
//
// Two important details that real implementations get wrong:
//
//  1. The failure counter must reset on success, not just on close. A
//     single success in Closed state breaks the streak.
//  2. While Open, ErrOpen must be returned without calling fn. The point
//     of the breaker is to avoid load on the downstream.
//
// Like the other exercises in this repo, time is passed in explicitly
// so tests don't depend on the wall clock.
package circuitbreaker

import (
	"errors"
	"time"
)

// ErrOpen is returned by Call when the breaker is Open and the call is
// short-circuited.
var ErrOpen = errors.New("circuitbreaker: open")

// State of the breaker.
type State int

const (
	StateClosed State = iota
	StateOpen
	StateHalfOpen
)

// Breaker is a state machine; not safe for concurrent use unless you add
// your own locking. Real production breakers usually do.
type Breaker struct {
	// TODO: store failureThreshold, openDuration, state, failure count,
	// and the time at which the breaker opened.
}

// New returns a Breaker that starts in Closed state.
func New(failureThreshold int, openDuration time.Duration) *Breaker {
	panic("TODO: implement New")
}

// State reports the current state at `now`. Calling State should be
// pure: it may compute the effective state (e.g. Open -> HalfOpen after
// cooldown) but should not consume the half-open trial slot.
func (b *Breaker) State(now time.Time) State {
	panic("TODO: implement State")
}

// Call invokes fn unless the breaker is currently Open at `now`, in which
// case it returns ErrOpen immediately without calling fn.
//
// State transitions on Call:
//
//   - Closed + success         -> Closed (failure count resets)
//   - Closed + failure         -> Closed; if streak >= threshold, -> Open
//   - Open (still cooling)     -> ErrOpen, fn NOT called
//   - Open (cooldown elapsed)  -> HalfOpen for this trial call
//   - HalfOpen + success       -> Closed
//   - HalfOpen + failure       -> Open (restart cooldown)
func (b *Breaker) Call(now time.Time, fn func() error) error {
	panic("TODO: implement Call")
}
