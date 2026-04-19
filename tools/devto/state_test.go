package main

import (
	"path/filepath"
	"testing"
)

func TestStateRoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "nested", "posted.json")

	// Fresh load on a missing file returns empty state.
	s, err := loadState(path)
	if err != nil {
		t.Fatalf("loadState: %v", err)
	}
	if s == nil || s.Posted == nil {
		t.Fatal("expected non-nil state")
	}
	if s.Has("anything") {
		t.Error("expected empty state to not contain keys")
	}

	// Populate and save.
	s.Set("foo", 1)
	s.Set("bar", 2)
	if err := s.Save(path); err != nil {
		t.Fatalf("Save: %v", err)
	}

	// Reload and verify contents.
	s2, err := loadState(path)
	if err != nil {
		t.Fatalf("reload: %v", err)
	}
	if !s2.Has("foo") || s2.Get("foo") != 1 {
		t.Errorf("foo missing or wrong id: %+v", s2)
	}
	if s2.Get("bar") != 2 {
		t.Errorf("bar wrong id: %d", s2.Get("bar"))
	}
	if s2.Has("unknown") {
		t.Error("unknown key should not be present")
	}
	if s2.Get("unknown") != 0 {
		t.Error("missing key should yield zero value")
	}
}
