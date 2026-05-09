package tokenbucket

import (
	"testing"
	"time"
)

func TestStartsFull(t *testing.T) {
	l := New(1, 3)
	t0 := time.Unix(1000, 0)
	for i := 0; i < 3; i++ {
		if !l.Allow(t0) {
			t.Fatalf("call %d should be allowed (bucket starts full)", i)
		}
	}
	if l.Allow(t0) {
		t.Fatalf("4th call at the same instant should be denied")
	}
}

func TestRefillIsLazy(t *testing.T) {
	l := New(1, 3)
	t0 := time.Unix(1000, 0)
	for i := 0; i < 3; i++ {
		l.Allow(t0)
	}
	if l.Allow(t0.Add(500 * time.Millisecond)) {
		t.Fatalf("0.5s @ 1 token/s should NOT yet refill a whole token")
	}
	if !l.Allow(t0.Add(1100 * time.Millisecond)) {
		t.Fatalf("1.1s @ 1 token/s should refill at least 1 token")
	}
}

func TestBurstIsCapped(t *testing.T) {
	l := New(10, 2)
	t0 := time.Unix(1000, 0)
	l.Allow(t0)
	l.Allow(t0)
	if l.Allow(t0) {
		t.Fatalf("3rd call at t0 must be denied; burst=2")
	}
	later := t0.Add(1 * time.Hour)
	if !l.Allow(later) {
		t.Fatalf("after long idle, 1st call should pass")
	}
	if !l.Allow(later) {
		t.Fatalf("after long idle, 2nd call should pass (burst=2)")
	}
	if l.Allow(later) {
		t.Fatalf("after long idle, 3rd call must still be denied; tokens cap at burst")
	}
}

func TestAllowNAtomic(t *testing.T) {
	l := New(1, 5)
	t0 := time.Unix(1000, 0)
	if !l.AllowN(t0, 5) {
		t.Fatalf("AllowN(5) on full bucket should succeed")
	}
	if l.AllowN(t0, 1) {
		t.Fatalf("AllowN(1) right after draining should fail")
	}
	if l.AllowN(t0, 3) {
		t.Fatalf("AllowN that cannot be satisfied must NOT consume tokens")
	}
	if !l.AllowN(t0.Add(2*time.Second), 2) {
		t.Fatalf("after 2s @ 1 token/s, AllowN(2) should succeed")
	}
}

func TestAllowNZeroOrNegative(t *testing.T) {
	l := New(1, 3)
	t0 := time.Unix(1000, 0)
	if !l.AllowN(t0, 0) {
		t.Fatalf("AllowN(0) should be a no-op success")
	}
	if !l.AllowN(t0, -5) {
		t.Fatalf("AllowN(<=0) should be a no-op success")
	}
	for i := 0; i < 3; i++ {
		if !l.Allow(t0) {
			t.Fatalf("the no-op AllowN must not have consumed tokens")
		}
	}
}

func TestAllowNLargerThanBurst(t *testing.T) {
	l := New(100, 3)
	t0 := time.Unix(1000, 0)
	if l.AllowN(t0.Add(1*time.Hour), 10) {
		t.Fatalf("AllowN(n>burst) is impossible to satisfy and must return false")
	}
}
