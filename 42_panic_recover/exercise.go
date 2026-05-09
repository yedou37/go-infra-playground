// Package safego is about the "safe goroutine" pattern.
//
// Why this matters:
//
//   In Go, an unrecovered panic in ANY goroutine crashes the WHOLE
//   process, no matter where it was started. This is one of the easiest
//   ways to take down a production server: a single buggy callback runs
//   in `go fn()` and the entire pod restarts.
//
//   So infra code (k8s, etcd, every serious server) wraps long-lived
//   background goroutines in a small helper that defers a recover() and
//   reports the panic somewhere safe (logger, metric, error channel).
//   client-go's `runtime.HandleCrash` is exactly this idea.
//
// You will implement two flavors:
//
//   - Go(fn): fire-and-forget, panics get reported to a global handler.
//   - GoSafe(fn): run fn in a goroutine and return a channel that yields
//     either a nil error (clean exit) or an error describing the panic.
package safego

// PanicHandler is invoked from within the recover() of a Go() goroutine.
// It receives whatever was passed to panic(). If nil, panics are silently
// swallowed.
//
// Typical production wiring sets this to a function that logs the panic
// + stack trace and bumps a metric. For this exercise it is enough to
// store the panic value somewhere observable.
var PanicHandler func(v any)

// Go starts fn in a new goroutine. If fn panics, the panic must NOT
// propagate to the runtime; instead, PanicHandler is invoked (if set).
//
// Required behavior:
//   - The new goroutine must defer recover().
//   - If recover() returns non-nil, PanicHandler(v) must be called
//     (only if PanicHandler != nil).
//   - Go must not block the caller waiting for fn to finish.
func Go(fn func()) {
	panic("TODO: implement Go")
}

// GoSafe runs fn in a new goroutine and returns a channel of size 1
// that will receive exactly one value:
//
//   - nil if fn returned normally
//   - a non-nil error of the form "recovered from panic: <value>" if fn
//     panicked
//
// The channel must be closed after the value is sent so that the caller
// can use range/ok semantics to detect completion.
func GoSafe(fn func()) <-chan error {
	panic("TODO: implement GoSafe")
}
