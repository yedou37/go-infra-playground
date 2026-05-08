package slicescopy

import "testing"

func TestCloneIntsCreatesIndependentMemory(t *testing.T) {
	src := []int{1, 2, 3}
	got := CloneInts(src)
	if len(got) != len(src) {
		t.Fatalf("len = %d, want %d", len(got), len(src))
	}
	got[0] = 99
	if src[0] != 1 {
		t.Fatalf("src was modified through clone: %v", src)
	}
}

func TestCloneIntsPreservesNil(t *testing.T) {
	var src []int
	got := CloneInts(src)
	if got != nil {
		t.Fatalf("got %#v, want nil", got)
	}
}

func TestWindowReturnsCopy(t *testing.T) {
	src := []int{10, 20, 30, 40}
	got, ok := Window(src, 1, 3)
	if !ok {
		t.Fatal("Window returned ok=false")
	}
	if len(got) != 2 || got[0] != 20 || got[1] != 30 {
		t.Fatalf("got %v, want [20 30]", got)
	}
	got[0] = 999
	if src[1] != 20 {
		t.Fatalf("Window returned alias into original slice: src=%v", src)
	}
}

func TestWindowRejectsInvalidBounds(t *testing.T) {
	src := []int{1, 2, 3}
	cases := [][2]int{
		{-1, 1},
		{2, 1},
		{0, 4},
	}
	for _, tc := range cases {
		if got, ok := Window(src, tc[0], tc[1]); ok || got != nil {
			t.Fatalf("Window(%d, %d) = (%v, %v), want (nil, false)", tc[0], tc[1], got, ok)
		}
	}
}

func TestWindowAllowsEmptyRange(t *testing.T) {
	src := []int{1, 2, 3}

	got, ok := Window(src, 1, 1)
	if !ok {
		t.Fatal("Window returned ok=false for empty range")
	}
	if got == nil || len(got) != 0 {
		t.Fatalf("got %v, want empty non-nil slice", got)
	}
}

func TestPush(t *testing.T) {
	got := Push([]int{1, 2}, 3)
	if len(got) != 3 || got[2] != 3 {
		t.Fatalf("got %v, want [1 2 3]", got)
	}
}
