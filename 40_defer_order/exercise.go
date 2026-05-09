// Package deferorder is about Go's defer mechanism, which is one of the
// few language features that has no direct C++/Java analog and is the
// canonical way to write robust resource cleanup in Go.
//
// Things this exercise pins down:
//
//   - defer runs in LIFO order (last-deferred runs first).
//   - defer arguments are evaluated AT THE defer STATEMENT, not when the
//     deferred function actually runs. Closures however capture the
//     variable, not the value at defer time.
//   - With NAMED return values, a deferred function can MUTATE the
//     return value before the caller sees it. This is the standard way
//     to do "translate panic to error" or "wrap error on the way out".
//   - recover() only works inside a deferred function, and only catches
//     a panic from the same goroutine.
//
// You will implement four small functions that exercise each of these.
package deferorder

// DeferOrder appends the values 1, 2, 3 in the order they are produced
// by a sequence of defers. The function returns a []int.
//
// Required behavior:
//   - Use exactly 3 deferred calls, each one appending a single int.
//   - Defer them in the order 1, 2, 3.
//   - The returned slice must reflect ACTUAL execution order (LIFO).
//   - i.e. you should observe [3, 2, 1] in the returned slice.
//
// The point of this is to make you write the code that proves you know
// LIFO, not just to know it intellectually.
func DeferOrder() []int {
	panic("TODO: implement DeferOrder using 3 defers")
}

// CapturedAtDeferTime demonstrates that defer arguments are evaluated
// at the moment of the defer statement.
//
// Required behavior:
//   - Inside the function, define x := 1.
//   - defer a call that "records" x by VALUE (e.g. by passing x as an
//     argument to a closure that appends it).
//   - Then set x = 2 and defer another call that records x BY CLOSURE
//     CAPTURE (no argument; the closure should reference the outer x).
//   - Then set x = 3 and return.
//
// Returned slice (in the order the defers actually fire, LIFO) must be:
//   [3, 1]
//
// Reading: the closure-capturing defer (deferred LAST) fires first and
// sees x == 3 because it captures the variable; the by-value defer
// (deferred FIRST) fires second and sees the old snapshot, x == 1.
func CapturedAtDeferTime() []int {
	panic("TODO: implement CapturedAtDeferTime")
}

// WrapErrorOnExit is the canonical "named return + defer" pattern.
//
// Required behavior:
//   - The function takes a bool `fail`.
//   - It uses a NAMED return value `err error`.
//   - It defers a function that, if err != nil, REPLACES it with a new
//     error whose Error() string is exactly: "wrapped: " + original.Error().
//   - In the body: if fail is true, set err = errors.New("boom") and
//     return; otherwise return nil.
//
// This is exactly how production code does "always wrap errors leaving
// this function" without sprinkling the wrap at every return site.
func WrapErrorOnExit(fail bool) (err error) {
	panic("TODO: implement WrapErrorOnExit")
}

// SafeCall runs fn and converts a panic into an error.
//
// Required behavior:
//   - If fn returns normally, SafeCall returns nil.
//   - If fn panics with value v, SafeCall must return a non-nil error
//     whose Error() string contains "panic:" followed by something that
//     identifies v (fmt.Sprint(v) is fine).
//   - SafeCall must NOT propagate the panic to the caller.
//   - Use defer + recover().
func SafeCall(fn func()) (err error) {
	panic("TODO: implement SafeCall")
}
