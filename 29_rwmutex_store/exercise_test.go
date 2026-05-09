package rwmutexstore

import "testing"

func TestStoreSetGetLen(t *testing.T) {
	var s Store
	s.Set("a", 1)
	s.Set("b", 2)

	if got := s.Len(); got != 2 {
		t.Fatalf("Len = %d, want 2", got)
	}
	if got, ok := s.Get("a"); !ok || got != 1 {
		t.Fatalf("Get = (%d, %v), want (1, true)", got, ok)
	}
}

func TestStoreMissingKey(t *testing.T) {
	var s Store
	got, ok := s.Get("missing")
	if ok || got != 0 {
		t.Fatalf("Get = (%d, %v), want (0, false)", got, ok)
	}
}
