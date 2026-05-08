package pipelineshutdown

import (
	"context"
	"testing"
	"time"
)

func TestStartPipelineTransformsValues(t *testing.T) {
	ctx := context.Background()
	in := make(chan int, 2)
	in <- 1
	in <- 2
	close(in)

	var got []int
	for v := range StartPipeline(ctx, in) {
		got = append(got, v)
	}

	want := []int{4, 6}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestStartPipelineStopsOnCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	in := make(chan int)
	out := StartPipeline(ctx, in)

	cancel()

	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("output channel should be closed after cancellation")
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timed out waiting for pipeline shutdown")
	}
}
