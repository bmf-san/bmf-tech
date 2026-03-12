---
title: Production Deployment with VPS, Docker Compose, Docker Machine, Golang, Nginx, and Let's Encrypt
description: 'A step-by-step guide on Production Deployment with VPS, Docker Compose, Docker Machine, Golang, Nginx, and Let''s Encrypt, with practical examples and configuration tips.'
slug: vps-docker-compose-golang-nginx-lets-encrypt
date: 2020-06-07T00:00:00Z
author: bmf-san
categories:
  - Poem
tags:
  - Docker
  - Docker Compose
  - Golang
  - Let's Encrypt
  - Nginx
  - Docker Machine
  - VPS
translation_key: vps-docker-compose-golang-nginx-lets-encrypt
---


# Overview
I wanted to try running a Golang application in production using Docker Compose on a VPS, so I gave it a shot.

# Environment
Here's a summary of the environment I actually tried.

- VPS (Conoha)
- Onamae.com (Domain Management)
- Docker Compose (Separate files for production and local configurations)
- Docker Machine (Used for deployment)
- Let's Encrypt (TLS/SSL)
- Nginx (Reverse Proxy)

# Repository
I created a sample.
[github - bmf-san/go-production-boilerplate](https://github.com/bmf-san/go-production-boilerplate)

On the production server, as long as you create users and open ports, you should be able to deploy for now...

By the way, since downtime occurs during deployment, separate consideration is necessary.

For deployment using docker-machine, this article was helpful.
[Qiita - Deploy an app from Mac to Docker on VPS using Docker Machine](https://qiita.com/momotaro98/items/5b902afea3530b6f0b93)

I got a bit stuck with Let's Encrypt, but it wasn't a particularly difficult problem to solve since it wasn't an issue specific to containers.

# Impressions
Using the generic driver of docker-machine allows for easy deployment. Considering downtime measures might be necessary, but for private application operations, it could be one of the considerations.