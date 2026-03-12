---
title: Measuring Scalability
description: An in-depth look at Measuring Scalability, covering key concepts and practical insights.
slug: measuring-scalability
date: 2025-06-08T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Scalability
translation_key: measuring-scalability
---

# Overview
I researched methods for measuring scalability and summarized them.

# Basic Understanding of Scalability
- **Scalability** is a quality attribute related to how a system can handle increased workloads.
- Scalability directly affects **Efficiency**, not just by adding resources.
- The key is how the system can **maintain and expand workload (processing capacity and responsiveness)** when resources like **CPU, memory, storage, and network** are increased or decreased.
- Scalability is closely tied to performance, cost, and maintainability, making it a core concern in architectural design.

# Key Metrics
- **Throughput**
  - Number of requests processed per unit time (e.g., RPS, TPS)
  - Example: 10,000 RPS with 100 instances ⇒ 100 RPS per instance
- **Latency**
  - Response times such as p50, p95, p99
  - How much low latency can be maintained while increasing throughput
- **Efficiency**
  - Performance improvement rate per added resource
  - **Speedup**
    $S(n) = \frac{T(1)}{T(n)}$
  - **Efficiency (E)**
    $E(n) = \frac{S(n)}{n}$
- **Cost Performance**
  - Cost per request (\$/RPS)
  - Or throughput cost (RPS/\$)
- **Elasticity**
  - Speed and stability of scale-out/in in response to load fluctuations
  - Auto-scaling startup time and frequency of oscillations

# Quantitative Evaluation Procedure
- **Baseline Measurement**
  - Obtain throughput and latency with minimum configuration (e.g., 1 instance)
- **Resource Addition Experiment**
  - Gradually increase the number of instances (1→2→4→8...) and benchmark at each step
- **Calculate Speedup and Efficiency**
  - Compute $S(n), E(n)$ for each $n$
- **Bottleneck Analysis**
  - Observe resource utilization (CPU, memory, DB connections, etc.)
- **Scenario-based Evaluation**
  - Conduct tests with read/write/mixed loads
- **Cost Estimation**
  - Compare actual operational costs and performance to determine optimal operation points

# Example of Quantification
| Number of Instances $n$ | Throughput RPS $R(n)$ | p95 Latency (ms) | Speedup $S(n)$ | Efficiency $E(n)$ |
| ------------------ | ----------------------- | ------------------- | ----------------- | ----------- |
| 1                  | 500                     | 120                 | 1.0               | 1.00        |
| 2                  | 1000                    | 125                 | 2.0               | 1.00        |
| 4                  | 1900                    | 140                 | 3.8               | 0.95        |
| 8                  | 3500                    | 180                 | 7.0               | 0.88        |

# Utilization of Mathematical Models
Mathematical models help clarify the theoretical limits and expectations of scalability.
- **Understand the limits of ideal scaling**
  - Amdahl's Law allows quantitative prediction of maximum speedup when there are non-parallelizable processes.
- **Identify gaps with reality**
  - Comparing measured values with theoretical values can identify causes of efficiency decline (bottlenecks).
- **Discuss cost-effectiveness of resource addition**
  - Gustafson's Law allows design evaluation based on efficiency improvement with problem size expansion.
- **Explain the limits of unlimited scaling**
  - Mathematical formulas can substantiate the point where increasing resources no longer yields benefits.

Below are representative models.

- **Amdahl's Law**

  $$
  S_{max} = \frac{1}{\alpha + \frac{1 - \alpha}{n}}
  $$

  (α is the non-parallelizable portion)

- **Gustafson's Law**
  A model where efficiency improves as problem size increases

# Operational and Maintainability Aspects (People & Maintainability)
- **Need for Maintenance Personnel**
  - Operational workload, skill set, presence of automation
- **Failure Response Flow**
  - Alert definitions, escalation routes, SLO compliance
- **CI/CD and Release Operational Load**
  - Deployment automation, feasibility of safe release strategies
- **Log and Monitoring Setup**
  - Level of observability, ease of monthly reviews

# Storage and Cost Efficiency
- **Storage Configuration and Scalability**
  - Optimization of cache, object/block storage
- **Storage Cost Optimization**
  - Application of tiered storage, lifecycle rules
- **Cost Observability and Optimization**
  - Monitoring cost per unit (\$/RPS), idle cost ratio
- **Cost-effectiveness Analysis (ROI)**
  - Visualize additional cost per additional performance

# Supplementary Evaluation Metrics
| Metric                          | Meaning                 | Supplement                        |
| ----------------------------- | -------------------- | --------------------------- |
| Maintenance Effort (man-days/month)           | Labor required for operation     | Includes alerts and configuration changes      |
| Cost Efficiency (\$ / RPS)         | Processing unit cost             | Compare under high load and idle conditions  |
| Idle Resource Rate            | Ratio of unused resources | Caution when setting min\_instances |
| Storage Unit Cost (\$ / GB / month) | Storage cost           | Use of compression and retention periods        |

# Conclusion
- Obtain throughput and latency through benchmarking
- Quantify performance improvement and efficiency with resource addition
- Identify bottlenecks from scale curves
- Add maintainability and cost efficiency as evaluation metrics to develop realistic expansion strategies

# References
- [Software Architecture Metrics](https://amzn.to/3ZWhVMP)
- [www.issoh.co.jp - Basic Points and Guides for Building Scalable Systems](https://www.issoh.co.jp/column/details/4129/?utm_source=chatgpt.com#i-18)
- [www.frugaltesting.com - How to Test the Scalability of Software Systems: Best Practices](https://www.frugaltesting.com/blog/how-to-test-the-scalability-of-software-systems-best-practices?utm_source=chatgpt.com)