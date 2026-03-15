package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// State tracks which slugs have already been posted to dev.to and the
// corresponding article IDs, so re-runs are idempotent.
type State struct {
	Posted map[string]int `json:"posted"` // slug → dev.to article ID
}

// loadState reads the state file from disk. If the file does not exist an
// empty state is returned (not an error).
func loadState(path string) (*State, error) {
	s := &State{Posted: make(map[string]int)}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return s, nil
	}
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, s); err != nil {
		return nil, err
	}
	if s.Posted == nil {
		s.Posted = make(map[string]int)
	}
	return s, nil
}

// Has returns true if the slug has already been posted.
func (s *State) Has(slug string) bool {
	_, ok := s.Posted[slug]
	return ok
}

// Get returns the dev.to article ID for a slug (0 if not found).
func (s *State) Get(slug string) int {
	return s.Posted[slug]
}

// Set records slug → id in memory.
func (s *State) Set(slug string, id int) {
	s.Posted[slug] = id
}

// Save writes the state to disk as indented JSON.
func (s *State) Save(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}
