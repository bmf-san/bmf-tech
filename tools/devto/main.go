// devto is a CLI tool to cross-post bmf-tech.com articles to dev.to.
//
// Usage:
//
//	devto --api-key=<key> --all [--dry-run] [--delay=1200] [--state=posted.json]
//	devto --api-key=<key> --file=content/en/posts/my-post.md [--dry-run]
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	apiKey := flag.String("api-key", os.Getenv("DEV_TO_API_KEY"), "dev.to API key (or set DEV_TO_API_KEY env var)")
	all := flag.Bool("all", false, "post all non-draft articles in content/en/posts/")
	file := flag.String("file", "", "path to a single markdown file to post")
	dryRun := flag.Bool("dry-run", false, "parse and print what would be posted without calling the API")
	stateFile := flag.String("state", "tools/devto/posted.json", "path to the state file tracking already-posted slugs")
	delayMs := flag.Int("delay", 1200, "delay in milliseconds between API calls (rate limit is ~30 req/30 sec)")
	flag.Parse()

	if !*dryRun && *apiKey == "" {
		log.Fatal("--api-key is required (or set DEV_TO_API_KEY env var)")
	}
	if !*all && *file == "" {
		log.Fatal("either --all or --file=<path> must be specified")
	}

	state, err := loadState(*stateFile)
	if err != nil {
		log.Fatalf("failed to load state file: %v", err)
	}

	poster := &Poster{
		APIKey: *apiKey,
		DryRun: *dryRun,
	}

	if *file != "" {
		if err := postFile(*file, poster, state, *stateFile); err != nil {
			log.Fatalf("failed to post %s: %v", *file, err)
		}
		return
	}

	// --all: glob all .md files in content/en/posts/
	pattern := filepath.Join("content", "en", "posts", "*.md")
	files, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatalf("failed to glob %s: %v", pattern, err)
	}

	succeeded := 0
	skipped := 0
	failed := 0
	for i, f := range files {
		slug := slugFromPath(f)
		if state.Has(slug) {
			fmt.Printf("[skip] %s (already posted as article %d)\n", slug, state.Get(slug))
			skipped++
			continue
		}

		err := postFile(f, poster, state, *stateFile)
		if err != nil {
			log.Printf("[error] %s: %v", f, err)
			failed++
		} else {
			succeeded++
		}

		if i < len(files)-1 && !*dryRun {
			time.Sleep(time.Duration(*delayMs) * time.Millisecond)
		}
	}

	fmt.Printf("\nDone. posted=%d skipped=%d failed=%d\n", succeeded, skipped, failed)
}

func postFile(path string, poster *Poster, state *State, stateFile string) error {
	article, err := ParseArticle(path)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}
	if article == nil {
		fmt.Printf("[skip] %s (draft)\n", path)
		return nil
	}

	if poster.DryRun {
		fmt.Printf("[dry-run] %s\n  title:         %s\n  canonical_url: %s\n  tags:          %v\n  body (first 120 chars): %.120s\n\n",
			path, article.Title, article.CanonicalURL, article.Tags, article.Body)
		return nil
	}

	id, err := poster.Post(article)
	if err != nil {
		return fmt.Errorf("post: %w", err)
	}

	state.Set(article.Slug, id)
	if err := state.Save(stateFile); err != nil {
		log.Printf("[warn] could not save state: %v", err)
	}
	fmt.Printf("[posted] %s → dev.to article %d\n", article.Slug, id)
	return nil
}
