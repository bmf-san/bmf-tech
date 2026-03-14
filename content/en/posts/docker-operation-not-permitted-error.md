---
title: Handling Docker 'Operation not permitted' Error
description: 'Fix the Docker Operation not permitted error in docker-compose by adding privileged: true or using cap_add: SYS_ADMIN with seccomp:unconfined to grant required Linux capabilities.'
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
translation_key: docker-operation-not-permitted-error
---

# Overview
While running tests in golang using Docker Compose, I encountered an **Operation not permitted** error.

# Solution
[Docker Documentation - runtime-privilege-and-linux-capabilities](https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities)

Adjusting the privilege settings of the Docker container resolves the issue.

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

Since I wasn't entirely sure about the security implications of the above, I configured it to restrict permissions further.

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

`cap_add` is an option to add Linux capabilities, and here it adds permissions for system administration operations.

Linux capabilities are a feature that subdivides superuser privileges.

`seccomp` is a security feature that restricts system call issuance in the Linux kernel.

Here, it is set to unconfined, which means disabled.

Unconfined literally translates to "not confined."

# Thoughts
[speakerdeck - Fully Understanding Containers](https://speakerdeck.com/bmf_san/kontenawan-quan-nili-jie-sita)

I studied containers briefly before, but my understanding is still shallow.

# References
- [Qiita - systemd in docker container without --privileged](https://qiita.com/shusugmt/items/92ece6874ba5aeff2b41)