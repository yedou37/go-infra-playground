// Package gracefulshutdown is about correctly stopping an http.Server,
// the same shape every production Go service needs.
//
// What this exercise pins down:
//
//   - http.Server.ListenAndServe blocks until the server is stopped.
//     When Shutdown() is called, ListenAndServe returns http.ErrServerClosed.
//     This sentinel must be treated as a CLEAN exit, not a real failure.
//   - http.Server.Shutdown(ctx) tries to drain in-flight requests up to
//     the deadline carried by ctx, then forcibly stops accepting new ones.
//   - Production code typically:
//       1. Starts the server in a goroutine
//       2. Blocks the main goroutine on a "stop" signal (SIGTERM in real
//          life; in this exercise: a stopCh you control)
//       3. Calls Shutdown with a bounded-timeout context
//       4. Waits for the serving goroutine to return
//
// You will implement that lifecycle as a single Run helper.
package gracefulshutdown

import (
	"context"
	"net/http"
	"time"
)

// Run starts srv, blocks until stopCh is closed, then calls
// srv.Shutdown with a context that uses `drainTimeout`. It returns:
//
//   - nil if everything went cleanly (server returned ErrServerClosed
//     and Shutdown returned nil).
//   - the first non-clean error otherwise (either from ListenAndServe
//     or from Shutdown).
//
// Required behavior:
//   - srv.ListenAndServe() must run in its own goroutine.
//   - http.ErrServerClosed from ListenAndServe is NOT an error here.
//   - When stopCh is closed, build ctx, _ := context.WithTimeout(...)
//     with drainTimeout and call srv.Shutdown(ctx).
//   - Wait until the serving goroutine has actually returned BEFORE
//     returning from Run. This is what callers rely on to know it's
//     safe to exit the process / release resources.
func Run(srv *http.Server, stopCh <-chan struct{}, drainTimeout time.Duration) error {
	panic("TODO: implement Run")
}

// Helper used in tests to assert that ctx with a deadline was used.
// Not part of the exercise per se, but exported so the test file can
// build a server with a known handler.
func _ctxIsDeadlineExceeded(err error) bool {
	return err == context.DeadlineExceeded
}
