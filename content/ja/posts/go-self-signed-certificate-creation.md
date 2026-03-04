---
title: "Goでオレオレ証明書がほしいときの一手"
slug: "go-self-signed-certificate-creation"
date: 2024-02-16
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Golang"
  - "Tips"
draft: false
---

GoでHTTPサーバーを書いているときなどオレオレ証明書がほしいときに役立つワンライナー。

```go
go run $(go env GOROOT)/src/crypto/tls/generate_cert.go -rsa-bits 2048 -host localhost
```

cert.pemとkey.pemが用意できる。

openssl使ったりmkcert使ったりしていたけどGo使っていたらこれで良さそう。

cf. [Source file src/crypto/tls/generate_cert.go](https://go.dev/src/crypto/tls/generate_cert.go)
