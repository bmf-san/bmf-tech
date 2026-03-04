---
title: "Laravel Homesteadを使ったLaravelの環境構築"
slug: "laravel-homestead-setup"
date: 2018-04-11
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Laravel"
  - "Vagrant"
  - "VirtualBox"
  - "composer"
  - "homestead"
draft: false
---

# 概要
Homesteadを触る機会があったのでざっくりまとめる。

# 準備
以下のツールは用意しておきましょう。
- composer
- vagrant
- virtualbox
- ssh key
    - 鍵つくっておいてください

# 手順
## Laravelをインストール
laravelのインストールと`composer install`を実行しましょう。

`composer create-project "laravel/laravel=5.5.*" projectname`

`cd app`
`composer install`

## HomeSteadを用意
Vagrant boxを用意します。

`vagrant box add laravel/homestead`

homesteadを利用するためのリポジトリをクローンし、初期化スクリプトを実行します。

`cd ~`
`git clone https://github.com/laravel/homestead.git Homestead`
`bash init.sh`

YAMLファイルで仮想環境の設定を調整します。

`vi Homestead.yaml`

```:bash
ip: "192.167.10.99" // edit
memory: 2048
cpus: 1
provider: virtualbox

authorize: ~/.ssh/id_rsa.pub

keys:
    - ~/.ssh/id_rsa

folders:
    - map: ~/localdev/project/laravel // edit
      to: /home/vagrant/code

sites:
    - map: laravel // edit
      to: /home/vagrant/code/public

databases:
    - homestead

# blackfire:
#     - id: foo
#       token: bar
#       client-id: foo
#       client-token: bar

# ports:
#     - send: 50000
#       to: 5000
#     - send: 7777
#       to: 777
#       protocol: udp
```

次に、ホストファイルを編集します。

`vi /etc/hosts`

```
##
# Host Database
#
# localhost is used to configure the loopback interface
# when the system is booting.  Do not change this entry.
##
127.0.0.1       localhost
255.255.255.255 broadcasthost
::1     localhost
192.167.10.99  laravel // edit
192.167.10.99  homestead  # VAGRANT: e78b975204ccde53ef11f1ffe284d4f4 (homestead-7) / b212bb82-8dc8-405b-8507-fa284ddb9aa8
```

## Vagrantを起動
`cd ~/Homestead`
`vagrant up`

以下にアクセスするとLaravelのデフォルトのウェルカムページが見れるはずです。

`http://laravel`

# リポジトリ
- [Github - bmf-san/laravel-homestead-boilerplate](https://github.com/bmf-san/laravel-homestead-boilerplate)

