package semaphore

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestAcquireRelease(t *testing.T) {
	l := New(1)

	if err := l.Acquire(context.Background()); err != nil {
		t.Fatalf("Acquire returned error: %v", err)
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		if err := l.Acquire(context.Background()); err != nil {
			t.Errorf("second Acquire returned error: %v", err)
			return
		}
		l.Release()
	}()

	select {
	case <-done:
		t.Fatal("second Acquire should block until Release")
	case <-time.After(20 * time.Millisecond):
	}

	l.Release()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("second Acquire did not proceed after Release")
	}
}

func TestAcquireRespectsCanceledContext(t *testing.T) {
	l := New(1)
	if err := l.Acquire(context.Background()); err != nil {
		t.Fatalf("Acquire returned error: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := l.Acquire(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
}

func TestNewTreatsNonPositiveCapacityAsOne(t *testing.T) {
	l := New(0)
	if err := l.Acquire(context.Background()); err != nil {
		t.Fatalf("Acquire returned error: %v", err)
	}
	l.Release()
}
