package broadcaster

import (
	"testing"
	"time"
)

func recv[T any](t *testing.T, ch <-chan T, want T, timeout time.Duration) {
	t.Helper()
	select {
	case got, ok := <-ch:
		if !ok {
			t.Fatalf("channel closed unexpectedly")
		}
		var anyGot any = got
		var anyWant any = want
		if anyGot != anyWant {
			t.Fatalf("got %v, want %v", got, want)
		}
	case <-time.After(timeout):
		t.Fatalf("timeout waiting for value")
	}
}

func TestSingleSubscriberReceivesPublish(t *testing.T) {
	b := New[int]()
	ch := b.Subscribe(2)
	b.Publish(1)
	recv(t, ch, 1, time.Second)
}

func TestMultipleSubscribersAllReceive(t *testing.T) {
	b := New[int]()
	a := b.Subscribe(4)
	c := b.Subscribe(4)
	b.Publish(7)
	recv(t, a, 7, time.Second)
	recv(t, c, 7, time.Second)
}

func TestSlowSubscriberDropsButOthersStillReceive(t *testing.T) {
	b := New[int]()
	slow := b.Subscribe(1) // capacity 1, never drained
	fast := b.Subscribe(8)

	for i := 0; i < 5; i++ {
		b.Publish(i)
	}

	// fast must see all 5 events in order.
	for i := 0; i < 5; i++ {
		recv(t, fast, i, time.Second)
	}

	// slow only had a buffer of 1, so it must see at most 1 event,
	// the rest must have been dropped (NOT block the publisher).
	count := 0
loop:
	for {
		select {
		case <-slow:
			count++
		default:
			break loop
		}
	}
	if count > 1 {
		t.Fatalf("slow subscriber expected at most 1 event due to drop, got %d", count)
	}
}

func TestUnsubscribeStopsDelivery(t *testing.T) {
	b := New[int]()
	a := b.Subscribe(4)
	c := b.Subscribe(4)

	b.Unsubscribe(a)

	b.Publish(99)
	recv(t, c, 99, time.Second)

	// `a` should be closed; reads must not block.
	select {
	case _, ok := <-a:
		if ok {
			t.Fatalf("after Unsubscribe, channel should be closed (or empty + closed)")
		}
	case <-time.After(time.Second):
		t.Fatalf("after Unsubscribe, reads must not block")
	}
}

func TestCloseClosesAllSubscribers(t *testing.T) {
	b := New[int]()
	a := b.Subscribe(2)
	c := b.Subscribe(2)
	b.Close()

	for _, ch := range []<-chan int{a, c} {
		select {
		case _, ok := <-ch:
			if ok {
				t.Fatalf("after Close, subscriber channel must be closed")
			}
		case <-time.After(time.Second):
			t.Fatalf("after Close, reads on subscriber channel must not block")
		}
	}

	// Publish after Close must be a silent no-op (no panic).
	b.Publish(1)

	// Subscribe after Close must return an already-closed channel.
	d := b.Subscribe(1)
	select {
	case _, ok := <-d:
		if ok {
			t.Fatalf("Subscribe after Close should return a closed channel")
		}
	case <-time.After(time.Second):
		t.Fatalf("Subscribe after Close must return a closed channel, not a live one")
	}
}

func TestCloseIsIdempotent(t *testing.T) {
	b := New[int]()
	_ = b.Subscribe(1)
	b.Close()
	b.Close() // must not panic
}
