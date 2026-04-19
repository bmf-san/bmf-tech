---
title: AnsibleでVagrant開発環境（CentOS7.3）を構築する
description: "構築するVagrant開発環境。Ansible自動化でPHP・Ruby・Python・Nginx・MySQL・Redis・Mailcatcher、CentOS7.3インスタンス構成。"
slug: setup-vagrant-environment-ansible
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Ansible
  - CentOS
  - Vagrant
translation_key: setup-vagrant-environment-ansible
---


# 概要
VagrantのCentOS7.3に開発環境をAnsibleで構築します。

# 環境
+ PHP7
+ Ruby
+ Python
+ Nginx
+ MySQL5.7
+ Redis
+ Mailcatcher

# 構築
ベストプラクティスをある程度模倣した形のディレクトリです。

```
ansible/
├── group_vars
│   └── vagrant.yml
├── host
├── roles
│   ├── common
│   │   └── tasks
│   │       ├── add_remi_repo.yml
│   │       ├── install_common.yml
│   │       ├── install_epel_release.yml
│   │       └── main.yml
│   ├── composer
│   │   └── tasks
│   │       ├── install_composer.yml
│   │       └── main.yml
│   ├── mailcatcher
│   │   └── tasks
│   │       ├── install_mailcatcher.yml
│   │       └── main.yml
│   ├── mysql
│   │   └── tasks
│   │       ├── install_mysql.yml
│   │       └── main.yml
│   ├── nginx
│   │   ├── tasks
│   │   │   ├── install_nginx.yml
│   │   │   └── main.yml
│   │   └── templates
│   │       ├── bmf-tech.com.conf
│   │       └── localdev.conf
│   ├── php
│   │   └── tasks
│   │       ├── install_php.yml
│   │       └── main.yml
│   ├── python
│   │   └── tasks
│   │       ├── install_python.yml
│   │       └── main.yml
│   ├── redis
│   │   └── tasks
│   │       ├── install_redis.yml
│   │       └── main.yml
│   └── ruby
│       └── tasks
│           ├── install_ruby.yml
│           └── main.yml
├── site.retry
└── site.yml
```

[github - my-ansible-vagrant](https://github.com/bmf-san/my-ansible-vagrant)にソースを上げているので中身はそちらをご参照ください。

Vagrantfileはこんな感じです。

```
# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "centos7.3"

  config.vm.network "private_network", ip: "192.168.33.10"

  config.vm.synced_folder "/path/to/directory", "/var/www/html",:mount_options => ["dmode=775,fmode=664"]

  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "ansible/site.yml"
    ansible.inventory_path = "ansible/host"
    ansible.limit = 'all'
  end

  config.vm.network :private_network, ip: "192.168.33.10"
  config.vm.hostname = "localdev"
  config.hostsupdater.aliases = ["localdev"]
end
```

`vagrant provision`でプロビジョニングを実行できます。

# 追記

## php-fpmの設定がおかしい？？
nginxでphp7を使うにはphp-fpmとかいうCGIをかます必要があるらしいのですが、これがハマりやすかったです。500エラーが出た時などは、このあたりを設定を見直すと解決するかもです。

[VAGRANTにてCENTOS7にNGINX+PHP-FPM+PHP7でLARAVELの開発環境構築(前編)](https://namaikinamaiki.wordpress.com/2015/11/02/vagrant%E3%81%AB%E3%81%A6centos7%E3%81%ABnginxphp-fpmphp7%E3%81%A7laravel%E3%81%AE%E9%96%8B%E7%99%BA%E7%92%B0%E5%A2%83%E6%A7%8B%E7%AF%89%EF%BC%91/)

## Vagrantfileに指定したipアドレスにアクセスできない
構築できたものの、Vagrantfileに指定したipにアクセスできずに結構ハマりました。
以下の記事を参考にipの設定を見直したり、firewalldの設定を調整したら何とか解決できました。
（vagrant1.9.0のバグを踏んでしまっていたのが原因だったみたいです。）

+ [[Vagrant]Vagrantfileで指定したipアドレスでアクセスができない場合の対応](http://to-developer.com/blog/?p=1827)
+ [Vagrantでpingが通らない！ゆえにVagrantネットワークを学び直したよ](http://www.kaasan.info/archives/3665)
+ [vagrant + centos7 でprivate_networkで設定したIPに接続ができない](http://qiita.com/junqiq/items/a19d3ea48b072a1b28d3)
+ ~~Vagrant で CentOS7 + PHP + MySQL の仮想環境を構築する~~

# 所感
CentOS7はそれまでのOSバージョンと異なる部分が結構あるのですが、その対応にはそんなにハマりませんでした。
むしろ、MySQL5.7の対応にハマりました。
とりあえず動くきますが、まだまだ改善の余地があるかと思います。

# 参考
+ [github - MiyaseTakurou/vagrant_ansible_laravel](http://web.archive.org/web/20180611161928/https://github.com/MiyaseTakurou/vagrant_ansible_laravel) ーベストプラクティスのディレクトリ構成が分かりやすかったです。
+ [vagrant + ansible + CentOS7.0 + VirtualBox 環境で仮想マシンを良い感じに初期化する](http://qiita.com/omochimetaru/items/94bda388dbd05d782f7a)
+ [CentOS7でansibleを使ってnginx+wordpressを構築](http://qiita.com/tamanugi/items/2a7fa9701f414ed663c0)
+ [Vagrant+ansibleでLAMP環境構築（４）](http://qiita.com/k-serenade/items/0ab59f9563493f0cf293)
+ ~~AnsibleでVagrantのubuntu16.04環境にPHP7とNginxをインストールする~~
+ [Vagrant + Ansible でCentOSにRubyの環境構築してみる。](http://qiita.com/yoshiokaCB/items/772bfadf6b7505cb8ba9)
+ [ansible mysql5.7](http://astail.net/?p=1178)
+ ~~Mysql 5.7 on Ansible~~
+ [Redisをいじくり倒す環境をVagrant+Ansibleで構築する](http://qiita.com/master-of-sugar/items/e78b173553f5233cd8bd)
+ [AnsibleでRedisの2.5以降のバージョンをインストールするPlaybook](http://qiita.com/joytomo/items/d0cb45074c61dd8935fd)
+ ~~github - heybigname/ansible~~

