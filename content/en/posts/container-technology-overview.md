---
title: "Container Technology Explained: How Docker and OCI Containers Work"
description: 'Learn how container technology works—namespaces, cgroups, OCI standards—and why containers have become essential for modern application deployment.'
slug: container-technology-overview
date: 2023-06-05T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Docker
  - libcontainer
  - lxc
  - lxd
  - Container
translation_key: container-technology-overview
---



# Overview
A summary of container technology. Experimenting with creating and interacting with containers without using Docker.

# What is a Container
- A set of processes that are isolated from the host OS, bundling applications and runtime together.

# History of Containers
1979: chroot introduced in UNIX OS.

2000: [FreeBSD jails](https://www.freebsd.org/doc/handbook/jails.html) appeared in FreeBSD 4.0, an evolution of chroot.

2001: Technology forming the basis of Linux containers appeared in Linux through the [VServer Project](http://linux-vserver.org/Welcome_to_Linux-VServer.org).

2004: LXC 1.0 released. [Linux Containers](https://linuxcontainers.org/)

2008: Docker emerged.

Besides the above, container technologies like Virtuozzo, OpenVZ, HP-UX Container, and Solaris Container also exist.

# Differences Between Containers and Virtualization
- **Containers**
  - A set of processes isolated from the host OS, bundling applications and runtime together.
  - Shares the kernel part of the host OS
    - The OS library part can be chosen by the container

- **Virtualization**
  - Configuration differs between host-based and hypervisor-based, but virtualization allows multiple OS setups, running applications on guest OS.

A rough summary is also available at [bmf-tech - What is Docker](https://bmf-tech.com/posts/Docker%E3%81%A8%E3%81%AF).

## Linux Kernel Features for Realizing Container Technology
### Kernel namespaces
- A feature that separates processes into six types of resources (ipc, uts, mount, pid, network, user)
- A mechanism that makes it appear as if users have their own isolated resources.
- Isolated resources cannot interfere with each other.

### Apparmor and SELinux profiles
- **Apparmor**
  - A type of Linux Security Modules (a framework for security in the Linux kernel).
  - Securely manages application access permissions (mandatory access control)
- **SELinux (Security Enhanced Linux)**
  - A module that adds mandatory access control features to the Linux kernel

### Seccomp policies
- A feature that restricts the issuance of system calls by processes

### Chroots (using pivot_root)
- An operation that changes the root directory for the current process and its child processes
- Processes with changed roots cannot access files outside the range => Realization of process isolation

### Kernel capabilities
- Permission management for processes
- Allows more granular permission management than just root or not root

### CGroups (control groups)
- A feature to group processes for common management

# Docker's Container Technology
Previously, Docker used lxc, but from v0.9, it seems to use libcontainer implemented in Go. (cf. [Docker blog - DOCKER 0.9: INTRODUCING EXECUTION DRIVERS AND LIBCONTAINER](https://blog.docker.com/2014/03/docker-0-9-introducing-execution-drivers-and-libcontainer/) [github - opencontainers/runc/libcontainer/](https://github.com/opencontainers/runc/tree/master/libcontainer))

# Standard Specifications
## OCI (Open Container Initiative)
The [Open Container Initiative](https://opencontainers.org/) is an organization aimed at creating industry standards for containers and runtimes.

It defines the following specifications:

- OCI Runtime Specification
- OCI Image Format Specification
- OCI Distribution Specification

OCI is involved in the specifications of low-level runtimes.
Examples: runC, gVisor, Kata Containers, Nabla Containers, etc.

## CRI (Container Runtime Interface)
[CRI](https://kubernetes.io/ja/docs/concepts/architecture/cri/) defines the interface for communication between kubelet and container runtime.

CRI is involved in the specifications of high-level runtimes.
Examples: docker, containerd, cri-o

# Summary
- Containers are processes with isolated resources
- Containers share the kernel part of the host OS, and the library part can be freely chosen
- Related specifications for containers include OCI and CRI

# Gave a Lightning Talk
Gave a lightning talk at Makuake LT Party (an internal LT event).

[speaker-deck - Fully Understanding Containers](https://speakerdeck.com/bmf_san/kontenawan-quan-nili-jie-sita)

# References
- [bmf-tech - What is Docker](https://bmf-tech.com/posts/Docker%E3%81%A8%E3%81%AF)
- [[History of Containers] Until Docker was Created Part 2 ~Let's Gather Collective Knowledge and Learn History~](https://hackmd.io/s/ryPfDLU77)
- [redhat - What is a Linux Container](https://www.redhat.com/ja/topics/containers/whats-a-linux-container)
- [redhat - What is a Linux Container](https://www.redhat.com/ja/topics/containers)
- [Linux Container](https://linuxcontainers.org/ja/)
- [IT Solution School - [Diagram] Understand Container-based Virtualization and Docker with One Sheet](http://blogs.itmedia.co.jp/itsolutionjuku/2015/05/docker.html)
- [SELinux Project Wiki](http://selinuxproject.org/page/Main_Page)
- [opensuse - AppArmor](https://ja.opensuse.org/AppArmor)
- [kernel.org - SECure COMPuting with filters](https://www.kernel.org/doc/Documentation/prctl/seccomp_filter.txt)
- [man7.org - Linux Capabilities](http://man7.org/linux/man-pages/man7/capabilities.7.html)
- [gihyo.jp - Learning Containers with LXC - Technology for Realizing Lightweight Virtualization Environments](https://gihyo.jp/admin/serial/01/linux_containers/0001)
- [Yuuki Blog - The Era of DIY Linux Containers](https://blog.yuuk.io/entry/diy-container)
- [Think IT - Basic Knowledge of Container Technology](https://thinkit.co.jp/story/2015/08/11/6285)
- ~~Linux Containers - What is LXD?~~
- [Hewlett Packard Enterprise - What is the Difference Between Docker Containers and Virtualization? Synergy and DevOps](https://community.hpe.com/t5/Enterprise-Topics/Docker%E3%82%B3%E3%83%B3%E3%83%86%E3%83%8A%E3%81%A8%E4%BB%AE%E6%83%B3%E5%8C%96%E3%81%AE%E9%81%95%E3%81%84%E3%81%A8%E3%81%AF-Synergy%E3%81%A8DevOps/ba-p/6980068?profile.language=ja#.XD6Zks8zZTY)
- [www.publickey1.jp - Mechanism of Container Runtime and Why Firecracker, gVisor, and Unikernel are Attracting Attention. Container Runtime Meetup #2](https://www.publickey1.jp/blog/20/firecrackergvisorunikernel_container_runtime_meetup_2.html)
- [thinkit.co.jp - Prepare for the Deprecation of Docker Runtime Starting from Kubernetes 1.20! What We Should Know and Do](https://thinkit.co.jp/article/18024)
- [container-security.dev - Container Security Books](https://container-security.dev/)
- [github.com - opencontainers/runtime-spec](https://github.com/opencontainers/runtime-spec/blob/main/spec.md)
- [udzura.hatenablog.jp - Reading the OCI Runtime Specification](https://udzura.hatenablog.jp/entry/2016/08/02/155913)
- [medium.com - Overview of the Container Runtime "runc" Used by Every Container User [Container Runtime Meetup #1 Presentation Report]](https://medium.com/nttlabs/runc-overview-263b83164c98)
- [The End of Docker's Dominance: Key Container Trends](https://zenn.dev/ttnt_1013/articles/f36e251a0cd24e)
- [gkuga.hatenablog.com - I Read the OCI Runtime Specification and Wrote an Overview](https://gkuga.hatenablog.com/entry/2020/01/24/032122)
- [yohgami.hateblo.jp - Using chroot and unshare, Create a Simple Container with 7 Commands on Shell](https://yohgami.hateblo.jp/entry/20161215/1481755818)
- ~~Introduction to Container Technology - Learn the Differences with Virtualization and Explore the Elemental Technologies~~
- [Creating a Container Without Using Docker Commands - Part 1](https://zenn.dev/chemimotty/articles/51788231854a5e)
- [kaminashi-developer.hatenablog.jp - [Go Language] DIY Container Swamp. Let's Create a Mini Docker from Scratch](https://kaminashi-developer.hatenablog.jp/entry/dive-into-swamp-container-scratch)
- [www.youtube.com - Building a Container from Scratch in Go - Liz Rice (Microscaling Systems)](https://www.youtube.com/watch?v=Utf-A4rODH8)
- [medium.com - Understand the Design of Container Runtime](https://medium.com/@ikeda.morito/understand-the-design-of-containerruntime-eb79161545ef)
