---
title: Building a Vagrant Development Environment (CentOS 7.3) with Ansible
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
We will build a development environment on Vagrant's CentOS 7.3 using Ansible.

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

The source is available at [github - my-ansible-vagrant](https://github.com/bmf-san/my-ansible-vagrant).

The Vagrantfile looks like this:

```
# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "centos7.3"

  config.vm.network "private_network", ip: "192.168.33.10"

  config.vm.synced_folder "/path/to/directory", "/var/www/html", :mount_options => ["dmode=775,fmode=664"]

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

## Is the php-fpm configuration incorrect??
To use PHP7 with nginx, it seems necessary to use something called php-fpm as a CGI, which can be tricky. If you encounter a 500 error, reviewing this configuration might resolve the issue.

## Unable to access the IP address specified in the Vagrantfile
Although the setup was successful, I struggled to access the IP specified in the Vagrantfile. By reviewing the IP settings and adjusting the firewalld configuration based on the following articles, I managed to resolve it. (It seems I encountered a bug in vagrant 1.9.0.)

# Impressions
CentOS 7 has quite a few differences from previous OS versions, but I didn't struggle too much with those. Rather, I had more trouble with MySQL 5.7. It works for now, but I believe there is still room for improvement.

# References
+ [github - MiyaseTakurou/vagrant_ansible_laravel](https://github.com/MiyaseTakurou/vagrant_ansible_laravel) - The directory structure following best practices was easy to understand.
+ [Initializing virtual machines nicely in vagrant + ansible + CentOS7.0 + VirtualBox environment](http://qiita.com/omochimetaru/items/94bda388dbd05d782f7a)
+ [Building nginx + wordpress with ansible on CentOS7](http://qiita.com/tamanugi/items/2a7fa9701f414ed663c0)
+ [Building a LAMP environment with Vagrant + ansible (4)](http://qiita.com/k-serenade/items/0ab59f9563493f0cf293)
+ [Installing PHP7 and Nginx in a Vagrant Ubuntu 16.04 environment with Ansible](http://koltatt.net/programing/ansible_ubuntu_php/)
+ [Setting up a Ruby environment on CentOS with Vagrant + Ansible](http://qiita.com/yoshiokaCB/items/772bfadf6b7505cb8ba9)
+ [ansible mysql5.7](http://astail.net/?p=1178)
+ [Mysql 5.7 on Ansible](https://www.rennetti.com/howto/139/mysql-5-7-on-ansible)
+ [Building an environment to manipulate Redis with Vagrant + Ansible](http://qiita.com/master-of-sugar/items/e78b173553f5233cd8bd)
+ [Ansible Playbook to install Redis version 2.5 and later](http://qiita.com/joytomo/items/d0cb45074c61dd8935fd)
+ [github - heybigname/ansible](https://github.com/heybigname/ansible/blob/master/tasks/mailcatcher.yml)