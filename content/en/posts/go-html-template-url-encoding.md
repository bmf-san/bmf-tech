---
title: Prevent URL Encoding in Go's html/template
description: 'An in-depth look at Prevent URL Encoding in Go''s html/template, covering key concepts and practical insights.'
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
When you use Go's [html/template](https://pkg.go.dev/html/template) to pass a URL to a template, it is designed to be encoded.

cf. https://pkg.go.dev/html/template#hdr-Contexts

I believe this is due to security reasons, but there might be cases where you want to avoid this in HTML.

In such cases, you can use `template.URL` to avoid it.

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
