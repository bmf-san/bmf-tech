---
title: Vagrant+CentOS7.3+Ansible
description: "Configure Vagrant with CentOS and Ansible provisioning using SSH connections and automated playbook deployment."
slug: vagrant-centos-ansible-setup
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Ansible
  - CentOS
  - Vagrant
translation_key: vagrant-centos-ansible-setup
---


# Overview
This is the first step in setting up a Vagrant environment using Ansible.
We will prepare an environment that can perform provisioning.

# Environment
+ Vagrant 1.9.1
+ CentOS 7.3
+ Ansible 2.2.1.0

# Preparing the CentOS 7.3 Vagrant Box
We will set up the Vagrant environment in any directory (for this example, we will use centos7.3).

`vagrant box add https://atlas.hashicorp.com/centos/boxes/7`
`vagrant init`


Directory structure up to this point

```
centos7.3/
├── .vagrant.d
├── Vagrantfile
```

*Since the default box name contains a slash, it might be better to rename it.

# Installing Ansible and Preparing for Provisioning
There are various ways to install Ansible, such as using Homebrew, pip, or downloading from GitHub.
Install Ansible on the host OS using one of these methods.
I installed it using pip for no particular reason.

I will skip the installation details.

Once Ansible is installed, prepare a `provisioning` directory and create two files: `hosts` and `site.yml`.

Next, since we will use Ansible to SSH into Vagrant, prepare the SSH configuration file in the root of the development directory.
`vagrant ssh-config > ssh.config`

*The location of ssh.config can be anywhere you prefer.

Contents of the hosts file

```
[vagrants]
127.0.0.1 ansible_ssh_port=2200 ansible_ssh_user=vagrant ansible_ssh_private_key_file=.vagrant/machines/default/virtualbox/private_key
```

Contents of the site.yml file

```
---
- hosts: vagrants
  become: true
  user: vagrant
  tasks:
     - name: install packages zsh
       ping:
```


Directory structure up to this point

```
centos7.3/
├── Vagrantfile
├── provisioning
│   ├── hosts
│   └── site.yml
└── ssh.config
```

*The ssh.config file does not necessarily have to be within this directory; you can write it in ~/.ssh/config, for example.

# Provisioning with Ansible
Let's try running the provisioning.

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

Wow! This is so much fun!


# Challenges Faced
I encountered quite a few issues when trying to SSH into Vagrant with Ansible, but I was helped by a question on Teratail.
[Failed to SSH into Vagrant with Ansible](https://teratail.com/questions/46676)

# Thoughts
For now, I have set up an environment that can perform provisioning using Ansible locally, so I can focus on creating the application.
It seems necessary to configure provisioning for different hosts like VPS and to familiarize myself with best practices.
Next time, I plan to create a Laravel environment and write an article about it (planned).

# Additional Note
It seems that when you run `vagrant destroy` and then `vagrant up` to rebuild, the port number in `ssh-config` may change.
If one day you suddenly find that provisioning is not working, it might be a good idea to check the SSH connection information.

# References
+ ~~Ansible Documentation~~
+ ~~Getting Started with Ansible on CentOS 7~~
