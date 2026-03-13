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
This is the first step in setting up a Vagrant environment with Ansible. We will prepare an environment capable of provisioning.

# Environment
+ Vagrant1.9.1
+ CentOS7.3
+ Ansible2.2.1.0

# Prepare a CentOS7.3 Vagrant Box
We will set up a Vagrant environment in any directory (for example, centos7.3).

`vagrant box add https://atlas.hashicorp.com/centos/boxes/7`
`vagrant init`

Directory structure so far:

```
centos7.3/
├── .vagrant.d
├── Vagrantfile
```

*Note: Since the default box name contains a slash, it might be better to rename it.*

# Install Ansible and Prepare for Provisioning
There are various ways to get the source from Homebrew, pip, or GitHub. Install Ansible on the host OS using one of these methods. I installed it using pip for no particular reason.

Installation details are omitted.

Once Ansible is installed, prepare a `provisioning` directory and create two files: `hosts` and `site.yml`.

Then, since Ansible will SSH into Vagrant, prepare an SSH configuration file directly under the development directory.
`vagrant ssh-config > ssh.config`

*Note: The location of ssh.config can be anywhere.*

Contents of hosts:

```
[vagrants]
127.0.0.1 ansible_ssh_port=2200 ansible_ssh_user=vagrant ansible_ssh_private_key_file=.vagrant/machines/default/virtualbox/private_key
```

Contents of site.yml:

```
---
- hosts: vagrants
  become: true
  user: vagrant
  tasks:
     - name: install packages zsh
       ping:
```

Directory structure so far:

```
centos7.3/
├── Vagrantfile
├── provisioning
│   ├── hosts
│   └── site.yml
└── ssh.config
```

*Note: ssh.config does not necessarily have to be in this directory; it can be written in ~/.ssh/config, for example.*

# Provisioning with Ansible
Let's try executing provisioning.

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

Amazing! So much fun!!

# Challenges
I struggled quite a bit with SSHing into Vagrant with Ansible, but I was helped by a question on teratail.
[vagrantにansbileでsshしようとすると失敗する](https://teratail.com/questions/46676)

# Thoughts
Now that I have an environment where I can use Ansible for provisioning locally, I can focus on creating roles. I need to set up provisioning for different hosts like VPS and get a grasp of best practices. Next time, I plan to create a role for my custom Laravel environment and write an article about it (planned).

# Additional Notes
When you execute `vagrant destroy` and then `vagrant up` to rebuild, the port number in `ssh-config` might change. If one day you suddenly can't provision, it might be a good idea to check the SSH connection information.

# References
+ [Ansible Documentation](http://docs.ansible.com/ansible/intro_installation.html)
+ [CentOS 7 - Ansible Getting Started](http://centos.sabakan.red/entry/2015/07/01/140000)