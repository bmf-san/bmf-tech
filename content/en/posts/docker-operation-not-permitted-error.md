---
title: Handling the Docker 'Operation not permitted' Error
slug: docker-operation-not-permitted-error
date: 2019-09-27T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Docker
  - Docker Compose
  - Linux
  - Linux Capabilities
  - Seccomp
  - Tips
description: How to resolve the 'Operation not permitted' error encountered when using Docker Compose.
translation_key: docker-operation-not-permitted-error
---

# Overview
While running golang tests using Docker Compose, I encountered an **Operation not permitted** error.

# Resolution
[Docker Documentation - runtime-privilege-and-linux-capabilities](https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities)

Adjusting the privilege settings of the Docker container resolved the issue.

```yml
  gobel_test_db:
    container_name: "gobel_test_db"
    build: ./docker/mysql
    ports:
      - "3305:3306"
    volumes:
      - mysql_gobel_test_db:/var/lib/mysql:delegated
      - ./docker/mysql/initdb.d/gobel_test_db:/docker-entrypoint-initdb.d
    environment:
      - MYSQL_DATABASE=gobel_test
      - MYSQL_ROOT_PASSWORD=password
    privileged: true  // add this option
```

However, I wasn't entirely sure about the security implications of the above configuration, so I opted for a more restrictive setup.

```yml
  gobel_test_db:
    container_name: "gobel_test_db"
    build: ./docker/mysql
    ports:
      - "3305:3306"
    volumes:
      - mysql_gobel_test_db:/var/lib/mysql:delegated
      - ./docker/mysql/initdb.d/gobel_test_db:/docker-entrypoint-initdb.d
    environment:
      - MYSQL_DATABASE=gobel_test
      - MYSQL_ROOT_PASSWORD=password
    cap_add:
      - SYS_ADMIN
    security_opt:
      - seccomp:unconfined
```

`cap_add` is an option to add Linux capabilities. In this case, it adds permissions for system administration operations.

Linux capabilities are a feature that allows fine-grained control of superuser privileges.

`seccomp` is a security feature in the Linux kernel that restricts system call execution. Here, the setting is `unconfined`, which disables restrictions.

The term "unconfined" literally means "not confined".

# Thoughts
[speakerdeck - Fully Understanding Containers](https://speakerdeck.com/bmf_san/kontenawan-quan-nili-jie-sita)

I studied containers briefly before, but my understanding is still shallow.

# References
- [Qiita - systemd in docker container without --privileged](https://qiita.com/shusugmt/items/92ece6874ba5aeff2b41)