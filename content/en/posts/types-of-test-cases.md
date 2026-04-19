---
title: Types of Test Cases and Identification
slug: types-of-test-cases
date: 2018-04-11T00:00:00Z
author: bmf-san
categories:
  - Testing
tags:
  - Design
description: A concise summary of basic types of test cases and how to identify them.
translation_key: types-of-test-cases
---

# Overview
A concise summary of the basic types of test cases and how to identify them.

# Purpose of Testing
- Discover bugs
- Quality assurance
- Quality improvement (refactoring)

# Types of Tests

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
    - Check if the system can handle unexpected inputs
- Equivalence Partitioning
    - Divide into equivalence classes... Group test results (divide by criteria, e.g., language, presence of symbols, domain of address, etc.)
    - Select representative values from each equivalence class
- Boundary Value Analysis
    - Use values that are boundaries between equivalence classes as input

Identifying test cases for unit tests and integration tests should consider the above perspectives, along with business factors such as quality and effort.

# Reference
- ~~How to Identify and Write Test Cases~~