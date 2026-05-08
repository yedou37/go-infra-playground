package arrayvsslice

import "testing"

func TestBumpArrayDoesNotMutateCaller(t *testing.T) {
	orig := [3]int{1, 2, 3}

	got := BumpArray(orig)

	if orig != [3]int{1, 2, 3} {
		t.Fatalf("orig = %v, want unchanged", orig)
	}
	if got != [3]int{2, 3, 4} {
		t.Fatalf("got = %v, want [2 3 4]", got)
	}
}

func TestBumpSliceMutatesCaller(t *testing.T) {
	xs := []int{1, 2, 3}

	BumpSlice(xs)

	want := []int{2, 3, 4}
	for i := range want {
		if xs[i] != want[i] {
			t.Fatalf("xs = %v, want %v", xs, want)
		}
	}
}

func TestToSliceDoesNotAliasOriginalArrayVariable(t *testing.T) {
	orig := [3]int{5, 6, 7}

	got := ToSlice(orig)
	got[0] = 99

	if orig[0] != 5 {
		t.Fatalf("orig = %v, want unchanged", orig)
	}
	if got[0] != 99 {
		t.Fatalf("got = %v, want first element updated", got)
	}
}

func TestToSliceHasExpectedLength(t *testing.T) {
	got := ToSlice([3]int{1, 2, 3})
	if len(got) != 3 {
		t.Fatalf("len = %d, want 3", len(got))
	}
}
