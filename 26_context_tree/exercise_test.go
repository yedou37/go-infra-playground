package contexttree

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRunSubtaskSuccess(t *testing.T) {
	err := RunSubtask(context.Background(), time.Second, func(ctx context.Context) error {
		return nil
	})
	if err != nil {
		t.Fatalf("RunSubtask returned error: %v", err)
	}
}

func TestRunSubtaskParentAlreadyCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	called := false
	err := RunSubtask(ctx, time.Second, func(ctx context.Context) error {
		called = true
		return nil
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
	if called {
		t.Fatal("fn should not be called when parent is already canceled")
	}
}

func TestRunSubtaskChildTimeout(t *testing.T) {
	err := RunSubtask(context.Background(), 20*time.Millisecond, func(ctx context.Context) error {
		<-ctx.Done()
		return nil
	})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("err = %v, want context.DeadlineExceeded", err)
	}
}
