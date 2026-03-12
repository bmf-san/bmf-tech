---
title: "データベースインデックスとは？仕組みと必要なタイミング"
description: 'データベースインデックスとは何か、B-Tree・ハッシュインデックスの仕組み・効果的に使う場面・作り過ぎのデメリットを解説します。'
slug: what-is-index
date: 2024-04-01T00:00:00Z
author: bmf-san
categories:
  - データベース
tags:
  - DB
  - インデックス
  - MySQL
translation_key: what-is-index
---


# インデックスとは
テーブルに格納されているレコードを高速に取り出すための仕組み

以下のようなO(n)問題を抱えたクエリがあるとする。

`SELECT * FROM users WHERE first_name = ‘Tom’`

このクエリのパフォーマンスを上げるためには、以下のようにIndexを貼る。

`ALTER TABLE users ADD INDEX (first_name)`

# メリット・デメリット
## メリット
- データの読み込み・取得の速度向上

## デメリット
- 容量の増加
- 書き込み速度の低下

データの作成・更新時には、同時にインデックスの追加・更新も行われるため、上記のようなデメリットが生じる。

# インデックスのパターン
## 通常（1つのカラムに対して貼る）

`ALTER TABLE users ADD INDEX (first_name)`

## 部分的インデックス
容量の増加を抑えつつ、性能を向上させたい時に有効なパターン

最初の4バイトだけにインデックスを貼る例
`ALTER TABLE users ADD INDEX (first_name(4))`

## マルチカラムインデックス（複合または合成インデックスとも呼ぶ）

`ALTER TABLE users ADD INDEX (last_name, first_name)`

MySQLでは、**1つのクエリを実行する際、1テーブルについて1インデックスしか使用できない**が、マルチインデックスを適用していれば、有効なインデックスがクエリ実行の際に選択される。

マルチカラムインデックスの先頭に指定するカラムはカーディナリティの高いものにしておくのが通常良い。

## ユニークインデックス
NULLを除いて値が重複して出現しなくなる。
レコードの作成・更新において、全ての値を調べて同じ値が既に存在しないことを確認する。
MySQLでは、ユニークキーを指定するとユニークインデックスも指定される。

`ALTER TABLE users ADD UNIQUE (first_name)`

# インデックスの効果測定
EXPLAIN句でクエリの実行計画を確認。

`EXPLAIN SELECT * FROM users WHERE first_name = ‘Tom’`

以下の項目を確認
- possible_keys
 - 選択可能なインデックス
- key
 - 実際に選択されたインデックス
- extra
 - 以下のような表示がある場合はクエリの最適化をしたほうが良い
  - using filesort
   - ソートに必要なメモリが不足し、物理ファイルに書き出しソートを行っている
  - using temporary
   - クエリの実行のためにテンポラリテーブルが作成されている

# インデックスの基準
インデックスを検討したほうが良いかもしれない判断基準をリストアップ。
あくまで推測するための基準なので、EXPLAINでの計測をしたほうが良い。

- テーブルのデータ量が多く、検索対象のレコードが少量になる時
- WHERE、JOIN、ORDER BYなどに利用されるカラムがある合
- NULLを多く含むデータからNULL以外の検索をする（IndexはNULLを含まないため有効な可能性がある）
- データの追加・更新・削除などは頻繁に行われないデータの場合（Indexの更新負荷を考慮）

# クラスタインデックスとセカンダリインデックス
## クラスタインデックス
以下に該当するインデックスがクラスタインデックスとなる。

- 主キーで定義されたカラム
- NOT NULLのユニークキーのカラム
- 上記に該当するカラムがない場合は、InnoDBはGEN_CLUST_INDEXという非表示のクラスタインデックスを作成する

## セカンダリインデックス
クラスタインデックス以外のインデックスをセカンダリインデックスという。
セカンダリインデックスには主キーの値が含まれている。
EXPLAINで計測することが前提がだが、主キーの値が含まれているので、カバンリングインデックスを狙う際は複合インデックスに主キーを含めなくてもセカンダリインデックスだけでカバリングインデックスになる可能性を覚えておくと良いかも。
cf. [知って得するInnoDBセカンダリインデックス活用術！](https://nippondanji.blogspot.com/2010/10/innodb.html)

 # カバリングインデックス
クエリの実行結果に必要なすべてのカラムを含むインデックスのこと。

データファイルを読まず、インデックスだけでカバーできるので、検索が高速化される。

# インデックスを貼る時の注意
## インデックス列の演算やSQL関数
```sql
SELECT * FROM users WHERE amount * 2 > 10;
```

amountにインデックスを貼った場合、インデックスを活用するには演算子を避ける。amount自体がインデックスに保持されており、演算結果が保持されているわけではない。SQL関数についても同様。

```sql
SELECT * FROM users WHERE amount > 10/2;
```

## IS_NULL
```sql
SELECT * FROM users WHERE amount IS NULL;
```

IS NULLやIS NOT NULLは基本的にはインデックスが有効に活用されない（DBMSの仕様に依る）。

## 否定形やOR
```sql
SELECT * FROM users WHERE amount <> 10;
```

否定形はインデックスを活用できない。ORについても同様。

## LIKE
```sql
SELECT * FROM users WHERE name = 'a%';
```

LIKEを使う場合は、B-Treeの性質に依るため前方一致のみインデックスが活用される。

## 暗黙的な型変換
```sql
SELECT * FROM users WHERE age = '10'
```

ageが数値型の場合に文字列から数値に暗黙的な型がされるとインデックスが活用されなくなる。

# 参考
- [amzn.to - 理論から学ぶデータベース実践入門 ~リレーショナルモデルによる効率的なSQL (WEB+DB PRESS plus) ](https://amzn.to/3TEltzx)
- [www.hi-ho.ne.jp - インデックスの基礎知識](http://www.hi-ho.ne.jp/tsumiki/doc_1.html)
- [kiyotakubo.hatenablog.com - MySQLパフォーマンス・チューニングのためのインデックスの基礎知識](http://kiyotakakubo.hatenablog.com/entry/20101117/1289952549)
- [qiita.com - MySQLのインデックスを貼るコツ](https://qiita.com/katsukii/items/3409e3c3c96580d37c2b#%E9%80%9A%E5%B8%B8)
- [dev.mysql.com - EXPLAIN出力フォーマット](https://dev.mysql.com/doc/refman/5.6/ja/explain.html)
- [dev.mysql.com - カバリングインデックス](https://dev.mysql.com/doc/refman/8.0/ja/glossary.html#glos_covering_index)
