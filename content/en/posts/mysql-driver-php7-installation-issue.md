---
title: No MySQL Driver Found When Installing PHP7
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

# No MySQL Driver Found When Installing PHP7

I referred to this for installing PHP7.
* [Install PHP5.6/PHP7 on CentOS6/CentOS7 using yum](http://qiita.com/ozawan/items/caf6e7ddec7c6b31f01e)

By the way, my environment is...
* Sakura VPS
* CentOS6.7

# If the Driver is Missing, Just Install It

```
yum install --enablerepo=remi,remi-php70 php-mysqlnd
```

This seems to have resolved the issue _(:3」∠)_.

# Additional Notes on Installing PHP7
There might be some missing packages, but I feel like there are no shortages for using Laravel with PHP7 via Composer.

```
yum -y install --enablerepo=remi-php70 php php-mbstring php-pear php-fpm php-mcrypt php-devel php-xml
```

References:
* [Updated from PHP5.6 to PHP7](https://monochrome-design.jp/126)

After switching to PHP7, I might just be imagining it, but it feels faster. No, it clearly seems like the perceived speed has changed...