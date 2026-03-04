---
title: "MySQLで全文検索"
slug: "full-text-search-with-mysql"
date: 2023-04-30
author: bmf-san
categories:
  - "データベース"
tags:
  - "MySQL"
draft: false
---

# 概要
MySQLにはだいぶ前から全文検索が使えるになっているが、最近まで全然触ってもいなかったので軽く素振りしてみた。

#　MySQLで全文検索を始める
MySQLで全文検索を利用するのはElasticSearchよりも圧倒的に手間が掛からない。

検索対象としたいカラムに**FULLTEXT INDEX**を付与、**MATCH (col1,col2,...) AGAINST (expr [search_modifier])** で検索クエリを投げるだけで全文検索がお手軽にできてしまう。

ex. 
```sql
// FULLTEXT INDEX付与対象のカラムを持つテーブル
CREATE TABLE `posts` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `title` varchar(255) DEFAULT NULL,
  `body` longtext DEFAULT NULL,
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

// FULLTEXT INDEXを付与
ALTER TABLE posts ADD FULLTEXT INDEX index_title_md_body (title, md_body) WITH PARSER ngram;

// MATCH ... AGINSTで検索クエリ
SELECT
  *
FROM
  posts
WHERE MATCH (title, body)
AGAINST ("MySQLで全文検索" IN BOOLEAN MODE)
```

FULLTEXT INDEXはALTER TABLE以外にもCREATE TABLEやCREATE INDEXで付与することもできる。

# 全文パーサー
MySQLの全文検索のパーサーにはngramとMeCabに対応している。

- [ngram 全文パーサー](https://dev.mysql.com/doc/refman/8.0/ja/fulltext-search-ngram.html)
- [MeCab フルテキストパーサープラグイン](https://dev.mysql.com/doc/refman/8.0/ja/fulltext-search-mecab.html)

デフォルトではngramが設定される。

MeCabを使いたい場合はプラグインの導入が必要。
 
# 全文検索のモード
3つのモードがあり、使いたいモードを指定することができる。

モードによって検索結果に差が出るので、どういう検索体験にしたいかによって選択の余地がある。

- NATURAL LANGUAGE MODE（自然言語検索）
  -　検索ワードを自然言語処理によって検索するモード
  - cf. [自然言語全文検索](https://dev.mysql.com/doc/refman/8.0/ja/fulltext-natural-language.html)
- BOOLEAN MODE（ブール検索）
  - 検索ワードをAND、OR、NOTなどの条件によって検索するモード
  - cf. [ブール全文検索](https://dev.mysql.com/doc/refman/8.0/ja/fulltext-boolean.html)
- QUERY EXPANSION（クエリー拡張検索）
  - 検索ワードに類義語や関連語を追加する形で検索するモード
  - cf. [クエリー拡張を使用した全文検索](https://dev.mysql.com/doc/refman/8.0/ja/fulltext-query-expansion.html)

# 検索の精度調整
パーサーや全文検索のモード以外の検索の性質を調整するアプローチとしては、

- パーサーの設定値変更
  - ex. ngramのトークンサイズ変更
- 文字コードの指定
  - ex. utf8_general_ci、utf8_unicode_ci etc...
- 全文ストップワードの調整
  - cf. [全文ストップワード](https://dev.mysql.com/doc/refman/8.0/ja/fulltext-stopwords.html)

などある模様。

[MySQL の全文検索の微調整](https://dev.mysql.com/doc/refman/8.0/ja/fulltext-fine-tuning.html)も参照。

# 所感
このブログにもMySQLの全文検索機能を取り入れてみた。

[検索記事一覧](https://bmf-tech.com/posts/search?keyword=)

LIKE検索より性能が良いと思われるが、実際どこまでパフォーマンスが維持できるかは環境ごとに性能検証が必要と思われるが、要件が満たせるのであれば十分に使える機能だということが分かった。

