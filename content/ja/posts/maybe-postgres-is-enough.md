---
title: "Maybe Postgres Is Enough：別のDBを足す前にPostgresで足りるか考える"
slug: maybe-postgres-is-enough
date: 2026-06-29
author: bmf-san
categories:
  - データベース
tags:
  - PostgreSQL
  - アーキテクチャ
description: "キャッシュや全文検索のために別のミドルウェアを足す前に、Postgresで足りないか確認したい。postgresisenough.devを参考に、用途ごとの代替機能を一覧で簡潔にまとめる。"
translation_key: maybe-postgres-is-enough
draft: false
---

# 概要
キャッシュにRedis、全文検索にElasticsearch、ジョブキューにSidekiq、と用途ごとに別のミドルウェアを足していくと、気づけばアプリは多数のデータストアに依存している。それぞれにデプロイ・バックアップ・監視・障害対応が必要になり、運用コストは膨らむ。

[postgresisenough.dev](https://postgresisenough.dev/)は「別のDBを足す前にPostgresで足りるか確認しよう」という考え方を提示している。本記事では、その用途ごとの対応を簡潔にまとめる。

# 用途ごとの対応一覧
やりたいことに対して別ミドルウェアへ飛びつく前に、まずPostgresの機能で代替できないか検討するとよい。

| やりたいこと | よく足すもの | Postgresの機能 |
| --- | --- | --- |
| キャッシュ | Redis、Memcached | `UNLOGGED`テーブル、マテリアライズドビュー |
| ジョブキュー | Redis+Sidekiq、RabbitMQ | `SKIP LOCKED`、pgmq、pgflow |
| 全文検索 | Elasticsearch、Algolia | tsvector、pg_trgm、ParadeDB |
| ドキュメントストア | MongoDB、CouchDB | JSONB、FerretDB |
| ベクトル検索/AI | Pinecone、Weaviate | pgvector、pgvectorscale |
| 時系列データ | InfluxDB、TimescaleDB | TimescaleDB、pg_partman |
| 分析/OLAP | Snowflake、BigQuery | pg_analytics、DuckDB連携 |
| グラフDB | Neo4j、Neptune | Apache AGE、再帰CTE |
| 地理空間 | 専用GISシステム | PostGIS |

# 各機能の概要
表のPostgres機能を、用途ごとに簡単に解説する。

## キャッシュ
`UNLOGGED`テーブルはWAL(先行書き込みログ)を省くため書き込みが速い。クラッシュで中身が消える代わりに、揮発キャッシュとして使える。マテリアライズドビューは重い集計結果を保存し、`REFRESH`で更新しながら高速に読み出せる。

## ジョブキュー
`SELECT ... FOR UPDATE SKIP LOCKED`はロック中の行を飛ばして取得する。複数のワーカーが競合せずに仕事を取り出せるため、テーブルを簡易キューにできる。pgmqはキュー、pgflowはワークフローを提供する拡張である。

## 全文検索
tsvectorは文書を検索用に正規化し、転置インデックスで高速に探す。pg_trgmはトライグラムで部分一致やあいまい検索を助ける。ParadeDBはBM25スコアの本格的な検索を載せる拡張である。

## ドキュメントストア
JSONBはJSONをバイナリで格納し、インデックスを張れる。スキーマレスなデータをそのまま扱える。FerretDBはMongoDB互換のAPIをPostgresの上で提供する。

## ベクトル検索/AI
pgvectorはベクトル型と近傍検索を追加し、埋め込みの類似検索を可能にする。pgvectorscaleは索引を強化し、大規模なベクトルでも速度を保つ。

## 時系列データ
TimescaleDBは時系列向けの拡張で、テーブルを自動分割して書き込みと集計を速くする。pg_partmanは時間や範囲でパーティションを自動管理する。

## 分析/OLAP
pg_analyticsは列指向で分析クエリを速くする。DuckDB連携を使えば、集計やファイルの直読みをPostgresから扱える。

## グラフDB
Apache AGEはopenCypherでグラフを問い合わせる拡張である。再帰CTE(`WITH RECURSIVE`)を使えば、拡張なしでも階層やグラフを辿れる。

## 地理空間
PostGISは地理空間型と関数、空間インデックスを追加する。距離や交差の計算を標準SQLで書ける。

# まとめ
ストアを1つ増やすたびに、デプロイ・監視・障害対応の面が増える。まずはPostgresで足りるかを確認する。これだけで構成はかなり単純になる。

# 参考
- [Postgres Is Enough](https://postgresisenough.dev/)
- [Tools 一覧](https://postgresisenough.dev/tools)
