---
title: Static Analysis of Clean Architecture Layers in Go
slug: go-clean-architecture-static-analysis
date: 2022-09-04T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Clean Architecture
translation_key: go-clean-architecture-static-analysis
---



# Overview
A note on how to perform static analysis of Clean Architecture layers in Go.

# Using go-cleanarch
While I could have created a custom static analysis tool, I found an easy-to-use tool and decided to try it out.

[roblaszczak/go-cleanarch](https://github.com/roblaszczak/go-cleanarch)

I introduced it to my custom CMS, [gobel-api](https://github.com/bmf-san/gobel-api).
cf. [PR](https://github.com/bmf-san/gobel-api/pull/74/files)

Install it with `go install github.com/roblaszczak/go-cleanarch@latest`.

Since the layer naming differs from the default, run the check with options.
`go-cleanarch -application usecase`

If there is a violation of layer dependencies, the check will catch it and result in an error.

If an error occurs, you'll be told, `Uncle Bob is not happy.`

# Thoughts
I plan to create my own tool eventually, but for now, I want to rely on this tool.

I believe that static analysis tools like this should be introduced early to maintain the design and structure of applications.

Since Go makes it easy to implement static analysis tools, it seems beneficial to actively create such tools if you want to maintain the structure of specific layers, not limited to Clean Architecture.
