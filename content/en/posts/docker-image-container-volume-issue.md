---
title: Notes on When Volumes Remain After Deleting Images and Containers in Docker
slug: docker-image-container-volume-issue
date: 2019-04-28T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Docker
  - Docker Compose
translation_key: docker-image-container-volume-issue
---

# Overview
Even after deleting containers and images in Docker, I often forget that volumes are not deleted, so I'm jotting this down as a reminder.

I usually use docker-compose.

```bash
docker-compose build
docker-compose up -d
```

Then I tidy up with:

```bash
docker rm **
docker rmi **
```

However, it seems there is an option to delete the mounted volumes as well.

# Solution
Check if volumes remain:

```bash
docker volume ls
```

Then remove them:

```bash
docker volume rm **
```

# Aside
There is a way to clean up all containers, network images, and volumes defined in the docker-compose.yml at once:

```bash
docker-compose down --rmi all -v
```

# References
- [About Docker Volumes -v --rm -d Garbage Leftover Issue Container Won't Start](https://stlisacity.hatenablog.com/entry/2018/09/10/145101)