---
title: "CQRSとは？コマンドクエリ責務分離を使うべき場面"
slug: cqrs-research-notes
date: 2023-09-18T00:00:00Z
author: bmf-san
categories:
  - アーキテクチャ
tags:
  - CQRS
  - マイクロサービス
translation_key: cqrs-research-notes
---


# 概要
CQRSについて調べたことをメモ。

# CQRSとは
- Command Query Responsibility Segregation
  - CommandとQueryに責務分離するパターン
    - Command
       - データの更新を行う
       - データではなくタスクにフォーカスしたコマンドを設計
	   - DDDのドメインモデルを内包することが想定されている
    - Query
       - データの参照を行う
       - ユースケースごとに最適化されたオブジェクトとして設計
         - DTO
- 実装方法
  - イベントソーシングパターン
    - CQRSでは、コマンドとクエリの連携のためにイベントが用いられることが一般的
  - [www.ibm.com - Command Query Responsibility Segregation (CQRS) pattern](https://www.ibm.com/cloud/architecture/architectures/event-driven-cqrs-pattern/)
    - 段階的にCQRSへ移行していく話が語られている
- メリット
  - ReadとWriteをそれぞれスケーリングすることができる
  - ReadとWriteのそれぞれでデータソースを選択することができる
    - データソースを分けずに同一にすることも可能
    - データソースを分ける場合は同期または非同期でデータを同期する必要がある
  - アプリケーション側のモデルはシンプルになる
- デメリット
  - システムの構成要素が増えるためコストがかかる
  - ネットワーク全体の構成が複雑になる

# 所感
このパターンを導入すると決めるのは結構ハードルの高さを感じる。。。

海外では事例が多そうだったが、日本ではまだまだ事例が少ないというのが分かった。

# 参考
- [microservices.io - Pattern: Command Query Responsibility Segregation (CQRS)](https://microservices.io/patterns/data/cqrs.html)
- [learn.microsoft.com - CQRS パターン](https://learn.microsoft.com/ja-jp/azure/architecture/patterns/cqrs)
- [learn.microsoft.com - イベントソーシングパターン](https://learn.microsoft.com/ja-jp/azure/architecture/patterns/event-sourcing)
- [zenn.dev - CQRSを完全に理解した](https://zenn.dev/shmi593/articles/c1baeb2d453929)
- [martinfowler.com - 「イベントドリブン」とはどういう意味ですか?](https://martinfowler.com/articles/201701-event-driven.html)
- [martinfowler.com - CQRS](https://martinfowler.com/bliki/CQRS.html)
- [docs.aws.amazon.com - CQRS パターン](https://docs.aws.amazon.com/ja_jp/prescriptive-guidance/latest/modernization-data-persistence/cqrs-pattern.html)
- [logmi.jp - Chatworkのテックリードが語る、CQRSを上手に使うため方法](https://logmi.jp/tech/articles/322972)
- [logmi.jp - 「CQRSをやる」は「Event Sourcingをやる」とほぼ同義 リアクティブシステムとCQRSを反映した新アーキテクチャの設計思想](https://logmi.jp/tech/articles/324798)
- [little-hands.hatenablog.com - CQRS実践入門 [ドメイン駆動設計]](https://little-hands.hatenablog.com/entry/2019/12/02/cqrs)
- [hireoo.io -CQRSパターンを用いて統計データを処理するマイクロサービスを作成したお話](https://hireroo.io/journal/tech/statistics-service-using-cqrs)
- [speakerdeck.com - ざっくりCQRS/Event Sourcingを解説する](https://speakerdeck.com/j5ik2o/event-sourcingwojie-shuo-suru)
- [blog.j5ik2o.me - CQRSはなぜEvent Sourcingになってしまうのか](https://blog.j5ik2o.me/entry/2020/09/18/172612)
- [note.com - CQRSはEvent Sourcingなしで実現できるのか？](https://note.com/j5ik2o/n/n20aadb440a9b)
- [appmaster.io - マイクロサービスにおけるCQRSとイベントソーシングの適用](https://appmaster.io/ja/blog/cqrs-ibentososhingu-maikurosabisu)
- [www.ibm.com - Command Query Responsibility Segregation (CQRS) pattern](https://www.ibm.com/cloud/architecture/architectures/event-driven-cqrs-pattern/)
- [www.eventstore.com - CQRS](https://www.eventstore.com/cqrs-pattern)
- [blog.risingstack.com - When should you use CQRS?](https://blog.risingstack.com/when-to-use-cqrs)
- [logmi.jp - “自律性”は独立で機能しないと意味がない コマンド・クエリの要件から考えるCQRSの利点と欠点](https://logmi.jp/tech/articles/324797)
- [pages.awscloud.com - CQRS & Event Sourcing モダンアーキテクチャにおける役割と実装](https://pages.awscloud.com/rs/112-TZM-766/images/DevAx_connect_jp_season1_day4_CQRS%26EventSourcing.pdf)
