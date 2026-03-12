---
title: Writing to Remote Files with Ansible
description: An in-depth look at Writing to Remote Files with Ansible, covering key concepts and practical insights.
slug: write-remote-file-ansible
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Ansible
translation_key: write-remote-file-ansible
---


# Overview
A task to write to remote files using Ansible. Frequently used.

# Playbook

```
---
- hosts: vps
  become: yes
  user: root
  tasks:
  - name: Add text
    blockinfile:
     dest: /path/to/file
     insertafter: '^# Add Here'
     content: |
        # New Line
         Here is a new line.
```

# Thoughts
It's easy to write.

# References
+ [lineinfile - Ensure a particular line is in a file, or replace an existing line using a back-referenced regular expression.](http://docs.ansible.com/ansible/lineinfile_module.html)
