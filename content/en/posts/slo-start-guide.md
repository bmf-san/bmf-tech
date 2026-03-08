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
This guide aims to help you understand SLOs and start operating them.

# Steps from Introducing to Operating SLOs
To introduce and start operating SLOs in an organization or team, several steps need to be followed.

1. Share knowledge about SLOs
2. Agree on the purpose of introducing SLOs
3. Agree on the operational policy of SLOs
4. Design the SLOs
5. Start operating the SLOs

When introducing SLOs to an organization for the first time, it is important to share knowledge about SLOs and agree on the purpose of introduction and operational policy. Based on my experience, I feel that the difficulty of operating SLOs partly stems from the knowledge and purpose of SLOs, and it is important to proceed carefully through the introduction process. I believe that without creating an atmosphere that can be called a culture, it is difficult to achieve an operational process that leads to the improvement of services and ultimately the organization.

# Knowledge of SLOs
## Users and Services
Define the users and services that are prerequisites when considering SLOs.

Users refer to **all service users**, not just humans using the service, but also systems and robots using the service.

On the other hand, **services refer to any system where users exist**. Examples include web services, APIs, batch jobs, databases, networks, and devices.

## Mindset about SLOs
Organize the mindset to keep in mind when operating SLOs.

### SLOs are Guiding Data
SLOs are not demands that force something but are data that provide insights and guidance.

You need to think for yourself about how to behave by operating SLOs.

### SLOs are a Continuous Process
SLOs are not something that ends once achieved but are continuously operated.

It is required to update during operation, and if unnecessary, to abolish them, and they are not tasks with completion conditions.

SLOs need to change according to changes in the service.

Moreover, even with the introduction, the benefits of operation are not immediately obtained, and it requires perseverance to continue and improve the process over time.

### SLOs Affect People
SLOs ultimately affect people's behavior.

They are intended to bring positive effects to stakeholders involved in the service, such as engineers and business members, not just the users targeted by the service.

## What is Reliability?
Reliability is "**the system performing the actions that users expect**."

Providing 100% reliability is impossible, and it is necessary to provide reliability considering the cost to provide the actions users expect.

The level of reliability required varies by service. It is necessary to quantify it with SLOs and provide appropriate reliability.

## What is a Service Level?
A service level is "**an indication of the quality or performance provided by the service**," showing the level of reliability provided by the service.

The service level is positioned as a concept that encompasses the concepts of SLO and SLI.

ex. Ensuring that user login processes are completed within a certain time frame.

## What is an SLI (Service Level Indicator)?
An SLI is "**an indicator showing the characteristics of a service over a certain period**." It provides data to measure the service level.

Indicators include availability, latency, response time, throughput, and error rate.

SLI is expressed as the percentage obtained by dividing good events by valid events and multiplying by 100.

SLI is an indicator showing how the service functions from the user's perspective. Effective SLIs positively impact users, engineers, and businesses. They improve user experience, indicate problem identification and improvement directions to engineers, and demonstrate service reliability to businesses. Such SLIs are discovered from the perspective of what users need, not what the service provides.

## What is an SLO (Service Level Objective)?
An SLO is "**a target value of the service level over a certain period**." An SLO indicates the target value for the SLI.

If the SLO is exceeded, users are considered satisfied with the service, and if it is not met, users are considered dissatisfied.

SLOs are not agreements but aspirational goals, and changes are allowed.

The purpose of SLOs is to collect data to quantify service reliability. Operating SLOs leads to discovering reliability improvement points and improving development and operation from the perspective of reliability.

SLO target values are set as percentages below 100% (sometimes set by time). Since 100% reliability is realistically impossible, it is necessary to consider the cost of reliability.

| Availability | Annual Downtime | Monthly Downtime |
| ------------ | --------------- | ---------------- |
| 99.0%        | 87.6 hours      | 7.6 hours        |
| 99.5%        | 43.8 hours      | 3.65 hours       |
| 99.9%        | 8.76 hours      | 43.8 minutes     |
| 99.95%       | 4.38 hours      | 21.9 minutes     |
| 99.99%       | 52.56 seconds   | 4.38 minutes     |
| 99.999%      | 5.256 seconds   | 26.28 seconds    |
| 99.9999%     | 31.536 seconds  | 2.628 seconds    |

Up to 99.9%, reliability recovery is possible manually, but beyond that, it becomes a target value that is difficult to achieve without assuming automation.

## What is an Error Budget?
An error budget is "**an indicator showing the allowable loss of service reliability over a certain period**." An error budget is the margin for achieving SLOs, indicating the cumulative allowable amount of errors until users become dissatisfied.

The error budget is calculated as 100% - SLO.

The error budget serves as a decision-making tool for prioritizing reliability or adding features. The response based on the error budget status is not enforced but serves as data for decision-making.

Measuring the error budget requires setting a time frame. There are two types of time frames: event-based and time-based. Event-based measures the number of occurrences, allowing observation of the number of errors within a time frame (e.g., 500 remaining errors in a week). Time-based measures time, allowing observation of error time within a time frame (e.g., 10 minutes remaining in a week).

Creating an error budget policy is beneficial for effective error budget operation.

- Clarification of ownership and stakeholders
  - Clearly define who owns the error budget and who is interested in it
- Error budget consumption policy
  - Clearly define the response based on the consumption status of the error budget
- Error budget overage policy
  - Clearly define the response if the error budget is exceeded

In consumption and overage policies, it is good to indicate the temperature of behavioral standards such as "should (recommended)", "ought to (advised)", and "must (mandatory)".

## Benefits of Operating SLOs
SLOs provide a means for decision-making based on SLOs from business, development, and operational perspectives.

From a business perspective, they serve as indicators for deciding whether to invest in reliability or feature additions.

From a development perspective, they serve as indicators for balancing feature additions and reliability improvements.

From an operational perspective, they help discover reliability improvement points, indicate directions for improving reliability, and serve as indicators for exploring appropriate reliability.

Operating SLOs may lead to positive changes such as:

- Mitigating user dissatisfaction in advance
- Preventing excessive responses to user expectations
- Making it easier to decide whether to improve user expectations
- Detecting the constant impact of feature development on reliability
- Understanding the impact of reliability degradation on business

# Introducing SLOs
## Designing SLOs
SLO design can be done through the following steps.

1. Define the user journey
2. Define the architecture involved in the user journey
3. Define SLIs
4. Define SLOs

### 1. Define the User Journey
The user journey shows the flow when a user uses the service.

Defining the user journey clarifies the functions provided by the service and the actions users expect.

ex. User journey when purchasing a product: Add product to cart -> Enter address -> Select payment method -> Purchase

### 2. Define the Architecture Involved in the User Journey
Defining the architecture involved in the user journey clarifies the system configuration to realize the functions provided by the service and the actions users expect. Understanding system dependencies and data flow helps select metrics that are candidates for SLIs.

ex. Architecture when purchasing a product: Web server -> Database -> Payment service -> Message queue -> Notification service

### 3. Define SLIs
Define SLIs based on the defined user journey and architecture.

ex. SLIs when purchasing a product: Response time for adding a product to the cart request, success rate for entering address request, success rate for selecting payment method request, success rate for purchase request

### 4. Define SLOs
Define SLOs based on the defined SLIs.

ex. SLOs when purchasing a product: Response time for adding a product to the cart request is 100ms or less, success rate for entering address request is 99.9% or more, success rate for selecting payment method request is 99.9% or more, success rate for purchase request is 99.9% or more

Since SLOs are improved through a continuous process, it is not necessary to aim for perfection from the beginning, and starting with a small scope is a step to begin the initiative.

## Operating SLOs
Monitor the set SLOs, analyze them regularly, and operate them continuously.

Regularly check the achievement status of SLOs, investigate based on the relevant SLIs if there are problems, and analyze the factors that influenced them.

Consider responses to the achievement status of SLOs based on the error budget operation policy.

Adjust SLO target values and SLIs, and add or delete SLOs as part of continuous operation. Changes in the environment surrounding the service may render previously defined SLOs and SLIs inappropriate.

The data obtained and discussions arising from SLO operation are assets that lead to improved service reliability, so it is important to store and share information appropriately.

# References
- [amazon.com - SLO サービスレベル目標 ―SLI、SLO、エラーバジェット導入の実践ガイド](https://amzn.to/4cSZVr1)
  - A particularly referenced book
  - Mainly referenced chapters 1-7, 13-16
- [cloud.google.com - メンテナンスの時間枠がエラー バジェットに与える影響 - SRE のヒント](https://cloud.google.com/blog/ja/products/management-tools/sre-error-budgets-and-maintenance-windows)
- [cloud.google.com - SRE の導入: SLO の設計プロセスを標準化する](https://cloud.google.com/blog/ja/products/devops-sre/how-to-design-good-slos-according-to-google-sres)
- [docs.google.com - SLO Documentation](https://docs.google.com/document/d/1SNgnAjRT1jrMa7vGHK0J_0jJEDvKJ5JmTEXFvNRDaHE/edit#heading=h.x9snb54sjlu9)
  - Template for creating SLOs
- [static.googleusercontent.com - The Art of SLOs](https://static.googleusercontent.com/media/sre.google/ja//static/pdf/jp-art-of-slos-handbook-pdf-a4.pdf)
- [newrelic.com - New Relic ハンズオン:SLI/SLOの設計の基本](https://newrelic.com/sites/default/files/2023-05/20230510_NRU303.pdf)
  - Publicly available hands-on material
  - Well-organized and easy to understand
- [newrelic.com - SLO、SLI、SLAとは何か?](https://newrelic.com/jp/blog/best-practices/what-are-slos-slis-slas)
- [newrelic.com - モダンなシステムにSLI/SLOを設定するときのベストプラクティス](https://newrelic.com/jp/blog/best-practices/best-practices-for-setting-slos-and-slis-for-modern-complex-systems)
- [bmf-tech.com - SLI・SLO・SLAについて](https://bmf-tech.com/posts/SLI%E3%83%BBSLO%E3%83%BBSLA%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6)
- [sreake.com - SLI、SLO、エラーバジェット導入の前に知っておきたいこと](https://sreake.com/blog/sli-slo-good-practices/)
  - Detailed from design to operation and useful
  - Released on the same day as this article
- Case Studies
  - [blog.smartbank.co.jp - 「我々はこうしてSLI/SLOを設計して運用を始めました -これからSLI/SLOの運用を始める人に向けて-」というタイトルで登壇してきました](https://blog.smartbank.co.jp/entry/2023/05/25/104024)
  - [inside.dmm.com - SLI/SLO文化を組織に浸透させる！ 〜プロダクト憲章作成から始める4ステップ〜](https://inside.dmm.com/articles/sli-slo/)