---
title: Trying Managed Kubernetes on DigitalOcean
slug: managed-kubernetes-on-digitalocean
date: 2021-03-07T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Kubernetes
  - VPS
  - DigitalOcean
description: Exploring the use of managed Kubernetes on DigitalOcean for personal development and learning.
translation_key: managed-kubernetes-on-digitalocean
---

# Trying Managed Kubernetes on DigitalOcean
I wanted to use k8s for personal development and learning, and after considering various options, I found that DigitalOcean seemed like a good choice, so I decided to give it a try.

[www.digitalocean.com - The best managed Kubernetes service is the one that’s made for you](https://www.digitalocean.com/blog/best-managed-kubernetes/)

If you're starting fresh, I think it's a good idea to use promotional links or coupons.

I forgot to register through the promotional link and missed out on getting credits initially, but when I inquired, they were kind enough to assist me (I received credits). I'm grateful for that.

I'll leave a referral link here just in case.

https://m.do.co/c/9fbf85c22695

# Good Points About DigitalOcean
Before discussing managed Kubernetes, I'll outline some of the good points about DigitalOcean.

There are many features that make you want to use DigitalOcean even without using Kubernetes.

- Affordable
  - Droplets (VMs) are billed by the hour based on specifications.
  - There is a free transfer limit based on specifications. If exceeded, it's $0.02/1GB.
    - With Conoha, there are no pay-as-you-go charges for transfer, so you only pay the monthly fee.
    - You cannot check transfer amounts from the dashboard. This is a concern. It seems better to set up a monitoring system using vmstat.
  - I think the cost performance is excellent.
- Budget alerts are available
  - Can be set from the management screen. It seems like only email notifications are available, but you might be able to adjust it for Slack notifications using the API? (I'm just guessing since I haven't looked into it thoroughly.)
- Metrics can be collected
  - Basic metrics are well-equipped.
  - Moreover, alerts can be set. Slack and email notifications are possible.
- There is an OpenStack-based API
  - There are providers for Terraform, making Infrastructure as Code (IaC) easier.
    - [developers.digitalocean.com - documentation](https://developers.digitalocean.com/documentation/)
  - Conoha also has a well-structured API for VPS in Japan.
- Free DNS service
- A rich ecosystem
  - [marketplace.digitalocean.com](https://marketplace.digitalocean.com/)
  - There are quite a few launch templates that can be started with one click. It could be convenient if you want to try out applications.
- Various documents are well-organized and easy to understand
  - There may be personal differences, but it provides a sense of security.
- Support is quite good (personal opinion)
  - I created a support ticket regarding the credit issue at the beginning, and they responded within a day. It might have been a coincidence, but it seems promising.
  - There are some negative reviews about support.
     - [www.websiteplanet.com - digitalocean/#support](https://www.websiteplanet.com/ja/web-hosting/digitalocean/#support)
- I have grown fond of the shark icon. (personal opinion)

# Overview of DigitalOcean's Managed Kubernetes
I will summarize the overview based on [www.digitalocean.com - docs/kubernetes](https://www.digitalocean.com/docs/kubernetes/).

- Master nodes are free.
- Worker nodes are charged at the droplet rate (DigitalOcean refers to scalable VMs as droplets).
  - [www.digitalocean.com - plans-and-pricing](https://www.digitalocean.com/docs/droplets/#plans-and-pricing)
- The minimum configuration is $10.
  - Since k8s requires at least 2GB of memory, you need to choose a worker with more than 2GB for the minimum configuration.
  - With a specification of 1vCPU/2GB Memory, a configuration with one worker costs $10/month.
    - It is recommended to have at least two workers. With only one, there is a possibility of downtime during cluster updates and maintenance.
- It seems that Cilium is adopted for CNI.
- There is no region in Japan.
  - [www.digitalocean.com - Regional Availability Matrix](https://www.digitalocean.com/docs/platform/availability-matrix/)
- A Kubernetes product compliant with CNCF.
  - [github.com - cncf/k8s-conformance](https://github.com/cncf/k8s-conformance)
- Load balancers are also provided.
  - [www.digitalocean.com - docs/kubernetes/how-to/add-load-balancers](https://www.digitalocean.com/docs/kubernetes/how-to/add-load-balancers/)
- Log rotation
  - Cluster logs are rotated at 10MB.
  - Active logs and the last two are retained.
- Limitations
  - [www.digitalocean.com - docs/kubernetes/#limits](https://www.digitalocean.com/docs/kubernetes/#limits)
  - There isn't much that would be a concern for personal use.
- Known issues
  - [www.digitalocean.com - docs/kubernetes/#known-issues](https://www.digitalocean.com/docs/kubernetes/#known-issues)
  - I don't think there are any critical issues, but it's better to review them.
- There is an automatic cluster upgrade feature.
  - [www.digitalocean.com - docs/kubernetes/how-to/upgrade-cluster](https://www.digitalocean.com/docs/kubernetes/how-to/upgrade-cluster/)

# Trying It Out
Now, I will actually try out DigitalOcean's managed Kubernetes service.

## Starting a Kubernetes Cluster
I started a Kubernetes cluster with the following configuration.

- Region
  - Singapore
  - I chose the closest location from Japan.
- VPC Network: default-sgp1 DEFAULT
  - This is the default. There is no room for change.
- Cluster capacity
  - Node pool name
    - test-k8s-node-pool
      - Any name you like.
  - Machine type
    - Basic nodes
  - Node plan
    - 1GB RAM usable (2GB Total)/1vCPU
    - $10/month per node ($0.015/hr)
  - Number of Nodes
    - 2
- Tags
  - test-k8s
  - You can assign multiple arbitrary tags to the cluster.
- Name
  - test-k8s-cluster
  - You can assign any name to the cluster.

The monthly fee is about this much.
MONTHLY RATE $20.00/month $0.03/hour

## Installing kubectl and doctl on Local Environment
Refer to the following for kubectl.
[kubernetes.io - install-kubectl](https://kubernetes.io/ja/docs/tasks/tools/install-kubectl/)

```sh
brew install doctl
```

## Preparing Personal Access Tokens
You can check Personal access tokens from the API menu on the dashboard.
Initially, it will only have READ permissions, so grant WRITE permissions as well and keep the token safe.
*It seems that the token is not generated on the first time, so you need to regenerate the token to issue it. Alternatively, creating a new token should also be fine.

## Connecting to the Cluster
First, authenticate.

```sh
doctl auth init
```

Check the list of clusters.
```sh
doctl kubernetes cluster list
```

Specify the cluster name to connect to and add the context (./kube/config will be updated).
```sh
doctl kubernetes cluster kubeconfig save CLUSTER_NAME
```

Check the nodes.
```sh
kubectl get no
```

Now that I am ready to deploy a sample application, it would be good to try something like [github.com - digitalocean/doks-example](https://github.com/digitalocean/doks-example). Be careful as a load balancer will be created, which will incur charges.

# Thoughts
Creating a k8s environment with the lowest configuration on GKE is also attractive, but for personal use, DigitalOcean might be a better choice for managed services.

I might be a bit sensitive, but I'm concerned about the pay-as-you-go transfer charges, so I will continue to consider Conoha.

# References
- [zenn.dev - Notes from when I started with DigitalOcean Kubernetes](https://zenn.dev/gosarami/articles/94475cc82d73b5e3f453)
- [www.slideshare.net - Starting with DigitalOcean Today](https://www.slideshare.net/zembutsu/all-about-degital-ocean-introduction-distribution-hbstudy)
