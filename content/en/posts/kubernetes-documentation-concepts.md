---
title: Kubernetes Documentation Reading - Summary of Concepts
slug: kubernetes-documentation-concepts
date: 2020-10-20T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Docker
  - Containers
  - Kubernetes
translation_key: kubernetes-documentation-concepts
---

# Overview
To properly catch up with Kubernetes, I read the documentation and will leave my personal notes. Since it's long, I will only summarize the concepts section.

[kubernetes.io](https://kubernetes.io/ja/docs/home/)

# What is Kubernetes?
cf. [What is Kubernetes?](https://kubernetes.io/ja/docs/concepts/overview/what-is-kubernetes/)

## What is Kubernetes?
- Declarative configuration management
- Promotion of automation
- A platform for managing containerized workloads and services

## Looking Back at the Past
- Deployment in the era before virtualization (Traditional deployment)
  - No resource limitations for applications on physical servers
    - Resource allocation issues
  - Difficult to scale
  - High maintenance costs
- Deployment using virtualization (Virtualized deployment)
  - Applications can be isolated per VM
    - Data access restrictions between applications
  - Improved resource utilization within physical servers due to virtualization
  - Easier to add or update applications
    - Reduced hardware costs, improved scalability
- Deployment using containers (Container deployment)
  - Applications can share the OS
    - Lightweight
  - Containers have their own filesystem, CPU, memory, process space, etc.
  - Not dependent on cloud or OS distributions
  - Benefits of containers
    - Container images are easier and more efficient to create than VM images
    - Continuous building and deployment of container images is easier
    - Separation of concerns between development and operations
      - Creation of application container images occurs during build/release
    - High observability
      - In addition to OS-level information and metrics, the operational state of applications and other warnings
    - Consistency of environments
      - Can run the same way in development, testing, and production
    - Portability across cloud and OS distributions
      - Can run in any environment, whether on-premises or public cloud
    - Application-centric management
      - Shift from running OS on virtual machines to running applications on OS using logical resources.
    - High affinity with microservices
      - Compatible with loosely coupled, distributed, scalable, and flexible microservices
    - Resource partitioning
      - Predictable application performance
    - Efficient use and aggregation of resources

## Reasons Why Kubernetes is Needed and Its Features
- Service discovery and load balancing
  - Can expose containers via DNS names or IP addresses
  - Can distribute network traffic
- Storage orchestration
  - Can freely choose the storage to mount
- Automated rollouts and rollbacks
  - Can define the state of the containers to be deployed
- Automated bin packing
  - Can declare the CPU and memory (RAM) required by containers
  - Can adjust according to nodes, efficiently utilizing resources
- Self-healing
  - Can restart, replace, or kill containers that fail to start
- Secret and configuration management
  - Can update application configuration information without needing to recreate container images

## What Kubernetes Does Not Provide
- Kubernetes does not...
  - Limit the types of applications it supports
  - Deploy source code or build applications
  - Provide application-level (middleware, database, cache, etc.) functionality
  - Specify logging, monitoring, or alerting features
  - Provide configuration languages
  - Offer systems for machine configuration, maintenance, management, or self-healing
  - Assume orchestration

# Kubernetes Components

cf. [Kubernetes Components](https://kubernetes.io/ja/docs/concepts/overview/components/)
- When you deploy Kubernetes, a cluster is created
  - A cluster is a collection of nodes that run containerized applications
    - Every cluster has at least one worker node
  - Worker nodes host Pods, which are components of applications
  - Master nodes manage the worker nodes and Pods within the cluster
    - Using multiple master nodes can provide failover and high availability to the cluster
  - The control plane manages the worker nodes and Pods within the cluster
    - In production environments, multiple nodes can be used to provide fault tolerance and high availability
  - [Kubernetes Cluster Diagram](https://d33wubrfki0l68.cloudfront.net/7016517375d10c702489167e704dcb99e570df85/7bb53/images/docs/components-of-kubernetes.png)

## Control Plane Components
- Makes overall decisions about the cluster (scheduling, etc.)

### kube-apiserver
- Component that provides the Kubernetes API externally
- Designed to scale horizontally

### etcd
- A key-value store with consistency and high availability
- The storage location for all cluster information in Kubernetes
- Always create a backup plan when using etcd as a data store for Kubernetes

### kube-scheduler
- Monitors whether a newly created Pod has been assigned to a node; if not, selects a node to run that Pod

### kube-controller-manager
- Runs multiple controller processes
- Operates as a single process
  - Logically, each controller is compiled into a single executable file
- Includes the following controllers:
  - Node controller
    - Notifies and responds when a node goes down
  - Replication controller
    - Maintains the correct number of Pods for all replication controller objects
  - Endpoint controller
    - Links Services and Pods
  - Service account and token controller
    - Creates default accounts and API access tokens for new namespaces

### cloud-controller-manager
- Runs controllers that interact with the underlying cloud provider
- The following controllers have dependencies on the cloud provider:
  - Node controller
    - Checks with the cloud provider to determine if a node has been deleted after it stops responding
  - Routing controller
    - Sets up routing in the underlying cloud infrastructure
  - Service controller
    - Creates, updates, and deletes cloud provider load balancers
  - Volume controller
    - Creates, attaches, and mounts volumes, and manages volume adjustments with the cloud provider

## Node Components
- Provides the execution environment for Kubernetes and manages Pods running on all nodes

### kubelet
- An agent that runs on each node in the cluster
- Ensures that each container is running in a Pod

### kube-proxy
- A network proxy running on each node in the cluster that implements part of the Kubernetes Service concept

### Container Runtime
- Software responsible for running containers
- ex. Docker, containerd, CRI-O, etc...

## Add-ons
- Uses Kubernetes resources (DaemonSet, Deployment, etc.) to implement cluster features
- Provides cluster-level functionality
  - Add-on resources that require namespaces belong to the kube-system namespace
- Some add-ons include:
  - DNS
  - Web UI
  - Container resource monitoring
  - Cluster-level logging

# Kubernetes API

cf. [Kubernetes API](https://kubernetes.io/ja/docs/concepts/overview/kubernetes-api/)

- [API Reference](https://kubernetes.io/docs/reference/) for more information

# About Kubernetes Objects

cf. [About Kubernetes Objects](https://kubernetes.io/ja/docs/concepts/overview/working-with-objects/)

## Object spec and status
- Kubernetes objects have two nested object fields that manage the object's configuration:
  - spec
    - Describes the characteristics that the object should have as its desired state
  - status
    - Indicates the current state of the object

# Managing Kubernetes Objects

cf. [Managing Kubernetes Objects](https://kubernetes.io/ja/docs/concepts/overview/working-with-objects/object-management/)

## Management Methods
- Imperative commands
    - Targets current objects
    - Recommended for development project environments
- Imperative object configuration
    - Targets individual files
    - Recommended for production project environments
- Declarative object configuration
    - Targets directories of files
    - Recommended for production project environments

## Imperative Commands
- Users perform actions on the current objects in the cluster

## Imperative Object Configuration
- Specify the processing content, optional flags, and one or more file names in the kubectl command

## Declarative Object Configuration
- Users operate on configuration files stored locally
- The operations are not recorded in the files

# Object Names and IDs

cf. [Object Names and IDs](https://kubernetes.io/ja/docs/concepts/overview/working-with-objects/names/)

# Namespace

cf. [Namespace](https://kubernetes.io/ja/docs/concepts/overview/working-with-objects/namespaces/)

- Supports the operation of multiple virtual clusters on the same physical cluster
  - These virtual clusters are called Namespaces

# Nodes

cf. [Nodes](https://kubernetes.io/ja/docs/concepts/architecture/nodes/)

- Worker machines
- A single node can be a single VM or physical machine, depending on the nature of the cluster
- Each node includes the services necessary to run Pods and is managed by the master components

# Appearance of Pods

cf. [Appearance of Pods](https://kubernetes.io/ja/docs/concepts/workloads/pods/pod-overview/)

- The smallest deployable unit in Kubernetes' object model

## Understanding Pods
- Pods are the basic execution unit of Kubernetes applications
- Encapsulates application containers, storage resources, unique network IPs, and options for how containers are run

# ReplicaSet

cf. [ReplicaSet](https://kubernetes.io/ja/docs/concepts/workloads/controllers/replicaset/)

- Aims to maintain a stable set of replica Pods at all times

## When to Use ReplicaSet
- Guarantees that a specified number of Pod replicas are running at all times

# Deployment

cf. [Deployment](https://kubernetes.io/ja/docs/concepts/workloads/controllers/deployment/)

- Provides declarative update functionality for Pods and ReplicaSets

# StatefulSet

cf. [StatefulSet](https://kubernetes.io/ja/docs/concepts/workloads/controllers/statefulset/)

- A workload API for managing stateful applications
- Manages scaling of sets of Deployments and Pods, ensuring order and uniqueness of Pods

# DaemonSet

cf. [DaemonSet](https://kubernetes.io/ja/docs/concepts/workloads/controllers/statefulset/)

- Ensures that all (or some) nodes run a copy of a single Pod

# Job

cf. [Job](https://kubernetes.io/ja/docs/concepts/workloads/controllers/statefulset/)

- Creates one or more Pods and ensures that a specified number of Pods successfully terminate
- Tracks the successful termination of Pods

# Service

cf. [Service](https://kubernetes.io/ja/docs/concepts/services-networking/service/)

- An abstract way to expose applications running on a set of Pods as a network service

## Motivation for Using Services
- Pods are designed to stop, and when Pods are created and stopped, they are not recreated
  - Using Deployment for application operation dynamically creates and deletes Pods
- Each Pod has its own IP address

## Service Resources
- In Kubernetes, a Service defines a logical set of Pods and a policy to access that set of Pods

## Service Exposure (Service Types)
- ClusterIP
  - Exposes the Service on a cluster-internal IP
- NodePort
  - Exposes the Service on a static port on each Node's IP
- LoadBalancer
  - Exposes the Service externally using a cloud provider's load balancer
- ExternalName
  - Associates the Service with content specified in the externalName field by returning a CNAME record

# Configuration

cf. [Configuration](https://kubernetes.io/ja/docs/concepts/configuration/)

## Best Practices for Configuration
- General configuration tips
  - Use the latest stable API version
  - Configuration files should be stored in a version control system
  - Use YAML instead of JSON; while compatibility is similar, YAML is more user-friendly.
  - Group related objects into a single file when meaningful
  - Remember that many kubectl commands can also be called on directories
  - Do not specify default values unnecessarily; simpler and minimal configurations are less prone to errors.
  - Include annotations in object descriptions

## ConfigMap
- An API object used to store non-sensitive data as key-value pairs.
  - ConfigMap does not provide confidentiality or encryption. Use Secrets for sensitive information or additional third-party tools.
- Pods can use ConfigMaps as environment variables, command-line arguments, or configuration files in volumes.

## Secrets
- Allows you to store and manage sensitive information such as passwords, OAuth tokens, and SSH keys
- Can be included in Pod definitions or images

# Security

cf. https://kubernetes.io/ja/docs/concepts/security/

## Overview of Cloud-Native Security

### The 4Cs of Cloud-Native Security
- Security can be thought of in terms of layers.
- The 4Cs of cloud-native:
  - Cloud
  - Cluster
  - Container
  - Code

### Infrastructure Security
- Concerns regarding Kubernetes infrastructure include:
  - Network access to the API Server (control plane)
  - Network access to Nodes
  - Access to cloud provider APIs from Kubernetes
  - Access to etcd
  - Encryption of etcd

### Components within the Cluster (Applications)
- Concerns regarding workload security include:
  - RBAC authorization (access to Kubernetes API)
- Authentication
- Management of application Secrets (and encryption when stored in etcd)
- PodSecurityPolicy
- Quality of Service (and cluster resource management)
- NetworkPolicy
- TLS for Kubernetes Ingress

### Containers
- Concerns regarding containers include:
  - Vulnerability scanning and OS-dependent security for containers
  - Image signing and enforcement
  - Avoiding privileged users

### Code
- Concerns regarding code include:
  - Access only via TLS
  - Limiting the range of communication ports
  - Security dependencies on third parties
  - Static code analysis
  - Dynamic probing attacks

# Personal Thoughts
The notes are quite abbreviated. It took a fair amount of time to read through the documentation...

# References
While progressing through the Kubernetes documentation, I also looked at external materials, so here are some references that were helpful:

- [slideshare.net - Understanding Kubernetes: Learning the Internal Structure and Architecture](https://www.slideshare.net/ToruMakabe/kubernetes-120907020)
- [qiita.com - Kubernetes Dojo Advent Calendar 2018](https://qiita.com/advent-calendar/2018/k8s-dojo)
- [qiita.com - Hands-on Introduction for Kubernetes Beginners](https://qiita.com/mihirat/items/ebb0833d50c882398b0f)
- [qiita.com - Complete Understanding in a Few Hours! A Relatively Heavy Kubernetes Hands-on!!](https://qiita.com/Kta-M/items/ce475c0063d3d3f36d5d)
- [www.netone.co.jp - Introduction to Kubernetes Networking](https://www.netone.co.jp/knowledge-center/netone-blog/20191226-1/)
- [www.slideshare.net - Understanding Kubernetes in About 30 Minutes](https://www.slideshare.net/YuyaOhara/30kubernetes-81054893)
- [cloud.google.com - Overview of Networking](https://cloud.google.com/kubernetes-engine/docs/concepts/network-overview?hl=ja)