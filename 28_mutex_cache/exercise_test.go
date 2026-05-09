package mutexcache

import "testing"

func TestCacheSetGet(t *testing.T) {
	var c Cache
	c.Set("a", "1")

	got, ok := c.Get("a")
	if !ok || got != "1" {
		t.Fatalf("Get = (%q, %v), want (%q, true)", got, ok, "1")
	}
}

func TestCacheMissingKey(t *testing.T) {
	var c Cache
	got, ok := c.Get("missing")
	if ok || got != "" {
		t.Fatalf("Get = (%q, %v), want (\"\", false)", got, ok)
	}
}

func TestSnapshotReturnsCopy(t *testing.T) {
	var c Cache
	c.Set("team", "api")

	snap := c.Snapshot()
	snap["team"] = "platform"

	again := c.Snapshot()
	if again["team"] != "api" {
		t.Fatalf("Snapshot returned aliased map: %v", again)
	}
}

func TestSnapshotPreservesNil(t *testing.T) {
	var c Cache
	if got := c.Snapshot(); got != nil {
		t.Fatalf("got %#v, want nil", got)
	}
}
