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
	"strings"
	"time"
	"unicode"

	"gopkg.in/yaml.v3"
)

var (
	jaDir   = flag.String("ja-dir", "../../content/ja/posts", "JA article dir")
	enDir   = flag.String("en-dir", "../../content/en/posts", "EN article dir")
	limitN  = flag.Int("limit", 0, "max articles (0=all)")
	inputF  = flag.String("input", "", "translate specific file")
	dryRun  = flag.Bool("dry-run", false, "show without translating")
	delayMs = flag.Int("delay", 500, "ms between API calls")
)

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
	if err != nil { return }
	content := string(data)
	if !strings.HasPrefix(content, "---") { body = content; return }
	rest := content[3:]
	idx := strings.Index(rest, "\n---")
	if idx == -1 { body = content; return }
	rawFM := strings.TrimSpace(rest[:idx])
	body = strings.TrimPrefix(rest[idx+4:], "\n")
	err = yaml.Unmarshal([]byte(rawFM), &fm)
	return
}

func writeMarkdown(path string, fm FrontMatter, body string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil { return err }
	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(fm); err != nil { return err }
	fmStr := strings.TrimRight(buf.String(), "\n")
	content := "---\n" + fmStr + "\n---\n\n" + body
	return os.WriteFile(path, []byte(content), 0o644)
}

func addTranslationKey(path, key string) error {
	fm, body, err := parseMarkdown(path)
	if err != nil { return err }
	if fm.TranslationKey == key { return nil }
	fm.TranslationKey = key
	return writeMarkdown(path, fm, body)
}

type aiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type aiRequest struct {
	Model       string      `json:"model"`
	Messages    []aiMessage `json:"messages"`
	Temperature float64     `json:"temperature"`
}

type aiChoice struct { Message aiMessage `json:"message"` }

type aiResponse struct {
	Choices []aiChoice `json:"choices"`
	Error   *struct{ Message string `json:"message"` } `json:"error,omitempty"`
}

type TranslationResult struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	Tags        []string `json:"tags"`
	Body        string   `json:"body"`
}

type apiConfig struct { baseURL, token, model string }

func loadAPIConfig() (apiConfig, error) {
	cfg := apiConfig{baseURL: os.Getenv("AI_BASE_URL"), model: os.Getenv("AI_MODEL")}
	if gt := os.Getenv("GITHUB_TOKEN"); gt != "" {
		cfg.token = gt
		if cfg.baseURL == "" { cfg.baseURL = "https://models.inference.ai.azure.com" }
		if cfg.model == "" { cfg.model = "gpt-4o" }
		return cfg, nil
	}
	if ak := os.Getenv("OPENAI_API_KEY"); ak != "" {
		cfg.token = ak
		if cfg.baseURL == "" { cfg.baseURL = "https://api.openai.com/v1" }
		if cfg.model == "" { cfg.model = "gpt-4o" }
		return cfg, nil
	}
	return cfg, fmt.Errorf("set GITHUB_TOKEN or OPENAI_API_KEY")
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
		"- Keep English tech terms as-is\n" +
		"- Translate Japanese category/tag names to English equivalents\n\n" +
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
	if err != nil { return TranslationResult{}, err }
	endpoint := strings.TrimRight(cfg.baseURL, "/") + "/chat/completions"
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(reqJSON))
	if err != nil { return TranslationResult{}, err }
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cfg.token)
	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil { return TranslationResult{}, err }
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil { return TranslationResult{}, err }
	var aiResp aiResponse
	if err := json.Unmarshal(raw, &aiResp); err != nil {
		return TranslationResult{}, fmt.Errorf("parse response: %w: %s", err, raw)
	}
	if aiResp.Error != nil {
		return TranslationResult{}, fmt.Errorf("API error: %s", aiResp.Error.Message)
	}
	if len(aiResp.Choices) == 0 {
		return TranslationResult{}, fmt.Errorf("empty response: %s", raw)
	}
	c := strings.TrimSpace(aiResp.Choices[0].Message.Content)
	if strings.HasPrefix(c, "```") {
		sc := bufio.NewScanner(strings.NewReader(c))
		var out []string
		first := true
		for sc.Scan() {
			ln := sc.Text()
			if strings.HasPrefix(ln, "```") {
				if first { first = false; continue }
				break
			}
			out = append(out, ln)
		}
		c = strings.Join(out, "\n")
	}
	var res TranslationResult
	if err := json.Unmarshal([]byte(c), &res); err != nil {
		return TranslationResult{}, fmt.Errorf("parse result: %w: %s", err, c)
	}
	return res, nil
}

func hasJapanese(s string) bool {
	for _, r := range s {
		if unicode.In(r, unicode.Hiragana, unicode.Katakana, unicode.Han) { return true }
	}
	return false
}

func main() {
	flag.Parse()
	cfg, cfgErr := loadAPIConfig()
	if cfgErr != nil && !*dryRun { fmt.Fprintln(os.Stderr, cfgErr); os.Exit(1) }

	var files []string
	if *inputF != "" {
		files = []string{*inputF}
	} else {
		entries, err := os.ReadDir(*jaDir)
		if err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(1) }
		for _, e := range entries {
			if !e.IsDir() && strings.HasSuffix(e.Name(), ".md") {
				files = append(files, filepath.Join(*jaDir, e.Name()))
			}
		}
	}

	if err := os.MkdirAll(*enDir, 0o755); err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(1) }

	skipped, translated, errors := 0, 0, 0
	for _, jaPath := range files {
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
		if slug == "" { slug = strings.TrimSuffix(filepath.Base(jaPath), ".md") }
		enPath := filepath.Join(*enDir, slug+".md")
		if _, err := os.Stat(enPath); err == nil {
			fmt.Printf("[SKIP] %s\n", slug)
			skipped++
			continue
		}
		if *dryRun {
			mark := ""
			if hasJapanese(fm.Title) { mark = " [JA]" }
			fmt.Printf("[DRY-RUN] %s%s\n", slug, mark)
			translated++
			continue
		}
		fmt.Printf("[TRANSLATING] %s ...\n", slug)
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
}
