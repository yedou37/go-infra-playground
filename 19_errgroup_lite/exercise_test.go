package errgrouplite

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRunAllSuccess(t *testing.T) {
	var a, b bool

	err := Run(context.Background(),
		func(ctx context.Context) error {
			a = true
			return nil
		},
		func(ctx context.Context) error {
			b = true
			return nil
		},
	)
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if !a || !b {
		t.Fatalf("expected both functions to run, got a=%v b=%v", a, b)
	}
}

func TestRunCancelsSiblingsOnFirstError(t *testing.T) {
	errBoom := errors.New("boom")
	stopped := make(chan struct{})

	err := Run(context.Background(),
		func(ctx context.Context) error {
			return errBoom
		},
		func(ctx context.Context) error {
			defer close(stopped)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(200 * time.Millisecond):
				t.Fatal("sibling should have been canceled early")
				return nil
			}
		},
	)
	if !errors.Is(err, errBoom) {
		t.Fatalf("err = %v, want errBoom", err)
	}

	select {
	case <-stopped:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Run did not wait for siblings to stop")
	}
}

func TestRunEmpty(t *testing.T) {
	if err := Run(context.Background()); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
}
