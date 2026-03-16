---
title: DBドキュメント（ER図など）を自動生成してくれるツールーschemaspy, tbls
description: DBドキュメント（ER図など）を自動生成してくれるツールーschemaspy, tbls
slug: db-documentation-tools-schemaspy-tbls
date: 2020-07-09T00:00:00Z
author: bmf-san
categories:
  - データベース
tags:
  - Docker
  - ER
translation_key: db-documentation-tools-schemaspy-tbls
---


# 概要
DBドキュメントを自動生成できるツールの有名所でいうと、MySQL Workbenchが筆頭に上がると思うが、それ以外にも良いOSSがないか漁ってみてちょっと触ってみたので感想を残す。

# Schemaspy
- [schemaspy.org](http://schemaspy.org/)
- [github.com - schemaspy/schemaspy](https://github.com/schemaspy/schemaspy)

DBに接続してhtmlでDBドキュメントを生成してくれるJava製のツール。

Dockerhubにイメージがあるので、それを使って簡単に試してみることができる。

MySQL5.7（多分5.8も大丈夫だと思う・・）は、こんな感じでいけるはず。
`docker run -v "$PWD/schema:/output" --net="host" schemaspy/schemaspy:latest \
 -t mysql -host {DBHOST}:{DBPORT} -db {DBNAME} -u {DBUSER} -p {DBPASSWORD}`

MySQL5.6環境下ではコマンドをちょっといじる必要がある。
`docker run -v "$PWD/schema:/output" --net="host" schemaspy/schemaspy:latest -t mysql -host {DBHOST}:{DBPORT} -db {DBNAME} -u {DBUSER} -p {DBPASSWORD} -connprops useSSL\\=false -s {DBNAME}`

いずれもワンライナーでお試しできるので簡単。

もちろんmysql以外でもOK。

# tbls
- [github.com - k1Low/tbls](https://github.com/k1LoW/tbls)

CIフレンドリーなDBドキュメンテーションツールで、markdownでドキュメントを生成してくれる。

depでもrpmでもbrewでもgoでもdockerでもインストールできる。

使い方は簡単なのでgithubのREADME参照。

ドキュメントはmarkdownですべて管理したいので個人のアプリケーションのドキュメンテーションに採用している。

# 所感
- ER図だけで比較すると、結論はMySQL Workbenchが一番見やすいかもしれない。ただしテーブル数による。少ないテーブル数であれば
どれでも良さそう。
- MySQL WorkbenchのER生成はリレーション関係にないテーブルを自動で落としてくれないが、schemaspyはちゃんと落として生成してくれるぽい。
- schemaspyのER図はテーブル数が多いとリレーションの線が追いづらいので、UIでなんとか調整できると良いと思った。MySQL Workbenchは手動で調整できるのでそのへん柔軟。（自動整頓の機能はあるが、限界があるぽい...）
- ERについては全テーブルのリレーションを吐き出すのは物理的限界を感じるので、そこは使う側の工夫が必要なのかなぁと思う。一度に全テーブルを対象としてリレーションをみたいというケースはそんなにないと思うので、テーブルを絞るなどすれば良いと思う。schemaspyはERとして生成するテーブルを絞ったりできるのかな・・？ちゃんと見ていないがパッと見できなそう・・
- tblsはCIフレンドリーなのでCIに組み込みやすいと思うが、schemaspyの方も比較的に容易に組み込むことができそう。
- [rarejob-tech-dept.hatenablog.com - ER図 の自動作成を CI に組み込んだ話](https://rarejob-tech-dept.hatenablog.com/entry/2020/01/24/190000)


# 参考
- [ベジプロ - SchemaSpyでER図を出力できない](https://www.blog.v41.me/posts/749a3607-aa12-47d6-9441-8f7497602325)
- [sys-guard.com - SchemaSpyでER図を自動作成 on Docker
2019年7月22日 2019年7月22日 AWS, Linux/UNIX 129view
](https://sys-guard.com/post-17119/)
