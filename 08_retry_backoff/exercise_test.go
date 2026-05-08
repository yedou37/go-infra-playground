package retrybackoff

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRetrySucceedsAfterBackoff(t *testing.T) {
	var attempts int
	var sleeps []time.Duration

	err := Retry(context.Background(), 4, 10*time.Millisecond, func(ctx context.Context, d time.Duration) error {
		sleeps = append(sleeps, d)
		return nil
	}, func(ctx context.Context) error {
		attempts++
		if attempts < 3 {
			return errors.New("transient")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Retry returned error: %v", err)
	}
	if attempts != 3 {
		t.Fatalf("attempts = %d, want 3", attempts)
	}
	want := []time.Duration{10 * time.Millisecond, 20 * time.Millisecond}
	if len(sleeps) != len(want) {
		t.Fatalf("sleeps = %v, want %v", sleeps, want)
	}
	for i := range want {
		if sleeps[i] != want[i] {
			t.Fatalf("sleeps = %v, want %v", sleeps, want)
		}
	}
}

func TestRetryDoesNotSleepAfterLastAttempt(t *testing.T) {
	var sleeps int

	err := Retry(context.Background(), 2, time.Second, func(ctx context.Context, d time.Duration) error {
		sleeps++
		return nil
	}, func(ctx context.Context) error {
		return errors.New("still failing")
	})
	if err == nil {
		t.Fatal("Retry should return an error")
	}
	if sleeps != 1 {
		t.Fatalf("sleep calls = %d, want 1", sleeps)
	}
	if !errors.Is(err, ErrExhausted) {
		t.Fatalf("err = %v, want ErrExhausted", err)
	}
}

func TestRetryReturnsContextError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	called := 0
	err := Retry(ctx, 3, time.Second, func(ctx context.Context, d time.Duration) error {
		t.Fatal("sleep should not be called after cancellation")
		return nil
	}, func(ctx context.Context) error {
		called++
		return nil
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
	if called != 0 {
		t.Fatalf("fn calls = %d, want 0", called)
	}
}

func TestRetryReturnsSleepError(t *testing.T) {
	errSleep := errors.New("sleep failed")

	err := Retry(context.Background(), 3, time.Second, func(ctx context.Context, d time.Duration) error {
		return errSleep
	}, func(ctx context.Context) error {
		return errors.New("transient")
	})
	if !errors.Is(err, errSleep) {
		t.Fatalf("err = %v, want errSleep", err)
	}
}

func TestRetryTreatsNonPositiveAttemptsAsOne(t *testing.T) {
	calls := 0
	errBoom := errors.New("boom")

	err := Retry(context.Background(), 0, time.Second, func(ctx context.Context, d time.Duration) error {
		t.Fatal("sleep should not be called for a single effective attempt")
		return nil
	}, func(ctx context.Context) error {
		calls++
		return errBoom
	})
	if calls != 1 {
		t.Fatalf("calls = %d, want 1", calls)
	}
	if !errors.Is(err, ErrExhausted) || !errors.Is(err, errBoom) {
		t.Fatalf("err = %v, want wrapped ErrExhausted and errBoom", err)
	}
}
