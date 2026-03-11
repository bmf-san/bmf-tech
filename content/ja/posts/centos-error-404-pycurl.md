---
title: 'CentOS6.7で[Errno 14] PYCURL ERROR 22 - \"The requested URL returned error: 404 Not Found\"とかいうエラーがでた'
slug: centos-error-404-pycurl
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - CentOS
  - os
  - zabbix
  - Tips
translation_key: centos-error-404-pycurl
---


# CentOS6.7で[Errno 14] PYCURL ERROR 22 - "The requested URL returned error: 404 Not Found"とかいうエラーがでた

zabbixを導入しようと色々試行錯誤していたら以下のようなエラーがでてyumが使えなくなりました。


```
http://mirror.centos.org/centos/6/SCL/x86_64/repodata/repomd.xml: [Errno 14] PYCURL ERROR 22 - "The requested URL returned error: 404 Not Found"
他のミラーを試します。
エラー: Cannot retrieve repository metadata (repomd.xml) for repository: scl. Please verify its path and try again
```


こうなってしまってはOSに疎い私は絶望しましたが、以下のリンクが参考になり、解決に至ったのでシェアします。


[[tips][Linux]旧バージョンCentOSでyum更新できなくなった時](http://luozengbin.github.io/blog/2015-08-29-%5Btips%5D%5Blinux%5D%E6%97%A7%E3%83%90%E3%83%BC%E3%82%B8%E3%83%A7%E3%83%B3centos%E3%81%A7yum%E6%9B%B4%E6%96%B0%E3%81%A7%E3%81%8D%E3%81%AA%E3%81%8F%E3%81%AA%E3%81%A3%E3%81%9F%E6%99%82.html)

ちなみにzabbixはphpの設定周りで面倒なことになって結局インストールできていません。。。


サーバー監視ツールとかプロファイラとかサーバーにインストールして使う系のプログラムはインストールの段階で敷居が高いです。。。

