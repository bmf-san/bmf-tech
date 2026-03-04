---
title: "SaaSのアーキテクチャについて知るためのリスト"
slug: "saas-architecture-resource-list"
date: 2024-02-24
author: bmf-san
categories:
  - "アーキテクチャ"
tags:
  - "アーキテクチャ"
  - "SaaS"
draft: false
---

SaaSなんもわからん！からSaaS完全に理解した！に近づくために参考になりそうな資料をまとめておく。

# 資料
- [docs.aws.amazon.com - SaaS アーキテクチャの基礎](https://docs.aws.amazon.com/ja_jp/whitepapers/latest/saas-architecture-fundamentals/saas-architecture-fundamentals.html)
- [docs.aws.amazon.com - SaaSレンズ](https://docs.aws.amazon.com/ja_jp/wellarchitected/latest/saas-lens/saas-lens.html)
- [docs.aws.amazon.com - SaaS Tenant Isolation Strategies: Isolating Resources in a Multi-Tenant Environment](https://docs.aws.amazon.com/whitepapers/latest/saas-tenant-isolation-strategies/saas-tenant-isolation-strategies.html)
- [docs.aws.amazon.com - SaaS Storage Strategies](https://docs.aws.amazon.com/whitepapers/latest/multi-tenant-saas-storage-strategies/multi-tenant-saas-storage-strategies.html)
- [d1.awsstatic.com - SaaS のテナント分離戦略](https://d1.awsstatic.com/whitepapers/ja_JP/saas-tenant-isolation-strategies.pdf)
- [speakerdeck.com - SaaS アーキテクチャ概要](https://speakerdeck.com/ryurock/saas-akitekutiyagai-yao)
- [dev.classmethod.jp - [レポート] SaaSアーキテクチャーのパターンを学ぶ #reinvent #ARC306](https://dev.classmethod.jp/articles/reinvent2021-arc306/)


# SaaSのビジネスモデル上アーキテクチャ検討で重要であると思われること

- ユーザーへの柔軟なサービス提供（サブスプリションプラン、機能など）ができるか？
- ユーザーへの導入（オンボーディング。サービス利用開始。）を早く行うことができるか？
- 分析やメトリクス、請求が適切に管理できるか？
- ノイジーネイバーを考慮した非機能要件の担保ができるか？
- データ漏えい（テナント間などで他ユーザーに見えてはいけないデータが見えてしまう）を適切に対策できるか？
  - テナントが利用するリソースをどのように分離するか色々な方法がある
    - cf. [docs.aws.amazon.com - 基本的な分離の概念](https://docs.aws.amazon.com/ja_jp/wellarchitected/latest/saas-lens/core-isolation-concepts.html)
- 機能、性能、データなど分離したい対象リソースは何か？

# 所感
- コスト最適化が難しそう。特定のテナントがやたらリソースを消費しているけど、他のテナントと変わらないような請求になってしまう、とかしてしまうと・・・
- 親会社、小会社でテナントを共有したい、データを一部連携したい、とかそういったニーズを汲むのとかウッてなりそう。トレードオフがある。
