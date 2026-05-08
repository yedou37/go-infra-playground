package contextloop

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestWaitReceivesValue(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 7

	got, err := Wait(context.Background(), ch)
	if err != nil {
		t.Fatalf("Wait returned error: %v", err)
	}
	if got != 7 {
		t.Fatalf("got %d, want 7", got)
	}
}

func TestWaitContextCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	got, err := Wait[int](ctx, make(chan int))
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
	if got != 0 {
		t.Fatalf("got %d, want zero value", got)
	}
}

func TestWaitClosedChannel(t *testing.T) {
	ch := make(chan int)
	close(ch)

	got, err := Wait(context.Background(), ch)
	if !errors.Is(err, ErrChannelClosed) {
		t.Fatalf("err = %v, want ErrChannelClosed", err)
	}
	if got != 0 {
		t.Fatalf("got %d, want zero value", got)
	}
}

func TestSend(t *testing.T) {
	ch := make(chan int, 1)
	if err := Send(context.Background(), ch, 11); err != nil {
		t.Fatalf("Send returned error: %v", err)
	}
	if got := <-ch; got != 11 {
		t.Fatalf("got %d, want 11", got)
	}
}

func TestSendContextCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if err := Send(ctx, make(chan int), 11); !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
}

func TestRunOnTicksStopsOnCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ticks := make(chan time.Time)
	called := make(chan struct{}, 2)
	done := make(chan struct{})

	go func() {
		RunOnTicks(ctx, ticks, func() {
			called <- struct{}{}
		})
		close(done)
	}()

	ticks <- time.Time{}
	ticks <- time.Time{}

	for i := 0; i < 2; i++ {
		select {
		case <-called:
		case <-time.After(100 * time.Millisecond):
			t.Fatal("timed out waiting for tick callback")
		}
	}

	cancel()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("RunOnTicks did not stop after cancellation")
	}
}
