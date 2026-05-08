package bufferreuse

import "testing"

func TestEncodeIntEncodesOneDigit(t *testing.T) {
	var e Encoder

	got := e.EncodeInt(7)

	if string(got) != "7" {
		t.Fatalf("got %q, want %q", got, "7")
	}
}

func TestEncodeIntEncodesTwoDigits(t *testing.T) {
	var e Encoder

	got := e.EncodeInt(42)

	if string(got) != "42" {
		t.Fatalf("got %q, want %q", got, "42")
	}
}

func TestEncodeIntReturnsIndependentBytes(t *testing.T) {
	var e Encoder

	first := e.EncodeInt(8)
	second := e.EncodeInt(12)

	if string(first) != "8" {
		t.Fatalf("first = %q, want %q", first, "8")
	}
	if string(second) != "12" {
		t.Fatalf("second = %q, want %q", second, "12")
	}
}

func TestEncodeIntReturnedSliceCanBeMutatedByCaller(t *testing.T) {
	var e Encoder

	got := e.EncodeInt(9)
	got[0] = '1'

	again := e.EncodeInt(9)
	if string(again) != "9" {
		t.Fatalf("again = %q, want %q", again, "9")
	}
}

func TestEncodeIntDoesNotLeakPreviousLongerResult(t *testing.T) {
	var e Encoder

	_ = e.EncodeInt(42)
	got := e.EncodeInt(7)

	if string(got) != "7" {
		t.Fatalf("got %q, want %q", got, "7")
	}
}
