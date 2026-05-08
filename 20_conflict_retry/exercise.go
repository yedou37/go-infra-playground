package conflictretry

import (
	"context"
	"errors"
)

// ErrConflict reports an optimistic concurrency conflict.
var ErrConflict = errors.New("conflict")

// Object is a tiny versioned resource.
type Object struct {
	Version int
	Value   string
}

// RetryUpdate loads the latest object, mutates it, and retries updates on
// conflict.
//
// This is one of the most common controller/APIServer patterns:
// read the newest version, change a copy, try to update, and retry only if the
// server reports a conflict.
//
// TODO:
// - Treat maxAttempts <= 0 as 1.
// - On each attempt, call get to load the latest object.
// - Mutate a local copy using mutate.
// - Call update with the mutated object.
// - Retry only when update returns ErrConflict.
// - Return immediately on any non-conflict error.
// - Respect ctx cancellation.
func RetryUpdate(
	ctx context.Context,
	maxAttempts int,
	get func(context.Context) (Object, error),
	update func(context.Context, Object) (Object, error),
	mutate func(*Object),
) (Object, error) {
	panic("TODO: implement RetryUpdate")
}
