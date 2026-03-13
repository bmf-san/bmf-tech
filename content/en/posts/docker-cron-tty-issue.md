---
title: Encountering 'the input device is not a TTY' When Running Docker Command with Cron
description: 'Fix the input device is not a TTY error when running docker exec via cron. The cause is the -it flags: cron has no TTY, so remove -t and -i and use plain docker exec.'
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
When I tried to run a docker command with cron, I encountered the error "the input device is not a TTY".

An example of what I tried to set in cron is as follows:

```sh
* * * * * user docker exec -it container-name mysqldump dbname -uuser -ppassword  > backup.sql
```

# Cause
The `-t` option assigns a TTY, and `-i` opens standard input, but these are unnecessary for cron execution.

# Solution
Removing the `-it` options resolves the issue.
