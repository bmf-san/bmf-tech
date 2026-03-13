---
title: Ubuntu Initial Setup Notes
description: "Set up Ubuntu servers with SSH key authentication, custom ports, and firewall rules for secure remote access management."
slug: ubuntu-initial-setup-notes
date: 2019-02-15T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Ubuntu
translation_key: ubuntu-initial-setup-notes
---



# Overview
I often forget, so I'm taking notes. I'll add more if needed.

# Environment
- ConoHa
- Ubuntu 18.04.2 LTS (Bionic Beaver)

# Preparation
## Install Ubuntu & Verify Root Login
Prepare an Ubuntu server on ConoHa and verify that you can log in as root.

## Prepare SSH Keys on the Client Side
Create a private key and a public key.

`ssh-keygen -t rsa`
`ssh root@<ip address>`

# Setup
## Server Update
Make sure to update.

`sudo apt update && sudo apt upgrade -y`

## Create a User
Create a user with sudo privileges.

`adduser <username>`
`usermod -aG sudo <username>`

Verify if the user belongs to the wheel group.
`groups <username>`

※ Check the list of users
`cat /etc/passwd`

## Transfer the Public Key to the Server
Log in as the created user.
`su <username>`

Prepare the `.ssh` directory.
`mkdir .ssh`
`touch .ssh/authorized_keys`
`chmod 700 .ssh`
`chmod 600 .ssh/authorized_keys`

Paste the public key created on the client side into `./ssh/authorized_keys`.

## Configure sshd_config and Open Ports
Change the SSH settings.

`sudo vi /etc/ssh/sshd_config`

```
Port 5005                                     // Change from default 22 to a custom number
PermitRootLogin no                    // Change from yes to no
PubkeyAuthentication yes         // Change from no to yes
PasswordAuthentication no      // Change from yes to no
UserPAM no                               // Change from yes to no
```

Restart SSH.
`sudo /etc/init.d/ssh restart`

Continue to open ports.

```
sudo ufw allow 5005
sudo ufw allow 443
sudo ufw default deny    // Default setting might be deny...
sudo ufw enable
```

Check port settings.
`sudo ufw status`

# Verify SSH Connection
Edit the `~/.ssh/config` file like this.

```
ServerAliveInterval 300
TCPKeepAlive yes
AddKeysToAgent yes
ForwardAgent yes
UseKeychain yes

Host conoha-demo
    Hostname    <ip address>
    User         <username>
    Port         5005   // The custom port number set above
    IdentityFile ~/.ssh/<pubkey name>
```

Verify SSH connection with `ssh conoha-demo`.

# Thoughts
I feel like I wrote a similar note when I first touched CentOS.