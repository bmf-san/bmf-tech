---
title: "契約テストとPactについて"
slug: "pact"
date: 2024-07-21
author: bmf-san
categories:
  - "テスト"
tags:
  - "契約テスト"
  - "Pact"
  - "プロデューサー駆動契約テスト"
  - "コンシューマー駆動契約テスト"
draft: false
---

# 概要
契約テストとPactについて調べたことをまとめる。

# 契約テスト（Contract Testing）とは
コンシューマ（サービスを呼び出す側）とプロバイダー（サービスを提供する側）の間の動作をテストするテスト手法のことである。

コンシューマーとプロバイダー間での合意を契約として記述し、その契約に基づいてテストを行う。

コンシューマーが契約を発行し、プロバイダーはその契約に従ってテストを行う形式の契約テストを**コンシューマー駆動契約テスト**と呼ぶ。

一方で、プロバイダーが契約を発行し、コンシューマーはその契約に従ってテストを行う形式の契約テストを**プロバイダー駆動契約テスト**と呼ぶ。

コンシューマー駆動契約テストとプロバイダー駆動契約テストのどちらが向いているかは、コンシューマーの数やAPIの安定性（変更頻度）、開発プロセスなど主導権をどちらが持ちやすいかによって異なる。

双方向契約をサポートする[pactflow](https://pactflow.io/)というサービスもある。[pactflow](https://pactflow.io/)は、コンシューマ駆動契約テストをサポートするPactをベースにしている。[pactflow](https://pactflow.io/)はpactと違ってOSSではない。

cf.
- [www.xlsoft.com - SmartBear Pactflow: コントラクト テストとは? 使うべき理由](https://www.xlsoft.com/jp/blog/blog/2022/10/18/smartbear-32158-post-32158/)
- [ipsj.ixsq.nii.ac.jp - 消費者駆動契約テストパターンとその課題](https://ipsj.ixsq.nii.ac.jp/ej/?action=repository_uri&item_id=193867&file_id=1&file_no=1)
- [www.ibm.com - 契約テスト](https://www.ibm.com/docs/ja/devops-test-workbench/11.0.0?topic=tasks-contract-testing)
- [gitlab-docs.creationline.com - 契約テスト](https://gitlab-docs.creationline.com/ee/development/testing_guide/contract/)
- [docs.pactflow.io - Bi-Directional Contract Testing Guide](https://docs.pactflow.io/docs/bi-directional-contract-testing)
- [pactflow.io - Bi-Directional Contract Testing](https://pactflow.io/bi-directional-contract-testing/)
- [pactflow.io - Pact is dead, long live Pact](https://pactflow.io/blog/bi-directional-contracts/)
- [www.thoughtworks.com - Pactflow](https://www.thoughtworks.com/radar/tools/pactflow)
- [technology.lastminute.com - Impacts of contract tests in our microservice architecture](https://technology.lastminute.com/impacts-of-contract-tests-in-a-microservice-architecture/)
- [alexromanov.github.io - Should You Use Contract Testing?](https://alexromanov.github.io/2021/07/12/should-you-use-contract-testing/)

## APIスキーマと契約テスト
Open APIやProtoBufなどのスキーマ駆動のAPI開発が整っていれば、契約テストの必要性に疑問を感じるかもしれない。

APIスキーマはAPIの仕様を定義する文書であり、エンドポイントやリクエスト・レスポンス、データモデルなどの要素を含んでいる。

スキーマ駆動をサポートするツールを使えば、APIの定義を元にコンシューマー側とプロバイダー側のコードそれぞれ生成することができる。

APIスキーマ通りに実装がされていることについて一定の保証を得ることができても、APIの振る舞い（実装されたAPIが期待通りに動作するか）を保証することができない。

スキーマー駆動のみではプロバイダーの変更がコンシューマーに通知されることをコードファーストで担保することができないため、コミュニケーションやテストの手間が発生する。

契約テストはサービス間の動作をテストするための手法であり、APIスキーマはAPIの仕様を定義するための文書であるため、APIスキーマと契約テストは異なる目的を持っている。

cf.
- [Schemas are not contracts](https://pactflow.io/blog/schemas-are-not-contracts/)
- [pactflow.io - Schema-based contract testing with JSON schemas and Open API (Part 1)](https://pactflow.io/blog/contract-testing-using-json-schemas-and-open-api-part-1/)
- [pactflow.io - Schema-based contract testing with JSON schemas and Open API (Part 2)](https://pactflow.io/blog/contract-testing-using-json-schemas-and-open-api-part-2/)
- [フロントエンドへのPact導入 - 契約によるAPIのテストを始めよう](https://blog.techscore.com/entry/2020/10/16/080000#OpenAPI%E3%81%A8%E3%81%AE%E9%81%95%E3%81%84%E3%81%AF%E3%81%AA%E3%81%AB)

## 契約テストのメリットとデメリット
コンシューマーまたはプロバイダー主導のどちらかに寄らない契約テストのメリットとデメリットについて整理する。

cf.
- [docs.pact.io - Convince me](https://docs.pact.io/faq/convinceme)
- [www.infoq.com - マイクロサービスアプリケーションにコントラクトテストを使用する](https://www.infoq.com/jp/news/2019/04/contract-testing-microservices/)

### メリット
- サービス間の信頼性・整合性を保てる
- コンシューマーまたはプロバイダーの変更を自動で検知できる
  - チーム間のコミュニケーションコストを削減できる
- E2Eよりも実行速度が速い
- サービス間の依存関係を明確にできる

### デメリット
- ツール依存・ツール導入のコスト
- 導入に当たって組織やチーム間の合意形成が必要
- 開発フローの中の1つのプロセスとして組み込む必要がある

# Pactとは
契約テストのためのツールで、コンシューマー駆動契約テストをサポートする。プロバイダー駆動契約テストはサポートされていない。

多言語に対応しており、テスティングフレームワークやビルドツールなどとの連携も可能になっている。

Pactでは、HTTPとメッセージ（非HTTPなメッセージキュー。Pact自体はRabbit MQやSQS、Kafkaなどの実装の詳細を把握せずにメッセージそのものをテストできるようになっている）をサポートしている。

Pactの動作フローは以下のような流れになっている。

1. コンシューマーは契約を記述する
2. コンシューマー側のテスト実行時にPactファイルを生成する
3. Pactファイルをpact_brokerにアップロードする
4. プロバイダーはアップロードされたPactファイルを使ってテストを実行する

cf.
- [docs.pact.io](https://docs.pact.io/)
- [github.com - pact-foundation](https://github.com/pact-foundation)
- [pactflow.io - How Pact contract testing works](https://pactflow.io/how-pact-works/?utm_source=ossdocs&utm_campaign=getting_started#slide-1)
- [dius.com.au - Simplifying Microservice testing with Pacts](https://dius.com.au/2014/05/20/simplifying-microservice-testing-with-pacts/)
- [pactflow.io - The curious case for the Provider Driven Contract](https://pactflow.io/blog/the-curious-case-for-the-provider-driven-contract/)

# 所感
組織がある程度の規模で、複数チームで開発しているとサービス間の整合性をどのように保つか？という課題が自然に発生すると思う。

契約テストはそのような課題に対して有効な手法であると思う。

特に、APIの仕様変更を自動で検知できるようになるのは大きなメリットだと思う。これはチームあるいはサービスが増えれば増えるほどレバレッジが効いてくるのではないかと思った。

そんなメリットを感じる契約テストではあるが、調べてみた感じではあまい流行しているような印象は受けなかった。

過去のデータであるが、Technology Radarでも積極的に採用するようなステータスではなかった。

契約テストの導入にあたっては、ツール導入が前提となり得るが、そのツールの中でもメジャーなのがPactのようであった。

Pactが流行っていないように見える理由としては、コンシューマー駆動契約テストしかサポートしていないことと、導入の旨味が大きく感じられない（E2Eなど他の代替のインテグレーションテストツールで間に合ってしまっている）ことが挙げられるのではないかと思った。実際に利用してみたわけではないので感覚でしかないし、海外では実はもっと普及しているのかもしれないが、少なくとも国内ではあまり事例は多くなかった。

Pactを拡張したPactFlowについてはまだ発展途中段階という印象を受けたが、契約テストのメリットや導入障壁を下げる良いツールになり得るのではないかと思った。

サービスがある程度の規模感になってくると複数のサービスから利用されるようなAPIを提供する基盤のようなシステムが登場することがあると思うが、コンシューマー駆動契約テストよりプロバイダー駆動契約テストの方が向いているように思われる。このようなケースの場合ではPactは最適解ではないかもしれない。
