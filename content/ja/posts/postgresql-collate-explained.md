---
title: "PostgreSQLにおけるCOLLATE（照合順序）について"
slug: "postgresql-collate-explained"
date: 2025-03-05
author: bmf-san
categories:
  - "データベース"
tags:
  - "PostgreSQL"
draft: false
---

PostgreSQLにおけるCOLLATE（照合順序）について調べたことをまとめる。

## 1. COLLATE（照合順序）とは

COLLATEとは、文字列の並び順や比較の仕方（大文字・小文字の扱いやアクセント、濁点など）を指定する仕組みである。たとえば`ORDER BY`による並び順や比較演算子の結果にも影響するため、日本語環境下でのアプリケーション開発においては正しいCOLLATEを設定しておくことが推奨される。

### COLLATE指定例

- **データベース作成時に指定する**
  ```sql
  CREATE DATABASE dbname
    LC_COLLATE='ja_JP.UTF-8'
    LC_CTYPE='ja_JP.UTF-8'
    TEMPLATE=template0;
  ```

- **テーブルやカラム作成時に指定する**
  ```sql
  CREATE TABLE example (
    name TEXT COLLATE "ja_JP.UTF-8"
  );
  ```

- **クエリ内で一時的に指定する**
  ```sql
  SELECT name
    FROM example
   ORDER BY name COLLATE "en_US.UTF-8";
  ```

## 2. COLLATEの設定状況を確認するクエリ

PostgreSQLでは、COLLATEがどこでどのように使われているかをクエリによって確認できる。以下に主な方法を示す。

### 2-1. データベース単位のCOLLATE設定

現在接続中のデータベース（`current_database()`）について、どの照合順序（`LC_COLLATE`）および文字種（`LC_CTYPE`）が設定されているかを確認するには、次のクエリを実行する。

```sql
SELECT datname,
       datcollate AS collate,
       datctype   AS ctype
  FROM pg_database
 WHERE datname = current_database();
```

### 2-2. テーブルおよびカラムにおけるCOLLATE設定

どのテーブル・カラムがCOLLATEを持っているかを調べるには、`information_schema.columns`テーブルを参照するとよい。以下のクエリでは、`collation_name`が設定されているカラムのみを取得している。

```sql
SELECT table_schema,
       table_name,
       column_name,
       collation_name
  FROM information_schema.columns
 WHERE collation_name IS NOT NULL
   AND table_schema NOT IN ('information_schema','pg_catalog')
 ORDER BY table_schema, table_name, column_name;
```

### 2-3. COLLATE付きINDEXの確認

COLLATEを指定しているインデックスを確認したい場合は、`pg_index`や`pg_class`を参照し、インデックス定義（`pg_get_indexdef`）を検索する方法がある。次のクエリでは、インデックス定義に`COLLATE`が含まれているかどうかを確認している。

```sql
SELECT idx.relname       AS index_name,
       tbl.relname       AS table_name,
       pg_get_indexdef(idx.oid) AS index_definition
  FROM pg_index i
  JOIN pg_class idx ON idx.oid = i.indexrelid
  JOIN pg_class tbl ON tbl.oid = i.indrelid
 WHERE idx.relkind = 'i'
   AND pg_get_indexdef(idx.oid) ILIKE '%COLLATE%'
 ORDER BY table_name, index_name;
```

該当データがない場合は結果が空になる。

### 2-4. 関数でのCOLLATE利用確認

ユーザー定義関数（`prokind = 'f'`）の定義内にCOLLATE指定が含まれているかどうかを調べるには、以下のようなクエリを利用する。

```sql
SELECT proname,
       pg_get_functiondef(oid) AS function_definition
  FROM pg_proc
 WHERE prokind = 'f'
   AND pg_get_functiondef(oid) ILIKE '%COLLATE%';
```

同様に結果が空であれば、COLLATEを含む関数は存在しないことになる。

### 2-5. トリガーでのCOLLATE利用確認

トリガー定義の中にCOLLATEが含まれているかを調べる方法は、以下のクエリを利用する。`pg_trigger`を参照し、内部トリガー（`tgisinternal`）を除外する条件を設定している。

```sql
SELECT tgname,
       pg_get_triggerdef(oid) AS trigger_definition
  FROM pg_trigger
 WHERE NOT tgisinternal
   AND pg_get_triggerdef(oid) ILIKE '%COLLATE%';
```

こちらも空の場合は、COLLATEを含むトリガーがないことを示す。

## 3. まとめ

PostgreSQLで日本語を扱う際は、照合順序（COLLATE）の適切な設定が重要である。英語圏のデフォルトのままでは、大文字・小文字の扱いや濁点・半濁点などの順序付けが期待どおりにならない可能性が高い。そのため、

- データベース作成時に`LC_COLLATE='ja_JP.UTF-8'`などを指定しておく
- テーブルやカラムにCOLLATEを設定し、必要に応じてクエリ単位でも指定する
- インデックスや関数、トリガーでもCOLLATEを利用している場合は適宜確認する

といった運用が推奨される。もし異なる言語設定が求められる場合には、英語のCOLLATEや他言語のCOLLATEを動的に指定してソート順や比較の挙動を切り替えることも可能である。開発・運用時には、実際に日本語のデータを格納してソートや比較結果をテストし、意図どおりの動作になることを確認することが大切である。

# 参考
- [PostgreSQL公式ドキュメント - Collation Support](https://www.postgresql.org/docs/current/collation.html))*
