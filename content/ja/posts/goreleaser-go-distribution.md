---
title: "GoReleaserを使ってのGoのアプリケーションを配布する"
slug: "goreleaser-go-distribution"
date: 2023-11-11
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Golang"
  - "GitHub Actions"
draft: false
---

[GoReleaser](https://goreleaser.com/)を使ってGoのアプリケーションをクロスコンパイルしてバイナリ配布をやってみた。

# GoReleaserとは
[GoReleaser](https://goreleaser.com/)は、Go言語で書かれたアプリケーションのビルド、パッケージング、およびリリースを自動化するツール。

クロスコンパイル、バイナリの圧縮、アーカイブの作成、GitHubなどのプラットフォームへのアーティファクトのアップロードができる。

# GitHub Actionsを使ってバイナリを配布する
Github ActionsにGoReleaserの公式Actionが用意されているので、それを使うことができる。GoReleaserは設定ファイルを用意することもできるが、特に用意しなくても使うことができる。

cmdディレクトリ配下をビルドする想定でworkflowの実装例を記載する。

## Dry run
ビルドしてバイナリを配布することができるかどうかCIのプロセスに組み込んでおくと、リリースする際に配布できなかった・・なんてことが避けれる。

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
タグリリース時にバイナリ配布を実行する。このジョブが完了すると、GitHubのリリースタグのページに成果物が添付される。

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

# バイナリ配布例
これはまだ開発中のアプリケーションだが、こんな感じで配布することできる。

https://github.com/bmf-san/gondola/releases/tag/0.0.3

# 所感
アプリケーションの実装がツールに依存することもなく、簡単に使うことができるので気に入った。

類似のツールは他にもあるが、とりあえずGoReleaserをしばらく使ってみようと思う。
