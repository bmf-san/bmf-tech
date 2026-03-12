---
title: "BFF（Backend For Frontend）とは？メリットと実装ガイド"
description: 'BFF（Backend For Frontend）とは何か、なぜ必要か、フロントエンドとバックエンドの間に BFF レイヤーを置くことで API 契約を簡素化する実装方法を解説します。'
slug: bff-explained
date: 2023-08-29T00:00:00Z
author: bmf-san
categories:
  - アーキテクチャ
tags:
  - BFF
translation_key: bff-explained
---


# 概要
BFFについて調べたことをまとめる。

# BFFとは
Backends For Frontendsの略。Best Friends Forever（ズッ友だよ）ではない。
‌
名前の通り、フロントエンドのためのバックエンドサーバーのことで、フロントエンドのためのAPIやHTMLをレスポンスするなどUI・UXのための役割を担っている。
‌
クライアント（サーバーの呼び出し側）の多様性に応えるのが難しいという問題を、BFFはクライアントごとの要求を整理する形で解決することができる。
‌
# 気になったこと
- 採用言語
  - BFFはフロンドエンドのためのバックエンドである故フロントエンド寄りの技術で構成するケースが多いように見受けられる
- 再構成
  - 一度BFFにするとバラすのは大変そう
  - 本当にBFFが必要となるまでは採用するのは先延ばしするのが良さそう（本当に必要か？という判断が難しいわけではあるが・・）
- アンチパターンとして考えられそうなこと
  - バックエンドエンジニアとフロントエンドエンジニア間のコミュニケーション不足
  - BFFにUI以外のロジックが多く乗ってしまう
  - バックエンドとフロントエンドの結合を一気に行ってしまうビックバンジョイント
- フロントエンドの最適化がしやすそう
  - APIの呼び出しを最適化することでUIの表示パフォーマンスを改善したりできそう
- BFFとDDD
  - フロントエンド側でのドメイン整理が必要？このへんはわからない..
- APIの集約単位
  - どのAPIをどういう単位でグルーピングするか考える難しさがありそう
  - マイクロサービスをやっているのなら、BFFではなく別のマイクロサービスを立てるでも良かったのでは？とかなるとBFFのメリットが損なわれそう
- マイクロフロントエンドとの相性
  - マイクロフロントエンドの知見がなくてﾅﾆﾓﾜｶﾗﾅｲ
  - マイクロフロントエンドのコンポーネント構成に影響されそう？
- GraphQLとの相性が良さそう
  - graphQLを使うならスキーマファーストではなく、コードファーストが向いているという事例
    - cf. [なぜGraphQLをコードファーストに統一したのか？型定義の一貫性を保つためのBFF／FE大整理](https://logmi.jp/tech/articles/326592))
- 可用性
  - BFFが複数のバックエンドの集約であるということは、複数のバックエンドの障害に影響を受ける、依存しているということであると思う
  - この懸念について、ZOZOさんでは正常に返却できるデータだけはレスポンスをするように工夫しているらしい
    - cf. [Backends For Frontends（BFF）はじめました](https://techblog.zozo.com/entry/zozo-aggregation-api-bff)
- キャッシュ
  - BFF側でのキャッシュも考慮する必要がありそう
- タイムアウト・リトライ制御
  - 通常のAPIでも考える事項ではあるが、BFFの場合は設定値の調整が少し頭を悩ませそう
- デプロイ
  - BFFのデプロイとバックエンドのデプロイの足並みを調整する必要がありそう
  - デリバリのスピードに関わる
- ビジネスロジックを持たない
  - BFFではビジネスロジックを持たないのが基本のようだが、ビジネスロジックを完全に切り離すことはできる？できないケースもありそう

# 参考
- [Pattern: Backends For Frontends](https://samnewman.io/patterns/architectural/bff/)
- [BFF（Backends For Frontends）超入門――Netflix、Twitter、リクルートテクノロジーズが採用する理由](https://atmarkit.itmedia.co.jp/ait/articles/1803/12/news012.html)
- [BFF(Backend for Frontend)](https://speakerdeck.com/dena_tech/bff-backend-for-frontend)
- [BFFに取り組む開発者たちが語る「UIT#3 The “Backends for Frontends” sharing」](https://engineering.linecorp.com/ja/blog/bff-talk-uit-3-backends-for-front-ends-sharing)
- [BFF/SSRの話](https://speakerdeck.com/yosuke_furukawa/ssrfalsehua?slide=2)
- [セッションレポート「Backends for Frontendsこそサーバーレスで楽をしよう!!～本来の目的に集中するために～」を見てBFFについて学習したことまとめ](https://tsd.mitsue.co.jp/blog/2021-10-28-event-report-aws-innovate-modern-app-edition-bff/)
- [Dividing frontend from backend is an antipattern](https://www.thoughtworks.com/insights/blog/dividing-frontend-backend-antipattern)
- [BFFとmicroservicesアーキテクチャ](https://zenn.dev/hirac/articles/7bd857ab904d66)
- [GraphQLを活用したBackend for Frontendへ リアーキテクチャした話](https://www.docswell.com/s/hireroo/Z6YJR4-techtalk2-1#p1)
- [BFF（Backends For Frontends）実践における3つのアンチパターンと、その回避策](https://atmarkit.itmedia.co.jp/ait/articles/1808/31/news013.html)
- [BFF（Backends For Frontends）の5つの便利なユースケース](https://atmarkit.itmedia.co.jp/ait/articles/1805/18/news022.html)
- [Backends for Frontends (BFF) 再考](https://zenn.dev/morinokami/scraps/20a4eab9415a50)
- [More coverage on BFFs](https://samnewman.io/blog/2016/02/14/more-coverage-on-bffs/)
- [Embracing the Differences : Inside the Netflix API Redesign](https://netflixtechblog.com/embracing-the-differences-inside-the-netflix-api-redesign-15fd8b3dc49d)
- [BFF \\@ SoundCloud](https://www.thoughtworks.com/insights/blog/bff-soundcloud)
- [Moving to Microservices at SoundCloud with Lukasz Plotnicki](https://softwareengineeringdaily.com/2016/02/04/moving-to-microservices-at-soundcloud-with-lukasz-plotnicki/)
- [Backends For Frontends（BFF）はじめました](https://techblog.zozo.com/entry/zozo-aggregation-api-bff)
- [Backends for Frontends パターン](https://aws.amazon.com/jp/blogs/news/backends-for-frontends-pattern/)
- [BFFとはなんなのか？](https://qiita.com/souhei-etou/items/d5de99bb8cba1c59d393)
- [新規開発においてBFF(Backend for Frontend) を採用すべきか](https://vivit.hatenablog.com/entry/2021/11/10/101530)
- [Backend for frontend (BFF) pattern— why do you need to know it?](https://medium.com/mobilepeople/backend-for-frontend-pattern-why-you-need-to-know-it-46f94ce420b0)
- [流行りのBFFアーキテクチャとは？｜Offers Tech Blog](https://zenn.dev/overflow_offers/articles/20220418-what-is-bff-architecture)
- [なぜGraphQLをコードファーストに統一したのか？型定義の一貫性を保つためのBFF／FE大整理](https://logmi.jp/tech/articles/326592)
- [Are BFF (Backend for Frontend) and DDD mutually exclusive?](https://stackoverflow.com/questions/76940683/are-bff-backend-for-frontend-and-ddd-mutually-exclusive)
- [フロントエンドエンジニアは Micro Frontends の夢を見るか](https://engineering.mercari.com/blog/entry/2018-12-06-162827/)

# 所感
BFF自体は知っていたのでさらっとググって終わろうと思っていたのだが、アーキテクチャの可用性や、ビジネスロジックの扱い、クライアントの適切な集約、組織構成との関連など色々考えるポイントが多く面白かった。

自分としてはBFFは結構慎重にならないと落とし穴が多そうという印象を持った。罠みたいなところは見えるけどそれに引っかからないようにうまく作るのは難しそうという感覚を持った。

もしBFFを検討する機会があれば振り返ってみようと思う。

