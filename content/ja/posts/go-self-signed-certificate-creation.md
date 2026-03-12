---
title: Goでオレオレ証明書がほしいときの一手
description: Goでオレオレ証明書がほしいときの一手について、基本的な概念から実践的な知見まで詳しく解説します。
slug: go-self-signed-certificate-creation
date: 2024-02-16T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
  - Tips
translation_key: go-self-signed-certificate-creation
---


GoでHTTPサーバーを書いているときなどオレオレ証明書がほしいときに役立つワンライナー。

```go
go run $(go env GOROOT)/src/crypto/tls/generate_cert.go -rsa-bits 2048 -host localhost
```

cert.pemとkey.pemが用意できる。

openssl使ったりmkcert使ったりしていたけどGo使っていたらこれで良さそう。

cf. [Source file src/crypto/tls/generate_cert.go](https://go.dev/src/crypto/tls/generate_cert.go)
