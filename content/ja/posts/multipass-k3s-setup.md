---
title: multipassでk3sを動かす
description: multipassでk3sを動かすについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: multipass-k3s-setup
date: 2023-08-17T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - multipass
  - k3s
translation_key: multipass-k3s-setup
---


# 概要
docker-composeで構成されている個人開発のアプリケーションをKubernetes（k3s）へ移行するという試みの際に、multipassを使ってみたのでそれについてメモを残す。

結局移行はしなかったが...

[k3s](https://k3s.io/)はCNCFに認証されたKubernetesディストリビューション。IoTやエッジコンピューティング用途向け。メモリを節約したい、Kubernetesほどのスケールは不要、あるいは気軽にKubernetesを触ってみたいときなどにも有用なので、個人がVPSなどでKubernetesを導入してみたいときの選択肢にもなりうると思う。

cf. [K3s on ConoHa](https://qiita.com/yhirokw/items/fd5dcb28d3f57de0cc40)

k3sは機能的にはKurbenetesとほとんど変わらないが、いくつか制約がある。詳しくはドキュメント参照。

cf. [docs.k3s.io](https://docs.k3s.io/)

# multipassとは
Ubuntuの仮想環境を気軽に構築できるツール。Linux、macOS、Windowsに対応している。

[multipass.run](https://multipass.run/)

# なぜmultipassを使ったのか
k3sの実行環境として、macOS上で仮想環境を容易する必要があったため。

cf. [Can I install k3s on macos (big sur) with m1 chip?](https://www.reddit.com/r/kubernetes/comments/qa2f8d/can_i_install_k3s_on_macos_big_sur_with_m1_chip/)

いくつか代替手段はあるが、気軽に簡単に触れそうだったmultipassを使ってみた。

# multipassでk3sを動かす
macOSならbrewでmultipassをインストールして、以下のステップだけでk3sを動かすことができるようになる。

1. mutipass find // ubuntuイメージを探す
2. multipass launch -c 2 -m 4G -d 50G -n example 22.10 // 22.10はubuntuのバージョン
3. multipass mount ./k3s/ example:~/k3s // マウントする
4. multipass shell example // 仮想マシンに接続
5. curl -sfL https://get.k3s.io | sh - // k8sを仮想マシン上にインストール

cf.
- [M1 MacでUbuntuの仮想マシンを使うためにMultipassをインストールする](https://virment.com/how-to-install-multipass-to-m1-mac/#%E3%83%9B%E3%82%B9%E3%83%88%E3%83%9E%E3%82%B7%E3%83%B3%E3%81%AE%E3%83%87%E3%82%A3%E3%83%AC%E3%82%AF%E3%83%88%E3%83%AA%E3%82%92%E3%83%9E%E3%82%A6%E3%83%B3%E3%83%88%E3%81%99%E3%82%8B)
- [M1 Macでk3s](https://qiita.com/tkuribayashi/items/4eb664631254fa58df57)
- [M1 MacにおけるDocker Desktopを使わないMultipassを使ったKubernetes環境の構築](https://zenn.dev/kkoudev/articles/b001c36c7d7005)
