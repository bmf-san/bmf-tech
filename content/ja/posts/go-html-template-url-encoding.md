---
title: Goのhtml／templateでURLをエンコードさせない
description: Goのhtml／templateでURLをエンコードさせない
slug: go-html-template-url-encoding
date: 2023-04-23T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
  - Tips
translation_key: go-html-template-url-encoding
---


# 概要
html/templateを使っているときに、テンプレートに渡すURLをエンコードさせたくなかった。

# template.URLを使う
Goの[html/template](https://pkg.go.dev/html/template)を使って、URLをテンプレートに渡すとエンコードされてしまう仕様になっている。

cf. https://pkg.go.dev/html/template#hdr-Contexts

セキュリティ上の理由でこのような仕様になっていると思うが、HTML上でこれを回避したいようなケースがあると思う。

そういうときは`template.URL`を使うと回避できる。

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

# 所感
地味にハマった。

