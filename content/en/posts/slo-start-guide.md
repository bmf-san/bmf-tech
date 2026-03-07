---
title: 'Getting Started with SLO: A Guide'
slug: slo-start-guide
date: 2024-07-23T00:00:00Z
author: bmf-san
categories:
  - Operations
tags:
  - Reliability
  - SLO
translation_key: slo-start-guide
---

# Overview
This post summarizes a guide to understanding SLOs and starting their operation.

# Steps from Introduction to Operation of SLOs
To introduce SLOs in an organization or team and start their operation, several stages must be followed.

1. Share knowledge about SLOs
2. Agree on the purpose of introducing SLOs
3. Agree on the operational policy for SLOs
4. Design the SLOs
5. Start the operation of SLOs

When introducing SLOs for the first time in an organization, it is crucial to first share knowledge about SLOs and reach an agreement on the purpose of introduction and operational policy. Based on the author's experience, the difficulty of operating SLOs partly stems from the knowledge and objectives related to SLOs, and it is essential to proceed carefully through the introduction process. I feel that creating an atmosphere that can be described as a culture of SLOs is necessary to achieve an operational process that leads to improvements in services and ultimately the organization.

# Knowledge of SLOs
## Users and Services
Clearly define the users and services that are prerequisites for considering SLOs.

Users refer not only to humans using the service but also to any service users, including systems and robots.

On the other hand, a service is defined as any system where users exist. Examples include web services, APIs, batch jobs, databases, networks, and devices.

## Mindset about SLOs
Organize the mindset that should be kept in mind when operating SLOs.

### SLOs are Guiding Data
SLOs are not demands that impose something but are data that provide insights and guidance.

It is necessary to think about how to behave when operating SLOs.

### SLOs are a Continuous Process
SLOs are not a one-time achievement but are continuously operated.

It is required to update or abolish them if they become unnecessary, and they are not tasks with completion criteria.

SLOs must change according to changes in services.

Moreover, it is not the case that the benefits of operation can be obtained immediately upon introduction; patience is needed to continue and improve the process over time.

### SLOs Affect People
Ultimately, SLOs influence human behavior.

They aim to have a positive impact not only on the users targeted by the service but also on stakeholders involved in the service, such as engineers and business-side members.

## What is Reliability?
Reliability is defined as "the system executing the actions that users demand."

Providing 100% reliability is impossible, and it is necessary to consider the cost of providing reliability based on the actions users demand.

The level of reliability required varies by service. It is necessary to quantify this with SLOs and provide appropriate reliability.

## What is Service Level?
Service level indicates "the quality and performance provided by the service" and represents the level of reliability the service offers.

Service level encompasses the concepts of SLO and SLI.

For example, the user login process completing within a certain time.

## What is SLI (Service Level Indicator)?
SLI is defined as "an indicator that shows the characteristics of the service over a certain period." It provides data to measure service levels.

Indicators include availability, latency, response time, throughput, and error rate.

SLI is expressed as the ratio of good events divided by valid events, multiplied by 100.

SLI serves as an indicator of how the service functions from the user's perspective. Effective SLIs positively impact users, engineers, and the business. They can enhance user experience, guide engineers in identifying problems and improvement directions, and demonstrate service reliability to the business. Such SLIs are discovered from the perspective of what users need rather than what the service provides.

## What is SLO (Service Level Objective)?
SLO is defined as "the target value of service level over a certain period." SLO indicates the target value for SLI.

If the SLO is exceeded, users are satisfied with the service; if it falls below, users are dissatisfied.

SLO is not a consensus but a target for effort, and changes are permissible.

The purpose of SLOs is to collect data to quantify service reliability. By operating SLOs, it becomes possible to discover points for improvement in reliability and enhance development and operations from the perspective of reliability.

SLO target values are set at less than 100% (sometimes set in time). Since 100% reliability is realistically impossible, it is necessary to consider the cost of reliability.

| Availability | Annual Downtime | Monthly Downtime |
| ------------ | --------------- | ---------------- |
| 99.0%       | 87.6 hours      | 7.6 hours        |
| 99.5%       | 43.8 hours      | 3.65 hours       |
| 99.9%       | 8.76 hours      | 43.8 minutes     |
| 99.95%      | 4.38 hours      | 21.9 minutes     |
| 99.99%      | 52.56 seconds   | 4.38 minutes     |
| 99.999%     | 5.256 seconds   | 26.28 seconds    |
| 99.9999%    | 31.536 seconds  | 2.628 seconds    |

Reliability recovery is possible manually up to 99.9%, but beyond that, achieving such target values becomes difficult without automation.

## What is Error Budget?
Error budget is defined as "an indicator that shows the allowable loss of service reliability over a certain period." The error budget is the leeway to achieve SLOs and indicates the amount of cumulative errors until users become dissatisfied.

The error budget is calculated as 100% - SLO value.

The error budget serves as a basis for deciding whether to prioritize reliability or feature additions. The response based on the state of the error budget is not mandatory but serves as data for decision-making.

Measuring the error budget requires setting a time frame. There are two types: event-based and time-based. Event-based measures the number of occurrences, allowing observation of error counts within the time frame (e.g., 500 remaining errors in a week). Time-based measures the duration, allowing observation of error time within the time frame (e.g., 10 minutes remaining in a week).

Creating an error budget policy is beneficial for effective error budget operation.

- Clarification of ownership and stakeholders
  - Clearly define who owns the error budget and who is interested in it.
- Error budget consumption policy
  - Clearly define how to respond based on the consumption status of the error budget.
- Error budget exceeding policy
  - Clearly define the response when the error budget is exceeded.

In consumption and exceeding policies, it is good to indicate the temperature of action criteria such as "recommended," "advised," and "mandatory."

## Benefits of Operating SLOs
SLOs provide a means for decision-making based on SLOs from the perspectives of business, development, and operations.

From a business perspective, it serves as an indicator for decisions such as whether to invest in reliability or feature additions.

From a development perspective, it serves as an indicator for balancing feature additions and reliability improvements.

From an operations perspective, it helps discover points for improving reliability and indicates directions for enhancing reliability or seeking appropriate reliability.

By operating SLOs, the following positive changes may be expected:

- Prevent user dissatisfaction in advance
- Prevent over-delivering on user expectations
- Make it easier to judge whether user expectations should be improved
- Detect the ongoing impact of feature development on reliability
- Understand the impact of reliability decline on the business

# Introduction of SLOs
## Designing SLOs
SLO design can be carried out in the following steps.

1. Define the user journey
2. Define the architecture involved in the user journey
3. Define the SLI
4. Define the SLO

### 1. Define the User Journey
The user journey illustrates the flow when a user utilizes the service.

By defining the user journey, it becomes possible to clarify the functions provided by the service and the actions users demand.

For example, the user journey when purchasing a product: Add product to cart -> Enter address -> Select payment method -> Purchase

### 2. Define the Architecture Involved in the User Journey
By defining the architecture involved in the user journey, it becomes possible to clarify the system configuration necessary to realize the functions provided by the service and the actions users demand. Understanding the dependencies between systems and the flow of data helps in selecting metrics that could serve as SLIs.

For example, the architecture when a user purchases a product: Web server -> Database -> Payment service -> Message queue -> Notification service

### 3. Define the SLI
Based on the definitions of the user journey and architecture, define the SLI.

For example, the SLIs when a user purchases a product: Response time for the request to add a product to the cart, success rate for the request to enter the address, success rate for the request to select a payment method, success rate for the request to purchase.

### 4. Define the SLO
Based on the definition of the SLI, define the SLO.

For example, the SLOs when a user purchases a product: Response time for the request to add a product to the cart must be below 100ms, success rate for the request to enter the address must be above 99.9%, success rate for the request to select a payment method must be above 99.9%, success rate for the request to purchase must be above 99.9%.

Since SLOs are improved through a continuous process, there is no need to aim for perfection from the start; starting with a small scope can be a first step to begin the initiative.

## Operating SLOs
The established SLOs should be monitored, analyzed regularly, and continuously operated.

Regularly check the status of SLO achievement, and if there are issues, investigate based on the relevant SLIs and analyze the factors that affected them.

Consider responses to the status of SLO achievement based on the error budget operation policy.

Adjustments to SLO target values and SLIs, as well as the addition or deletion of SLOs, should also be made continuously. Changes in the environment surrounding the service may render previously defined SLOs or SLIs inappropriate.

The data obtained from operating SLOs and the discussions that arise are assets that contribute to improving service reliability, so it is essential to store and share information appropriately.

# References
- [amazon.com - SLO Service Level Objectives: A Practical Guide to Implementing SLI, SLO, and Error Budgets](https://amzn.to/4cSZVr1)
  - Particularly referenced book
  - Mainly chapters 1-7, 13-16 were referenced
- [cloud.google.com - Impact of Maintenance Windows on Error Budgets - SRE Tips](https://cloud.google.com/blog/ja/products/management-tools/sre-error-budgets-and-maintenance-windows)
- [cloud.google.com - Implementing SRE: Standardizing the SLO Design Process](https://cloud.google.com/blog/ja/products/devops-sre/how-to-design-good-slos-according-to-google-sres)
- [docs.google.com - SLO Documentation](https://docs.google.com/document/d/1SNgnAjRT1jrMa7vGHK0J_0jJEDvKJ5JmTEXFvNRDaHE/edit#heading=h.x9snb54sjlu9)
  - Template for creating SLOs
- [static.googleusercontent.com - The Art of SLOs](https://static.googleusercontent.com/media/sre.google/ja//static/pdf/jp-art-of-slos-handbook-pdf-a4.pdf)
- [newrelic.com - New Relic Hands-On: Basics of SLI/SLO Design](https://newrelic.com/sites/default/files/2023-05/20230510_NRU303.pdf)
  - Publicly available hands-on material
  - The content is well organized and easy to understand
- [newrelic.com - What are SLOs, SLIs, and SLAs?](https://newrelic.com/jp/blog/best-practices/what-are-slos-slis-slas)
- [newrelic.com - Best Practices for Setting SLI/SLO in Modern Systems](https://newrelic.com/jp/blog/best-practices/best-practices-for-setting-slos-and-slis-for-modern-complex-systems)
- [bmf-tech.com - About SLI, SLO, and SLA](https://bmf-tech.com/posts/SLI%E3%83%BBSLO%E3%83%BBSLA%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6)
- [sreake.com - What You Should Know Before Introducing SLI, SLO, and Error Budgets](https://sreake.com/blog/sli-slo-good-practices/)
  - Detailed writing from design to operation, which is helpful
  - Released on the same day as this article
- Case Studies
  - [blog.smartbank.co.jp - We Designed and Started Operating SLI/SLO - For Those Who Are About to Start Operating SLI/SLO](https://blog.smartbank.co.jp/entry/2023/05/25/104024)
  - [inside.dmm.com - Instilling SLI/SLO Culture in the Organization! 〜4 Steps Starting from Creating a Product Charter〜](https://inside.dmm.com/articles/sli-slo/)