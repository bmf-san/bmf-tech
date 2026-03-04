// migrate converts a gobel-api MySQL dump into gohan-flavoured Markdown files.
//
// It reads:
//   - A MySQL dump (--sql)   ← bmf-tech_2026-03-01.sql
//   - A slug map CSV (--csv) ← tools/slug_map.csv  (id,category_id,title_ja,slug,created_at,reviewed)
//
// And produces:
//   - One Markdown file per published post under --out-dir (default: ../../content/ja/posts/)
//   - A Cloudflare Pages _redirects file under --redirects (default: ../../_redirects)
//
// Usage:
//
//	go run main.go \
//	  [-sql   ../../bmf-tech_2026-03-01.sql] \
//	  [-csv   ../../tools/slug_map.csv] \
//	  [-out   ../../content/ja/posts] \
//	  [-redir ../../_redirects]
package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ─── Data types ──────────────────────────────────────────────────────────────

type post struct {
	id         string
	categoryID string
	title      string
	body       string // raw Markdown from md_body
	status     string // "publish", "draft", etc.
	createdAt  time.Time
}

type slugEntry struct {
	id   string
	slug string
}

// ─── Main ─────────────────────────────────────────────────────────────────────

func main() {
	sqlPath := flag.String("sql", "../../bmf-tech_2026-03-01.sql", "MySQL dump path")
	csvPath := flag.String("csv", "../../tools/slug_map.csv", "slug_map.csv path")
	outDir := flag.String("out", "../../content/ja/posts", "output directory for Markdown files")
	redirPath := flag.String("redir", "../../_redirects", "_redirects file path")
	flag.Parse()

	// ── Load slug map ──────────────────────────────────────────────────────────
	slugMap, err := loadSlugMap(*csvPath)
	if err != nil {
		log.Fatalf("load slug map: %v", err)
	}
	fmt.Printf("slug map: %d entries\n", len(slugMap))

	// ── Load posts ─────────────────────────────────────────────────────────────
	catMap, tagMap, tagPostMap, posts, err := parseDump(*sqlPath)
	if err != nil {
		log.Fatalf("parse dump: %v", err)
	}
	fmt.Printf("categories: %d, tags: %d, posts: %d\n", len(catMap), len(tagMap), len(posts))

	// ── Write Markdown files ───────────────────────────────────────────────────
	if err := os.MkdirAll(*outDir, 0o755); err != nil {
		log.Fatalf("mkdir: %v", err)
	}

	var redirects []string
	written, skipped := 0, 0

	for _, p := range posts {
		slug, ok := slugMap[p.id]
		if !ok || slug == "" {
			slug = "post-" + p.id
		}

		catName := catMap[p.categoryID]
		tags := tagPostMap[p.id] // []string of tag names
		tagNames := expandNames(tags, tagMap)

		fm := buildFrontMatter(p, slug, catName, tagNames)
		mdContent := fm + "\n" + p.body + "\n"

		filename := slug + ".md"
		path := filepath.Join(*outDir, filename)
		if err := os.WriteFile(path, []byte(mdContent), 0o644); err != nil {
			log.Printf("write %s: %v", path, err)
			skipped++
			continue
		}
		written++

		// Old gobel URL: /posts/{title-slug} (title was URL-encoded Japanese).
		// We redirect from the numeric id path as a safe catch-all since we
		// don't know the exact old slugs.  Individual known redirects can be
		// added manually to _redirects.
		oldURL := fmt.Sprintf("/posts/%s", p.id)
		newURL := fmt.Sprintf("/ja/posts/%s/", slug)
		redirects = append(redirects, fmt.Sprintf("%s  %s  301", oldURL, newURL))
	}

	fmt.Printf("written: %d, skipped: %d\n", written, skipped)

	// ── Write _redirects ───────────────────────────────────────────────────────
	if err := writeRedirects(*redirPath, redirects); err != nil {
		log.Fatalf("write redirects: %v", err)
	}
	fmt.Printf("wrote %s (%d redirects)\n", *redirPath, len(redirects))
}

// ─── Helpers ──────────────────────────────────────────────────────────────────

func loadSlugMap(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	header, err := r.Read()
	if err != nil {
		return nil, err
	}
	// Find column indices by header name.
	idIdx, slugIdx := -1, -1
	for i, h := range header {
		switch h {
		case "id":
			idIdx = i
		case "slug":
			slugIdx = i
		}
	}
	if idIdx < 0 || slugIdx < 0 {
		return nil, fmt.Errorf("slug map CSV must have 'id' and 'slug' columns")
	}

	m := make(map[string]string)
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if idIdx < len(row) && slugIdx < len(row) {
			m[row[idIdx]] = row[slugIdx]
		}
	}
	return m, nil
}

// parseDump reads categories, tags, tag_post associations, and posts from a MySQL dump.
func parseDump(path string) (
	catMap map[string]string, // category_id → name
	tagMap map[string]string, // tag_id → name
	tagPostMap map[string][]string, // post_id → []tag_id
	posts []post,
	err error,
) {
	catMap = make(map[string]string)
	tagMap = make(map[string]string)
	tagPostMap = make(map[string][]string)

	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	buf := make([]byte, 128*1024*1024)
	scanner.Buffer(buf, cap(buf))

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "INSERT INTO `categories` VALUES"):
			parseCategories(line, catMap)
		case strings.HasPrefix(line, "INSERT INTO `tags` VALUES"):
			parseTags(line, tagMap)
		case strings.HasPrefix(line, "INSERT INTO `tag_post` VALUES"):
			parseTagPost(line, tagPostMap)
		case strings.HasPrefix(line, "INSERT INTO `posts` VALUES"):
			pp := parsePosts(line)
			posts = append(posts, pp...)
		}
	}
	err = scanner.Err()
	return
}

// ─── INSERT parsers ───────────────────────────────────────────────────────────

// parseCategories parses: (id,'name','ts','ts'),...
func parseCategories(line string, m map[string]string) {
	pos := valueStart(line)
	for pos < len(line) {
		if line[pos] != '(' {
			pos++
			continue
		}
		pos++
		id, p := readInt(line, pos)
		pos = commaSkip(line, p)
		name, p := readString(line, pos+1) // skip opening '
		pos = p
		m[id] = name
		pos = skipToTupleEnd(line, pos)
		if pos < len(line) && line[pos] == ',' {
			pos++
		}
	}
}

// parseTags parses: (id,'name','ts','ts'),...
func parseTags(line string, m map[string]string) {
	pos := valueStart(line)
	for pos < len(line) {
		if line[pos] != '(' {
			pos++
			continue
		}
		pos++
		id, p := readInt(line, pos)
		pos = commaSkip(line, p)
		name, p := readString(line, pos+1)
		pos = p
		m[id] = name
		pos = skipToTupleEnd(line, pos)
		if pos < len(line) && line[pos] == ',' {
			pos++
		}
	}
}

// parseTagPost parses: (id,tag_id,post_id,'ts','ts'),...
func parseTagPost(line string, m map[string][]string) {
	pos := valueStart(line)
	for pos < len(line) {
		if line[pos] != '(' {
			pos++
			continue
		}
		pos++
		_, p := readInt(line, pos) // id
		pos = commaSkip(line, p)
		tagID, p := readInt(line, pos)
		pos = commaSkip(line, p)
		postID, p := readInt(line, pos)
		pos = p
		m[postID] = append(m[postID], tagID)
		pos = skipToTupleEnd(line, pos)
		if pos < len(line) && line[pos] == ',' {
			pos++
		}
	}
}

// parsePosts parses all (id,admin_id,category_id,'title','md_body','html_body','status','created_at','updated_at') tuples.
func parsePosts(line string) []post {
	var posts []post
	pos := valueStart(line)
	for pos < len(line) {
		if line[pos] != '(' {
			pos++
			continue
		}
		pos++

		id, p := readInt(line, pos)
		pos = commaSkip(line, p)
		_, p = readInt(line, pos) // admin_id
		pos = commaSkip(line, p)
		catID, p := readInt(line, pos)
		pos = commaSkip(line, p)

		title := ""
		if pos < len(line) && line[pos] == '\'' {
			title, p = readString(line, pos+1)
			pos = p
		} else if strings.HasPrefix(line[pos:], "NULL") {
			pos += 4
		}

		pos = commaSkip(line, pos)

		body := ""
		if pos < len(line) && line[pos] == '\'' {
			body, p = readString(line, pos+1)
			pos = p
		} else if strings.HasPrefix(line[pos:], "NULL") {
			pos += 4
		}

		// skip html_body
		pos = commaSkip(line, pos)
		pos = skipField(line, pos)

		// status
		pos = commaSkip(line, pos)
		status := ""
		if pos < len(line) && line[pos] == '\'' {
			status, p = readString(line, pos+1)
			pos = p
		} else if strings.HasPrefix(line[pos:], "NULL") {
			pos += 4
		}

		// created_at
		pos = commaSkip(line, pos)
		createdAtStr := ""
		if pos < len(line) && line[pos] == '\'' {
			createdAtStr, p = readString(line, pos+1)
			pos = p
		}

		createdAt, _ := time.Parse("2006-01-02 15:04:05", createdAtStr)

		title = strings.Map(func(r rune) rune {
			if r == '\n' || r == '\r' {
				return ' '
			}
			return r
		}, title)

		posts = append(posts, post{
			id:         id,
			categoryID: catID,
			title:      title,
			body:       body,
			status:     status,
			createdAt:  createdAt,
		})

		pos = skipToTupleEnd(line, pos)
		if pos < len(line) && line[pos] == ',' {
			pos++
		}
	}
	return posts
}

// ─── Front matter builder ─────────────────────────────────────────────────────

func buildFrontMatter(p post, slug, catName string, tagNames []string) string {
	var sb strings.Builder
	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("title: %q\n", p.title))
	sb.WriteString(fmt.Sprintf("slug: %q\n", slug))
	if !p.createdAt.IsZero() {
		sb.WriteString(fmt.Sprintf("date: %s\n", p.createdAt.Format("2006-01-02")))
	}
	sb.WriteString(fmt.Sprintf("author: bmf-san\n"))
	if catName != "" {
		sb.WriteString("categories:\n")
		sb.WriteString(fmt.Sprintf("  - %q\n", catName))
	}
	if len(tagNames) > 0 {
		sb.WriteString("tags:\n")
		for _, t := range tagNames {
			sb.WriteString(fmt.Sprintf("  - %q\n", t))
		}
	}
	if p.status == "draft" {
		sb.WriteString("draft: true\n")
	} else {
		sb.WriteString("draft: false\n")
	}
	sb.WriteString("---\n")
	return sb.String()
}

// ─── _redirects writer ────────────────────────────────────────────────────────

func writeRedirects(path string, redirects []string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	fmt.Fprintln(w, "# Cloudflare Pages _redirects")
	fmt.Fprintln(w, "# Format: /old-path  /new-path  301")
	fmt.Fprintln(w)
	for _, r := range redirects {
		fmt.Fprintln(w, r)
	}
	return w.Flush()
}

// ─── Low-level string helpers (shared with slug generator) ───────────────────

func valueStart(line string) int {
	vi := strings.Index(line, " VALUES ")
	if vi < 0 {
		return len(line)
	}
	return vi + len(" VALUES ")
}

func readInt(s string, pos int) (string, int) {
	start := pos
	for pos < len(s) && s[pos] >= '0' && s[pos] <= '9' {
		pos++
	}
	return s[start:pos], pos
}

func commaSkip(s string, pos int) int {
	if pos < len(s) && s[pos] == ',' {
		return pos + 1
	}
	return pos
}

// readString parses a MySQL single-quoted string starting after the opening quote.
func readString(s string, pos int) (string, int) {
	var b strings.Builder
	for pos < len(s) {
		c := s[pos]
		if c == '\\' && pos+1 < len(s) {
			nc := s[pos+1]
			switch nc {
			case '\'':
				b.WriteByte('\'')
			case '\\':
				b.WriteByte('\\')
			case 'n':
				b.WriteByte('\n')
			case 'r':
				b.WriteByte('\r')
			case '"':
				b.WriteByte('"')
			default:
				b.WriteByte('\\')
				b.WriteByte(nc)
			}
			pos += 2
		} else if c == '\'' {
			if pos+1 < len(s) && s[pos+1] == '\'' {
				b.WriteByte('\'')
				pos += 2
			} else {
				return b.String(), pos + 1
			}
		} else {
			b.WriteByte(c)
			pos++
		}
	}
	return b.String(), pos
}

// skipField advances past a NULL or single-quoted string field.
func skipField(s string, pos int) int {
	if pos >= len(s) {
		return pos
	}
	if s[pos] == '\'' {
		_, p := readString(s, pos+1)
		return p
	}
	if strings.HasPrefix(s[pos:], "NULL") {
		return pos + 4
	}
	return pos
}

// skipToTupleEnd advances pos to the ')' closing the current tuple.
func skipToTupleEnd(s string, pos int) int {
	for pos < len(s) && s[pos] != ')' {
		if s[pos] == '\'' {
			_, pos = readString(s, pos+1)
		} else {
			pos++
		}
	}
	if pos < len(s) {
		pos++ // consume ')'
	}
	return pos
}

// expandNames converts a slice of tag IDs to tag names using tagMap.
func expandNames(ids []string, m map[string]string) []string {
	names := make([]string, 0, len(ids))
	seen := make(map[string]bool)
	for _, id := range ids {
		name := m[id]
		if name == "" || seen[name] {
			continue
		}
		seen[name] = true
		names = append(names, name)
	}
	return names
}
