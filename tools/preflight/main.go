// preflight — Cloudflare Pages 移行前チェックツール
//
// Usage:
//
//	go run ./tools/preflight          # リポジトリルートから実行
//	go run ./tools/preflight -root /path/to/bmf-tech  # 任意のパスを指定
//	go run ./tools/preflight -skip-build              # gohan build を省略
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

// ────────────────────────────────────────────────────────────
// データ構造
// ────────────────────────────────────────────────────────────

type PostFrontMatter struct {
	Title          string   `yaml:"title"`
	Slug           string   `yaml:"slug"`
	Date           string   `yaml:"date"`
	Author         string   `yaml:"author"`
	TranslationKey string   `yaml:"translation_key"`
	Categories     []string `yaml:"categories"`
	Tags           []string `yaml:"tags"`
	Draft          bool     `yaml:"draft"`
}

type TaxonomyItem struct {
	Name string `yaml:"name"`
}

// ────────────────────────────────────────────────────────────
// 結果記録
// ────────────────────────────────────────────────────────────

type Level int

const (
	INFO Level = iota
	WARN
	FAIL
)

type Finding struct {
	Level   Level
	Section string
	Message string
}

var findings []Finding

func add(l Level, section, msg string) {
	findings = append(findings, Finding{l, section, msg})
}
func info(s, m string)  { add(INFO, s, m) }
func warn(s, m string)  { add(WARN, s, m) }
func fail(s, m string)  { add(FAIL, s, m) }
func passf(s, m string) { add(INFO, s, "✅ "+m) }

// ────────────────────────────────────────────────────────────
// ヘルパー
// ────────────────────────────────────────────────────────────

func loadTaxonomy(path string) (map[string]bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var items []TaxonomyItem
	if err := yaml.Unmarshal(data, &items); err != nil {
		return nil, err
	}
	m := make(map[string]bool, len(items))
	for _, it := range items {
		m[it.Name] = true
	}
	return m, nil
}

// parseFrontMatter は --- で囲まれた YAML フロントマターを返す
func parseFrontMatter(path string) (*PostFrontMatter, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	s := string(data)
	if !strings.HasPrefix(s, "---") {
		return nil, fmt.Errorf("no frontmatter")
	}
	end := strings.Index(s[3:], "\n---")
	if end < 0 {
		return nil, fmt.Errorf("frontmatter not closed")
	}
	fmYAML := s[3 : end+3]
	var fm PostFrontMatter
	if err := yaml.Unmarshal([]byte(fmYAML), &fm); err != nil {
		return nil, err
	}
	return &fm, nil
}

func slugFromPath(path string) string {
	base := filepath.Base(path)
	return strings.TrimSuffix(base, ".md")
}

// ────────────────────────────────────────────────────────────
// [CONTENT] JA / EN 記事チェック
// ────────────────────────────────────────────────────────────

func checkContent(root string) (jaSlugs, enSlugs map[string]bool) {
	sec := "CONTENT"
	jaDir := filepath.Join(root, "content", "ja", "posts")
	enDir := filepath.Join(root, "content", "en", "posts")

	jaSlugs = make(map[string]bool)
	enSlugs = make(map[string]bool)
	jaTranslationKey := make(map[string]string) // slug → translation_key
	enTranslationKey := make(map[string]string)

	// JA
	jaPaths, _ := filepath.Glob(filepath.Join(jaDir, "*.md"))
	jaMissingFields := 0
	jaSlugMismatch := 0
	jaFMError := 0
	for _, p := range jaPaths {
		fileSlug := slugFromPath(p)
		fm, err := parseFrontMatter(p)
		if err != nil {
			fail(sec, fmt.Sprintf("JA frontmatter parse error: %s — %v", filepath.Base(p), err))
			jaFMError++
			continue
		}
		// 必須フィールドチェック
		missing := []string{}
		if fm.Title == "" {
			missing = append(missing, "title")
		}
		if fm.Slug == "" {
			missing = append(missing, "slug")
		}
		if fm.Date == "" {
			missing = append(missing, "date")
		}
		if fm.TranslationKey == "" {
			missing = append(missing, "translation_key")
		}
		if len(fm.Categories) == 0 {
			missing = append(missing, "categories")
		}
		if len(missing) > 0 {
			fail(sec, fmt.Sprintf("JA missing fields [%s]: %s", strings.Join(missing, ","), filepath.Base(p)))
			jaMissingFields++
		}
		// slug と ファイル名が一致するか
		if fm.Slug != "" && fm.Slug != fileSlug {
			warn(sec, fmt.Sprintf("JA slug mismatch (file=%s, slug=%s): %s", fileSlug, fm.Slug, filepath.Base(p)))
			jaSlugMismatch++
		}
		jaSlugs[fileSlug] = true
		if fm.TranslationKey != "" {
			jaTranslationKey[fileSlug] = fm.TranslationKey
		}
	}
	info(sec, fmt.Sprintf("JA posts: %d files", len(jaPaths)))
	if jaFMError == 0 && jaMissingFields == 0 && jaSlugMismatch == 0 {
		passf(sec, fmt.Sprintf("All %d JA posts have valid frontmatter", len(jaPaths)))
	} else {
		if jaFMError > 0 {
			fail(sec, fmt.Sprintf("JA frontmatter parse errors: %d", jaFMError))
		}
		if jaMissingFields > 0 {
			fail(sec, fmt.Sprintf("JA posts with missing required fields: %d", jaMissingFields))
		}
		if jaSlugMismatch > 0 {
			warn(sec, fmt.Sprintf("JA posts with slug/filename mismatch: %d", jaSlugMismatch))
		}
	}

	// EN
	enPaths, _ := filepath.Glob(filepath.Join(enDir, "*.md"))
	enMissingFields := 0
	enFMError := 0
	enSlugMismatch := 0
	for _, p := range enPaths {
		fileSlug := slugFromPath(p)
		fm, err := parseFrontMatter(p)
		if err != nil {
			fail(sec, fmt.Sprintf("EN frontmatter parse error: %s — %v", filepath.Base(p), err))
			enFMError++
			continue
		}
		missing := []string{}
		if fm.Title == "" {
			missing = append(missing, "title")
		}
		if fm.Slug == "" {
			missing = append(missing, "slug")
		}
		if fm.Date == "" {
			missing = append(missing, "date")
		}
		if fm.TranslationKey == "" {
			missing = append(missing, "translation_key")
		}
		if len(fm.Categories) == 0 {
			missing = append(missing, "categories")
		}
		if len(missing) > 0 {
			fail(sec, fmt.Sprintf("EN missing fields [%s]: %s", strings.Join(missing, ","), filepath.Base(p)))
			enMissingFields++
		}
		if fm.Slug != "" && fm.Slug != fileSlug {
			warn(sec, fmt.Sprintf("EN slug mismatch (file=%s, slug=%s): %s", fileSlug, fm.Slug, filepath.Base(p)))
			enSlugMismatch++
		}
		enSlugs[fileSlug] = true
		if fm.TranslationKey != "" {
			enTranslationKey[fileSlug] = fm.TranslationKey
		}
	}
	info(sec, fmt.Sprintf("EN posts: %d files", len(enPaths)))
	if enFMError == 0 && enMissingFields == 0 && enSlugMismatch == 0 {
		passf(sec, fmt.Sprintf("All %d EN posts have valid frontmatter", len(enPaths)))
	} else {
		if enFMError > 0 {
			fail(sec, fmt.Sprintf("EN frontmatter parse errors: %d", enFMError))
		}
		if enMissingFields > 0 {
			fail(sec, fmt.Sprintf("EN posts with missing required fields: %d", enMissingFields))
		}
		if enSlugMismatch > 0 {
			warn(sec, fmt.Sprintf("EN posts with slug/filename mismatch: %d", enSlugMismatch))
		}
	}

	// 翻訳ペアチェック: JA に対応する EN があるか
	missing := []string{}
	for slug := range jaSlugs {
		if !enSlugs[slug] {
			missing = append(missing, slug)
		}
	}
	sort.Strings(missing)
	info(sec, fmt.Sprintf("JA posts without EN translation: %d / %d", len(missing), len(jaSlugs)))
	if len(missing) > 0 {
		warn(sec, fmt.Sprintf("JA posts without EN counterpart (%d articles):", len(missing)))
		for _, s := range missing {
			warn(sec, fmt.Sprintf("  - %s", s))
		}
	} else {
		passf(sec, "All JA posts have EN translations")
	}

	// translation_key 整合性チェック
	tkMismatch := 0
	for slug, tk := range enTranslationKey {
		jaTK, ok := jaTranslationKey[slug]
		if !ok {
			continue // JA がない場合はスキップ
		}
		if jaTK != tk {
			warn(sec, fmt.Sprintf("translation_key mismatch for %s (JA=%s, EN=%s)", slug, jaTK, tk))
			tkMismatch++
		}
	}
	if tkMismatch == 0 && len(enPaths) > 0 {
		passf(sec, "All translation_key values are consistent between JA and EN")
	}

	return jaSlugs, enSlugs
}

// ────────────────────────────────────────────────────────────
// [TAXONOMY] タグ・カテゴリー整合性チェック
// ────────────────────────────────────────────────────────────

func checkTaxonomy(root string) {
	sec := "TAXONOMY"
	tagsPath := filepath.Join(root, "taxonomies", "tags.yaml")
	catsPath := filepath.Join(root, "taxonomies", "categories.yaml")

	knownTags, err := loadTaxonomy(tagsPath)
	if err != nil {
		fail(sec, fmt.Sprintf("Cannot load tags.yaml: %v", err))
		return
	}
	knownCats, err := loadTaxonomy(catsPath)
	if err != nil {
		fail(sec, fmt.Sprintf("Cannot load categories.yaml: %v", err))
		return
	}
	info(sec, fmt.Sprintf("tags.yaml: %d tags, categories.yaml: %d categories", len(knownTags), len(knownCats)))

	unknownTags := make(map[string][]string) // tag → []slug
	unknownCats := make(map[string][]string) // cat → []slug

	dirs := []string{
		filepath.Join(root, "content", "ja", "posts"),
		filepath.Join(root, "content", "en", "posts"),
	}
	for _, dir := range dirs {
		paths, _ := filepath.Glob(filepath.Join(dir, "*.md"))
		for _, p := range paths {
			fm, err := parseFrontMatter(p)
			if err != nil {
				continue
			}
			slug := slugFromPath(p)
			locale := "en"
			if strings.Contains(dir, "/ja/") {
				locale = "ja"
			}
			key := locale + "/" + slug
			for _, t := range fm.Tags {
				if !knownTags[t] {
					unknownTags[t] = append(unknownTags[t], key)
				}
			}
			for _, c := range fm.Categories {
				if !knownCats[c] {
					unknownCats[c] = append(unknownCats[c], key)
				}
			}
		}
	}

	if len(unknownTags) == 0 {
		passf(sec, "All post tags exist in tags.yaml")
	} else {
		tagNames := make([]string, 0, len(unknownTags))
		for t := range unknownTags {
			tagNames = append(tagNames, t)
		}
		sort.Strings(tagNames)
		fail(sec, fmt.Sprintf("Unknown tags (not in tags.yaml): %d", len(unknownTags)))
		for _, t := range tagNames {
			fail(sec, fmt.Sprintf("  tag %q used in: %s", t, strings.Join(unknownTags[t], ", ")))
		}
	}

	if len(unknownCats) == 0 {
		passf(sec, "All post categories exist in categories.yaml")
	} else {
		catNames := make([]string, 0, len(unknownCats))
		for c := range unknownCats {
			catNames = append(catNames, c)
		}
		sort.Strings(catNames)
		fail(sec, fmt.Sprintf("Unknown categories (not in categories.yaml): %d", len(unknownCats)))
		for _, c := range catNames {
			fail(sec, fmt.Sprintf("  category %q used in: %s", c, strings.Join(unknownCats[c], ", ")))
		}
	}
}

// ────────────────────────────────────────────────────────────
// [REDIRECTS] _redirects チェック
// ────────────────────────────────────────────────────────────

func checkRedirects(root string) {
	sec := "REDIRECTS"
	path := filepath.Join(root, "_redirects")

	f, err := os.Open(path)
	if err != nil {
		fail(sec, fmt.Sprintf("Cannot open _redirects: %v", err))
		return
	}
	defer f.Close()

	type Rule struct {
		Line   int
		Source string
		Dest   string
		Status string
		Raw    string
	}

	var rules []Rule
	sources := make(map[string]int) // source → first line
	spaceErrors := 0
	statusErrors := 0
	duplicates := 0
	totalComments := 0
	lineNum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineNum++
		raw := scanner.Text()
		trimmed := strings.TrimSpace(raw)
		if trimmed == "" {
			continue
		}
		if strings.HasPrefix(trimmed, "#") {
			totalComments++
			continue
		}
		// フィールドはホワイトスペース区切り
		fields := strings.Fields(trimmed)
		if len(fields) < 2 {
			warn(sec, fmt.Sprintf("L%d: incomplete rule: %s", lineNum, raw))
			continue
		}
		src := fields[0]
		dest := fields[1]
		status := "301"
		if len(fields) >= 3 {
			status = fields[2]
		}

		// スペースが混入しているか:
		// Cloudflare Pages のフォーマットは「source  dest  [status]」の3フィールド。
		// strings.Fields でトークン数が4以上になる場合は dest か source に生スペースがある。
		// また、dest フィールドの中に全角スペース(U+3000)が混入しているケースも検出する。
		if len(fields) > 3 {
			fail(sec, fmt.Sprintf("L%d: unencoded space in source/destination (%d tokens, expected 2-3): %s", lineNum, len(fields), raw))
			spaceErrors++
		} else if len(fields) >= 2 {
			// 全角スペース(U+3000)チェック: fields[1] の中に含まれていないか
			if strings.ContainsRune(fields[1], '\u3000') {
				fail(sec, fmt.Sprintf("L%d: fullwidth space (U+3000) in destination: %s", lineNum, raw))
				spaceErrors++
			}
		}

		// ステータスコード
		if status != "301" && status != "302" && status != "200" {
			if _, err := strconv.Atoi(status); err != nil {
				fail(sec, fmt.Sprintf("L%d: invalid status code %q: %s", lineNum, status, raw))
				statusErrors++
			} else {
				warn(sec, fmt.Sprintf("L%d: unusual status code %s: %s", lineNum, status, raw))
			}
		}

		// 重複チェック
		if first, exist := sources[src]; exist {
			warn(sec, fmt.Sprintf("L%d: duplicate source (first at L%d): %s", lineNum, first, src))
			duplicates++
		} else {
			sources[src] = lineNum
		}

		rules = append(rules, Rule{lineNum, src, dest, status, raw})
	}

	info(sec, fmt.Sprintf("Total redirect rules: %d (comments: %d)", len(rules), totalComments))

	// Cloudflare Pages 制限: 2000ルール
	if len(rules) > 2000 {
		fail(sec, fmt.Sprintf("EXCEEDS Cloudflare Pages limit of 2000 rules! (current: %d)", len(rules)))
	} else {
		passf(sec, fmt.Sprintf("Rule count OK: %d / 2000", len(rules)))
	}

	if spaceErrors == 0 {
		passf(sec, "No unencoded spaces in redirect destinations")
	} else {
		fail(sec, fmt.Sprintf("Unencoded spaces in destinations: %d rules", spaceErrors))
	}

	if duplicates == 0 {
		passf(sec, "No duplicate source paths")
	} else {
		fail(sec, fmt.Sprintf("Duplicate source paths: %d", duplicates))
	}

	if statusErrors == 0 {
		passf(sec, "All redirect status codes are valid")
	}

	// 内部リンクの宛先チェック（/で始まる宛先が public/ に存在するか）
	// ただし、タグ・カテゴリーページはビルドで生成されるのでスキップ
	// 記事ページ (/ja/posts/ 宛先) で対応するJA / EN コンテンツが存在するか
	publicDir := filepath.Join(root, "public")
	if _, err := os.Stat(publicDir); err == nil {
		missingDest := 0
		checked := 0
		for _, r := range rules {
			if !strings.HasPrefix(r.Dest, "/ja/posts/") && !strings.HasPrefix(r.Dest, "/posts/") {
				continue
			}
			// /ja/posts/{slug}/ → public/ja/posts/{slug}/index.html
			destPath := filepath.Join(root, "public", strings.TrimPrefix(r.Dest, "/"), "index.html")
			destPathAlt := strings.TrimSuffix(destPath, "/index.html")
			if r.Dest == "/" {
				continue
			}
			checked++
			if _, err := os.Stat(destPath); err != nil {
				if _, err2 := os.Stat(destPathAlt); err2 != nil {
					fail(sec, fmt.Sprintf("Destination not found in public/: %s (from %s)", r.Dest, r.Source))
					missingDest++
				}
			}
		}
		if checked > 0 {
			if missingDest == 0 {
				passf(sec, fmt.Sprintf("All %d article redirect destinations exist in public/", checked))
			} else {
				fail(sec, fmt.Sprintf("Redirect destinations missing in public/: %d", missingDest))
			}
		}
	} else {
		warn(sec, "public/ directory not found — skipping destination existence check (run 'gohan build' first)")
	}
}

// ────────────────────────────────────────────────────────────
// [CONFIG] config.yaml チェック
// ────────────────────────────────────────────────────────────

func checkConfig(root string) {
	sec := "CONFIG"
	path := filepath.Join(root, "config.yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		fail(sec, fmt.Sprintf("Cannot read config.yaml: %v", err))
		return
	}
	content := string(data)

	// noindex チェック
	if strings.Contains(content, `noindex: "true"`) || strings.Contains(content, "noindex: true") {
		fail(sec, `noindex is set to "true" — remove this before deploying to Cloudflare Pages!`)
	} else {
		passf(sec, "noindex is not set (OK)")
	}

	// GA ID チェック
	if strings.Contains(content, `# ga_id:`) || strings.Contains(content, `#ga_id:`) {
		warn(sec, "ga_id is commented out — enable it before production deploy")
	} else if strings.Contains(content, "ga_id:") {
		passf(sec, "ga_id is configured")
	} else {
		warn(sec, "ga_id not found in config.yaml")
	}

	// AdSense ID チェック
	if strings.Contains(content, `# adsense_id:`) || strings.Contains(content, `#adsense_id:`) {
		warn(sec, "adsense_id is commented out — enable it before production deploy")
	} else if strings.Contains(content, "adsense_id:") {
		passf(sec, "adsense_id is configured")
	} else {
		warn(sec, "adsense_id not found in config.yaml")
	}

	// base_url チェック
	if strings.Contains(content, `base_url: "https://bmf-tech.com"`) {
		passf(sec, `base_url is set to https://bmf-tech.com`)
	} else if strings.Contains(content, "base_url:") {
		warn(sec, "base_url is set but not https://bmf-tech.com — verify before deploy")
	} else {
		fail(sec, "base_url not found in config.yaml")
	}

	// github_repo / github_branch チェック
	if strings.Contains(content, "github_repo:") {
		passf(sec, "github_repo is configured")
	} else {
		warn(sec, "github_repo not found in config.yaml")
	}
}

// ────────────────────────────────────────────────────────────
// [ASSETS] アセットチェック
// ────────────────────────────────────────────────────────────

func checkAssets(root string) {
	sec := "ASSETS"

	// 外部画像参照チェック
	externalRefs := 0
	externalDomains := make(map[string]int)
	dirs := []string{
		filepath.Join(root, "content", "ja", "posts"),
		filepath.Join(root, "content", "en", "posts"),
	}
	for _, dir := range dirs {
		paths, _ := filepath.Glob(filepath.Join(dir, "*.md"))
		for _, p := range paths {
			data, err := os.ReadFile(p)
			if err != nil {
				continue
			}
			lines := strings.Split(string(data), "\n")
			for _, line := range lines {
				// Markdown image pattern: ![...](http...)
				if strings.Contains(line, "](http") {
					idx := strings.Index(line, "](http")
					if idx >= 0 {
						url := line[idx+2:]
						if end := strings.IndexAny(url, ") "); end >= 0 {
							url = url[:end]
						}
						externalRefs++
						// ドメイン抽出
						if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {
							parts := strings.SplitN(strings.TrimPrefix(strings.TrimPrefix(url, "https://"), "http://"), "/", 2)
							if len(parts) > 0 {
								externalDomains[parts[0]]++
							}
						}
					}
				}
			}
		}
	}

	if externalRefs > 0 {
		warn(sec, fmt.Sprintf("External image references found: %d (risk of future URL expiry)", externalRefs))
		domains := make([]string, 0, len(externalDomains))
		for d := range externalDomains {
			domains = append(domains, d)
		}
		sort.Strings(domains)
		for _, d := range domains {
			warn(sec, fmt.Sprintf("  %s: %d image(s)", d, externalDomains[d]))
		}
	} else {
		passf(sec, "No external image references found")
	}
}

// ────────────────────────────────────────────────────────────
// [BUILD] gohan build チェック
// ────────────────────────────────────────────────────────────

func checkBuild(root string) {
	sec := "BUILD"
	gohanBin, err := exec.LookPath("gohan")
	if err != nil {
		warn(sec, "gohan binary not found in PATH — skipping build check")
		return
	}
	info(sec, fmt.Sprintf("Running: gohan build (binary: %s)", gohanBin))
	cmd := exec.Command(gohanBin, "build")
	cmd.Dir = root
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		fail(sec, fmt.Sprintf("gohan build FAILED: %v", err))
		// 最大30行のエラー出力を表示
		lines := strings.Split(out.String(), "\n")
		if len(lines) > 30 {
			lines = lines[:30]
		}
		for _, l := range lines {
			if l != "" {
				fail(sec, "  "+l)
			}
		}
	} else {
		passf(sec, "gohan build succeeded")
	}
}

// ────────────────────────────────────────────────────────────
// [LIVE] リダイレクト先 HTTP 200 チェック
// ────────────────────────────────────────────────────────────

func checkLive(root string, baseURL string) {
	sec := "LIVE"
	path := filepath.Join(root, "_redirects")

	f, err := os.Open(path)
	if err != nil {
		fail(sec, fmt.Sprintf("Cannot open _redirects: %v", err))
		return
	}
	defer f.Close()

	// 宛先パスを収集（重複除去）
	destSet := make(map[string]bool)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		dest := fields[1]
		if strings.HasPrefix(dest, "/") {
			destSet[dest] = true
		}
	}

	destList := make([]string, 0, len(destSet))
	for d := range destSet {
		destList = append(destList, d)
	}
	sort.Strings(destList)

	info(sec, fmt.Sprintf("Checking %d unique redirect destinations against %s ...", len(destList), baseURL))

	type result struct {
		dest   string
		status int
		err    error
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
		// リダイレクトをたどらず、宛先そのものの応答だけを確認する
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	sem := make(chan struct{}, 20) // 同時接続数上限
	resultsCh := make(chan result, len(destList))
	var wg sync.WaitGroup

	for _, d := range destList {
		wg.Add(1)
		go func(dest string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			url := baseURL + dest
			resp, err := client.Get(url)
			if err != nil {
				resultsCh <- result{dest, 0, err}
				return
			}
			resp.Body.Close()
			resultsCh <- result{dest, resp.StatusCode, nil}
		}(d)
	}

	wg.Wait()
	close(resultsCh)

	notOK := 0
	for r := range resultsCh {
		if r.err != nil {
			fail(sec, fmt.Sprintf("GET %s → error: %v", r.dest, r.err))
			notOK++
		} else if r.status != 200 {
			fail(sec, fmt.Sprintf("GET %s → HTTP %d (expected 200)", r.dest, r.status))
			notOK++
		}
	}

	if notOK == 0 {
		passf(sec, fmt.Sprintf("All %d redirect destinations return HTTP 200", len(destList)))
	} else {
		fail(sec, fmt.Sprintf("%d / %d destinations did NOT return HTTP 200", notOK, len(destList)))
	}
}

// ────────────────────────────────────────────────────────────
// 出力
// ────────────────────────────────────────────────────────────

func printReport() {
	const (
		colorReset  = "\033[0m"
		colorRed    = "\033[31m"
		colorYellow = "\033[33m"
		colorGreen  = "\033[32m"
		colorCyan   = "\033[36m"
		colorBold   = "\033[1m"
	)

	fmt.Println()
	fmt.Println(colorBold + "═══════════════════════════════════════════════════════" + colorReset)
	fmt.Println(colorBold + "  bmf-tech Cloudflare Pages 移行前チェックレポート" + colorReset)
	fmt.Println(colorBold + "═══════════════════════════════════════════════════════" + colorReset)

	currentSection := ""
	failCount := 0
	warnCount := 0
	passCount := 0

	for _, f := range findings {
		if f.Section != currentSection {
			currentSection = f.Section
			fmt.Printf("\n%s[%s]%s\n", colorCyan+colorBold, f.Section, colorReset)
		}
		switch f.Level {
		case FAIL:
			fmt.Printf("  %s❌ %s%s\n", colorRed, f.Message, colorReset)
			failCount++
		case WARN:
			fmt.Printf("  %s⚠️  %s%s\n", colorYellow, f.Message, colorReset)
			warnCount++
		case INFO:
			if strings.HasPrefix(f.Message, "✅") {
				fmt.Printf("  %s%s%s\n", colorGreen, f.Message, colorReset)
				passCount++
			} else {
				fmt.Printf("  ℹ️  %s\n", f.Message)
			}
		}
	}

	fmt.Println()
	fmt.Println(colorBold + "───────────────────────────────────────────────────────" + colorReset)
	fmt.Printf("  %s✅ PASS: %d%s  %s⚠️  WARN: %d%s  %s❌ FAIL: %d%s\n",
		colorGreen, passCount, colorReset,
		colorYellow, warnCount, colorReset,
		colorRed, failCount, colorReset,
	)
	fmt.Println(colorBold + "───────────────────────────────────────────────────────" + colorReset)

	if failCount > 0 {
		fmt.Printf("\n%s%s🚫 移行不可: %d 件の重大な問題を修正してください。%s\n\n", colorRed, colorBold, failCount, colorReset)
		os.Exit(1)
	} else if warnCount > 0 {
		fmt.Printf("\n%s%s⚠  移行前に警告事項を確認してください。%s\n\n", colorYellow, colorBold, colorReset)
	} else {
		fmt.Printf("\n%s%s🚀 全チェック通過！Cloudflare Pages への移行準備が整っています。%s\n\n", colorGreen, colorBold, colorReset)
	}
}

// ────────────────────────────────────────────────────────────
// main
// ────────────────────────────────────────────────────────────

func main() {
	var (
		root      = flag.String("root", ".", "bmf-tech リポジトリのルートパス")
		skipBuild = flag.Bool("skip-build", false, "gohan build チェックをスキップ")
		live      = flag.Bool("live", false, "localhost への HTTP 200 チェックを実行")
		port      = flag.Int("port", 8080, "gohan serve のポート番号 (-live 時に使用)")
	)
	flag.Parse()

	absRoot, err := filepath.Abs(*root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid root path: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("チェック対象: %s\n", absRoot)

	checkContent(absRoot)
	checkTaxonomy(absRoot)
	checkRedirects(absRoot)
	checkConfig(absRoot)
	checkAssets(absRoot)
	if !*skipBuild {
		checkBuild(absRoot)
	} else {
		warn("BUILD", "gohan build チェックはスキップされました (-skip-build フラグ)")
	}

	if *live {
		checkLive(absRoot, fmt.Sprintf("http://localhost:%d", *port))
	} else {
		warn("LIVE", "HTTP 200 チェックはスキップ (実行するには -live フラグを追加、サーバー起動後に使用)")
	}

	printReport()
}
