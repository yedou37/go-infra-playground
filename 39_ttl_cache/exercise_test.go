package ttlcache

import (
	"testing"
	"time"
)

func TestSetGetWithinTTL(t *testing.T) {
	c := New[string, int]()
	t0 := time.Unix(1000, 0)
	c.Set("a", 1, t0.Add(5*time.Second))

	v, ok := c.Get("a", t0.Add(2*time.Second))
	if !ok || v != 1 {
		t.Fatalf("want (1,true), got (%d,%v)", v, ok)
	}
}

func TestGetExpiredReturnsMiss(t *testing.T) {
	c := New[string, int]()
	t0 := time.Unix(1000, 0)
	c.Set("a", 1, t0.Add(5*time.Second))

	if _, ok := c.Get("a", t0.Add(10*time.Second)); ok {
		t.Fatalf("expired entry must be invisible to Get")
	}
}

func TestSetWithPastExpiryNeverVisible(t *testing.T) {
	c := New[string, int]()
	t0 := time.Unix(1000, 0)
	c.Set("dead", 9, t0.Add(-1*time.Second))

	if _, ok := c.Get("dead", t0); ok {
		t.Fatalf("entry born expired must not be visible")
	}
	if got := c.Len(t0); got != 0 {
		t.Fatalf("Len at t0 want 0, got %d", got)
	}
	if removed := c.GC(t0); removed != 1 {
		t.Fatalf("GC want 1, got %d", removed)
	}
}

func TestDeleteRemovesEntry(t *testing.T) {
	c := New[string, int]()
	t0 := time.Unix(1000, 0)
	c.Set("a", 1, t0.Add(time.Hour))
	c.Delete("a")
	if _, ok := c.Get("a", t0); ok {
		t.Fatalf("after Delete, key should be missing")
	}
	c.Delete("nope") // no-op, must not panic
}

func TestLenIgnoresExpired(t *testing.T) {
	c := New[string, int]()
	t0 := time.Unix(1000, 0)
	c.Set("a", 1, t0.Add(2*time.Second))
	c.Set("b", 2, t0.Add(10*time.Second))
	c.Set("c", 3, t0.Add(10*time.Second))

	if got := c.Len(t0.Add(5 * time.Second)); got != 2 {
		t.Fatalf("Len want 2 at t0+5s, got %d", got)
	}
}

func TestGCReturnsCountAndKeepsLive(t *testing.T) {
	c := New[string, int]()
	t0 := time.Unix(1000, 0)
	c.Set("a", 1, t0.Add(2*time.Second))
	c.Set("b", 2, t0.Add(10*time.Second))
	c.Set("c", 3, t0.Add(2*time.Second))

	if got := c.GC(t0.Add(5 * time.Second)); got != 2 {
		t.Fatalf("GC want 2 expired removed, got %d", got)
	}
	if _, ok := c.Get("b", t0.Add(5*time.Second)); !ok {
		t.Fatalf("live entry b must survive GC")
	}
	if _, ok := c.Get("a", t0.Add(5*time.Second)); ok {
		t.Fatalf("expired entry a must be gone after GC")
	}
}

func TestSetOverwritesValueAndExpiry(t *testing.T) {
	c := New[string, int]()
	t0 := time.Unix(1000, 0)
	c.Set("a", 1, t0.Add(1*time.Second))
	c.Set("a", 2, t0.Add(10*time.Second))

	v, ok := c.Get("a", t0.Add(5*time.Second))
	if !ok || v != 2 {
		t.Fatalf("Set must overwrite value+expiry; got (%d,%v)", v, ok)
	}
}
