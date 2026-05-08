package customjson

import (
	"encoding/json"
	"testing"
	"time"
)

func TestPolicyMarshalJSON(t *testing.T) {
	got, err := json.Marshal(Policy{
		Level:   LevelWarn,
		Timeout: MillisDuration(1500 * time.Millisecond),
	})
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if string(got) != `{"level":"warn","timeoutMillis":1500}` {
		t.Fatalf("got %s, want %s", got, `{"level":"warn","timeoutMillis":1500}`)
	}
}

func TestPolicyUnmarshalJSON(t *testing.T) {
	var got Policy
	err := json.Unmarshal([]byte(`{"level":"error","timeoutMillis":250}`), &got)
	if err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if got.Level != LevelError {
		t.Fatalf("Level = %q, want error", got.Level)
	}
	if time.Duration(got.Timeout) != 250*time.Millisecond {
		t.Fatalf("Timeout = %v, want 250ms", time.Duration(got.Timeout))
	}
}

func TestLevelRejectsUnknownValue(t *testing.T) {
	var got Level
	if err := json.Unmarshal([]byte(`"debug"`), &got); err == nil {
		t.Fatal("Unmarshal should reject unknown enum values")
	}
}

func TestMillisDurationRejectsNegativeValue(t *testing.T) {
	var got MillisDuration
	if err := json.Unmarshal([]byte(`-1`), &got); err == nil {
		t.Fatal("Unmarshal should reject negative durations")
	}
}
