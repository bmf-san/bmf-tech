---
title: About Capacity Planning
slug: capacity-planning
date: 2024-03-30T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Capacity Planning
description: An overview of capacity planning for efficient resource management in systems.
translation_key: capacity-planning
---

# Purpose of Capacity Planning
The goal is to manage system resources and utilize them efficiently.

# Capacity Planning Process
Capacity planning involves balancing between rough estimates and detailed analysis. Avoid unnecessary deep dives and prepare detailed information only when needed.

1. Define application performance requirements
   - Response time, throughput, concurrent connections, load performance, etc.

2. Understand the current infrastructure's operational status
   - Measure the load characteristics of each architectural component (e.g., web servers, databases, storage) that make up the application and compare them with performance requirements.

3. Consider future needs to maintain acceptable performance
   - Based on past data, evaluate what resources will be needed, in what quantity, and when, while considering budget and schedule.

4. Sizing
   - Adjust the architecture (modifications, additions, removals) according to the capacity plan.

# Measurement
System metrics should be evaluated in relation to the application's specific characteristics (e.g., how users utilize features, peak usage times, etc.).

By closely observing application performance trends, it becomes easier to determine which capacity areas should be prioritized.

Additionally, linking metrics to business requirements and indicators enhances the persuasiveness of technical investments.

# Performance Tuning vs. Capacity Planning
Performance tuning aims to optimize the performance of existing systems.

Capacity planning, on the other hand, predicts what resources and when they will be needed based on the performance of existing systems, without focusing on optimization.

In capacity planning, observations based on actual usage are prioritized over benchmarks like load testing.

# Goal Setting in Capacity Planning
The goals to be set in capacity planning include the following:

- **Performance**
  - Performance requirements based on monitoring
    - Predict future performance needs
  - SLA
    - For web applications, primarily targets availability and performance metrics
  - Business requirements
  - User-expected performance

- **Capacity**
  - System metrics
  - Resource limits

# Safety Margin
The safety margin refers to the buffer percentage added as a risk hedge against uncertainties in the predicted capacity during capacity planning.

Determine a reasonable percentage based on past system performance trends and the nature of the application.

# Requirements for Measurement Tools
- Ability to store data over a certain period
- Customizable metrics
- Capability to collect and compare metrics from arbitrary sources
- Ability to import and export metrics

The "observer effect" caused by measurement should be accepted as a necessary cost in capacity planning.

# References
- [www.oreilly.co.jp - Capacity Planning](https://www.oreilly.co.jp/books/9784873113999/)