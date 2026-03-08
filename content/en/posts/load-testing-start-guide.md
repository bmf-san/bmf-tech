---
title: Starting Guide to Load Testing
slug: load-testing-start-guide
date: 2024-08-24T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Load Testing
translation_key: load-testing-start-guide
---

# Overview
This post summarizes the content to help you understand load testing and get started with it.

# What is Load Testing?
Load testing is a **testing method to verify system performance**.

It aims to clarify whether the system can maintain its expected capacity (performance tolerance), what system impacts occur under excessive load, and where performance-related bottlenecks are.

*Note: In this article, load testing and performance testing are treated as synonymous, while load test and stress test are also considered equivalent.*

# Methods
There are four main testing methods for load testing.

While there are various definitions, these four methods generally apply.

Choose the testing method based on the purpose of load testing (what you want to test).

## Load Testing
This method tests with a **constant level of requests**.

The goal is to **verify whether the expected performance is met**.

## Stress Testing
This method tests with requests that assume **limits**.

The goal is to **clarify the performance limits and verify what kind of failures occur when exceeding tolerance levels**.

## Soak Testing
This method tests with a **constant level of requests over a certain period (long duration)**.

The goal is to **verify what impacts occur based on the system's resource status**.

Soak testing is synonymous with long-run testing.

## Baseline Testing
This method tests based on **single user requests**.

The goal is to **identify performance benchmarks or analyze the degree of performance degradation compared to other test results**.

*Note: While benchmark testing and baseline testing seem synonymous, there are differences.*

cf. [loadview - Explanation of Benchmark Testing](https://www.loadview-testing.com/ja/blog/%E3%83%91%E3%83%95%E3%82%A9%E3%83%BC%E3%83%9E%E3%83%B3%E3%82%B9-%E3%83%86%E3%82%B9%E3%83%88-%E3%83%99%E3%83%BC%E3%82%B9%E3%83%A9%E3%82%A4%E3%83%B3-%E3%83%86%E3%82%B9%E3%83%88%E3%81%A8%E3%83%99/) 

# Flow of Load Testing
The flow of load testing is explained in terms of planning and execution.

## Planning
### Setting Objectives
Clarify the **purpose of conducting load testing**, such as whether you want to test for performance degradation, know the limit performance, or meet required non-functional requirements.

Consider **what risks exist if the objectives are not met**.

Examples:
- Confirm that there is no performance degradation due to modifications to existing APIs.
- Validate whether new features can withstand expected traffic.
- Understand performance trends to define API throttling and quotas.

### Investigation
Investigate the configuration and specifications of the system to be tested, as well as access conditions based on the objectives.

- Infrastructure configuration
- Functional specifications
- System resource usage
- Access conditions

Organize sufficient information necessary for planning.

Insufficient preparation may lead to ad-hoc responses during testing, affecting the credibility of the test results.

Also, check the testing environment.

### Setting Performance Goals and Metrics
Clearly define performance goals as criteria for achieving objectives.

Examples:
- When there are requests on the top screen with over 1000 simultaneous connections, the response time should be below an average of 800ms.
- When continuously receiving request volumes during business days, the CPU usage should hover around an average of 20%.

The more specific the performance goals, the easier it is to anticipate measurement metrics.

### Creating Test Scenarios
Clarify how to apply load to the system with specific requests.

Depending on the purpose and method, consider how the test subject will be used by users.

For example, in a load test to verify the response time of a My Page:

1. Access the login page.
2. Press the login button.
    - Authentication information is set in cookies.
    - User information retrieval API is called.
3. Transition to My Page.
    - After authentication processing, redirect occurs.
    - User setting information retrieval API is called.

You need to clearly outline the processes that need to be emulated based on the state users will have when making requests.

## Execution
### 1. Preparation
Prepare the following for testing execution:

- Prepare the execution environment.
    - Set up the infrastructure environment for load testing and load testing tools.
- Prepare test data.
- Share with stakeholders.
    - Communicate the impact of load testing to relevant parties both inside and outside the organization.
    - → Estimate the impact range in advance.
- Confirm monitoring targets.
    - Organize the monitoring targets (logs, metrics, etc.) to be used during testing.
    - It’s good to prepare a monitoring dashboard in advance.

### 2. Conduct Testing
Execute the tests.

### 3. Analyze Results
Observe where spikes occurred, whether there are increasing or decreasing trends, or if unfamiliar errors are frequently occurring, and consider causes and countermeasures.

### Tuning
Based on the analysis results, perform tuning if necessary and re-execute the tests.

# Conclusion
Above all, the purpose is crucial. If the purpose becomes unclear, the subsequent processes lose their meaning.

# References
- [Capacity Planning](https://www.oreilly.co.jp/books/9784873113999/)
    - An old book, but a rare one that discusses capacity planning.
    - The significance of capacity planning has not diminished in the cloud-dominated modern era, so it may be worth keeping.
- [Web API Testing Techniques](https://www.shoeisha.co.jp/book/detail/9784798179728)
    - There aren’t many books that discuss methods for conducting load testing, but this book dedicates a chapter to it.