---
title: Convert Markdown Files to PDF (with mermaid, emoji, and TOC support)
slug: markdown-to-pdf-conversion
date: 2022-09-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - markdown
  - marked
  - emoji
  - mermaid
  - JavaScript
translation_key: markdown-to-pdf-conversion
---

# Overview
I created a simple document management tool to meet the demand for converting Markdown files to PDF files.

[bmf-san/docs-md-to-pdf-example](https://github.com/bmf-san/docs-md-to-pdf-example)

Since I utilized existing libraries without much thought, the structure feels quite unsustainable.

# Motivation
If you simply want to convert Markdown files to PDF, you can just use the [md-to-pdf](https://github.com/simonhaenisch/md-to-pdf) library.

This library has also been helpful for managing my resume.
cf. [Managing my resume on GitHub](https://bmf-tech.com/posts/Github%e3%81%a7%e3%83%ac%e3%82%b8%e3%83%a5%e3%83%a1%e3%82%92%e7%ae%a1%e7%90%86%e3%81%99%e3%82%8b%e3%82%88%e3%81%86%e3%81%ab%e3%81%97%e3%81%9f)

I wanted to support mermaid syntax and use emojis that are not registered in Unicode, so I wanted to create something that accommodates those needs.

Using the [vscode-markdown-pdf](https://github.com/yzane/vscode-markdown-pdf) extension makes it easy to solve this, but it requires VSCode, which means some people would need to install it.

I thought it was nonsensical to use VSCode just for conversion, so I decided to implement it myself.

# Design
The [md-to-pdf](https://github.com/simonhaenisch/md-to-pdf) library is easy to use and a great library, but currently, the following features are not supported by default:

- mermaid syntax
- emojis (other than those registered in Unicode)
- TOC generation

Since [md-to-pdf](https://github.com/simonhaenisch/md-to-pdf) allows for configuration extensions of [markedjs/marked](https://github.com/markedjs/marked), it seems possible to achieve all of these by customizing [md-to-pdf](https://github.com/simonhaenisch/md-to-pdf).

It seems that TOC support is planned.
[Generate TOC (table of contents) #74](https://github.com/simonhaenisch/md-to-pdf/issues/74)

While it would have been fine to use [md-to-pdf](https://github.com/simonhaenisch/md-to-pdf), it seemed a bit cumbersome, so I wanted to implement it quickly in a hackathon-like manner, so I decided to use the [md-to-pdf-ng](https://github.com/mikewootc/md-to-pdf-ng) library.

This library extends [md-to-pdf](https://github.com/simonhaenisch/md-to-pdf) to support mermaid syntax, and while it doesn't seem to be maintained much, it works without issues.

Based on [md-to-pdf-ng](https://github.com/mikewootc/md-to-pdf-ng), I used [node-emojify](https://github.com/jesselpalmer/node-emojify) for emoji support and [doctoc](https://github.com/thlorenz/doctoc) for TOC generation.

# Implementation
Install the following via npm:

- [md-to-pdf-ng](https://github.com/mikewootc/md-to-pdf-ng)
- [node-emojify](https://github.com/jesselpalmer/node-emojify)
- [doctoc](https://github.com/thlorenz/doctoc)

*Note: I also included textlint as a bonus, but I will skip that part.*

Emoji support is implemented by extending marked, so I prepared the following configuration file:

```js
const marked = require('marked');
const { emojify } = require('node-emoji');

const renderer = new marked.Renderer();

renderer.text = emojify;

module.exports = {
	marked_options: { renderer },
};
```

Define the following command in the scripts section of package.json:

```json
doctoc --notitle md/ && md-to-pdf md/*.md --config-file config.js && mv md/*.pdf pdf/
```

First, generate the TOC with doctoc, then convert the Markdown to PDF, and finally move the directory.

It would be nice if md-to-pdf allowed specifying the output directory for generated PDFs, but it seems there is no such option, so I handled it with the straightforward method of `mv md/*.pdf pdf/`.

# Thoughts
When trying to create something like this, it tends to rely heavily on external libraries. Ideally, I would like to implement everything myself, but it seems quite challenging. If I get the chance, I would like to learn about the PDF data structure and try to create a similar CLI tool in Go.