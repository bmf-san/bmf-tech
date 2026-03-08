---
title: Introduction to the Go-based Git Tool 'ggc'
slug: introducing-go-git-tool-ggc
date: 2025-06-15T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Git
  - CLI
  - TUI
translation_key: introducing-go-git-tool-ggc
---

# Introduction to the Go-based Git Tool 'ggc'

## What is ggc?

[ggc](https://github.com/bmf-san/ggc) is a Git operation support tool implemented in Go. Its purpose is to be "easy to remember, easy to use, and improve work efficiency," aiming to make everyday Git operations more comfortable.

Existing Git client tools often have either too many features, resulting in a high learning cost, or are too simple to be practical. ggc aims to fill this gap by providing a **command system that is simple and easy to remember, focusing on the functions used daily**.

### Features

1. **Dual Interface**: Fast operations via CLI, intuitive operations
2. **Compound Commands**: Execute multiple Git operations with a single command
3. **Incremental Search**: Select commands without memorizing them

### Convenience Comparison

| Regular Git Operations                                     | Operations with ggc                   |
| ---------------------------------------------------- | --------------------------------- |
| `git add .` → `git commit -m "..."` → `git push` | `ggc add-commit-push`           |
| `git branch` → `git checkout branch-name`              | `ggc branch checkout` (interactive selection) |
| `git stash` → `git pull` → `git stash pop`         | `ggc stash-pull-pop`            |

As shown above, common operations can be executed concisely with a single command.

## Main Features

* **Dual Interface**: Execute commands directly with arguments, or start interactive mode without arguments
* **Interactive Operations**: Supports branch/file selection and commit message input
* **Rich Command Set**: Covers basic Git operations
* **Compound Commands**: Provides commands like `add-commit-push`, `stash-pull-pop`, etc.
* **Lightweight Design**: Uses only the Go standard library and `golang.org/x/term`
* **Operating Environment**: Confirmed on macOS (Apple Silicon/Intel)

## Usage Examples

```bash
# Update to the latest state
ggc pull current

# Start working on a new branch (select interactively)
ggc branch checkout
```

```bash
# Push changes in bulk
ggc add-commit-push
```

```bash
# Merge safely
ggc stash-pull-pop
```

## Installation Steps

### Installation via `go install`

The easiest way to install is as follows:

```sh
go install github.com/bmf-san/ggc@latest
```

Set the PATH if necessary:

```sh
export PATH=$PATH:$(go env GOBIN)
```

### Build from Source

```sh
git clone https://github.com/bmf-san/ggc
cd ggc
make build
```

After building, place the generated binary in a directory included in your PATH.

## How to Use

### Switching Between CLI and Interactive Mode

The CLI or interactive mode is automatically started based on the presence of arguments.

```sh
# CLI (specify command directly)
ggc branch current

# Interactive mode
ggc
```

Both modes are supported by a single binary, allowing flexible operations according to your needs.

### Command Selection in Interactive Mode

When `ggc` is executed without arguments, a command selection screen appears with incremental search.

```sh
ggc
```

Example display:

```
Please select a command (incremental search: narrow down by typing, ctrl+n: move down, ctrl+p: move up, enter: execute, ctrl+c: exit)
Search: branch

> branch current
  branch checkout
  branch checkout-remote
  branch delete
  branch delete-merged
```

Operation steps:

* Typing narrows down the candidates
* Use `Ctrl+n`/`Ctrl+p` to move up and down
* Press `Enter` to execute
* If arguments are needed, a prompt will be displayed
* After execution, return to the selection screen after checking the results

There is no need to memorize commands, as candidates are displayed based on input, allowing for intuitive operation.

### Representative Commands

| ggc Command                      | Actual git Command                                                         | Description                     |
| ---------------------------- | ------------------------------------------------------------------- | ---------------------- |
| `ggc add <file>`             | `git add <file>`                                                    | Stage a file            |
| `ggc add .`                  | `git add .`                                                         | Stage all files           |
| `ggc add -p`                 | `git add -p`                                                        | Interactive staging              |
| `ggc branch current`         | `git rev-parse --abbrev-ref HEAD`                                   | Get current branch name             |
| `ggc branch checkout`        | `git branch ... → git checkout <selection>`                                | Interactive branch switching            |
| `ggc branch checkout-remote` | `git branch -r ... → git checkout -b <n> --track <remote>/<branch>` | Create/switch from remote branch    |
| `ggc branch delete`          | `git branch ... → git branch -d <selection>`                               | Interactively delete local branch         |
| `ggc push current`           | `git push origin <branch>`                                          | Push current branch           |
| `ggc pull current`           | `git pull origin <branch>`                                          | Pull current branch             |
| `ggc log simple`             | `git log --oneline`                                                 | Simple log display              |
| `ggc commit <message>`       | `git commit -m <message>`                                           | Create a commit                 |
| `ggc fetch --prune`          | `git fetch --prune`                                                 | Fetch while deleting old remote tracking branches |
| `ggc clean files`            | `git clean -f`                                                      | Clean up files           |
| `ggc remote add <n> <url>`   | `git remote add <n> <url>`                                          | Add remote                 |
| `ggc stash`                  | `git stash`                                                         | Temporarily stash work              |
| `ggc rebase interactive`     | `git rebase -i`                                                     | Interactive rebase                |

### Examples of Compound Commands

| ggc Command                       | Executed Git Operations                             | Description                      |
| ----------------------------- | -------------------------------------- | ----------------------- |
| `ggc add-commit-push`         | `git add . → git commit → git push`    | Execute stage → commit → push in one go |
| `ggc commit-push-interactive` | Interactive stage → commit → push                  |                         |
| `ggc pull-rebase-push`        | `git pull → git rebase → git push`     | Execute pull → rebase → push in one go   |
| `ggc stash-pull-pop`          | `git stash → git pull → git stash pop` | Execute stash → pull → restore in one go     |

## Completion Scripts

Completion scripts for Bash and Zsh are included.

### Setup Method

```sh
# For bash
source /path/to/ggc/tools/completions/ggc.bash

# For zsh (same script can be used)
source /path/to/ggc/tools/completions/ggc.bash
```

By adding this to `.bashrc` or `.zshrc`, completion will be enabled when the terminal starts.

## Summary

* Intuitive operations without memorizing commands
* Execute routine tasks with a single command
* Interactive support for branch and file selection
* Improved work efficiency through compound commands

### Related Links

* **GitHub Repository**: [https://github.com/bmf-san/ggc](https://github.com/bmf-san/ggc)
* **Issues, Feature Requests, Bug Reports, etc.**: [https://github.com/bmf-san/ggc/issues](https://github.com/bmf-san/ggc/issues)