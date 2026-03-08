---
title: MySQL 'MySQL Daemon failed to start' on Vagrant
slug: vagrant-mysql-daemon-start-failure
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - MySQL
  - Vagrant
  - Tips
translation_key: vagrant-mysql-daemon-start-failure
---



While repeatedly running migrations and importing SQL files with Laravel, and interrupting them recklessly, MySQL started acting up.

# Solution
It was resolved by setting the log size in the MySQL configuration file.

`innodb_log_file_size=5M`

# Thoughts
Let's avoid reckless actions!

# References
+ ["Plugin 'InnoDB' registration as a STORAGE ENGINE failed" error prevents MySQL service from starting](https://support.plesk.com/hc/ja/articles/213398969--Plugin-InnoDB-registration-as-a-STORAGE-ENGINE-failed-%E3%81%A8%E3%81%84%E3%81%86%E3%82%A8%E3%83%A9%E3%83%BC%E3%81%A7-MySQL-%E3%82%B5%E3%83%BC%E3%83%93%E3%82%B9%E3%82%92%E8%B5%B7%E5%8B%95%E3%81%A7%E3%81%8D%E3%81%BE%E3%81%9B%E3%82%93)

+ [Changing innodb_log_file_size casually can be fatal](http://masasuzu.hatenadiary.jp/entry/2014/06/13/innodb_log_file_size%E3%82%92%E6%B0%97%E8%BB%BD%E3%81%AB%E5%A4%89%E3%81%88%E3%82%8B%E3%81%A8%E6%AD%BB%E3%81%AC%E3%82%88)
