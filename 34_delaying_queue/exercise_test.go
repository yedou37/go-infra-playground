package delayingqueue

import (
	"testing"
	"time"
)

func TestAddAfterReadyImmediately(t *testing.T) {
	q := New()
	now := time.Unix(1000, 0)
	q.AddAfter("a", 0, now)
	item, ok, wait := q.Get(now)
	if !ok || item != "a" || wait != 0 {
		t.Fatalf("want (a,true,0), got (%q,%v,%v)", item, ok, wait)
	}
}

func TestAddAfterNotYetReady(t *testing.T) {
	q := New()
	now := time.Unix(1000, 0)
	q.AddAfter("a", 5*time.Second, now)
	_, ok, wait := q.Get(now)
	if ok {
		t.Fatalf("expected not ready")
	}
	if wait != 5*time.Second {
		t.Fatalf("want wait=5s, got %v", wait)
	}
}

func TestEarliestDeadlineFirst(t *testing.T) {
	q := New()
	now := time.Unix(1000, 0)
	q.AddAfter("late", 10*time.Second, now)
	q.AddAfter("soon", 1*time.Second, now)
	q.AddAfter("mid", 5*time.Second, now)

	if got := q.Len(); got != 3 {
		t.Fatalf("len want 3, got %d", got)
	}

	item, ok, _ := q.Get(now.Add(2 * time.Second))
	if !ok || item != "soon" {
		t.Fatalf("first ready want soon, got %q ok=%v", item, ok)
	}

	item, ok, _ = q.Get(now.Add(6 * time.Second))
	if !ok || item != "mid" {
		t.Fatalf("second ready want mid, got %q ok=%v", item, ok)
	}

	_, ok, wait := q.Get(now.Add(6 * time.Second))
	if ok {
		t.Fatalf("expected not ready yet")
	}
	if wait != 4*time.Second {
		t.Fatalf("want wait=4s, got %v", wait)
	}
}

func TestEmptyQueueWaitIsZero(t *testing.T) {
	q := New()
	_, ok, wait := q.Get(time.Unix(0, 0))
	if ok || wait != 0 {
		t.Fatalf("empty queue must return (false, 0); got ok=%v wait=%v", ok, wait)
	}
}

func TestShutdownPreventsFurtherWork(t *testing.T) {
	q := New()
	now := time.Unix(1000, 0)
	q.AddAfter("a", 0, now)
	q.Shutdown()

	if !q.ShuttingDown() {
		t.Fatalf("ShuttingDown must report true")
	}
	_, ok, wait := q.Get(now)
	if ok || wait != 0 {
		t.Fatalf("after shutdown Get must return (false, 0); got ok=%v wait=%v", ok, wait)
	}

	q.AddAfter("b", 0, now)
	if q.Len() > 1 {
		t.Fatalf("AddAfter after Shutdown must be a no-op")
	}
}
