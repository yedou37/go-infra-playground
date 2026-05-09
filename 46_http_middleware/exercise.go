// Package middleware is about the http.Handler middleware pattern, the
// idiomatic way to compose cross-cutting concerns (logging, auth,
// timeouts, recovery, request ID, ...) on top of net/http.
//
// What this exercise pins down:
//
//   - http.Handler is just an interface with one method: ServeHTTP.
//   - http.HandlerFunc is a function type that satisfies http.Handler;
//     it is the canonical "function-as-interface" adapter in std lib.
//   - A Middleware is `func(http.Handler) http.Handler`.
//   - Chain(m1, m2, m3)(h) should produce a handler whose request flow
//     is m1 -> m2 -> m3 -> h, and whose response flow is the reverse.
//   - Recovery middleware MUST defer-recover BEFORE writing the response,
//     so that a panic in the handler turns into a 500 response instead
//     of crashing the process and leaving the client hanging.
//
// You will implement the chaining helper plus three middlewares.
package middleware

import (
	"net/http"
)

// Middleware wraps an http.Handler with extra behavior.
type Middleware func(http.Handler) http.Handler

// Chain composes middlewares so that the FIRST one in the list is the
// OUTERMOST wrapper.
//
// Example:
//   Chain(Logger, Auth)(handler)
// must call Logger first (which then calls Auth, which then calls
// handler).
//
// TODO:
//   - Iterate the list in reverse order, wrapping the result so that
//     mws[0] ends up as the outermost layer.
func Chain(mws ...Middleware) Middleware {
	panic("TODO: implement Chain")
}

// LogMiddleware records the incoming request method+path into `sink`
// before calling the next handler.
//
// TODO:
//   - On every request, append fmt.Sprintf("%s %s", r.Method, r.URL.Path)
//     to *sink.
//   - Then call next.ServeHTTP(w, r).
func LogMiddleware(sink *[]string) Middleware {
	panic("TODO: implement LogMiddleware")
}

// AuthMiddleware short-circuits any request that does not carry the
// header "X-Token: <expectedToken>".
//
// TODO:
//   - If the header is missing or wrong: respond with 401 and DO NOT
//     call next.
//   - Otherwise call next.ServeHTTP(w, r).
func AuthMiddleware(expectedToken string) Middleware {
	panic("TODO: implement AuthMiddleware")
}

// RecoverMiddleware turns a panic in the downstream handler into a
// 500 response. It must NOT propagate the panic.
//
// TODO:
//   - In the wrapper handler, defer a function that recovers and, on
//     non-nil recover(), writes http.StatusInternalServerError.
//   - Then call next.ServeHTTP(w, r).
func RecoverMiddleware() Middleware {
	panic("TODO: implement RecoverMiddleware")
}
