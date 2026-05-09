package safego

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestGoNormalDoesNotInvokeHandler(t *testing.T) {
	var mu sync.Mutex
	var captured []any
	PanicHandler = func(v any) {
		mu.Lock()
		captured = append(captured, v)
		mu.Unlock()
	}
	t.Cleanup(func() { PanicHandler = nil })

	done := make(chan struct{})
	Go(func() { close(done) })
	<-done
	time.Sleep(20 * time.Millisecond)

	mu.Lock()
	defer mu.Unlock()
	if len(captured) != 0 {
		t.Fatalf("PanicHandler must not be called for clean exits; got %v", captured)
	}
}

func TestGoRecoversAndCallsHandler(t *testing.T) {
	var (
		mu  sync.Mutex
		got any
	)
	done := make(chan struct{})
	PanicHandler = func(v any) {
		mu.Lock()
		got = v
		mu.Unlock()
		close(done)
	}
	t.Cleanup(func() { PanicHandler = nil })

	Go(func() { panic("boom") })

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for PanicHandler to be called")
	}

	mu.Lock()
	defer mu.Unlock()
	if fmt.Sprint(got) != "boom" {
		t.Fatalf("PanicHandler got %v, want boom", got)
	}
}

func TestGoWithNilHandlerDoesNotCrash(t *testing.T) {
	PanicHandler = nil
	done := make(chan struct{})
	Go(func() {
		defer close(done)
		panic("silent")
	})
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("goroutine did not finish")
	}
	// If we got here without the test process dying, the recover worked.
}

func TestGoSafeNormalReturnsNilThenClose(t *testing.T) {
	ch := GoSafe(func() {})
	err, ok := <-ch
	if !ok {
		t.Fatal("expected one value before close")
	}
	if err != nil {
		t.Fatalf("normal exit should send nil, got %v", err)
	}
	if _, ok := <-ch; ok {
		t.Fatal("channel must be closed after the single value")
	}
}

func TestGoSafePanicReportsError(t *testing.T) {
	ch := GoSafe(func() { panic("kapow") })
	err, ok := <-ch
	if !ok || err == nil {
		t.Fatalf("expected non-nil err, got ok=%v err=%v", ok, err)
	}
	if !strings.Contains(err.Error(), "recovered from panic") {
		t.Fatalf("err should describe a recovered panic; got %q", err.Error())
	}
	if !strings.Contains(err.Error(), "kapow") {
		t.Fatalf("err should mention the panic value; got %q", err.Error())
	}
	if _, ok := <-ch; ok {
		t.Fatal("channel must be closed after the single value")
	}
}
