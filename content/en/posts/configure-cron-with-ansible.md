---
title: Setting Up Cron with Ansible
description: 'Configure cron jobs with Ansible playbooks using YAML syntax and the cron module for task scheduling and automation.'
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
This task involves setting up cron with Ansible.

# Playbook

Below is an example configuration to execute a task every minute.

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

A point to note is that if you want to write `*/1`, you need to enclose it in double quotes. If not enclosed, a syntax error will occur. (Need to study YAML more... lol)

# Thoughts
The cron setup went smoothly without any issues.

# Reference
+ [cron - Manage cron.d and crontab entries.](http://docs.ansible.com/ansible/cron_module.html)
