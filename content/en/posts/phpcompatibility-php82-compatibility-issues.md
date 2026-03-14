---
title: Challenges Faced When Checking PHP 8.2 Compatibility with PHPCompatibility
description: 'Explore PHP 8.2 compatibility checks with PHPCompatibility CodeSniffer, addressing tool limitations and the develop branch.'
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


[PHPCompatibilitiy](https://github.com/PHPCompatibility/PHPCompatibility) is a tool used for checking compatibility with PHP 8.2, and here are some challenges I faced.

# Challenges
As of October 2023, even if you install the latest version [9.3.5](https://github.com/PHPCompatibility/PHPCompatibility/releases/tag/9.3.5), you cannot check compatibility with PHP 8.2.

Version 9.3.5 was released in 2019 and apparently does not yet support recent PHP versions...

# Solution
I briefly wondered if development had stopped, but that doesn't seem to be the case.

There are commits being added to the develop branch, so it seems using develop is the way to go.

cf. [Should I use develop or 9.3.5 sniffs? #1653](https://github.com/PHPCompatibility/PHPCompatibility/issues/1653)

# Thoughts
While it may not support 100% compatibility checks, it's a useful tool that I plan to continue using in the future.
