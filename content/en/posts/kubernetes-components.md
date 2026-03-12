---
title: "Kubernetes Components Explained: Pods, Nodes, and the Control Plane"
description: 'Understand the core Kubernetes components—pods, nodes, API server, scheduler, etcd, and kubelet—and how they work together to run containerized workloads.'
slug: kubernetes-components
date: 2024-03-27T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Kubernetes
translation_key: kubernetes-components
---


# About Kubernetes Components

![Screenshot 2024-03-27 22 40 42](https://github.com/bmf-san/bmf-tech-client/assets/13291041/cf2f9712-2bf0-4c5e-9212-c66c387110b9)

## Control Plane Components
### kube-apiserver
A server that provides the Kubernetes API to operate the cluster. It is designed to scale horizontally.

### etcd
A highly available key-value store for managing the state of all cluster data.

### kube-scheduler
Decides which node an unassigned newly created Pod should be scheduled to.

### kube-controller-manager
A process that controls the state of the cluster, with multiple types available.

- Node controller
  - Notifies when a node is down
- Job controller
  - Monitors and creates Jobs (one-time tasks) and runs Pods to complete tasks
- EndpointSlice controller
  - Creates EndpointSlices (references to collections of network endpoints)
- Service Account controller
  - Creates ServiceAccounts in new Namespaces

### cloud-controller-manager
A controller with cloud provider-specific control logic. It is used to integrate services provided by cloud providers into Kubernetes.

## Node Components
### kubelet
An agent that runs on each node. It communicates with the kube-apiserver and manages the execution of Pods.

### kube-proxy
A network proxy within the cluster. It controls communication within and outside the cluster.

### Container Runtime
Software for running containers.

# References
- [kubernetes.io - Kubernetes Components](https://kubernetes.io/docs/concepts/overview/components/)
- [www.redhat.com - Overview of Kubernetes Architecture](https://www.redhat.com/ja/topics/containers/kubernetes-architecture)
- [www.rworks.jp - What is Kubernetes Architecture? A detailed explanation from basic components to data protection methods](https://www.rworks.jp/cloud/kubernetes-op-support/kubernetes-column/kubernetes-entry/29132/)
- [speakerdeck.com - Learning the Big Picture of Kubernetes from Architecture](https://speakerdeck.com/bells17/akitekutiyakaraxue-hukubernetesnoquan-ti-xiang)
