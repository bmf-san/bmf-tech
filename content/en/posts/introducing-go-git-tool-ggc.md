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

[ggc](https://github.com/bmf-san/ggc) is a Git operation support tool implemented in Go. It aims to "be easy to remember, easy to use, and improve work efficiency," making everyday Git operations more comfortable.

Existing Git client tools can be either too feature-rich, resulting in high learning costs, or too simple to be practical. ggc bridges this gap by providing a **simple and memorable command system focused on daily-use features**.

### Features

1. **Dual Interface**: Fast operation via CLI, intuitive operation
2. **Composite Commands**: Execute multiple Git operations with a single command
3. **Incremental Search**: Select commands without memorizing them

### Convenience Comparison

| Standard Git Operations                                      | Operations with ggc                |
| ------------------------------------------------------------ | ---------------------------------- |
| `git add .` → `git commit -m "..."` → `git push` | `ggc add-commit-push`              |
| `git branch` → `git checkout <branch>`                       | `ggc branch checkout` (interactive selection) |
| `git stash` → `git pull` → `git stash pop`                   | `ggc stash-pull-pop`               |

As shown above, common operations can be executed concisely with a single command.

## Main Features

* **Dual Interface**: Execute commands directly with arguments, launch interactive mode without arguments
* **Interactive Operations**: Supports branch/file selection and commit message input
* **Rich Command Set**: Covers basic Git operations
* **Composite Commands**: Offers `add-commit-push`, `stash-pull-pop`, etc.
* **Lightweight Design**: Uses only Go standard library and `golang.org/x/term`
* **Operating Environment**: Confirmed on macOS (Apple Silicon/Intel)

## Usage Examples

```bash
# Update to the latest state
ggc pull current

# Start working on a new branch (interactive selection)
ggc branch checkout
```

```bash
# Push changes in bulk
ggc add-commit-push
```

```bash
# Safe merge
ggc stash-pull-pop
```

## Installation Instructions

### Installation via `go install`

The simplest installation method is as follows:

```sh
go install github.com/bmf-san/ggc@latest
```

Set the PATH as needed:

```sh
export PATH=$PATH:$(go env GOBIN)
```

### Building from Source

```sh
git clone https://github.com/bmf-san/ggc
cd ggc
make build
```

After building, place the generated binary in a directory included in PATH.

## Usage

### Switching Between CLI and Interactive Mode

CLI or interactive mode is automatically launched depending on the presence of arguments.

```sh
# CLI (specify command directly)
ggc branch current

# Interactive mode
ggc
```

Both modes are supported by a single binary, allowing flexible operations according to the use case.

### Command Selection in Interactive Mode

Executing `ggc` without arguments displays a command selection screen with incremental search.

```sh
ggc
```

Example display:

```
Select a command (Incremental Search: Narrow down by typing, ctrl+n: move down, ctrl+p: move up, enter: execute, ctrl+c: exit)
Search: branch

> branch current
  branch checkout
  branch checkout-remote
  branch delete
  branch delete-merged
```

Operation steps:

* Narrow down candidates by typing
* Move up/down with `Ctrl+n`/`Ctrl+p`
* Execute with `Enter`
* Prompt displayed if arguments are needed
* Return to selection screen after confirming results

There's no need to memorize commands, as candidates are displayed based on input, allowing intuitive operation.

### Representative Commands

| ggc Command                      | Actual git Command                                                          | Description                  |
| -------------------------------- | --------------------------------------------------------------------------- | ---------------------------- |
| `ggc add <file>`                 | `git add <file>`                                                            | Stage a file                 |
| `ggc add .`                      | `git add .`                                                                 | Stage all files              |
| `ggc add -p`                     | `git add -p`                                                                | Interactive staging          |
| `ggc branch current`             | `git rev-parse --abbrev-ref HEAD`                                           | Get current branch name      |
| `ggc branch checkout`            | `git branch ... → git checkout <selection>`                                 | Interactive branch switch    |
| `ggc branch checkout-remote`     | `git branch -r ... → git checkout -b <n> --track <remote>/<branch>`         | Create/switch from remote branch |
| `ggc branch delete`              | `git branch ... → git branch -d <selection>`                                | Interactive local branch deletion |
| `ggc push current`               | `git push origin <branch>`                                                  | Push current branch          |
| `ggc pull current`               | `git pull origin <branch>`                                                  | Pull current branch          |
| `ggc log simple`                 | `git log --oneline`                                                         | Simple log display           |
| `ggc commit <message>`           | `git commit -m <message>`                                                   | Create a commit              |
| `ggc fetch --prune`              | `git fetch --prune`                                                         | Fetch while removing old remote-tracking branches |
| `ggc clean files`                | `git clean -f`                                                              | Clean up files               |
| `ggc remote add <n> <url>`       | `git remote add <n> <url>`                                                  | Add a remote                 |
| `ggc stash`                      | `git stash`                                                                 | Temporarily save work        |
| `ggc rebase interactive`         | `git rebase -i`                                                             | Interactive rebase           |

### Examples of Composite Commands

| ggc Command                       | Executed Git Operations                             | Description                  |
| --------------------------------- | --------------------------------------------------- | ---------------------------- |
| `ggc add-commit-push`             | `git add . → git commit → git push`                 | Execute stage → commit → push in one go |
| `ggc commit-push-interactive`     | Interactive stage → commit → push                   |                              |
| `ggc pull-rebase-push`            | `git pull → git rebase → git push`                  | Execute pull → rebase → push in one go |
| `ggc stash-pull-pop`              | `git stash → git pull → git stash pop`              | Execute stash → pull → restore in one go |

## Completion Script

Completion scripts for Bash and Zsh are included.

### Setup Method

```sh
# For bash
source /path/to/ggc/tools/completions/ggc.bash

# For zsh (same script can be used)
source /path/to/ggc/tools/completions/ggc.bash
```

By adding this to `.bashrc` or `.zshrc`, completion is enabled when the terminal starts.

## Summary

* Intuitive operation without memorizing commands
* Execute routine tasks with a single command
* Interactive support for branch and file selection
* Composite commands can improve work efficiency

### Related Links

* **GitHub Repository**: [https://github.com/bmf-san/ggc](https://github.com/bmf-san/ggc)
* **Issues, Feature Requests, Bug Reports**: [https://github.com/bmf-san/ggc/issues](https://github.com/bmf-san/ggc/issues)
