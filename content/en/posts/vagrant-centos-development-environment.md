---
title: Setting Up a CentOS 6.7 Development Environment with Vagrant
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
Recently, when I rebuilt my Box, I thought, "I should properly document the Vagrant development environment workflow," so I decided to summarize it. There is a similar workflow memo in the repository.

[github - bmf-san/vagrant-development-workflow](https://github.com/bmf-san/vagrant-development-workflow)

# Prerequisites

The following applications should be installed on the host machine (Mac):

* [Vagrant](https://www.vagrantup.com/downloads.html)
* [VirtualBox](https://www.virtualbox.org/)

# Environment
## Host Machine (Mac)
+ macOS Sierra v10.12.2

## Virtual Environment
+ CentOS 6.7

# Setup Procedure
+ Create a Vagrantfile in the development environment directory
  + `vagrant init`

+ Obtain a Box template and create a Box
  + `vagrant box add BOX_NAME /path/to/box/url`
  + Specify any name for BOX_NAME
  + Download the Box from [Vagrantbox.es](http://www.vagrantbox.es/)
  + Use CentOS 6.7

+ Install [vagrant-hostupdater](https://github.com/cogitatio/vagrant-hostsupdater)
  + `vagrant plugin install vagrant-hostsupdater`
  + [Hosts](http://permanentmarkers.nl/software.html) - A useful app for GUI management of hosts on the host machine (Mac)

+ Edit the Vagrantfile
  + Set Box Name
  + Configure Network
  + Configure Synced Folder
  + Configure Provider (to improve Vagrant performance)
  + Configure Host Updater
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
  + `vagrant up` - Start
  + `vagrant ssh` - SSH connection
  + `vagrant reload` - Restart
  + `vagrant halt` - Stop
  + `vagrant provision` - Provisioning (update host)

+ Install Apache
  + `yum install httpd` - Install Apache
  + `service httpd start` - Start server
  + `chkconfig httpd on` - Set to start automatically on login

+ Configure Apache
  + `cd /etc/httpd/conf.d`
  + `vim localdev-hoge.conf` - Create a configuration file for each host (set to access with `localdev-hoge`)

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

  + `service httpd restart` - Restart server to apply settings

# Learn More
  + Install various necessary applications (php, mailcatcher, xdebug, webgrind, etc.) to build the development environment
