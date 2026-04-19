package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestNormalizeTags(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want []string
	}{
		{"empty", nil, []string{}},
		{"lowercased", []string{"Go", "Docker"}, []string{"go", "docker"}},
		{"strip_non_alnum", []string{"C++", "C#", "Node.js"}, []string{"c", "nodejs"}},
		{"dedupe", []string{"Go", "GO", "go"}, []string{"go"}},
		{"cap_at_four", []string{"a", "b", "c", "d", "e", "f"}, []string{"a", "b", "c", "d"}},
		{"japanese_drops_to_empty_then_skipped", []string{"日本語", "go"}, []string{"go"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeTags(tt.in)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("normalizeTags(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestSlugFromPath(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"content/en/posts/foo-bar.md", "foo-bar"},
		{"/abs/path/baz.md", "baz"},
		{"no-ext", "no-ext"},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			if got := slugFromPath(tt.path); got != tt.want {
				t.Errorf("slugFromPath(%q) = %q, want %q", tt.path, got, tt.want)
			}
		})
	}
}

func TestSplitFrontmatter(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		in := []byte("---\ntitle: Hi\n---\nbody here\n")
		fm, body, err := splitFrontmatter(in)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if string(fm) != "title: Hi" {
			t.Errorf("frontmatter = %q", fm)
		}
		if body != "body here\n" {
			t.Errorf("body = %q", body)
		}
	})

	t.Run("no_frontmatter", func(t *testing.T) {
		_, _, err := splitFrontmatter([]byte("no fm\n"))
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("unterminated_frontmatter", func(t *testing.T) {
		_, _, err := splitFrontmatter([]byte("---\ntitle: Hi\n"))
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func writeTemp(t *testing.T, name, content string) string {
	t.Helper()
	dir := t.TempDir()
	p := filepath.Join(dir, name)
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	return p
}

func TestParseArticle(t *testing.T) {
	t.Run("happy_path", func(t *testing.T) {
		path := writeTemp(t, "my-post.md", `---
title: My Post
slug: my-post
description: A post
tags:
  - Go
  - Docker
---
Hello ![pic](/assets/pic.png) world.
`)
		a, err := ParseArticle(path)
		if err != nil {
			t.Fatalf("parse: %v", err)
		}
		if a == nil {
			t.Fatal("expected non-nil article")
		}
		if a.Slug != "my-post" || a.Title != "My Post" {
			t.Errorf("unexpected meta: %+v", a)
		}
		if a.CanonicalURL != "https://bmf-tech.com/posts/my-post/" {
			t.Errorf("canonical URL: %s", a.CanonicalURL)
		}
		if a.MainImage != "https://bmf-tech.com/ogp/my-post.png" {
			t.Errorf("main image: %s", a.MainImage)
		}
		want := []string{"go", "docker"}
		if !reflect.DeepEqual(a.Tags, want) {
			t.Errorf("tags = %v, want %v", a.Tags, want)
		}
		// absolute-path image is rewritten
		if !containsAll(a.Body, []string{
			"This article was originally published",
			"https://bmf-tech.com/posts/my-post/",
			"![pic](https://bmf-tech.com/assets/pic.png)",
		}) {
			t.Errorf("body missing expected pieces:\n%s", a.Body)
		}
	})

	t.Run("draft_returns_nil", func(t *testing.T) {
		path := writeTemp(t, "draft.md", `---
title: Draft
draft: true
---
hi
`)
		a, err := ParseArticle(path)
		if err != nil {
			t.Fatalf("parse: %v", err)
		}
		if a != nil {
			t.Errorf("expected nil for draft, got %+v", a)
		}
	})

	t.Run("poem_category_skipped", func(t *testing.T) {
		path := writeTemp(t, "poem.md", `---
title: Poem
slug: p
categories:
  - Poem
---
`)
		a, err := ParseArticle(path)
		if err != nil {
			t.Fatalf("parse: %v", err)
		}
		if a != nil {
			t.Errorf("expected nil for poem category, got %+v", a)
		}
	})

	t.Run("slug_defaults_to_filename", func(t *testing.T) {
		path := writeTemp(t, "fallback-slug.md", `---
title: X
---
`)
		a, err := ParseArticle(path)
		if err != nil {
			t.Fatalf("parse: %v", err)
		}
		if a == nil || a.Slug != "fallback-slug" {
			t.Errorf("expected slug from filename, got %+v", a)
		}
	})
}

func containsAll(s string, subs []string) bool {
	for _, sub := range subs {
		if !containsString(s, sub) {
			return false
		}
	}
	return true
}

func containsString(s, sub string) bool {
	return len(s) >= len(sub) && (func() bool {
		for i := 0; i+len(sub) <= len(s); i++ {
			if s[i:i+len(sub)] == sub {
				return true
			}
		}
		return false
	})()
}
