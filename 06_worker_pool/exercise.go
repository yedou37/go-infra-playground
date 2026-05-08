package workerpool

import "context"

// Task is a small unit of work.
type Task struct {
	ID    int
	Value int
}

// Result is the processed form of one Task.
type Result struct {
	ID    int
	Value int
}

// Map applies fn to tasks using at most workers goroutines and returns the
// results in the same order as the input slice.
//
// This exercise mirrors a very common infrastructure pattern:
// do bounded parallel work, stop on cancellation or first error, and still
// produce a predictable result shape for callers.
//
// TODO:
// - Treat workers <= 0 as 1.
// - Preserve input order in the returned slice even if tasks finish out of
//   order.
// - Stop early and return the first error from fn.
// - Respect ctx cancellation.
// - Avoid leaking goroutines when returning early.
func Map(
	ctx context.Context,
	workers int,
	tasks []Task,
	fn func(context.Context, Task) (Result, error),
) ([]Result, error) {
	panic("TODO: implement Map")
}
