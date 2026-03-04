---
title: "ssh接続セットアップメモ"
slug: "ssh-connection-setup-notes"
date: 2018-09-18
author: bmf-san
categories:
  - "インフラストラクチャ"
tags:
  - "ssh"
  - "sshd"
draft: false
---

# 概要
ssh接続のセットアップ方法についてメモ。

# 準備　
サーバーに接続してwheelグループに所属するユーザーを作成しておく

# 手順
## ホスト側
`~/.ssh/`にて`ssh-keygen`で公開鍵・秘密鍵を作成。
ここでは公開鍵を`id_rsa.pub`、秘密鍵を`id_rsa`として作成する。

公開鍵の中身をコピーしておく。

`~/.ssh/`にて`config`ファイルを作成しておく。

Ex.
```
Host bmf
 HostName 123.45.679.012
 User bmf
 Port 22
 IdentityFile ~/.ssh/id_rsa
```

## サーバー側
`~/.ssh/`が存在しない場合はディレクトリを作成する。
パーミッションは700を指定。
`mkdir .ssh && chmod 700 .ssh`

続いて、`~/.ssh/`にて`authorized_keys`という名前のファイルを作成。
パーミッションは600を指定。
`authorized_keys`には公開鍵の中身を貼り付ける。

次に`/etc/ssh/sshd_config`の設定を調整する。
以下の設定を調整。
- Port
  - コメントアウトを外す。デフォルトでは22番ポートだがセキュリティを考慮して別の番号をしてしたほうがよい。
- PasswordAuthentication
  - パスワード認証による接続をoffにする。noを指定。
- PermitRootLogin
  - rootユーザーでのログイン許可をoffにする。noを指定。

ssh接続で使用するポート番号が空いている確認。
`firewall-cmd --list-all`

空いていない場合は開放。
`firewall-cmd --permanent --zone=public --add-port=22/tcp`

リロード。

`firewall-cmd --reload`

# 接続してみる
`ssh bmf`

