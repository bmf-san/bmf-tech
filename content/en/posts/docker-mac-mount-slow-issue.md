---
title: Addressing Slow Mounting Issues in Docker for Mac
description: 'Resolve slow Docker for Mac volume mount performance caused by the macOS filesystem API. Workarounds: delegated/cached mount options, docker-sync, or using a Linux-based VM.'
slug: docker-mac-mount-slow-issue
date: 2018-08-19T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Docker
  - Tips
translation_key: docker-mac-mount-slow-issue
---



# Overview
Mounting in Docker for Mac is slow. npm and other operations are painfully slow. This is a memo.

# Cause
Refer to the comments from Docker staff. (See the middle section of the linked page)

[Docker - File access in mounted volumes extremely slow, CPU bound](https://forums.docker.com/t/file-access-in-mounted-volumes-extremely-slow-cpu-bound/8076/158)

It seems related to the file system API of MacOS.

# Solutions
- Use Windows or Linux
- docker-sync
- Set up a virtual environment built on an OS different from MacOS (like Vagrant)
- Utilize options like cached, delegated, consistent (~~Docker - Performance tuning for volume mounts (shared filesystems)~~)

# Thoughts
I feel like I want to use Linux.