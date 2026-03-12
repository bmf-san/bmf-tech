---
title: PostgreSQL 論理レプリケーションの仕様まとめ
description: PostgreSQL 論理レプリケーションの仕様まとめについて調査・整理したメモ。基本概念と重要ポイントを解説します。
slug: postgresql-logical-replication-summary
date: 2025-05-23T00:00:00Z
author: bmf-san
categories:
  - データベース
tags:
  - PostgreSQL
translation_key: postgresql-logical-replication-summary
---


# 概要

PostgreSQLにおける論理レプリケーションは、データベース内の特定のテーブルに対するDML操作（INSERT、UPDATE、DELETEなど）を、他のPostgreSQLインスタンスに複製する仕組みである。物理レプリケーションとは異なり、論理レプリケーションはテーブル単位で柔軟に対象を選定できるため、データ統合、分散処理、マイグレーションといった用途に適している。

> 参考: [PostgreSQL Logical Replication Documentation - Overview](https://www.postgresql.org/docs/current/logical-replication.html)

# ハンズオン
## 環境構築
```yaml
# docker-compose.yml
version: '3.8'
services:
  publisher:
    container_name: publisher
    image: postgres:17
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: demo
    ports:
      - "5432:5432"
    volumes:
      - publisher_data:/var/lib/postgresql/data
    command: >
      postgres -c wal_level=logical
               -c max_replication_slots=4
               -c max_wal_senders=4
               -c hot_standby=off

  subscriber:
    container_name: subscriber
    image: postgres:17
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: demo
    ports:
      - "5433:5432"
    volumes:
      - subscriber_data:/var/lib/postgresql/data

volumes:
  publisher_data:
  subscriber_data:
```

`docker-compose up -d`で起動する。

# 動作検証
## Publisher側の設定
`docker exec -it publisher psql -U postgres -d demo`でpublisherに接続する。

以下のSQLでテーブルとデータを作成する。

```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT
);

INSERT INTO users (name) VALUES ('Alice'), ('Bob');
```

次のSQLでパブリケーションを作成する。

```sql
CREATE PUBLICATION my_pub FOR TABLE users;
```

## Subscriber側の設定
`docker exec -it subscriber psql -U postgres -d demo`でsubscriberに接続する。

次のSQLでテーブルを作成する。

```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT
);
```

次のSQLでサブスクリプションを作成する。

```sql
CREATE SUBSCRIPTION my_sub
CONNECTION 'host=publisher port=5432 user=postgres password=postgres dbname=demo'
PUBLICATION my_pub;
```

## 論理レプリケーションを検証
Publisher側でデータを挿入する。

```sql
INSERT INTO users (name) VALUES ('Charlie');
```

Subscriber側でデータを確認する。

```sql
SELECT * FROM users;
```

Publisher側で挿入されたデータがSubscriber側に反映されていることを確認する。

# アーキテクチャ：パブリッシャ・サブスクライバモデル

論理レプリケーションは、\*\*パブリッシャ（publisher）**と**サブスクライバ（subscriber）\*\*という2つの役割によって構成される。パブリッシャは、指定されたテーブルの変更を「パブリケーション（publication）」として公開する。サブスクライバは、これらのパブリケーションを購読し、変更内容を自身のデータベースに反映する。

レプリケーションはPostgreSQLのWAL（Write-Ahead Log）に基づき、論理デコーディングされた変更がリアルタイムで送信される。サブスクライバは、`apply`プロセスによって変更を受け取り、同一トランザクション単位で反映するため、整合性が維持される。

> 参考: [Architecture - Publisher and Subscriber](https://www.postgresql.org/docs/current/logical-replication.html#LOGICAL-REPLICATION-PUBLISHER)

# サポート対象と制限

論理レプリケーションの対象は通常のテーブルに限られ、ビュー、マテリアライズドビュー、シーケンス、外部テーブル、大きなオブジェクト（LOB）は対象外である。また、UPDATEやDELETEを正しくレプリケートするためには、レプリカID（通常は主キー）が必須である。`REPLICA IDENTITY FULL` を用いることで主キーがないテーブルにも対応可能だが、パフォーマンスの低下に留意すべきである。

DDL（テーブル定義の変更）はレプリケーションされないため、パブリッシャとサブスクライバで同一のスキーマ管理が必要となる。

> 参考: [Restrictions](https://www.postgresql.org/docs/current/logical-replication-restrictions.html)

# パブリケーションの構成

パブリケーションは、公開対象のテーブルと対象とする操作（INSERT, UPDATE, DELETE, TRUNCATE）を指定して作成される。特定スキーマ内の全テーブルを対象にした一括公開も可能である。PostgreSQL 15以降では、テーブルごとの**行フィルタ**（WHERE句）や**列リスト**による転送対象の絞り込みにも対応しており、より細かい制御が可能となっている。

> 参考: [CREATE PUBLICATION](https://www.postgresql.org/docs/current/sql-createpublication.html)

# サブスクリプションの構成と同期

サブスクライバ側では、`CREATE SUBSCRIPTION` によって購読を開始する。接続情報と対象のパブリケーション名を指定することで、初期同期（スナップショットのコピー）と、その後の継続的なストリーミングが行われる。初期同期は並列で実行されるため、大量のデータでも比較的短時間で同期が完了する。

レプリケーションスロットがパブリッシャ側に作成され、変更ログの保持と送信の役割を果たす。サブスクライバ側では、`apply`ワーカプロセスがこのデータを逐次反映する。

> 参考: [CREATE SUBSCRIPTION](https://www.postgresql.org/docs/current/sql-createsubscription.html)

# トランザクションと整合性

論理レプリケーションは、トランザクション単位での複製を行う。すなわち、複数テーブルにまたがる変更であっても、パブリッシャ上のコミット単位でサブスクライバに適用される。そのため、データ整合性が維持され、サブスクライバ側でも一貫した状態を保てる。

ただし、複数のサブスクリプションで同一テーブルを同時に更新するような構成を採用する場合、競合が発生する可能性がある。PostgreSQL 15以降では `ALTER SUBSCRIPTION ... SKIP` によるエラー処理の制御が可能である。

> 参考: [Replication Conflicts](https://www.postgresql.org/docs/current/logical-replication-conflicts.html)

# 行・列レベルのフィルタリング（PostgreSQL 15以降）

より柔軟なレプリケーションを実現するため、PostgreSQL 15では**行レベルフィルタ（WHERE句）**と**列リスト**の指定が可能となった。これにより、たとえばマルチテナント環境でテナントごとに別のサブスクライバへ必要な行だけを転送することが可能である。

ただし、行フィルタでは非イミュータブルな関数や副作用を含む式は利用できない。また、UPDATE/DELETE を行う場合はレプリカIDで使用される列のみ条件に含めることができる。

> 参考: [Row and Column Filtering](https://www.postgresql.org/docs/current/logical-replication-row-filter.html)

# 主なユースケース

論理レプリケーションは、以下のようなユースケースに適している。

- **マイグレーション**：バージョンアップや異なるプラットフォームへの移行時に、稼働中のシステムに影響を与えずにデータを移行可能。
- **データ統合・集約**：複数のデータベースから特定のテーブルだけを収集し、BIツールやデータウェアハウスで分析。
- **マルチテナント分離**：1つのDBに存在する複数テナントのデータを、テナントごとのDBに分離して配布。
- **イベント駆動アーキテクチャ**：データベースの変更をトリガとして、外部システムと連携した処理を構築。

> 参考: [Logical Replication Use Cases](https://www.postgresql.org/docs/current/logical-replication-use-cases.html)

# バージョンごとの主な機能追加

- PostgreSQL 10：論理レプリケーション初登場。
- PostgreSQL 13：パーティションテーブルへの対応。
- PostgreSQL 14：ストリーミング適用とバイナリ転送。
- PostgreSQL 15：行・列フィルタ、スキーマ単位公開、`SKIP`など。
- PostgreSQL 16：スタンバイからの複製、並列適用、オリジン指定のループ防止。

> 参考: [Release Notes](https://www.postgresql.org/docs/release/)

# 導入における注意点

論理レプリケーションを導入するには、以下の点に注意すべきである。

- `wal_level = logical` を有効にすること。
- `max_replication_slots`、`max_wal_senders`、`max_logical_replication_workers` など、適切なリソース設定を行うこと。
- DDLの変更はレプリケーションされないため、手動で同期すること。
- サブスクライバはデフォルトで `session_replication_role = replica` となるため、ステートメントトリガは発火しない。

> 参考: [Logical Replication Setup](https://www.postgresql.org/docs/current/logical-replication-setup.html)

このように、PostgreSQLの論理レプリケーションは高い柔軟性を持ちつつも、仕様に基づいた設計と設定が求められる機能である。マイグレーション、データ集約、マルチテナント対応といった幅広いニーズに応える手段として有効である一方、制限事項やバージョン差異にも留意し、適切に運用すべきである。

# 参考
- [www.postgresql.org - Architecture - Publisher and Subscriber](https://www.postgresql.org/docs/current/logical-replication.html#LOGICAL-REPLICATION-PUBLISHER)
- [www.postgresql.org - PostgreSQL Logical Replication Documentation - Overview](https://www.postgresql.org/docs/current/logical-replication.html)
- [www.postgresql.org - Restrictions](https://www.postgresql.org/docs/current/logical-replication-restrictions.html)
- [www.postgresql.org - CREATE PUBLICATION](https://www.postgresql.org/docs/current/sql-createpublication.html)
- [www.postgresql.org - CREATE SUBSCRIPTION](https://www.postgresql.org/docs/current/sql-createsubscription.html)
- [www.postgresql.org - Replication Conflicts](https://www.postgresql.org/docs/current/logical-replication-conflicts.html)
- [www.postgresql.org - Row and Column Filtering](https://www.postgresql.org/docs/current/logical-replication-row-filter.html)
- [www.postgresql.org - Logical Replication Use Cases](https://www.postgresql.org/docs/current/logical-replication-use-cases.html)
- [www.postgresql.org - Release Notes](https://www.postgresql.org/docs/release/)
- [www.postgresql.org - Logical Replication Setup](https://www.postgresql.org/docs/current/logical-replication-setup.html)
