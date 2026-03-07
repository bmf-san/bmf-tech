---
title: Ubuntu Initial Setup Notes
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
I often forget these steps, so I'm making a note of them.
I will add more as needed.

# Environment
- ConoHa
- Ubuntu 18.04.2 LTS (Bionic Beaver)

# Preparation
## Confirm Ubuntu Installation & Root Login
Prepare an Ubuntu server on ConoHa and confirm that you can log in as root.

## Prepare SSH Keys on Client Side
Create a private key and a public key.

```bash
ssh-keygen -t rsa
ssh root@<ip address>
```

# Setup
## Server Update
Make sure to update the server.

```bash
sudo apt update && sudo apt upgrade -y
```

## Create User
Create a user with sudo privileges.

```bash
adduser <username>
usermod -aG sudo <username>
```

Check if the user is part of the wheel group.
```bash
groups <username>
```

*To check the list of users*
```bash
cat /etc/passwd
```

## Transfer Public Key to Server
Log in with the created user.
```bash
su <username>
```

Prepare the `.ssh` directory.
```bash
mkdir .ssh
touch .ssh/authorized_keys
chmod 700 .ssh
chmod 600 .ssh/authorized_keys
```

Paste the public key created on the client side into `./ssh/authorized_keys`.

## Configure sshd_config and Open Ports
Change the SSH configuration.

```bash
sudo vi /etc/ssh/sshd_config
```

```
Port 5005                                     // Change from default 22 to any number
PermitRootLogin no                    // Change from yes to no
PubkeyAuthentication yes         // Change from no to yes
PasswordAuthentication no      // Change from yes to no
UserPAM no                               // Change from yes to no
```

Restart SSH.
```bash
sudo /etc/init.d/ssh restart
```

Next, open the ports.

```
sudo ufw allow 5005
sudo ufw allow 443
sudo ufw default deny    // Default setting may be deny...
sudo ufw enable
```

Check the port settings.
```bash
sudo ufw status
```

# Confirm SSH Connection
Edit the `~/.ssh/config` file like this:

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

Confirm that you can connect via SSH using `ssh conoha-demo`.

# Thoughts
I feel like I wrote similar notes when I first touched CentOS a long time ago.