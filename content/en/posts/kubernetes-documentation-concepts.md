---
title: Reading Kubernetes Documentation - Summary of Concepts
slug: kubernetes-documentation-concepts
date: 2020-10-20T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Docker
  - Container
  - Kubernetes
translation_key: kubernetes-documentation-concepts
---



# Overview
To seriously catch up with Kubernetes, I read the documentation and left my personal notes. Since it's long, I only took notes on the concepts section.

[kubernetes.io](https://kubernetes.io/ja/docs/home/)

# What is Kubernetes?
cf. [What is Kubernetes?](https://kubernetes.io/ja/docs/concepts/overview/what-is-kubernetes/)

## What is Kubernetes?
- Declarative configuration management
- Promotion of automation
- A platform for managing containerized workloads and services

## Looking Back
- Deployment before virtualization (Traditional deployment)
  - No resource limitations for applications on physical servers
    - Resource allocation issues
  - Difficult to scale
  - Maintenance costs
- Deployment using virtualization (Virtualized deployment)
  - Applications can be isolated per VM
    - Restriction on data access between applications
  - Improved resource utilization within physical servers through virtualization
  - Easy to add or update applications
    - Reduced hardware costs, improved scalability
- Deployment using containers (Container deployment)
  - OS can be shared between applications
    - Lightweight
  - Containers have their own filesystem, CPU, memory, process space, etc.
  - Independent of cloud or OS distribution
  - Benefits of containers
    - Easier and more efficient to create container images than VM images
    - Continuous build and deployment of container images
    - Separation of concerns between development and operations
      - Application container images are created during build and release
    - High observability
      - In addition to OS-level information and metrics, also includes application status and other alerts
    - Consistency across environments
      - Can be run the same way in development, testing, and production
    - Portability across cloud and OS distributions
      - Can be run in any environment, whether on-premises or public cloud
    - Application-centric management
      - From running OS on virtual machines to running applications using logical resources on OS
    - High affinity with microservices
      - Compatible with loosely coupled, distributed, scalable, and flexible microservices
    - Resource partitioning
      - Predictable application performance
    - Efficient use and aggregation of resources

## Why Kubernetes is Needed and Its Features
- Service discovery and load balancing
  - Containers can be exposed using DNS names or IP addresses
  - Network traffic can be distributed
- Storage orchestration
  - Freedom to choose storage to mount
- Automated rollouts and rollbacks
  - Can define the state of containers to be deployed
- Automatic bin packing
  - Can declare CPU and memory (RAM) required by containers
  - Can adjust according to nodes, efficiently utilizing resources
- Self-healing
  - Can restart, replace, or terminate containers that fail to start
- Secrets and configuration management
  - Can update application configuration without recreating container images

## What Kubernetes Does Not Include
- Kubernetes does not...
  - Limit the types of applications it supports
  - Deploy source code or build applications
  - Provide built-in application-level (middleware, databases, caches, etc.) features
  - Specify logging, monitoring, or alerting features
  - Provide a configuration language
  - Provide or adopt systems for machine configuration, maintenance, management, or self-healing
  - Assume orchestration

# Kubernetes Components

cf. [Kubernetes Components](https://kubernetes.io/ja/docs/concepts/overview/components/)
- Deploying Kubernetes results in a cluster
  - A cluster is a set of nodes that run containerized applications
    - Every cluster has at least one worker node
  - Worker nodes host Pods, which are components of applications
  - Master nodes manage worker nodes and Pods within the cluster
    - Using multiple master nodes provides failover and high availability to the cluster
  - The control plane manages worker nodes and Pods within the cluster
    - In production environments, multiple nodes can be used to provide fault tolerance and high availability
  - [Diagram of a Kubernetes Cluster](https://d33wubrfki0l68.cloudfront.net/7016517375d10c702489167e704dcb99e570df85/7bb53/images/docs/components-of-kubernetes.png)

## Control Plane Components
- Makes overall decisions about the cluster (e.g., scheduling)

### kube-apiserver
- Component that exposes the Kubernetes API externally
- Designed to scale horizontally

### etcd
- Consistent, highly available key-value store
- Stores all cluster information for Kubernetes
- When using etcd as a data store for Kubernetes, always create a backup plan

### kube-scheduler
- Monitors newly created Pods that have no node assigned, and selects a node for them to run on

### kube-controller-manager
- Runs multiple controller processes
- Operates as a single process
  - Logically, each controller is compiled into a single executable file
- Includes the following controllers:
  - Node Controller
    - Notifies and responds when nodes go down
  - Replication Controller
    - Maintains the correct number of Pods for all replication controller objects
  - Endpoint Controller
    - Links Services and Pods
  - Service Account and Token Controller
    - Creates default accounts and API access tokens for new namespaces

### cloud-controller-manager
- Runs controllers that interact with the underlying cloud provider
- The following controllers have dependencies on the cloud provider:
  - Node Controller
    - Checks with the cloud provider to determine if a node has been deleted after it stops responding
  - Routing Controller
    - Sets up routing in the underlying cloud infrastructure
  - Service Controller
    - Creates, updates, and deletes cloud provider load balancers
  - Volume Controller
    - Creates, attaches, mounts volumes, and coordinates with the cloud provider

## Node Components
- Provides the runtime environment for managing Pods on all nodes

### kubelet
- Agent that runs on each node in the cluster
- Ensures that containers are running in a Pod

### kube-proxy
- Network proxy that runs on each node in the cluster, implementing part of the Kubernetes Service concept

### Container Runtime
- Software responsible for running containers
- ex. Docker, containerd, CRI-O etc...

## Add-ons
- Uses Kubernetes resources (DaemonSet, Deployment, etc.) to implement cluster features
- Provides cluster-level features
  - Add-ons that require namespaces belong to the kube-system namespace
- Some add-ons include:
  - DNS
  - Web UI
  - Container Resource Monitoring
  - Cluster-level Logging

# Kubernetes API

cf. [Kubernetes API](https://kubernetes.io/ja/docs/concepts/overview/kubernetes-api/)

- Refer to [API Reference](https://kubernetes.io/docs/reference/)

# About Kubernetes Objects

cf. [About Kubernetes Objects](https://kubernetes.io/ja/docs/concepts/overview/working-with-objects/)

## Object spec and status
- Kubernetes objects have two nested object fields that manage the object's configuration
  - spec
    - Describes the desired state and characteristics of the object
  - status
    - Indicates the current state of the object

# Kubernetes Object Management

cf. [Kubernetes Object Management](https://kubernetes.io/ja/docs/concepts/overview/working-with-objects/object-management/)

## Management Methods
- Imperative Commands
    - Targets existing objects
    - Recommended for development project environments
- Imperative Object Configuration
    - Targets individual files
    - Recommended for production project environments
- Declarative Object Configuration
    - Targets directories of files
    - Recommended for production project environments

## Imperative Commands
- Users perform operations on existing objects in the cluster

## Imperative Object Configuration
- Specify the operation, optional flags, and one or more filenames with the kubectl command

## Declarative Object Configuration
- Users operate on configuration files located locally
- Operations are not recorded in the files

# Object Names and IDs

cf. [Object Names and IDs](https://kubernetes.io/ja/docs/concepts/overview/working-with-objects/names/)

# Namespace

cf. [Namespace](https://kubernetes.io/ja/docs/concepts/overview/working-with-objects/namespaces/)

- Supports running multiple virtual clusters on the same physical cluster
  - These virtual clusters are called Namespaces

# Nodes

cf. [Nodes](https://kubernetes.io/ja/docs/concepts/architecture/nodes/)

- Worker machines
- A node can be a VM or a physical machine, depending on the cluster
- Each node includes services necessary to run Pods and is managed by master components

# Pod Overview

cf. [Pod Overview](https://kubernetes.io/ja/docs/concepts/workloads/pods/pod-overview/)

- The smallest deployable unit in the Kubernetes object model

## Understanding Pods
- Pods are the basic execution units of Kubernetes applications
- Encapsulate application containers, storage resources, unique network IPs, and options for managing container execution

# ReplicaSet

cf. [ReplicaSet](https://kubernetes.io/ja/docs/concepts/workloads/controllers/replicaset/)

- Aims to maintain a stable set of replica Pods at all times

## When to Use ReplicaSet
- Ensures that a specified number of Pod replicas are running at all times

# Deployment

cf. [Deployment](https://kubernetes.io/ja/docs/concepts/workloads/controllers/deployment/)

- Provides declarative updates for Pods and ReplicaSets

# StatefulSet

cf. [StatefulSet](https://kubernetes.io/ja/docs/concepts/workloads/controllers/statefulset/)

- A workload API for managing stateful applications
- Manages scaling of a set of Pods and ensures order and uniqueness

# DaemonSet

cf. [DaemonSet](https://kubernetes.io/ja/docs/concepts/workloads/controllers/statefulset/)

- Ensures that all (or some) nodes run a copy of a Pod

# Job

cf. [Job](https://kubernetes.io/ja/docs/concepts/workloads/controllers/statefulset/)

- Creates one or more Pods and ensures that a specified number of them terminate successfully
- Tracks successful completion of Pods

# Service

cf. [Service](https://kubernetes.io/ja/docs/concepts/services-networking/service/)

- An abstract way to expose an application running on a set of Pods as a network service

## Motivation for Using Service
- Pods are designed to be ephemeral, and when they are created and stopped, they are not recreated
  - Using Deployment for application operation allows dynamic creation and deletion of Pods
- Each Pod has its own IP address

## Service Resources
- In Kubernetes, a Service defines a logical set of Pods and a policy for accessing them

## Service Exposure (Service Types)
- ClusterIP
  - Exposes the Service on a cluster-internal IP
- NodePort
  - Exposes the Service on a static port on each Node's IP
- LoadBalancer
  - Uses a cloud provider's load balancer to expose the Service externally
- ExternalName
  - Maps the Service to the contents specified in the externalName field by returning a CNAME record

# Configuration

cf. [Configuration](https://kubernetes.io/ja/docs/concepts/configuration/)

## Best Practices for Configuration
- General configuration tips
  - Use the latest stable API version
  - Configuration files should be stored in a version control system
  - Use YAML instead of JSON. While compatibility is similar, YAML is more user-friendly.
  - Group related objects into a single file when meaningful
  - Remember that many kubectl commands can be called on directories
  - Avoid specifying default values unnecessarily. Simplicity and minimalism reduce errors.
  - Add annotations to describe objects

## ConfigMap
- An API object used to store non-confidential data in key-value pairs.
  - ConfigMap does not provide confidentiality or encryption. Use Secret for sensitive data or additional third-party tools.
- Pods can use ConfigMap as environment variables, command-line arguments, or configuration files within volumes

## Secrets
- Allows storage and management of sensitive information like passwords, OAuth tokens, and SSH keys
- Can be included in Pod definitions or images

# Security

cf. https://kubernetes.io/ja/docs/concepts/security/

## Overview of Cloud Native Security

### The 4Cs of Cloud Native Security
- Security can be thought of in layers.
- The 4Cs of Cloud Native
  - Cloud
  - Cluster
  - Container
  - Code

### Infrastructure Security
- Concerns related to Kubernetes infrastructure
  - Network access to the API Server (control plane)
  - Network access to Nodes
  - Access to cloud provider APIs from Kubernetes
  - Access to etcd
  - Encryption of etcd

### Components Within the Cluster (Applications)
- Concerns related to workload security
  - RBAC authorization (access to Kubernetes API)
- Authentication
- Secret management for applications (and encryption when stored in etcd)
- PodSecurityPolicy
- Quality of Service (and cluster resource management)
- NetworkPolicy
- TLS for Kubernetes Ingress

### Containers
- Concerns related to containers
  - Vulnerability scanning and OS-dependent security
  - Image signing and enforcement
  - Do not allow privileged users

### Code
- Concerns related to code
  - Access only via TLS
  - Restrict communication port ranges
  - Security dependencies on third parties
  - Static code analysis
  - Dynamic probing attacks

# Impressions
The notes are quite abbreviated. It took a fair amount of time to read through the documentation...

# References
While progressing through the Kubernetes documentation, I also looked at external materials that were helpful.

- [slideshare.net - Understanding Kubernetes: Learn the Internal Structure and Architecture](https://www.slideshare.net/ToruMakabe/kubernetes-120907020)
- [qiita.com - Kubernetes Dojo Advent Calendar 2018](https://qiita.com/advent-calendar/2018/k8s-dojo)
- [qiita.com - Introduction Hands-on for Kubernetes Beginners](https://qiita.com/mihirat/items/ebb0833d50c882398b0f)
- [qiita.com - Fully Understand in a Few Hours! A Rather Intense Kubernetes Hands-on!!](https://qiita.com/Kta-M/items/ce475c0063d3d3f36d5d)
- [www.netone.co.jp - Introduction to Kubernetes Networking](https://www.netone.co.jp/knowledge-center/netone-blog/20191226-1/)
- [www.slideshare.net - "About Kubernetes" in About 30 Minutes](https://www.slideshare.net/YuyaOhara/30kubernetes-81054893)
- [cloud.google.com - Network Overview](https://cloud.google.com/kubernetes-engine/docs/concepts/network-overview?hl=ja)
