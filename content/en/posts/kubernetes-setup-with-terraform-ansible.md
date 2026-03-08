---
title: Building a Kubernetes Environment with Terraform and Ansible
slug: kubernetes-setup-with-terraform-ansible
date: 2021-04-06T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Ansible
  - Kubernetes
  - Terraform
  - kubeadm
translation_key: kubernetes-setup-with-terraform-ansible
---

# Overview
I worked on building a Kubernetes environment using Terraform and Ansible.

I started from the desire to run my custom application on Kubernetes.

# Server Selection
Since this is for private development, I want to keep the budget around 2000 yen per month at most.

The major considerations are whether to use cloud or VPS, managed or unmanaged, but I don't think it will be too difficult to decide as long as I consider the cost and operational benefits. As mentioned later, the biggest headache was the load balancer...

Three candidates came up this time.

## GCP
- GKE
  - Using it as is could lead to cloud bankruptcy... but there is a high possibility of exceeding the budget.
  - There is a way to use it cheaply by utilizing preemptible VMs, but I wonder how it would be operationally.
    - [ludwig125.hatenablog.com - What I did to use GKE cheaply](https://ludwig125.hatenablog.com/entry/2019/11/30/073458)
    - [sleepless-se.net - How to create the cheapest Kubernetes cluster on GKE](https://sleepless-se.net/2018/12/11/gke-kubernetes/)
    - [blog.a-know.me - Creating a cheap GKE (k8s) cluster for hobby development](https://blog.a-know.me/entry/2018/06/17/220222)

## Digital Ocean
- A VPS that allows managed Kubernetes.
- The master node is free, and worker nodes are pay-as-you-go.
- If downtime is acceptable, one worker costs $10 per month ($0.01/hour).
- Load balancer is also pay-as-you-go.
- You can get about $100 for signing up through a promo link, allowing you to try various things.
- The free monitoring is quite solid and good.
- The ecosystem is good
  - [marketplace.digitalocean.com](https://marketplace.digitalocean.com/)
  - [community](https://www.digitalocean.com/community)
- Openstack API compatible
- Might want to use it regardless of whether I use Kubernetes or not.

## Conoha VPS
- No data transfer fees.
- Openstack API compatible.
- The UI is easy to understand.
- There are also DB servers and object storage.

In addition to the above, I considered the option of managed k3s provided by [civo.com](https://www.civo.com/), but since I wanted to touch k8s, I excluded it from consideration.

I was torn between Digital Ocean and Conoha, but I chose Conoha because I was captivated by the reassuring pricing structure with no pay-as-you-go.

I think GKE and Digital Ocean provide a good environment for quickly building and studying Kubernetes, so I decided to consider using them for that purpose.

# Building a Kubernetes Environment on Conoha VPS
Since I chose not to use managed Kubernetes, I decided to build Kubernetes myself.

I adopted kubeadm as the tool for building.

Using Terraform and Ansible, I coded everything from instance creation to initial setup (user creation, SSH key adjustments, etc.) and the construction of Kubernetes using kubeadm, which can be found here:

[github.com - bmf-san/setup-kubernetes-cluster-on-vps-boilerplate](https://github.com/bmf-san/setup-kubernetes-cluster-on-vps-boilerplate)

It is assumed to have one master node and multiple worker nodes.

Conoha provides an API that supports Openstack, so it should be easy to rewrite for other servers that support Openstack (e.g., Digital Ocean).

Building Kubernetes with kubeadm wasn't too difficult as long as I read the official Kubernetes documentation to understand the prerequisites.

# Issues I Couldn't Resolve
I couldn't handle the load balancer, so I couldn't reach the point of publishing the application and operating Kubernetes.

In the case of a self-hosted Kubernetes cluster, I cannot use the load balancer provided by the cloud, so I need to prepare an OSS one myself, but I couldn't set it up successfully and gave up...

I lost almost a week of sleep but couldn't make any progress..w

The unresolved issue is this.
https://github.com/kubernetes/ingress-nginx/issues/5401

I decided to temporarily shift to operating my custom application with docker-compose and plan to delve deeper into Kubernetes operations before proceeding further...