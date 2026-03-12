---
title: Mackerelで監視対象から外れたホストを再度復帰させる方法
description: Mackerelで監視対象から外れたホストを再度復帰させる方法の手順と実践例を詳しく解説します。
slug: mackerel-recover-excluded-hosts
date: 2019-09-17T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Mackerel
  - Tips
translation_key: mackerel-recover-excluded-hosts
---


# 概要
Mackerelの監視対象から誤ったホストを退役させてしまったり、意図せず監視対象から外してしまったときなどに復帰させる方法についてメモっておく。

# 手順
mackerel-agentを再起動しても自動で復帰しないのでホストに割り振られたhostIdを更新させる必要がある。

```sh
service mackerel-agent stop
cd /var/lib/mackerel-agent
mv id /tmp/
service mackerel-agent start
```

# 所感
たまに焦るので気をつけたい...
