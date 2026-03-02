---
title: "PostgreSQLのCOLLATEとglibcのバージョン差によるソート順の違い"
slug: "postgresql-collate-glibc"
date: 2025-03-05
author: bmf-san
categories:
  - "データベース"
tags:
  - "PostgreSQL"
draft: false
---

# PostgreSQLのCOLLATEとglibcのバージョン差によるソート順の違い
同じCOLLATE設定を指定しているにもかかわらず、環境によってソート順が異なる問題に遭遇したので、その時の調査記録をメモしておく。

## 事象
### Cloud SQLとローカル環境のPostgreSQLコンテナでCOLLATEが同じでもソート結果が違う

Cloud SQL for PostgreSQL 17上でテーブルの文字列カラムを`ORDER BY カラム名 COLLATE "en_US.utf8"`でソートしたところ、ローカルで動かしているPostgreSQLコンテナの結果と並びが異なる問題に遭遇した。
同じバージョンのPostgreSQL・同じCOLLATE設定（データベースのCOLLATEやクエリで直指定するCOLLATE設定が同じ）なのに、なぜか順序が違うという事象である。

## 調査
最初はCOLLATEが効いていないことを疑ったが、COLLATEはEXPLAIN ANALYZEで確認すると正しく適用されていることが確認できたり、データベースレベル・クエリレベルでのCOLLATE指定の違いを検証したことでCOLLATEの問題ではないことがわかった。

### glibcバージョンを確認する
PostgreSQLが文字列比較を行う際、**glibc**（collprovider = `c`）か**ICU**（collprovider = `i`）を利用する。もしglibcを利用している場合、glibcのバージョンが違うと同じロケール名（`en_US.UTF8`など）でも実際のソート順が変わる可能性がある。

どのバージョンのglibcが使われているかは、PostgreSQL内で以下のクエリを実行すれば確認できる。

```sql
SELECT oid, collname, collprovider, collversion
FROM pg_collation
WHERE collname = 'en_US.utf8';
```

- Cloud SQL (PostgreSQL 17)では`collversion`が`2.19`だった
- ローカル環境のコンテナでは`collversion`が`2.31`だった

同じ`en_US.UTF8`でも、Cloud SQLは**glibc 2.19**、ローカルは**glibc 2.31**を使っていたというわけである。

### glibc 2.19環境で再検証
ローカルで簡単に利用できるDockerイメージの多くは、glibcが比較的新しいバージョンになっている。そこで、glibc 2.19を使う**PostgreSQL 14**のコンテナイメージを独自に用意し（または[こちら](https://hub.docker.com/r/bmfsan/collversion-2.19-postgres-v14)を使用し）、同じクエリを実行したところ、Cloud SQLと同じソート結果になった。
つまり、**glibc 2.19**と**glibc 2.31**の違いが、この問題の原因であることが確かめられた。

## glibc依存を回避する方法: ICU
PostgreSQL 10以降では、ICU (collprovider = `i`)を用いて文字列の照合を行う機能が追加されている。ICUを使えば、glibcのバージョンによるソート順の変化を回避できる可能性が高い。

ICUを使うには、PostgreSQLがICU対応でビルドされている必要がある。また、`CREATE DATABASE`時やカラム定義時に`en-US-x-icu`のようなICUロケールを指定することで、glibcの影響を受けないソートを実現できる。（未検証だがたぶん・・・。）

## まとめ
- **glibcのバージョン差により、同じロケール名でもソート順が変わる**
- glibc依存を回避したいなら**ICU**を使う手段がある
