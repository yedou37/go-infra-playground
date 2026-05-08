package channels

import "testing"

func TestStartGeneratorAndDrain(t *testing.T) {
	got := Drain(StartGenerator(4, 2))
	want := []int{0, 1, 2, 3}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestStartGeneratorWithNonPositiveN(t *testing.T) {
	got := Drain(StartGenerator(0, 1))
	if len(got) != 0 {
		t.Fatalf("got %v, want empty slice", got)
	}
}

func TestDrainClosedChannel(t *testing.T) {
	ch := make(chan int)
	close(ch)

	got := Drain(ch)
	if len(got) != 0 {
		t.Fatalf("got %v, want empty slice", got)
	}
}

func TestTrySend(t *testing.T) {
	ch := make(chan int, 1)
	if ok := TrySend(ch, 10); !ok {
		t.Fatal("first send should succeed")
	}
	if ok := TrySend(ch, 20); ok {
		t.Fatal("second send should fail because buffer is full")
	}
	if got := <-ch; got != 10 {
		t.Fatalf("got %d, want 10", got)
	}
}

func TestTrySendUnbufferedWithoutReceiver(t *testing.T) {
	ch := make(chan int)
	if ok := TrySend(ch, 1); ok {
		t.Fatal("send on unbuffered channel without receiver should not succeed")
	}
}
