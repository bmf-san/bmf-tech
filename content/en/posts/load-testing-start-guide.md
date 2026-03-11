---
title: Starting Guide to Load Testing
slug: load-testing-start-guide
date: 2024-08-24T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Load Testing
description: A guide to understanding and starting load testing.
translation_key: load-testing-start-guide
---



# Overview
This guide compiles content to help you understand and start load testing.

# What is Load Testing?
It is a **testing method to verify the performance of a system**.

This method aims to clarify whether the system can maintain its expected capacity (performance tolerance), what system impacts occur under excessive load, and where performance bottlenecks are.

※ In this article, load testing and performance testing, as well as load test and stress test, are treated as synonymous.

# Methods
Load testing can be broadly divided into four testing methods.

There are various definitions, but they generally fall into these four methods.

Choose the testing method according to the purpose of the load test (what you want to test).

## Load Testing
A method that **tests with a constant level of requests**.

The purpose is to **verify whether the expected performance is met**.

## Stress Testing
A method that **tests with requests assuming limits**.

The purpose is to **clarify the performance limit level and verify what kind of failures the system causes when the tolerance is exceeded**.

## Soak Testing
A method that **tests with a constant level of requests over a certain period (long duration)**.

The purpose is to **verify what impact occurs according to the system's resource status**.

Soak testing is synonymous with long-run testing.

## Baseline Testing
A method that **tests based on a single user's request**.

The purpose is to **identify performance benchmarks and analyze the degree of performance degradation by comparing with other test results**.

※ Although benchmark testing and baseline testing seem synonymous, there appears to be a difference.

cf. [loadview - Explanation of Benchmark Testing](https://www.loadview-testing.com/ja/blog/%E3%83%91%E3%83%95%E3%82%A9%E3%83%BC%E3%83%9E%E3%83%B3%E3%82%B9-%E3%83%86%E3%82%B9%E3%83%88-%E3%83%99%E3%83%BC%E3%82%B9%E3%83%A9%E3%82%A4%E3%83%B3-%E3%83%86%E3%82%B9%E3%83%88%E3%81%A8%E3%83%99/)

# Flow of Load Testing
The flow of load testing is explained by dividing it into planning and execution.

## Planning
### Setting Objectives
Clarify the **purpose** of conducting load testing, such as whether you want to test for no performance degradation, know the performance limits, or meet the required non-functional requirements.

Consider the **risks if the objectives are not met**. It's good to think about it in conjunction with risks.

ex.
- Want to confirm that there is no performance degradation due to modifications of existing APIs
- Want to verify if a new feature can withstand expected traffic
- Want to understand performance trends to set API throttling or quotas

### Investigation
Investigate the system configuration, functional specifications, access status, etc., to be tested according to the purpose.

- Infrastructure configuration
- Functional specifications
- System resource usage
- Access status

Organize the necessary information for planning sufficiently.

Lack of preparation can lead to ad-hoc responses during testing, affecting the credibility of test results.

It is also good to check the test execution environment.

### Setting Performance Goals and Metrics
Clarify performance goals as achievement criteria for the purpose.

ex.
- When there are requests to the top page with more than 1000 concurrent connections, the average response time should be below 800ms
- When continuously accepting the request volume during regular hours over a period assuming business days, the CPU usage should hover around an average of 20%

The more specific the performance goals, the easier it is to anticipate metrics.

### Creating Test Scenarios
Clarify what kind of requests will load the system.

Depending on the purpose and method, consider how the test subject will be used by users.

For example, in a load test to verify the response time of a My Page:

1. Access the login page
2. Press the login button
    - Authentication information is set in the cookie
    - User information retrieval API is called
2. Transition to My Page
    - Redirect after authentication processing
    - User settings information retrieval API is called

It is necessary to clarify what state the user is in when making requests about the test subject and write out the processes that need to be emulated.

## Execution
### 1. Preparation
Prepare the following for test execution.

- Preparation of the execution environment
    - Prepare the infrastructure environment for load testing and load testing tools
- Preparation of test data
- Sharing with stakeholders
  - Share the impact of load testing with stakeholders both inside and outside the company
    - → Estimate the impact range in advance
- Confirmation of monitoring targets
    - Organize the monitoring targets (logs, metrics, etc.) used during test execution
    - It is good to prepare a monitoring dashboard in advance

### 2. Test Execution
Execute the test.

### 3. Result Analysis
Observe where spikes occurred, if there are increasing or decreasing trends, or if unfamiliar errors are frequent, and consider causes and countermeasures.

### Tuning
Based on the analysis results, perform tuning if necessary and re-execute.

# Conclusion
In any case, the purpose is important. If the purpose wavers, the subsequent process becomes meaningless.

# References
- [Capacity Planning](https://www.oreilly.co.jp/books/9784873113999/)
    - An old book but a rare one written about capacity planning
    - The significance of capacity planning has not disappeared in the cloud-dominated modern era, so it seems worth having
- [Web API Testing Techniques](https://www.shoeisha.co.jp/book/detail/9784798179728)
    - There aren't many books that write about how to conduct load testing, but this book dedicates a chapter to it
