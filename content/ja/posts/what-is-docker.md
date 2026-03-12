---
title: "Dockerとは？コンテナ入門の完全ガイド"
slug: what-is-docker
date: 2018-04-01T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Docker
  - 仮想環境
translation_key: what-is-docker
---


# Dockerとは

- Docker社が開発している、コンテナ型の仮想環境を作成、配布、実行するためのプラットフォーム
- Linuxのコンテナ技術を使用

  - コンテナは、ホストマシンのカーネルを利用し、プロセスやユーザーなどを隔離する

    - 軽量・高速

- ミドルウェアや各種環境設定をコード化して管理できる（=Infrastructure as Code）

  - ローカル・本番環境問わず

    - 誰でも同じ環境が作れる
    - 環境の再配布・再利用が容易

- Dockerの正体

  - Linux Container(LXC)
  - LXCにおけるコンテナのカスタマイズ（設定ファイル作成やシェルスクリプトによるファイルインストール作業など）を楽に行うためのコンテナ管理ツール

- Docker for Mac
 
  - Macにデフォルトで入っているHyperKitという仮想化ツールで仮想マシンを立ち上げ、Linuxを起動してDockerを使えるようにしている

# コンテナ（Linuxコンテナ）とは

- システムのその他の部分から分離された一連のプロセス

  - プロセスのサポートに必要な全てのファイルを提供する個別のイメージから実行される

- OSとカーネルを共有し、アプリケーションプロセスをシステムの他の部分から独立させる（単一のOSで実行される）

# 他の仮想環境とコンテナの違い

## 一般的なPC

- 構成

  ```
  [ホストOS]
  [ハードディスク]
  ```

## ホストOS型（広義の意味でホストOSを使うハイパーバイザ型に定義されることもある）

- 構成

  ```
  [ゲストOS]
  [仮想化ソフト]
  [ホストOS]
  [ハードディスク]
  ```

- メリット

  - 手軽に仮想化を実現できる
  - OS選択の自由度が高い

- デメリット

  - ディスクやメモリの消費多い

## ハイパーバイザ型（ハードディスクを用いるパターン）

- 構成

  ```
  [ゲストOS]
  [ハイパーバイザ]
  [ハードディスク]
  ```

- メリット

  - ホストOSの処理を必要とせず直接ハードウェアを制御できるため処理速度が速い

- デメリット

  - 手軽に仮想環境を実現できない（ホストOSをそのまま使えなかったり、専用の物理サーバーを必要とする場合がある）

## コンテナ型

```
[コンテナ管理ソフトウェア]
[ホストOS]
[ハードディスク]
```

- メリット

  - 手軽に仮想環境を実現できる
  - ディスクやメモリの消費が少ない

- デメリット

  - Linuxカーネルを使ったOSしか使えない

# Dockerのイメージとコンテナの概要

```
                Docker repository(Ex. Docker Hub)

                         ↓ (pull)

Dockerfile   →     Docker Image    →    Docker Container
                 (build)                           (run)   
                                      ↓ (commit)

                             Docker Image    →    Docker Container
                                                       (run)
```

※ホストOS型とハイパーバイザ型の違い
[Think IT - ホスト型とハイパーバイザー型の違いは何？VMware vSphere Hypervisor の概要](https://thinkit.co.jp/story/2012/10/17/3722)が参考になる。

- ホストOS型
    - OS上に仮想化ソフトウェアをインストールする
    - ハードウェアへのアクセスはホストOSを経由するので、オーバーヘッドが発生し、パフォーマンスが大きく発揮されない。
- ハイパーバイザ型
    - ホストOSを必要とせず、サーバーへ直接インストールする。
    -  ハードウェアを直接制御できるので、パフォーマンスが出やすい。


# 基本コマンド
- `docker build`
   - Dockerfileからイメージを作成

- `docker run`
   - イメージからコンテナを起動

- `docker commit`
   - コンテナからイメージを作成

# 補足

- カーネルとは

  - OSの基本機能の役割を担うソフトウェア

    - カーネルだけではOSの利用が困難なので別途ソフトウェアと組み合わせて利用するのが一般的

  - カーネル＋ソフトウェア = ディストリビューション

    - CentOSやUbuntu→Linuxカーネル＋ソフトウェアのディストリビューション

# 参考

- [Docker公式サイト](https://www.docker.com/what-docker)
- [Docker入門（第一回）～Dockerとは何か、何が良いのか～](https://knowledge.sakura.ad.jp/13265/)
- [Linuxコンテナとは](https://www.redhat.com/ja/topics/containers/whats-a-linux-container)
- [サーバーの仮想化とは？仕組み、メリット・デメリットをわかりやすく解説します](https://www.kagoya.jp/howto/rentalserver/virtualization/)
- [15分で分かるLXC（Linux Containers）の仕組みと基本的な使い方 エンジニア向け 2014.06.16](https://knowledge.sakura.ad.jp/2108/)
- [LXCを使った権限分離とテンプレートのカスタマイズ](https://knowledge.sakura.ad.jp/2163/)
- [Dockerのイメージとコンテナの概要と各種コマンド（随時更新するかも）](https://yoshinorin.net/2016/10/03/docker-image-and-container-command/)
- [Linux コンテナの内部を知ろう / OSC 2018 Kyoto](https://speakerdeck.com/tenforward/osc-2018-kyoto)
- [Think IT - ホスト型とハイパーバイザー型の違いは何？VMware vSphere Hypervisor の概要](https://thinkit.co.jp/story/2012/10/17/3722)
