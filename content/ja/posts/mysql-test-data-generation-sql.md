---
title: MySQLでテストデータを生成するSQL
slug: mysql-test-data-generation-sql
date: 2019-07-16T00:00:00Z
author: bmf-san
categories:
  - データベース
tags:
  - MySQL
  - sql
  - cross join
translation_key: mysql-test-data-generation-sql
---


 概要
MySQLだけでテストデータを生成するSQLについてメモしておく。
テストデータをスクリプトで生成する方法は柔軟性が高く、上等手段な気がするが、
数万件のレコードを用いてパフォーマンステストをしたい時などはSQLだけでも十分かもしれない。

# SQL
クエリはこんな感じ。

```sql
DROP TABLE IF EXISTS `tests`;

CREATE TABLE `tests` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `value` int(5) NOT NULL DEFAULT 0,
  PRIMARY KEY (id)
);

INSERT INTO tests(value)
VALUES (1), (2), (3), (4), (5), (6), (7), (8), (9), (10);

DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `title` varchar(255) DEFAULT NULL,
  `body` text DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  FOREIGN KEY (admin_id) REFERENCES admins(id),
  FOREIGN KEY (category_id) REFERENCES categories(id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

INSERT INTO posts(title, body, created_at, updated_at)
SELECT
  (@rownum := @rownum + 1),
  @rownum,
  CONCAT(@rownum, 'title'),
  CONCAT(@rownum, 'md_body'),
  CONCAT(@rownum, 'html_body')
FROM
  tests AS t1,
  tests AS t2,
  tests AS t3,
  tests AS t4,
  (SELECT @rownum := 0) AS v;
```

ユーザー定義変数を使って行番号を取りつつ、直積(CROSS JOIN)と`INSERT INTO ... SELECT`を使ってレコードを生成する方法。
色々なパターンがあったがこれが比較的わかりやすい、というか書きやすい気がする。
パッと見て何をしているのか動作をイメージするのが難点ではある。

# 参考
- [SQLで大量のテストデータ作成](https://qiita.com/cobot00/items/8d59e0734314a88d74c7)
