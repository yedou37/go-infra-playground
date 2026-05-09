package waitgroupfanout

import "testing"

func TestFanOutPreservesOrder(t *testing.T) {
	got := FanOut([]int{1, 2, 3}, func(v int) int {
		return v * 10
	})
	want := []int{10, 20, 30}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestFanOutEmpty(t *testing.T) {
	got := FanOut[int, int](nil, func(v int) int { return v })
	if len(got) != 0 {
		t.Fatalf("got %v, want empty slice", got)
	}
}
