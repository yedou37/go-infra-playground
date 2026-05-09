package receivers

import (
	"fmt"
	"testing"
)

func TestIncByValueDoesNotMutateCaller(t *testing.T) {
	c := Counter{N: 5}
	c.IncByValue()
	c.IncByValue()
	if c.N != 5 {
		t.Fatalf("value receiver must NOT mutate caller; got N=%d, want 5", c.N)
	}
}

func TestIncByPointerMutatesCaller(t *testing.T) {
	c := &Counter{N: 5}
	c.IncByPointer()
	c.IncByPointer()
	if c.N != 7 {
		t.Fatalf("pointer receiver must mutate caller; got N=%d, want 7", c.N)
	}
}

func TestNewCounterReturnsPointer(t *testing.T) {
	c := NewCounter(3)
	c.IncByPointer()
	if c.N != 4 {
		t.Fatalf("NewCounter should return *Counter so mutations stick; got N=%d", c.N)
	}
}

func TestPointerSatisfiesStringer(t *testing.T) {
	var s fmt.Stringer = NewCounter(9)
	got := s.String()
	if got != "Counter(9)" {
		t.Fatalf("String() = %q, want %q", got, "Counter(9)")
	}
}
