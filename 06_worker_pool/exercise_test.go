package workerpool

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestMapPreservesInputOrder(t *testing.T) {
	tasks := []Task{
		{ID: 1, Value: 2},
		{ID: 2, Value: 4},
		{ID: 3, Value: 6},
	}

	got, err := Map(context.Background(), 2, tasks, func(ctx context.Context, task Task) (Result, error) {
		if task.ID == 1 {
			time.Sleep(20 * time.Millisecond)
		}
		return Result{
			ID:    task.ID,
			Value: task.Value * 10,
		}, nil
	})
	if err != nil {
		t.Fatalf("Map returned error: %v", err)
	}

	want := []Result{
		{ID: 1, Value: 20},
		{ID: 2, Value: 40},
		{ID: 3, Value: 60},
	}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestMapReturnsFirstError(t *testing.T) {
	errBoom := errors.New("boom")
	tasks := []Task{
		{ID: 1, Value: 1},
		{ID: 2, Value: 2},
		{ID: 3, Value: 3},
	}

	_, err := Map(context.Background(), 2, tasks, func(ctx context.Context, task Task) (Result, error) {
		if task.ID == 2 {
			return Result{}, errBoom
		}

		select {
		case <-ctx.Done():
			return Result{}, ctx.Err()
		case <-time.After(50 * time.Millisecond):
		}

		return Result{
			ID:    task.ID,
			Value: task.Value,
		}, nil
	})
	if !errors.Is(err, errBoom) {
		t.Fatalf("err = %v, want errBoom", err)
	}
}

func TestMapRespectsCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := Map(ctx, 4, []Task{{ID: 1, Value: 10}}, func(ctx context.Context, task Task) (Result, error) {
		t.Fatal("fn should not be called after cancellation")
		return Result{}, nil
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
}

func TestMapTreatsNonPositiveWorkersAsOne(t *testing.T) {
	tasks := []Task{{ID: 1, Value: 3}}

	got, err := Map(context.Background(), 0, tasks, func(ctx context.Context, task Task) (Result, error) {
		return Result{ID: task.ID, Value: task.Value + 1}, nil
	})
	if err != nil {
		t.Fatalf("Map returned error: %v", err)
	}
	if len(got) != 1 || got[0].Value != 4 {
		t.Fatalf("got %v, want one processed result", got)
	}
}

func TestMapEmptyTasks(t *testing.T) {
	called := false
	got, err := Map(context.Background(), 2, nil, func(ctx context.Context, task Task) (Result, error) {
		called = true
		return Result{}, nil
	})
	if err != nil {
		t.Fatalf("Map returned error: %v", err)
	}
	if called {
		t.Fatal("fn should not be called for empty task list")
	}
	if len(got) != 0 {
		t.Fatalf("got %v, want nil or empty slice", got)
	}
}
