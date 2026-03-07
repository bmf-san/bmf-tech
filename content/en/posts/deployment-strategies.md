---
title: Deployment Strategies
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
Clarifying the definitions of terms as prerequisite knowledge for deployment strategies.

- **Deploy**: "Placing executable programs in the execution environment"
- **Release**: "Making it accessible to users"
- **Rollback**: "Releasing an older version"

# Types of Deployment Strategies
Highlighting some representative strategies.

## In-Place Deployment
A method of directly deploying a new version to the existing environment.

As a side note, bmf-tech.com uses docker-compose, and the deployment is done via in-place deployment...
cf. [bmf-tech Technologies](https://bmf-tech.com/posts/bmf-tech%e3%82%92%e6%94%af%e3%81%88%e3%82%8b%e6%8a%80%e8%a1%93#5-%E3%83%87%E3%83%97%E3%83%AD%E3%82%A4)

## Symbolic Link Deployment
A method of switching between old and new versions using symbolic links.

## Blue-Green Deployment
Preparing two environments, Blue and Green, deploying the new version to one, and temporarily running both versions. If there are no issues with the new version, traffic is switched from the old version to the new version. Both Blue and Green environments are maintained.

## Immutable Deployment
Generally similar to Blue-Green, but the key difference is that the old environment is deleted after the traffic switch.

## Rolling Deployment
A method of deploying and releasing the new version in increments. Traffic to both the old and new environments remains active until all are completed.

## Canary Deployment
A method that allows deploying and releasing a new version only to a subset of users or traffic.

# Deployment Strategy Perspective Table
Picking up important perspectives for selecting each deployment strategy and summarizing them in a table.

| Deployment Method          | Zero Downtime | Production Testing | Rollback Time          | Operational Cost |
| -------------------------- | -------------- | ------------------ | --------------------- | ---------------- |
| In-Place Deployment        | ×              | ×                  | High                  | Low              |
| Symbolic Link Deployment   | ○              | ×                  | Low                   | Medium           |
| Blue-Green Deployment      | ○              | ○                  | Low (both before and after traffic switch) | High             |
| Immutable Deployment       | ○              | ○                  | Low (until old environment is deleted) | High             |
| Rolling Deployment         | ○              | ×                  | Low                   | Medium           |
| Canary Deployment          | ○              | ○                  | Medium                | High             |

- **Zero Downtime**
  - Whether service interruption occurs during deployment.
  - If zero downtime is possible, mark as ○; if not, mark as ✗.
- **Production Testing**
  - Whether traffic can be routed to the new version.
  - If possible, mark as ○; if not, mark as ✗.
- **Rollback Time**
  - The time cost required for rollback.
  - Judged as Low, Medium, or High.
- **Operational Cost**
  - The operational cost for organizing the deployment method.
  - Judged as Low, Medium, or High.

# Release Strategies
- [Summary of Deployment / Release Methods](https://garafu.blogspot.com/2018/11/release-strategy.html)
- [Meaning and Differences of Build / Deploy / Release / Rollback](https://garafu.blogspot.com/2018/11/build-deploy-release-rollback.html)
- [What is Deployment](https://cmc-japan.co.jp/blog/%E3%83%87%E3%83%97%E3%83%AD%E3%82%A4%E3%81%A8%E3%81%AF/)
- [Strategies for Application Deployment and Testing](https://cloud.google.com/architecture/application-deployment-and-testing-strategies?hl=ja)
- [Software Deployment](https://ja.wikipedia.org/wiki/%E3%82%BD%E3%83%95%E3%83%88%E3%82%A6%E3%82%A7%E3%82%A2%E3%83%87%E3%83%97%E3%83%AD%E3%82%A4%E3%83%A1%E3%83%B3%E3%83%88)