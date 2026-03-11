---
title: Ansibleでリモートのファイルに書き込みをする
slug: write-remote-file-ansible
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Ansible
translation_key: write-remote-file-ansible
---


# 概要
Ansibleでリモートのファイルに書き込みをするタスク。よく使うやつ。

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

# 所感
さらっとかけますねー

# 参考
+ [lineinfile - Ensure a particular line is in a file, or replace an existing line using a back-referenced regular expression.](http://docs.ansible.com/ansible/lineinfile_module.html)

