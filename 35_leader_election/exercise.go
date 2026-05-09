// Package leaderelection is a tiny model of lease-based leader election,
// the same idea used by Kubernetes (k8s.io/client-go/tools/leaderelection)
// and by etcd's `concurrency` package on top of leases.
//
// Design notes:
//
// In a real cluster, leader election is implemented on top of a strongly
// consistent store: a candidate writes a lease object that contains
// (holder, expiry, etc.) using compare-and-swap semantics. Whichever
// candidate's CAS wins becomes the leader. The leader periodically
// renews the lease before it expires; if it fails to renew in time, any
// other candidate may take over once the old lease is observed expired.
//
// This exercise strips that down to a single in-memory Election that
// holds at most one valid lease at a time. Time is passed in explicitly
// (`now time.Time`) so tests don't depend on the wall clock.
//
// Semantics to implement:
//
//   - TryAcquire(holder, now, ttl): if no current valid lease (none ever
//     held, or the previous one has expired at `now`), set the lease to
//     (holder, now+ttl) and return true. Otherwise return false.
//   - Renew(holder, now, ttl): if `holder` is the current valid leader at
//     `now`, extend the lease to now+ttl and return true. Otherwise return
//     false. Renew must NOT let an expired holder silently re-acquire.
//   - Release(holder): if `holder` is the current leader (regardless of
//     expiry, to allow voluntary release), clear the lease and return true.
//   - CurrentHolder(now): returns the current valid leader at `now`.
//     Returns ("", false) if there is no leader or the lease has expired.
package leaderelection

import "time"

// Election holds at most one Lease at a time.
type Election struct {
	// TODO: you decide. Typically a struct holding (holder, expiry).
}

// New returns an Election with no current leader.
func New() *Election {
	panic("TODO: implement New")
}

// TryAcquire tries to become the leader at `now` with the given TTL.
// Returns true on success.
func (e *Election) TryAcquire(holder string, now time.Time, ttl time.Duration) bool {
	panic("TODO: implement TryAcquire")
}

// Renew extends the lease if `holder` is currently the valid leader.
func (e *Election) Renew(holder string, now time.Time, ttl time.Duration) bool {
	panic("TODO: implement Renew")
}

// Release clears the lease if `holder` is the recorded leader.
// This is voluntary step-down: the lease is cleared even if it has not
// yet expired, so other candidates can take over immediately.
func (e *Election) Release(holder string) bool {
	panic("TODO: implement Release")
}

// CurrentHolder reports who the valid leader is at `now`.
func (e *Election) CurrentHolder(now time.Time) (string, bool) {
	panic("TODO: implement CurrentHolder")
}
