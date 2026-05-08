package fakeclockretry

import (
	"context"
	"errors"
	"testing"
)

type manualClock struct {
	ch chan struct{}
}

func newManualClock() *manualClock {
	return &manualClock{ch: make(chan struct{}, 10)}
}

func (c *manualClock) After() <-chan struct{} {
	return c.ch
}

func (c *manualClock) Tick() {
	c.ch <- struct{}{}
}

func TestRetryWithClockRetriesAfterTick(t *testing.T) {
	clock := newManualClock()
	attempts := 0
	done := make(chan error, 1)

	go func() {
		done <- RetryWithClock(context.Background(), clock, 2, func(ctx context.Context) error {
			attempts++
			if attempts == 1 {
				return errors.New("transient")
			}
			return nil
		})
	}()

	if attempts != 1 {
		t.Fatalf("attempts = %d, want 1 before first tick", attempts)
	}
	clock.Tick()

	if err := <-done; err != nil {
		t.Fatalf("RetryWithClock returned error: %v", err)
	}
	if attempts != 2 {
		t.Fatalf("attempts = %d, want 2", attempts)
	}
}

func TestRetryWithClockReturnsLastError(t *testing.T) {
	clock := newManualClock()
	errBoom := errors.New("boom")
	done := make(chan error, 1)

	go func() {
		done <- RetryWithClock(context.Background(), clock, 2, func(ctx context.Context) error {
			return errBoom
		})
	}()

	clock.Tick()
	err := <-done
	if !errors.Is(err, errBoom) {
		t.Fatalf("err = %v, want errBoom", err)
	}
}

func TestRetryWithClockRespectsCanceledContext(t *testing.T) {
	clock := newManualClock()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := RetryWithClock(ctx, clock, 2, func(ctx context.Context) error {
		t.Fatal("fn should not be called after cancellation")
		return nil
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
}
