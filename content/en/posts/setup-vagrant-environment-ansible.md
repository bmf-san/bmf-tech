---
title: Building a Vagrant Development Environment (CentOS7.3) with Ansible
description: "Configure Vagrant development environments using Ansible automation with PHP, Ruby, Nginx, MySQL, and Redis infrastructure."
slug: setup-vagrant-environment-ansible
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Ansible
  - CentOS
  - Vagrant
translation_key: setup-vagrant-environment-ansible
---


# Overview
We will set up a development environment on Vagrant's CentOS7.3 using Ansible.

# Environment
+ PHP7
+ Ruby
+ Python
+ Nginx
+ MySQL5.7
+ Redis
+ Mailcatcher

# Setup
This is a directory structure that somewhat mimics best practices.

```
ansible/
├── group_vars
│   └── vagrant.yml
├── host
├── roles
│   ├── common
│   │   └── tasks
│   │       ├── add_remi_repo.yml
│   │       ├── install_common.yml
│   │       ├── install_epel_release.yml
│   │       └── main.yml
│   ├── composer
│   │   └── tasks
│   │       ├── install_composer.yml
│   │       └── main.yml
│   ├── mailcatcher
│   │   └── tasks
│   │       ├── install_mailcatcher.yml
│   │       └── main.yml
│   ├── mysql
│   │   └── tasks
│   │       ├── install_mysql.yml
│   │       └── main.yml
│   ├── nginx
│   │   ├── tasks
│   │   │   ├── install_nginx.yml
│   │   │   └── main.yml
│   │   └── templates
│   │       ├── bmf-tech.com.conf
│   │       └── localdev.conf
│   ├── php
│   │   └── tasks
│   │       ├── install_php.yml
│   │       └── main.yml
│   ├── python
│   │   └── tasks
│   │       ├── install_python.yml
│   │       └── main.yml
│   ├── redis
│   │   └── tasks
│   │       ├── install_redis.yml
│   │       └── main.yml
│   └── ruby
│       └── tasks
│           ├── install_ruby.yml
│           └── main.yml
├── site.retry
└── site.yml
```

Please refer to the contents on [github - my-ansible-vagrant](https://github.com/bmf-san/my-ansible-vagrant).

The Vagrantfile looks like this.

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

You can execute provisioning with `vagrant provision`.

# Additional Notes

## Issues with php-fpm Configuration?
To use php7 with nginx, it seems necessary to use a CGI called php-fpm, which was quite tricky. If you encounter a 500 error, reviewing these settings might solve the issue.

[VAGRANTにてCENTOS7にNGINX+PHP-FPM+PHP7でLARAVELの開発環境構築(前編)](https://namaikinamaiki.wordpress.com/2015/11/02/vagrant%E3%81%AB%E3%81%A6centos7%E3%81%ABnginxphp-fpmphp7%E3%81%A7laravel%E3%81%AE%E9%96%8B%E7%99%BA%E7%92%B0%E5%A2%83%E6%A7%8B%E7%AF%89%EF%BC%91/)

## Cannot Access the IP Address Specified in Vagrantfile
Although the setup was completed, I struggled quite a bit because I couldn't access the IP specified in the Vagrantfile. By referring to the following articles, I reviewed the IP settings and adjusted the firewalld settings, which eventually resolved the issue. (It seems the cause was a bug in vagrant1.9.0.)

+ [[Vagrant]Vagrantfileで指定したipアドレスでアクセスができない場合の対応](http://to-developer.com/blog/?p=1827)
+ [Vagrantでpingが通らない！ゆえにVagrantネットワークを学び直したよ](http://www.kaasan.info/archives/3665)
+ [vagrant + centos7 でprivate_networkで設定したIPに接続ができない](http://qiita.com/junqiq/items/a19d3ea48b072a1b28d3)
+ [Vagrant で CentOS7 + PHP + MySQL の仮想環境を構築する](http://fnya.cocolog-nifty.com/blog/2015/12/vagrant-centos7.html)

# Impressions
There are quite a few differences in CentOS7 compared to previous OS versions, but I didn't struggle much with those. Instead, I had more trouble with MySQL5.7. It works for now, but I think there's still room for improvement.

# References
+ [github - MiyaseTakurou/vagrant_ansible_laravel](https://github.com/MiyaseTakurou/vagrant_ansible_laravel) - The directory structure for best practices was easy to understand.
+ [vagrant + ansible + CentOS7.0 + VirtualBox 環境で仮想マシンを良い感じに初期化する](http://qiita.com/omochimetaru/items/94bda388dbd05d782f7a)
+ [CentOS7でansibleを使ってnginx+wordpressを構築](http://qiita.com/tamanugi/items/2a7fa9701f414ed663c0)
+ [Vagrant+ansibleでLAMP環境構築（４）](http://qiita.com/k-serenade/items/0ab59f9563493f0cf293)
+ [AnsibleでVagrantのubuntu16.04環境にPHP7とNginxをインストールする](http://koltatt.net/programing/ansible_ubuntu_php/)
+ [Vagrant + Ansible でCentOSにRubyの環境構築してみる。](http://qiita.com/yoshiokaCB/items/772bfadf6b7505cb8ba9)
+ [ansible mysql5.7](http://astail.net/?p=1178)
+ [Mysql 5.7 on Ansible](https://www.rennetti.com/howto/139/mysql-5-7-on-ansible)
+ [Redisをいじくり倒す環境をVagrant+Ansibleで構築する](http://qiita.com/master-of-sugar/items/e78b173553f5233cd8bd)
+ [AnsibleでRedisの2.5以降のバージョンをインストールするPlaybook](http://qiita.com/joytomo/items/d0cb45074c61dd8935fd)
+ [github - heybigname/ansible](https://github.com/heybigname/ansible/blob/master/tasks/mailcatcher.yml)
