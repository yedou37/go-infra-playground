package generics

import "testing"

func TestSwap(t *testing.T) {
	got := Swap(Pair[int]{First: 1, Second: 2})
	if got.First != 2 || got.Second != 1 {
		t.Fatalf("got %+v, want {First:2 Second:1}", got)
	}
}

func TestMapSlice(t *testing.T) {
	got := MapSlice([]int{1, 2, 3}, func(v int) string {
		return string(rune('0' + v))
	})
	want := []string{"1", "2", "3"}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestMapSlicePreservesNil(t *testing.T) {
	got := MapSlice[int, int](nil, func(v int) int { return v * 2 })
	if got != nil {
		t.Fatalf("got %#v, want nil", got)
	}
}

func TestLast(t *testing.T) {
	got, ok := Last([]string{"a", "b", "c"})
	if !ok || got != "c" {
		t.Fatalf("got (%q, %v), want (%q, true)", got, ok, "c")
	}
}

func TestLastEmpty(t *testing.T) {
	got, ok := Last[int](nil)
	if ok || got != 0 {
		t.Fatalf("got (%d, %v), want (0, false)", got, ok)
	}
}
