---
title: Introduction to grep
slug: introduction-to-grep
date: 2019-01-22T00:00:00Z
author: bmf-san
categories:
  - Operating Systems
tags:
  - Linux
  - grep
translation_key: introduction-to-grep
---

# Overview
grep is a command I used casually, so I decided to investigate it more thoroughly.

# Basics
`grep search-pattern filename`

You can use wildcards, so for example, if you want to target all files in the current directory, you can do it like this:

`grep "foo" ./*`

If you want to include directories under the current directory, use the `-r` option.

`grep -r "foo" ./*`

# Options
Here are some commonly used options.

## -i
- Ignore case distinctions

## -v
- Select non-matching lines

## -n
- Display line numbers with output lines

## -l
- Display file names with matches

## -L
- Display file names without matches

## -r
- Include directories in the search

## -w
- Search for whole words (exact match)

# Practical Use
## OR Search
`grep "foo\|bar" ./*`

You need to escape with `\`.

## AND Search
`grep "foo" ./* | grep "bar"`

## Exclude Specific Directories
`grep "foo" ./* --exclude-dir=vendor`

# References
- [@IT - 【 grep 】Command: Extract lines containing specific characters](http://www.atmarkit.co.jp/ait/articles/1604/07/news018.html)