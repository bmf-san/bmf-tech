---
title: Setting Up Cron with Ansible
slug: configure-cron-with-ansible
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Ansible
  - cron
translation_key: configure-cron-with-ansible
---

# Overview
This is a task to set up cron with Ansible.

# Playbook

Here is an example configuration that runs a task every minute.

```
---
- hosts: vps
  become: yes
  user: root
  tasks:
  - name: Output recently logined users
    cron:
     name: last.sh
     job: last.sh
     minute: "*/1"
```

One point to note is that if you want to write `*/1`, you need to enclose it in double quotes. If you don't, a syntax error will occur. (Lack of YAML study... lol)

# Thoughts
The cron setup went smoothly without any issues.

# Reference
+ [cron - Manage cron.d and crontab entries.](http://docs.ansible.com/ansible/cron_module.html)