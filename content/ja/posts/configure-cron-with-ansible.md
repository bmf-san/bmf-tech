---
title: Ansibleでcronを設定
description: Ansibleでcronを設定の手順と実践例を詳しく解説します。
slug: configure-cron-with-ansible
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Ansible
  - cron
translation_key: configure-cron-with-ansible
---


# 概要
Ansibleでcronを設定するタスクです。

# Playbook

以下は毎分タスクを実行する設定の例です。

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

注意点としては、`*/1`と書きたい場合はダブルクォーテーションで囲う必要がある点です。囲わないとシンタックスエラーが出ます。（YAMLの勉強不足。。ｗ）

# 所感
cronの設定も問題なくすんなりいけました。

# 参考
+ [cron - Manage cron.d and crontab entries.](http://docs.ansible.com/ansible/cron_module.html)

