package contexttree

import (
	"context"
	"time"
)

// RunSubtask creates a child context and runs fn with it.
//
// This exercise focuses on the parent/child context model that appears
// everywhere in industrial Go code.
//
// The mental model is:
// - The parent context describes the lifetime of the larger operation.
// - The child context describes one shorter sub-operation inside it.
// - The child inherits cancellation from the parent.
// - The child may also add a stricter timeout of its own.
//
// In practice this means:
//   - If the parent is already canceled, the subtask should not even start.
//   - If the parent is canceled later, the child should stop too.
//   - Even if the parent stays alive, the child may still stop earlier because
//     its own timeout expires.
//
// Another important rule is resource cleanup:
// context.WithTimeout allocates internal timer-related state, so callers should
// almost always call the returned cancel function. In short-lived helper code,
// the usual pattern is:
//
//	child, cancel := context.WithTimeout(parent, timeout)
//	defer cancel()
//
// That defer is not only for "manual cancellation"; it also releases internal
// resources promptly after the subtask finishes.
//
// This function also highlights a subtle but important semantic point:
// fn returning nil does not automatically mean the overall subtask succeeded.
// If fn returned because the child context was canceled or timed out, then the
// correct result is usually child.Err(), not nil.
//
// This shape is common in request handling, RPC fan-out, storage calls,
// controller helpers, and any code that needs one operation to live inside the
// lifetime budget of another.
//
// TODO:
// - If ctx is already done, return ctx.Err() without calling fn.
// - Create a child context with the provided timeout.
// - Defer the child cancel function.
// - Call fn with the child context.
// - If fn returns nil but the child context is done, return child.Err().
// - Otherwise return fn's result.
func RunSubtask(ctx context.Context, timeout time.Duration, fn func(context.Context) error) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	child, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	err := fn(child)
	if err != nil {
		return err
	}
	if err := child.Err(); err != nil {
		return err
	}
	return nil
}
