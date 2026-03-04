// slug_map_generator reads a MySQL dump and generates a CSV of post IDs +
// Japanese titles + English slug candidates for human review.
//
// Usage:
//
//	go run main.go [-sql PATH] [-out PATH] [-ai]
//
// Defaults:
//
//	-sql  ../../bmf-tech_2026-03-01.sql
//	-out  ../../tools/slug_map.csv
//	-ai   false (set true to use OpenAI GPT-4o-mini; requires OPENAI_API_KEY env)
package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"unicode"
)

type postRecord struct {
	id, categoryID, title, createdAt, slug string
}

func main() {
	sqlPath := flag.String("sql", "../../bmf-tech_2026-03-01.sql", "MySQL dump path")
	outPath := flag.String("out", "../../tools/slug_map.csv", "output CSV path")
	useAI := flag.Bool("ai", false, "use OpenAI GPT-4o-mini to generate meaningful English slugs (requires OPENAI_API_KEY)")
	flag.Parse()

	records, err := parsePosts(*sqlPath)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}
	fmt.Printf("parsed %d post records\n", len(records))

	if *useAI {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("OPENAI_API_KEY is not set")
		}
		fmt.Println("generating slugs with OpenAI GPT-4o-mini...")
		if err := assignAISlugs(records, apiKey); err != nil {
			log.Fatalf("ai slug generation: %v", err)
		}
	}

	if err := writeCSV(*outPath, records); err != nil {
		log.Fatalf("write csv: %v", err)
	}
	fmt.Printf("wrote %s\n", *outPath)
}

func parsePosts(path string) ([]postRecord, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	buf := make([]byte, 128*1024*1024) // 128 MB for large longtext fields
	scanner.Buffer(buf, cap(buf))

	var records []postRecord
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "INSERT INTO `posts` VALUES") {
			continue
		}
		records = append(records, parseTuples(line)...)
	}
	return records, scanner.Err()
}

// parseTuples extracts all (id,admin_id,cat_id,'title',...) rows from a bulk INSERT line.
func parseTuples(line string) []postRecord {
	var records []postRecord
	vi := strings.Index(line, " VALUES ")
	if vi < 0 {
		return nil
	}
	pos := vi + len(" VALUES ")

	for pos < len(line) {
		if line[pos] != '(' {
			pos++
			continue
		}
		pos++ // consume '('

		id, p := readInt(line, pos)
		pos = commaSkip(line, p)
		_, p = readInt(line, pos) // admin_id
		pos = commaSkip(line, p)
		catID, p := readInt(line, pos)
		pos = commaSkip(line, p)

		var title string
		if pos < len(line) && line[pos] == '\'' {
			title, p = readString(line, pos+1)
			pos = p
		} else if strings.HasPrefix(line[pos:], "NULL") {
			pos += 4
		}
		title = strings.Map(func(r rune) rune {
			if r == '\n' || r == '\r' {
				return ' '
			}
			return r
		}, title)

		// Skip md_body, html_body, status (3 fields)
		for i := 0; i < 3; i++ {
			pos = commaSkip(line, pos)
			pos = skipField(line, pos)
		}
		// created_at
		pos = commaSkip(line, pos)
		var createdAt string
		if pos < len(line) && line[pos] == '\'' {
			createdAt, p = readString(line, pos+1)
			pos = p
		}

		records = append(records, postRecord{
			id:         id,
			categoryID: catID,
			title:      title,
			createdAt:  createdAt,
		})

		// Advance to end of tuple ')'
		for pos < len(line) && line[pos] != ')' {
			if line[pos] == '\'' {
				_, pos = readString(line, pos+1)
			} else {
				pos++
			}
		}
		if pos < len(line) {
			pos++ // ')'
		}
		if pos < len(line) && line[pos] == ',' {
			pos++ // inter-tuple comma
		}
	}
	return records
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

func writeCSV(path string, records []postRecord) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	_ = w.Write([]string{"id", "category_id", "title_ja", "slug", "created_at", "reviewed"})
	for _, r := range records {
		slug := r.slug
		if slug == "" {
			slug = generateSlug(r.id, r.title)
		}
		_ = w.Write([]string{r.id, r.categoryID, r.title, slug, r.createdAt, "false"})
	}
	return nil
}

func generateSlug(id, title string) string {
	var words []string
	var cur strings.Builder
	flush := func() {
		if cur.Len() > 1 {
			words = append(words, cur.String())
		}
		cur.Reset()
	}
	for _, r := range title {
		switch {
		case r == '+':
			flush()
			words = append(words, "plus")
		case r <= 127 && (unicode.IsLetter(r) || unicode.IsDigit(r)):
			cur.WriteRune(unicode.ToLower(r))
		default:
			flush()
		}
	}
	flush()

	// deduplicate adjacent
	var deduped []string
	for _, w := range words {
		if len(deduped) == 0 || deduped[len(deduped)-1] != w {
			deduped = append(deduped, w)
		}
	}
	if len(deduped) > 6 {
		deduped = deduped[:6]
	}
	if len(deduped) == 0 {
		return "post-" + id
	}
	return strings.Join(deduped, "-")
}

// assignAISlugs calls OpenAI GPT-4o-mini in batches to generate meaningful
// English slugs and writes them into each record's slug field.
func assignAISlugs(records []postRecord, apiKey string) error {
	const batchSize = 20
	for i := 0; i < len(records); i += batchSize {
		end := i + batchSize
		if end > len(records) {
			end = len(records)
		}
		batch := records[i:end]

		titles := make([]string, len(batch))
		for j, r := range batch {
			titles[j] = r.title
		}

		slugs, err := callOpenAI(titles, apiKey)
		if err != nil {
			return fmt.Errorf("batch %d-%d: %w", i, end, err)
		}
		for j := range batch {
			if j < len(slugs) && slugs[j] != "" {
				records[i+j].slug = slugs[j]
			}
		}
		fmt.Printf("  processed %d/%d\n", end, len(records))
	}
	return nil
}

var slugClean = regexp.MustCompile(`[^a-z0-9-]+`)

func sanitizeSlug(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	// strip list prefix like "1. " that the model might emit
	if idx := strings.Index(s, ". "); idx >= 0 && idx < 4 {
		s = s[idx+2:]
	}
	s = slugClean.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	// collapse consecutive hyphens
	for strings.Contains(s, "--") {
		s = strings.ReplaceAll(s, "--", "-")
	}
	return s
}

type openAIRequest struct {
	Model       string          `json:"model"`
	Messages    []openAIMessage `json:"messages"`
	Temperature float32         `json:"temperature"`
}

type openAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openAIResponse struct {
	Choices []struct {
		Message openAIMessage `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error"`
}

const systemPrompt = `You are a URL slug generator for a Japanese tech blog.
Given a list of blog post titles (some in Japanese, some mixed), generate a concise, meaningful English slug for each.

Rules:
- Return exactly one slug per line, in the same order as the input.
- Each slug: 2–5 lowercase words, hyphen-separated.
- Use English words that clearly describe the article content.
- Treat "+" as a separator between tech names — do NOT use "plus".
- Drop Japanese grammatical particles (の, を, が, は, で, に, と, から, まで, etc.).
- Keep well-known tech names as-is (laravel, react, es6, ansible, docker, golang, etc.).
- Return ONLY the slugs, one per line — no numbers, no explanations.`

func callOpenAI(titles []string, apiKey string) ([]string, error) {
	var sb strings.Builder
	for _, t := range titles {
		sb.WriteString(t)
		sb.WriteByte('\n')
	}

	reqBody := openAIRequest{
		Model: "gpt-4o-mini",
		Messages: []openAIMessage{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: sb.String()},
		},
		Temperature: 0.2,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/chat/completions", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ai openAIResponse
	if err := json.Unmarshal(raw, &ai); err != nil {
		return nil, fmt.Errorf("unmarshal: %w — body: %s", err, string(raw))
	}
	if ai.Error != nil {
		return nil, fmt.Errorf("openai api error: %s", ai.Error.Message)
	}
	if len(ai.Choices) == 0 {
		return nil, fmt.Errorf("empty choices in response")
	}

	lines := strings.Split(strings.TrimSpace(ai.Choices[0].Message.Content), "\n")
	slugs := make([]string, 0, len(lines))
	for _, l := range lines {
		if s := sanitizeSlug(l); s != "" {
			slugs = append(slugs, s)
		}
	}
	return slugs, nil
}
