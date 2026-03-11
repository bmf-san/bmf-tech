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
This post summarizes various findings about SLI, SLO, and SLA.

# What are SLO, SLI, and SLA?
SLO, SLI, and SLA are indicators, objectives, and agreements related to service levels. A service level is a measure of the service provided over a certain period, expressed in a specific way.

- SLI (Service Level Indicator)
  - Service Level Indicator
  - Metrics for measuring service levels
  - ex. Availability, latency, error rate, throughput
- SLO (Service Level Objective)
  - Service Level Objective
    - Quantitative or qualitative values set as service level goals
    - Consider external dependencies
      - Communication with external services, SLOs of managed services, etc.
- SLA (Service Level Agreement)
  - Service Level Agreement
    - Agreements or guarantees regarding service levels between providers and users
    - It is better to set looser target values than SLOs

# How to Set SLI and SLO
NewRelic's proposed best practices are easy to implement and effective.

[newrelic.com - Best Practices for Setting SLOs and SLIs for Modern Complex Systems](https://newrelic.com/jp/blog/best-practices/best-practices-for-setting-slos-and-slis-for-modern-complex-systems)

The method for formulating SLI and SLO is introduced, including defining system boundaries, defining functions for each boundary, defining availability for each function, and defining SLIs for measuring availability.

When starting to operate SLI and SLO, it is recommended to start with simple and loose values.

cf. [sre.google - Chapter 4 - Service Level Objectives](https://sre.google/sre-book/service-level-objectives/#indicators-o8seIAcZ)

When I actually formulated SLI and SLO in my work, I followed this NewRelic practice but adjusted the functional units to avoid becoming too detailed.

If you make the functional units too detailed from the start, it becomes difficult to operate, so I think it's better to adjust the granularity as needed during operation.

# Tips
Tips on keywords related to SLI and SLO.

## Difference Between Reliability and Availability
- Reliability
  - The degree of fault tolerance inherent in a system
- Availability
  - The degree to which a system can continue to operate

## List of Uptime and Downtime, Availability Calculation
|  Uptime  | Annual Downtime | Monthly Downtime |
| -------- | --------------- | ---------------- |
| 99.0%    | 87.6 hours      | 7.6 hours        |
| 99.5%    | 43.8 hours      | 3.65 hours       |
| 99.9%    | 8.76 hours      | 43.8 minutes     |
| 99.95%   | 4.38 hours      | 21.9 minutes     |
| 99.99%   | 52.56 seconds   | 4.38 minutes     |
| 99.999%  | 5.256 seconds   | 26.28 seconds    |
| 99.9999% | 31.536 seconds  | 2.628 seconds    |

## What is an Error Budget?
An error budget is a permissible reliability indicator calculated based on the SLO.
ex. SLO 99.99% → Error Budget less than 0.01%

# Impressions
By making service levels measurable, it becomes possible to observe whether the service users (users or systems) are satisfactorily provided with the service, and it can also serve as an indicator for service providers to determine whether service level improvements are necessary.

# References
- [newrelic.com - What are SLOs, SLIs, SLAs?](https://newrelic.com/jp/topics/what-are-slos-slis-slas#:~:text=%E3%82%B5%E3%83%BC%E3%83%93%E3%82%B9%E3%83%AC%E3%83%99%E3%83%AB%E7%9B%AE%E6%A8%99%EF%BC%88SLO%3AService,%E6%B8%AC%E5%AE%9A%E5%80%A4%E3%81%8A%E3%82%88%E3%81%B3%E3%83%A1%E3%83%88%E3%83%AA%E3%82%AF%E3%82%B9%E3%81%A7%E3%81%99%E3%80%82)
- [newrelic.com - New Relic Hands-on: Basics of SLI/SLO Design](https://newrelic.com/sites/default/files/2022-02/NRU303_SLISLO_20220222.pdf)
- [cloud.google.com - Thoughts on SLO, SLI, SLA: Lessons Learned by CRE](https://cloud.google.com/blog/ja/products/gcp/availability-part-deux-cre-life-lessons)
- [cloud.google.com - SRE Fundamentals (2021 Edition): Comparison of SLI, SLA, SLO](https://cloud.google.com/blog/ja/products/devops-sre/sre-fundamentals-sli-vs-slo-vs-sla)
- [cloud.google.com - SLOs, SLIs, SLAs, oh my—CRE life lessons](https://cloud.google.com/blog/products/devops-sre/availability-part-deux-cre-life-lessons)
- [cloud.google.com - How to Face Availability: Lessons Learned by CRE](https://cloud.google.com/blog/ja/products/gcp/available-or-not-that-is-the-question-cre-life-lessons)
- [engineering.mericari.com - 2018/12/25 SLI/SLO in Mercari's Web Microservices](https://engineering.mercari.com/blog/entry/2018-12-25-150405/)
- [sre.google - sre-book](https://sre.google/sre-book/service-level-objectives/#indicators-o8seIAcZ)
- [qiita.com - Thoughts on Formulating SLI/SLO](https://qiita.com/t-okibayashi/items/9a5085803ac0b11554a0)
- [qiita.com - Learning About SRE - Error Budget Edition](https://qiita.com/katsulang/items/feb3070666607b7c924c#:~:text=%E3%82%A8%E3%83%A9%E3%83%BC%E3%83%90%E3%82%B8%E3%82%A7%E3%83%83%E3%83%88%E3%81%A8%E3%81%AF%E3%80%81%E3%82%B5%E3%83%BC%E3%83%93%E3%82%B9,%E6%8A%91%E3%81%88%E3%82%8B%E3%81%93%E3%81%A8%E3%81%AB%E3%81%AA%E3%82%8A%E3%81%BE%E3%81%99%E3%80%82)
- [bongineer.net - Difference Between Reliability and Availability](https://bongineer.net/entry/rasis/)
- [mathwords.net - How Much Downtime for 99%, 99.9% Availability](https://mathwords.net/kadouritu)
- [wnkhs.net - Availability Calculation and Assumptions (with Representative Numbers)](https://wnkhs.net/availability-calculation/)
