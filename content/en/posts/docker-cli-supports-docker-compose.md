---
title: Docker CLI Now Supports Docker Compose
slug: docker-cli-supports-docker-compose
date: 2021-05-06T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Docker
  - Docker Compose
translation_key: docker-cli-supports-docker-compose
---

Recently, it seems that the Docker CLI has started to support Docker Compose.

cf. [docs.docker.com - Compose CLI Tech Preview](https://docs.docker.com/compose/cli-command/)

Since this is a Tech Preview, it is not recommended for production use yet.

For compatibility details, see here.
[docs.docker.com - Compose command compatibility with docker-compose](https://docs.docker.com/compose/cli-command-compatibility/)

Docker Compose is implemented in Python, but the supported Docker Compose this time appears to be built with Golang.

cf. https://github.com/docker/compose
cf. https://github.com/docker/compose-cli