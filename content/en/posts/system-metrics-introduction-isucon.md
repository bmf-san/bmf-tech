---
title: Introduction to System Metrics with ISUCON Environment
slug: system-metrics-introduction-isucon
date: 2024-04-14T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - ISUCON
description: A summary of regular study sessions using the ISUCON environment to properly understand system metrics.
translation_key: system-metrics-introduction-isucon
---

# Overview
We regularly hold study sessions using the ISUCON environment to learn how to properly monitor system metrics, and this is a summary of those sessions.

Using [The Basics of Infrastructure Every Web Engineer Should Know: From Infrastructure Design to Configuration, Monitoring, and Tuning](https://book.mynavi.jp/ec/products/detail/id=33857) as a reference book, we conduct hands-on sessions from Chapter 5 onwards.

The environment is set up on Conoha using the ISUCON8 image.

Plan: Memory 512MB/CPU 1Core

```sh
[root@160-251-16-96 ~]# cat /etc/redhat-release 
CentOS Linux release 7.5.1804 (Core) 
```

# Checking the Current System Status
cf. Chapter 5

## iptables
Check the server's listening ports and firewall settings to understand which ports are open.

```sh
iptables -nv -L
```

In this environment, it was confirmed that ports 22 and 80 are open.

The pre-installed ufw (Uncomplicated FireWall) on Ubuntu is a wrapper for iptables.

## ss
Next, check the listening ports from external sources.

Use the network status check command ss (formerly netstat) to verify.

```sh
ss -lnp
```

It was confirmed that h2o is listening on port 80 without an IP address, the isucon application is listening on port 8080 without an IP address, mysqld is listening on port 3306 without an IP address, and sshd is listening on port 22 without an IP address.

Entries with :: at the beginning of the Local Address:Port field are also listening on IPv6.

You can also use lsof as an alternative to ss.

```sh
lsof -i
lsof -i:port number
```

## ps
Check the processes and their startup commands.

```
ps aufx | grep -v grep | grep -e isucon -e h2o -e mysql
```

It was confirmed that isucon is running as the isucon user with `/home/isucon/torb/webapp/perl/local/bin/plackup`, h2o is running as root with `perl -x /usr/share/h2o/start_server --pid-file=/var/run/h2o/h2o.pid --log-file=/var/log/h2o/error.log --port=0.0.0.0:80 -- /usr/sbin/h2o -c /etc/h2o/h2o.conf`, and mysql is running as the mysql user with `/bin/sh /usr/bin/mysqld_safe --basedir=/usr`.

## df
Check disk usage.

```sh
df -h
```

It was confirmed that the disk capacity is 30G and 22% is used.

You can list directories using capacity like this:

```sh
df -sh /*
```

## top
Check CPU usage, memory usage, and processes with high CPU usage.

```sh
top -b -d 1 -n 1
```

## dstat
Check CPU usage, network usage, disk I/O, paging, and trends.

```sh
dstat -taf 1 10
```

Running the benchmarker while checking confirmed that there is a load on disk I/O.

# Status Monitoring
// TODO:: Continuously updating
