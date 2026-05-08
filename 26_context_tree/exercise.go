package contexttree

import (
	"context"
	"time"
)

// RunSubtask creates a child context and runs fn with it.
//
// This exercise focuses on parent/child context behavior:
// short-lived sub-operations should inherit cancellation, possibly add a
// timeout, and always release resources via cancel.
//
// TODO:
// - If ctx is already done, return ctx.Err() without calling fn.
// - Create a child context with the provided timeout.
// - Defer the child cancel function.
// - Call fn with the child context.
// - If fn returns nil but the child context is done, return child.Err().
// - Otherwise return fn's result.
func RunSubtask(ctx context.Context, timeout time.Duration, fn func(context.Context) error) error {
	panic("TODO: implement RunSubtask")
}
