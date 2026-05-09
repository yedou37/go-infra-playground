package deferorder

import (
	"reflect"
	"strings"
	"testing"
)

func TestDeferOrderIsLIFO(t *testing.T) {
	got := DeferOrder()
	want := []int{3, 2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestCapturedAtDeferTime(t *testing.T) {
	got := CapturedAtDeferTime()
	want := []int{3, 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestWrapErrorOnExitOnFailure(t *testing.T) {
	err := WrapErrorOnExit(true)
	if err == nil {
		t.Fatal("expected non-nil err")
	}
	if !strings.HasPrefix(err.Error(), "wrapped: ") {
		t.Fatalf("expected wrapped prefix, got %q", err.Error())
	}
}

func TestWrapErrorOnExitOnSuccess(t *testing.T) {
	if err := WrapErrorOnExit(false); err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestSafeCallNormalReturnsNil(t *testing.T) {
	if err := SafeCall(func() {}); err != nil {
		t.Fatalf("normal return should give nil, got %v", err)
	}
}

func TestSafeCallRecoversPanic(t *testing.T) {
	err := SafeCall(func() { panic("explode") })
	if err == nil {
		t.Fatal("expected non-nil err on panic")
	}
	if !strings.Contains(err.Error(), "panic:") {
		t.Fatalf("err should contain 'panic:', got %q", err.Error())
	}
	if !strings.Contains(err.Error(), "explode") {
		t.Fatalf("err should mention the panic value, got %q", err.Error())
	}
}
