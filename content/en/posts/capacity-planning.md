---
title: About Capacity Planning
description: An in-depth look at About Capacity Planning, covering key concepts and practical insights.
slug: capacity-planning
date: 2024-03-30T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Capacity Planning
translation_key: capacity-planning
---

# Purpose of Capacity Planning
The goal is to manage system resources and aim for efficient utilization of those resources.

# Process of Capacity Planning
Capacity planning uses both rough estimates and detailed analysis. It avoids unnecessary detail while preparing appropriate specifics when needed.

1. Define application performance requirements

   Response time, throughput, concurrent connections, load performance, etc.

2. Understand the current infrastructure's operational status

   Measure the load characteristics of each architectural element that makes up the application (Web servers, DB, storage, etc.) and compare them with performance requirements.

3. Consider future needs to maintain acceptable performance

   Based on past performance, determine what will be needed, when, and how much, in relation to budget and schedule.

4. Sizing

   Adjust the architecture according to the capacity plan, including changes, additions, or deletions of architecture.

# Measurement
System metrics should be compared with application-specific characteristics (how users utilize features, when they are most used, etc.).

This allows for deeper observation of application performance trends, making it easier to prioritize which capacities to consider.

Additionally, it becomes easier to relate to business requirements and business metrics, increasing the persuasiveness of technical investments.

# Performance Tuning and Capacity Planning
Performance tuning aims to optimize the performance of existing systems.

Capacity planning, on the other hand, predicts what will be needed for the system and when, based on the performance of existing systems, without considering what to optimize.

In capacity planning, observations based on actual usage take precedence over observations from benchmarks like load testing.

# Goal Setting in Capacity Planning
Goals to be set in capacity planning include the following items:

- Performance
  - Performance requirements based on monitoring
    - Predict future performance needs
  - SLA
    - Target values mainly focused on availability and performance for web applications
  - Business requirements
  - Performance expected by users
- Capacity
  - System metrics
  - Resource limits

# Safety Factor
The safety factor is the percentage of buffer added as a risk hedge against the uncertainty of the capacity predicted in capacity planning.

Determine a reasonable percentage based on past system performance trends and application characteristics.

# Requirements for Measurement Tools
- Ability to store data for a certain period
- Customizable metrics
- Ability to obtain and compare metrics from any source
- Ability to import and export metrics

The "observer effect" from measurement is best accepted as a necessary cost in capacity planning.

# References
- [www.oreilly.co.jp - Capacity Planning](https://www.oreilly.co.jp/books/9784873113999/)