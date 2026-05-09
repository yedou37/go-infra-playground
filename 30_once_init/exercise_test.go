package onceinit

import "testing"

func TestGetInitializesOnce(t *testing.T) {
	var l Loader
	calls := 0

	first := l.Get(func() string {
		calls++
		return "ready"
	})
	second := l.Get(func() string {
		calls++
		return "again"
	})

	if calls != 1 {
		t.Fatalf("calls = %d, want 1", calls)
	}
	if first != "ready" || second != "ready" {
		t.Fatalf("got (%q, %q), want (%q, %q)", first, second, "ready", "ready")
	}
}
