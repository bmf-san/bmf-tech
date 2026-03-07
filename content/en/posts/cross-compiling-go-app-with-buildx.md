---
title: Cross-compiling Go Application Images with buildx
slug: cross-compiling-go-app-with-buildx
date: 2023-04-20T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Docker
description: A guide on using Docker buildx for cross-compiling Go application images.
translation_key: cross-compiling-go-app-with-buildx
---

# Overview
Driven by the need to cross-compile the image of an application I am developing privately (due to differences in architecture between the local development environment and the production environment), I implemented the solution and took notes.

# buildx
Docker Desktop comes with buildx as a standard feature, so I used it.

[Docker Buildx](https://matsuand.github.io/docs.docker.jp.onthefly/buildx/working-with-buildx/)

With buildx, you can easily create multi-architecture compatible images.

# Example
Let's say we have a Dockerfile like this (this is the actual Dockerfile I use):

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

For information on environment variables, refer to the link below.
cf. https://matsuand.github.io/docs.docker.jp.onthefly/engine/reference/builder/

Building and pushing can be done as follows. You can specify multiple platforms.

```sh
// Create a builder instance
docker buildx create --name buildx-builder
docker buildx use buildx-builder

// Build and push to Docker Hub
docker buildx build --no-cache --push --platform linux/amd64,linux/arm64 --file app/Dockerfile --tag bmfsan/gobel-api app/
```

# Miscellaneous
The official MySQL image has, at some point, become compatible with ARM. This is great news for M1 users.

https://hub.docker.com/layers/library/mysql/8.0.29/images/sha256-44f98f4dd825a945d2a6a4b7b2f14127b5d07c5aaa07d9d232c2b58936fb76dc?context=explore

# References
- [AverageMarcus/Dockerfile](https://gist.github.com/AverageMarcus/78fbcf45e72e09d9d5e75924f0db4573)