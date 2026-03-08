---
title: Issues Encountered When Checking PHP 8.2 Compatibility Using PHPCompatibility
slug: phpcompatibility-php82-compatibility-issues
date: 2024-04-14T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - PHP CodeSniffer
  - PHPCompatibility
translation_key: phpcompatibility-php82-compatibility-issues
---

[PHPCompatibility](https://github.com/PHPCompatibility/PHPCompatibility) notes on issues encountered while checking PHP 8.2 compatibility.

# Issues Encountered
As of October 2023, even after installing the latest version [9.3.5](https://github.com/PHPCompatibility/PHPCompatibility/releases/tag/9.3.5), compatibility checks for PHP 8.2 cannot be performed.

Version 9.3.5 was released in 2019 and seems to not yet support recent PHP versions...

# Solution
I briefly wondered if development had stopped, but that doesn't seem to be the case.

Commits are being made to the develop branch, so it appears that using develop is the way to go.

cf. [Should I use develop or 9.3.5 sniffs? #1653](https://github.com/PHPCompatibility/PHPCompatibility/issues/1653)

# Thoughts
I don't think it supports 100% compatibility checks, but it's a useful tool that I want to continue using in the future.