---
title: Introducing ggc — A Go-Based Git Tool (2026 Edition)
description: 'A complete walkthrough of ggc v8: the CLI/interactive split architecture, fuzzy-search engine implementation, Workflow Mode internals, customisable aliases, and cross-platform keybinding profiles.'
slug: introducing-go-git-tool-ggc
date: 2025-06-15T00:00:00Z
lastmod: 2026-03-20
author: bmf-san
categories:
  - Tools
tags:
  - Golang
  - Git
  - CLI
  - TUI
translation_key: introducing-go-git-tool-ggc
---

# Introducing ggc — A Go-Based Git Tool (2026 Edition)

## What is ggc?

[ggc](https://github.com/bmf-san/ggc) is a Git workflow tool written in Go. It wraps everyday Git sub-commands under a consistent surface and adds an interactive fuzzy-search TUI so you can find and execute commands without memorising their exact names. Version 8 introduced Workflow Mode, customisable aliases with placeholder support, and a layered keybinding profile system — this article covers all of them.

## CLI/Interactive Architecture

ggc has two execution paths: the **CLI path** for direct commands and the **interactive path** for the full-screen TUI. The `Execute()` method in `cmd/execute.go` is the single entry point that decides which path to take:

```go
func (c *Cmd) Execute(args []string) error {
    if len(args) == 0 {
        c.Interactive()
        return nil
    }

    cmdName, cmdArgs := args[0], args[1:]

    // Check if this is an alias
    if c.configManager != nil && c.configManager.GetConfig().IsAlias(cmdName) {
        return c.executeAlias(cmdName, cmdArgs)
    }

    // Regular command
    return c.Route(args)
}
```

When no arguments are supplied, `Interactive()` launches the TUI. When an alias name is detected, `executeAlias()` resolves and runs it. Everything else goes to `Route()`, which dispatches to individual Git command handlers. The same `Route()` function is shared between the CLI path, the alias expander, and the Workflow executor.

## Interactive Mode Design

The interactive TUI is split into two sub-modes that share a single terminal screen, toggled with `Ctrl+t`.

### Search Mode — Fuzzy Search

Search Mode is the default. It shows a list of all ggc commands and updates the filtered list as you type. The scoring is done by the `matchPattern()` function in `internal/interactive/fuzzy.go`:

```go
func matchPattern(textRunes, patternRunes []rune) (bool, matchMetadata) {
    meta := matchMetadata{firstIndex: -1, lastIndex: -1}
    textIdx := 0
    patternIdx := 0

    for textIdx < len(textRunes) && patternIdx < len(patternRunes) {
        if textRunes[textIdx] == patternRunes[patternIdx] {
            if meta.firstIndex == -1 {
                meta.firstIndex = textIdx
            }
            if meta.lastIndex != -1 {
                meta.gapScore += textIdx - meta.lastIndex - 1
            }
            meta.lastIndex = textIdx
            patternIdx++
        }
        textIdx++
    }

    if patternIdx != len(patternRunes) {
        return false, meta
    }
    return true, meta
}
```

The algorithm is a classic *subsequence* fuzzy matcher: a pattern matches if all its characters appear in the text in order, but not necessarily consecutively. The `matchMetadata` struct accumulates three signals — `firstIndex` (where the match starts), `gapScore` (total inter-character gaps), and `lastIndex` — which are combined into a `matchScore` for sorting. A tighter, earlier match ranks higher than a sparse, late one.

### Workflow Mode

Pressing `Ctrl+t` calls `ToggleWorkflowView()` in `internal/interactive/ui_mode.go`:

```go
func (ui *UI) ToggleWorkflowView() {
    if ui.state.IsWorkflowMode() {
        ui.enterSearchMode()
        return
    }
    ui.enterWorkflowMode()
}
```

In Workflow Mode you can build up a sequence of ggc commands (e.g. `add` → `commit` → `push`) by pressing `Tab` on any item in the search list. When you press `x`, the `WorkflowExecutor.Execute()` method runs them in order:

```go
func (we *WorkflowExecutor) Execute(workflow *Workflow) error {
    steps := workflow.GetSteps()
    if len(steps) == 0 {
        return fmt.Errorf("workflow is empty")
    }

    for i, step := range steps {
        resolvedArgs, canceled := resolveStepPlaceholders(we.ui, step)
        if canceled {
            return ErrWorkflowCanceled
        }

        parts := append([]string{step.Command}, resolvedArgs...)
        if err := we.router.Route(parts); err != nil {
            return fmt.Errorf("step %d/%d failed: %w", i+1, len(steps), err)
        }
    }
    return nil
}
```

Before each step, `resolveStepPlaceholders()` scans the step arguments for `<name>` tokens and prompts you interactively if any are found. A workflow step like `commit -m <message>` will pause and ask for the commit message at runtime.

## Other New Features in v8

### Command Aliases

Aliases are defined in `~/.ggcconfig.yaml`. Both simple and sequence formats are supported:

```yaml
aliases:
  br: branch          # simple — ggc br list
  ci: commit

  quick:              # sequence
    - status
    - add .
    - commit

  deploy:             # sequence with placeholders
    - "branch checkout {0}"
    - "pull current"
    - "push {0}"
```

### Keybinding Profiles

The interactive UI supports four built-in keybinding profiles (`default`, `emacs`, `vi`, `readline`). Configuration resolves in layers — defaults → profile → platform → user config:

```yaml
interactive:
  profile: emacs
  keybindings:
    move_up: "ctrl+p"
    move_down: "ctrl+n"
    toggle_workflow_view: "ctrl+t"
    add_to_workflow: "tab"
```

### Cross-Platform Builds

ggc distributes pre-built binaries for Linux, macOS, and Windows (amd64 / arm64 / 386) via GoReleaser. All interactive TUI features work on all three platforms.

## Installation

```bash
# Homebrew (macOS/Linux)
brew install bmf-san/tap/ggc

# Go install
go install github.com/bmf-san/ggc/v8@latest
```

## Summary

ggc v8 is an opinionated but flexible Git workflow tool. The CLI path and the interactive path share the same `Route()` layer, keeping behaviours consistent. The fuzzy-search scorer is a pure function with no dependencies, and the Workflow executor reuses `Route()` rather than duplicating command dispatch logic.

- **GitHub**: [bmf-san/ggc](https://github.com/bmf-san/ggc)
