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
I wanted to use k8s for personal development and learning, and after considering various options, DigitalOcean seemed promising, so I decided to try it out.

[www.digitalocean.com - The best managed Kubernetes service is the one that’s made for you](https://www.digitalocean.com/blog/best-managed-kubernetes/)

If you're starting fresh, it's a good idea to use promo links or coupons.

I forgot to register through a promo link initially and missed out on credits, but after contacting support, they kindly provided the credits. Much appreciated.

Here's a referral link just in case.

https://m.do.co/c/9fbf85c22695

# Good Points of Digital Ocean
Before talking about managed Kubernetes, let's discuss the good points of Digital Ocean.

Even without using Kubernetes, DigitalOcean offers features that make you want to use it.

- Affordable
  - Droplets (VMs) are billed hourly based on specs.
  - Transfer volume has a free tier per spec. Exceeding it costs $0.02/1GB.
    - With Conoha, there's no pay-as-you-go for transfer volume, just a monthly fee.
    - You can't check transfer volume from the dashboard, which is concerning. It might be good to set up monitoring with vmstat.
  - I think the cost performance is excellent.
- Budget alerts
  - Can be set from the management screen. It seems to be email notifications only, but you might be able to adjust it for Slack notifications using the API? (Just a guess as I haven't checked thoroughly)
- Metrics available
  - Basic metrics are well-equipped.
  - Alerts can be set up, allowing notifications via Slack or email.
- OpenStack-based API
  - Providers like Terraform make IaC easy.
    - [developers.digitalocean.com - documentation](https://developers.digitalocean.com/documentation/)
  - Conoha in Japan also has a well-developed API.
- Free DNS service
- Rich ecosystem
  - [marketplace.digitalocean.com](https://marketplace.digitalocean.com/)
  - There are plenty of launch templates that can be started with 1-click. Convenient if you want to try out applications.
- Well-organized and easy-to-understand documentation
  - This might vary by person, but it provides a sense of security.
- Good support (personal opinion)
  - I submitted a support ticket regarding the credits issue mentioned earlier, and they responded within a day. It might have been a coincidence, but it seems promising.
  - There are some negative reviews about support.
     - [www.websiteplanet.com - digitalocean/#support](www.websiteplanet.com/ja/web-hosting/digitalocean/#support)
- The shark icon is endearing. (personal opinion)

# Overview of Digital Ocean's Managed Kubernetes
Referencing [www.digitalocean.com - docs/kubernetes](https://www.digitalocean.com/docs/kubernetes/), here's a summary.

- Master nodes are free
- Worker nodes are billed at the rate of droplets (Digital Ocean's scalable VMs)
  - [www.digitalocean.com - plans-and-pricing](https://www.digitalocean.com/docs/droplets/#plans-and-pricing)
- Minimum configuration is $10
  - k8s requires more than 2GB of memory, so for the minimum configuration, choose a worker with more than 2GB.
  - With a 1vCPU/2GB Memory spec, a single worker setup costs $10/Month
    - More than one worker is recommended. With just one, there's a risk of downtime during cluster updates or maintenance.
- CNI seems to use Cilium
- No region in Japan
  - [www.digitalocean.com - Regional Availability Matrix](https://www.digitalocean.com/docs/platform/availability-matrix/)
- Kubernetes product compliant with CNCF
  - [github.com - cncf/k8s-conformance](https://github.com/cncf/k8s-conformance)
- Load balancer available
  - [www.digitalocean.com - docs/kubernetes/how-to/add-load-balancers](https://www.digitalocean.com/docs/kubernetes/how-to/add-load-balancers/)
- Log rotation
  - Cluster logs rotate at 10MB.
  - Active logs and the last two are retained.
- Limitations
  - [www.digitalocean.com - docs/kubernetes/#limits](https://www.digitalocean.com/docs/kubernetes/#limits)
  - Not much of a concern for personal use
- Known issues
  - [www.digitalocean.com - docs/kubernetes/#known-issues](https://www.digitalocean.com/docs/kubernetes/#known-issues)
  - No critical issues, but it's good to review them
- Automatic cluster upgrade feature
  - [www.digitalocean.com - docs/kubernetes/how-to/upgrade-cluster](https://www.digitalocean.com/docs/kubernetes/how-to/upgrade-cluster/)


# Trying It Out
Actually trying out DigitalOcean's managed Kubernetes service.

## Starting a Kubernetes Cluster
I started a Kubernetes cluster with the following configuration.

- Region
  - Singapore
  - Chose the closest location to Japan
- VPC Network: default-sgp1 DEFAULT
  - This is default. No option to change.
- Cluster capacity
  - Node pool name
    - test-k8s-node-pool
      - Arbitrary name
  - Machine type
    - Basic nodes
  - Node plan
    - 1GB RAM usable(2GB Total)/1vCPU
    - $10/Month per node($0.015/hr)
  - Number Nodes
    - 2
- Tags
  - test-k8s
  - You can assign multiple arbitrary tags to the cluster.
- Name
  - test-k8s-cluster
  - You can assign an arbitrary cluster name.

Monthly cost is about this.
MONTHLY RATE $20.00/month $0.03/hour

## Installing kubectl and doctl Locally
Refer to the following for kubectl.
[kubernetes.io - install-kubectl](https://kubernetes.io/ja/docs/tasks/tools/install-kubectl/)

```sh
brew install doctl
```

## Preparing Personal Access Tokens
You can check Personal access tokens from the API menu on the dashboard.
Initially, it's READ only, so grant WRITE as well and keep the token.
※ Initially, it seems no token is generated, so you need to Regenerate token to issue it. Alternatively, creating a new token is also fine.

## Connecting to the Cluster
First, authenticate.

```sh
doctl auth init
```

Check the list of clusters.
```sh
doctl kubernetes cluster list
```

Specify the cluster name to connect to and add the context. (./kube/config will be updated)
```sh
doctl kubernetes cluster kubeconfig save CLUSTER_NAME
```

Check the nodes.
```sh
kubectl get no
```

With this, you're ready to deploy a sample app, so trying something like [github.com - digitalocean/doks-example](https://github.com/digitalocean/doks-example) might be good. Note that creating a load balancer will incur charges.

# Impressions
Creating the cheapest k8s environment with GKE is also attractive, but for personal use, Digital Ocean might be better.

I might be a bit sensitive, but the pay-as-you-go transfer volume is concerning, so I'll continue considering Conoha.

# References
- [zenn.dev - DigitalOcean Kubernetesに入門した時のメモ](https://zenn.dev/gosarami/articles/94475cc82d73b5e3f453)
- [www.slideshare.net - 今日から始めるDigitalOcean](https://www.slideshare.net/zembutsu/all-about-degital-ocean-introduction-distribution-hbstudy)
