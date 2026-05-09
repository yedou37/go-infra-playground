package gracefulshutdown

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"sync/atomic"
	"testing"
	"time"
)

func newServer(t *testing.T, h http.Handler) (*http.Server, string) {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	srv := &http.Server{Handler: h}
	addr := fmt.Sprintf("http://%s", ln.Addr().String())

	go func() {
		_ = srv.Serve(ln)
	}()

	// Give the server a moment to be ready.
	time.Sleep(20 * time.Millisecond)
	return srv, addr
}

func TestRunReturnsCleanlyOnStopSignal(t *testing.T) {
	srv, _ := newServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "ok")
	}))

	stop := make(chan struct{})
	done := make(chan error, 1)
	go func() {
		done <- Run(srv, stop, time.Second)
	}()

	close(stop)

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Run returned error: %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Run did not return after Shutdown")
	}
}

func TestRunWaitsForInFlightRequest(t *testing.T) {
	var inflight atomic.Int32
	release := make(chan struct{})

	srv, addr := newServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inflight.Add(1)
		<-release
		inflight.Add(-1)
		_, _ = io.WriteString(w, "done")
	}))

	stop := make(chan struct{})
	done := make(chan error, 1)
	go func() {
		done <- Run(srv, stop, 5*time.Second)
	}()

	clientDone := make(chan struct{})
	go func() {
		defer close(clientDone)
		resp, err := http.Get(addr + "/slow")
		if err != nil {
			return
		}
		_, _ = io.ReadAll(resp.Body)
		resp.Body.Close()
	}()

	// Wait until handler is in-flight.
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) && inflight.Load() == 0 {
		time.Sleep(5 * time.Millisecond)
	}
	if inflight.Load() == 0 {
		t.Fatal("handler never observed as in-flight")
	}

	// Trigger shutdown; release request shortly after.
	close(stop)
	time.AfterFunc(50*time.Millisecond, func() { close(release) })

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Run returned error: %v", err)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("Run did not return after drain")
	}

	<-clientDone
	if inflight.Load() != 0 {
		t.Fatal("handler should have completed during drain")
	}
}
