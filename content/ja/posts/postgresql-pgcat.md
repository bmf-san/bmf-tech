---
title: "PostgreSQLとPgCatをローカルで素振りできる環境を作った"
slug: "postgresql-pgcat"
date: 2024-09-15
author: bmf-san
categories:
  - "インフラストラクチャ"
tags:
  - "負荷試験"
  - "PgCat"
  - "PostgreSQL"
  - "Grafana"
  - "Prometheus"
draft: false
---

PostgreSQLとPgCatをローカルで実験できる環境を作った。

[bmf-san/postgresql-pgcat-example](https://github.com/bmf-san/postgresql-pgcat-example)

PostgreSQLやPgCatのパラメータをチューニングしてローカルで負荷検証することできるようになっている。

構成は、

- Web
  - Rubyで雑APIサーバー
- PostgreSQL
  - MySQLも同じだが、volumeにデータが存在するとinit.sqlが走らない罠に小一時間ハマった。ドキュメント見るとちゃんと書いてあるので次回から気をつける...
- PgCat
- Prometheus
  - postgres_exporterのメトリクス収集
  - PgCatのメトリクス収集
    - PgCatはexporterが内包されている
- [postgres_exporter](https://github.com/prometheus-community/postgres_exporter)
- Grafana
  - メトリクスの可視化
  - Prometheusで必要なやつだけピックアップしてみるでも十分かも
- Locust
  - Pythonでシナリオを書いて負荷試験できるツール
  - WorkerコンテナをScaleさせれば並列リクエスト数を増やすことができる
  - 便利で手軽に使える。UIもいい感じで良い。OSSで負荷試験するならこれが良さげで気に入ったので、個人プロジェクトでも使ってみようと思った

となっている。

PgBouncerも扱えるようにしたかったのだが、設定が面倒になって途中で断念した。


コネクションプーリングが性能に与える影響をいい感じに観測できるようにしたいという動機があったのだが、まだいい感じに検証できていない。

あとはPostgreSQLやPgCatのパラメータチューニングが性能に与える影響なども観測できるといいかなと思っているが、まだ満足にやれていない。（環境を色々調整していたら力尽きてしまった..orz）






