package sliceownership

import "testing"

func TestSetClonesInput(t *testing.T) {
	src := []int{1, 2, 3}
	var s Store

	s.Set(src)
	src[0] = 99

	got := s.Snapshot()
	if got[0] != 1 {
		t.Fatalf("store observed caller mutation: %v", got)
	}
}

func TestSnapshotReturnsCopy(t *testing.T) {
	var s Store
	s.Set([]int{4, 5, 6})

	snap := s.Snapshot()
	snap[0] = 100

	again := s.Snapshot()
	if again[0] != 4 {
		t.Fatalf("Snapshot returned aliased storage: %v", again)
	}
}

func TestAppendUpdatesInternalState(t *testing.T) {
	var s Store
	s.Set([]int{7})
	s.Append(8)

	got := s.Snapshot()
	want := []int{7, 8}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestSetPreservesNil(t *testing.T) {
	var s Store
	s.Set(nil)
	if got := s.Snapshot(); got != nil {
		t.Fatalf("got %#v, want nil", got)
	}
}

func TestSnapshotOfZeroValueStoreIsNil(t *testing.T) {
	var s Store
	if got := s.Snapshot(); got != nil {
		t.Fatalf("got %#v, want nil", got)
	}
}
