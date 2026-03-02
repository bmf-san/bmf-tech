---
title: "VagrantでCentOS6.7の開発環境を構築する"
slug: "vagrant-centos6"
date: 2017-10-01
author: bmf-san
categories:
  - "インフラストラクチャ"
tags:
  - "apache"
  - "CentOS"
  - "Vagrant"
draft: false
---

# 概要
最近、Boxを再構築した際に、「Vagrant開発環境のワークフローをちゃんとまとめておきたい」と思ったのでまとめてみました。
レポジトリにほぼ同じワークフローメモがあります。

[github - bmf-san/vagrant-development-workflow](https://github.com/bmf-san/vagrant-development-workflow)

# 前提

以下のアプリケーションがホストマシン（Mac）にインストールされていること

* [Vagrant](https://www.vagrantup.com/downloads.html)
* [VirtualBox](https://www.virtualbox.org/)

# 環境
## ホストマシン（Mac）
+ macOS Sierra v10.12.2

## 仮想環境
+ CentOS6.7

# 構築手順
+ 開発環境ディレクトリにて、Vagrantfileを作成する
  + `vagrant init`

+ Boxテンプレートを取得し、Boxを作成する
  + `vagrant box add BOX_NAME /path/to/box/url`
  + BOX_NAMEは任意の名前を指定
  + Boxのダウンロード先は[Vagrantbox.es](http://www.vagrantbox.es/)
  + CentOS6.7を使用

+ [vagrant-hostupdater](https://github.com/cogitatio/vagrant-hostsupdater)のインストール
  + `vagrant plugin install vagrant-hostsupdater`
  + [Hosts](http://permanentmarkers.nl/software.html)ーホストマシン（Mac）でホストのGUI管理に便利なアプリ

+ Vagrantfileの編集
  + Box Nameの設定
  + Networkの設定
  + Synced Folder（共有フォルダ）の設定
  + Providerの設定（Vagrantパフォーマンス向上が目的）
  + Host Updaterの設定
  + xdebugの設定（任意）

```
# -*- mode: ruby -*-
# vi: set ft=ruby :
Vagrant.configure(2) do |config|
  # Box Name
  config.vm.box = "centos6.7"

  # Network
  config.vm.network "private_network", ip: "192.168.33.10"

  # Synced Folder
  config.vm.synced_folder "/path/to/directory", "/var/www/html",
    :owner => "apache",
    :group => "apache",
    :mount_options => ["dmode=775,fmode=664"]

  # Provider(Optional)
  config.vm.provider "virtualbox" do |vb|
        vb.customize ["modifyvm", :id, "--paravirtprovider", "kvm"]
  end

  # Host Updater
  config.vm.network :private_network, ip: "192.168.33.10"
  config.vm.hostname = "localdev"
  config.hostsupdater.aliases = ["localdev-hoge"]

  # xdebug(Optional)
  config.vm.network :forwarded_port, host: 3000, guest: 3000
end
```

+ Vagrantの起動と接続
  + `vagrant up`ー起動
  + `vagrant ssh`ーssh接続
  + `vagrant reload`ー再起動
  + `vagrant halt`ー停止
  + `vagrant provision`ープロビジョニング（ホストの更新）

+ Apacheのインストール
  + `yum install httpd`ーApacheのインストール
  + `service httpd start`ーサーバー起動
  + `chkconfig httpd on`ーログイン時自動起動設定

+ Apacheの設定
  + `cd /etc/httpd/conf.d`
  + `vim localdev-hoge.conf`ーホスト別の設定ファイルを作成（`localdev-hoge`でアクセスできるように設定）

```
<VirtualHost *:80>
  ServerName localdev-hoge
  ServerAdmin localhost
  DocumentRoot /var/www/html/path/to/directory

  <Directory "/var/www/html/path/to/directory">
    AllowOverride All
  </Directory>
</VirtualHost>
```

  + `service httpd restart`ーサーバー再起動で設定を反映

# Learn More
  + 各種必要なアプリケーション（php, mailcatcher, xdebug, webgrind etc...）をインストールし、開発環境を構築すること

