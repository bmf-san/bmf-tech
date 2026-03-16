---
title: Ansibleでローカルのファイルをリモートにコピーする
description: Ansibleでローカルのファイルをリモートにコピーする
slug: copy-local-file-to-remote-ansible
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Ansible
translation_key: copy-local-file-to-remote-ansible
---


# 概要
Ansibleでローカルにあるファイル（ディレクトリの中身）をリモートにコピーするタスクです。

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

ディレクトリの中身をリモートの`/usr/local/bin`以下に全てコピーするタスクです。パーミッションも指定しています。

# 所感
ドキュメント通りで特にハマるようなポイントはなさそうです。

# 参考
+ [copy - Copies files to remote locations.](http://docs.ansible.com/ansible/copy_module.html)

