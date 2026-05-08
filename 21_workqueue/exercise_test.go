package workqueue

import (
	"context"
	"errors"
	"testing"
)

func TestAddDeduplicatesPendingKeys(t *testing.T) {
	q := New(4)
	if !q.Add("a") {
		t.Fatal("first Add should succeed")
	}
	if !q.Add("a") {
		t.Fatal("duplicate Add should report queue still open")
	}

	got, err := q.Get(context.Background())
	if err != nil {
		t.Fatalf("Get returned error: %v", err)
	}
	if got != "a" {
		t.Fatalf("got %q, want %q", got, "a")
	}
}

func TestDoneAllowsReenqueue(t *testing.T) {
	q := New(2)
	q.Add("a")

	got, err := q.Get(context.Background())
	if err != nil || got != "a" {
		t.Fatalf("Get = (%q, %v), want (%q, nil)", got, err, "a")
	}
	q.Done("a")

	if !q.Add("a") {
		t.Fatal("Add should succeed after Done")
	}
	got, err = q.Get(context.Background())
	if err != nil || got != "a" {
		t.Fatalf("Get = (%q, %v), want (%q, nil)", got, err, "a")
	}
}

func TestGetReturnsErrClosedAfterShutdown(t *testing.T) {
	q := New(1)
	q.ShutDown()

	_, err := q.Get(context.Background())
	if !errors.Is(err, ErrClosed) {
		t.Fatalf("err = %v, want ErrClosed", err)
	}
}

func TestAddReturnsFalseAfterShutdown(t *testing.T) {
	q := New(1)
	q.ShutDown()
	if ok := q.Add("a"); ok {
		t.Fatal("Add should return false after shutdown")
	}
}
