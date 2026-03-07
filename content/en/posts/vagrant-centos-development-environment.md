---
title: Building a Development Environment with Vagrant on CentOS 6.7
slug: vagrant-centos-development-environment
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - apache
  - CentOS
  - Vagrant
translation_key: vagrant-centos-development-environment
---

# Overview
Recently, while rebuilding my Box, I thought, "I want to properly summarize the Vagrant development environment workflow," so I compiled this information. There is a nearly identical workflow memo in the repository.

[github - bmf-san/vagrant-development-workflow](https://github.com/bmf-san/vagrant-development-workflow)

# Prerequisites
The following applications must be installed on the host machine (Mac):

* [Vagrant](https://www.vagrantup.com/downloads.html)
* [VirtualBox](https://www.virtualbox.org/)

# Environment
## Host Machine (Mac)
+ macOS Sierra v10.12.2

## Virtual Environment
+ CentOS 6.7

# Setup Steps
+ In the development environment directory, create a Vagrantfile
  + `vagrant init`

+ Obtain the Box template and create the Box
  + `vagrant box add BOX_NAME /path/to/box/url`
  + Specify any name for BOX_NAME
  + The download source for the Box is [Vagrantbox.es](http://www.vagrantbox.es/)
  + Use CentOS 6.7

+ Install [vagrant-hostupdater](https://github.com/cogitatio/vagrant-hostsupdater)
  + `vagrant plugin install vagrant-hostsupdater`
  + [Hosts](http://permanentmarkers.nl/software.html) - A convenient app for GUI management of hosts on the host machine (Mac)

+ Edit the Vagrantfile
  + Set the Box Name
  + Configure the Network
  + Set up Synced Folder (shared folder)
  + Configure the Provider (to improve Vagrant performance)
  + Set up Host Updater
  + Configure xdebug (optional)

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

+ Start and connect to Vagrant
  + `vagrant up` - start
  + `vagrant ssh` - ssh connection
  + `vagrant reload` - restart
  + `vagrant halt` - stop
  + `vagrant provision` - provisioning (updating the host)

+ Install Apache
  + `yum install httpd` - install Apache
  + `service httpd start` - start the server
  + `chkconfig httpd on` - set to start automatically on login

+ Configure Apache
  + `cd /etc/httpd/conf.d`
  + `vim localdev-hoge.conf` - create a separate configuration file for the host (set up to access with `localdev-hoge`)

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

  + `service httpd restart` - restart the server to apply the configuration

# Learn More
  + Install various necessary applications (php, mailcatcher, xdebug, webgrind, etc.) to build the development environment.