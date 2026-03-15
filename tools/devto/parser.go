package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

const baseURL = "https://bmf-tech.com"

// maxDevToTags is the maximum number of tags dev.to allows per article.
const maxDevToTags = 4

// Article holds the data needed to create a dev.to article.
type Article struct {
	Slug         string
	Title        string
	Description  string
	CanonicalURL string
	Tags         []string
	Body         string // includes the prepended origin notice
}

type frontmatter struct {
	Title       string   `yaml:"title"`
	Slug        string   `yaml:"slug"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags"`
	Draft       bool     `yaml:"draft"`
}

var fmSep = []byte("---")

// ParseArticle reads a Hugo markdown file and returns an Article ready for
// posting to dev.to. Returns nil (no error) when the article is a draft.
func ParseArticle(path string) (*Article, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	fm, body, err := splitFrontmatter(raw)
	if err != nil {
		return nil, fmt.Errorf("split frontmatter: %w", err)
	}

	var meta frontmatter
	if err := yaml.Unmarshal(fm, &meta); err != nil {
		return nil, fmt.Errorf("parse yaml: %w", err)
	}

	if meta.Draft {
		return nil, nil
	}

	slug := meta.Slug
	if slug == "" {
		slug = slugFromPath(path)
	}

	canonicalURL := fmt.Sprintf("%s/en/posts/%s", baseURL, slug)
	notice := fmt.Sprintf("> This article was originally published on [bmf-tech.com](%s).\n\n", canonicalURL)
	fullBody := notice + strings.TrimSpace(body)

	return &Article{
		Slug:         slug,
		Title:        meta.Title,
		Description:  meta.Description,
		CanonicalURL: canonicalURL,
		Tags:         normalizeTags(meta.Tags),
		Body:         fullBody,
	}, nil
}

// splitFrontmatter splits a Hugo markdown file into YAML frontmatter bytes
// and the body string.
func splitFrontmatter(raw []byte) ([]byte, string, error) {
	// Expect the file to start with ---
	if !bytes.HasPrefix(bytes.TrimLeft(raw, "\n"), fmSep) {
		return nil, "", fmt.Errorf("no frontmatter found")
	}

	// Find the closing ---
	content := bytes.TrimLeft(raw, "\n")
	// Strip the opening ---\n
	after := content[len(fmSep):]
	if len(after) > 0 && after[0] == '\n' {
		after = after[1:]
	}

	idx := bytes.Index(after, []byte("\n---"))
	if idx == -1 {
		return nil, "", fmt.Errorf("frontmatter closing --- not found")
	}

	fmBytes := after[:idx]
	bodySrc := after[idx+4:] // skip \n---
	// Skip optional \n after closing ---
	if len(bodySrc) > 0 && bodySrc[0] == '\n' {
		bodySrc = bodySrc[1:]
	}

	return fmBytes, string(bodySrc), nil
}

// slugFromPath derives a slug from the file basename without extension.
func slugFromPath(path string) string {
	base := filepath.Base(path)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

var nonAlphaNum = regexp.MustCompile(`[^a-z0-9]`)

// normalizeTags converts Hugo tags to dev.to-compatible tags (lowercase,
// alphanumeric only, max 4 tags).
func normalizeTags(tags []string) []string {
	seen := make(map[string]bool)
	out := make([]string, 0, maxDevToTags)
	for _, t := range tags {
		normalized := nonAlphaNum.ReplaceAllString(strings.ToLower(t), "")
		if normalized == "" || seen[normalized] {
			continue
		}
		seen[normalized] = true
		out = append(out, normalized)
		if len(out) == maxDevToTags {
			break
		}
	}
	return out
}
