---
title: Ubuntu 20.04.2 LTSでmysqlコンテナを起動するとPermission deniedで怒られた
description: Ubuntu 20.04.2 LTSでmysqlコンテナを起動するとPermission deniedで怒られたについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: permission-denied-when-starting-mysql-container-on-ubuntu
date: 2021-09-12T00:00:00Z
author: bmf-san
categories:
  - データベース
tags:
  - Ubuntu
  - MySQL
  - Tips
translation_key: permission-denied-when-starting-mysql-container-on-ubuntu
---


# 概要
Ubuntu 20.04.2 LTSでmysqlコンテナを起動しようとと以下のようなエラーが出てコンテナ起動に失敗する。

```sh
Could not open file '/var/log/mysql/mysql-error.log' for error logging: Permission denied”
```

# dockerfile
問題が発生したdockerfile。

docker-compose.yml（一部抜粋）

```yml
version: '3.2'
services:
  mysql:
    container_name: "example-mysql"
    env_file: ./mysql/.env
    build:
        context: "./mysql"
        dockerfile: "Dockerfile"
    ports:
      - "3306:3306"
    volumes:
      - ./mysql/data:/var/lib/mysql
      - ./mysql/initdb.d:/docker-entrypoint-initdb.d
      - ./mysql/log:/var/log/mysql
```

Dockerfile
```yml
FROM --platform=linux/amd64 mysql:8.0.26

ADD ./conf.d/my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]
```

# マウント元を確認してみる
```sh
$ ls -la mysql
drwxrwxrwx  7 systemd-coredump example-app 4096 Sep 12 23:10 data
```

`systemd-coredump`という見慣れないユーザーが。

# systemd-coredump
ホストでユーザーを確認すると、`systemd-coredump`はuid 999。

```sh
$ cat /etc/passwd | grep systemd-coredump
systemd-coredump:x:999:999:systemd Core Dumper:/:/usr/sbin/nologin
```

mysqlのコンテナ内のユーザーがuid 999を持っているのがおそらく原因？

# 対応
docker-compose.ymlに`user: 1000:1000`を追加。

docker-compose.yml（一部抜粋）
```yml
version: '3.2'
services:
  mysql:
    container_name: "example-mysql"
    env_file: ./mysql/.env
    build:
        context: "./mysql"
        dockerfile: "Dockerfile"
    ports:
      - "3306:3306"
    volumes:
      - ./mysql/data:/var/lib/mysql
      - ./mysql/initdb.d:/docker-entrypoint-initdb.d
      - ./mysql/log:/var/log/mysql
    user: 1000:1000
```

ハードコードしないでホストからuidとgidを渡すようにしたほうが良い気はする。

# 所感
docker for macではこの問題は発生していなかったので、気づくことができてよかった。

# 参考
- [gitmemory.com - redis uid issue on ubuntu20.04](https://www.gitmemory.com/issue/docker-library/redis/279/833602815)
- [tsaeki.hatenablog.com - https://tsaeki.hatenablog.com/entry/2020/02/29/182620](https://tsaeki.hatenablog.com/entry/2020/02/29/182620)
- [stackoverflow.com - Permission issue with PostgreSQL in docker container](https://stackoverflow.com/questions/56188573/permission-issue-with-postgresql-in-docker-container)
