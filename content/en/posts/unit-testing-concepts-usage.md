---
title: Thinking About/Using Unit Tests
slug: unit-testing-concepts-usage
date: 2024-05-21T00:00:00Z
author: bmf-san
categories:
  - Testing
tags:
  - Design
  - Books
translation_key: unit-testing-concepts-usage
---

[Thinking About/Using Unit Tests](https://amzn.to/44Y35Xt) was read.

It summarizes the necessary thinking and approaches for high-quality test design. This book goes beyond just testing methods and is quite unique among test-related literature.

Good tests provide protection against regression, resilience to refactoring, quick feedback, and maintainability. However, since it is difficult to meet all these criteria, it is necessary to consider cost-effectiveness. This book contains many practical examples that illustrate this thinking.

Until now, I had been conscious of writing tests for production code, but I realized that I did not have a proper framework for deciding whether to write them or not. I felt that writing high-quality tests requires a deep understanding of design, which is likely still challenging for AI.

Chapter 10 on databases was impressive. Most of the concerns I had regarding testing databases were thoroughly covered.