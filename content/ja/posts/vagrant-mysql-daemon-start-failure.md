---
title: Vagrant上のMySQLで"MySQL Daemon failed to start
description: 'Vagrant上のMySQLで"MySQL Daemon failed to start'
slug: vagrant-mysql-daemon-start-failure
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - MySQL
  - Vagrant
  - Tips
translation_key: vagrant-mysql-daemon-start-failure
---


LaravelでマイグレーションやSQLファイルのインポートを繰り返したり、中断したりと無茶をやっていたらMySQLがおかしくなりました。

# 対応
MySQLの設定ファイルにログサイズを設定したら直りました。

`innodb_log_file_size=5M`

# 所感
無茶はやめよう！


# 参考
+ ~~"Plugin 'InnoDB' registration as a STORAGE ENGINE failed" というエラーで MySQL サービスを起動できません
症状~~

+ [innodb_log_file_sizeを気軽に変えると死ぬよ](http://masasuzu.hatenadiary.jp/entry/2014/06/13/innodb_log_file_size%E3%82%92%E6%B0%97%E8%BB%BD%E3%81%AB%E5%A4%89%E3%81%88%E3%82%8B%E3%81%A8%E6%AD%BB%E3%81%AC%E3%82%88)

