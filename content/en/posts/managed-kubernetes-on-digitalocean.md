---
title: Using Managed Kubernetes on DigitalOcean
slug: managed-kubernetes-on-digitalocean
date: 2021-03-07T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Kubernetes
  - VPS
  - DigitalOcean
translation_key: managed-kubernetes-on-digitalocean
---

# Using Managed Kubernetes on DigitalOcean
I wanted to use k8s for learning in private development, and after considering various options, DigitalOcean seemed like a good choice, so I decided to give it a try.

[www.digitalocean.com - The best managed Kubernetes service is the one that’s made for you](https://www.digitalocean.com/blog/best-managed-kubernetes/)

If you're starting fresh, I recommend using promo links or coupons.

I forgot to register through a promo link and missed out on initial credits, but after contacting support, they were kind enough to assist me (I received credits). I'm grateful.

Here's a referral link just in case.

https://m.do.co/c/9fbf85c22695

# Advantages of Digital Ocean
Before discussing managed Kubernetes, let me outline the benefits of Digital Ocean.

There are many features that make you want to use DigitalOcean even without Kubernetes.

- Affordable
  - Droplets (VMs) are billed by the hour based on specifications.
  - There is a free transfer limit based on specifications. If exceeded, it's $0.02/1GB.
    - Conoha has no pay-as-you-go billing for transfer, so you only pay a monthly fee.
    - You cannot check transfer amounts from the dashboard, which is concerning. It seems better to set up monitoring using vmstat.
  - I think the cost performance is excellent.
- Budget alerts
  - Can be set from the management screen. It seems to only notify via email, but you might be able to adjust it for Slack notifications using the API? (I haven't looked into it thoroughly, so this is a guess.)
- Metrics available
  - Basic metrics are well-provided.
  - You can even set alerts. Notifications via Slack or email are possible.
- OpenStack-based API
  - There are providers for Terraform, making IaC easier.
    - [developers.digitalocean.com - documentation](https://developers.digitalocean.com/documentation/)
  - Conoha also has a well-structured API for Japanese VPS.
- Free DNS service
- Rich ecosystem
  - [marketplace.digitalocean.com](https://marketplace.digitalocean.com/)
  - There are many launch templates that can be started with one click. It might be convenient if you want to try out applications.
- Various documents are well-organized and easy to understand
  - This may vary by individual, but it provides a sense of security.
- Good support (personal opinion)
  - I submitted a support ticket regarding the credit issue, and they responded within a day. It might have been a coincidence, but it seems promising.
  - There are some negative opinions about support.
     - [www.websiteplanet.com - digitalocean/#support](www.websiteplanet.com/ja/web-hosting/digitalocean/#support)
- I have developed an attachment to the shark icon. (personal opinion)

# Overview of Digital Ocean's Managed Kubernetes
I will summarize the overview based on [www.digitalocean.com - docs/kubernetes](https://www.digitalocean.com/docs/kubernetes/).

- Master nodes are free.
- Worker nodes are billed at the rate of droplets (Digital Ocean refers to scalable VMs as droplets).
  - [www.digitalocean.com - plans-and-pricing](https://www.digitalocean.com/docs/droplets/#plans-and-pricing)
- The minimum configuration is $10.
  - k8s requires at least 2GB of memory, so you need to choose a worker with 2GB or more for the minimum configuration.
  - With a specification of 1vCPU/2GB Memory, a configuration with one worker costs $10/month.
    - It is recommended to have at least two workers. With only one, there is a possibility of downtime during cluster updates and maintenance.
- It seems that Cilium is adopted for CNI.
- There are no regions in Japan.
  - [www.digitalocean.com - Regional Availability Matrix](https://www.digitalocean.com/docs/platform/availability-matrix/)
- Kubernetes product compliant with CNCF.
  - [github.com - cncf/k8s-conformance](https://github.com/cncf/k8s-conformance)
- Load balancers are also available.
  - [www.digitalocean.com - docs/kubernetes/how-to/add-load-balancers](https://www.digitalocean.com/docs/kubernetes/how-to/add-load-balancers/)
- Log rotation
  - Cluster logs are rotated at 10MB.
  - The active log and the last two are retained.
- Limitations
  - [www.digitalocean.com - docs/kubernetes/#limits](https://www.digitalocean.com/docs/kubernetes/#limits)
  - For personal use, there isn't much to worry about.
- Known issues
  - [www.digitalocean.com - docs/kubernetes/#known-issues](https://www.digitalocean.com/docs/kubernetes/#known-issues)
  - I don't think there are any critical issues, but it might be good to review them.
- There is a cluster auto-upgrade feature.
  - [www.digitalocean.com - docs/kubernetes/how-to/upgrade-cluster](https://www.digitalocean.com/docs/kubernetes/how-to/upgrade-cluster/)

# Trying it Out
Let's actually try the managed Kubernetes service on DigitalOcean.

## Starting a Kubernetes Cluster
I started a Kubernetes cluster with the following configuration.

- Region
  - Singapore
  - I chose the closest location from Japan.
- VPC Network: default-sgp1 DEFAULT
  - This is the default. No room for change.
- Cluster capacity
  - Node pool name
    - test-k8s-node-pool
      - Any name you like.
  - Machine type
    - Basic nodes
  - Node plan
    - 1GB RAM usable (2GB Total)/1vCPU
    - $10/Month per node ($0.015/hr)
  - Number of Nodes
    - 2
- Tags
  - test-k8s
  - You can assign multiple arbitrary tags to the cluster.
- Name
  - test-k8s-cluster
  - You can give any name to the cluster.

The monthly fee is about this.
MONTHLY RATE $20.00/month $0.03/hour

## Installing kubectl and doctl Locally
Refer to the following for kubectl.
[kubernetes.io - install-kubectl](https://kubernetes.io/ja/docs/tasks/tools/install-kubectl/)

```sh
brew install doctl
```

## Preparing Personal Access Tokens
You can check Personal access tokens from the API menu on the dashboard.
Initially, it will only have READ permissions, so grant WRITE permissions and keep the token.
* It seems that the token is not generated initially, so you need to regenerate the token to issue one. Alternatively, creating a new token should also work.

## Connecting to the Cluster
First, authenticate.

```sh
doctl auth init
```

Check the list of clusters.
```sh
doctl kubernetes cluster list
```

Specify the cluster name to connect and add the context (./kube/config will be updated).
```sh
doctl kubernetes cluster kubeconfig save CLUSTER_NAME
```

Check the nodes.
```sh
kubectl get no
```

Now that you're ready to deploy a sample application, it would be good to try something like [github.com - digitalocean/doks-example](https://github.com/digitalocean/doks-example). Be careful as a load balancer will be created, and charges for the load balancer will apply.

# Thoughts
While creating the cheapest k8s environment on GKE is attractive, for personal use, Digital Ocean might be a better choice for managed services.

I might be a bit sensitive, but I'm concerned about the pay-as-you-go transfer, so I will continue to consider Conoha.

# References
- [zenn.dev - Notes from when I started with DigitalOcean Kubernetes](https://zenn.dev/gosarami/articles/94475cc82d73b5e3f453)
- [www.slideshare.net - Starting with DigitalOcean Today](https://www.slideshare.net/zembutsu/all-about-degital-ocean-introduction-distribution-hbstudy)