---
title: Disadvantages of Jest Snapshot Testing
slug: jest-snapshot-testing-drawbacks
date: 2024-07-18T00:00:00Z
author: bmf-san
categories:
  - Testing
tags:
  - JavaScript
  - jest
description: A casual note on why you might want to avoid Jest snapshot testing.
translation_key: jest-snapshot-testing-drawbacks
---

A casual note on the disadvantages of Jest snapshot testing and why you might want to avoid it.

# Disadvantages

## Tests Become Fragile
Snapshot testing makes it easy to detect UI changes, but even changes that are not relevant to the test (e.g., renaming a class name without altering the appearance) can cause the test to fail.

## Unclear Assertion Expectations
Snapshot testing uses the `toMatchSnapshot()` assertion, but when a test fails, it can be difficult to understand the expected behavior. You have to judge based on the snapshot diff, but the assertion is often too concise to determine the correct state.

## Delayed Test Writing
Snapshot tests tend to fail repeatedly unless the implementation is complete, making it necessary to wait until the implementation is finished before writing the tests.

# References
There aren't many Japanese articles discussing the drawbacks of Jest, but the following article is well-organized and highly informative:

- [azukiazusa.dev - スナップショットテストとアサーションテスト](https://azukiazusa.dev/blog/snapshot-test-vs-assertion-test/)