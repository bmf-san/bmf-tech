---
title: Setting Up a Docker Environment on Sakura VPS
description: 'Step-by-step guide to installing Docker CE on Sakura VPS running CentOS. Covers yum repository setup, stable vs. edge channel configuration, and Docker CE installation commands.'
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
Set up a Docker environment on Sakura VPS. Initial server settings, etc., are omitted.

# Setup Procedure
All operations are assumed to be performed by a regular user with sudo privileges.

Docker has a free CE version and a commercial EE version, but we will use the CE version this time.

# Set Up the Repository

## Installation

```
sudo yum install -y yum-utils \
device-mapper-persistent-data \
lvm2
```

## Set Up the Stable Repository

```
sudo yum-config-manager \
--add-repo \    https://download.docker.com/linux/centos/docker-ce.repo
```

## Set Up the Edge and Test Repositories

```
sudo yum-config-manager --enable docker-ce-edge
```

```
sudo yum-config-manager --enable docker-ce-test
```

Since we only want to use stable, disable them with `--disable`.

```
sudo yum-config-manager --disable docker-ce-edge
```

```
sudo yum-config-manager --disable docker-ce-test
```

# Install Docker CE

```
sudo yum install docker-ce
```

You can check the available versions for installation with the following command.

```
yum list docker-ce --showduplicates | sort -r
```

To install a specified version, specify the version as follows.

```
sudo yum install docker-ce-<VERSION STRING>
```

# Start Docker

```
sudo systemctl start docker
```

Check if it's running.

```
sudo docker run hello-world
```

# Uninstall Docker CE

```
sudo yum remove docker-ce
```

Docker images, volumes, containers, and configuration files are not automatically deleted, so manually delete the following directory.

```
sudo rm -rf /var/lib/docker
```

# References
- [docker docs - Get Docker CE for CentOS](https://docs.docker.com/install/linux/docker-ce/centos/)
