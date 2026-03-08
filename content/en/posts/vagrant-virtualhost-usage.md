---
title: Using VirtualHost with Vagrant
slug: vagrant-virtualhost-usage
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Vagrant
translation_key: vagrant-virtualhost-usage
---

I have been maintaining a virtual environment with a VagrantBox that I created and left unattended, but I finally set up the VirtualHost configuration.

# What to Do
* Install vagrant-hostupdater and configure Vagrantfile
* Configure Apache's conf files (VirtualHost settings)

# What Not to Do
* Install Vagrant
* Set up symbolic links or other tedious tasks

# Environment
* Vagrant
* vagrant-hostsupdater ... this is the Vagrant plugin we will install this time.
* CentOS 6 series
* Apache
* Hosts ... If you install this, you can check the Hosts from the GUI in your MAC's environment settings. You can also configure Hosts. This time, we will configure Hosts using the Vagrant plugin, so we won't set it up from the GUI. Reference: [Editing Mac's hosts with Hosts is good【Mac】](http://blog.sou-lab.com/mac-hosts/)

# Installing vagrant-hostupdater and Editing Vagrantfile

Install vagrant-hostupdater (you may need sudo or root privileges)
`vagrant plugin install vagrant-hostsupdater`

Edit the Vagrantfile
Add the following lines at the end of the file.
```
# vagrant-hostupdater
 config.vm.network :private_network, ip: "192.168.33.10"
 config.vm.hostname = "localdev"
 config.hostsupdater.aliases = ["dev", "hoge"]
```

Feel free to set the hostname and aliases. The aliases will be used later in the Apache conf files. Think of them as domain names.

Once the configuration is complete, start or restart Vagrant.

You should be able to confirm that the Hosts are automatically set by checking /private/etc/hosts in the terminal or by looking at Hosts if you have it installed.

# VirtualHost Configuration
Some of you may be familiar with this from using VPS. The steps are the same.

Uncomment `#NameVirtualHost *:80` in /etc/httpd/conf/httpd.conf.

Open httpd.conf with vi or vim, type /NameV, press enter, and hit n a few times to find it.

You can write the virtual host settings directly in httpd.conf, but I will consolidate the settings in /etc/httpd/conf.d.

Since we set the aliases to dev and hoge, we will create separate configuration files for each.

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

I will omit the log file output destination (if omitted, it will default to the standard output destination).

/var/www/html is a shared folder or a symbolic link, which can be troublesome with Vagrant. Please adjust it according to your environment.

If you can display dev/ and hoge/ in your browser, you're good to go.

# Thoughts
I regretted not doing this sooner (´・ω・`)

# References
* [Configured VirtualHost with Vagrant](http://raichel.hatenablog.com/entry/2015/06/06/205958)
* [Setting up VirtualHost using Vagrant](http://sk51.jp/how-to-setup-virtualhost-using-vagrant/)