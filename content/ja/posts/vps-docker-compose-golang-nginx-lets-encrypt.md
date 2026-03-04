---
title: "VPS×Docker Compose×Docker Machine×Golang×Nginx×Let's Encryptでの本番運用"
slug: "vps-docker-compose-golang-nginx-lets-encrypt"
date: 2020-06-07
author: bmf-san
categories:
  - "ポエム"
tags:
  - "Docker"
  - "Docker Compose"
  - "Golang"
  - "Let's Encrypt"
  - "Nginx"
  - "Docker Machine"
  - "VPS"
draft: false
---

# 概要
GolangのアプリケーションをVPSでDocker-Composeを使って本番運用してみたかったので、トライ
してみた。

# 環境
実際に自分がトライした環境をまとめておく。

- VPS（Conoha）
- お名前ドットコム（ドメイン管理）
- Docker Compose（本番とローカルの構成を別ファイルで構築）
- Docker Machine（デプロイで使う）
- Let's Encrypt（TLS/SSL）
- Nginx（リバースプロキシ）

# リポジトリ
サンプルをつくった。
[github - bmf-san/go-production-boilerplate](https://github.com/bmf-san/go-production-boilerplate)

本番環境のサーバーではユーザー作成とかポートの開放くらいやっておけば、とりあえずデプロイできるはず・・・

ちなみにデプロイでダウンタイムが発生してしまうのでそちらは別途考慮が必要。

docker-machineを使ったデプロイについては、こちらの記事がわかりやすいので参考にした。
[Qiita - Docker MachineでMacからVPS上のDockerへアプリをデプロイしよう](https://qiita.com/momotaro98/items/5b902afea3530b6f0b93)

Let's Encrypt周りはちょっとハマったが、コンテナだからハマるという部分ではないところでハマったので、特に解決が難しい問題ではなかった。

# 所感
docker-machineのgenericドライバーを使えば気軽にデプロイできる。ダウンタイムの対策が必要かなと思うのが、プライベートのアプリケーションの運用であれば、考慮の1つになるかなと思う。
