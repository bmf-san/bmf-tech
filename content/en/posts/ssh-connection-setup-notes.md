---
title: SSH Connection Setup Memo
description: A step-by-step guide on SSH Connection Setup Memo, with practical examples and configuration tips.
slug: ssh-connection-setup-notes
date: 2018-09-18T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - ssh
  - sshd
translation_key: ssh-connection-setup-notes
---

# Overview
A memo on how to set up SSH connections.

# Preparation
Connect to the server and create a user belonging to the wheel group.

# Steps
## Host Side
In `~/.ssh/`, use `ssh-keygen` to create a public key and a private key. Here, the public key is created as `id_rsa.pub` and the private key as `id_rsa`.

Copy the contents of the public key.

Create a `config` file in `~/.ssh/`.

Ex.
```
Host bmf
 HostName 123.45.679.012
 User bmf
 Port 22
 IdentityFile ~/.ssh/id_rsa
```

## Server Side
If `~/.ssh/` does not exist, create the directory. Set the permission to 700.
`mkdir .ssh && chmod 700 .ssh`

Next, create a file named `authorized_keys` in `~/.ssh/`. Set the permission to 600. Paste the contents of the public key into `authorized_keys`.

Then adjust the settings in `/etc/ssh/sshd_config`.
Adjust the following settings:
- Port
  - Uncomment. The default is port 22, but for security reasons, it's better to specify a different number.
- PasswordAuthentication
  - Turn off password authentication. Specify no.
- PermitRootLogin
  - Turn off root user login permission. Specify no.

Check if the port number used for SSH connection is open.
`firewall-cmd --list-all`

If not open, open it.
`firewall-cmd --permanent --zone=public --add-port=22/tcp`

Reload.

`firewall-cmd --reload`

# Try Connecting
`ssh bmf`
