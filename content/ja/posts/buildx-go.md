---
title: "buildxを使ったGoアプリケーションイメージのクロスコンパイル"
slug: "buildx-go"
date: 2023-04-20
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Golang"
  - "Docker"
draft: false
---

# 概要
プライベートで開発しているアプリケーションのイメージをクロスコンパイルする必要性に駆られて（ローカルの開発環境と本番の環境でアーキテクチャが異なっていることが起因）対応したのでメモ。

# buildx
Docker Desktopにはbuildxが標準で備わっているのでそちらを利用する。

[Docker Buildx](https://matsuand.github.io/docs.docker.jp.onthefly/buildx/working-with-buildx/)

buildxを使うことでマルチアーキテクチャ対応のイメージを簡単に作ることができる。

# 例
こんな感じのDockerfileがあったとする。（実際に使っているDockerfileなのだが..）

```yaml
FROM --platform=$BUILDPLATFORM golang:1.20.0-alpine as builder

WORKDIR /go/gobel-api/app

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

COPY . .

RUN apk update && \
    apk upgrade && \
    apk add --no-cache libc-dev gcc git openssh openssl bash

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o app

FROM --platform=$TARGETPLATFORM alpine

COPY --from=builder /go/gobel-api/app ./

ENTRYPOINT ["/app"]
```

環境変数については下記参照。
cf. https://matsuand.github.io/docs.docker.jp.onthefly/engine/reference/builder/

buildとpushはこんな感じ。platformは複数指定することができる。

```sh
// ビルダーインスタンスの作成
docker buildx create --name buildx-builder
docker buildx use buildx-builder

// ビルドしてdockerhubにpush
docker buildx build --no-cache --push --platform linux/amd64,linux/arm64 --file app/Dockerfile --tag bmfsan/gobel-api app/
```

# 余談
MySQLのオフィシャルイメージがいつの間にかARMにも対応するようになっていた。
M1ユーザーは嬉しい。

https://hub.docker.com/layers/library/mysql/8.0.29/images/sha256-44f98f4dd825a945d2a6a4b7b2f14127b5d07c5aaa07d9d232c2b58936fb76dc?context=explore

# 参考
- [AverageMarcus/Dockerfile](https://gist.github.com/AverageMarcus/78fbcf45e72e09d9d5e75924f0db4573)

