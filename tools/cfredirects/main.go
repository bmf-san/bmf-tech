// cfredirects syncs bulk-redirects.txt to Cloudflare Bulk Redirects via API.
//
// Usage:
//
//	cfredirects --file=bulk-redirects.txt --list=bmf-tech-redirects --domain=https://bmf-tech.com
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const cfAPI = "https://api.cloudflare.com/client/v4"

type redirect struct {
	Source string
	Target string
	Code   int
}

type cfClient struct {
	accountID string
	token     string
}

func main() {
	file := flag.String("file", "bulk-redirects.txt", "path to redirect rules file")
	listName := flag.String("list", "bmf-tech-redirects", "Cloudflare Redirect List name")
	domain := flag.String("domain", "https://bmf-tech.com", "base domain for relative source/target URLs")
	flag.Parse()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if accountID == "" || apiToken == "" {
		log.Fatal("CLOUDFLARE_ACCOUNT_ID and CLOUDFLARE_API_TOKEN must be set")
	}

	rules, err := loadRules(*file)
	if err != nil {
		log.Fatalf("load rules: %v", err)
	}
	log.Printf("Loaded %d rules from %s", len(rules), *file)

	c := &cfClient{accountID: accountID, token: apiToken}

	listID, err := c.findOrCreateList(*listName)
	if err != nil {
		log.Fatalf("find/create list: %v", err)
	}
	log.Printf("Redirect List '%s': %s", *listName, listID)

	items := buildItems(rules, *domain)
	opID, err := c.putListItems(listID, items)
	if err != nil {
		log.Fatalf("PUT list items: %v", err)
	}
	log.Printf("Waiting for operation %s...", opID)
	if err := c.waitOperation(opID); err != nil {
		log.Fatalf("operation failed: %v", err)
	}
	log.Printf("Uploaded %d items.", len(items))

	if err := c.ensureRule(*listName); err != nil {
		log.Fatalf("ensure rule: %v", err)
	}
	log.Println("Sync complete!")
}

func loadRules(path string) ([]redirect, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var rules []redirect
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		code := 301
		if len(parts) >= 3 {
			fmt.Sscanf(parts[2], "%d", &code)
		}
		rules = append(rules, redirect{Source: parts[0], Target: parts[1], Code: code})
	}
	return rules, sc.Err()
}

func buildItems(rules []redirect, domain string) []map[string]any {
	items := make([]map[string]any, 0, len(rules))
	for _, r := range rules {
		src := r.Source
		if !strings.HasPrefix(src, "http") {
			src = domain + src
		}
		dst := r.Target
		if !strings.HasPrefix(dst, "http") {
			dst = domain + dst
		}
		items = append(items, map[string]any{
			"redirect": map[string]any{
				"source_url":            src,
				"target_url":            dst,
				"status_code":           r.Code,
				"preserve_query_string": false,
				"include_subpaths":      false,
				"subpath_matching":      false,
				"preserve_path_suffix":  false,
			},
		})
	}
	return items
}

func (c *cfClient) do(method, path string, body any) (map[string]any, error) {
	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		r = bytes.NewReader(b)
	}
	req, err := http.NewRequest(method, cfAPI+path, r)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, fmt.Errorf("parse response: %w (body: %s)", err, raw)
	}
	if success, _ := result["success"].(bool); !success {
		return nil, fmt.Errorf("API error (HTTP %d): %s", resp.StatusCode, raw)
	}
	return result, nil
}

func (c *cfClient) findOrCreateList(name string) (string, error) {
	result, err := c.do("GET", fmt.Sprintf("/accounts/%s/rules/lists", c.accountID), nil)
	if err != nil {
		return "", err
	}
	for _, item := range result["result"].([]any) {
		lst := item.(map[string]any)
		if lst["name"] == name {
			return lst["id"].(string), nil
		}
	}
	result, err = c.do("POST", fmt.Sprintf("/accounts/%s/rules/lists", c.accountID), map[string]any{
		"name":        name,
		"kind":        "redirect",
		"description": "Static redirects for bmf-tech.com",
	})
	if err != nil {
		return "", err
	}
	return result["result"].(map[string]any)["id"].(string), nil
}

func (c *cfClient) putListItems(listID string, items []map[string]any) (string, error) {
	result, err := c.do("PUT", fmt.Sprintf("/accounts/%s/rules/lists/%s/items", c.accountID, listID), items)
	if err != nil {
		return "", err
	}
	return result["result"].(map[string]any)["operation_id"].(string), nil
}

func (c *cfClient) waitOperation(opID string) error {
	for i := range 60 {
		time.Sleep(3 * time.Second)
		result, err := c.do("GET", fmt.Sprintf("/accounts/%s/rules/lists/bulk_operations/%s", c.accountID, opID), nil)
		if err != nil {
			return err
		}
		status := result["result"].(map[string]any)["status"].(string)
		log.Printf("  [%ds] %s", (i+1)*3, status)
		if status == "completed" {
			return nil
		}
		if status == "failed" || status == "error" {
			return fmt.Errorf("operation %s: %v", status, result["result"])
		}
	}
	return fmt.Errorf("timeout waiting for operation")
}

func (c *cfClient) ensureRule(listName string) error {
	result, err := c.do("GET", fmt.Sprintf("/accounts/%s/rulesets", c.accountID), nil)
	if err != nil {
		return err
	}

	rule := map[string]any{
		"ref":         "bulk_redirect_" + listName,
		"expression":  fmt.Sprintf("http.request.full_uri in $%s", listName),
		"description": fmt.Sprintf("Bulk Redirect rule for %s", listName),
		"action":      "redirect",
		"action_parameters": map[string]any{
			"from_list": map[string]any{
				"name": listName,
				"key":  "http.request.full_uri",
			},
		},
	}

	var rsID string
	var existingRules []any
	for _, item := range result["result"].([]any) {
		rs := item.(map[string]any)
		if rs["phase"] == "http_request_redirect" {
			rsID = rs["id"].(string)
			break
		}
	}

	if rsID != "" {
		rs, err := c.do("GET", fmt.Sprintf("/accounts/%s/rulesets/%s", c.accountID, rsID), nil)
		if err != nil {
			return err
		}
		rsDef := rs["result"].(map[string]any)
		if rules, ok := rsDef["rules"].([]any); ok {
			existingRules = rules
			for _, r := range existingRules {
				rm := r.(map[string]any)
				ap, _ := rm["action_parameters"].(map[string]any)
				fl, _ := ap["from_list"].(map[string]any)
				if fl["name"] == listName {
					log.Printf("Bulk Redirect Rule already exists in ruleset %s — no change.", rsID)
					return nil
				}
			}
		}
		existingRules = append(existingRules, rule)
		_, err = c.do("PUT", fmt.Sprintf("/accounts/%s/rulesets/%s", c.accountID, rsID), map[string]any{
			"name":  rsDef["name"],
			"kind":  rsDef["kind"],
			"phase": rsDef["phase"],
			"rules": existingRules,
		})
		if err != nil {
			return err
		}
		log.Printf("Added Bulk Redirect Rule to existing ruleset %s.", rsID)
		return nil
	}

	newRS, err := c.do("POST", fmt.Sprintf("/accounts/%s/rulesets", c.accountID), map[string]any{
		"name":  "bmf-tech bulk redirects",
		"kind":  "root",
		"phase": "http_request_redirect",
		"rules": []any{rule},
	})
	if err != nil {
		return err
	}
	log.Printf("Created new ruleset: %s", newRS["result"].(map[string]any)["id"])
	return nil
}
