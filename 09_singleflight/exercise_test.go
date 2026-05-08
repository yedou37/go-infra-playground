package singleflight

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestDoSharesConcurrentCalls(t *testing.T) {
	var g Group[int]

	start := make(chan struct{})
	var fnCalls int

	type output struct {
		v      int
		err    error
		shared bool
	}

	results := make(chan output, 2)
	run := func() {
		v, err, shared := g.Do("alpha", func() (int, error) {
			fnCalls++
			close(start)
			time.Sleep(20 * time.Millisecond)
			return 42, nil
		})
		results <- output{v: v, err: err, shared: shared}
	}

	go run()
	<-start
	go run()

	first := <-results
	second := <-results

	if fnCalls != 1 {
		t.Fatalf("fn calls = %d, want 1", fnCalls)
	}
	if first.err != nil || second.err != nil {
		t.Fatalf("got errors (%v, %v), want nil", first.err, second.err)
	}
	if first.v != 42 || second.v != 42 {
		t.Fatalf("got values (%d, %d), want (42, 42)", first.v, second.v)
	}
	if first.shared == second.shared {
		t.Fatalf("want exactly one shared caller, got (%v, %v)", first.shared, second.shared)
	}
}

func TestDoSharesErrors(t *testing.T) {
	var g Group[int]
	errBoom := errors.New("boom")

	start := make(chan struct{})
	done := make(chan struct{})

	go func() {
		defer close(done)
		_, err, shared := g.Do("alpha", func() (int, error) {
			close(start)
			time.Sleep(20 * time.Millisecond)
			return 0, errBoom
		})
		if !errors.Is(err, errBoom) {
			t.Errorf("err = %v, want errBoom", err)
		}
		if shared {
			t.Errorf("first caller should not be shared")
		}
	}()

	<-start
	_, err, shared := g.Do("alpha", func() (int, error) {
		t.Fatal("second concurrent caller should not execute fn")
		return 0, nil
	})
	if !errors.Is(err, errBoom) {
		t.Fatalf("err = %v, want errBoom", err)
	}
	if !shared {
		t.Fatal("second caller should be shared")
	}

	<-done
}

func TestDoAllowsNewCallAfterCompletion(t *testing.T) {
	var g Group[int]
	var mu sync.Mutex
	fnCalls := 0

	run := func() int {
		v, err, _ := g.Do("alpha", func() (int, error) {
			mu.Lock()
			defer mu.Unlock()
			fnCalls++
			return fnCalls, nil
		})
		if err != nil {
			t.Fatalf("Do returned error: %v", err)
		}
		return v
	}

	first := run()
	second := run()

	if first != 1 || second != 2 {
		t.Fatalf("got (%d, %d), want (1, 2)", first, second)
	}
}

func TestDoDoesNotShareDifferentKeys(t *testing.T) {
	var g Group[int]
	var mu sync.Mutex
	calls := 0

	run := func(key string) (int, bool) {
		v, err, shared := g.Do(key, func() (int, error) {
			mu.Lock()
			defer mu.Unlock()
			calls++
			return calls, nil
		})
		if err != nil {
			t.Fatalf("Do returned error: %v", err)
		}
		return v, shared
	}

	first, sharedFirst := run("alpha")
	second, sharedSecond := run("beta")

	if calls != 2 {
		t.Fatalf("calls = %d, want 2", calls)
	}
	if sharedFirst || sharedSecond {
		t.Fatalf("different keys should not be shared: (%v, %v)", sharedFirst, sharedSecond)
	}
	if first == second {
		t.Fatalf("got equal results (%d, %d), want distinct executions", first, second)
	}
}
