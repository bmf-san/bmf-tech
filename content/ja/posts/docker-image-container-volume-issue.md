---
title: dockerでイメージとコンテナを削除してもボリュームが消えていなかったときのメモ
description: "Dockerボリューム削除問題を解説。イメージ・コンテナ削除でもボリューム残存する理由、docker volume rm、docker-compose down -v オプション活用を紹介します。"
slug: docker-image-container-volume-issue
date: 2019-04-28T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Docker
  - Docker Compose
translation_key: docker-image-container-volume-issue
draft: false
---


# 概要
dockerでコンテナとイメージを削除してもvolumeが削除されていなくてちょいちょい忘れてハマるのでメモっておく。

普段は、docker-composeを使っている。

`docker-compose build`
`docker-compose up -d`

して、

`docker rm **`
`docker rmi **`

という感じにお片付けしているのだが、どうやらマウントしているvolumeを削除するオプションがあったらしい。

# 対応
ボリュームが残っているか確認。
`docker volume ls`
　
`docker volume rm **`

# 余談
docker-composeでdocker-compose.ymlに記述されているコンテナとネットワークイメージとボリュームを一気に片付ける方法があった。

`docker-compose down --rmi all -v`

# Dockerボリュームの削除方法
イメージやコンテナを削除しても、ボリュームは残る。残ったボリュームは次の手順で削除する。

まず、残っているボリュームを確認する。

```
docker volume ls
```

次に、不要なボリュームを名前で指定して削除する。

```
docker volume rm <volume-name>
```

どのコンテナからも参照されていないボリュームをまとめて削除する場合に使う。

```
docker volume prune
```

Docker Composeでは、`down`に`-v`を付けると、定義されたボリュームも削除できる。

```
docker compose down -v
```

イメージまで含めて片付けるなら、`--rmi all`も付ける。

```
docker compose down --rmi all -v
```


# 参考
- [DockerのVolumeに関して -v --rm -d ゴミが残る問題 コンテナが起動しない](https://stlisacity.hatenablog.com/entry/2018/09/10/145101)
