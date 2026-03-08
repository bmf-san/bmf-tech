---
title: NewSQL関連の資料を読み漁った
slug: newsql-resources-review
date: 2023-03-29T00:00:00Z
author: bmf-san
categories:
  - データベース
tags:
  - NewSQL
  - リンク集
translation_key: newsql-resources-review
---


NewSQLについての色々と調べて読み漁ってみたので読んだものをまとめておく。

元々いくつかのNewSQLのDBについての比較をしてみたいと思って漁っていたのだが、内部で使われている技術について知っておく必要があったので、関連技術についての記事が多めになっている。

# 資料一覧
- [Hybrid Clock](https://martinfowler.com/articles/patterns-of-distributed-systems/hybrid-clock.html)
- [Hybrid Logical Clock (HLC)](https://sergeiturukin.com/2017/06/26/hybrid-logical-clocks.html)
- [本当は恐ろしい分散システムの話](https://www.slideshare.net/kumagi/ss-81368169)
- [分散処理システムの検証をサービスとして提供するJepsenに注目 | Think IT（シンクイット）](https://thinkit.co.jp/article/17532)
- [分散システムの課題](https://aws.amazon.com/jp/builders-library/challenges-with-distributed-systems/)
- [分散システムについて語らせてくれ](https://www.slideshare.net/kumagi/ss-78765920)
- [クラスタに3ノード必要な理由 - Qiita](https://qiita.com/ntoreg/items/ec6f1eca87ba5c5c0399)
- [雰囲気で分散システム使ってるやついる!? - Qiita](https://qiita.com/muson0110/items/2379595f9bc1d5720478)
- [Navigating the 8 fallacies of distributed computing](https://ably.com/blog/8-fallacies-of-distributed-computing)
- [分散コンピューティングの落とし穴 - Wikipedia](https://ja.wikipedia.org/wiki/%E5%88%86%E6%95%A3%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%81%AE%E8%90%BD%E3%81%A8%E3%81%97%E7%A9%B4)
- [Percolator vs Spanner. Implementing Distributed Transactions the Google Way | YugabyteDB](https://www.yugabyte.com/blog/implementing-distributed-transactions-the-google-way-percolator-vs-spanner/)
- [Low Latency Reads in Geo-Distributed SQL with Raft Leader Leases | Yugabyte](https://www.yugabyte.com/blog/low-latency-reads-in-geo-distributed-sql-with-raft-leader-leases/)
- [A Busy Developer’s Guide to Database Storage Engines — The Basics | Yugabyte](https://www.yugabyte.com/blog/a-busy-developers-guide-to-database-storage-engines-the-basics/)
- [分散合意アルゴリズム Raft を理解する - Qiita](https://qiita.com/torao@github/items/5e2c0b7b0ea59b475cce)
- [4 Data Sharding Strategies We Analyzed When Building YugabyteDB](https://www.yugabyte.com/blog/four-data-sharding-strategies-we-analyzed-in-building-a-distributed-sql-database/)
- [CockroachDB's consistency model](https://www.cockroachlabs.com/blog/consistency-model/)
- [Spanner vs. Calvin: Distributed Consistency at Scale](https://fauna.com/blog/distributed-consistency-at-scale-spanner-vs-calvin)
- [Fauna | The distributed serverless database](https://fauna.com/)
- [Fauna の紹介 （Temporal Database ） なぜ定番のSQLを使わずFQLという謎言語を使うのか？ - Qiita](https://qiita.com/masakinihirota/items/c232832bd72ae11905e7)
- [Living without atomic clocks: Where CockroachDB and Spanner diverge](https://www.cockroachlabs.com/blog/living-without-atomic-clocks/)
- [オープンソースの分散型NewSQLデータベースによるタイムサービス配信の仕組み - TiDB - PingCAP株式会社](https://pingcap.co.jp/how-an-open-source-distributed-newsql-database-delivers-time-services/)
- [The Vitess Docs | What Is Vitess](https://vitess.io/docs/17.0/overview/whatisvitess/)
- [金融トランザクションにNewSQLが使えるか検証してみた - 分散型データベース PingCAP株式会社](https://pingcap.co.jp/project/customer-story-tis/)
- [金融業界のミッションクリティカルなシナリオでTiDBを活用（パート1） - PingCAP株式会社](https://pingcap.co.jp/blog-using-tidb-in-mission-critical-scenarios-of-the-financial-industry-part-1/)
- [CAP定理の壁 ～NewSQLへの道～ | データベース アクセス パフォーマンス ブログ](https://www.climb.co.jp/blog_dbmoto/archives/5077)
- [分散ノードの整合性 ～NewSQLへの道～ | データベース アクセス パフォーマンス ブログ](https://www.climb.co.jp/blog_dbmoto/archives/5193)
- [LSM-TreeとRocksDB、TiDB、CockroachDBが気になる &middot; hnakamur&#39;s blog](https://hnakamur.github.io/blog/2016/06/20/lsm-tree-and-rocksdb/)
- [Amazon Aurora アーキテクチャ概要](https://pages.awscloud.com/rs/112-TZM-766/images/01_Amazon%20Aurora%20%E3%82%A2%E3%83%BC%E3%82%AD%E3%83%86%E3%82%AF%E3%83%81%E3%83%A3%E6%A6%82%E8%A6%81.pdf)
- [NewSQLのコンポーネント詳解 - Qiita](https://qiita.com/tzkoba/items/3e875e5a6ccd99af332f)
- [Amazon Auroraの先進性を誰も解説してくれないから解説する - Qiita](https://qiita.com/kumagi/items/67f9ac0fb4e6f70c056d)
- [TiDB on AWS EKS 〜DMM動画のPoCレポート〜 - DMM inside](https://inside.dmm.com/articles/tidb-on-aws-eks-poc-report/)
- [CockroachDB vs. TiDB vs. YugabyteDB Comparison](https://db-engines.com/en/system/CockroachDB%3BTiDB%3BYugabyteDB)
- [2020年現在のNewSQLについて - Qiita](https://qiita.com/tzkoba/items/5316c6eac66510233115)
- [Architecture Overview](https://www.cockroachlabs.com/docs/stable/architecture/overview.html)
- [Explore YugabyteDB | YugabyteDB Docs](https://docs.yugabyte.com/preview/explore/)
- [[論文紹介] TiDB:a Raft-based HTAP database](https://zenn.dev/tzkoba/articles/4e20ad7a514022)
- [TiDB の紹介](https://docs.pingcap.com/ja/tidb/stable/overview)
- [SQL・NoSQL・NewSQLの違いは？【データベース言語】](https://zenn.dev/umi_mori/books/331c0c9ef9e5f0/viewer/da6ec2#newsql%E3%81%A8%E3%81%AF%EF%BC%9F)
- [100億円キャンペーンで学んだ“教訓”　PayPayのスケーラブルな巨大決済システムを支える工夫 - ログミーTech](https://logmi.jp/tech/articles/321524)

# 所感
実際にDB選定する際は要件に合わせた検証や性能比較が必要と思われるが、比較する際のポイントがいくつか分かった。

- 採用しているSQLインタフェース
  - MySQLやPostgreSQLといったSQLへの互換性
- 分散トランザクションの仕様
  - ちょっと難しくて理解できていないが、RaftやPaxosといった分散合意プロトコルや分散クロックなど採用している技術や方針によって整合性をどういう形で担保するかといった仕様が異なる、はず
- computeとストレージが別になっているか、同居しているアーキテクチャ構成か
  - Spannerしか使ったことがないので、Spannerのように別になっているのが当たり前だと思っていたが、そうではないものもあるらしい
  - スケーラビリティの担保に影響するかと思う
- HTAPか否か
  - 分析用途も考慮するなら大事なポイントかと思われる
    - とはいってもTiDB以外でサポートしているNew SQLあったっけ・・？？

