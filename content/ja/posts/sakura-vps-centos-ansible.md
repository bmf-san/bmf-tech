---
title: さくらVPS+CentOS7.3+Ansible
description: さくらVPS+CentOS7.3+Ansible
slug: sakura-vps-centos-ansible
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Ansible
  - CentOS
  - さくらのVPS
  - Iaas
translation_key: sakura-vps-centos-ansible
---


# 概要
AnsibleでさくらVPSの初期セットアップを自動化します。

# 環境
+ さくらVPS
+ CentOS7.3
+ Ansible2.2.1.0

# 前提知識
+ さくらVPSをの初期設定の流れを理解していること。
  + さくらVPSの初期設定

# さくらVPSにCentOS7をインストール
さくらVPSのコンソール画面から`OSインストール>カスタムOSインストール`を選択してCentOS7をインストールしておきます。
インストールが開始されると、CentOS7のインストール用コンソール画面（VNCコンソールのHTML5版かJava Applet版）を開くことができるので、環境に合わせて好きな方を選びます。

CentOS7のインストールでは、言語設定やディスクの初期化など行う必要があります。
rootユーザーのパスワード設定と新規ユーザー作成をする画面がありますが、新規ユーザー作成はAnsibleで行うので、rootユーザーのパスワード設定のみだけでOKです。

次に、公開鍵をansibleホスト側からさくらVPSに送ります。（鍵は事前に作成しておいてください。ここでは割愛します。）
こちらの鍵はAnsibleで新規に作成するユーザー用の鍵です。
`ssh-copy-id -i ~/.ssh/id_rsa.pub root@123.45.678.910`

`ssh root@123.45.678.901`でさくらVPSにssh接続できれば準備OKです。

# AnsibleでさくらVPSの初期設定をする

## hostsファイルを定義
hosts

```
[sakura]
123.45.678.910 ansible_ssh_user=root ansible_ssh_private_key_file=~/.ssh/id_rsa
```

## Playbookを定義
タスク内容はこんな感じです。

+ ユーザーを新規作成する
+ authorize_keysファイルを作成する
+ authorize_keysファイルのパーミッションを調整する
+ wheelグループにsudo権限を与える
+ rootユーザーでのssh接続を禁止する
+ ssh接続のポート番号を変更する
+ iptablesのtcpポートを変更する
+ sshポートをシャットダウンする
+ SELinuxをdisabledに設定する

1点注意点があります。
`ssh_user_password`は`openssl`で暗号化したものを指定する必要があります。

`openssl passwd -salt hoge -1 moge`

Playbookは参考サイトを大いに参考にさせて頂きました。m(_ _)m

init.yml

```
---
- hosts: sakura
  become: yes
  user: root
  vars:
    ssh_user: bmf
    ssh_user_password: hogehogemogemoge
    ssh_port: 50055
  tasks:
  - name: Add a new user
    user:
     name="{{ ssh_user }}"
     groups=wheel
     password="{{ ssh_user_password }}"
     generate_ssh_key=yes
     ssh_key_bits=2048

  - name: Create an authorize_keys file
    command: /bin/cp /home/{{ ssh_user }}/.ssh/id_rsa.pub /home/{{ ssh_user}}/.ssh/authorized_keys

  - name: Change attributes of an authorized_keys file
    file:
     path: /home/{{ ssh_user }}/.ssh/authorized_keys
     owner: "{{ ssh_user }}"
     group: "{{ ssh_user }}"
     mode: 0600

  - name: Allow wheel group to use sudo
    lineinfile:
     dest: /etc/sudoers
     state: present
     insertafter: "^# %wheel\\s+ALL=\\(ALL\\)\\s+NOPASSWD:\\s+ALL"
     line: "%wheel ALL=(ALL) NOPASSWD: ALL"
     validate: "visudo -cf %s"
     backup: yes

  - name: Forbid root to access via ssh
    lineinfile:
     dest: /etc/ssh/sshd_config
     state: present
     regexp: "^PermitRootLogin without-password"
     line: "PermitRootLogin no"
     backrefs: yes
     validate: "sshd -T -f %s"
     backup: yes
    notify:
     - restart sshd

  - name: Permit only specific user to access via ssh
    lineinfile:
     dest: /etc/ssh/sshd_config
     state: present
     insertafter: "^PasswordAuthentication no"
     regexp: "^AllowUsers"
     line: "AllowUsers {{ ssh_user }}"
     validate: "sshd -T -f %s"
     backup: yes
    notify:
     - restart sshd

  - name: Change ssh port number
    lineinfile:
     dest: /etc/ssh/sshd_config
     state: present
     insertafter: "^#Port 22"
     regexp: "^Port"
     line: "Port {{ ssh_port }}"
     validate: "sshd -T -f %s"
     backup: yes
    notify:
     - restart sshd

  - name: Change acceptable tcp port for ssh on iptables
    firewalld: port={{ ssh_port }}/tcp permanent=true state=enabled immediate=yes

  - name: shutdown ssh port
    firewalld: service=sshd permanent=true state=disabled immediate=yes

  - name: disable selinux
    selinux: state=disabled
```

## handlerを定義
ハンドラーを定義します。

main.yml

```
---
- name: restart sshd
  service: name=iptables start restarted
```

# Ansible実行
`ansible-playbook sakura.yml -i hosts -k -c paramiko`

タスクの実行が全て完了したら、サーバーを一度再起動して完了です。

# 参考
+ [AnsibleでさくらVPSの初期設定](http://qiita.com/meganii/items/8c91a43e52bd5d61cdde#hosts)
+ [Ansibleで光の速さのWEBサーバーを光の速さで構成してみる。](http://qiita.com/sak_2/items/7dd3dcd864f93103f0db#%E3%81%95%E3%81%8F%E3%82%89vps%E5%81%B4%E3%81%AE%E6%BA%96%E5%82%99)

