#!/usr/bin/env python3
"""Sync bulk-redirects.txt to Cloudflare Bulk Redirects via API.

Required environment variables:
  CLOUDFLARE_ACCOUNT_ID  - Cloudflare account ID
  CLOUDFLARE_API_TOKEN   - API token with permissions:
                           Account > Bulk URL Redirects > Edit
                           Account > Account Filter Lists > Edit
"""

import json
import os
import sys
import time
import urllib.error
import urllib.request

ACCOUNT_ID = os.environ["CLOUDFLARE_ACCOUNT_ID"]
API_TOKEN = os.environ["CLOUDFLARE_API_TOKEN"]
BASE_DOMAIN = "https://bmf-tech.com"
LIST_NAME = "bmf-tech-redirects"
SOURCE_FILE = "bulk-redirects.txt"

HEADERS = {
    "Authorization": f"Bearer {API_TOKEN}",
    "Content-Type": "application/json",
}


def cf_request(method, path, body=None):
    url = f"https://api.cloudflare.com/client/v4{path}"
    data = json.dumps(body).encode("utf-8") if body is not None else None
    req = urllib.request.Request(url, data=data, method=method, headers=HEADERS)
    try:
        with urllib.request.urlopen(req) as resp:
            result = json.loads(resp.read())
            if not result.get("success"):
                print(f"API error: {result.get('errors')}", file=sys.stderr)
                sys.exit(1)
            return result
    except urllib.error.HTTPError as e:
        body_text = e.read().decode()
        print(f"HTTP {e.code} for {method} {path}: {body_text}", file=sys.stderr)
        sys.exit(1)


def wait_for_operation(op_id):
    for i in range(60):
        time.sleep(3)
        result = cf_request(
            "GET", f"/accounts/{ACCOUNT_ID}/rules/lists/bulk_operations/{op_id}"
        )
        status = result["result"]["status"]
        print(f"  [{i * 3}s] {status}")
        if status == "completed":
            return
        if status in ("failed", "error"):
            print(f"Operation failed: {result['result']}", file=sys.stderr)
            sys.exit(1)
    print("Timeout waiting for operation", file=sys.stderr)
    sys.exit(1)


# ── 1. Load rules ────────────────────────────────────────────────────────────
rules = []
with open(SOURCE_FILE, encoding="utf-8") as f:
    for line in f:
        line = line.strip()
        if not line or line.startswith("#"):
            continue
        parts = line.split()
        if len(parts) >= 2:
            src = parts[0]
            dst = parts[1]
            code = int(parts[2]) if len(parts) >= 3 else 301
            rules.append((src, dst, code))

print(f"Loaded {len(rules)} rules from {SOURCE_FILE}")

# ── 2. Find or create Redirect List ─────────────────────────────────────────
result = cf_request("GET", f"/accounts/{ACCOUNT_ID}/rules/lists")
existing = [lst for lst in result["result"] if lst["name"] == LIST_NAME]

if existing:
    list_id = existing[0]["id"]
    print(f"Found existing list '{LIST_NAME}': {list_id}")
else:
    result = cf_request(
        "POST",
        f"/accounts/{ACCOUNT_ID}/rules/lists",
        {
            "name": LIST_NAME,
            "kind": "redirect",
            "description": "Static redirects for bmf-tech.com",
        },
    )
    list_id = result["result"]["id"]
    print(f"Created new list '{LIST_NAME}': {list_id}")

# ── 3. Replace all items (PUT is atomic/idempotent) ──────────────────────────
items = []
for src, dst, code in rules:
    full_src = BASE_DOMAIN + src
    full_dst = dst if dst.startswith("http") else BASE_DOMAIN + dst
    items.append(
        {
            "redirect": {
                "source_url": full_src,
                "target_url": full_dst,
                "status_code": code,
                "preserve_query_string": False,
                "include_subpaths": False,
                "subpath_matching": False,
                "preserve_path_suffix": False,
            }
        }
    )

print(f"Uploading {len(items)} items to list '{LIST_NAME}'...")
result = cf_request(
    "PUT", f"/accounts/{ACCOUNT_ID}/rules/lists/{list_id}/items", items
)
op_id = result["result"]["operation_id"]
print(f"Waiting for operation {op_id}...")
wait_for_operation(op_id)
print(f"Items uploaded successfully.")

# ── 4. Ensure Bulk Redirect Rule exists ──────────────────────────────────────
result = cf_request("GET", f"/accounts/{ACCOUNT_ID}/rulesets")
redirect_rulesets = [
    r for r in result.get("result", []) if r.get("phase") == "http_request_redirect"
]

rule_entry = {
    "ref": f"bulk_redirect_{LIST_NAME}",
    "expression": f"http.request.full_uri in ${LIST_NAME}",
    "description": f"Bulk Redirect rule for {LIST_NAME}",
    "action": "redirect",
    "action_parameters": {
        "from_list": {
            "name": LIST_NAME,
            "key": "http.request.full_uri",
        }
    },
}

if redirect_rulesets:
    rs_id = redirect_rulesets[0]["id"]
    rs = cf_request("GET", f"/accounts/{ACCOUNT_ID}/rulesets/{rs_id}")
    existing_rules = rs["result"].get("rules", [])
    our_rules = [
        r
        for r in existing_rules
        if r.get("action_parameters", {}).get("from_list", {}).get("name") == LIST_NAME
    ]
    if our_rules:
        print(f"Bulk Redirect Rule already exists in ruleset {rs_id} — no change needed.")
    else:
        existing_rules.append(rule_entry)
        cf_request(
            "PUT",
            f"/accounts/{ACCOUNT_ID}/rulesets/{rs_id}",
            {
                "name": rs["result"]["name"],
                "kind": rs["result"]["kind"],
                "phase": rs["result"]["phase"],
                "rules": existing_rules,
            },
        )
        print(f"Added Bulk Redirect Rule to existing ruleset {rs_id}.")
else:
    result = cf_request(
        "POST",
        f"/accounts/{ACCOUNT_ID}/rulesets",
        {
            "name": "bmf-tech bulk redirects",
            "kind": "root",
            "phase": "http_request_redirect",
            "rules": [rule_entry],
        },
    )
    print(f"Created new ruleset: {result['result']['id']}")

print("Sync complete!")
