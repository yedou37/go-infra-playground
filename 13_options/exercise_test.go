package options

import (
	"testing"
	"time"
)

func TestNewConfigDefaults(t *testing.T) {
	got := NewConfig()
	if got.Name != "default" {
		t.Fatalf("Name = %q, want default", got.Name)
	}
	if got.Timeout != 5*time.Second {
		t.Fatalf("Timeout = %v, want 5s", got.Timeout)
	}
	if got.MaxInFlight != 1 {
		t.Fatalf("MaxInFlight = %d, want 1", got.MaxInFlight)
	}
}

func TestNewConfigAppliesOptionsInOrder(t *testing.T) {
	got := NewConfig(
		WithName("apiserver"),
		WithTimeout(2*time.Second),
		WithMaxInFlight(8),
		WithName("scheduler"),
	)
	if got.Name != "scheduler" {
		t.Fatalf("Name = %q, want scheduler", got.Name)
	}
	if got.Timeout != 2*time.Second {
		t.Fatalf("Timeout = %v, want 2s", got.Timeout)
	}
	if got.MaxInFlight != 8 {
		t.Fatalf("MaxInFlight = %d, want 8", got.MaxInFlight)
	}
}

func TestOptionsIgnoreInvalidValues(t *testing.T) {
	got := NewConfig(
		WithTimeout(0),
		WithTimeout(-1*time.Second),
		WithMaxInFlight(0),
		WithMaxInFlight(-3),
	)
	if got.Timeout != 5*time.Second {
		t.Fatalf("Timeout = %v, want default 5s", got.Timeout)
	}
	if got.MaxInFlight != 1 {
		t.Fatalf("MaxInFlight = %d, want default 1", got.MaxInFlight)
	}
}

func TestWithNameAllowsEmptyString(t *testing.T) {
	got := NewConfig(WithName(""))
	if got.Name != "" {
		t.Fatalf("Name = %q, want empty string", got.Name)
	}
}
