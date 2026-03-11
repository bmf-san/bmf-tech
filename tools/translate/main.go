package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
	"unicode"

	"gopkg.in/yaml.v3"
)

var (
	jaDir   = flag.String("ja-dir", "../../content/ja/posts", "JA article dir")
	enDir   = flag.String("en-dir", "../../content/en/posts", "EN article dir")
	taxDir  = flag.String("tax-dir", "../../taxonomies", "taxonomy dir (tags.yaml / categories.yaml)")
	limitN  = flag.Int("limit", 0, "max articles to translate (0=all)")
	inputF  = flag.String("input", "", "translate specific file only")
	dryRun  = flag.Bool("dry-run", false, "show what would be translated without doing it")
	delayMs = flag.Int("delay", 1000, "ms between API calls")
)

// ---- Frontmatter ----

type FrontMatter struct {
	Title          string      `yaml:"title"`
	Slug           string      `yaml:"slug"`
	Date           interface{} `yaml:"date"`
	Author         string      `yaml:"author,omitempty"`
	Categories     []string    `yaml:"categories,omitempty"`
	Tags           []string    `yaml:"tags,omitempty"`
	Description    string      `yaml:"description,omitempty"`
	Draft          bool        `yaml:"draft,omitempty"`
	Template       string      `yaml:"template,omitempty"`
	TranslationKey string      `yaml:"translation_key,omitempty"`
}

func parseMarkdown(path string) (fm FrontMatter, body string, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	content := string(data)
	if !strings.HasPrefix(content, "---") {
		body = content
		return
	}
	rest := content[3:]
	idx := strings.Index(rest, "\n---")
	if idx == -1 {
		body = content
		return
	}
	rawFM := strings.TrimSpace(rest[:idx])
	body = strings.TrimPrefix(rest[idx+4:], "\n")
	err = yaml.Unmarshal([]byte(rawFM), &fm)
	return
}

func writeMarkdown(path string, fm FrontMatter, body string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(fm); err != nil {
		return err
	}
	fmStr := strings.TrimRight(buf.String(), "\n")
	content := "---\n" + fmStr + "\n---\n\n" + body
	return os.WriteFile(path, []byte(content), 0o644)
}

func addTranslationKey(path, key string) error {
	fm, body, err := parseMarkdown(path)
	if err != nil {
		return err
	}
	if fm.TranslationKey == key {
		return nil
	}
	fm.TranslationKey = key
	return writeMarkdown(path, fm, body)
}

// ---- Taxonomy ----

type Taxonomy struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description,omitempty"`
}

// loadTaxonomyNames reads a YAML taxonomy file and returns a set of names.
func loadTaxonomyNames(path string) (map[string]bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]bool{}, nil
		}
		return nil, err
	}
	var entries []Taxonomy
	if err := yaml.Unmarshal(data, &entries); err != nil {
		return nil, err
	}
	set := make(map[string]bool, len(entries))
	for _, e := range entries {
		set[e.Name] = true
	}
	return set, nil
}

// appendTaxonomyNames appends new names to a taxonomy YAML file, keeping existing entries.
func appendTaxonomyNames(path string, newNames []string) error {
	data, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	var entries []Taxonomy
	if len(data) > 0 {
		if err := yaml.Unmarshal(data, &entries); err != nil {
			return err
		}
	}
	existing := make(map[string]bool, len(entries))
	for _, e := range entries {
		existing[e.Name] = true
	}
	for _, n := range newNames {
		if !existing[n] {
			entries = append(entries, Taxonomy{Name: n})
		}
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name < entries[j].Name })
	out, err := yaml.Marshal(entries)
	if err != nil {
		return err
	}
	return os.WriteFile(path, out, 0o644)
}

// updateTaxonomies scans all EN articles and adds missing tags/categories to YAML files.
func updateTaxonomies(enDir, taxDir string) error {
	tagsPath := filepath.Join(taxDir, "tags.yaml")
	catsPath := filepath.Join(taxDir, "categories.yaml")

	existingTags, err := loadTaxonomyNames(tagsPath)
	if err != nil {
		return fmt.Errorf("tags.yaml: %w", err)
	}
	existingCats, err := loadTaxonomyNames(catsPath)
	if err != nil {
		return fmt.Errorf("categories.yaml: %w", err)
	}

	newTags := map[string]bool{}
	newCats := map[string]bool{}

	entries, err := os.ReadDir(enDir)
	if err != nil {
		return err
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		fm, _, err := parseMarkdown(filepath.Join(enDir, e.Name()))
		if err != nil {
			continue
		}
		for _, t := range fm.Tags {
			if !existingTags[t] {
				newTags[t] = true
			}
		}
		for _, c := range fm.Categories {
			if !existingCats[c] {
				newCats[c] = true
			}
		}
	}

	if len(newTags) > 0 {
		newTagList := make([]string, 0, len(newTags))
		for t := range newTags {
			newTagList = append(newTagList, t)
		}
		sort.Strings(newTagList)
		if err := appendTaxonomyNames(tagsPath, newTagList); err != nil {
			return fmt.Errorf("update tags: %w", err)
		}
		fmt.Printf("[TAXONOMY] tags.yaml: added %d new tags: %v\n", len(newTagList), newTagList)
	}
	if len(newCats) > 0 {
		newCatList := make([]string, 0, len(newCats))
		for c := range newCats {
			newCatList = append(newCatList, c)
		}
		sort.Strings(newCatList)
		if err := appendTaxonomyNames(catsPath, newCatList); err != nil {
			return fmt.Errorf("update categories: %w", err)
		}
		fmt.Printf("[TAXONOMY] categories.yaml: added %d new categories: %v\n", len(newCatList), newCatList)
	}
	if len(newTags) == 0 && len(newCats) == 0 {
		fmt.Println("[TAXONOMY] All tags and categories already registered.")
	}
	return nil
}

// ---- AI Translation ----

type aiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type aiRequest struct {
	Model       string      `json:"model"`
	Messages    []aiMessage `json:"messages"`
	Temperature float64     `json:"temperature"`
}

type aiChoice struct {
	Message aiMessage `json:"message"`
}

type aiResponse struct {
	Choices []aiChoice `json:"choices"`
	Error   *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

type TranslationResult struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	Tags        []string `json:"tags"`
	Body        string   `json:"body"`
}

const (
	openAIBaseURL = "https://api.openai.com/v1"
	openAIModel   = "gpt-4o"
)

type apiConfig struct{ token, model string }

func loadAPIConfig() (apiConfig, error) {
	ak := os.Getenv("OPENAI_API_KEY")
	if ak == "" {
		return apiConfig{}, fmt.Errorf("OPENAI_API_KEY is not set")
	}
	return apiConfig{token: ak, model: openAIModel}, nil
}

func buildPrompt(fm FrontMatter, body string) string {
	catJSON, _ := json.Marshal(fm.Categories)
	tagJSON, _ := json.Marshal(fm.Tags)
	return "Translate the following Japanese blog post to English.\n\n" +
		"Return ONLY a valid JSON object with exactly these fields:\n" +
		"  title (string), description (string), categories ([]string), tags ([]string), body (string)\n\n" +
		"Rules:\n" +
		"- Preserve all markdown formatting\n" +
		"- Do NOT translate content inside code blocks or URLs\n" +
		"- Keep English tech terms, product names, and command names as-is\n" +
		"- Translate Japanese category/tag names to English equivalents\n" +
		"- Use concise English category/tag names (1-3 words)\n\n" +
		"FRONTMATTER:\ntitle: " + fm.Title + "\ndesc: " + fm.Description + "\n" +
		"categories: " + string(catJSON) + "\ntags: " + string(tagJSON) + "\n\nBODY:\n" + body
}

func callAPI(cfg apiConfig, fm FrontMatter, body string) (TranslationResult, error) {
	reqBody := aiRequest{
		Model: cfg.model,
		Messages: []aiMessage{
			{Role: "system", Content: "You are a professional technical translator. Translate Japanese developer blog posts to natural English. Always respond with valid JSON only."},
			{Role: "user", Content: buildPrompt(fm, body)},
		},
		Temperature: 0.2,
	}
	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return TranslationResult{}, err
	}
	endpoint := openAIBaseURL + "/chat/completions"
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(reqJSON))
	if err != nil {
		return TranslationResult{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cfg.token)
	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return TranslationResult{}, err
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return TranslationResult{}, err
	}
	var aiResp aiResponse
	if err := json.Unmarshal(raw, &aiResp); err != nil {
		return TranslationResult{}, fmt.Errorf("parse response: %w: %s", err, raw[:min(200, len(raw))])
	}
	if aiResp.Error != nil {
		return TranslationResult{}, fmt.Errorf("API error: %s", aiResp.Error.Message)
	}
	if len(aiResp.Choices) == 0 {
		return TranslationResult{}, fmt.Errorf("empty response: %s", raw[:min(200, len(raw))])
	}
	c := strings.TrimSpace(aiResp.Choices[0].Message.Content)
	if strings.HasPrefix(c, "```") {
		sc := bufio.NewScanner(strings.NewReader(c))
		var out []string
		first := true
		for sc.Scan() {
			ln := sc.Text()
			if strings.HasPrefix(ln, "```") {
				if first {
					first = false
					continue
				}
				break
			}
			out = append(out, ln)
		}
		c = strings.Join(out, "\n")
	}
	c = sanitizeJSONContent(c)
	var res TranslationResult
	if err := json.Unmarshal([]byte(c), &res); err != nil {
		return TranslationResult{}, fmt.Errorf("parse result: %w: %s", err, c)
	}
	return res, nil
}

// sanitizeJSONContent fixes common issues in AI-generated JSON:
// 1. Literal tab characters (0x09) inside JSON strings → \t escape
// 2. Invalid backslash escapes (e.g., \T, \( from PHP/LaTeX) → \\T, \\(
func sanitizeJSONContent(s string) string {
	// Replace literal tab bytes
	s = strings.ReplaceAll(s, "\t", `\t`)
	// Scan byte by byte and fix invalid JSON escape sequences
	var out strings.Builder
	out.Grow(len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && i+1 < len(s) {
			switch s[i+1] {
			case '"', '\\', '/', 'b', 'f', 'n', 'r', 't', 'u':
				// valid JSON escape — copy both chars
				out.WriteByte(s[i])
				out.WriteByte(s[i+1])
				i++
			default:
				// invalid escape — double the backslash
				out.WriteByte('\\')
				out.WriteByte('\\')
				out.WriteByte(s[i+1])
				i++
			}
		} else {
			out.WriteByte(s[i])
		}
	}
	return out.String()
}

func hasJapanese(s string) bool {
	for _, r := range s {
		if unicode.In(r, unicode.Hiragana, unicode.Katakana, unicode.Han) {
			return true
		}
	}
	return false
}

// ---- Main ----

func main() {
	flag.Parse()
	cfg, cfgErr := loadAPIConfig()
	if cfgErr != nil && !*dryRun {
		fmt.Fprintln(os.Stderr, cfgErr)
		os.Exit(1)
	}

	var files []string
	if *inputF != "" {
		files = []string{*inputF}
	} else {
		entries, err := os.ReadDir(*jaDir)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, e := range entries {
			if !e.IsDir() && strings.HasSuffix(e.Name(), ".md") {
				files = append(files, filepath.Join(*jaDir, e.Name()))
			}
		}
	}

	if err := os.MkdirAll(*enDir, 0o755); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	total := len(files)
	skipped, translated, errors := 0, 0, 0
	for i, jaPath := range files {
		if *limitN > 0 && translated >= *limitN {
			fmt.Printf("--limit %d reached\n", *limitN)
			break
		}
		fm, body, err := parseMarkdown(jaPath)
		if err != nil {
			fmt.Printf("[ERROR] %s: %v\n", jaPath, err)
			errors++
			continue
		}
		slug := fm.Slug
		if slug == "" {
			slug = strings.TrimSuffix(filepath.Base(jaPath), ".md")
		}
		enPath := filepath.Join(*enDir, slug+".md")
		if _, err := os.Stat(enPath); err == nil {
			fmt.Printf("[SKIP] %s\n", slug)
			skipped++
			continue
		}
		if *dryRun {
			mark := ""
			if hasJapanese(fm.Title) {
				mark = " [JA]"
			}
			fmt.Printf("[DRY-RUN] %s%s\n", slug, mark)
			translated++
			continue
		}
		fmt.Printf("[%d/%d] %s ...\n", i+1, total, slug)
		result, err := callAPI(cfg, fm, body)
		if err != nil {
			fmt.Printf("[ERROR] %s: %v\n", slug, err)
			errors++
			time.Sleep(time.Duration(*delayMs) * time.Millisecond)
			continue
		}
		enFM := FrontMatter{
			Title:          result.Title,
			Slug:           slug,
			Date:           fm.Date,
			Author:         fm.Author,
			Categories:     result.Categories,
			Tags:           result.Tags,
			Description:    result.Description,
			Draft:          fm.Draft,
			Template:       fm.Template,
			TranslationKey: slug,
		}
		if err := writeMarkdown(enPath, enFM, result.Body); err != nil {
			fmt.Printf("[ERROR] write %s: %v\n", enPath, err)
			errors++
			continue
		}
		if err := addTranslationKey(jaPath, slug); err != nil {
			fmt.Printf("[WARN] translation_key %s: %v\n", slug, err)
		}
		fmt.Printf("[OK] %s\n", slug)
		translated++
		time.Sleep(time.Duration(*delayMs) * time.Millisecond)
	}

	fmt.Printf("\ndone: translated=%d skipped=%d errors=%d\n", translated, skipped, errors)

	// Update taxonomies from all EN articles
	if !*dryRun {
		fmt.Println("\n[TAXONOMY] Scanning EN articles for new tags/categories...")
		if err := updateTaxonomies(*enDir, *taxDir); err != nil {
			fmt.Fprintf(os.Stderr, "[WARN] taxonomy update: %v\n", err)
		}
	}
}
