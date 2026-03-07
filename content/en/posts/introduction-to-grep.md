---
title: Introduction to grep
slug: introduction-to-grep
date: 2019-01-22T00:00:00Z
author: bmf-san
categories:
  - OS
tags:
  - Linux
  - grep
translation_key: introduction-to-grep
---

# Overview
I have been using grep without fully understanding it, so I did a rough investigation.

# Basics
`grep search_regex filename`

You can use wildcards, so for example, if you want to target all files in the current directory:

`grep "foo" ./*`

If you want to include directories under the current directory, use the `-r` option.

`grep -r "foo" ./*`

# Options
I picked out only the ones that seem useful.

## -i
- Case insensitive

## -v
- Target non-matching items

## -n
- Display line numbers in search results

## -l
- Display filenames in search results

## -L
- Display files that did not match in search results

## -r
- Include directories in the search target

## -w
- Search by whole word (search for exact matches)

# Practical Use
## OR Search
`grep "foo\|bar" ./*`

You need to escape with `\`.

## AND Search
`grep "foo" ./* | grep "bar"`

## Exclude Specific Directory
`grep "foo" ./* --exclude-dir=vendor`

# References
- [@IT - grep command: Extract lines containing specific characters](http://www.atmarkit.co.jp/ait/articles/1604/07/news018.html)