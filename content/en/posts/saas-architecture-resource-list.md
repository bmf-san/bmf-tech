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


SaaSなんもわからん！からSaaS完全に理解した！に近づくために参考になりそうな資料をまとめておく。

# Resources
- [docs.aws.amazon.com - SaaS アーキテクチャの基礎](https://docs.aws.amazon.com/ja_jp/whitepapers/latest/saas-architecture-fundamentals/saas-architecture-fundamentals.html)
- [docs.aws.amazon.com - SaaSレンズ](https://docs.aws.amazon.com/ja_jp/wellarchitected/latest/saas-lens/saas-lens.html)
- [docs.aws.amazon.com - SaaS Tenant Isolation Strategies: Isolating Resources in a Multi-Tenant Environment](https://docs.aws.amazon.com/whitepapers/latest/saas-tenant-isolation-strategies/saas-tenant-isolation-strategies.html)
- [docs.aws.amazon.com - SaaS Storage Strategies](https://docs.aws.amazon.com/whitepapers/latest/multi-tenant-saas-storage-strategies/multi-tenant-saas-storage-strategies.html)
- [d1.awsstatic.com - SaaS のテナント分離戦略](https://d1.awsstatic.com/whitepapers/ja_JP/saas-tenant-isolation-strategies.pdf)
- [speakerdeck.com - SaaS アーキテクチャ概要](https://speakerdeck.com/ryurock/saas-akitekutiyagai-yao)
- [dev.classmethod.jp - [レポート] SaaSアーキテクチャーのパターンを学ぶ #reinvent #ARC306](https://dev.classmethod.jp/articles/reinvent2021-arc306/)

# Important Considerations for SaaS Architecture in Business Models

- Can flexible services (subscription plans, features, etc.) be provided to users?
- Can user onboarding (starting service use) be expedited?
- Can analytics, metrics, and billing be managed appropriately?
- Can non-functional requirements considering noisy neighbors be ensured?
- Can data breaches (such as data visible to other users between tenants) be appropriately addressed?
  - There are various methods to isolate resources used by tenants
    - cf. [docs.aws.amazon.com - 基本的な分離の概念](https://docs.aws.amazon.com/ja_jp/wellarchitected/latest/saas-lens/core-isolation-concepts.html)
- What are the resources to be isolated in terms of functionality, performance, data, etc.?

# Thoughts
- Cost optimization seems challenging. If a specific tenant consumes excessive resources but the billing remains similar to other tenants, it could become problematic...
- There might be a need to share tenants between parent and subsidiary companies or partially integrate data, which could be tricky. There are trade-offs.