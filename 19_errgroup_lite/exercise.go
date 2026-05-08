package errgrouplite

import "context"

// Run executes fns concurrently with a child context.
//
// This mirrors the most common errgroup usage shape in infrastructure code:
// start sibling tasks, cancel the rest on the first error, and wait for every
// goroutine to exit before returning.
//
// TODO:
// - Return nil if fns is empty.
// - Create a child context from ctx.
// - Start all functions concurrently.
// - On the first non-nil error, cancel the child context.
// - Wait for all functions to exit before returning.
// - Return the first non-nil error, or nil if all succeed.
func Run(ctx context.Context, fns ...func(context.Context) error) error {
	panic("TODO: implement Run")
}
