package interfaces

import (
	"context"
	"errors"
	"testing"
)

type recordingHandler struct {
	calls *[]string
}

func (h recordingHandler) Handle(ctx context.Context, event Event) error {
	*h.calls = append(*h.calls, event.Name)
	return nil
}

func TestHandlerFuncImplementsHandler(t *testing.T) {
	var h Handler = HandlerFunc(func(ctx context.Context, event Event) error {
		if event.Name != "added" {
			t.Fatalf("event = %+v, want Name=added", event)
		}
		return nil
	})

	if err := h.Handle(context.Background(), Event{Name: "added"}); err != nil {
		t.Fatalf("Handle returned error: %v", err)
	}
}

func TestDispatchCallsHandlersInOrder(t *testing.T) {
	var calls []string

	err := Dispatch(
		context.Background(),
		Event{Name: "sync"},
		recordingHandler{calls: &calls},
		HandlerFunc(func(ctx context.Context, event Event) error {
			calls = append(calls, "fn:"+event.Name)
			return nil
		}),
	)
	if err != nil {
		t.Fatalf("Dispatch returned error: %v", err)
	}

	want := []string{"sync", "fn:sync"}
	if len(calls) != len(want) {
		t.Fatalf("calls = %v, want %v", calls, want)
	}
	for i := range want {
		if calls[i] != want[i] {
			t.Fatalf("calls = %v, want %v", calls, want)
		}
	}
}

func TestDispatchStopsOnFirstError(t *testing.T) {
	errBoom := errors.New("boom")
	var called bool

	err := Dispatch(
		context.Background(),
		Event{Name: "delete"},
		HandlerFunc(func(ctx context.Context, event Event) error {
			return errBoom
		}),
		HandlerFunc(func(ctx context.Context, event Event) error {
			called = true
			return nil
		}),
	)
	if !errors.Is(err, errBoom) {
		t.Fatalf("err = %v, want errBoom", err)
	}
	if called {
		t.Fatal("later handlers should not be called after an error")
	}
}

func TestDispatchWithNoHandlers(t *testing.T) {
	err := Dispatch(context.Background(), Event{Name: "noop"})
	if err != nil {
		t.Fatalf("Dispatch returned error: %v", err)
	}
}
