---
title: Basics of Regular Expressions
description: 'Master regular expression fundamentals using character classes, anchors, quantifiers, and grouping constructs for pattern matching.'
slug: basic-regular-expressions
date: 2018-12-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - ERE
  - Regular Expressions
translation_key: basic-regular-expressions
---

# Overview
Basics of regular expressions.
This summarizes commonly used notations in ERE (Extended Regular Expression).

# Character Classes
## [string]
- Matches any one character in the string.
- Ex:
  - [きつね]
    - きつねたぬきねこ
      - き,つ,ね,き,ね

## [^string]
- Matches any one character not in the string.
- Ex:
  - [^きつね]
    - きつねたぬきねこ
      - た,ね,こ

## [string-string]
- Matches any one character in the specified range.
- Ex:
  - [あ-ん]
    - きつねたぬきcat
      - き,つ,ね,た,ぬ,き

## \d
- Matches decimal digits.
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
- Matches all alphanumeric characters and underscores.
- Ex:
  - \w
    - abc_*
      - a,b,c,_

## \W
- Matches all characters that are not alphanumeric or underscores.
- Ex:
  - \W
    - abc_*
      - *

## \s
- Matches whitespace characters.
- Ex:
  - \s
    - a b c
      - Matches two whitespace characters (between a and b, between b and c).

## \S
- Matches any character that is not whitespace.
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
## (subexpression)
- Captures the string that matches the subexpression.
- Ex:
  - (りり){2}
    - ありりりりがとう ○
    - ありりがとう ✗

# Quantifiers
## *
- Matches 0 or more occurrences of the preceding element.
- Greedy quantifier.
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
- Matches 1 or more occurrences of the preceding element.
- Greedy quantifier.
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
- Matches 0 or 1 occurrence of the preceding element.
- Greedy quantifier.
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
- Matches 0 or more occurrences of the preceding element.
- Lazy quantifier.
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
- Matches 1 or more occurrences of the preceding element.
- Lazy quantifier.
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
- Matches 0 or 1 occurrence of the preceding element.
- Lazy quantifier.
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
- Matches exactly n occurrences of the preceding element.
- Ex:
  - b{2}
    - abba
      - bb

## {n,}
- Matches n or more occurrences of the preceding element.
- Ex:
  - b{2,}
    - abbba
      - bbb

## {n, m}
- Matches between n and m occurrences of the preceding element.
- Ex:
  - b{2,4}
    - abbba
      - bbb
    - abbbba
      - bbbb

# Alternation Constructs
## |
- Matches any one of the separated strings.
- Ex:
  - ab|cd
    - abcd
      - ab,cd
    - aaccd
      - cd

# References
- [Regular Expression Language - Quick Reference](https://docs.microsoft.com/ja-jp/dotnet/standard/base-types/regular-expression-language-quick-reference)
- [Regular Expression Notes](http://www.kt.rim.or.jp/~kbk/regex/regex.html#BRE)
- [Qiita - Regular Expressions Usable in Any UNIX Command](https://qiita.com/richmikan@github/items/b6fb641e5b2b9af3522e)
- [SE Academy - Regular Expressions](http://itref.fc2web.com/unix/regular-expression.html#BRE)