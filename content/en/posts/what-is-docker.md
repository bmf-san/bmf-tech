---
title: What is Docker? A Beginner's Complete Guide to Containers
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

- A platform developed by Docker Inc. for creating, distributing, and running container-based virtual environments
- Uses Linux container technology

  - Containers utilize the host machine's kernel to isolate processes and users

    - Lightweight and fast

- Middleware and various environment settings can be managed as code (=Infrastructure as Code)

  - Regardless of local or production environments

    - Anyone can create the same environment
    - Easy redistribution and reuse of environments

- The essence of Docker

  - Linux Container (LXC)
  - A container management tool to simplify customization of containers in LXC (such as creating configuration files and installing files via shell scripts)

- Docker for Mac
 
  - Uses the virtualization tool HyperKit, which is included by default on Mac, to launch a virtual machine and run Linux to enable Docker

# What is a Container (Linux Container)

- A set of processes isolated from the rest of the system

  - Runs from a separate image that provides all the files necessary to support the processes

- Shares the OS and kernel, isolating application processes from the rest of the system (runs on a single OS)

# Differences Between Other Virtual Environments and Containers

## General PC

- Configuration

  ```
  [Host OS]
  [Hard Disk]
  ```

## Host OS Type (sometimes defined as hypervisor type using host OS in a broad sense)

- Configuration

  ```
  [Guest OS]
  [Virtualization Software]
  [Host OS]
  [Hard Disk]
  ```

- Advantages

  - Easy to achieve virtualization
  - High freedom in OS selection

- Disadvantages

  - High consumption of disk and memory

## Hypervisor Type (using hard disk pattern)

- Configuration

  ```
  [Guest OS]
  [Hypervisor]
  [Hard Disk]
  ```

- Advantages

  - Can control hardware directly without needing host OS processing, resulting in faster processing speed

- Disadvantages

  - Not easy to achieve virtual environments (may not be able to use the host OS as is, or may require a dedicated physical server)

## Container Type

```
[Container Management Software]
[Host OS]
[Hard Disk]
```

- Advantages

  - Easy to achieve virtual environments
  - Low consumption of disk and memory

- Disadvantages

  - Can only use OS that utilizes the Linux kernel

# Overview of Docker Images and Containers

```
                Docker repository(Ex. Docker Hub)

                         ↓ (pull)

Dockerfile   →     Docker Image    →    Docker Container
                 (build)                           (run)   
                                      ↓ (commit)

                             Docker Image    →    Docker Container
                                                       (run)
```

※ Differences between Host OS Type and Hypervisor Type
[Think IT - What are the differences between Host Type and Hypervisor Type? Overview of VMware vSphere Hypervisor](https://thinkit.co.jp/story/2012/10/17/3722) is a useful reference.

- Host OS Type
    - Install virtualization software on the OS
    - Access to hardware is via the host OS, causing overhead and not fully utilizing performance.
- Hypervisor Type
    - Install directly on the server without needing a host OS.
    - Can directly control hardware, making it easier to achieve performance.

# Basic Commands
- `docker build`
   - Create an image from a Dockerfile

- `docker run`
   - Launch a container from an image

- `docker commit`
   - Create an image from a container

# Additional Information

- What is a Kernel

  - Software responsible for the basic functions of the OS

    - It's generally difficult to use the OS with just the kernel, so it's commonly used in combination with other software

  - Kernel + Software = Distribution

    - CentOS or Ubuntu → Linux kernel + software distribution

# References

- [Docker Official Site](https://www.docker.com/what-docker)
- [Introduction to Docker (Part 1) ~What is Docker and What are its Benefits~](https://knowledge.sakura.ad.jp/13265/)
- [What is a Linux Container](https://www.redhat.com/ja/topics/containers/whats-a-linux-container)
- [What is Server Virtualization? Explaining the Mechanism, Advantages, and Disadvantages in an Easy-to-Understand Manner](https://www.kagoya.jp/howto/rentalserver/virtualization/)
- [Understanding LXC (Linux Containers) in 15 Minutes and Basic Usage for Engineers 2014.06.16](https://knowledge.sakura.ad.jp/2108/)
- [Using LXC for Privilege Separation and Template Customization](https://knowledge.sakura.ad.jp/2163/)
- [Overview of Docker Images and Containers and Various Commands (May be Updated)](https://yoshinorin.net/2016/10/03/docker-image-and-container-command/)
- [Understanding the Inside of Linux Containers / OSC 2018 Kyoto](https://speakerdeck.com/tenforward/osc-2018-kyoto)
- [Think IT - What are the Differences Between Host Type and Hypervisor Type? Overview of VMware vSphere Hypervisor](https://thinkit.co.jp/story/2012/10/17/3722)
