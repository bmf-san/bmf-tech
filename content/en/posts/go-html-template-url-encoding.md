---
title: Prevent URL Encoding in Go's html/template
slug: go-html-template-url-encoding
date: 2023-04-23T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Tips
translation_key: go-html-template-url-encoding
---

# Overview
When using html/template, I wanted to prevent the URL passed to the template from being encoded.

# Using template.URL
When you pass a URL to a template using Go's [html/template](https://pkg.go.dev/html/template), it is encoded by default.

cf. https://pkg.go.dev/html/template#hdr-Contexts

I believe this behavior is for security reasons, but there are cases where you might want to avoid this in HTML.

In such cases, you can use `template.URL` to bypass this.

```golang
package main

import (
	"html/template"
	"os"
)

func main() {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>bmf-tech.com</title>
	</head>
	<body>
		<a href="{{ .URL }}">bmf-tech.com</a>
	</body>
</html>`

	t, _ := template.New("index").Parse(tpl)
	data := struct {
		URL template.URL
	}{
		URL: template.URL("http://bmf-tech/posts/search?keyword=something"),
	}

	t.Execute(os.Stdout, data)
}
```

# Thoughts
I got stuck on this unexpectedly.