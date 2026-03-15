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
I embarked on building a Kubernetes environment using Terraform and Ansible.

Driven by the desire to run my custom app on Kubernetes, I started from setting up the environment.

# Server Selection
Since this is a private development, I want to keep the budget around 2000 yen per month.

The major considerations are whether to use cloud or VPS, managed or unmanaged. However, these decisions can be made by weighing the cost and operational benefits, so I don't think it's a big dilemma. As I'll mention later, the biggest headache was the load balancer...

Three candidates emerged this time.

## GCP
- GKE
  - Using it as-is might not lead to cloud bankruptcy, but there's a high chance of exceeding the budget.
  - There's a cheaper way using preemptible VMs, but how viable is it operationally?
    - [ludwig125.hatenablog.com - GKE を格安で使うためにやったこと](https://ludwig125.hatenablog.com/entry/2019/11/30/073458)
    - [sleepless-se.net - GKEで最安値のKubernetesクラスタを作る方法](https://sleepless-se.net/2018/12/11/gke-kubernetes/)
    - [blog.a-know.me - 安価なGKE（k8s）クラスタを作って趣味開発に活用する](https://blog.a-know.me/entry/2018/06/17/220222)

## Digital Ocean
- Managed Kubernetes available on VPS.
- Free master node, pay-as-you-go worker nodes.
- If downtime is acceptable, one worker node costs $10/month ($0.01/hour).
- Load balancer is also pay-as-you-go.
- You can try various things with about $100 from a promo link for new registrations.
- Free and robust monitoring is available.
- Good ecosystem
  - [marketplace.digitaloceancom.com](https://marketplace.digitalocean.com/)
  - [community](https://www.digitalocean.com/community)
- Openstack API compatible
- Might want to use it regardless of using Kubernetes

## Conoha VPS
- No data transfer charges.
- Openstack API compatible
- Easy-to-understand UI.
- Also offers DB servers, object storage, etc.

Besides the above, I considered [civo.com](https://www.civo.com/) which offers managed k3s, but since I wanted to work with k8s, I excluded it from consideration.

I was torn between Digital Ocean and Conoha, but was captivated by the peace of mind of a non-pay-as-you-go pricing model, so I chose Conoha.

GKE and Digital Ocean provide a well-prepared environment for quickly setting up Kubernetes and studying, so I decided to consider using them for such purposes.

# Building a Kubernetes Environment on Conoha VPS
Since I opted not to use managed Kubernetes, I decided to build Kubernetes myself.

I adopted kubeadm as the tool for the build.

Using Terraform and Ansible, I coded everything from instance setup to initial setup (user creation, SSH key adjustments, etc.) and Kubernetes setup with kubeadm:

[github.com - bmf-san/setup-kubernetes-cluster-on-vps-boilerplate](https://github.com/bmf-san/setup-kubernetes-cluster-on-vps-boilerplate)

It assumes one master node and multiple worker nodes.

Conoha provides an API that supports Openstack, so rewriting should be easy if you use other servers that support Openstack (e.g., Digital Ocean).

Building Kubernetes with kubeadm wasn't too difficult as long as you read the official Kubernetes documentation and understand the prerequisites.

# Unresolved Issues
I couldn't handle the load balancer, so I didn't reach the point of publishing the application and operating Kubernetes.

In the case of a self-hosted Kubernetes cluster, you can't use the load balancer provided by the cloud, so you need to prepare an OSS one yourself, but I couldn't set it up successfully and gave up...

I spent nearly a week cutting into my sleep time, but it was beyond my reach...w

This is the unresolved issue:
https://github.com/kubernetes/ingress-nginx/issues/5401

I've decided to temporarily switch to operating my custom app with docker-compose and deepen my understanding of Kubernetes operations before proceeding further...
