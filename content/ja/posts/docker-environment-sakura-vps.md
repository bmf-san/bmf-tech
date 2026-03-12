---
title: さくらVPSでDocker環境を構築する
description: さくらVPSでDocker環境を構築するの手順と実践例を詳しく解説します。
slug: docker-environment-sakura-vps
date: 2018-06-09T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Docker
  - さくらのVPS
translation_key: docker-environment-sakura-vps
---


# 概要
さくらVPS上にDocker環境を構築する。
サーバーの初期設定等は割愛。

# 構築手順
操作はすべてsudo権限を持った一般ユーザーで行うものとする。

Dockerには無償のCE版と商用版のEE版があるが、今回はCE版を使用する。

# リポジトリをセットアップ

## インストール

```
sudo yum install -y yum-utils \
device-mapper-persistent-data \
lvm2
```

## stableリポジトリのセットアップ

```
sudo yum-config-manager \
--add-repo \    https://download.docker.com/linux/centos/docker-ce.repo
```

## edgeとtestリポジトリのセットアップ

```
sudo yum-config-manager --enable docker-ce-edge
```

```
sudo yum-config-manager --enable docker-ce-test
```

今回はstableだけ使いたいので`--disable`で無効化しておく。

```
sudo yum-config-manager --disable docker-ce-edge
```

```
sudo yum-config-manager --disable docker-ce-test
```

# Docker CEのインストール

```
sudo yum install docker-ce
```

以下のコマンドでインストール可能なバージョンを確認できる。

```
yum list docker-ce --showduplicates | sort -r
```

指定したバージョンをインストールするには次のようにバージョンを指定してインストールする。

```
sudo yum install docker-ce-<VERSION STRING>
```

# Dockerの起動

```
sudo systemctl start docker
```

起動しているか確認。

```
sudo docker run hello-world
```

# Docker CEのアンインストール

```
sudo yum remove docker-ce
```

dockerのイメージやボリューム、コンテナ等や設定ファイルなどは自動で削除されないので、以下のディレクトリを手動で削除する。

```
sudo rm -rf /var/lib/docker
```

# 参考
- [docker docs - Get Docker CE for CentOS](https://docs.docker.com/install/linux/docker-ce/centos/)


