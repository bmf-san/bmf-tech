---
title: 'Legacy Code Improvement Guide: Refactoring for Maintenance Development'
description: "A review of Working Effectively with Legacy Code. A book that systematizes practical techniques for tackling legacy code that lacks tests. Centered on concepts like seams, places to insert tests without changing behavior, it teaches how to make changes safely, step by step."
slug: legacy-code-improvement-guide
date: 2017-03-12T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Legacy Code
  - Book Review
translation_key: legacy-code-improvement-guide
books:
  - asin: "4798116831"
    title: "レガシーコード改善ガイド: 保守開発のためのリファクタリング"
draft: false
---

[Legacy Code Improvement Guide: Refactoring for Maintenance Development](https://amzn.to/4adL0FR) is a book I read.

It's a must-read for anyone facing legacy code.

As a prerequisite, I feel that a certain level of ability to write tests is required.

  - A seam is a place where you can change the behavior of the program even if you cannot directly edit that location.
  - Every seam has an enabling point. At the enabling point, you can decide which behavior to use.