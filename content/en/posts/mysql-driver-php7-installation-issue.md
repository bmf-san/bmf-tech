---
title: No MySQL Driver After Installing PHP7
slug: mysql-driver-php7-installation-issue
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - CentOS
  - MySQL
  - PHP
  - PHP7
  - Sakura VPS
  - yum
translation_key: mysql-driver-php7-installation-issue
---



# No MySQL Driver After Installing PHP7

For PHP7 installation, I referred to this guide.
* [CentOS6/CentOS7: Installing PHP5.6/PHP7 with yum](http://qiita.com/ozawan/items/caf6e7ddec7c6b31f01e)

By the way, my environment is...
* Sakura VPS
* CentOS6.7


# If There's No Driver, Just Install It

```
yum install yum install --enablerepo=remi,remi-php70 php-mysqlnd
```

This seems to have resolved the issue _(:3」∠)_

# Note: Regarding PHP7 Installation
There might be some missing packages, but for using Laravel with PHP7 via Composer, it seems sufficient.

```
yum -y install --enablerepo=remi-php70 php php-mbstring php-pear php-fpm php-mcrypt php-devel php-xml
```

Reference
* [Updated from PHP5.6 to PHP7](https://monochrome-design.jp/126)

After switching to PHP7, it might just be my imagination, but it feels faster. No, the perceived speed has clearly changed...
