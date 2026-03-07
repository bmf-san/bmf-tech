---
title: Basics of Regular Expressions
slug: basic-regular-expressions
date: 2018-12-01T00:00:00Z
author: bmf-san
categories:
  - Applications
tags:
  - ERE
  - Regular Expressions
description: An overview of the basics of regular expressions, focusing on commonly used syntax in ERE (Extended Regular Expressions).
translation_key: basic-regular-expressions
---

# Overview
An introduction to the basics of regular expressions. This post summarizes commonly used syntax in ERE (Extended Regular Expressions).

# Character Classes
## [Characters]
- Matches any single character within the brackets.
- Ex:
  - [きつね]
    - きつねたぬきねこ
      - き,つ,ね,き,ね

## [^Characters]
- Matches any single character not within the brackets.
- Ex:
  - [^きつね]
    - きつねたぬきねこ
      - た,ね,こ

## [Character-Character]
- Matches any single character within the specified range.
- Ex:
  - [あ-ん]
    - きつねたぬきcat
      - き,つ,ね,た,ぬ,き

## \d
- Matches any decimal digit.
- Ex:
  - \d
    - りんごが10個
      - 1, 0

## \D
- Matches any character that is not a decimal digit.
- Ex:
  - \D
    - りんごが10個
      - り,ん,ご,が,個

## \w
- Matches any alphanumeric character or underscore.
- Ex:
  - \w
    - abc_*
      - a,b,c,_

## \W
- Matches any character that is not alphanumeric or an underscore.
- Ex:
  - \W
    - abc_*
      - *

## \s
- Matches any whitespace character.
- Ex:
  - \s
    - a b c
      - Matches two spaces (between a and b, and b and c).

## \S
- Matches any non-whitespace character.
- Ex:
  - \S
    - a b c
      - a,b,c

# Anchors
## ^
- Matches the beginning of a line.
- Ex:
  - ^ありがとう
    - ありがとう友よ
    - 昨日はありがとう ✗

## $
- Matches the end of a line.
- Ex:
  - ありがとう$
    - ありがとう友よ ✗
    - 昨日はありがとう ○

# Grouping Constructs
## (Subexpression)
- Captures the substring that matches the subexpression.
- Ex:
  - (りり){2}
    - ありりりりがとう ○
    - ありりがとう ✗

# Quantifiers
## *
- Matches the preceding element zero or more times (greedy).
- Ex:
  - ab*
    - ab
      - ab
    - abab
      - ab, ab
    - aabb
      - ab
    - abbb
      - abbb
    - a
      - a
    - ba
      - a

## +
- Matches the preceding element one or more times (greedy).
- Ex:
  - ab+
    - ab
      - ab
    - abab
      - ab,ab
    - aabb
      - abb
    - abbb
      - abbb
    - a
      - No match
    - ba
      - No match

## ?
- Matches the preceding element zero or one time (greedy).
- Ex:
  - ab?
    - ab
      - ab
    - abab
      - ab,ab
    - aabb
      - a, ab
    - abbb
      - ab
    - a
      - a
    - ba
      - a

## *?
- Matches the preceding element zero or more times (lazy).
- Ex:
  - ab*?
    - ab
      - a
    - abab
      - a,a
    - aabb
      - a,a
    - abbb
      - a
    - a
      - a
    - ba
      - a

## +?
- Matches the preceding element one or more times (lazy).
- Ex:
  - ab+?
    - ab
      - ab
    - abab
      - ab,ab
    - aabb
      - ab
    - abbb
      - ab
    - a
      - No match
    - ba
      - No match

## ??
- Matches the preceding element zero or one time (lazy).
- Ex:
  - ab??
    - ab
      - a
    - abab
      - a,a
    - aabb
      - a, a
    - abbb
      - a
    - a
      - a
    - ba
      - a

## {n}
- Matches the preceding element exactly n times.
- Ex:
  - b{2}
    - abba
      - bb

## {n,}
- Matches the preceding element n or more times.
- Ex:
  - b{2,}
    - abbba
      - bbb

## {n, m}
- Matches the preceding element between n and m times.
- Ex:
  - b{2,4}
    - abbba
      - bbb
    - abbbba
      - bbbb

# Alternation Constructs
## |
- Matches any one of the separated alternatives.
- Ex:
  - ab|cd
    - abcd
      - ab,cd
    - aaccd
      - cd

# References
- [Regular Expression Language - Quick Reference](https://docs.microsoft.com/ja-jp/dotnet/standard/base-types/regular-expression-language-quick-reference)
- [Regular Expression Notes](http://www.kt.rim.or.jp/~kbk/regex/regex.html#BRE)
- [Qiita - Regular Expressions for Any UNIX Command](https://qiita.com/richmikan@github/items/b6fb641e5b2b9af3522e)
- [SE Academy - Regular Expressions](http://itref.fc2web.com/unix/regular-expression.html#BRE)