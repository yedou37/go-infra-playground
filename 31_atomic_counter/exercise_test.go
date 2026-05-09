package atomiccounter

import "testing"

func TestCounterAddLoadReset(t *testing.T) {
	var c Counter
	if got := c.Add(3); got != 3 {
		t.Fatalf("Add = %d, want 3", got)
	}
	if got := c.Add(2); got != 5 {
		t.Fatalf("Add = %d, want 5", got)
	}
	if got := c.Load(); got != 5 {
		t.Fatalf("Load = %d, want 5", got)
	}
	c.Reset()
	if got := c.Load(); got != 0 {
		t.Fatalf("Load = %d, want 0", got)
	}
}
