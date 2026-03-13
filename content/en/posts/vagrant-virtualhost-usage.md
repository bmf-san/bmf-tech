---
title: Using VirtualHost with Vagrant
description: "Configure VirtualHost with Vagrant using vagrant-hostupdater and Apache for multi-domain local development environments."
slug: vagrant-virtualhost-usage
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Vagrant
translation_key: vagrant-virtualhost-usage
---


I have been maintaining a virtual environment with a VagrantBox left as is, but I finally set up VirtualHost.

# What to Do
* Install vagrant-hostupdater and configure Vagrantfile
* Configure Apache's conf file (VirtualHost settings)


# What Not to Do
* Install vagrant
* Set up symbolic links or other tedious tasks

# Environment
* Vagrant
* vagrant-hostsupdater... This is the Vagrant plugin to be installed this time.
* CentOS6 series
* Apache
* Hosts... If installed, you can check Hosts from the MAC environment settings via GUI. You can also configure Hosts. This time, Hosts will be configured with the Vagrant plugin, so no GUI configuration will be done. Reference: [Macのhosts編集はHostsが良さげ【Mac】](http://blog.sou-lab.com/mac-hosts/)


# Installing vagrant-hostupdater and Editing Vagrantfile

Install vagrant-hostupdater (may require sudo or root privileges)
`vagrant plugin install vagrant-hostsupdater`


Edit Vagrantfile
Add the following description at the end of the file.
```
# vagrant-hostupdater
 config.vm.network :private_network, ip: "192.168.33.10"
 config.vm.hostname = "localdev"
 config.hostsupdater.aliases = ["dev", "hoge"]
```

Feel free to set the hostname and alias. The alias will be used later in the Apache conf file. Think of it as a domain name.

Once the configuration is complete, start or restart Vagrant.


Check /private/etc/hosts in the terminal, or if you have Hosts installed, you should be able to confirm that the Host is automatically configured.


# VirtualHost Configuration
Some of you may be familiar with this on a VPS. It's the same procedure.


In /etc/httpd/conf, uncomment `#NameVirtualHost *:80` in httpd.conf.

Open httpd.conf with vi or vim, type /NameV, press enter, and hit n a few times to find it.


You can write the virtual host settings directly in httpd.conf, but let's consolidate the settings in /etc/httpd/conf.d.


Since we set aliases to dev and hoge, create configuration files for each.


/etc/httpd/conf.d/dev.conf

```
<VirtualHost *:80>
        ServerName dev
        ServerAdmin localhost
        DocumentRoot /var/www/html/dev

        <Directory "/var/www/html/dev">
                AllowOverride All
        </Directory>
</VirtualHost>
```

/etc/httpd/conf.d/hoge.conf
```
<VirtualHost *:80>
        ServerName hoge
        ServerAdmin localhost
        DocumentRoot /var/www/html/hoge

        <Directory "/var/www/html/hoge">
                AllowOverride All
        </Directory>
</VirtualHost>
```


The log file output destination is omitted. (If omitted, it will default to the default output destination.)

/var/www/html is a shared folder or symbolic link, a tricky part of Vagrant. Please adjust according to your environment.

If you can display dev/ and hoge/ in your browser, it's OK.


# Thoughts
I regretted not doing this sooner (´・ω・`)


# References
* [Vagrantでバーチャルホストを設定してみたｗ](http://raichel.hatenablog.com/entry/2015/06/06/205958)
* [Vagrantを利用したVirtualHostの設定](http://sk51.jp/how-to-setup-virtualhost-using-vagrant/)
