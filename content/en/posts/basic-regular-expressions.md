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
- Matches any single character in the string.
- Ex:
  - [fox]
    - fox raccoon cat
      - f, o, x

## [^string]
- Matches any single character not in the string.
- Ex:
  - [^fox]
    - fox raccoon cat
      - r, a, c, o, n

## [string-string]
- Matches any single character in the specified range.
- Ex:
  - [a-z]
    - fox raccoon cat
      - f, o, x, r, a, c, o, n

## \d
- Matches decimal digits.
- Ex:
  - \d
    - There are 10 apples
      - 1, 0

## \D
- Matches any character that is not a decimal digit.
- Ex:
  - \D
    - There are 10 apples
      - T, h, e, r, e, a, p, l, e, s

## \w
- Matches all alphanumeric characters and underscores.
- Ex:
  - \w
    - abc_*
      - a, b, c, _

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
      - Matches two whitespace characters (between a and b, and between b and c)

## \S
- Matches any character that is not a whitespace.
- Ex:
  - \S
    - a b c
      - a, b, c

# Anchors
## ^
- Matches the beginning of a line.
- Ex:
  - ^Thank you
    - Thank you, my friend
    - Yesterday, thank you ✗

## $
- Matches the end of a line.
- Ex:
  - thank you$
    - Thank you, my friend ✗
    - Yesterday, thank you ○

# Grouping Constructs
## (subexpression)
- Captures the string that matches the subexpression.
- Ex:
  - (ri){2}
    - aririri thank you ○
    - ariri thank you ✗

# Quantifiers
## *
- Matches when the preceding element is repeated 0 or more times.
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
- Matches when the preceding element is repeated 1 or more times.
- Greedy quantifier.
- Ex:
  - ab+
    - ab
      - ab
    - abab
      - ab, ab
    - aabb
      - abb
    - abbb
      - abbb
    - a
      - No match
    - ba
      - No match

## ?
- Matches when the preceding element is repeated 0 or 1 time.
- Greedy quantifier.
- Ex:
  - ab?
    - ab
      - ab
    - abab
      - ab, ab
    - aabb
      - a, ab
    - abbb
      - ab
    - a
      - a
    - ba
      - a

## *?
- Matches when the preceding element is repeated 0 or more times.
- Lazy quantifier.
- Ex:
  - ab*?
    - ab
      - a
    - abab
      - a, a
    - aabb
      - a, a
    - abbb
      - a
    - a
      - a
    - ba
      - a

## +?
- Matches when the preceding element is repeated 1 or more times.
- Lazy quantifier.
- Ex:
  - ab+?
    - ab
      - ab
    - abab
      - ab, ab
    - aabb
      - ab
    - abbb
      - ab
    - a
      - No match
    - ba
      - No match

## ??
- Matches when the preceding element is repeated 0 or 1 time.
- Lazy quantifier.
- Ex:
  - ab??
    - ab
      - a
    - abab
      - a, a
    - aabb
      - a, a
    - abbb
      - a
    - a
      - a
    - ba
      - a

## {n}
- Matches when the preceding element is repeated n times.
- Ex:
  - b{2}
    - abba
      - bb

## {n,}
- Matches when the preceding element is repeated n or more times.
- Ex:
  - b{2,}
    - abbba
      - bbb

## {n, m}
- Matches when the preceding element is repeated between n and m times.
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
      - ab, cd
    - aaccd
      - cd

# References
- [Regular Expression Language - Quick Reference](https://docs.microsoft.com/en-us/dotnet/standard/base-types/regular-expression-language-quick-reference)
- [Regular Expression Notes](http://www.kt.rim.or.jp/~kbk/regex/regex.html#BRE)
- [Qiita - Regular Expressions Usable in Any UNIX Command](https://qiita.com/richmikan@github/items/b6fb641e5b2b9af3522e)
- [SE Academy - Regular Expressions](http://itref.fc2web.com/unix/regular-expression.html#BRE)
