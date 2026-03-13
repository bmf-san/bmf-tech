---
title: One-Liner for Self-Signed Certificates in Go
description: 'Generate a self-signed TLS certificate in Go with one command using the built-in generate_cert.go tool. Produces cert.pem and key.pem for local HTTPS without openssl or mkcert.'
slug: go-self-signed-certificate-creation
date: 2024-02-16T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Tips
translation_key: go-self-signed-certificate-creation
---

A handy one-liner for when you need a self-signed certificate while writing an HTTP server in Go.

```go
go run $(go env GOROOT)/src/crypto/tls/generate_cert.go -rsa-bits 2048 -host localhost
```

This will generate cert.pem and key.pem.

I used to rely on openssl or mkcert, but this seems to work well when using Go.

cf. [Source file src/crypto/tls/generate_cert.go](https://go.dev/src/crypto/tls/generate_cert.go)