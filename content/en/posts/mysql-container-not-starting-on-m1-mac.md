---
title: Cannot Start MySQL 8.0.17 Container on M1 Mac
description: An in-depth look at Cannot Start MySQL 8.0.17 Container on M1 Mac, covering key concepts and practical insights.
slug: mysql-container-not-starting-on-m1-mac
date: 2021-09-05T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - Docker
  - Docker Compose
  - MySQL
  - M1
  - Tips
translation_key: mysql-container-not-starting-on-m1-mac
---



I recently upgraded to an M1 Mac and tried to run a MySQL container for my local development environment, but it didn't work.

Here's the error:

```sh
runtime: failed to create new OS thread (have 2 already; errno=22)
```

Since it was a Go error, I guessed it might be an architecture-related issue.

I checked Docker Hub for a version newer than 8.0.17 and found that the latest patch version 8.0.26 had been released.

https://hub.docker.com/layers/mysql/library/mysql/8.0.26/images/sha256-75e71ac9b332935f396d85965213a64f1bd6fc7c42e9900b106f7af462c599b0?context=explore

It seems it was released just two days ago.

It looks like MySQL 8.0.26 includes support for M1, so it might work with this version.
cf. https://dev.mysql.com/doc/relnotes/mysql/8.0/en/news-8-0-26.html

```
macOS: It is now possible to build MySQL for macOS 11 on ARM (that is, for Apple M1 systems). (Bug #32386050, Bug #102259)
```

Specify --platform and set the image to 8.0.26.

```yaml
FROM --platform=linux/amd64 mysql:8.0.26

ADD ./conf.d/my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]
```

This worked for now.
