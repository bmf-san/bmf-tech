---
title: Unable to Start MySQL 8.0.17 Container on M1 Mac
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

After upgrading to an M1 Mac, I tried to run a MySQL container for my development environment, but it didn't work.

The error looked like this:

```sh
runtime: failed to create new OS thread (have 2 already; errno=22)
```

Since it was a Go error, I guessed that it was not working due to some architecture issue.

I checked Docker Hub to find a newer version than 8.0.17 and confirmed that the latest patch version 8.0.26 has been released.

https://hub.docker.com/layers/mysql/library/mysql/8.0.26/images/sha256-75e71ac9b332935f396d85965213a64f1bd6fc7c42e9900b106f7af462c599b0?context=explore

It seems to have been released just two days ago.

Since MySQL 8.0.26 includes support for M1, it should probably work now.
cf. https://dev.mysql.com/doc/relnotes/mysql/8.0/en/news-8-0-26.html

```
macOS: It is now possible to build MySQL for macOS 11 on ARM (that is, for Apple M1 systems). (Bug #32386050, Bug #102259)
```

I specified the --platform option and set the image to 8.0.26.

```yaml
FROM --platform=linux/amd64 mysql:8.0.26

ADD ./conf.d/my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]
```

This worked for now.