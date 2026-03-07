---
title: Controlling Container Startup Order in Docker Compose with Dockerize
slug: dockerize-control-docker-compose-start-order
date: 2019-09-17T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Docker
  - Docker Compose
  - Dockerize
description: A guide on using Dockerize to control the startup order of containers in Docker Compose.
translation_key: dockerize-control-docker-compose-start-order
---

# Overview
This is a memo on how to control the startup order of containers in Docker Compose using Dockerize.

[github.com - jwilder/dockerize](https://github.com/jwilder/dockerize)

# Why Use Dockerize
Instead of Dockerize, you could also use a pure bash script method like [wait-for-it](https://github.com/vishnubob/wait-for-it).

cf. [Docker-docs-ja - Controlling Startup Order in Compose](http://docs.docker.jp/compose/startup-order.html)

The purpose of using Dockerize is to intentionally control the startup order of multiple containers.

For example, if you have an application container and a test database container, and the application container needs to perform tests using the database, the database container must start before the application container.

In short, the goal is to resolve the dependency of startup order between containers.

Docker Compose provides options like `depends_on` and `links`. However, `depends_on` only controls the order in which containers are created, and `links` adds name resolution between containers in addition to the functionality of `depends_on`. Neither of these options controls the actual startup order.

As I later found out, `links` is automatically executed from version 2 onwards and has become somewhat legacy.

# Example
Here’s an example from one of my applications.

docker-compose.yml

```yaml
version: "3"
services:
  gobel_app:
    container_name: "gobel_app"
    build: ./docker/go
    volumes:
      - ./app:/go/src/github.com/bmf-san/Gobel/app
    ports:
      - "8080:8080"
    depends_on:
      - gobel_db
      - gobel_test_db
    entrypoint:
      - dockerize
      - -timeout
      - 10s
      - -wait
      - tcp://gobel_test_db:3306
    command: realize start
  gobel_db:
    container_name: "gobel_db"
    build: ./docker/mysql
    ports:
      - "3306:3306"
    volumes:
      - mysql_gobel_db:/var/lib/mysql:delegated
      - ./docker/mysql/initdb.d/gobel_db:/docker-entrypoint-initdb.d
    environment:
      - MYSQL_DATABASE=gobel
      - MYSQL_ROOT_PASSWORD=password
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
volumes:
  mysql_gobel_db:
    driver: local
  mysql_gobel_test_db:
    driver: local
```

The `gobel_app` container needs to wait for the `gobel_test_db` to start. The `entrypoint` specifies the Dockerize command.

Here is the Dockerfile used to build the `gobel_app` container:

```Dockerfile
FROM golang:1.13.0-alpine

WORKDIR /go/src/github.com/bmf-san/Gobel/app/

RUN apk add --no-cache git \
    binutils-gold \
    curl \
    g++ \
    gcc \
    gnupg \
    libgcc \
    linux-headers \
    make
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/oxequa/realize
RUN go get -u golang.org/x/lint/golint
ENV DOCKERIZE_VERSION v0.6.0
RUN apk add --no-cache openssl \
 && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz
```

This includes the installation of Dockerize.

# Thoughts
It’s easy to implement. While it might be feasible to prepare your own script like wait-for-it, considering the potential complexity of container management, using Dockerize seems like a reasonable choice.