---
title: Aurora Serverlessについて調べた
description: Aurora Serverlessについて調べたについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: aurora-serverless-research
date: 2023-12-27T00:00:00Z
author: bmf-san
categories:
  - データベース
tags:
  - Amazon Web Service
translation_key: aurora-serverless-research
---


Aurora Serverlessについて興味を持ったので記事を漁ってみた。

# チェックした記事
- [aws.amazon.com - Amazon Aurora Serverless](https://aws.amazon.com/jp/rds/aurora/serverless/)
- [techblog.zozo.com - Aurora Serverless v2を本番導入した話 〜検討や導入時のポイント・得られた効果について〜](https://techblog.zozo.com/entry/aurora-serverless-v2)
- [www.plan-b.co.jp - Aurora Serverlessを実際に使ってみたメリットとデメリット](https://www.plan-b.co.jp/blog/tech/28232/)
- [www.stylez.co.jp - AWS Aurora Serverlessを最新機能(v2)を含めて解説](https://www.stylez.co.jp/columns/aws_aurora_serverless_including_the_latest_features_v2/)
- [dev.classmethod.jp - Aurora Serverlessについての整理](https://dev.classmethod.jp/articles/aurora-serverless-summary/)
- [qiita.com - 待たせたな！噂のAuroraサーバーレスv2がGA。初心者にも分かりやすくまとめてみた](https://qiita.com/minorun365/items/2a548f6138b6869de51a)
- [engineering.dena.com - 新サービス Aurora Serverless v2 の検証とその評価 \[DeNA インフラ SRE\]](https://engineering.dena.com/blog/2022/06/aurora-serverless-v2/)
- [logmi.jp - Aurora Serverless v2とAurora Provisionedのコスト比較 3つの変数を“羅針盤”に優位境界面を探る](https://logmi.jp/tech/articles/328375)
- [zenn.dev - Aurora Serverless v2を運用して気づいたこと](https://zenn.dev/yumemi_inc/articles/7ad423906ec202)

# Aurora Serverless v2の特徴
ざっくり調べた感じ次のような特徴があると分かった。

- マルチAZ対応
- オートスケール対応（スケールアウトではなく、スケールアップ・ダウン）
- VPC外からのアクセスも可能（public ipの付与が可能）
- Aurora Provisionedと比較するとコストは高め
- 特定の時間帯にスパイクが予測されるようなケースがユースケースとしては合いそう
