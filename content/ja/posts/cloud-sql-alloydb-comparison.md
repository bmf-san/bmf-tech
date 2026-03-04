---
title: "Cloud SQLとAlloyDBの比較"
slug: "cloud-sql-alloydb-comparison"
date: 2024-11-17
author: bmf-san
categories:
  - "データベース"
tags:
  - "AlloyDB"
  - "Cloud SQL"
  - "Google Cloud Platform"
draft: false
---

Cloud SQLと比較したAlloyDBの優位性について知りたかったので簡単に調査してみた。

# 前提
Cloud SQLについてはCloud SQL Enterprise editionを対象とする。

# Cloud SQL vs AlloyDB
仕様について単純比較する。

|                        | Cloud SQL                                  | AlloyDB                                                     |
| ---------------------- | ------------------------------------------ | ------------------------------------------------------------ |
| サービス形態           | マネージド型RDBMS                          | PostgreSQLベースの分散型に近いDB                             |
| 用途                   | OLTP                                       | OLTPとOLAP                                                   |
| 構成単位               | 単一インスタンス                           | クラスタ                                                     |
| スケーリング           | 垂直・水平（制限有り）                     | 垂直・水平（制限有り）                                       |
| 可用性                 | SLA 99.99%                                 | SLA 99.95%                                                   |
| メンテナンス           | ダウンタイム有り                           | 無停止                                                       |
| バックアップとリストア | オンデマンドバックアップ・自動バックアップ | オンデマンドバックアップ・自動バックアップ・継続バックアップ |
| 互換性                 | MySQL/PostgreSQL/SQL Server                | PostgreSQL                                                   |
| コスト                 | 単純比較では全体的にAlloyDBより安い       | 単純比較では全体的にCloud SQLより高い                        |

# AlloyDBの特徴
AlloyDBの主な特徴についてCloud SQLと比較しながらまとめる。

## 用途
Cloud SQLは基本的にOLTP向けだが、**AlloyDBはHTAPに対応**しており、OLTPとOLAPの両方の用途で利用できる。

OLAPはデータベースフラグ（`google_columnar_engine.enabled`）を設定すると有効になる。

OLAPで利用できる型には制限があるが、一般的な型は網羅されている。

## 構成単位
Cloud SQLは単一インスタンスをベースとしており、コンピューティングやストレージの性能が単一インスタンスの限界に制限される。

一方AlloyDBは**コンピューティングとストレージを分離した構成となっているため、高いスケーラビリティを備えた構成**となっている。

## スケーリング
Cloud SQLもAlloyDBもスペックアップによる垂直スケーリング、リードレプリカによる水平スケーリングも可能である。

どちらも共通して書き込みについては水平スケーリングができない。

大きな違いとしては、**AlloyDBはリードレプリカによるレプリケーション遅延が構成上低減される**仕組みとなっている。

## バックアップ・リストア
Cloud SQLとAlloyDBのバックアップの違いとしては、AlloyDBには、**継続的バックアップ**がある点となる。

> 継続的なバックアップとリカバリは、すべてのクラスターでデフォルトで有効になっており、同じプロジェクトとリージョン内の別のクラスターの最新の状態に基づいて新しいクラスターを作成できる AlloyDB の機能です。

> AlloyDB を使用すると、既存のクラスターをマイクロ秒単位の精度で最近の履歴の任意の時点に復元できます。デフォルトでは、AlloyDB では過去 14 日間までの任意の時点を選択できます。クラスターを構成して、このウィンドウを最長 35 日間、または最短 1 日間に変更できます。

他のバックアップとの違いがあまり理解できなかったが、おそらくマイクロ秒単位まで細かく指定してポイントインタイムリカバリが迅速に行える、という点が特徴なのかもしれない。

## ネットワーク
Cloud SQLのCloud SQL Auth Proxyのように、AlloyDB Auth Proxyが提供されている。

DBに接続するクライアント側のローカルにインストールするプロキシソフトウェアで、IAMベースの接続認証や暗号化通信（TLS）を利用することができる。

このプロキシの利用は必須ではないが、推奨されている。

## データのインポート・エクスポート
AlloyDBでは、CSV・DMP（PostgreSQLのダンプファイル）・SQLによるデータのインポート・エクスポートに対応している。

ファイル形式以外では、Database Migration Service（DMS）を利用することで、他のデータベースからAlloyDBへのデータ移行が可能となる。

一部制約があるが、Cloud SQLからAlloyDBへの移行も可能である。

cf. [cloud.google.com - Migrate from Cloud SQL for PostgreSQL to AlloyDB for PostgreSQL](https://cloud.google.com/alloydb/docs/migrate-cloud-sql-to-alloydb#required-roles)

## チューニング
AlloyDBにはadaptive autovacuumとindex advisorという機能が用意されている。

adaptive autovacuumは、PostgreSQLのautovacuumの運用を改善する機能（オートパイロット機能。マネージドautovacuumといったところ。）で、DBのパフォーマンスや可用性を改善するために、自動的に適切なタイミングでautovacuumを実行する。

index advisorは、クエリの実行計画を分析し、インデックスの提案を行う機能である。

# 所感
- Cloud SQLの上位互換と捉えて良さそう
- Readは無限にスケールする。レプリケーション遅延を気にする必要はなさそう
- 既にCloud SQLを利用している場合、Cloud SQL Enterprise edition Plusでも性能改善が期待できる可能性もあり、移行対象候補の一つとして考えても良いかもしれない

# 参考
- [cloud.google.com - AlloyDB overview](https://cloud.google.com/alloydb/docs/overview)
- [cloud.google.com - AlloyDB for PostgreSQL の仕組み: データベース対応のインテリジェントなストレージ](https://cloud.google.com/blog/ja/products/databases/alloydb-for-postgresql-intelligent-scalable-storage)
- [cloud.google.com - AlloyDB for PostgreSQL の仕組み: 適応型自動バキューム](https://cloud.google.com/blog/ja/products/databases/alloydb-for-postgresql-under-the-hood-adaptive-autovacuum)
- [blog.g-gen.co.jp - AlloyDB for PostgreSQLを徹底解説！](https://blog.g-gen.co.jp/entry/alloydb-for-postgresql-explained)
- [【Google Cloud】AlloyDB と Cloud SQL を徹底比較してみた！！(第1回：AlloyDB の概要、性能検証編)
](https://sight-r.sts-inc.co.jp/google_cloud_article/google-cloud-compare-alloydb-1/)
- [cloud-ace.jp - これを見れば分かる！Google Cloud データベース選定](https://cloud-ace.jp/column/detail469/)
- [medium.com - AlloyDB Adaptive AutoVacuum and how AlloyDB Cluster Storage Space is Released.](https://medium.com/google-cloud/alloydb-adaptive-autovacuum-and-how-alloydb-cluster-storage-space-is-released-41be54b8b8c8)
