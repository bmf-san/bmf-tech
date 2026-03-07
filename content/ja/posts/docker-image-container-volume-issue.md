---
title: dockerでイメージとコンテナを削除してもボリュームが消えていなかったときのメモ
slug: docker-image-container-volume-issue
date: 2019-04-28T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Docker
  - Docker Compose
translation_key: docker-image-container-volume-issue
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


# 参考
- [DockerのVolumeに関して -v --rm -d ゴミが残る問題 コンテナが起動しない](https://stlisacity.hatenablog.com/entry/2018/09/10/145101)
