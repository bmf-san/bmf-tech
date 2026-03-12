---
title: "MySQLのロック解説：デッドロックの防ぎ方とパフォーマンス改善"
description: 'MySQL のロック（行ロック・テーブルロック・GAP ロック）を解説。デッドロックの防ぎ方とトランザクション設計のベストプラクティスを紹介します。'
slug: mysql-locking
date: 2024-04-05T00:00:00Z
author: bmf-san
categories:
  - データベース
tags:
  - MySQL
  - トランザクション
  - ロック
translation_key: mysql-locking
---


# 概要
MySQLのロックについてまとめる。
MySQLのバージョンは8系を想定する。

# 検証環境
検証に使う環境はdocker-composeで用意した。（1コンテナだけなのでcomposeを使わなくも良いのだが..）

```sh
.
├── docker-compose.yml
└── initdb.d
    └── 1_schema.sql
```

docker-compose.yml
```yml
version: '3'

services:
  mysql:
    image: mysql:8.0.33
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: example
      TZ: "Asia/Tokyo"
    command: mysqld
    ports:
      - 3306:3306
    volumes:
      - ./initdb.d:/docker-entrypoint-initdb.d
```

1_schema.sql
```sql
CREATE DATABASE IF NOT EXISTS example;
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(255) NOT NULL UNIQUE
) ENGINE = InnoDB DEFAULT CHARSET = utf8;
```

`docker compose up`でお手元にMySQL8系のコンテナが用意できる。

# ロック
## 内部レベルロック
MySQLにおける排他制御の手法としては、行レベルロックとテーブルレベルロックがある。

cf. [dev.mysql.com - 8.11.1 内部ロック方法](https://dev.mysql.com/doc/refman/8.0/ja/internal-locking.html)

- 行レベルロック
  - テーブル内の個々の行を対象としたロック
  - ロック対象が狭いのでロックの競合、ロールバックする変更が少なくなる
  - 1つの行を長時間ロック可能
- テーブルレベルロック
  - テーブルを対象としたロック
  - 必要になるメモリーが比較的に少ない（行ロックはロックされた行また行のグループごとにメモリーが必要）
  - 単一のロックだけが必要となるため、テーブルの大部分を対象に使用する場合は高速
  - データの大部分を対象にGROUP BYを頻繁に実行する場合やテーブル全体を頻繁にスキャンする場合は高速

## InnoDBロック
cf. [dev.mysql.com - 15.7.1 InnoDB ロック](https://dev.mysql.com/doc/refman/8.0/ja/innodb-locking.html)

### 共有（READ）ロック
共有ロックは、データのREADは可能だが、WRITEはできないロック。Shared lock（IS）。

#### 検証
1. TX1でトランザクションを開始、共有ロックをかける
```sql
// TX1
mysql> INSERT INTO users(name) VALUES('foo'); // 初期データ投入
mysql> START TRANSACTION;
mysql> SELECT * FROM users WHERE id = 1 LOCK IN SHARE MODE;
```

2. TX2でトランザクションを開始、WRITEを行う
```sql
// TX2
mysql> START TRANSACTION;
mysql> UPDATE users SET name = 'bar' WHERE id = 1;
```

TX1がCOMMITするまでTX2の更新はロックされる。

### 占有（排他・WRITE）ロック
排他ロックは、データのREADもWRITEもできないロック。Exclusive lock（IX）。

#### 検証
1. TX1にてトランザクションを開始、占有ロックをかける
```sql
// TX1
mysql> INSERT INTO users(name) VALUES('foo'); // 初期データ投入
mysql> START TRANSACTION;
mysql> SELECT * FROM users WHERE id = 1 FOR UPDATE;
```

2. TX2でトランザクションを開始、READ、WRITEを行う
```sql
// TX2
mysql> START TRANSACTION;
mysql> SELECT * FROM users WHERE id = 1; // これは許容されるが
mysql> SELECT * FROM users WHERE id = 1 FOR UPDATE; // 許容されない
mysql> UPDATE users SET name = 'bar' WHERE id = 1; // 許容されない
```

TX1のロックが解放されるまでTX2ではREAD（単純なSELECT以外）やWRITEができないことが確認できる。

### インテンションロック
トランザクションがテーブルの行に必要とするロックタイプ（共有または排他）を示すテーブルレベルのロック。
行ロックとテーブルロックの共存をサポートするために用意されている。

インテンションロックには、

- インテンション共有ロック
- インテンション排他ロック

の2つがある。

#### 検証
SQLで明示的に操作できるものではなく、基本的にはデータベース内部で管理されるものであるので、検証は割愛。

いくつか検証パターンがあるが、以下の記事で色々と検証されている。

cf. [qiita.com - MySQLのロックについて公式ドキュメントを読みながら動作検証してみた〜行レベルロック: インテンションロック〜](https://qiita.com/ham0215/items/2f38a2949d9012074c3d)

### レコードロック
インデックスレコードのロック。インデックスレコードとはクラスタインデックスとセカンダリインデックスのこと。スキャンしたインデックスに対してロックする。

#### 検証
データベースの内部的な動作であるため割愛。

### ギャップロック
インデックスレコード間のギャップのロック。または、インデックスレコードの前または後ろのギャップのロック。

#### 検証
1. TX1でトランザクション開始、READを行う
```sql
// TX1
mysql> INSERT INTO users(id, name) VALUES(1, 'foo'), (2, 'bar'), (4, 'qux'), (5, 'quux'), (6, 'corge'); // 初期データ投入
mysql> START TRANSACTION;
mysql> SELECT * FROM users WHERE ID between 1 AND 5 FOR UPDATE;
```

2. TX2でトランザクション開始、WRITEを行う
```sql
// TX2
mysql> START TRANSACTION;
mysql> INSERT INTO users(id, name) VALUES(3, 'baz');
```

行単位のロックかと思いきや、範囲でロックされているのが確認できる。

### ネクストキーロック
インデックスレコードのレコードロックとインデックスレコードの前のギャップのギャップロックの組み合わせ。

#### 検証
1. TX1でトランザクション開始、READを行う
```sql
// TX1
mysql> INSERT INTO users(id, name) VALUES(1, 'foo'), (2, 'bar'), (3, 'baz'), (4, 'qux'); // 初期データ投入
mysql> START TRANSACTION;
mysql> SELECT * FROM users WHERE ID < 5 FOR UPDATE;
```

2. TX2でトランザクション開始、WRITEを行う
```sql
// TX2
mysql> START TRANSACTION;
mysql> INSERT INTO users(id, name) VALUES(5, 'quux');
```

idが5未満の行だけでなく、末尾のインデックス値を持つ行の後のギャップもロックされることが確認できる。

### インテンションロックの挿入
行の挿入前のINSERTによって設定されるギャップロックのタイプ。INSERTのインテンションロック。

#### 検証
データベースの内部的な動作であるため割愛。

こちらの記事で検証されているので参照。
cf. [MySQLのロックについて公式ドキュメントを読みながら動作検証してみた〜レコードロック / ギャップロック / ネクストキーロック / 他〜](https://qiita.com/ham0215/items/99679d499869365446ec#%E3%82%A4%E3%83%B3%E3%83%86%E3%83%B3%E3%82%B7%E3%83%A7%E3%83%B3%E3%83%AD%E3%83%83%E3%82%AF%E3%81%AE%E6%8C%BF%E5%85%A5)

### AUTO-INCロック
AUTO_INCREMENTカラムを含むテーブルに挿入されるトランザクションによって取得されるテーブルロック。
TX1でのトランザクションでINSERTするためにAUTO_INCREMENTの値を取得している間はTX2でのAUTO_INCREMENTの値を取得できないようするロック。

#### 検証
内部的な動作である&再現方法が分からなかったので割愛。

### 空間インデックスの述語ロック
これはドキュメント参照。（空間インデックスに触りなれていないのものあってイマイチ分からなかった。。。）

cf. [空間インデックスの述語ロック](https://dev.mysql.com/doc/refman/8.0/ja/innodb-locking.html#innodb-auto-inc-locks)

# ロックの確認方法
ロックは以下のクエリで確認することができる。

```sql
// ロックの状態確認
SELECT * FROM performance_schema.data_locks;

// ロック件数確認+スレッドID
SHOW ENGINE INNODB STATUS;

// ロック件数確認
SELECT trx_id, trx_rows_locked, trx_mysql_thread_id FROM information_schema.INNODB_TRX;
```

デッドロックを確認するには、`SHOW ENGINE INNODB STATUS`を実行し、`LATEST DETECTED DEADLOCK`と記載されている部分を確認する。

# まとめ
MySQLには明示的・暗黙的にロックされるパターンがある。

何が（行なのかテーブルなのか）対象なのか、範囲はどこまでなのかといったことにまずは目を向けると良さそう。

# 参照
- [dev.mysql.com - 8.11.1 内部ロック方法](https://dev.mysql.com/doc/refman/8.0/ja/internal-locking.html)
- [dev.mysql.com - 15.7.1 InnoDB ロック](https://dev.mysql.com/doc/refman/8.0/ja/innodb-locking.html)
- [zenn.dev - データベースのロックの基礎からデッドロックまで](https://zenn.dev/gibjapan/articles/1d8dfb7520dabc)
- [qiita.com - MySQLのロックについて公式ドキュメントを読みながら動作検証してみた〜レコードロック / ギャップロック / ネクストキーロック / 他〜](https://qiita.com/ham0215/items/99679d499869365446ec)
- [qiita.com - MySQL のロックについて補足（注：すでに語りつくされている内容です）](https://qiita.com/hmatsu47/items/f5eb64428494686d4ad3)
- [qiita.com - MySQLのロックについて公式ドキュメントを読みながら動作検証してみた〜行レベルロック: 共有ロック(S) / 排他ロック(X) 〜](https://qiita.com/ham0215/items/b9efc718670b1d2d48c1#%E8%A1%8C%E3%83%AC%E3%83%99%E3%83%AB%E3%83%AD%E3%83%83%E3%82%AF%E3%82%92%E6%A4%9C%E8%A8%BC%E3%81%99%E3%82%8B)
- [qiita.com - MySQLのロックについて公式ドキュメントを読みながら動作検証してみた〜行レベルロック: インテンションロック〜](https://qiita.com/ham0215/items/2f38a2949d9012074c3d)
- [techblog.cartaholdings.co.jp - DBのロックについてあまり意識したことがない人に向けた実は覚えておきたいロックについての知識](https://techblog.cartaholdings.co.jp/entry/2022/12/14/113000)
- [www.wakuwakubank.com - 占有ロック(FOR UPDATE)と共有ロック(LOCK IN SHARE MODE)](https://www.wakuwakubank.com/posts/201-mysql-lock/)
- [saekis.hatenablog.com - MySQLの排他ロックの挙動を確認する](https://saekis.hatenablog.com/entry/2019/02/06/191454)
- [bizstation.hatenablog.com - MySQL/MariaDBとTransactdのInnoDBロック制御詳細　その1](https://bizstation.hatenablog.com/entry/2014/12/24/103641)
- [devsakaso.com - 【MySQL】ロックとデッドロックについて](https://devsakaso.com/mysql-about-lock-and-deadlock/)
- [nishinatoshiharu.com - 【MySQL】InnoDBの共有ロックと排他ロックの概要と挙動検証](https://nishinatoshiharu.com/db-lock-overview/)
- [www.wakuwakubank.com - 占有ロック(FOR UPDATE)と共有ロック(LOCK IN SHARE MODE)](https://www.wakuwakubank.com/posts/201-mysql-lock/)
- [free-engineer.life - MySQL(InnoDB)共有ロックと排他ロックとインテンションロック（テーブルロック）](https://free-engineer.life/mysql-innodb-lock-mode/)
- [free-engineer.life - MySQL(InnoDB)の行ロック](https://free-engineer.life/mysql-innodb-record-locks/)
- [github.com - MySQLのInnoDBのロック挙動調査](https://github.com/ichirin2501/doc/blob/master/innodb.md)
- [github.com - スレッドIDを利用したMySQLのデッドロック解析手法](https://github.com/ichirin2501/doc/blob/master/innodb-deadlock-thread-id.md)
