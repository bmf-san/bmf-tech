---
title: Dockerizeを使ってDocker Composeのコンテナの起動順を制御する
description: Dockerizeを使ってDocker Composeのコンテナの起動順を制御するについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: dockerize-control-docker-compose-start-order
date: 2019-09-17T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Docker
  - Docker Compose
  - Dockerize
translation_key: dockerize-control-docker-compose-start-order
---


# 概要
Dockerizeを使ってDocker Composeで起動するコンテナの順番を制御する方法についての覚え書き。

[github.com - jwilder/dockerize](https://github.com/jwilder/dockerize)

# なぜDockerizeを使うのか
Dockerizeではなく、[wait-for-it](https://github.com/vishnubob/wait-for-it) というピュアなbash scriptを使った方法も採用することもできる。

cf. [Docker-docs-ja - Compose の起動順番を制御](http://docs.docker.jp/compose/startup-order.html)

Dockerizeを使う目的は複数コンテナを起動する際、コンテナの起動順を意図的に制御したいようなときである。

例えば、アプリケーション用のコンテナとテスト用のデータベースコンテナがあったとして、アプリケーション側のコンテナがDBを使用したテストを行うようなとき、データベースのコンテナがアプリケーションのコンテナよりも先に起動されている必要がある。

要は、コンテナ間の起動順の依存関係を解決するような目的であるかと思う。

docker-composeには`depends_on`や`links`といったオプションがあるが、`depends_on`はコンテナの作成順序を、`links`は`depends_on`の機能に加えてコンテナ間の名前解決を行うもので、どちらも起動の順番はまでは制御しない。

後で知ったが`links`はversion2以降では自動的に実行されるらしく、レガシーになったらしい。

# Example
筆者の某アプリを実例にあげておく。

docker-compose.yml

```yaml
version: "3"
services:
  gobel_app:
    container_name: "gobel_app"
    build: ./docker/go
    volumes:
      - ./app:/go/src/github.com/bmf-san/Gobel/app
    ports:
      - "8080:8080"
    depends_on:
      - gobel_db
      - gobel_test_db
    entrypoint:
      - dockerize
      - -timeout
      - 10s
      - -wait
      - tcp://gobel_test_db:3306
    command: realize start
  gobel_db:
    container_name: "gobel_db"
    build: ./docker/mysql
    ports:
      - "3306:3306"
    volumes:
      - mysql_gobel_db:/var/lib/mysql:delegated
      - ./docker/mysql/initdb.d/gobel_db:/docker-entrypoint-initdb.d
    environment:
      - MYSQL_DATABASE=gobel
      - MYSQL_ROOT_PASSWORD=password
  gobel_test_db:
    container_name: "gobel_test_db"
    build: ./docker/mysql
    ports:
      - "3305:3306"
    volumes:
      - mysql_gobel_test_db:/var/lib/mysql:delegated
      - ./docker/mysql/initdb.d/gobel_test_db:/docker-entrypoint-initdb.d
    environment:
      - MYSQL_DATABASE=gobel_test
      - MYSQL_ROOT_PASSWORD=password
volumes:
  mysql_gobel_db:
    driver: local
  mysql_gobel_test_db:
    driver: local
```

gobel_appコンテナはgobel_test_dbの起動を待ちたい。 
entrypointにはdockerizeのコマンドを指定している。

gobel_appコンテナのビルドに使用しているDockerfile。

```Dockerfile
FROM golang:1.13.0-alpine

WORKDIR /go/src/github.com/bmf-san/Gobel/app/

RUN apk add --no-cache git \
    binutils-gold \
    curl \
    g++ \
    gcc \
    gnupg \
    libgcc \
    linux-headers \
    make
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/oxequa/realize
RUN go get -u golang.org/x/lint/golint
ENV DOCKERIZE_VERSION v0.6.0
RUN apk add --no-cache openssl \
 && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz
```

dockerizeのインストールを含めるようになっている。

# 所感
手軽に導入できる。
wait-for-itのようなスクリプトを自前で用意しても良さそうな気がするがコンテナ管理が複雑化することを見込んでこうしたdockerizeを入れておくのはありかなと思った。

