package queue

import "testing"

func TestQueueBasicFlow(t *testing.T) {
	q := New[int](2)
	q.Push(10)
	q.Push(20)
	if q.Len() != 2 {
		t.Fatalf("Len = %d, want 2", q.Len())
	}

	got, ok := q.Pop()
	if !ok || got != 10 {
		t.Fatalf("first Pop = (%d, %v), want (10, true)", got, ok)
	}

	got, ok = q.Pop()
	if !ok || got != 20 {
		t.Fatalf("second Pop = (%d, %v), want (20, true)", got, ok)
	}

	got, ok = q.Pop()
	if ok || got != 0 {
		t.Fatalf("empty Pop = (%d, %v), want (0, false)", got, ok)
	}
}

func TestSnapshotReturnsCopy(t *testing.T) {
	q := New[int](0)
	q.Push(1)
	q.Push(2)
	q.Push(3)

	snap := q.Snapshot()
	if len(snap) != 3 {
		t.Fatalf("len = %d, want 3", len(snap))
	}
	snap[0] = 999

	again := q.Snapshot()
	if again[0] != 1 {
		t.Fatalf("Snapshot returned aliased memory: %v", again)
	}
}

func TestLenExcludesPoppedItems(t *testing.T) {
	q := New[string](1)
	q.Push("a")
	q.Push("b")
	if _, ok := q.Pop(); !ok {
		t.Fatal("expected first pop to succeed")
	}
	if q.Len() != 1 {
		t.Fatalf("Len = %d, want 1", q.Len())
	}
}

func TestNewTreatsNegativeCapacityAsZero(t *testing.T) {
	q := New[int](-3)
	q.Push(1)

	got, ok := q.Pop()
	if !ok || got != 1 {
		t.Fatalf("Pop = (%d, %v), want (1, true)", got, ok)
	}
}

func TestSnapshotEmptyQueue(t *testing.T) {
	q := New[int](0)

	got := q.Snapshot()
	if got == nil || len(got) != 0 {
		t.Fatalf("got %v, want empty non-nil slice", got)
	}
}
