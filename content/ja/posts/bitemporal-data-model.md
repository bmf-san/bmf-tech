---
title: Bitemporal Data Modelについて
description: "バイテンポラルデータモデルを解説。トランザクション時間と有効時間の2軸管理、履歴追跡、監査要件対応、PostgreSQL実装のEXCLUDE制約での一貫性保証を紹介します。"
slug: bitemporal-data-model
date: 2024-05-27T00:00:00Z
author: bmf-san
categories:
  - データベース
tags:
  - Bi-Temporal
  - Uni-Temporal
  - Non-Temporal
  - 履歴
  - Temporal Data Model
translation_key: bitemporal-data-model
---


# 概要
Bitemporal Data Modelについて調べたことをまとめる。

# Bitemporal Data Modelとは
Temporal Data Modelsという分野で扱われるデータモデルの１つで、時間軸を取り扱うデータモデルのうち、2つの時間軸を持つデータモデルのことを指す。

- ノンテンポラル（Non-Temporal）
  - 現在の状態のみを保持する
  - 過去の履歴や将来の変更に関する情報は保持しない
- ユニテンポラル（Uni-Temporal）
  - 単一の時間軸を持つ
  - データが有効である開始日と有効期間が終了する終了日を保持する
- バイテンポラル（Bi-Temporal）
  - 2つの時間軸を持つ
  - データベースに記録された時間であるトランザクション時間（システム時間）と、イベントが発生した時間である有効時間を保持する
  - ユニテンポラルは、トランザクション時間と有効時間が同じであるのに対し、バイテンポラルは異なる時間軸を持つ

バイテンポラルデータモデルが採用される背景としては次のような要件が考えられる。

- データの変更履歴の追跡
- 法規制や監査要件
- 時系列分析の柔軟性向上

バイテンポラルなデータモデルをRDBで扱う場合、次のような難しさが伴う。

- データの一貫性の保証
  - 同一のエンティティに対して、有効時間が重複しないように制約を持たせる必要がある
  - トランザクション時間の連続性を保つ必要がある
  - PostgreSQLでは、排他制約（EXCLUDE制約）が使えるため、有用である
- クエリの複雑さ
  - 有効時間とトランザクション時間の両方を考慮したクエリを発行する必要がある
    - 複雑化しやすいためパフォーマンスチューニングも難しくなる
- アプリケーションロジックの複雑さ
  - データの取得、更新、削除の際に意識する時間軸が2つある
  - ロジックに一貫性を持たせるように実装する必要がある

過去の履歴を参照したり、過去や未来の履歴を追加したり、履歴自体の更新情報を残したり、履歴情報の柔軟性を持つことができるデータモデルであるが、実現するためには複雑さを覚悟する必要がある。（と感じた・・）

# 所感
今まで業務で柔軟性の高い履歴データを扱ったことがなかったので、履歴の扱いの世界に深いものを感じた。

# 参考
- [en.wikipedia.org - Bitemporal modeling](https://en.wikipedia.org/wiki/Bitemporal_modeling)
- [martinfowler.com - Bitemporal History](https://martinfowler.com/articles/bitemporal-history.html)
- [martinfowler.com - Temporal Patterns](https://martinfowler.com/eaaDev/timeNarrative.html)
- [wiki.postgresql.org - Temporal Data & Time Travel in PostgreSQL](https://wiki.postgresql.org/images/6/64/Fosdem20150130PostgresqlTemporal.pdf)
- [www.progress.com - バイテンポラルで何ができるようになるのか](https://www.progress.com/docs/default-source/marklogic-docs/Bitemporal-Whitepaper-JP.pdf)
- [matsu-chara.hatenablog.com - BiTemporal Data Model導入時の注意点](https://matsu-chara.hatenablog.com/entry/2022/06/25/110000)
- [matsu-chara.hatenablog.com - だいたいよくわからないブログ](https://matsu-chara.hatenablog.com/entry/2017/04/01/110000)
- [www.slideshare.net - データ履歴管理のためのテンポラルデータモデルとReladomoの紹介 #jjug_ccc #ccc_g3](https://www.slideshare.net/slideshow/jjug-ccc-2017-spring-bitemporal-data-modeling-and-reladomo/76145041)
- [tech.smarthr.jp - ActiveRecord::Bitemporalとの歩み方](https://tech.smarthr.jp/entry/2023/10/04/110000)
- [speakerdeck.com - 操作履歴/時点指定アクセスの実現 - BiTemporal Data Model の実践](https://speakerdeck.com/f440/implementing-command-history-and-temporal-access)
- [speakerdeck.com - SmartHRにおけるBiTemporal Data Modelの実践のその後](https://speakerdeck.com/wata727/after-the-practice-of-bitemporal-data-model-in-smarthr)
- [zenn.dev - 業務から見たテンポラルデータモデルの解釈と利用方法の紹介](https://zenn.dev/zahn/articles/6a3d2138e9fe68)
- [ontact-rajeshvinayagam.medium.com - Bi-Temporal Data Modeling: An Overview](https://contact-rajeshvinayagam.medium.com/bi-temporal-data-modeling-an-overview-cbba335d1947)
