---
title: Types of Test Cases and How to Identify Them
slug: types-of-test-cases
date: 2018-04-11T00:00:00Z
author: bmf-san
categories:
  - Testing
tags:
  - Design
translation_key: types-of-test-cases
---

# Overview
A brief summary of the basic types of test cases and how to identify them.

# Purpose of Testing
- Discover bugs
- Quality assurance
- Quality improvement (refactoring)

# Types of Testing

- Unit Test
    - Tests to verify the specifications of methods
    - Targets public methods in the domain layer

- Integration Test
    - Tests based on use cases (a series of processes)
    - Prioritizes business-critical cases

# Identifying Test Cases
- Normal Cases
    - Check if the expected output is produced for the expected input
- Abnormal Cases
    - Determine if the system can handle unexpected input
- Equivalence Partitioning
    - Dividing into equivalence classes... grouping test results (based on criteria such as language, presence of symbols, domain of addresses, etc.)
    - Select representative values from each equivalence class
- Boundary Value Analysis
    - Use boundary values between equivalence classes as input

Identifying test cases for unit tests and integration tests should be done considering the above perspectives, along with business factors (quality and effort).

# Reference
- [How to Identify Test Cases and Examples of Writing Them](http://post.simplie.jp/posts/41)