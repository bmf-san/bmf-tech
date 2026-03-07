---
title: ubuntu初期設定メモ
slug: ubuntu-initial-setup-notes
date: 2019-02-15T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Ubuntu
translation_key: ubuntu-initial-setup-notes
---


# 概要
ちょいちょい忘れてるのでメモ。
不足があれば随時追加。

# 環境
- ConoHa
- ubuntu 18.04.2 LTS (Bionic Beaver)

# 準備
## ubuntuインストール&rootログイン確認
conohaでubuntuサーバーを用意し、rootログインできることを確認しておく。

## ssh用の鍵をクライアント側で用意する
秘密鍵と公開鍵を作成。

`ssh-keygen -t rsa`
`ssh root@<ip address>`

# セットアップ
## サーバーアップデート
アップデートしておく。

`sudo apt update && sudo apt upgrade -y`

## ユーザー作成
sudo権限を持ったユーザーを作成しておく。

`adduser <username>`
`usermod -aG sudo <username>`

wheelに所属しているか確認しておく。
`groups <username>`

※ユーザー一覧確認
`cat /etc/passwd`

## 公開鍵をサーバーに持っていく
作成したユーザーにログイン。
`su <username>`

`.ssh`ディレクトリの準備。
`mkdir .ssh`
`touch .ssh/authorized_keys`
`chmod 700 .ssh`
`chmod 600 .ssh/authorized_keys`

`./ssh/authorized_keys`にクライアント側で作成しておいた公開鍵を貼り付ける。

## sshd_configの設定とポート開放
sshの設定を変更する。

`sudo vi /etc/ssh/sshd_config`

```
Port 5005                                     // デフォルト22から任意の番号に変更
PermitRootLogin no                    // yes → noに変更
PubkeyAuthentication yes         // no → yesに変更
PasswordAuthentication no      // yes → noに変更
UserPAM no                               // yes → noに変更
```

ssh再起動。
`sudo /etc/init.d/ssh restart`

続けてポート開放。

```
sudo ufw allow 5005
sudo ufw allow 443
sudo ufw default deny    // デフォルトの設定でdenyかも...
sudo ufw enable
```

ポート設定を確認。
`sudo ufw status`

# ssh接続確認
`~/.ssh/config`ファイルをこんな感じに編集。

```
ServerAliveInterval 300
TCPKeepAlive yes
AddKeysToAgent yes
ForwardAgent yes
UseKeychain yes

Host conoha-demo
    Hostname    <ip address>
    User         <username>
    Port         5005   // 上で設定した任意のポート番号
    IdentityFile ~/.ssh/<pubkey name>
```

`ssh conoha-demo`でもssh接続できるか確認。

# 所感
昔centosを初めて触ったときも似たようなメモ書いた気がする。
