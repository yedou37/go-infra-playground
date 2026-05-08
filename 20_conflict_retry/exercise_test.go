package conflictretry

import (
	"context"
	"errors"
	"testing"
)

func TestRetryUpdateRetriesConflicts(t *testing.T) {
	getCalls := 0
	updateCalls := 0

	got, err := RetryUpdate(
		context.Background(),
		3,
		func(ctx context.Context) (Object, error) {
			getCalls++
			return Object{Version: getCalls, Value: "old"}, nil
		},
		func(ctx context.Context, obj Object) (Object, error) {
			updateCalls++
			if updateCalls == 1 {
				return Object{}, ErrConflict
			}
			return Object{Version: obj.Version + 1, Value: obj.Value}, nil
		},
		func(obj *Object) {
			obj.Value = "new"
		},
	)
	if err != nil {
		t.Fatalf("RetryUpdate returned error: %v", err)
	}
	if getCalls != 2 || updateCalls != 2 {
		t.Fatalf("calls = get:%d update:%d, want 2/2", getCalls, updateCalls)
	}
	if got.Value != "new" {
		t.Fatalf("got %+v, want updated value", got)
	}
}

func TestRetryUpdateStopsOnNonConflictError(t *testing.T) {
	errBoom := errors.New("boom")

	_, err := RetryUpdate(
		context.Background(),
		3,
		func(ctx context.Context) (Object, error) {
			return Object{Version: 1, Value: "old"}, nil
		},
		func(ctx context.Context, obj Object) (Object, error) {
			return Object{}, errBoom
		},
		func(obj *Object) {
			obj.Value = "new"
		},
	)
	if !errors.Is(err, errBoom) {
		t.Fatalf("err = %v, want errBoom", err)
	}
}

func TestRetryUpdateTreatsNonPositiveAttemptsAsOne(t *testing.T) {
	updateCalls := 0

	_, err := RetryUpdate(
		context.Background(),
		0,
		func(ctx context.Context) (Object, error) {
			return Object{Version: 1, Value: "old"}, nil
		},
		func(ctx context.Context, obj Object) (Object, error) {
			updateCalls++
			return Object{}, ErrConflict
		},
		func(obj *Object) {},
	)
	if !errors.Is(err, ErrConflict) {
		t.Fatalf("err = %v, want ErrConflict", err)
	}
	if updateCalls != 1 {
		t.Fatalf("updateCalls = %d, want 1", updateCalls)
	}
}
