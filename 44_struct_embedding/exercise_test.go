package embed

import (
	"reflect"
	"testing"
)

func TestBaseLoggerCaptures(t *testing.T) {
	var b BaseLogger
	b.Log("a")
	b.Log("b")
	got := b.Lines()
	if !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Fatalf("got %v, want [a b]", got)
	}
}

func TestBaseLoggerLinesIsDefensiveCopy(t *testing.T) {
	var b BaseLogger
	b.Log("a")
	got := b.Lines()
	got[0] = "MUTATED"
	again := b.Lines()
	if again[0] != "a" {
		t.Fatalf("Lines() must return a defensive copy; internal slice was mutated to %v", again)
	}
}

func TestPrefixLoggerOverridesAndDelegates(t *testing.T) {
	p := NewPrefixLogger("ctrl")
	p.Log("first")
	p.Log("second")

	want := []string{"ctrl: first", "ctrl: second"}
	if got := p.Lines(); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	if got, want := joinLines(p.Lines()), "ctrl: first\nctrl: second"; got != want {
		t.Fatalf("joinLines mismatch:\n got %q\nwant %q", got, want)
	}
}

func TestPrefixLoggerSatisfiesLoggerInterface(t *testing.T) {
	var l Logger = NewPrefixLogger("x")
	l.Log("hi")
	if got := l.Lines(); len(got) != 1 || got[0] != "x: hi" {
		t.Fatalf("interface dispatch broke override; got %v", got)
	}
}
