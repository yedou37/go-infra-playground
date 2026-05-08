package selectpatterns

import (
	"testing"
	"time"
)

func TestOrDoneForwardsValues(t *testing.T) {
	done := make(chan struct{})
	in := make(chan int, 2)
	in <- 1
	in <- 2
	close(in)

	got := make([]int, 0, 2)
	for v := range OrDone(done, in) {
		got = append(got, v)
	}

	want := []int{1, 2}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestOrDoneStopsOnDone(t *testing.T) {
	done := make(chan struct{})
	in := make(chan int)
	out := OrDone(done, in)

	close(done)

	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("output channel should be closed after done")
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timed out waiting for output to close")
	}
}

func TestFanInMergesInputs(t *testing.T) {
	done := make(chan struct{})
	a := make(chan int, 2)
	b := make(chan int, 2)

	a <- 1
	a <- 2
	b <- 3
	b <- 4
	close(a)
	close(b)

	got := make(map[int]bool)
	for v := range FanIn(done, a, b) {
		got[v] = true
	}

	for _, want := range []int{1, 2, 3, 4} {
		if !got[want] {
			t.Fatalf("missing value %d in %v", want, got)
		}
	}
}

func TestFanInStopsOnDone(t *testing.T) {
	done := make(chan struct{})
	a := make(chan int)
	out := FanIn(done, a)

	close(done)

	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("output channel should be closed after done")
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timed out waiting for output to close")
	}
}

func TestFanInWithNoInputsClosesImmediately(t *testing.T) {
	done := make(chan struct{})
	out := FanIn[int](done)

	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("output channel should be closed when there are no inputs")
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timed out waiting for output to close")
	}
}
