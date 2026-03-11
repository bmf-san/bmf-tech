---
title: Distributing Go Applications with GoReleaser
slug: goreleaser-go-distribution
date: 2023-11-11T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - GitHub Actions
description: Using GoReleaser to cross-compile and distribute Go application binaries.
translation_key: goreleaser-go-distribution
---



[GoReleaser](https://goreleaser.com/) was used to cross-compile and distribute binaries for a Go application.

# What is GoReleaser
[GoReleaser](https://goreleaser.com/) is a tool that automates the build, packaging, and release of applications written in Go.

It can perform cross-compilation, compress binaries, create archives, and upload artifacts to platforms like GitHub.

# Distributing Binaries with GitHub Actions
GitHub Actions provides an official GoReleaser Action, which can be used. Although GoReleaser can be configured with a settings file, it can also be used without one.

Below is an example workflow implementation assuming a build under the cmd directory.

## Dry run
Incorporating a dry run into the CI process ensures that binaries can be distributed when releasing, avoiding distribution failures.

```yaml
name: Dry run GoReleaser

on: [push]

permissions:
  contents: write

jobs:
  goreleaser:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        go-version: [ '1.21.x' ]
    steps:
      - name: Checkout
        uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: ${{ matrix.go-version }}
      - name: Dry run GoReleaser
        uses: goreleaser/goreleaser-action@v5
        with:
          distribution: goreleaser
          version: latest
          args: release --rm-dist --skip-publish --snapshot
          workdir: cmd
```

## Release
Execute binary distribution upon tag release. Once this job is complete, artifacts will be attached to the GitHub release tag page.

```yml
name: GoReleaser

on:
  push:
    tags:
      - '*'

permissions:
  contents: write

jobs:
  goreleaser:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        go-version: [ '1.21.x' ]
    steps:
      - name: Checkout
        uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: ${{ matrix.go-version }}
      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v5
        with:
          distribution: goreleaser
          version: latest
          args: release --clean
          workdir: cmd
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

# Example of Binary Distribution
This is still a developing application, but it can be distributed like this.

https://github.com/bmf-san/gondola/releases/tag/0.0.3

# Impressions
I liked that the implementation of the application does not depend on the tool and is easy to use.

There are other similar tools, but I plan to use GoReleaser for a while.