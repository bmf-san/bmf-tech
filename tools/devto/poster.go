package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const devToAPIBase = "https://dev.to/api"

// Poster calls the dev.to API to create articles.
type Poster struct {
	APIKey string
	DryRun bool
	client http.Client
}

type devToRequest struct {
	Article devToArticlePayload `json:"article"`
}

type devToArticlePayload struct {
	Title        string   `json:"title"`
	BodyMarkdown string   `json:"body_markdown"`
	Published    bool     `json:"published"`
	Tags         []string `json:"tags"`
	CanonicalURL string   `json:"canonical_url"`
	Description  string   `json:"description,omitempty"`
	MainImage    string   `json:"main_image,omitempty"`
}

type devToResponse struct {
	ID int `json:"id"`
}

// Post creates a new article on dev.to and returns its numeric ID.
// It retries once on HTTP 429 (Too Many Requests) after honouring the
// Retry-After header (defaulting to 30 seconds).
func (p *Poster) Post(a *Article) (int, error) {
	payload := devToRequest{
		Article: devToArticlePayload{
			Title:        a.Title,
			BodyMarkdown: a.Body,
			Published:    true,
			Tags:         a.Tags,
			CanonicalURL: a.CanonicalURL,
			Description:  a.Description,
			MainImage:    a.MainImage,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return 0, fmt.Errorf("marshal payload: %w", err)
	}

	const maxAttempts = 3
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		id, retry, err := p.doPost(body)
		if err == nil {
			return id, nil
		}
		if !retry || attempt == maxAttempts {
			return 0, err
		}
		// rate limited – wait and retry
		fmt.Printf("  rate limited, waiting 30s before retry (attempt %d/%d)…\n", attempt, maxAttempts)
		time.Sleep(30 * time.Second)
	}
	return 0, fmt.Errorf("all retries exhausted")
}

// doPost performs a single POST. Returns (id, shouldRetry, error).
func (p *Poster) doPost(body []byte) (int, bool, error) {
	req, err := http.NewRequest(http.MethodPost, devToAPIBase+"/articles", bytes.NewReader(body))
	if err != nil {
		return 0, false, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", p.APIKey)

	resp, err := p.client.Do(req)
	if err != nil {
		return 0, false, fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, false, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		return 0, true, fmt.Errorf("rate limited (429)")
	}

	if resp.StatusCode != http.StatusCreated {
		return 0, false, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, truncate(string(respBody), 300))
	}

	var result devToResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return 0, false, fmt.Errorf("decode response: %w", err)
	}
	return result.ID, false, nil
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}
