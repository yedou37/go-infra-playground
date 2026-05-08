package mapownership

import "testing"

func TestSetLabelsClonesInput(t *testing.T) {
	src := map[string]string{"team": "api"}
	var s Store

	s.SetLabels(src)
	src["team"] = "platform"

	got := s.Labels()
	if got["team"] != "api" {
		t.Fatalf("store observed caller mutation: %v", got)
	}
}

func TestLabelsReturnsCopy(t *testing.T) {
	var s Store
	s.SetLabels(map[string]string{"env": "prod"})

	got := s.Labels()
	got["env"] = "dev"

	again := s.Labels()
	if again["env"] != "prod" {
		t.Fatalf("Labels returned aliased map: %v", again)
	}
}

func TestSetAllocatesMapIfNeeded(t *testing.T) {
	var s Store
	s.Set("zone", "us-east-1")

	got := s.Labels()
	if got["zone"] != "us-east-1" {
		t.Fatalf("got %v, want zone=us-east-1", got)
	}
}

func TestPreservesNil(t *testing.T) {
	var s Store
	s.SetLabels(nil)
	if got := s.Labels(); got != nil {
		t.Fatalf("got %#v, want nil", got)
	}
}
