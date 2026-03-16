---
title: dockerコマンドをcronで実行しようとしたらthe input device is not a TTY
description: dockerコマンドをcronで実行しようとしたらthe input device is not a TTY
slug: docker-cron-tty-issue
date: 2023-03-17T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Docker
translation_key: docker-cron-tty-issue
---


# 概要
dockerコマンドをcronで実行しようとしたら"the input device is not a TTY"と怒られてしまった。

cronに設定しようとした内容例は以下。

```sh
* * * * * user docker exec -it container-name mysqldump dbname -uuser -ppassword  > backup.sql
```

# 原因
`-t` がTTY割当、`-i`が標準入力を開くオプションだが、cronの実行では不要だった。

# 解決策
`-it`のオプションを削除すれば解決。


