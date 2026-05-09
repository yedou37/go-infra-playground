package leaderelection

import (
	"testing"
	"time"
)

func TestAcquireWhenEmpty(t *testing.T) {
	e := New()
	t0 := time.Unix(1000, 0)
	if !e.TryAcquire("a", t0, 5*time.Second) {
		t.Fatalf("first acquire should succeed")
	}
	h, ok := e.CurrentHolder(t0)
	if !ok || h != "a" {
		t.Fatalf("holder want a, got %q ok=%v", h, ok)
	}
}

func TestAcquireRejectedWhileHeld(t *testing.T) {
	e := New()
	t0 := time.Unix(1000, 0)
	e.TryAcquire("a", t0, 5*time.Second)
	if e.TryAcquire("b", t0.Add(1*time.Second), 5*time.Second) {
		t.Fatalf("b must not be able to acquire while a's lease is valid")
	}
	h, _ := e.CurrentHolder(t0.Add(1 * time.Second))
	if h != "a" {
		t.Fatalf("a must remain leader, got %q", h)
	}
}

func TestAcquireAfterExpiry(t *testing.T) {
	e := New()
	t0 := time.Unix(1000, 0)
	e.TryAcquire("a", t0, 5*time.Second)
	if !e.TryAcquire("b", t0.Add(10*time.Second), 5*time.Second) {
		t.Fatalf("b must be able to take over after a's lease expired")
	}
	h, ok := e.CurrentHolder(t0.Add(10 * time.Second))
	if !ok || h != "b" {
		t.Fatalf("holder want b, got %q ok=%v", h, ok)
	}
}

func TestRenewOnlyByCurrentHolder(t *testing.T) {
	e := New()
	t0 := time.Unix(1000, 0)
	e.TryAcquire("a", t0, 5*time.Second)

	if e.Renew("b", t0.Add(1*time.Second), 5*time.Second) {
		t.Fatalf("b is not the holder, Renew must fail")
	}
	if !e.Renew("a", t0.Add(1*time.Second), 5*time.Second) {
		t.Fatalf("a is the holder, Renew must succeed")
	}
	if _, ok := e.CurrentHolder(t0.Add(5 * time.Second)); !ok {
		t.Fatalf("after renew the lease should still be valid at t0+5s")
	}
}

func TestRenewFailsAfterExpiry(t *testing.T) {
	e := New()
	t0 := time.Unix(1000, 0)
	e.TryAcquire("a", t0, 5*time.Second)
	if e.Renew("a", t0.Add(10*time.Second), 5*time.Second) {
		t.Fatalf("expired holder must not be able to silently renew")
	}
}

func TestReleaseClearsLease(t *testing.T) {
	e := New()
	t0 := time.Unix(1000, 0)
	e.TryAcquire("a", t0, 100*time.Second)
	if !e.Release("a") {
		t.Fatalf("Release by current holder should succeed")
	}
	if _, ok := e.CurrentHolder(t0); ok {
		t.Fatalf("after release there should be no holder")
	}
	if !e.TryAcquire("b", t0, 5*time.Second) {
		t.Fatalf("b should be able to acquire after voluntary release")
	}
}
