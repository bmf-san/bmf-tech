---
title: Kubernetes Components Overview
slug: kubernetes-components
date: 2024-03-27T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Kubernetes
translation_key: kubernetes-components
---

# Kubernetes Components Overview

![Screenshot 2024-03-27 22 40 42](https://github.com/bmf-san/bmf-tech-client/assets/13291041/cf2f9712-2bf0-4c5e-9212-c66c387110b9)

## Control Plane Components
### kube-apiserver
A server that provides the Kubernetes API for operating the cluster. Designed to be horizontally scalable.

### etcd
A highly available key-value store for managing the state of the entire cluster.

### kube-scheduler
Determines which node to schedule newly created Pods that are unassigned.

### kube-controller-manager
A process that controls the state of the cluster, with several types existing.

- Node controller  
  - Notifies when a node is down  
- Job controller  
  - Monitors and creates Jobs (one-off tasks) and runs Pods to complete the tasks  
- EndpointSlice controller  
  - Creates EndpointSlices (references to a set of network endpoints)  
- Service Account controller  
  - Creates ServiceAccounts in new Namespaces  

### cloud-controller-manager
A controller with cloud provider-specific control logic. Used to integrate services provided by the cloud provider into Kubernetes.

## Node Components
### kubelet
An agent that runs on each node. Communicates with the kube-apiserver and manages the execution of Pods.

### kube-proxy
A network proxy in the cluster. Controls communication within and outside the cluster.

### Container Runtime
Software for running containers.

# References
- [kubernetes.io - Kubernetes Components](https://kubernetes.io/docs/concepts/overview/components/)
- [www.redhat.com - Overview of Kubernetes Architecture](https://www.redhat.com/ja/topics/containers/kubernetes-architecture)
- [www.rworks.jp - What is Kubernetes Architecture? Detailed explanation from features and basic components to data protection methods](https://www.rworks.jp/cloud/kubernetes-op-support/kubernetes-column/kubernetes-entry/29132/)
- [speakerdeck.com - Learning the Overall Picture of Kubernetes from Architecture](https://speakerdeck.com/bells17/akitekutiyakaraxue-hukubernetesnoquan-ti-xiang)