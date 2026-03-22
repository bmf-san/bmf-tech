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
We will build a development environment on CentOS 7.3 using Vagrant and Ansible.

# Environment
+ PHP 7
+ Ruby
+ Python
+ Nginx
+ MySQL 5.7
+ Redis
+ Mailcatcher

# Setup
This directory structure somewhat mimics best practices.

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

The source is available on [github - my-ansible-vagrant](https://github.com/bmf-san/my-ansible-vagrant), so please refer to that for the contents.

Here is how the Vagrantfile looks.

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

You can run provisioning with `vagrant provision`.

# Additional Notes

## Is the php-fpm configuration incorrect?
To use PHP 7 with Nginx, it seems necessary to use something called php-fpm as a CGI, which can be tricky. If you encounter a 500 error, reviewing the configuration in this area might resolve the issue.

[Building a Laravel Development Environment on CENTOS 7 with NGINX+PHP-FPM+PHP7 on VAGRANT (Part 1)](https://namaikinamaiki.wordpress.com/2015/11/02/vagrant%E3%81%AB%E3%81%A6centos7%E3%81%ABnginxphp-fpmphp7%E3%81%A7laravel%E3%81%AE%E9%96%8B%E7%99%BA%E7%92%B0%E5%A2%83%E6%A7%8B%E7%AF%89%EF%BC%91/)

## Unable to access the IP address specified in the Vagrantfile
Although I was able to build it, I struggled to access the IP specified in the Vagrantfile. After reviewing the IP settings and adjusting the firewalld configuration based on the following articles, I managed to resolve the issue. (It seems I encountered a bug in vagrant 1.9.0.)

+ [[Vagrant] How to deal with not being able to access the IP address specified in the Vagrantfile](http://to-developer.com/blog/?p=1827)
+ [Vagrant cannot ping! Therefore, I had to relearn Vagrant networking](http://www.kaasan.info/archives/3665)
+ [Cannot connect to the IP set in private_network with vagrant + centos7](http://qiita.com/junqiq/items/a19d3ea48b072a1b28d3)
+ [Building a virtual environment with CentOS 7 + PHP + MySQL using Vagrant](http://fnya.cocolog-nifty.com/blog/2015/12/vagrant-centos7.html)

# Impressions
CentOS 7 has quite a few differences from previous OS versions, but I didn't struggle too much with those adjustments. Instead, I had more trouble with the compatibility of MySQL 5.7. It works for now, but I think there is still room for improvement.

# References
+ [github - MiyaseTakurou/vagrant_ansible_laravel](https://github.com/MiyaseTakurou/vagrant_ansible_laravel) - The directory structure for best practices was easy to understand.
+ [Initializing a virtual machine nicely in a vagrant + ansible + CentOS 7.0 + VirtualBox environment](http://qiita.com/omochimetaru/items/94bda388dbd05d782f7a)
+ [Building nginx + wordpress with ansible on CentOS 7](http://qiita.com/tamanugi/items/2a7fa9701f414ed663c0)
+ [Building a LAMP environment with Vagrant + ansible (Part 4)](http://qiita.com/k-serenade/items/0ab59f9563493f0cf293)
+ [Installing PHP 7 and Nginx on an ubuntu 16.04 environment with Ansible](http://koltatt.net/programing/ansible_ubuntu_php/)
+ [Setting up a Ruby environment on CentOS with Vagrant + Ansible](http://qiita.com/yoshiokaCB/items/772bfadf6b7505cb8ba9)
+ [ansible mysql5.7](http://astail.net/?p=1178)
+ [Mysql 5.7 on Ansible](https://www.rennetti.com/howto/139/mysql-5-7-on-ansible)
+ [Building an environment to manipulate Redis with Vagrant + Ansible](http://qiita.com/master-of-sugar/items/e78b173553f5233cd8bd)
+ [Ansible playbook to install Redis version 2.5 and later](http://qiita.com/joytomo/items/d0cb45074c61dd8935fd)
+ [github - heybigname/ansible](https://github.com/heybigname/ansible/blob/master/tasks/mailcatcher.yml)
