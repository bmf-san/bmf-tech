---
title: Building a Kubernetes Environment with Terraform and Ansible
slug: kubernetes-setup-with-terraform-ansible
date: 2021-04-06T00:00:00Z
lastmod: 2026-03-15
author: bmf-san
categories:
  - Infrastructure
tags:
  - Ansible
  - Kubernetes
  - Terraform
  - kubeadm
description: A journey of setting up a Kubernetes environment using Terraform and Ansible.
translation_key: kubernetes-setup-with-terraform-ansible
---

# Overview
I worked on building a Kubernetes environment using Terraform and Ansible.

I started building the environment from the desire to run my own application on Kubernetes.

# Server Selection
Since this is for private development, I want to keep the budget around 2000 yen per month.

The major considerations are whether to use cloud or VPS, managed or unmanaged services, but I think I won't have too much trouble deciding that while considering cost and operational benefits. As mentioned later, the biggest headache was the load balancer...

Three options came up this time.

## GCP
- GKE
  - Using it as is could lead to cloud bankruptcy... but it is likely to exceed the budget.
  - There is a method to use preemptible VMs to keep costs low, but I wonder how feasible it is operationally.
    - [ludwig125.hatenablog.com - What I Did to Use GKE Cheaply](https://ludwig125.hatenablog.com/entry/2019/11/30/073458)
    - [sleepless-se.net - How to Create the Cheapest Kubernetes Cluster on GKE](https://sleepless-se.net/2018/12/11/gke-kubernetes/)
    - [blog.a-know.me - Creating an Inexpensive GKE (k8s) Cluster for Hobby Development](https://blog.a-know.me/entry/2018/06/17/220222)

## Digital Ocean
- A VPS that offers managed Kubernetes.
- The master node is free, and worker nodes are billed based on usage.
- If downtime is acceptable, one worker node costs $10 per month ($0.01/hour).
- Load balancer is also billed based on usage.
- New registrations through promotional links receive about $100, allowing for various trials.
- The free monitoring service is quite robust and good.
- The ecosystem is nice
  - [marketplace.digitalocean.com](https://marketplace.digitalocean.com/)
  - [community](https://www.digitalocean.com/community)
- OpenStack API compatible
- Might want to use it regardless of whether I use Kubernetes

## Conoha VPS
- No charges for data transfer.
- OpenStack API compatible
- The UI is user-friendly.
- Also offers DB servers, object storage, etc.

In addition to the above, I considered the option of managed k3s provided by [civo.com](https://www.civo.com/), but since I wanted to work with k8s, I decided against it.

I was torn between Digital Ocean and Conoha, but I was captivated by the reassuring pricing structure with no usage charges, so I chose Conoha.

I believe GKE and Digital Ocean provide a suitable environment for quickly setting up Kubernetes for study purposes, so I made the decision to consider using them for that purpose.

# Building a Kubernetes Environment on Conoha VPS
Since I chose not to use managed Kubernetes, I decided to build Kubernetes myself.

I adopted kubeadm as the tool for the setup.

Using Terraform and Ansible, I coded everything from instance creation to initial setup (user creation, SSH key adjustments, etc.) and the construction of Kubernetes using kubeadm, which can be found here:

[github.com - bmf-san/setup-kubernetes-cluster-on-vps-boilerplate](https://github.com/bmf-san/setup-kubernetes-cluster-on-vps-boilerplate)

It is designed for one master node and multiple worker nodes.

Conoha provides an API that supports OpenStack, so it should be easy to modify for other servers that support OpenStack (e.g., Digital Ocean).

Building Kubernetes with kubeadm was not too difficult as long as I read the official Kubernetes documentation to understand the prerequisites.

# Issues That Could Not Be Resolved
I couldn't address the load balancer issue, so I was unable to publish the application and operate it on Kubernetes.

In the case of a self-hosted Kubernetes cluster, I cannot use the load balancer provided by the cloud, so I need to prepare an open-source solution myself, but I was unable to successfully set that up and had to give up...

I spent nearly a week sacrificing sleep but couldn't make any progress... lol

The unresolved issue is this one.
https://github.com/kubernetes/ingress-nginx/issues/5401

I decided to temporarily switch to running my application with docker-compose and plan to deepen my understanding of Kubernetes operations before trying again...
