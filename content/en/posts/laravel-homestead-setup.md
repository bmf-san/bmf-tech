---
title: Setting Up a Laravel Environment with Laravel Homestead
description: An in-depth look at Setting Up a Laravel Environment with Laravel Homestead, covering key concepts and practical insights.
slug: laravel-homestead-setup
date: 2018-04-11T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
  - Vagrant
  - VirtualBox
  - composer
  - homestead
translation_key: laravel-homestead-setup
---



# Overview
I had the opportunity to work with Homestead, so here's a brief summary.

# Preparation
Make sure to have the following tools ready:
- composer
- vagrant
- virtualbox
- ssh key
    - Please create a key

# Steps
## Install Laravel
Let's install Laravel and run `composer install`.

`composer create-project "laravel/laravel=5.5.*" projectname`

`cd app`
`composer install`

## Prepare Homestead
Prepare the Vagrant box.

`vagrant box add laravel/homestead`

Clone the repository for using Homestead and run the initialization script.

`cd ~`
`git clone https://github.com/laravel/homestead.git Homestead`
`bash init.sh`

Adjust the virtual environment settings in the YAML file.

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

Next, edit the host file.

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

## Start Vagrant
`cd ~/Homestead`
`vagrant up`

You should be able to see the default Laravel welcome page by accessing the following:

`http://laravel`

# Repository
- [Github - bmf-san/laravel-homestead-boilerplate](https://github.com/bmf-san/laravel-homestead-boilerplate)
