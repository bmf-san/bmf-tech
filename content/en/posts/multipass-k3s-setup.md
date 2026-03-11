---
title: Running k3s with multipass
slug: multipass-k3s-setup
date: 2023-08-17T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - multipass
  - k3s
description: Exploring the use of multipass to run k3s for personal development applications.
translation_key: multipass-k3s-setup
---



# Overview
This is a note on using multipass during an attempt to migrate a personal development application, structured with docker-compose, to Kubernetes (k3s).

Ultimately, the migration was not completed...

[k3s](https://k3s.io/) is a CNCF-certified Kubernetes distribution. It's designed for IoT and edge computing use cases. It's useful when you want to save memory, don't need the scale of full Kubernetes, or just want to try Kubernetes easily. It can also be an option for individuals wanting to implement Kubernetes on a VPS.

cf. [K3s on ConoHa](https://qiita.com/yhirokw/items/fd5dcb28d3f57de0cc40)

Functionally, k3s is almost the same as Kubernetes, but it has some limitations. For details, refer to the documentation.

cf. [docs.k3s.io](https://docs.k3s.io/)

# What is multipass?
A tool that allows you to easily create Ubuntu virtual environments. It supports Linux, macOS, and Windows.

[multipass.run](https://multipass.run/)

# Why use multipass?
I needed to set up a virtual environment on macOS as a runtime for k3s.

cf. [Can I install k3s on macos (big sur) with m1 chip?](https://www.reddit.com/r/kubernetes/comments/qa2f8d/can_i_install_k3s_on_macos_big_sur_with_m1_chip/)

There are several alternatives, but I decided to try multipass because it seemed easy and accessible.

# Running k3s with multipass
On macOS, you can install multipass with brew and run k3s with just the following steps:

1. `mutipass find` // Search for Ubuntu images
2. `multipass launch -c 2 -m 4G -d 50G -n example 22.10` // 22.10 is the Ubuntu version
3. `multipass mount ./k3s/ example:~/k3s` // Mount
4. `multipass shell example` // Connect to the virtual machine
5. `curl -sfL https://get.k3s.io | sh -` // Install k8s on the virtual machine

cf.
- [Installing Multipass on M1 Mac for Ubuntu Virtual Machines](https://virment.com/how-to-install-multipass-to-m1-mac/#%E3%83%9B%E3%82%B9%E3%83%88%E3%83%9E%E3%82%B7%E3%83%B3%E3%81%AE%E3%83%87%E3%82%A3%E3%83%AC%E3%82%AF%E3%83%88%E3%83%AA%E3%82%92%E3%83%9E%E3%82%A6%E3%83%B3%E3%83%88%E3%81%99%E3%82%8B)
- [k3s on M1 Mac](https://qiita.com/tkuribayashi/items/4eb664631254fa58df57)
- [Building a Kubernetes Environment on M1 Mac without Docker Desktop using Multipass](https://zenn.dev/kkoudev/articles/b001c36c7d7005)
