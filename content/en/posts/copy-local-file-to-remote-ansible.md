---
title: Copy Local Files to Remote with Ansible
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
This task copies files (contents of a directory) from local to remote using Ansible.

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

This task copies all contents of the directory to the remote `/usr/local/bin`. Permissions are also specified.

# Thoughts
There doesn't seem to be any particular points where one could get stuck, as per the documentation.

# References
+ [copy - Copies files to remote locations.](http://docs.ansible.com/ansible/copy_module.html)