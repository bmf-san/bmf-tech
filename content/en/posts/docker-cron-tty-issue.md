---
title: Running Docker Commands in Cron Results in 'the input device is not a TTY'
slug: docker-cron-tty-issue
date: 2023-03-17T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Docker
translation_key: docker-cron-tty-issue
---

# Overview
When I tried to run a docker command in cron, I encountered the error "the input device is not a TTY".

An example of what I tried to set in cron is as follows:

```sh
* * * * * user docker exec -it container-name mysqldump dbname -uuser -ppassword  > backup.sql
```

# Cause
The `-t` option allocates a TTY, and the `-i` option opens standard input, but these were unnecessary for cron execution.

# Solution
Removing the `-it` options resolves the issue.