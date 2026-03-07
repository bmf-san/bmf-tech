---
title: Building a Docker Environment on Sakura VPS
slug: docker-environment-sakura-vps
date: 2018-06-09T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Docker
  - Sakura VPS
translation_key: docker-environment-sakura-vps
---

# Overview
Building a Docker environment on Sakura VPS.
Initial server setup will be omitted.

# Setup Steps
All operations will be performed by a regular user with sudo privileges.

Docker has a free CE version and a commercial EE version, but this time we will use the CE version.

# Setting Up the Repository

## Installation

```
sudo yum install -y yum-utils \
device-mapper-persistent-data \
lvm2
```

## Setting Up the Stable Repository

```
sudo yum-config-manager \
--add-repo \    https://download.docker.com/linux/centos/docker-ce.repo
```

## Setting Up the Edge and Test Repositories

```
sudo yum-config-manager --enable docker-ce-edge
```

```
sudo yum-config-manager --enable docker-ce-test
```

Since we only want to use stable, we will disable the others with `--disable`.

```
sudo yum-config-manager --disable docker-ce-edge
```

```
sudo yum-config-manager --disable docker-ce-test
```

# Installing Docker CE

```
sudo yum install docker-ce
```

You can check the available versions for installation with the following command.

```
yum list docker-ce --showduplicates | sort -r
```

To install a specified version, use the following command with the version specified.

```
sudo yum install docker-ce-<VERSION STRING>
```

# Starting Docker

```
sudo systemctl start docker
```

Check if it is running.

```
sudo docker run hello-world
```

# Uninstalling Docker CE

```
sudo yum remove docker-ce
```

Docker images, volumes, containers, and configuration files will not be automatically deleted, so manually delete the following directory.

```
sudo rm -rf /var/lib/docker
```

# References
- [docker docs - Get Docker CE for CentOS](https://docs.docker.com/install/linux/docker-ce/centos/)