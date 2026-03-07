---
title: GoReleaserを使ってDocker Imageをbuild&pushする
slug: goreleaser-docker-image-build-push
date: 2024-05-29T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
  - GitHub Actions
translation_key: goreleaser-docker-image-build-push
---


以前、[GoReleaserを使ってのGoのアプリケーションを配布する](https://bmf-tech.com/posts/GoReleaser%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E3%81%AEGo%E3%81%AE%E3%82%A2%E3%83%97%E3%83%AA%E3%82%B1%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%E3%82%92%E9%85%8D%E5%B8%83%E3%81%99%E3%82%8B)というタイトルの記事でGoReleaserを使ったバイナリ配布の方法について書いたが、Dockerhubへのイメージプッシュもやってみたのでまとめておく。

ソースコードの全体像は[bmf-san/gondola](https://github.com/bmf-san/gondola)を参照。

# .goreleaser.yamlを設定
```yaml
ockers:
  - image_templates:
      - bmfsan/gondola:latest-amd64
      - bmfsan/gondola:{{ .Version }}-amd64
    use: buildx
    goos: linux
    goarch: amd64
    build_flag_templates:
      - --platform=linux/amd64
  - image_templates:
      - bmfsan/gondola:latest-arm64
      - bmfsan/gondola:{{ .Version }}-arm64
    use: buildx
    goos: linux
    goarch: arm64
    build_flag_templates:
      - --platform=linux/arm64
  - image_templates:
      - bmfsan/gondola:latest-arm
      - bmfsan/gondola:{{ .Version }}-arm
    use: buildx
    goos: linux
    goarch: arm
    build_flag_templates:
      - --platform=linux/arm
```

ちょっと長たらしいが、こんな感じでdockersのオプションを書く。

長らしくなっているのは、builds（バイナリ作成）の方のオプションと同じように成果物を作成するように調整しているためである。（もうちょっとまとめて書いたりできるのかもしれないが・・）

```yaml
builds:
  - env:
      - CGO_ENABLED=0
    goos:
      - linux
      - darwin
      - windows
    goarch:
      - amd64
      - arm
      - arm64
    main: ./cmd/main.go
```

goreleaserのデフォルト値は以下を参照。

cf. [github.com/goreleaser/goreleaser/blob/main/.goreleaser.yaml](https://github.com/goreleaser/goreleaser/blob/main/.goreleaser.yaml)

# Dockerfileの作成
```
FROM gcr.io/distroless/static-debian12
COPY gondola /
ENTRYPOINT ["./gondola", "-config", "config.yaml"]
```

COPYしているバイナリは、goreleaserがバイナリを生成してくれるので、自分でビルドする必要はない。

# Github Actionsのworkflowを編集
```yaml
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
        go-version: [ '1.22.3' ]
    steps:
      - name: Checkout
        uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: ${{ matrix.go-version }}
      - name: Login to Docker Hub
        uses: docker/login-action@v3
        with:
          username: ${{ secrets.DOCKERHUB_USERNAME }}
          password: ${{ secrets.DOCKERHUB_TOKEN }}
      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v5
        with:
          distribution: goreleaser
          version: latest
          args: release --clean
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

DockerhubへのイメージプッシュがしたいのでDockerhubへログインするように調整。

secretsはGithubのリポジトリの設定から事前に設定しておく。



と、こんな感じで簡単に設定できる！ GoReleaser便利・・！！
