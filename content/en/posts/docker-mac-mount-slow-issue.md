---
title: Addressing the Slow Mount Issue in Docker for Mac
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
The mounts in Docker for Mac are slow.
npm is too slow and it's painful.
Just some notes.

# Cause
Referencing comments from Docker staff. (See link in the middle)

[Docker - File access in mounted volumes extremely slow, CPU bound](https://forums.docker.com/t/file-access-in-mounted-volumes-extremely-slow-cpu-bound/8076/158)

It seems related to the MacOS file system API.

# Solutions
- Use Windows or Linux
- docker-sync
- Prepare a virtual environment built on a different OS than MacOS (like Vagrant)
- Utilize options like cached, delegated, and consistent ([Docker - Performance tuning for volume mounts (shared filesystems)](https://docs.docker.com/docker-for-mac/osxfs-caching/#performance-implications-of-host-container-file-system-consistency))

# Thoughts
I feel like I want to use Linux.