---
title: Copy Local Files to Remote with Ansible
description: 'Learn Ansible copy module syntax to transfer directory contents to remote servers with proper file permissions configuration.'
slug: copy-local-file-to-remote-ansible
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Ansible
translation_key: copy-local-file-to-remote-ansible
---

# Overview
This task involves copying files (contents of a directory) from a local machine to a remote location using Ansible.

# Playbook

```
---
- hosts: vps
  become: yes
  user: root
  tasks:
  - name: Copy a directory
    copy:
     src: /path/to/directory/
     dest: /usr/local/bin/
     mode: u+x
```

This task copies all contents of a directory to the remote `/usr/local/bin` directory. Permissions are also specified.

# Thoughts
There don't seem to be any tricky points; it works as documented.

# References
+ [copy - Copies files to remote locations.](http://docs.ansible.com/ansible/copy_module.html)
