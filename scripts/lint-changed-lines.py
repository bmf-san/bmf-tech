#!/usr/bin/env python3
"""Run textlint on posts changed since BASE, reporting only problems on added/modified lines.

New posts are checked in full (every line is added). Editing an old post only
surfaces problems on the lines you touched, not its pre-existing ones.

Usage:
  scripts/lint-changed-lines.py            # BASE defaults to origin/main
  BASE=main scripts/lint-changed-lines.py
"""

import json
import os
import re
import subprocess
import sys

BASE = os.environ.get("BASE", "origin/main")
CONFIGS = {"ja": ".textlintrc-ja.json", "en": ".textlintrc-en.json"}
HUNK = re.compile(r"^@@ -\d+(?:,\d+)? \+(\d+)(?:,(\d+))? @@")


def git(*args):
    return subprocess.run(["git", *args], check=True, capture_output=True, text=True).stdout


def changed_lines(path):
    lines = set()
    for row in git("diff", "-U0", f"{BASE}...HEAD", "--", path).splitlines():
        m = HUNK.match(row)
        if m:
            start, count = int(m.group(1)), int(m.group(2) or 1)
            lines.update(range(start, start + count))
    return lines


def main():
    failed = 0
    for lang, config in CONFIGS.items():
        files = git("diff", "--name-only", "--diff-filter=ACM", f"{BASE}...HEAD",
                    "--", f"content/{lang}/posts/*.md").split()
        targets = {f: changed_lines(f) for f in files}
        targets = {f: ls for f, ls in targets.items() if ls}
        if not targets:
            print(f"No changed {lang.upper()} posts to lint.")
            continue
        out = subprocess.run(["npx", "textlint", "--config", config, "-f", "json", *targets],
                             capture_output=True, text=True)
        try:
            results = json.loads(out.stdout)
        except json.JSONDecodeError:
            sys.stderr.write(out.stdout + out.stderr)
            return 2
        for result in results:
            path = os.path.relpath(result["filePath"])
            hits = [m for m in result["messages"] if m["line"] in targets.get(path, ())]
            if not hits:
                continue
            print(path)
            for m in hits:
                print(f"  {m['line']}:{m['column']}  {m['message']}  {m['ruleId']}")
            failed += len(hits)
        print(f"{lang.upper()}: checked changed lines in {len(targets)} file(s)")
    if failed:
        print(f"\n✖ {failed} problem(s) on changed lines")
        return 1
    return 0


if __name__ == "__main__":
    sys.exit(main())
