---
title: "M1 MacでMySQL8.0.17のコンテナが起動できない"
slug: "mysql-container-not-starting-on-m1-mac"
date: 2021-09-05
author: bmf-san
categories:
  - "データベース"
tags:
  - "Docker"
  - "Docker Compose"
  - "MySQL"
  - "M1"
  - "Tips"
draft: false
---

M1 Macに買い替えたので手元の開発環境のためにmysqlコンテナを動かそうとしたら動かなかった。


エラーはこんな感じ。

```sh
runtime: failed to create new OS thread (have 2 already; errno=22)
```

goのエラーだったので、アーキテクチャの何かしらの問題で動いていないのだろうと推測。

とりあえずdocker hubを見て8.0.17より最新のバージョンを探してみると8.0.26の最新パッチバージョンまでリリースされているのを確認できた。

https://hub.docker.com/layers/mysql/library/mysql/8.0.26/images/sha256-75e71ac9b332935f396d85965213a64f1bd6fc7c42e9900b106f7af462c599b0?context=explore

ちょうど2日前にリリースされたらしい。

MySQL8.0.26でM1の対応が入った？ぽいので多分これで動くのでは。
cf. https://dev.mysql.com/doc/relnotes/mysql/8.0/en/news-8-0-26.html

```
macOS: It is now possible to build MySQL for macOS 11 on ARM (that is, for Apple M1 systems). (Bug #32386050, Bug #102259)
```

--platoform指定して、イメージは8.0.26指定。

```yaml
FROM --platform=linux/amd64 mysql:8.0.26

ADD ./conf.d/my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]
```

とりあえずこれで動いた。
