package errwrap

import (
	"errors"
	"strings"
	"testing"
)

func TestLookupNotFoundIsWrapped(t *testing.T) {
	err := LookupAndAnnotate("foo")
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("errors.Is(err, ErrNotFound) must be true; got err=%v", err)
	}
	if !strings.Contains(err.Error(), "foo") {
		t.Fatalf("annotated err should contain the key %q; got %q", "foo", err.Error())
	}
}

func TestLookupEmptyIsNotNotFound(t *testing.T) {
	err := LookupAndAnnotate("")
	if err == nil {
		t.Fatal("expected error")
	}
	if errors.Is(err, ErrNotFound) {
		t.Fatalf("empty-key error must NOT wrap ErrNotFound; got %v", err)
	}
}

func TestIsNotFoundOnNil(t *testing.T) {
	if IsNotFound(nil) {
		t.Fatal("IsNotFound(nil) must be false")
	}
}

func TestQuotaExceededIsExtractable(t *testing.T) {
	err := CheckQuota("pods", 10, 12)
	if err == nil {
		t.Fatal("expected error")
	}
	qe, ok := AsQuotaExceeded(err)
	if !ok || qe == nil {
		t.Fatalf("AsQuotaExceeded should extract *QuotaExceeded; got ok=%v qe=%v", ok, qe)
	}
	if qe.Resource != "pods" || qe.Limit != 10 || qe.Got != 12 {
		t.Fatalf("extracted struct mismatch: %+v", qe)
	}
}

func TestCheckQuotaNoViolation(t *testing.T) {
	if err := CheckQuota("pods", 10, 5); err != nil {
		t.Fatalf("under-limit must return nil; got %v", err)
	}
}

func TestAsQuotaExceededOnUnrelatedError(t *testing.T) {
	if _, ok := AsQuotaExceeded(errors.New("nope")); ok {
		t.Fatal("AsQuotaExceeded on unrelated error must return false")
	}
}

func TestQuotaExceededErrorMentionsFields(t *testing.T) {
	qe := &QuotaExceeded{Resource: "cpu", Limit: 4, Got: 7}
	msg := qe.Error()
	for _, sub := range []string{"cpu", "4", "7"} {
		if !strings.Contains(msg, sub) {
			t.Fatalf("Error() should mention %q; got %q", sub, msg)
		}
	}
}
