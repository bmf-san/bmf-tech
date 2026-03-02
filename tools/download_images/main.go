// download_images scans Markdown files in a content directory for external
// image URLs, downloads them into a local assets directory, and rewrites the
// Markdown source files to use the local paths.
//
// Usage:
//
//	go run main.go [-content PATH] [-assets PATH] [-dry-run]
//
// Defaults:
//
//	-content  ../../content/ja/posts
//	-assets   ../../assets/images/posts
//
// External images are saved as:
//
//	{assets}/{slug}/{filename}
//
// and the Markdown URLs are rewritten to:
//
//	/assets/images/posts/{slug}/{filename}
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// imgRe matches http(s):// URLs ending in a known image extension.
var imgRe = regexp.MustCompile(
	`(?i)(https?://[^\s)"']+\.(?:png|jpg|jpeg|gif|webp|svg))`,
)

func main() {
	contentDir := flag.String("content", "../../content/ja/posts", "directory of Markdown files to process")
	assetsDir := flag.String("assets", "../../assets/images/posts", "directory to save downloaded images")
	dryRun := flag.Bool("dry-run", false, "print actions without writing files")
	flag.Parse()

	entries, err := os.ReadDir(*contentDir)
	if err != nil {
		log.Fatalf("read content dir: %v", err)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	downloaded, skipped, unchanged := 0, 0, 0

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		filePath := filepath.Join(*contentDir, entry.Name())
		slug := strings.TrimSuffix(entry.Name(), ".md")

		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("read %s: %v", entry.Name(), err)
			continue
		}
		content := string(data)

		newContent, count, skip := processFile(content, slug, *assetsDir, client, *dryRun)
		downloaded += count
		skipped += skip

		if newContent != content {
			if !*dryRun {
				if err := os.WriteFile(filePath, []byte(newContent), 0o644); err != nil {
					log.Printf("write %s: %v", filePath, err)
					continue
				}
			}
			fmt.Printf("updated: %s (%d images)\n", entry.Name(), count)
		} else {
			unchanged++
		}
	}

	fmt.Printf("\ndownloaded: %d, skipped: %d, unchanged files: %d\n", downloaded, skipped, unchanged)
}

// processFile finds all external image URLs in content, downloads each image
// (unless already present), and returns the updated content.
func processFile(content, slug, assetsDir string, client *http.Client, dryRun bool) (string, int, int) {
	downloaded, skipped := 0, 0

	newContent := imgRe.ReplaceAllStringFunc(content, func(rawURL string) string {
		// Skip already-local paths.
		if strings.HasPrefix(rawURL, "/") || strings.HasPrefix(rawURL, "./") {
			return rawURL
		}

		filename := urlFilename(rawURL)
		if filename == "" {
			skipped++
			return rawURL
		}

		localDir := filepath.Join(assetsDir, slug)
		localPath := filepath.Join(localDir, filename)
		localURL := fmt.Sprintf("/assets/images/posts/%s/%s", slug, filename)

		// Already downloaded?
		if _, err := os.Stat(localPath); err == nil {
			downloaded++
			return localURL
		}

		if dryRun {
			fmt.Printf("  [dry-run] would download: %s → %s\n", rawURL, localURL)
			downloaded++
			return localURL
		}

		if err := os.MkdirAll(localDir, 0o755); err != nil {
			log.Printf("mkdir %s: %v", localDir, err)
			skipped++
			return rawURL
		}

		if err := downloadFile(client, rawURL, localPath); err != nil {
			log.Printf("download %s: %v", rawURL, err)
			skipped++
			return rawURL
		}

		downloaded++
		return localURL
	})

	return newContent, downloaded, skipped
}

// urlFilename extracts a safe filename from a URL.
func urlFilename(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	base := filepath.Base(u.Path)
	// Sanitize: keep only safe characters.
	var sb strings.Builder
	for _, r := range base {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' ||
			r == '-' || r == '_' || r == '.' {
			sb.WriteRune(r)
		} else {
			sb.WriteRune('_')
		}
	}
	s := sb.String()
	if s == "" || s == "." {
		return ""
	}
	return s
}

// downloadFile fetches url and writes to dest.
func downloadFile(client *http.Client, rawURL, dest string) error {
	resp, err := client.Get(rawURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d for %s", resp.StatusCode, rawURL)
	}

	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}
