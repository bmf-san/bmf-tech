---
title: Note on Volumes Not Being Deleted Even After Removing Docker Images and Containers
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
I often forget that volumes are not deleted even after removing containers and images in Docker, so I'm writing this note.

I usually use docker-compose.

`docker-compose build`
`docker-compose up -d`

Then, I clean up with:

`docker rm **`
`docker rmi **`

However, it seems there is an option to delete mounted volumes.

# Solution
Check if volumes remain.
`docker volume ls`

`docker volume rm **`

# Additional Notes
There is a way to clean up all containers, network images, and volumes described in docker-compose.yml at once using docker-compose.

`docker-compose down --rmi all -v`

# References
- [Docker Volume Issues -v --rm -d Garbage Remains, Container Won't Start](https://stlisacity.hatenablog.com/entry/2018/09/10/145101)