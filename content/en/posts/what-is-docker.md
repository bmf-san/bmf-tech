---
title: What is Docker
slug: what-is-docker
date: 2018-04-01T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Docker
  - Virtual Environment
translation_key: what-is-docker
---

# What is Docker

- A platform developed by Docker, Inc. for creating, distributing, and running container-based virtual environments.
- Uses Linux container technology.

  - Containers utilize the host machine's kernel to isolate processes and users.

    - Lightweight and fast.

- Middleware and various environment settings can be managed as code (Infrastructure as Code).

  - Applicable to both local and production environments.

    - Anyone can create the same environment.
    - Easy to redistribute and reuse environments.

- The essence of Docker:

  - Linux Container (LXC).
  - A container management tool that simplifies customization of containers in LXC (such as creating configuration files and installing files via shell scripts).

- Docker for Mac:

  - Uses a virtualization tool called HyperKit, which is included by default on Mac, to launch a virtual machine, start Linux, and enable the use of Docker.

# What is a Container (Linux Container)

- A set of processes isolated from other parts of the system.

  - Executed from individual images that provide all the files necessary to support the processes.

- Shares the OS and kernel, isolating application processes from other parts of the system (runs on a single OS).

# Differences Between Other Virtual Environments and Containers

## General PC

- Configuration

  ```
  [Host OS]
  [Hard Disk]
  ```

## Host OS Type (also defined broadly as hypervisor type using host OS)

- Configuration

  ```
  [Guest OS]
  [Virtualization Software]
  [Host OS]
  [Hard Disk]
  ```

- Advantages:

  - Easy to achieve virtualization.
  - High freedom of OS selection.

- Disadvantages:

  - High disk and memory consumption.

## Hypervisor Type (using hard disk)

- Configuration

  ```
  [Guest OS]
  [Hypervisor]
  [Hard Disk]
  ```

- Advantages:

  - Can control hardware directly without needing the host OS, resulting in faster processing speed.

- Disadvantages:

  - Virtual environments cannot be easily realized (may not be able to use the host OS as is or may require dedicated physical servers).

## Container Type

```
[Container Management Software]
[Host OS]
[Hard Disk]
```

- Advantages:

  - Easy to achieve virtual environments.
  - Low disk and memory consumption.

- Disadvantages:

  - Can only use OS that utilizes the Linux kernel.

# Overview of Docker Images and Containers

```
                Docker repository (Ex. Docker Hub)

                         ↓ (pull)

Dockerfile   →     Docker Image    →    Docker Container
                 (build)                           (run)   
                                      ↓ (commit)

                             Docker Image    →    Docker Container
                                                       (run)
```

*Difference between Host OS Type and Hypervisor Type*
[Think IT - What is the difference between host type and hypervisor type? Overview of VMware vSphere Hypervisor](https://thinkit.co.jp/story/2012/10/17/3722) is a useful reference.

- Host OS Type:
    - Install virtualization software on the OS.
    - Access to hardware goes through the host OS, causing overhead and preventing full performance.
- Hypervisor Type:
    - Installed directly on the server without needing a host OS.
    - Can control hardware directly, making it easier to achieve performance.

# Basic Commands
- `docker build`
   - Create an image from a Dockerfile.

- `docker run`
   - Start a container from an image.

- `docker commit`
   - Create an image from a container.

# Additional Information

- What is a Kernel?

  - Software that plays the role of the basic functions of an OS.

    - It is generally difficult to use an OS with just the kernel, so it is common to use it in combination with other software.

  - Kernel + Software = Distribution

    - CentOS and Ubuntu → Distribution of Linux kernel + software.

# References

- [Official Docker Website](https://www.docker.com/what-docker)
- [Introduction to Docker (Part 1) - What is Docker and what are its benefits?](https://knowledge.sakura.ad.jp/13265/)
- [What is a Linux Container?](https://www.redhat.com/ja/topics/containers/whats-a-linux-container)
- [What is Server Virtualization? An Easy-to-Understand Explanation of Mechanisms, Advantages, and Disadvantages](https://www.kagoya.jp/howto/rentalserver/virtualization/)
- [Understanding LXC (Linux Containers) in 15 Minutes: Mechanisms and Basic Usage for Engineers, June 16, 2014](https://knowledge.sakura.ad.jp/2108/)
- [Using LXC for Permission Separation and Template Customization](https://knowledge.sakura.ad.jp/2163/)
- [Overview of Docker Images and Containers and Various Commands (may be updated periodically)](https://yoshinorin.net/2016/10/03/docker-image-and-container-command/)
- [Let's Learn About the Internals of Linux Containers / OSC 2018 Kyoto](https://speakerdeck.com/tenforward/osc-2018-kyoto)
- [Think IT - What is the difference between host type and hypervisor type? Overview of VMware vSphere Hypervisor](https://thinkit.co.jp/story/2012/10/17/3722)