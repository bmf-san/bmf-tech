---
title: "DockerのOperation not permittedというエラー\bに対応する"
slug: docker-operation-not-permitted-error
date: 2019-09-27T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Docker
  - Docker Compose
  - Linux
  - linux capabilities
  - seccomp
  - Tips
translation_key: docker-operation-not-permitted-error
---


# 概要
Docker Composeを使ってgolangのtest実行していたら、**Operation not permitted**というエラーに遭遇した。

# 対応
[Docker Documentation - runtaime-privilege-and-linux-capabilities](https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities)

Dockerコンテナの特権設定をいじると解決する。

```yml
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
    privileged: true  // add a this option
```

上記だとセキュリティ的にどうなのかイマイチよくわかっていないので、もう少し権限を絞るような設定にした。

```yml
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
    cap_add:
      - SYS_ADMIN
    security_opt:
      - seccomp:unconfined
```

`cap_add`はLinux capabilitiesを追加するオプションで、ここではsystem administration operationsの権限を追加している。

Linux capabilitiesはスーパーユーザーの権限を細分化する機能。

`seccomp`はLinuxカーネルのシステムコール発行の制限をするセキュリティ関連の機能。

ここでは、unconfined、無効化の設定をしている。

unconfinedは直訳すると"監禁されていない"という意味らしい。

# 所感
[speakerdeck - コンテナ完全に理解した](https://speakerdeck.com/bmf_san/kontenawan-quan-nili-jie-sita)

コンテナことは前にちらっと勉強したが、まだまだ理解が浅い。

# 参考
- [Qiita - systemd in docker container without --privileged](https://qiita.com/shusugmt/items/92ece6874ba5aeff2b41)
