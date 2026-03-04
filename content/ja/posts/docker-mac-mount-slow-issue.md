---
title: "Docker for Macのマウントが遅い問題の対応"
slug: "docker-mac-mount-slow-issue"
date: 2018-08-19
author: bmf-san
categories:
  - "インフラストラクチャ"
tags:
  - "Docker"
  - "Tips"
draft: false
---

# 概要
Docker for Macのマウントが遅い。
npmとかスロー過ぎて辛い。
メモ書き。

# 原因
Dockerのスタッフの方のコメントを参照。（リンク先中段）

[Docker - File access in mounted volumes extremely slow, CPU bound](https://forums.docker.com/t/file-access-in-mounted-volumes-extremely-slow-cpu-bound/8076/158)

MacOSのファイルシステムのAPIが関連しているらしい。

# 解決策
- WindowsやLinuxを使う
- docker-sync
- MacOSとは異なるOSで構築した仮想環境を用意する（Vagrantとか）
- cached、delegated、consistentといったオプションを活用する（[Docker - Performance tuning for volume mounts (shared filesystems)](https://docs.docker.com/docker-for-mac/osxfs-caching/#performance-implications-of-host-container-file-system-consistency)）

# 所感
Linux使いたいなという気持ちになった。
