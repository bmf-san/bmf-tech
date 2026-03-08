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
translation_key: saas-architecture-resource-list
---

I have compiled some resources that may help you move from not understanding anything about SaaS to fully grasping it!

# Resources
- [docs.aws.amazon.com - SaaS Architecture Fundamentals](https://docs.aws.amazon.com/ja_jp/whitepapers/latest/saas-architecture-fundamentals/saas-architecture-fundamentals.html)
- [docs.aws.amazon.com - SaaS Lens](https://docs.aws.amazon.com/ja_jp/wellarchitected/latest/saas-lens/saas-lens.html)
- [docs.aws.amazon.com - SaaS Tenant Isolation Strategies: Isolating Resources in a Multi-Tenant Environment](https://docs.aws.amazon.com/whitepapers/latest/saas-tenant-isolation-strategies/saas-tenant-isolation-strategies.html)
- [docs.aws.amazon.com - SaaS Storage Strategies](https://docs.aws.amazon.com/whitepapers/latest/multi-tenant-saas-storage-strategies/multi-tenant-saas-storage-strategies.html)
- [d1.awsstatic.com - SaaS Tenant Isolation Strategies](https://d1.awsstatic.com/whitepapers/ja_JP/saas-tenant-isolation-strategies.pdf)
- [speakerdeck.com - Overview of SaaS Architecture](https://speakerdeck.com/ryurock/saas-akitekutiyagai-yao)
- [dev.classmethod.jp - [Report] Learning SaaS Architecture Patterns #reinvent #ARC306](https://dev.classmethod.jp/articles/reinvent2021-arc306/)

# Important Considerations for SaaS Architecture in Business Models

- Can you provide flexible services to users (subscription plans, features, etc.)?
- Can you onboard users quickly (start using the service)?
- Can you manage analytics, metrics, and billing appropriately?
- Can you ensure non-functional requirements considering noisy neighbors?
- Can you adequately address data leakage (ensuring that data not visible to other users between tenants remains secure)?
  - There are various methods for isolating resources used by tenants.
    - cf. [docs.aws.amazon.com - Basic Isolation Concepts](https://docs.aws.amazon.com/ja_jp/wellarchitected/latest/saas-lens/core-isolation-concepts.html)
- What resources do you want to separate in terms of functionality, performance, and data?

# Thoughts
- Cost optimization seems challenging. If a specific tenant consumes excessive resources but ends up with the same billing as others, it could be problematic...
- There may be needs to share tenants between parent and subsidiary companies or to partially integrate data, which could lead to trade-offs.