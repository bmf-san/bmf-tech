---
title: Encountered 'Permission denied' Error When Starting MySQL Container on Ubuntu 20.04.2 LTS
slug: permission-denied-when-starting-mysql-container-on-ubuntu
date: 2021-09-12T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - Ubuntu
  - MySQL
  - Tips
translation_key: permission-denied-when-starting-mysql-container-on-ubuntu
---



# Overview
When attempting to start a MySQL container on Ubuntu 20.04.2 LTS, the following error occurs, causing the container to fail to start.

```sh
Could not open file '/var/log/mysql/mysql-error.log' for error logging: Permission denied”
```

# Dockerfile
The Dockerfile where the issue occurred.

docker-compose.yml (partial excerpt)

```yml
version: '3.2'
services:
  mysql:
    container_name: "example-mysql"
    env_file: ./mysql/.env
    build:
        context: "./mysql"
        dockerfile: "Dockerfile"
    ports:
      - "3306:3306"
    volumes:
      - ./mysql/data:/var/lib/mysql
      - ./mysql/initdb.d:/docker-entrypoint-initdb.d
      - ./mysql/log:/var/log/mysql
```

Dockerfile
```yml
FROM --platform=linux/amd64 mysql:8.0.26

ADD ./conf.d/my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]
```

# Checking the Mount Source
```sh
$ ls -la mysql
drwxrwxrwx  7 systemd-coredump example-app 4096 Sep 12 23:10 data
```

An unfamiliar user named `systemd-coredump` appears.

# systemd-coredump
Checking the user on the host, `systemd-coredump` has uid 999.

```sh
$ cat /etc/passwd | grep systemd-coredump
systemd-coredump:x:999:999:systemd Core Dumper:/:/usr/sbin/nologin
```

The user inside the MySQL container likely has uid 999, which is probably the cause?

# Solution
Add `user: 1000:1000` to docker-compose.yml.

docker-compose.yml (partial excerpt)
```yml
version: '3.2'
services:
  mysql:
    container_name: "example-mysql"
    env_file: ./mysql/.env
    build:
        context: "./mysql"
        dockerfile: "Dockerfile"
    ports:
      - "3306:3306"
    volumes:
      - ./mysql/data:/var/lib/mysql
      - ./mysql/initdb.d:/docker-entrypoint-initdb.d
      - ./mysql/log:/var/log/mysql
    user: 1000:1000
```

It might be better to pass the uid and gid from the host instead of hardcoding them.

# Thoughts
This issue did not occur on Docker for Mac, so I'm glad I was able to notice it.

# References
- [gitmemory.com - redis uid issue on ubuntu20.04](https://www.gitmemory.com/issue/docker-library/redis/279/833602815)
- [tsaeki.hatenablog.com - https://tsaeki.hatenablog.com/entry/2020/02/29/182620](https://tsaeki.hatenablog.com/entry/2020/02/29/182620)
- [stackoverflow.com - Permission issue with PostgreSQL in docker container](https://stackoverflow.com/questions/56188573/permission-issue-with-postgresql-in-docker-container)
