---
title: Overview of Container Technology
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
Summary of container technology.
Creating and experimenting with containers without using Docker.

# What is a Container
- A series of processes that are separated from the host OS, bundling applications and runtimes.

# History of Containers
In 1979, chroot was introduced in UNIX OS.

In 2000, [FreeBSD jails](https://www.freebsd.org/doc/handbook/jails.html) appeared in FreeBSD 4.0, an evolution of chroot.

In 2001, technologies that would form the basis of Linux containers emerged through the [VServer Project](http://linux-vserver.org/Welcome_to_Linux-VServer.org).

In 2004, LXC 1.0 was released.
[Linux Containers](https://linuxcontainers.org/)

In 2008, Docker was introduced.

Other container technologies include Virtuozzo, OpenVZ, HP-UX Container, Solaris Container, etc.

# Differences Between Containers and Virtualization
- Containers
    - A series of processes that are separated from the host OS, bundling applications and runtimes.
    - Shares the kernel part of the host OS.
        - The library part of the OS can be chosen by the container.

- Virtualization
    - Composed of host-based and hypervisor-based structures, virtualization allows for multiple OS environments, running applications on guest OS.

A rough summary can also be found in [bmf-tech - What is Docker](https://bmf-tech.com/posts/Docker%E3%81%A8%E3%81%AF).

## Linux Kernel Features for Container Technology
### Kernel Namespaces
- A feature that separates processes into six types of resources (ipc, uts, mount, pid, network, user).
- A mechanism that makes it appear as if users have their own isolated resources.
- Isolated resources cannot interfere with each other.

### AppArmor and SELinux Profiles
- AppArmor
    - A type of Linux Security Module (a framework for security in the Linux kernel).
    - Securely manages application access permissions (mandatory access control).
- SELinux (Security Enhanced Linux)
    - A module that adds mandatory access control features to the Linux kernel.

### Seccomp Policies
- A feature that restricts system call issuance by processes.

### Chroots (using pivot_root)
- An operation that changes the root directory for the current process and its child processes.
- Processes with a changed root cannot access files outside the specified range, achieving process isolation.

### Kernel Capabilities
- Management of process permissions.
- Allows for more granular permission management than just root or non-root.

### CGroups (Control Groups)
- A feature that groups processes for common management.

# Docker's Container Technology
Previously, Docker used lxc, but since v0.9, it has been using libcontainer implemented in Go. (cf. [Docker blog - DOCKER 0.9: INTRODUCING EXECUTION DRIVERS AND LIBCONTAINER](https://blog.docker.com/2014/03/docker-0-9-introducing-execution-drivers-and-libcontainer/) [github - opencontainers/runc/libcontainer/](https://github.com/opencontainers/runc/tree/master/libcontainer))

# Standard Specifications
## OCI (Open Container Initiative)
The [Open Container Initiative](https://opencontainers.org/) is an organization aimed at creating industry standards for containers and runtimes.

The following specifications are defined:

- OCI Runtime Specification
- OCI Image Format Specification
- OCI Distribution Specification

OCI is involved with low-level runtime specifications.
ex. runC, gVisor, Kata Containers, Nabla Containers, etc...

## CRI (Container Runtime Interface)
[CRI](https://kubernetes.io/ja/docs/concepts/architecture/cri/) defines the interface for communication between kubelet and container runtimes.

CRI is involved with high-level runtime specifications.
ex. docker, containerd, cri-o

# Conclusion
- Containers are processes with isolated resources.
- Containers share the kernel part of the host OS, while the library part can be freely chosen.
- Specifications related to containers include OCI and CRI.

# LT Presentation
I gave a presentation at the Makuake LT Party (internal LT competition).

[speaker-deck - Fully Understanding Containers](https://speakerdeck.com/bmf_san/kontenawan-quan-nili-jie-sita)

# References
- [bmf-tech - What is Docker](https://bmf-tech.com/posts/Docker%E3%81%A8%E3%81%AF)
- [[History of Containers] Until Docker was Created Part 2 - Let's Gather Collective Knowledge and Learn History](https://hackmd.io/s/ryPfDLU77)
- [redhat - What is a Linux Container](https://www.redhat.com/ja/topics/containers/whats-a-linux-container)
- [redhat - What is a Linux Container](https://www.redhat.com/ja/topics/containers)
- [Linux Container](https://linuxcontainers.org/ja/)
- [IT Solution School - 【Diagram】 Understanding Container-based Virtualization and Docker in One Page](http://blogs.itmedia.co.jp/itsolutionjuku/2015/05/docker.html)
- [SELinux Project Wiki](http://selinuxproject.org/page/Main_Page)
- [opensuse - AppArmor](https://ja.opensuse.org/AppArmor)
- [kernel.org - SECure COMPuting with filters](https://www.kernel.org/doc/Documentation/prctl/seccomp_filter.txt)
- [man7.org - Linux Capabilities](http://man7.org/linux/man-pages/man7/capabilities.7.html)
- [gihyo.jp - Introduction to Containers with LXC - Technology for Lightweight Virtualization Environments](https://gihyo.jp/admin/serial/01/linux_containers/0001)
- [Yuuki Blog - The Era of DIY Linux Containers](https://blog.yuuk.io/entry/diy-container)
- [Think IT - Basic Knowledge of Container Technology](https://thinkit.co.jp/story/2015/08/11/6285)
- [Linux Containers - What is LXD?](https://linuxcontainers.org/ja/lxd/introduction/)
- [Hewlett Packard Enterprise - Differences Between Docker Containers and Virtualization? Synergy and DevOps](https://community.hpe.com/t5/Enterprise-Topics/Docker%E3%82%B3%E3%83%B3%E3%83%86%E3%83%8A%E3%81%A8%E4%BB%AE%E6%83%B3%E5%8C%96%E3%81%AE%E9%81%95%E3%81%84%E3%81%A8%E3%81%AF-Synergy%E3%81%A8DevOps/ba-p/6980068?profile.language=ja#.XD6Zks8zZTY)
- [www.publickey1.jp - The Mechanism of Container Runtime and Why Firecracker, gVisor, and Unikernel are Gaining Attention. Container Runtime Meetup #2](https://www.publickey1.jp/blog/20/firecrackergvisorunikernel_container_runtime_meetup_2.html)
- [thinkit.co.jp - Preparing for the Deprecation of Docker Runtime Starting from Kubernetes 1.20! What We Should Know and Do](https://thinkit.co.jp/article/18024)
- [container-security.dev - Container Security Books](https://container-security.dev/)
- [github.com - opencontainers/runtime-spec](https://github.com/opencontainers/runtime-spec/blob/main/spec.md)
- [udzura.hatenablog.jp - Reading the OCI Runtime Specification](https://udzura.hatenablog.jp/entry/2016/08/02/155913)
- [medium.com - An Overview of the Runtime 'runc' Used by All Container Users [Report on Container Runtime Meetup #1]](https://medium.com/nttlabs/runc-overview-263b83164c98)
- [As Docker's Dominance Ends, Key Container Issues to Understand](https://zenn.dev/ttnt_1013/articles/f36e251a0cd24e)
- [gkuga.hatenablog.com - Writing an Overview After Reading the OCI Runtime Specification](https://gkuga.hatenablog.com/entry/2020/01/24/032122)
- [yohgami.hateblo.jp - Using chroot and unshare, Creating a Simple Container with 7 Commands in Shell](https://yohgami.hateblo.jp/entry/20161215/1481755818)
- [Introduction to Container Technology - Understanding the Differences with Virtualization and Learning the Elemental Technologies](https://eh-career.com/engineerhub/entry/2019/02/05/103000)
- [Creating a Container Without Using Docker - 1](https://zenn.dev/chemimotty/articles/51788231854a5e)
- [kaminashi-developer.hatenablog.jp - [Go Language] DIY Container Swamp. Let's Create a Mini Docker from Scratch](https://kaminashi-developer.hatenablog.jp/entry/dive-into-swamp-container-scratch)
- [www.youtube.com - Building a Container from Scratch in Go - Liz Rice (Microscaling Systems)](https://www.youtube.com/watch?v=Utf-A4rODH8)
- [medium.com - Understand the Design of Container Runtime](https://medium.com/@ikeda.morito/understand-the-design-of-containerruntime-eb79161545ef)