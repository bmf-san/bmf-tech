---
title: A List to Understand SaaS Architecture
slug: saas-architecture-resource-list
date: 2024-02-24T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Architecture
  - SaaS
description: A compilation of resources to transition from knowing nothing about SaaS to fully understanding it.
translation_key: saas-architecture-resource-list
---

To move from "I don't understand anything about SaaS!" to "I completely understand SaaS!", here is a compilation of useful resources.

# Resources
- [docs.aws.amazon.com - Fundamentals of SaaS Architecture](https://docs.aws.amazon.com/ja_jp/whitepapers/latest/saas-architecture-fundamentals/saas-architecture-fundamentals.html)
- [docs.aws.amazon.com - SaaS Lens](https://docs.aws.amazon.com/ja_jp/wellarchitected/latest/saas-lens/saas-lens.html)
- [docs.aws.amazon.com - SaaS Tenant Isolation Strategies: Isolating Resources in a Multi-Tenant Environment](https://docs.aws.amazon.com/whitepapers/latest/saas-tenant-isolation-strategies/saas-tenant-isolation-strategies.html)
- [docs.aws.amazon.com - SaaS Storage Strategies](https://docs.aws.amazon.com/whitepapers/latest/multi-tenant-saas-storage-strategies/multi-tenant-saas-storage-strategies.html)
- [d1.awsstatic.com - SaaS Tenant Isolation Strategies](https://d1.awsstatic.com/whitepapers/ja_JP/saas-tenant-isolation-strategies.pdf)
- [speakerdeck.com - Overview of SaaS Architecture](https://speakerdeck.com/ryurock/saas-akitekutiyagai-yao)
- [dev.classmethod.jp - [Report] Learning SaaS Architecture Patterns #reinvent #ARC306](https://dev.classmethod.jp/articles/reinvent2021-arc306/)

# Considerations for Architecture in SaaS Business Models

- Can you provide flexible services to users (subscription plans, features, etc.)?
- Can you onboard users quickly (getting them started with the service)?
- Can you manage analytics, metrics, and billing appropriately?
- Can you ensure non-functional requirements considering noisy neighbors?
- Can you adequately address data leakage (ensuring data that should not be visible to other users remains hidden)?
  - There are various methods for isolating resources used by tenants.
    - cf. [docs.aws.amazon.com - Basic Concepts of Isolation](https://docs.aws.amazon.com/ja_jp/wellarchitected/latest/saas-lens/core-isolation-concepts.html)
- What resources do you want to separate in terms of functionality, performance, and data?

# Thoughts
- Cost optimization seems challenging. If a specific tenant consumes excessive resources but is billed similarly to others, it could lead to issues...
- There may be needs to share tenants between a parent company and a subsidiary or to partially integrate data, which could be complicated. There are trade-offs involved.
