---
title: "Vagrant+CentOS7.3+Ansible"
slug: "vagrant-centos-ansible-setup"
date: 2017-10-01
author: bmf-san
categories:
  - "インフラストラクチャ"
tags:
  - "Ansible"
  - "CentOS"
  - "Vagrant"
draft: false
---

# 概要
AnsibleでVagrantの環境構築をする最初の一歩です。
プロビジョニングができる環境を整えます。

# 環境
+ Vagrant1.9.1
+ CentOS7.3
+ Ansible2.2.1.0

# CentOS7.3のVagrnat Boxを用意する
任意のディレクトリ（例として今回はcentos7.3）にてVagrant環境を構築します。

`vagrant box add https://atlas.hashicorp.com/centos/boxes/7`
`vagrant init`


ここまでのディレクトリ構成

```
centos7.3/
├── .vagrant.d
├── Vagrantfile
```

※デフォルトのbox名にスラッシュが入っているのでリネームしたほうがいいかもです。

# AnsibleのインストールとProvisioningの準備
Homebrewかpipかgithubからソースを持ってくるか色々やり方があります。
いずれかの方法でansbileをホストOS側にインストールします。
私は何となくpipでインストールしました。

インストールは割愛します。

Ansibleのインストールが完了したら、`provisioning`ディレクトリを用意して、`hosts`、`site.yml`の2つのファイルを作成します。

それから、ansibleでvagrantにsshをするので、sshの設定ファイルを開発ディレクトリ直下に用意しておきます。
`vagrant ssh-config > ssh.config`

※ssh.configの場所は任意の場所でOK

hostsの中身

```
[vagrants]
127.0.0.1 ansible_ssh_port=2200 ansible_ssh_user=vagrant ansible_ssh_private_key_file=.vagrant/machines/default/virtualbox/private_key
```

site.ymlの中身

```
---
- hosts: vagrants
  become: true
  user: vagrant
  tasks:
     - name: install packages zsh
       ping:
```


ここまでのディレクトリ構成

```
centos7.3/
├── Vagrantfile
├── provisioning
│   ├── hosts
│   └── site.yml
└── ssh.config
```

※ssh.configは~/.ssh/configに記述するなど必ずしもこのディレクトリ内でなくともいいと思います。

# AnsibleでProvisioning
プロビジョニングを実行してみます。

`vagrant provision`

```
$ vagrant provision
==> default: [vagrant-hostsupdater] Checking for host entries
==> default: Running provisioner: ansible...
    default: Running ansible-playbook...

PLAY [vagrant] *****************************************************************

TASK [setup] *******************************************************************
ok: [127.0.0.1]

TASK [check ping] **************************************************************
ok: [127.0.0.1]

PLAY RECAP *********************************************************************
127.0.0.1                  : ok=2    changed=0    unreachable=0    failed=0
```

すごーーい! たのしーーー！！


# ハマったところ
vagrantにansibleでsshする時に結構ハマったのですが、teratailの質問に助けられました。
[vagrantにansbileでsshしようとすると失敗する](https://teratail.com/questions/46676)

# 所感
とりあえずローカル環境でAnsibleを使ったプロビジョニングができる環境が整ったので、
タレづくりに専念できそうです。
VPSなどホスト別にプロビジョニングができるように設定したり、ベストプラクティスをかじっておく必要がありそうです。
次回オレオレLaravel環境のタレを作って記事にしたいと思います（予定）

# 追記
`vagrant destroy`を実行して　`vagrant up`で再構築すると、`ssh-config`のポート番号が変わる場合があるみたいです。
ある日突然プロビジョニングできなくなった！なんて時はssh接続情報を確認してみると良いかもです。

# 参考
+ [Ansible Documentation](http://docs.ansible.com/ansible/intro_installation.html)
+ [CentOS 7 - Ansible のはじめ方](http://centos.sabakan.red/entry/2015/07/01/140000)   

