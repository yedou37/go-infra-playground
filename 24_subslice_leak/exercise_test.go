package subsliceleak

import "testing"

func TestStableWindowReturnsCopy(t *testing.T) {
	src := []byte("abcdef")

	got, ok := StableWindow(src, 1, 4)
	if !ok {
		t.Fatal("StableWindow returned ok=false")
	}
	if string(got) != "bcd" {
		t.Fatalf("got %q, want %q", got, "bcd")
	}

	src[1] = 'Z'
	if string(got) != "bcd" {
		t.Fatalf("got %q, want detached copy", got)
	}
}

func TestStableWindowHasTightCapacity(t *testing.T) {
	got, ok := StableWindow([]byte("abcdef"), 2, 5)
	if !ok {
		t.Fatal("StableWindow returned ok=false")
	}
	if cap(got) != len(got) {
		t.Fatalf("cap = %d, want len %d", cap(got), len(got))
	}
}

func TestStableWindowRejectsInvalidBounds(t *testing.T) {
	cases := [][2]int{{-1, 1}, {2, 1}, {0, 7}}
	for _, tc := range cases {
		got, ok := StableWindow([]byte("abcdef"), tc[0], tc[1])
		if ok || got != nil {
			t.Fatalf("StableWindow(%d, %d) = (%v, %v), want (nil, false)", tc[0], tc[1], got, ok)
		}
	}
}
