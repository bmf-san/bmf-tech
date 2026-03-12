---
title: Disadvantages of Jest Snapshot Testing
description: An in-depth look at Disadvantages of Jest Snapshot Testing, covering key concepts and practical insights.
slug: jest-snapshot-testing-drawbacks
date: 2024-07-18T00:00:00Z
author: bmf-san
categories:
  - Testing
tags:
  - JavaScript
  - Jest
translation_key: jest-snapshot-testing-drawbacks
---

Here are some rough notes on the disadvantages of avoiding Jest snapshot testing.

# Disadvantages
## Tests Become Fragile
While snapshot tests can easily detect changes in the UI, they can also fail due to changes that are not the main concern of the test (e.g., renaming a class that does not change the appearance).

## Unclear Assertion Expectations
In snapshot tests, the assertion `toMatchSnapshot()` is used, but when a test fails, it can be difficult to read the expected specifications. You have to look at the differences in the snapshot to make a judgment, but the assertion is too concise, making it hard to determine what the correct state is.

## Delayed Test Writing
Snapshot tests can only be written after the implementation is complete, as they will fail continuously otherwise, making the completion of the implementation a prerequisite.

# References
The disadvantages of Jest are not often discussed in Japanese articles, but the following article is well summarized and very helpful.

- [azukiazusa.dev - Snapshot Testing vs Assertion Testing](https://azukiazusa.dev/blog/snapshot-test-vs-assertion-test/)