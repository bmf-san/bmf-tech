---
title: "ISUCON環境で学ぶシステムメトリクス入門"
slug: "isucon"
date: 2024-04-14
author: bmf-san
categories:
  - "アーキテクチャ"
tags:
  - "ISUCON"
draft: false
---

# 概要　
ISUCON環境を利用してシステムメトリクスをちゃんと見れるようなろうという勉強会を定期的に行っているので、そのまとめを残す。

[Webエンジニアが知っておきたいインフラの基本 インフラの設計から構成、監視、チューニングまで](https://book.mynavi.jp/ec/products/detail/id=33857)を参考図書とし、第5章以降の内容を実際に手を動かして確認するような形で行っている。

環境はConohaでISUCON8のイメージを利用してサーバーを立てている。

プラン：メモリ 512MB/CPU 1Core

```sh
[root@160-251-16-96 ~]# cat /etc/redhat-release 
CentOS Linux release 7.5.1804 (Core) 
```

# システムの現状確認
cf. 第5章

## iptables
サーバーの待受ポートとファイアウォールの設定を確認して、開放されているポートを把握する。

```sh
iptables -nv -L
```

今回の環境では22と80が開放されていることが確認できた。

ubuntuでプリインストールされているufw（Uncomplicated FireWall）はiptablesのラッパー。

## ss
続いて、外部からの待受ポートを確認する。

ネットワークの状態確認コマンドであるss（旧netstat）を使って確認。

```sh
ss -lnp
```

h2oが80でIPアドレスなしで待受、isuconアプリケーションが8080でIPアドレスなしで待受、mysqldが3306でIPアドレスなしで待受、sshdが22でIPアドレスなしで待受しているのが確認できた。

Local Address:Portの項目で最初に::がついているものはIPv6でも待受していることになる。

ssの代用でlsofを使うでも良いと思う。

```sh
lsof -i
lsof -i:ポート番号
```

## ps
プロセスを見て起動コマンドを確認。

```
ps aufx | grep -v grep | grep -e isucon -e h2o -e mysql
```

isuconがisuconユーザーで`/home/isucon/torb/webapp/perl/local/bin/plackup`、h2oがrootで`perl -x /usr/share/h2o/start_server --pid-file=/var/run/h2o/h2o.pid --log-file=/var/log/h2o/error.log --port=0.0.0.0:80 -- /usr/sbin/h2o -c /etc/h2o/h2o.conf`、mysqlがmysqlユーザーで`/bin/sh /usr/bin/mysqld_safe --basedir=/usr`で起動されていることが確認できた。

## df
ディスク使用量の確認。

```sh
df -h
```

ディスク容量は30Gで22%使用されていることが確認できた。

こうすると容量を使っているディレクトリをリストアップできる。

```sh
df -sh /*
```

## top
CPU利用率、メモリ利用量、CPU使用率が高いプロセスを確認。

```sh
top -b -d 1 -n 1
```

## dstat
CPU使用率、ネットワーク利用量、ディスクI/O量、ページング量と推移を確認。

```sh
dstat -taf 1 10
```

ベンチマーカー実行しながら確認するとディスクI/Oに負荷がかかっていることが確認できた。

# ステータスモニタリング
// TODO:: 随時更新中



