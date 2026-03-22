---
title: A Starter Guide to Implementing SLOs
slug: slo-start-guide
date: 2024-07-23T00:00:00Z
author: bmf-san
categories:
  - Operations
tags:
  - Reliability
  - SLO
description: A guide to understanding and starting the operation of SLOs.
translation_key: slo-start-guide
---

# Overview
This content serves as a guide to understand SLO and to begin its operation.

# Steps from Introduction to Operation of SLO
To introduce SLO into an organization or team and start its operation, several stages need to be followed.

1. Share knowledge about SLO
2. Agree on the purpose of introducing SLO
3. Agree on the operational policy for SLO
4. Design the SLO
5. Start the operation of SLO

When introducing SLO for the first time into an organization, it is important to first share knowledge about SLO and reach an agreement on its purpose and operational policy. Based on the author's experience, the difficulty of operating SLO partly stems from the knowledge and objectives surrounding SLO, and it is crucial to carefully proceed through the introduction process. I feel that creating an atmosphere where SLO can be considered a culture is essential to achieving an operational process that leads to improvements in services and ultimately the organization.

# Knowledge of SLO
## Users and Services
Clearly define the users and services that are prerequisites for considering SLO.

Users refer not only to humans who utilize the service but also to any **users of the service**, including systems and robots that use the service.

On the other hand, "a **service is some system where users exist**." Examples include web services, APIs, batch jobs, databases, networks, and devices.

## Mindset for SLO
Organize the mindset that should be kept in mind when operating SLO.

### SLO is Guiding Data
SLO is not a demand that imposes something but rather data that provides insights and serves as guidance.

It is necessary to think for oneself about how to behave when operating SLO.

### SLO is a Continuous Process
SLO is not something that ends once achieved; it is something that is continuously operated.

It is required to update during operation or abolish it if it becomes unnecessary, and it is not a task with completion criteria.

SLO must also change according to changes in services.

Moreover, even with the introduction, immediate operational benefits are not guaranteed, and patience is needed to continue and improve the process over time.

### SLO Affects People
Ultimately, SLO influences human behavior.

It aims to positively impact not only the users targeted by the service but also stakeholders involved in the service, such as engineers and business-side members.

## What is Reliability?
Reliability is defined as "**the system executing the actions that users require**."

Providing 100% reliability is impossible, and it is necessary to consider the costs of providing reliability to deliver the actions that users seek.

The level of reliability needed varies by service. It is required to quantify this through SLO and provide appropriate reliability.

## What is Service Level?
Service level refers to "**the quality or performance that a service provides**," indicating the level of reliability that the service offers.

Service level is positioned as a concept that encompasses SLO and SLI.

For example, the completion of a user's login process within a certain time.

## What is SLI (Service Level Indicator)?
SLI is defined as "**an indicator that shows the characteristics of a service over a certain period**." It provides data to measure service levels.

Indicators include availability, latency, response time, throughput, and error rate.

SLI is expressed as the ratio of good events divided by valid events multiplied by 100.

SLI serves as an indicator of how the service functions from the user's perspective. Effective SLI positively influences users, engineers, and the business. It can enhance user experience, guide engineers in identifying problems and improvement directions, and demonstrate service reliability to the business. Such SLIs are discovered from the perspective of what users need rather than what the service provides.

## What is SLO (Service Level Objective)?
SLO is defined as "**the target value for service level over a certain period**." SLO indicates the target value for SLI.

If SLO is exceeded, users are expected to be satisfied with the service; if it is not met, users are likely to be dissatisfied.

SLO is not a binding agreement but rather a target for effort, and changes are permissible.

The purpose of SLO is to gather data to quantify the reliability of the service. By operating SLO, it enables the discovery of improvement points for reliability and improvements in development and operations from the perspective of reliability.

SLO target values are set at less than 100% (sometimes even in terms of time). Since 100% reliability is practically impossible, it is necessary to consider the costs of reliability.

| Availability | Annual Downtime | Monthly Downtime |
| ------------ | --------------- | ---------------- |
| 99.0%        | 87.6 hours      | 7.6 hours        |
| 99.5%        | 43.8 hours      | 3.65 hours       |
| 99.9%        | 8.76 hours      | 43.8 minutes     |
| 99.95%       | 4.38 hours      | 21.9 minutes     |
| 99.99%       | 52.56 seconds   | 4.38 minutes     |
| 99.999%      | 5.256 seconds   | 26.28 seconds    |
| 99.9999%     | 31.536 seconds  | 2.628 seconds    |

Reliability recovery by manual means is possible up to 99.9%, but beyond that, achieving such target values becomes difficult without automation.

## What is an Error Budget?
An error budget is defined as "**an indicator that shows the allowable loss of service reliability over a certain period**." The error budget is the leeway to achieve SLO and indicates the amount of cumulative errors that can occur before users become dissatisfied.

The error budget is calculated as 100% - SLO value.

The error budget serves as a basis for deciding whether to prioritize reliability or feature additions. The response based on the state of the error budget is not mandatory but serves as data for decision-making.

Setting a time frame is necessary for measuring the error budget. There are two types of time frames: event-based and time-based. Event-based measures the number of occurrences, allowing observation of the number of errors within a time frame (e.g., 500 remaining errors in a week). Time-based measures the duration, allowing observation of error time within a time frame (e.g., 10 minutes remaining in a week).

Creating an error budget policy is beneficial for effective error budget operation.

- Clarification of ownership and stakeholders
  - Clearly define who owns the error budget and who is concerned about it.
- Error budget consumption policy
  - Clearly define how to respond based on the consumption status of the error budget.
- Error budget exceeding policy
  - Clearly define the response when the error budget is exceeded.

In consumption and exceeding policies, it is good to indicate the temperature of action standards such as "recommended," "advised," and "mandatory."

## Benefits of Operating SLO
SLO provides a means for decision-making based on SLO from the perspectives of business, development, and operations.

From the business perspective, it serves as an indicator for decisions such as whether to invest in reliability or feature additions.

From the development perspective, it serves as an indicator for balancing feature additions and reliability improvements.

From the operations perspective, it helps discover points for improving reliability and indicates directions for enhancing reliability or seeking appropriate reliability.

By operating SLO, the following positive changes may be expected:

- Preventing user dissatisfaction in advance
- Preventing over-delivery on user expectations
- Making it easier to judge whether user expectations should be improved
- Detecting the ongoing impact of feature development on reliability
- Understanding the impact of reliability degradation on the business

# Introduction of SLO
## Designing SLO
The design of SLO can be carried out in the following steps.

1. Define the user journey
2. Define the architecture involved in the user journey
3. Define the SLI
4. Define the SLO

### 1. Define the User Journey
The user journey illustrates the flow of how users utilize the service.

By defining the user journey, it becomes possible to clarify the functions provided by the service and the actions required by users.

For example, the user journey when a user purchases a product: Add the product to the cart -> Enter the address -> Choose the payment method -> Purchase.

### 2. Define the Architecture Involved in the User Journey
By defining the architecture involved in the user journey, it becomes possible to clarify the system configuration required to realize the functions provided by the service and the actions required by users. Understanding the dependencies between systems and the flow of data helps in selecting candidate metrics for SLI.

For example, the architecture when a user purchases a product: Web server -> Database -> Payment service -> Message queue -> Notification service.

### 3. Define the SLI
Based on the definitions of the user journey and architecture, define the SLI.

For example, the SLI when a user purchases a product: Response time for the request to add the product to the cart, success rate of the request to enter the address, success rate of the request to choose the payment method, success rate of the request to purchase.

### 4. Define the SLO
Based on the definition of SLI, define the SLO.

For example, the SLO when a user purchases a product: The response time for the request to add the product to the cart should be less than 100ms, the success rate of the request to enter the address should be 99.9% or higher, the success rate of the request to choose the payment method should be 99.9% or higher, and the success rate of the request to purchase should be 99.9% or higher.

Since SLO is improved through a continuous process, there is no need to aim for perfection from the start; beginning with a small scope can be a first step to getting started.

## Operating SLO
The established SLO should be monitored, analyzed regularly, and operated continuously.

Regularly check the achievement status of SLO, and if there are issues, investigate based on the relevant SLI and analyze the factors that affected it.

Based on the error budget operation policy, consider responses to the achievement status of SLO.

Adjustments to SLO target values and SLI, as well as the addition or deletion of new SLOs, should also be made continuously. Changes in the environment surrounding the service may render previously defined SLOs and SLIs inappropriate.

The data obtained from the operation of SLO and the discussions that arise are assets that lead to improvements in service reliability, so it is important to store and share information appropriately.

# References
- [amazon.com - SLO Service Level Objectives: A Practical Guide to Implementing SLI, SLO, and Error Budgets](https://amzn.to/4cSZVr1)
  - Particularly referenced book
  - Mainly chapters 1-7, 13-16 were referenced
- [cloud.google.com - The Impact of Maintenance Windows on Error Budgets - SRE Tips](https://cloud.google.com/blog/ja/products/management-tools/sre-error-budgets-and-maintenance-windows)
- [cloud.google.com - Implementing SRE: Standardizing the SLO Design Process](https://cloud.google.com/blog/ja/products/devops-sre/how-to-design-good-slos-according-to-google-sres)
- [docs.google.com - SLO Documentation](https://docs.google.com/document/d/1SNgnAjRT1jrMa7vGHK0J_0jJEDvKJ5JmTEXFvNRDaHE/edit#heading=h.x9snb54sjlu9)
  - Template for creating SLO
- [static.googleusercontent.com - The Art of SLOs](https://static.googleusercontent.com/media/sre.google/ja//static/pdf/jp-art-of-slos-handbook-pdf-a4.pdf)
- [newrelic.com - New Relic Hands-On: Basics of Designing SLI/SLO](https://newrelic.com/sites/default/files/2023-05/20230510_NRU303.pdf)
  - Published hands-on materials
  - The content is well organized and easy to understand
- [newrelic.com - What are SLOs, SLIs, and SLAs?](https://newrelic.com/jp/blog/best-practices/what-are-slos-slis-slas)
- [newrelic.com - Best Practices for Setting SLI/SLO in Modern Systems](https://newrelic.com/jp/blog/best-practices/best-practices-for-setting-slos-and-slis-for-modern-complex-systems)
- [bmf-tech.com - About SLI, SLO, and SLA](https://bmf-tech.com/posts/SLI%E3%83%BBSLO%E3%83%BBSLA%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6)
- [sreake.com - What You Should Know Before Implementing SLI, SLO, and Error Budgets](https://sreake.com/blog/sli-slo-good-practices/)
  - Detailed writing from design to operation, which is helpful
  - Released on the same day as this article
- Case Studies
  - [blog.smartbank.co.jp - "We Designed and Started Operating SLI/SLO - For Those Who Are About to Start Operating SLI/SLO"](https://blog.smartbank.co.jp/entry/2023/05/25/104024)
  - [inside.dmm.com - Infiltrating SLI/SLO Culture into the Organization! - 4 Steps Starting from Creating a Product Charter](https://inside.dmm.com/articles/sli-slo/)
