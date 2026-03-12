---
title: Docker CLI Now Supports Docker Compose
description: An in-depth look at Docker CLI Now Supports Docker Compose, covering key concepts and practical insights.
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



Recently, an update has made it possible for the Docker CLI to support Docker Compose.

cf. [docs.docker.com - Compose CLI Tech Preview](https://docs.docker.com/compose/cli-command/)

Since it's a Tech Preview, it's not yet recommended for production use.

For compatibility details, see here.
[docs.docker.com - Compose command compatibility with docker-compose](https://docs.docker.com/compose/cli-command-compatibility/)

Docker Compose is implemented in Python, but the supported Docker Compose this time is apparently made with Golang.

cf. https://github.com/docker/compose
cf. https://github.com/docker/compose-cli
