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
This is a note on how to perform static analysis of Clean Architecture layers in Go.

# Using go-cleanarch
I could have created my own static analysis tool, but I found an easy-to-use tool, so I decided to try it out.

[roblaszczak/go-cleanarch](https://github.com/roblaszczak/go-cleanarch) 

I implemented it in my custom CMS, [gobel-api](https://github.com/bmf-san/gobel-api).
cf. [PR](https://github.com/bmf-san/gobel-api/pull/74/files)

Install it with `go install github.com/roblaszczak/go-cleanarch@latest`.

Since the naming of the layers differs from the default, I ran the check with options.
`go-cleanarch -application usecase`

If there is a violation of the layer dependencies, it will trigger a check and result in an error.

When an error occurs, I get told that `Uncle Bob is not happy.`

# Thoughts
I would like to create my own tool eventually, but for now, I want to rely on this tool.

I believe that static analysis tools like this should be introduced early on to maintain the design and structure of applications.

Since it seems easy to implement static analysis tools in Go, if you want to maintain the structure of specific layers, it would be good to actively create such tools.