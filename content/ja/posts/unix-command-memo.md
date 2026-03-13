---
title: Unixコマンドメモ
description: "Unixコマンド活用集、jqでJSON加工・teeで双方向出力・atでジョブスケジュール実行する実用的な例"
slug: unix-command-memo
date: 2018-07-07T00:00:00Z
author: bmf-san
categories:
  - OS
tags:
  - unixコマンド
  - jq
  - tee
translation_key: unix-command-memo
---


# 概要
Unixコマンドのメモ。

# jq
JSON形式のデータを加工するコマンド。

## JSONのPretty Print
```
echo '[{"name": "Tom", "age": 20}}]' | jq .
```

Pretty Print以外にもオブジェクトからプロパティを指定してデータを取り出したり、オブジェクトの長さを取得したり、色々な使い方がある。


# tee
標準入力を標準出力とファイルの両方に出力する。
sudoが使える。
オプションなしで上書き、オプション-aで追記。

```
echo 'hello world' | sudo tee ./sample.txt
```

リダイレクトの場合は、>が上書き、>>が追記。sudoは使えない。


# at
コマンドの実行時刻を予約することができる。

```
at -f ./sample.txt 2230
```

日時の部分はフォーマットが色々ある。


# mktemp
ランダムな名前で/tmpディレクトリ以下にファイルを作成する。

```
mktemp
```


# lsof
List of open files
プロセスがオープンしているファイルを出力する。

-i
ポート番号に絞って出力。

-i tcp or udp
TCPやUDPに絞って出力。

-P
ポート番号を数字で出力。

-n
IPアドレスをホスト名に逆引き変換せずに出力。

# nmap
ネットワーク経由で対象とするホストのポート状況を調べる。

```
nmap 192.168.33.10
```

# fsfeeze
ファイルシステムのI/Oを一時的に停止させるコマンド。

ファイルシステムのI/Oを一時的に停止する（フリーズ処理）。
```
fsfreeze -f /data
```

ファイルシステムのI/Oを解除する（アンフリーズ処理）。
```
fsfreeze -u /data
```

# findmnt
マウントしているファイルシステムの情報を出力する。




