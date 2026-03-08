---
title: Using Realize for Hot Reload in Go
slug: realize-hot-reload-go
date: 2019-04-11T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Realize
translation_key: realize-hot-reload-go
---

# Overview
This is a note on using [github - oxequa/realize](https://github.com/oxequa/realize).

# Preparation
`go get github.com/oxequa/realize`

# Usage
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

In the demo directory, running `realize start --server` will start a monitoring server, enabling hot reload.

Monitoring server
`http://localhost:5002/#/demo`
You can view logs, errors, and output.

Server started by main.go
`http://localhost:8080`

# Thoughts
It's easy to set up and convenient, so I think I'll use this for hot reload.