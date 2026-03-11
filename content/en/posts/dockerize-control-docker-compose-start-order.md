---
title: Controlling Container Startup Order with Dockerize in Docker Compose
slug: dockerize-control-docker-compose-start-order
date: 2019-09-17T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Docker
  - Docker Compose
  - Dockerize
translation_key: dockerize-control-docker-compose-start-order
---

# Overview
This is a note on how to control the startup order of containers launched with Docker Compose using Dockerize.

[github.com - jwilder/dockerize](https://github.com/jwilder/dockerize)

# Why Use Dockerize
It is also possible to use a pure bash script called [wait-for-it](https://github.com/vishnubob/wait-for-it) instead of Dockerize.

cf. [Docker-docs-ja - Controlling Startup Order](http://docs.docker.jp/compose/startup-order.html)

The purpose of using Dockerize is when you want to intentionally control the startup order of multiple containers. 

For example, if there is an application container and a test database container, the database container needs to start before the application container when the application container performs tests that use the DB.

In short, it aims to resolve the dependency of the startup order between containers.

Docker Compose has options like `depends_on` and `links`, but `depends_on` controls the creation order of containers, while `links` performs name resolution between containers in addition to the functionality of `depends_on`, but neither controls the startup order.

I later learned that `links` is automatically executed from version 2 onwards and has become legacy.

# Example
Here is an example from one of my applications.

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

gobel_app container wants to wait for the startup of gobel_test_db. 
The entrypoint specifies the dockerize command.

Dockerfile used for building the gobel_app container.

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

It includes the installation of dockerize.

# Thoughts
It can be easily introduced.
While it might be fine to prepare a script like wait-for-it on your own, I thought it would be reasonable to include dockerize considering that container management could become more complex.