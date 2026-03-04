---
title: "Realizeを使ってgoでホットリロードを実現するメモ"
slug: "realize-go"
date: 2019-04-11
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Golang"
  - "realize"
draft: false
---

# 概要
[github - oxequa/realize](https://github.com/oxequa/realize)を使ってみたメモ。

# 準備
`go get github.com/oxequa/realize`

# 使い方
```
./demo/
├── .realize.yaml
└── main.go
```

.realize.yaml
```yaml
settings:
  legacy:
    force: false
    interval: 0s
schema:
- name: demo
  path: .
  commands:
    run:
      status: true
  watcher:
    extensions:
    - go
    paths:
    - /
    ignored_paths:
    - .git
    - .realize
    - vendor
```

main.go
```go
package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello World")
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", handler)

  http.ListenAndServe(":8080", mux)
}
```

demoディレクトリにて、`realize start --server`をすると、監視用サーバーが立ち上がり、ホットリロードができるようになる。

監視サーバー
`http://localhost:5002/#/demo`
log、error、outputが見れる。

main.goで起動されるサーバー
`http://localhost:8080`

# 所感
簡単に導入できて便利なのでほっとリロードはこれ使っていこうかーという気持ち。


