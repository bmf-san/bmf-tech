---
title: About SLI, SLO, and SLA
slug: sli-slo-sla-explained
date: 2022-09-10T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - SLI
  - SLA
  - SLO
translation_key: sli-slo-sla-explained
---

# About SLI, SLO, and SLA
This is a summary of various things I researched about SLI, SLO, and SLA.

# What are SLO, SLI, and SLA?
SLO, SLI, and SLA are indicators, objectives, and agreements related to service levels (Service Level). The service level is a representation of the service provided over a specific period measured in a particular way.

- **SLI (Service Level Indicator)**
  - Service Level Indicator
    - An indicator or metric used to measure service levels.
    - e.g. Availability, Latency, Error Rate, Throughput
- **SLO (Service Level Objective)**
  - Service Level Objective
    - A quantitative or qualitative value that serves as the goal for service levels.
    - Takes external dependencies into account.
      - Parts that communicate with external services, SLOs of managed services, etc.
- **SLA (Service Level Agreement)**
  - Service Level Agreement
    - An agreement or guarantee regarding service levels made between the service provider and the user.
    - It is better to set a looser target value than the SLO.

# How to Set SLI and SLO
I think the best practices proposed by New Relic are easy to implement and good.

[New Relic - Best Practices for Setting SLOs and SLIs for Modern Complex Systems](https://newrelic.com/jp/blog/best-practices/best-practices-for-setting-slos-and-slis-for-modern-complex-systems)

It introduces a method for formulating SLI and SLO, such as defining system boundaries, defining functions for each boundary, defining availability for each function, and defining SLI for measuring availability.

When starting to operate SLI and SLO, it is recommended to start as simply as possible with looser values.

cf. [SRE Google - Chapter 4 - Service Level Objectives](https://sre.google/sre-book/service-level-objectives/#indicators-o8seIAcZ)

When I actually formulated SLI and SLO in my work, I followed this New Relic practice but adjusted the functional units to avoid being too detailed.

If you make the functional units too detailed from the beginning, it will become difficult to operate, so I think it is better to adjust the granularity as needed while operating.

# Tips
Tips related to keywords associated with SLI and SLO.

## Difference Between Reliability and Availability
- **Reliability**
  - A property of the system, indicating the degree of resistance to failure.
- **Availability**
  - The degree to which the system can continue to operate.

## Uptime and Downtime Overview, Availability Calculation
|  Uptime  | Annual Downtime | Monthly Downtime |
| -------- | -------------- | ---------------- |
| 99.0%    | 87.6 hours     | 7.6 hours        |
| 99.5%    | 43.8 hours     | 3.65 hours       |
| 99.9%    | 8.76 hours     | 43.8 minutes     |
| 99.95%   | 4.38 hours     | 21.9 minutes     |
| 99.99%   | 52.56 seconds   | 4.38 minutes     |
| 99.999%  | 5.256 seconds   | 26.28 seconds    |
| 99.9999% | 31.536 seconds  | 2.628 seconds    |

## Error Budget
An error budget is a measure of acceptable reliability calculated based on the SLO.
For example, SLO 99.99% → Error Budget less than 0.01%.

# Thoughts
By making service levels measurable, it becomes possible to observe whether service users (users or systems) can adequately provide the service, and it can also serve as an indicator of whether service providers need to improve service levels.

# References
- [New Relic - What are SLOs, SLIs, and SLAs?](https://newrelic.com/jp/topics/what-are-slos-slis-slas#:~:text=%E3%82%B5%E3%83%BC%E3%83%93%E3%82%B9%E3%83%AC%E3%83%99%E3%83%AB%E7%9B%AE%E6%A8%99%EF%BC%88SLO%3AService,%E6%B8%AC%E5%AE%9A%E5%80%A4%E3%81%8A%E3%82%88%E3%81%B3%E3%83%A1%E3%83%88%E3%83%AA%E3%82%AF%E3%82%B9%E3%81%A7%E3%81%99%E3%80%82)
- [New Relic - New Relic Hands-On: Basics of SLI/SLO Design](https://newrelic.com/sites/default/files/2022-02/NRU303_SLISLO_20220222.pdf)
- [Cloud Google - Thinking About SLOs, SLIs, and SLAs: Lessons Learned by CRE](https://cloud.google.com/blog/ja/products/gcp/availability-part-deux-cre-life-lessons)
- [Cloud Google - SRE Fundamentals (2021 Edition): Comparison of SLI, SLA, and SLO](https://cloud.google.com/blog/ja/products/devops-sre/sre-fundamentals-sli-vs-slo-vs-sla)
- [Cloud Google - SLOs, SLIs, SLAs, oh my—CRE life lessons](https://cloud.google.com/blog/products/devops-sre/availability-part-deux-cre-life-lessons)
- [Cloud Google - How to Face Availability, That is the Question: Lessons Learned by CRE](https://cloud.google.com/blog/ja/products/gcp/available-or-not-that-is-the-question-cre-life-lessons)
- [Engineering Mercari - SLI/SLO in Mercari's Web Microservices (2018/12/25)](https://engineering.mercari.com/blog/entry/2018-12-25-150405/)
- [SRE Google - sre-book](https://sre.google/sre-book/service-level-objectives/#indicators-o8seIAcZ)
- [Qiita - Thoughts on Formulating SLI/SLO](https://qiita.com/t-okibayashi/items/9a5085803ac0b11554a0)
- [Qiita - Learning About SRE - Error Budget Edition](https://qiita.com/katsulang/items/feb3070666607b7c924c#:~:text=%E3%82%A8%E3%83%A9%E3%83%BC%E3%83%90%E3%82%B8%E3%82%A7%E3%83%83%E3%83%88%E3%81%A8%E3%81%AF%E3%80%81%E3%82%B5%E3%83%BC%E3%83%93%E3%82%B9,%E6%8A%91%E3%81%88%E3%82%8B%E3%81%93%E3%81%A8%E3%81%AB%E3%81%AA%E3%82%8A%E3%81%BE%E3%81%99%E3%80%82)
- [Bongineer - Difference Between Reliability and Availability](https://bongineer.net/entry/rasis/)
- [Math Words - Downtime for Availability of 99%, 99.9%, etc.](https://mathwords.net/kadouritu)
- [Wnkhs - Availability Calculation and Expected Values (Representative Numbers)](https://wnkhs.net/availability-calculation/)