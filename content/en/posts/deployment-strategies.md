---
title: About Deployment Strategies
slug: deployment-strategies
date: 2023-06-03T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Deployment
translation_key: deployment-strategies
---



# Overview
Summarizing deployment strategies.

# Definitions of Deploy, Release, and Rollback
As a prerequisite for deployment strategies, let's clarify the definitions of these terms.

Deploy: "Placing an executable program in the execution environment."
Release: "Making it accessible to users."
Rollback: "Releasing an older version."

# Types of Deployment Strategies
Let's look at some representative strategies.

## In-Place Deployment
A method of deploying a new version directly to the existing environment.

As a side note, bmf-tech.com uses docker-compose, and the deployment is in-place deployment...
cf. [Technology Supporting bmf-tech](https://bmf-tech.com/posts/bmf-tech%e3%82%92%e6%94%af%e3%81%88%e3%82%8b%e6%8a%80%e8%a1%93#5-%E3%83%87%E3%83%97%E3%83%AD%E3%82%A4)

## Symbolic Link Deployment
A method of switching between old and new versions using symbolic links.

## Blue-Green Deployment
A method where two environments, blue and green, are prepared, and the new version is deployed to one, temporarily deploying both old and new versions. If there are no issues with the new version, traffic is switched from the old version to the new version. Both blue and green environments are maintained.

## Immutable Deployment
Similar to blue-green, but differs in that the old environment is deleted after traffic is switched.

## Rolling Deployment
A method of deploying and releasing the new version in increments. Until completion, traffic to both old and new environments remains active.

## Canary Deployment
A method where only a portion of users or traffic is exposed to the new version, allowing for partial deployment and release of the new version.

# Deployment Strategy Evaluation Table
For each deployment strategy, important criteria for selection are picked and summarized in a table.

|        Deployment Method        | Zero Downtime | Production Testing |           Rollback Time           | Operational Cost |
| ------------------------------- | ------------- | ------------------ | --------------------------------- | ---------------- |
| In-Place Deployment             | ×             | ×                  | High                              | Low              |
| Symbolic Link Deployment        | ○             | ×                  | Low                               | Medium           |
| Blue-Green Deployment           | ○             | ○                  | Low (even before/after traffic switch) | High           |
| Immutable Deployment            | ○             | ○                  | Low (limited to before old environment deletion) | High |
| Rolling Deployment              | ○             | ×                  | Low                               | Medium           |
| Canary Deployment               | ○             | ○                  | Medium                            | High             |

- Zero Downtime
  - Whether service downtime occurs during deployment
  - ○ if zero downtime is possible, ✗ if not
- Production Testing
  - Whether traffic can be routed to the new version
  - ○ if possible, ✗ if not
- Rollback Time
  - The time cost required for rollback
  - Judged as Low, Medium, or High
- Operational Cost
  - The operational cost to organize the deployment method
  - Judged as Low, Medium, or High

# Release Strategies
- [Deploy / Release Methods Summary](https://garafu.blogspot.com/2018/11/release-strategy.html)
- [Meaning and Differences of Build / Deploy / Release / Rollback](https://garafu.blogspot.com/2018/11/build-deploy-release-rollback.html)
- [What is Deployment](https://cmc-japan.co.jp/blog/%E3%83%87%E3%83%97%E3%83%AD%E3%82%A4%E3%81%A8%E3%81%AF/)
- [Application Deployment and Testing Strategies](https://cloud.google.com/architecture/application-deployment-and-testing-strategies?hl=ja)
- [Software Deployment](https://ja.wikipedia.org/wiki/%E3%82%BD%E3%83%95%E3%83%88%E3%82%A6%E3%82%A7%E3%82%A2%E3%83%87%E3%83%97%E3%83%AD%E3%82%A4%E3%83%A1%E3%83%B3%E3%83%88)
