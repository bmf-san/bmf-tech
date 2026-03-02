---
title: "Sagaパターンについて調べたことをメモ"
slug: "saga"
date: 2023-09-17
author: bmf-san
categories:
  - "アーキテクチャ"
tags:
  - "マイクロサービス"
  - "Sagaパターン"
  - "TCCパターン"
  - "分散トランザクション"
  - "2phase commit"
draft: false
---

# 概要
Sagaパターンについて調べたことをメモ。

# Sagaパターンとは
- マイクロサービスにおいては、分散トランザクション（2phase commitなど）は推奨されていない
  - 分散トランザクションを避け、整合性を担保するパターンとしてSagaパターンがある
- 長時間のロックを避け、結果整合性を利用する
- 補償トランザクション
  - 一連のトランザクションの取り消しを行う操作
  - Sagaパターンでは単純なロールバックを禁止している
- マイクロサービス特有のパターンではなく、SOAでも使われていた
- 実装パターン
  - Choreography（コレオグラフィ）　
    - 各サービスが自身の責務を持ってトランザクションを進めていく
	- SPOFを避けれるが、全体のテストがしづらくなる
  - Orchestration（オーケストレーション）とは、
    - 中央集権的なサービスを用意し、各サービスに指示を出してトランザクションを進めていく
    - SPOFになりやすいが、全体のテストがしやすくなる
- [ソフトウェアアーキテクチャ・ハードパーツ](https://amzn.to/3RmavPp)では8つの種類のSagaパターンが紹介されている

# その他のパターン
マイクロサービスの整合性を保つ別のパターンとして、Sagaパターン同じく結果整合性を利用するTCC（Try-Confirm/Cancel）パターンというものもある。

2phase commitに似ているが、TCCパターンでは、各サービスがトランザクションの準備、確認、キャンセルの3つのステップを持つ。

TCCパターンは補償トランザクションのようなロールバックは行わず、不整合の生じる処理を行わないことにすることで整合性を担保する。

# 参考
- [www.cs.cornell.edu - SAGAS](https://www.cs.cornell.edu/andru/cs711/2002fa/reading/sagas.pdf)
- [microservices.io - Pattern: Saga](https://microservices.io/patterns/data/saga.html)
- [qiita.com - マイクロサービスに移行した際の分散トランザクションの危険性](https://qiita.com/yoshii0110/items/3c86173dc53d93588b72)
- [qiita.com - マイクロサービスのトランザクション管理をデザインする（事前知識編）](https://qiita.com/Yoyo-kikuchi/items/c113aeeab3bf2daa0910#tcctry-confirmcancel%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3)
- [zenn.dev - awsで実現するsagaパターン(コレオグラフィ編)](https://zenn.dev/yoshii0110/articles/57ccc582f7fcd3)
- [learn.microsoft.com - saga 分散トランザクション パターン](https://learn.microsoft.com/ja-jp/azure/architecture/reference-architectures/saga/saga)
- [docs.aws.amazon.com - Sagaパターン](https://docs.aws.amazon.com/ja_jp/prescriptive-guidance/latest/modernization-data-persistence/saga-pattern.html)
- [wakatchi.dev - TCCパターンとSagaパターンでマイクロサービスのトランザクションをまとめてみた](https://wakatchi.dev/microservices-tx-pattern-saga-tcc/)
- [eikatou.net - 分散トランザクション Sagaパターンを学ぶ](https://eikatou.net/blog/20201205.html)
- [techblog.raksul.com - マイクロサービスにSagaパターンを用いて検証を行った](https://techblog.raksul.com/entry/2022/09/22/095433)
- [www.techscore.com - マイクロサービスにおけるトランザクション【Saga】](https://www.techscore.com/blog/2018/12/05/saga-in-microservices/)
- [www.12-technology.com - \[SAGAパターン\]コレオグラフィとオーケストレーションのメリット・デメリット
](https://www.12-technology.com/2021/08/dbsaga.html)
- [cloud.google.com - WorkflowsでSagaパターンを実装する](https://cloud.google.com/blog/ja/topics/developers-practitioners/implementing-saga-pattern-workflows)
- [medium.com - GCP でマイクロサービス Saga パターン編](https://medium.com/google-cloud-jp/gcp-saga-microservice-7c03a16a7f9d)
- [www.oracle.com - Microservicesアーキテクチャのトランザクション管理の現実](https://www.oracle.com/a/otn/docs/jp-dev-days-microservices.pdf)
- [speakerdeck.com - 気がついたらSagaパターンになっていた！？ 少人数で運用するサーバレスバックエンド](https://speakerdeck.com/miu_crescent/qi-gatuitarasagapatanninatuteita-shao-ren-shu-deyun-yong-surusabaresubatukuendo)
- [engineering.mercari.com - メルコイン決済基盤における分散トランザクション管理](https://engineering.mercari.com/blog/entry/20230614-distributed-transaction/)

