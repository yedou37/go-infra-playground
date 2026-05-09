package channelvsmutex

import "testing"

func TestMutexCounter(t *testing.T) {
	var c MutexCounter
	c.Inc()
	c.Inc()
	if got := c.Load(); got != 2 {
		t.Fatalf("Load = %d, want 2", got)
	}
}

func TestChannelCounter(t *testing.T) {
	c := NewChannelCounter()
	defer c.Close()

	c.Inc()
	c.Inc()
	c.Inc()

	if got := c.Load(); got != 3 {
		t.Fatalf("Load = %d, want 3", got)
	}
}
