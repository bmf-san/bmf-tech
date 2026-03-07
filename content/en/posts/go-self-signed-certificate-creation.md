---
title: A Handy Way to Generate Self-Signed Certificates in Go
slug: go-self-signed-certificate-creation
date: 2024-02-16T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Tips
description: A quick one-liner for generating self-signed certificates when working with Go HTTP servers.
translation_key: go-self-signed-certificate-creation
---

When writing an HTTP server in Go, you might need a self-signed certificate. Here's a useful one-liner for that:

```go
go run $(go env GOROOT)/src/crypto/tls/generate_cert.go -rsa-bits 2048 -host localhost
```

This will generate `cert.pem` and `key.pem` files.

While I used to rely on tools like OpenSSL or mkcert, this approach seems sufficient if you're already using Go.

cf. [Source file src/crypto/tls/generate_cert.go](https://go.dev/src/crypto/tls/generate_cert.go)