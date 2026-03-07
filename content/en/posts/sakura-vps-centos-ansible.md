---
title: Sakura VPS + CentOS 7.3 + Ansible
slug: sakura-vps-centos-ansible
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Ansible
  - CentOS
  - Sakura VPS
  - IaaS
translation_key: sakura-vps-centos-ansible
---

# Overview
Automating the initial setup of Sakura VPS with Ansible.

# Environment
+ Sakura VPS
+ CentOS 7.3
+ Ansible 2.2.1.0

# Prerequisites
+ Understanding the flow of initial setup for Sakura VPS.
  + [Initial Setup of Sakura VPS](google.com)

# Installing CentOS 7 on Sakura VPS
Select `OS Installation > Custom OS Installation` from the Sakura VPS console screen to install CentOS 7. Once the installation starts, you can open the console screen for CentOS 7 installation (either the HTML5 version of VNC console or the Java Applet version), so choose the one that suits your environment.

During the CentOS 7 installation, you will need to configure language settings and initialize the disk. There is a screen for setting the root user password and creating a new user, but since the new user creation will be done with Ansible, you only need to set the root user password.

Next, send the public key from the Ansible host to the Sakura VPS. (Please create the key in advance; this part is omitted here.) This key is for the new user that will be created by Ansible.
`ssh-copy-id -i ~/.ssh/id_rsa.pub root@123.45.678.910`

If you can SSH into Sakura VPS with `ssh root@123.45.678.901`, you're all set.

# Initial Setup of Sakura VPS with Ansible

## Define the hosts file
hosts

```
[sakura]
123.45.678.910 ansible_ssh_user=root ansible_ssh_private_key_file=~/.ssh/id_rsa
```

## Define the Playbook
The tasks will look like this:

+ Create a new user
+ Create an authorize_keys file
+ Adjust permissions of the authorize_keys file
+ Grant sudo privileges to the wheel group
+ Prohibit SSH access for the root user
+ Change the SSH connection port number
+ Change the TCP port for iptables
+ Shutdown the SSH port
+ Set SELinux to disabled

One point to note is that `ssh_user_password` needs to be specified as an encrypted value using `openssl`.

`openssl passwd -salt hoge -1 moge`

The Playbook was greatly inspired by reference sites. m(_ _)m

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
     insertafter: "^# %wheel\s+ALL=\(ALL\)\s+NOPASSWD:\s+ALL"
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

## Define the handler
Define the handler.

main.yml

```
---
- name: restart sshd
  service: name=iptables start restarted
```

# Execute Ansible
`ansible-playbook sakura.yml -i hosts -k -c paramiko`

Once all tasks are completed, restart the server to finish.

# References
+ [Initial Setup of Sakura VPS with Ansible](http://qiita.com/meganii/items/8c91a43e52bd5d61cdde#hosts)
+ [Setting up a web server at lightning speed with Ansible.](http://qiita.com/sak_2/items/7dd3dcd864f93103f0db#%E3%81%95%E3%81%8F%E3%82%89vps%E5%81%B8%E3%81%AE%E6%BA%96%E5%82%99)