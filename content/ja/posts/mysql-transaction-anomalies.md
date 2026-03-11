---
title: MySQLのトランザクションのアノマリーについて
slug: mysql-transaction-anomalies
date: 2023-06-08T00:00:00Z
author: bmf-san
categories:
  - データベース
tags:
  - トランザクション
  - MySQL
translation_key: mysql-transaction-anomalies
---


# 概要
MySQLのトランザクションのアノマリーについてまとめる。
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

#　トランザクション分離レベル
MySQLのInnoDBでは、ANSI/ISO SQL標準で定められている4つのトランザクション分離レベルが提供されている。

|    分離レベル     | ダーティリード | インコンシステントリード | ロストアップデート | ファントムリード |
| ----------------- | -------------- | ------------------------ | ------------------ | ---------------- |
| READ UNCOMMITTED  | ○              | ○                        | ○                  | ○                |
| READ COMMMITED    | ×              | ○                        | ○                  | ○                |
| REPEATABLE READ※1 | ×              | ×                        | ○                  | ○※               |
| SERIALIZABLE      | ×              | ×                        | ×                  | ×                |


※1MySQLではREPEATABLE READがデフォルトとなっている。

※2上記では○になっているが、MySQLではREPEATABLE READにおいてファントムリードが発生しないようになっている。

トランザクションの分離レベルはREAD UNCOMMITTEDが一番低く、SERIALIZABLEが一番高い。上記は上から低い順となっている。基本的には分離性が高いほど性能が低下する傾向にある。

トランザクションについては、[トランザクション概観](https://bmf-tech.com/posts/%E3%83%88%E3%83%A9%E3%83%B3%E3%82%B6%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3%E6%A6%82%E8%A6%B3)にもまとめている。

# アノマリー
トランザクションにおけるアノマリーについてMySQLで再現してみる。

アノマリーとは、「トランザクションの分離レベルや処理順序によって生じる期待しない結果や不整合」のこと。

アノマリーはANSI SQL標準やISO/IEC 9075によって定義されているものがあり、ここで取り上げるアノマリー以外にも色々ある。

インコンシステントリードについてはそれらの標準に定義されたものではない。（どこで定義されているのかはわからなかった。。.）

トランザクションはTXと表記する。複数トランザクションを区別するために数字を使う。（ex. TX1、TX2）

## ダーティリード
ダーティリードは、TX1がTX2のCOMMIT前のデータを読み取ってしまう現象。

### 検証
すべてのセッションはREAD UNCOMMITEDで行う。
```sql
mysql> SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
```

1. TX1、TX2にてトランザクションを開始
```sql
// TX1
mysql> START TRANSACTION;
// TX2
mysql> START TRANSACTION;
```

2. TX2にてデータ追加
```sql
// TX2
mysql> INSERT INTO users(name) VALUES('foo');
```

TX2にてデータを追加、COMMITはしない。

3. TX1にて再度データ読み取り
```sql
// TX1
mysql> SELECT * FROM users; // 1 row in set
```

TX1でTX2のCOMMIT前のデータが読み取れてしまっている。

## インコンシステントリード
インコンシステントリードは、読み取るデータに一貫性がない現象。

[いろんなAnomaly#Inconsistent Read Anomaly](https://qiita.com/kumagi/items/5ef5e404546736ebac49#inconsistent-read-anomaly)を参照とした。

これについては正確な定義がちょっと分からなかったので、理解が正しいか怪しい。。

COMMIT後の一貫性のなさということなので、インコンシステントリードはファジーリードやファントムリードの上位概念？？な感じがするが、厳密はおそらく違うはず・・。

### 検証
すべてのセッションはREAD UNCOMMITEDで行う。
```sql
mysql> SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
```

1. TX1にてトランザクション開始、データ読み取り
```sql
// TX1
mysql> START TRANSACTION;
mysql> SELECT * FROM users; // Empty set
```

2. TX2にてトランザクション開始、データ追加
```sql
// TX2
mysql> START TRANSACTION;
mysql> INSERT INTO users(name) VALUES('foo');
mysql> COMMIT;
```

3. TX1にて再度読み取り
```sql
// TX1
mysql> SELECT * FROM users; // 1 row in set
```

最初に読み取った結果と違う結果（TX2の処理結果）が取得され、一貫性がなくなっていることが確認できる。

## ファジーリード（ノンリピータブルリード）
ファジーリードは、TX1が他のTX2にて更新したデータを参照できてしまう現象。

すべてのセッションはREAD COMMITTEDで行う。
```sql
mysql> SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;
```

### 検証
1. TX1にてトランザクション開始、データ読み取り
```sql
// TX1
mysql> INSERT INTO users(name) VALUES('foo'); // 初期データ投入
mysql> START TRANSACTION;
mysql> SELECT * FROM users; // 1 row in set
```

初期データ投入結果。
```
+-----+------+
| id  | name |
+-----+------+
| 1   | foo  |
+-----+------+
```


2. TX2にてトランザクション開始、データ読み取り
```sql
// TX2
mysql> START TRANSACTION;
mysql> UPDATE users SET name = 'bar' WHERE id = 1;
mysql> COMMIT;
mysql> SELECT * FROM users; // 1 row in set
```

更新が完了。
```
+-----+------+
| id  | name |
+-----+------+
| 1   | bar  |
+-----+------+
```


4. TX1にて再度データ読み取り
```sql
// TX1
mysql> SELECT * FROM users; // 1 row in set
```

TX2のCOMMITが影響し、TX1の読み取り結果が変わったことが確認できる。
```
+-----+------+
| id  | name |
+-----+------+
| 1   | bar  |
+-----+------+
```

## ファントムリード
ファントムリードは、TX2が新規追加または削除をCOMMITした場合にTX1が読み取るデータが変わってしまう現象。
ファジーリードは更新処理、ファントムリードは新規追加または削除が対象とした現象である。

### 検証
すべてのセッションはREAD COMMITTEDで行う。
```sql
mysql> SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;
```

1. TX1にてトランザクション開始、データ読み取り
```sql
// TX1
mysql> START TRANSACTION;
mysql> SELECT * FROM users; // Empty set
```

2. TX2にてデータを追加、COMMIT
```sql
// TX2
mysql> START TRANSACTION;
mysql> INSERT INTO users(name) VALUES('foo');
mysql> COMMIT;
mysql> SELECT * FROM users;
```

追加が完了。
```
+-----+------+
| id  | name |
+-----+------+
| 1   | foo  |
+-----+------+
```

3. TX1にて再度データを取得
```sql
// TX1
mysql> SELECT * FROM users; // 1 row in set
```

TX2のCOMMITが影響し、TX1の読み取り結果が変わったことが確認できる。
```
+-----+------+
| id  | name |
+-----+------+
| 1   | foo  |
+-----+------+
```

## ロストアップデート
ロストアップデートは、TX1とTX2が同じデータを更新する際に競合が発生し、一部の更新が失われる現象。

### 検証
すべてのセッションはREPEATABLE READで行う。
```sql
mysql> SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ;
```

1. TX1でトランザクション開始、データ読み取り
```sql
// TX1
mysql> INSERT INTO users(name) VALUES('foo'); // 初期データ投入
mysql> START TRANSACTION;
mysql> SELECT * FROM users; // 1 row set
```

2. TX2にてトランザクション開始、データ読み取り
```sql
// TX2
mysql> START TRANSACTION;
mysql> SELECT * FROM users; // 1 row set
```

3. TX1、TX2にそれぞれデータを更新
```sql
// TX1
mysql> UPDATE users SET name = 'tx1' WHERE id = 1;

// TX2
mysql> UPDATE users SET name = 'tx2' WHERE id = 1;
```

4. TX1、TX2をそれぞれCOMMIT
```sql
// TX1
mysql> COMMIT;
// TX2
mysql> COMMIT;
```

5. データ読み取り
```sql
mysql> SELECT * FROM users; 1 row set
```

TX1のCOMMITが失われてTX2にCOMMITが反映されていることが確認できる。
```
+-----+------+
| id  | name |
+-----+------+
| 1   | tx2  |
+-----+------+
```

# まとめ
トランザクションの分離レベルによって、発生するアノマリーは異なる。

アノマリーはCOMMIT前後でのデータの読み取りや一貫性が変わる現象としていくつかのパターンがある。

トランザクションのアノマリーについて詳しく学ぶにはトランザクションに関する本か何かを参照したほうが良さそう。

# 参考
- [dev.mysql.com - 15.7.2.1 トランザクション分離レベル](https://dev.mysql.com/doc/refman/8.0/ja/innodb-transaction-isolation-levels.html)
- [zenn.dev - MySQL/Postgres におけるトランザクション分離レベルと発生するアノマリーを整理する](https://zenn.dev/mpyw/articles/rdb-transaction-isolations#%E7%90%86%E8%AB%96%E9%9D%A2%E3%81%AE%E5%89%8D%E6%8F%90%E7%9F%A5%E8%AD%98)


