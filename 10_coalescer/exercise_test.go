package coalescer

import (
	"testing"
	"time"
)

func TestNotifyDeliversSignal(t *testing.T) {
	c := New()

	c.Notify()

	select {
	case <-c.C():
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timed out waiting for notification")
	}
}

func TestNotifyCoalescesPendingSignals(t *testing.T) {
	c := New()

	c.Notify()
	c.Notify()
	c.Notify()

	select {
	case <-c.C():
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timed out waiting for notification")
	}

	select {
	case <-c.C():
		t.Fatal("expected notifications to be coalesced into one pending signal")
	default:
	}
}

func TestNotifyWorksAgainAfterDrain(t *testing.T) {
	c := New()

	c.Notify()
	<-c.C()
	c.Notify()

	select {
	case <-c.C():
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timed out waiting for second notification")
	}
}

func TestCReturnsSameChannel(t *testing.T) {
	c := New()
	ch1 := c.C()
	ch2 := c.C()
	if ch1 != ch2 {
		t.Fatal("C should return the same channel each time")
	}
}
