package circuitbreaker

import (
	"errors"
	"testing"
	"time"
)

var errBoom = errors.New("boom")

func TestClosedAllowsCalls(t *testing.T) {
	b := New(3, time.Second)
	t0 := time.Unix(1000, 0)
	for i := 0; i < 5; i++ {
		if err := b.Call(t0, func() error { return nil }); err != nil {
			t.Fatalf("call %d: unexpected err %v", i, err)
		}
	}
	if b.State(t0) != StateClosed {
		t.Fatalf("breaker should remain Closed on all-success")
	}
}

func TestOpensAfterConsecutiveFailures(t *testing.T) {
	b := New(2, time.Second)
	t0 := time.Unix(1000, 0)
	if err := b.Call(t0, func() error { return errBoom }); !errors.Is(err, errBoom) {
		t.Fatalf("want errBoom, got %v", err)
	}
	if err := b.Call(t0, func() error { return errBoom }); !errors.Is(err, errBoom) {
		t.Fatalf("want errBoom, got %v", err)
	}
	if b.State(t0) != StateOpen {
		t.Fatalf("breaker should be Open after 2 failures")
	}
}

func TestOpenShortCircuitsWithoutCallingFn(t *testing.T) {
	b := New(1, time.Second)
	t0 := time.Unix(1000, 0)
	_ = b.Call(t0, func() error { return errBoom })
	if b.State(t0) != StateOpen {
		t.Fatalf("expected Open after 1 failure (threshold=1)")
	}

	called := false
	err := b.Call(t0.Add(100*time.Millisecond), func() error {
		called = true
		return nil
	})
	if !errors.Is(err, ErrOpen) {
		t.Fatalf("want ErrOpen, got %v", err)
	}
	if called {
		t.Fatalf("fn must NOT be invoked while Open")
	}
}

func TestSuccessInClosedResetsFailureStreak(t *testing.T) {
	b := New(3, time.Second)
	t0 := time.Unix(1000, 0)
	b.Call(t0, func() error { return errBoom })
	b.Call(t0, func() error { return errBoom })
	b.Call(t0, func() error { return nil })
	b.Call(t0, func() error { return errBoom })
	b.Call(t0, func() error { return errBoom })
	if b.State(t0) != StateClosed {
		t.Fatalf("streak should have reset on success; want Closed")
	}
}

func TestHalfOpenSuccessClosesBreaker(t *testing.T) {
	b := New(1, 1*time.Second)
	t0 := time.Unix(1000, 0)
	b.Call(t0, func() error { return errBoom })
	if b.State(t0) != StateOpen {
		t.Fatalf("expected Open")
	}

	t1 := t0.Add(2 * time.Second)
	if err := b.Call(t1, func() error { return nil }); err != nil {
		t.Fatalf("trial call should be allowed and succeed; got err=%v", err)
	}
	if b.State(t1) != StateClosed {
		t.Fatalf("after successful trial, breaker should be Closed")
	}
}

func TestHalfOpenFailureReopensWithFreshCooldown(t *testing.T) {
	b := New(1, 1*time.Second)
	t0 := time.Unix(1000, 0)
	b.Call(t0, func() error { return errBoom })

	t1 := t0.Add(2 * time.Second)
	if err := b.Call(t1, func() error { return errBoom }); !errors.Is(err, errBoom) {
		t.Fatalf("trial failure should propagate fn's error; got %v", err)
	}

	if err := b.Call(t1.Add(100*time.Millisecond), func() error { return nil }); !errors.Is(err, ErrOpen) {
		t.Fatalf("after trial failure breaker should be Open again; got %v", err)
	}
}
