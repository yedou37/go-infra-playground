package jsoncodec

import (
	"strings"
	"testing"
)

func TestEncodeUsesTaggedFieldNames(t *testing.T) {
	got, err := Encode(LeaseSpec{
		Holder:     "scheduler",
		TTLSeconds: 15,
	})
	if err != nil {
		t.Fatalf("Encode returned error: %v", err)
	}

	want := `{"holder":"scheduler","ttlSeconds":15}`
	if string(got) != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestDecodeValidJSON(t *testing.T) {
	got, err := Decode([]byte(`{"holder":"controller","ttlSeconds":30,"labels":{"team":"api"}}`))
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if got.Holder != "controller" || got.TTLSeconds != 30 {
		t.Fatalf("got %+v, want holder=controller ttlSeconds=30", got)
	}
	if got.Labels["team"] != "api" {
		t.Fatalf("labels = %v, want team=api", got.Labels)
	}
}

func TestDecodeRejectsUnknownField(t *testing.T) {
	_, err := Decode([]byte(`{"holder":"controller","ttlSeconds":30,"unknown":true}`))
	if err == nil {
		t.Fatal("Decode should reject unknown fields")
	}
}

func TestDecodeRejectsTrailingGarbage(t *testing.T) {
	_, err := Decode([]byte(`{"holder":"controller","ttlSeconds":30} true`))
	if err == nil {
		t.Fatal("Decode should reject trailing garbage")
	}
}

func TestDecodeAllowsTrailingWhitespace(t *testing.T) {
	got, err := Decode([]byte("{\"holder\":\"controller\",\"ttlSeconds\":30}\n\t "))
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if got.Holder != "controller" || got.TTLSeconds != 30 {
		t.Fatalf("got %+v, want holder=controller ttlSeconds=30", got)
	}
}

func TestDecodeRejectsEmptyInput(t *testing.T) {
	_, err := Decode([]byte(strings.TrimSpace("   ")))
	if err == nil {
		t.Fatal("Decode should reject empty input")
	}
}
