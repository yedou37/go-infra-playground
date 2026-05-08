package slicesemantics

import "testing"

func TestSumNilAndEmpty(t *testing.T) {
	if got := Sum(nil); got != 0 {
		t.Fatalf("Sum(nil) = %d, want 0", got)
	}
	if got := Sum([]int{}); got != 0 {
		t.Fatalf("Sum(empty) = %d, want 0", got)
	}
}

func TestSumDoesNotMutateInput(t *testing.T) {
	xs := []int{1, 2, 3}

	got := Sum(xs)
	if got != 6 {
		t.Fatalf("sum = %d, want 6", got)
	}
	if xs[0] != 1 || xs[1] != 2 || xs[2] != 3 {
		t.Fatalf("Sum mutated input: %v", xs)
	}
}

func TestZeroInPlaceMutatesSharedBackingArray(t *testing.T) {
	xs := []int{3, 4, 5}

	ZeroInPlace(xs)

	want := []int{0, 0, 0}
	for i := range want {
		if xs[i] != want[i] {
			t.Fatalf("xs = %v, want %v", xs, want)
		}
	}
}

func TestZeroInPlaceOnSubsliceAffectsParentWindow(t *testing.T) {
	parent := []int{1, 2, 3, 4}
	view := parent[1:3]

	ZeroInPlace(view)

	want := []int{1, 0, 0, 4}
	for i := range want {
		if parent[i] != want[i] {
			t.Fatalf("parent = %v, want %v", parent, want)
		}
	}
}

func TestAppendValueReturnsNewHeaderToCaller(t *testing.T) {
	xs := make([]int, 1)
	xs[0] = 7

	got := AppendValue(xs, 9)

	if len(xs) != 1 {
		t.Fatalf("original len = %d, want 1", len(xs))
	}
	if len(got) != 2 || got[0] != 7 || got[1] != 9 {
		t.Fatalf("got = %v, want [7 9]", got)
	}
}

func TestAppendValueMayReuseBackingArrayWhenCapacityAllows(t *testing.T) {
	xs := make([]int, 1, 3)
	xs[0] = 7

	got := AppendValue(xs, 9)

	if len(xs) != 1 {
		t.Fatalf("original len = %d, want 1", len(xs))
	}
	if got[1] != 9 {
		t.Fatalf("got = %v, want appended value 9", got)
	}
	expanded := xs[:2]
	if expanded[1] != 9 {
		t.Fatalf("shared backing array was not updated as expected: xs[:2]=%v", expanded)
	}
}

func TestCloneReturnsIndependentBackingStorage(t *testing.T) {
	xs := []int{8, 9, 10}

	got := Clone(xs)
	if len(got) != len(xs) {
		t.Fatalf("len = %d, want %d", len(got), len(xs))
	}
	got[0] = 100
	if xs[0] != 8 {
		t.Fatalf("Clone returned aliased storage: xs=%v clone=%v", xs, got)
	}
}

func TestCloneOfSubsliceIsStillIndependent(t *testing.T) {
	parent := []int{1, 2, 3, 4}
	got := Clone(parent[1:3])

	got[0] = 99
	if parent[1] != 2 {
		t.Fatalf("parent = %v, clone = %v, want independent storage", parent, got)
	}
}

func TestClonePreservesNil(t *testing.T) {
	var xs []int
	got := Clone(xs)
	if got != nil {
		t.Fatalf("got %#v, want nil", got)
	}
}

func TestMutatePacketShowsMixedValueAndSliceBehavior(t *testing.T) {
	p := Packet{
		Version: 1,
		Data:    []byte{10, 20},
	}

	gotVersion := MutatePacket(p)

	if gotVersion != 2 {
		t.Fatalf("returned version = %d, want 2", gotVersion)
	}
	if p.Version != 1 {
		t.Fatalf("Version should not change in caller: %d", p.Version)
	}
	if p.Data[0] != 11 {
		t.Fatalf("Data[0] = %d, want 11", p.Data[0])
	}
}

func TestMutatePacketHandlesEmptyData(t *testing.T) {
	p := Packet{Version: 5}

	gotVersion := MutatePacket(p)

	if gotVersion != 6 {
		t.Fatalf("returned version = %d, want 6", gotVersion)
	}
	if p.Version != 5 {
		t.Fatalf("caller Version = %d, want 5", p.Version)
	}
	if p.Data != nil {
		t.Fatalf("caller Data = %v, want nil", p.Data)
	}
}
