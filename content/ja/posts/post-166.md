---
title: "コンテナ技術概要"
slug: "post-166"
date: 2023-06-05
author: bmf-san
categories:
  - "インフラストラクチャ"
tags:
  - "Docker"
  - "libcontainer"
  - "lxc"
  - "lxd"
  - "コンテナ"
draft: false
---

# 概要
コンテナ技術についてのまとめ。
Dockerを使わずにコンテナをつくって触ってみる。

# コンテナとは
- ホストOSからアプリケーションとランタイムをまとめて、分離した一連のプロセス。

# コンテナの歴史
1979年 UNIX OSにchrootが登場。

2000年 [FreeBSD jails](https://www.freebsd.org/doc/handbook/jails.html)がFreeBSD 4.0に登場。chrootの発展系。

2001年 [VServer Project](http://linux-vserver.org/Welcome_to_Linux-VServer.org)を通じてLinuxにもLinuxコンテナのベースとなる技術が登場。

2004年、LXC1.0がリリース。
[Linux Containers](https://linuxcontainers.org/)

2008年、Dockerが登場

コンテナ技術は上記以外にもVirtuozzo、OpenVZ、HP-UX Container、Solaris Containerなど存在する。

# コンテナと仮想化の違い
- コンテナ
    - ホストOSからアプリケーションとランタイムをまとめて分離した一連のプロセス。
    - ホストOSのカーネル部分を共有している
        - OSのライブラリ部分はコンテナ側が選択可能 
  
- 仮想化
    - ホスト型とハイパーバイザ型で構成が異なるが、仮想化は複数OSを用意できるような構成で、ゲストOS上でアプリケーションを実行する。

[bmf-tech - Dockerとは](https://bmf-tech.com/posts/Docker%E3%81%A8%E3%81%AF)にもざっくりまとめている。

## コンテナ技術を実現するためのLinuxカーネルの機能
### Kernel namespaces
- プロセスを6種類のリソース(ipc, uts, mount, pid, network, user)に分離する機能
- ユーザーがユーザー専用の分離されたリソースを持っているように見せる仕組み。
- 分離されたリソース同士は互いに干渉できない。

### Apparmor and SELinux profiles
- Apparmor
    - Linux Security Modules（Linuxカーネルにあるセキュリティのためのフレームワーク）の一種。
    - アプリケーションのアクセス権限をセキュアに管理（強制アクセス制御）
- SELinux（Security Enhanced Linux）
    - Linuxカーネルに強制アクセス制御機能を加えるモジュール

### Seccomp policies
- プロセスのシステムコール発行を制限する機能

### Chroots(using pivot_root)
-  現在のプロセスとその子プロセス群に対して、ルートディレクトリを変更する操作のこと
- ルートを変更されたプロセスは範囲外のファイルにアクセスできなくなる=>プロセス分離の実現

### Kernel capabilities
- プロセスの権限管理
- root or not rootよりももっと細かい権限管理ができる

### CGroups(control groups)
- プロセスを共通管理するために、プロセスをグループ化する機能

# Dockerのコンテナ技術
以前までDockerはlxcを使っていたが、v0.9からgoで実装されたlibcontainerを使っているらしい。（cf. [Docker blog - DOCKER 0.9: INTRODUCING EXECUTION DRIVERS AND LIBCONTAINER](https://blog.docker.com/2014/03/docker-0-9-introducing-execution-drivers-and-libcontainer/) [github - opencontainers/runc/libcontainer/](https://github.com/opencontainers/runc/tree/master/libcontainer)）

# 標準仕様
## OCI(Open Container Initiative)
[Open Container Initiative](https://opencontainers.org/)はコンテナとランタイムに関する業界標準の作成を目的として組織。

以下の仕様を定義している。

- OCI Runtime Specification
- OCI Image Format Specification
- OCI Distribution Specification

OCIはローレベルランタイムの仕様に関わっている。
ex. runC、gVisor、Kata Containers、Nabla Containers etc...

## CRI(Container Runtime Interface)
[CRI](https://kubernetes.io/ja/docs/concepts/architecture/cri/)は、kubeletとコンテナランタイム間の通信のインタフェースを規定している。

CRIはハイレベルランタイムの仕様に関わっている。
ex. docker、containerd、cri-o

# まとめ
- コンテナはリソースを分離されたプロセス
- コンテナはホストOSのカーネル部分を共有、ライブラリ部分は自由に選択できる
- コンテナに関連する仕様としては、OCIとCRIがある

# LTした
Makuake LT Party（社内LT大会）にてLTをした。

[speaker-deck - コンテナ完全に理解した](https://speakerdeck.com/bmf_san/kontenawan-quan-nili-jie-sita)

# 参考
- [bmf-tech - Dockerとは](https://bmf-tech.com/posts/Docker%E3%81%A8%E3%81%AF)
- [[コンテナの歴史]Dockerができるまで 第二回 〜集合知を集めて歴史を知ろう〜](https://hackmd.io/s/ryPfDLU77)
- [redhat - Linux コンテナとは](https://www.redhat.com/ja/topics/containers/whats-a-linux-container)
- [redhat - Linuxコンテナとは何か](https://www.redhat.com/ja/topics/containers)
- [Linux Container](https://linuxcontainers.org/ja/)
- [ITソリューション塾 - 【図解】コレ１枚で分かるコンテナ型仮想化とDocker](http://blogs.itmedia.co.jp/itsolutionjuku/2015/05/docker.html)
- [SELinux Project Wiki](http://selinuxproject.org/page/Main_Page)
- [opensuse - AppArmor](https://ja.opensuse.org/AppArmor)
- [kernel.org - SECure COMPuting with filters](https://www.kernel.org/doc/Documentation/prctl/seccomp_filter.txt)
- [man7.org - Linux Capabilities](http://man7.org/linux/man-pages/man7/capabilities.7.html)
- [gihyo.jp - LXCで学ぶコンテナ入門 －軽量仮想化環境を実現する技術](https://gihyo.jp/admin/serial/01/linux_containers/0001)
- [ゆううきブログ - 自作Linuxコンテナの時代](https://blog.yuuk.io/entry/diy-container)
- [Think IT - コンテナ技術の基礎知識](https://thinkit.co.jp/story/2015/08/11/6285)
- [Linux Containers - LXDとは？](https://linuxcontainers.org/ja/lxd/introduction/)
- [Hewlett Packard Enterprise - Dockerコンテナと仮想化の違いとは？SynergyとDevOps   ](https://community.hpe.com/t5/Enterprise-Topics/Docker%E3%82%B3%E3%83%B3%E3%83%86%E3%83%8A%E3%81%A8%E4%BB%AE%E6%83%B3%E5%8C%96%E3%81%AE%E9%81%95%E3%81%84%E3%81%A8%E3%81%AF-Synergy%E3%81%A8DevOps/ba-p/6980068?profile.language=ja#.XD6Zks8zZTY)
- [www.publickey1.jp - コンテナランタイムの仕組みと、Firecracker、gVisor、Unikernelが注目されている理由。 Container Runtime Meetup #2
](https://www.publickey1.jp/blog/20/firecrackergvisorunikernel_container_runtime_meetup_2.html)
- [thinkit.co.jp - Kubernetes 1.20から始まるDockerランタイムの非推奨化に備えよう！我々が知っておくべきこと・すべきこと](https://thinkit.co.jp/article/18024)
- [container-security.dev - Container Security Books](https://container-security.dev/)
- [github.com - opencontainers/runtime-spec](https://github.com/opencontainers/runtime-spec/blob/main/spec.md)
- [udzura.hatenablog.jp - OCI Runtime Specificationを読む](https://udzura.hatenablog.jp/entry/2016/08/02/155913)
- [medium.com - コンテナユーザなら誰もが使っているランタイム「runc」を俯瞰する[Container Runtime Meetup #1発表レポート]](https://medium.com/nttlabs/runc-overview-263b83164c98)
- [Docker一強の終焉にあたり、押さえるべきContainer事情](https://zenn.dev/ttnt_1013/articles/f36e251a0cd24e)
- [gkuga.hatenablog.com - OCI Runtime Specificationを読んだので概要を書く](https://gkuga.hatenablog.com/entry/2020/01/24/032122)
- [yohgami.hateblo.jp - chrootとunshareを使い、シェル上でコマンド7つで簡易コンテナ](https://yohgami.hateblo.jp/entry/20161215/1481755818)
- [コンテナ技術入門 - 仮想化との違いを知り、要素技術を触って学ぼう](https://eh-career.com/engineerhub/entry/2019/02/05/103000)
- [dockerコマンドを使わずにコンテナを作る - 1](https://zenn.dev/chemimotty/articles/51788231854a5e)
- [kaminashi-developer.hatenablog.jp - 【Go言語】自作コンテナ沼。スクラッチでミニDockerを作ろう
](https://kaminashi-developer.hatenablog.jp/entry/dive-into-swamp-container-scratch)
- [www.youtube.com - Building a container from scratch in Go - Liz Rice (Microscaling Systems)
](https://www.youtube.com/watch?v=Utf-A4rODH8)
- [medium.com - Understand the Design of Container Runtime](https://medium.com/@ikeda.morito/understand-the-design-of-containerruntime-eb79161545ef)
