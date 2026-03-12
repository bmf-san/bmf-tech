---
title: 'Encountered [Errno 14] PYCURL ERROR 22 - "The requested URL returned error: 404 Not Found" on CentOS6.7'
description: 'A troubleshooting guide for Encountered [Errno 14] PYCURL ERROR 22 - "The requested URL returned error: 404 Not Found" on CentOS6.7, explaining the root cause and how to resolve it.'
slug: centos-error-404-pycurl
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - CentOS
  - OS
  - Zabbix
  - Tips
translation_key: centos-error-404-pycurl
---

# Encountered [Errno 14] PYCURL ERROR 22 - "The requested URL returned error: 404 Not Found" on CentOS6.7

While trying to set up Zabbix, I encountered the following error, which rendered yum unusable.

```
http://mirror.centos.org/centos/6/SCL/x86_64/repodata/repomd.xml: [Errno 14] PYCURL ERROR 22 - "The requested URL returned error: 404 Not Found"
Trying other mirror.
Error: Cannot retrieve repository metadata (repomd.xml) for repository: scl. Please verify its path and try again
```

Faced with this, I was quite disheartened due to my limited OS knowledge, but the following link was helpful and led to a solution, so I am sharing it.

[[tips][Linux] When you can't update yum on older versions of CentOS](http://luozengbin.github.io/blog/2015-08-29-%5Btips%5D%5Blinux%5D%E6%97%A7%E3%83%90%E3%83%BC%E3%82%B8%E3%83%A7%E3%83%B3centos%E3%81%A7yum%E6%9B%B4%E6%96%B0%E3%81%A7%E3%81%8D%E3%81%AA%E3%81%8F%E3%81%AA%E3%81%A3%E3%81%9F%E6%99%82.html)

Incidentally, I couldn't install Zabbix due to complications with PHP settings...

Programs like server monitoring tools or profilers that need to be installed on servers have a high barrier to entry at the installation stage...