---
title: Notes on GCP
slug: gcp-notes-and-resources
date: 2023-05-22T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Google Cloud Platform
description: Miscellaneous notes taken while studying for the Google Cloud Certified Associate Cloud Engineer exam.
translation_key: gcp-notes-and-resources
---

# Overview
Miscellaneous notes taken while studying for the Google Cloud Certified Associate Cloud Engineer exam.

# Concepts
- Shared Responsibility Model
- Principle of Least Privilege
- Governance Across the Organization
- Delegation of Authority
- Stateless Applications and Immutable Infrastructure
- Reducing Toil Through Automation
- Scalability, High Availability, Robustness
- Monitoring and Automatic Notifications

cf. https://blog.g-gen.co.jp/entry/associate-cloud-engineer

# Network Service Tiers
- Premium Tier
  - Default
  - Best performance
    - Traffic between the internet and VMs within the VPC network is routed within Google's network whenever possible
  - Suitable for services requiring global availability
  - Unique to Google Cloud
- Standard Tier
  - Cost optimization
    - Traffic between the internet and VMs within the VPC network is usually routed via the internet
    - Suitable for services hosted entirely within a single region
    - Performance is comparable to other cloud providers

# Cloud Billing Account
- A Cloud Billing Account pays for the usage fees of Google Cloud projects and Google Maps Platform projects
- Cloud Billing Accounts are not billed from Google Workspace accounts
  - A separate Google Workspace billing account is required
- A Cloud Billing Account can be linked to one or more projects
- Types of Billing Accounts
  - Self-service
    - Automatically pays usage fees
  - Invoiced
    - Monthly invoicing
    - Requires a billing account

# Regions and Zones
## Regions
- Specific geographical locations where resources can be hosted
  - Locations of data centers
- Have three or more zones

## Zones
- Logical data centers
- Deployment areas for resources

## Examples of Multi-Region Services
- Firestore, Cloud Storage, BigQuery, Cloud Spanner, Cloud Bigtable

## Resource Placement
- Global Resources
  - Accessible from all resources in all zones belonging to the same project
  - Addresses, images, snapshots, instance templates, Cloud interconnect, VPC, firewalls, routes, global operations
- Regional Resources
  - Accessible from all resources within the same region
  - Addresses, Cloud interconnect attachments, subnets, regional managed instance groups, regional persistent disks, regional operations
- Zonal Resources
  - Usable only by other resources in the same zone
  - Instances, persistent disks, machine types, zonal managed instance groups, Cloud TPU, zonal operations
- Global Products
  - Not dependent on regions
  - Cloud CDN, Cloud DNS, Cloud Armor, Cloud Logging, Cloud Build
- Deployment per Region + One or More Multi-Region Deployments
  - Google App Engine and its features, Cloud Datastore, Cloud Storage, BigQuery

# IAM
## Members
- Google Account
- Service Account
- Google Group
- G Suite Domain
- Cloud Identity Domain

## Roles
- Basic Roles
  - Apply to all Google Cloud services within a project
  - Owner/Editor/Viewer
- Predefined Roles
  - Apply to specific Google Cloud services within a project
  - Allow fine-grained permissions for specific services
- Custom Roles
  - Define an exact set of permissions
- Policies
  - Bindings that associate members with roles

## Using IAM Securely
Refer to [Using IAM Securely](https://cloud.google.com/iam/docs/using-iam-securely?hl=ja#policy_management)

## Service Accounts
- Special accounts used by applications or computing workloads like Compute Engine instances, not users
- Identified by an account-specific email address
- Types of Service Accounts
  - User-managed Service Accounts
    - Created and managed by users
  - Default Service Accounts
    - Automatically created user-managed service accounts when enabling specific Google Cloud services
  - Google-managed Service Accounts
    - Created and managed by Google on behalf of users to allow services to access resources
- [Best Practices for Using Service Accounts](https://cloud.google.com/iam/docs/best-practices-service-accounts?hl=ja)

# VPC
## VPC Network
- Global resource
- Multiple VPC networks can be configured per project
- New projects are set up with one subnet per region and a default network (auto mode VPC network)
- Supports unicast addresses for both IPv4 and IPv6
- Does not support broadcast or multicast addresses within the network

## Subnet
- Regional resource
- Types of Subnets
  - IPv4-only subnets (IPv4 subnet range only)
  - Dual-stack subnets (both IPv4 and IPv6 subnet ranges)
- Subnet Creation Modes
  - Auto Mode VPC Network
    - Automatically creates one subnet per region within the network
  - Custom Mode VPC Network
    - Does not automatically create subnets
    - Allows full control over subnets and IP ranges
- Subnets within the same VPC can communicate even if they are in different segments
  - Possible even if they are in different regions within the same VPC

## Firewall
- Firewall Rules
  - Defined at the network level and apply only to the network where the rule was created
  - Names must be unique per project
  - Priority is specified as an integer from 0 to 65535, with lower numbers indicating higher priority

## Routes
- Inform instances and VPC networks how to send traffic from instances to destinations inside or outside Google Cloud

## VPC Network Peering
- Connects two VPC networks to allow resources within each network to communicate with each other
- Does not traverse the public internet
- All subnets can communicate using internal IPv4 addresses
- Dual-stack subnets can communicate using internal or external IPv6 addresses
- Allows private connections regardless of whether they belong to the same project or organization
- Use Cases
  - When an organization has multiple network management domains
  - When peering with another organization
- Benefits
  - Network Latency
    - Private connections have lower latency than public IP networks
  - Network Security
    - Services are not exposed to the internet
  - Network Cost
    - Communicating using internal IPs can save on Google Cloud egress (outbound) bandwidth costs

## Shared VPC
- Allows sharing a single VPC across multiple projects within the same organization
- Connects projects within the same organization
- Linked projects can be placed in the same or different folders
  - If in different folders, the administrator needs Shared VPC Admin permissions for both folders
- The project sharing the VPC is called the host project, and the projects using the shared VPC are called service projects. Service projects connect to the host project.

## Serverless VPC Access
- Service that allows direct connection from serverless environments like Cloud Run, App Engine, Cloud Functions to VPC networks

# Network Connectivity
## Connecting to Google Cloud
### Cloud VPN
- Securely connects to a VPC using an IPsec VPN connection
- Traverses the public internet
- Traffic between two networks is encrypted at one VPN gateway and decrypted at the other
  - Protects data transmitted over the internet
- Ipsec
  - Protocol for encrypting packets to ensure confidentiality and tamper detection

### Classic VPN
- On-premises communicates with VMs within a VPC network via one or more IPsec VPN tunnels

### HA VPN
- Requires the peer VPN device or service to support BGP
- High availability solution that securely connects on-premises networks to VPC networks using IPsec VPN connections within a single region
- Creating an HA VPN gateway automatically selects two external IPv4 addresses, one for each of the two interfaces

### Cloud Interconnect
- Connects on-premises to GCP (VPC) networks
- Does not traverse the public internet
- Low latency and high availability
- Dedicated Interconnect
  - Physically connects on-premises to GCP (VPC)
    - Requires physically connecting your network to Google's network at a supported colocation facility
- Partner Interconnect
  - Connects on-premises to GCP (VPC) via a supported service provider
  - Useful when unable to connect to a colocation facility or when a 10Gbs dedicated line is not required

### Cloud Router
- Fully distributed managed service that advertises IP address ranges using BGP (Border Gateway Protocol). Provides dynamic routing.
- Also functions as the control plane for Cloud NAT
- Provides BGP to the following services
  - Dedicated InterConnect
  - Partner Interconnect
  - Cloud VPN
  - Router Appliance
- Custom Route Advertisement
  - Select routes to advertise to on-premises routers via BGP

### Network Connectivity Center
- Allows applying hub-and-spoke architecture to network connectivity management
- Enables data transfer between sites

## Connecting to Google Workspace and Supported APIs
### Direct Peering
- Establishes a direct peering connection between your network and Google's edge network
- Available at over 100 locations in 33 countries worldwide
- Exists outside of Google Cloud
- Recommended to use Dedicated Interconnect or Partner Interconnect unless accessing Google Workspace applications

### Carrier Peering
- Uses a service provider to access Google applications like Google Workspace and obtain enterprise-class network services connecting infrastructure to Google
- Exists outside of Google Cloud

## Connecting to CDN Providers
### CDN Interconnect
- Third-party CDN providers establish direct peering with Google's edge network at various locations, allowing traffic to be transferred from VPC to the provider's network

# Cloud DNS
- High-performance, resilient global Domain Name System (DNS) service
- Zones
  - Public Zones
    - Exposed to the internet
  - Private Zones
    - Not discoverable from the internet
  - Delegated Zones
    - Allows zone owners to delegate subdomains to other name servers using NS records
  - Split Horizon DNS
    - Represents instances where two zones are created for the same domain, one for internal networks and another for external networks
    - Can return different information based on the query source
- Records
  - A: Address Record. Maps a hostname to an IPv4 address
  - AAAA: IPv6 Address Record. Maps a hostname to an IPv6 address
  - CNAME: Canonical Name Record. Specifies an alias name
  - MX: Mail Exchange Record. Used for routing requests to mail servers
  - NS: Name Server Record. Used for delegating DNS zones to authoritative servers
  - PTR: Pointer Record. Defines a name corresponding to an IP address. Used for reverse DNS
  - SOA: Start of Authority. Used for specifying the primary name server and administrator of a zone

# Cloud KMS
- Locations
  - Can be created in one of many locations
  - Some resources, like Cloud HSM keys, are not available in all locations
- Keys cannot be automatically deleted
- Automatic Key Rotation
  - Possible for keys used for symmetric encryption
  - Not possible for keys used for asymmetric encryption or signing
- Data is not re-encrypted during key rotation
- Keys and key rings cannot be deleted
  - Prevents resource name conflicts
- Key versions cannot be deleted
  - Resources become unusable by destroying key version material
  - After scheduling key version destruction, it is destroyed in 24 hours by default
    - Can be restored until destruction
- Keys cannot be exported
  - Can be imported, but only if protection level is HSM or SOFTWARE

# Cloud Deployment Manager
- Infrastructure deployment service for automating the creation and management of various GCP resources
- Deploy Resources
  - gcloud deployment-manager deployments create example --config example.yaml
- Verify New Deployment
  - gcloud deployment-manager deployments describe example
- Delete Resources
  - gcloud deployment-manager deployments delete example
- Preview New Configuration
  - gcloud deployment-manager deployments create example-config --config example.yaml --preview
  - Update Using Last Preview
    - gcloud deployment-manager deployments update example
  - Cancel Preview
    - gcloud deployment-manager deployments cancel-preview example
- Stop Updates
  - gcloud deployment-manager deployments stop my-first-deployment

# Cloud Foundation Toolkit
- Provides templates for Deployment Manager and Terraform that reflect Google Cloud best practices

# Resource Manager
## Resource Hierarchy
- Organization > Folder > Project > Resources for each service

## Tags and Labels
|                          |                                                  Tags                                                   |                                          Labels                                          |
| ------------------------ | ------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| Resource Structure       | Individual resources                                                                                     | Metadata of resources                                                                     |
| Definition               | Organization level                                                                                       | Per resource                                                                              |
| Access Control           | Requires IAM roles for management and attachment                                                         | Requires different IAM roles depending on the resource                                    |
| Attachment Prerequisites | Requires definition of tag keys and tag values before tags can be attached to resources                  | No prerequisites                                                                          |
| Deletion Requirements    | Cannot be deleted if tag bindings exist                                                                  | No conditions                                                                             |
| IAM Policy Support       | Available in IAM policy conditions                                                                       | Not supported                                                                             |
| Organization Policy Support | Available in organization policy conditional constraints                                               | Not supported                                                                             |
| Cloud Billing Integration | Chargeback, auditing, and other cost allocation analysis, exporting Cloud Billing cost data to BigQuery | Filter resources by label in Cloud Billing, export Cloud Billing data to BigQuery         |

Tags can be flexibly used for constraints, while labels are just labels.

# Compute Engine
- Instances
  - Virtual machines (VMs) hosted on Google's infrastructure
- Live Migration
  - Allows maintenance to be performed without interrupting workloads, restarting VMs, or changing VM properties
  - Prevents downtime due to maintenance
- OS Login
  - Allows SSH access to instances using IAM
  - No need to create and manage SSH authentication keys
  - Maintains Linux user IDs across instances when using OS Login
- Custom Images
  - Boot disk images owned and access-controlled by users
  - Can be created from the following sources
    - Persistent disk (including when the disk is attached to a VM)
    - Snapshot of a persistent disk
    - Another image in the project
    - Shared image from another project
    - Compressed RAW image in Cloud Storage
- Adding Extended Memory to Existing VM Instances
  - Stop existing instance → Update vCPU and memory → Restart
- Deletion Protection
  - Can be enabled to prevent specific instances from being deleted
  - Cannot be applied to MIGs, but can be applied to instances not belonging to MIGs
  - Cannot be specified in instance templates
  - Enable deletion protection when creating an instance
    - gcloud compute instances create [INSTANCE_NAME] --deletion-protection
  - Disable deletion protection when creating an instance
    - gcloud compute instances create [INSTANCE_NAME] --no-deletion-protection
  - Update deletion protection for existing instances
    - gcloud compute instances update [INSTANCE_NAME] [--deletion-protection | --no-deletion-protection]

## Compute Engine Related IAM Roles
- Compute Admin (roles/compute.admin)
  - Manages all Compute Engine resources
- Compute Image User (roles/compute.imageUser)
  - Lists and reads images
- Compute Instance Admin (roles/compute.instanceAdmin)
  - Beta version
  - Creates, modifies, and deletes virtual machine instances
  - Only manages instances, not networks or security
- Compute Instance Admin (roles/compute.instanceAdmin.v1)
  - v1
  - Manages all Compute Engine instances, instance groups, disks, snapshots, and images
- Compute Viewer (roles/compute.viewer)
  - Read access to retrieve and display Compute Engine resources

## Instance Groups
- Collection of VMs that can be managed as a single entity

### Managed Instance Groups (MIG)
- Allows application operations across multiple identical VMs
- Equipped with auto-scaling policies
  - Load judgment indicators
    - Target utilization indicators
      - Average CPU utilization
      - HTTP load balancing capacity
      - Cloud Monitoring indicators
  - Schedule
    - Specify capacity and schedule to scale according to expected load
- Predictive Autoscaling
  - When enabled, scales out based on past data predictions
- Apply rolling updates to all instances in a MIG
  - gcloud compute instance-groups managed rolling-action start-update INSTANCE_GROUP_NAME -version=template=INSTANCE_TEMPLATE_NAME [--zone=ZONE | --region=REGION]
- Enable automatic autoscaling for MIG
  - gcloud compute instance-groups managed set-autoscaling INSTANCE_GROUP_NAME
  - ex. Average CPU utilization 70%, maximum number of instances 3
    - gcloud compute instance-groups managed set-autoscaling INSTANCE_GROUP_NAME --max-num-replicas 3 --target-cpu-utilization 0.70

### Unmanaged Instance Groups
- Allows load balancing across a set of VMs managed by the user
  - Can include heterogeneous instances
- Does not support autoscaling, auto-healing, multi-zone support, rolling updates, or instance templates
- Create unmanaged instance group
  - gcloud compute instance-groups unmanaged create instance-group-name --zone=zone

## Instance Templates
- Resources used to create reservations for VMs or MIGs
- Define machine type, boot disk image or container image, labels, startup scripts, and other instance properties
- Can be used to
  - Create individual VMs
  - Create VMs within a MIG
  - Create VM reservations

## Discounts
- Sustained Use Discounts
  - Applied to resources used beyond 25% of the billing month and not covered by other discounts
- Committed Use Discounts
  - Applied by purchasing commitments
    - Allows committing to a minimum resource usage and minimum period of one or three years
    - No upfront payment, billed monthly
  - Two types of Compute Engine commitments
    - Resource-based
      - Suitable for predictable steady-state resource usage
    - Cost-based
      - Suitable for cases where cost needs are easily predictable

## Machine Types

| Type | Feature |  Machine Type  |
| ---- | ---- | ---- |
|  General Purpose |  Cost Optimization |  E2 |
|  General Purpose | Balanced (Cost Performance) | N2, N2D, N2 |
| General Purpose | Scale-out Optimization | Tau T2D, Tau T2A |
| Optimized Workloads | Memory Optimization | M3, M2, M1 |
| Optimized Workloads | Compute Optimization | C2, C2D |
| Optimized Workloads | Accelerator Optimization | A2 |

## Other Machine Types
- Custom Machine Types
  - Custom-defined by the user
  - Create instances with custom machine types
    - For N1 machine types (default machine type is N1)
      - ex. gcloud compute instances create example --custom-cpu=4 --custom-memory=5
    - For N2 machine types
      - ex. gcloud compute instances create example --custom-cpu=6 --custom-memory=3072MB --custom-vm-type=n2
      - ex. gcloud compute instances create example --machine-type n2-custom-6-3072
- Shared-core Machine Types
  - Share a single vCPU across multiple workloads
    - Burst capability allows additional vCPU usage for short periods
- Preemptible
  - Usable for up to 24 hours
  - May be interrupted
  - Inexpensive
  - No live migration

## Storage Options

| Type | Feature | Data Redundancy |
| ---- | ---- | ---- |
| Zonal Standard PD | Efficient, reliable block storage | Zone |
| Regional Standard PD | Synchronous replication across two zones within a region, efficient and reliable block storage | Multi-zone |
| Zonal Balanced PD | Cost-effective, reliable block storage | Zone |
| Regional Balanced PD | Synchronous replication across two zones within a region, cost-effective and reliable block storage | Multi-zone |
| Zonal SSD PD | Fast, reliable block storage | Zone |
| Regional SSD PD | Synchronous replication across two zones within a region, fast and reliable block storage | Multi-zone |
| Zonal Extreme PD | Highest performance persistent block storage | Zone |
| Local SSD | High-performance local block storage. Physically attached to the server, cannot be set as a boot disk. | None |
| Cloud Storage Bucket | Low-cost object storage | Region, Dual-region, Multi-zone |

Zonal XXX → Redundancy is zone
Region → Redundancy is multi-zone

cf. https://cloud.google.com/compute/docs/disks?hl=ja#disk-types

## Spot VM
- Latest version of Preemptible VM
- Unlike Preemptible VM, which lasts up to 24 hours after creation, Spot VM has no maximum expiration
- Detecting Preemption
  - gcloud compute operations list --filter="operationType=compute.instances.preempted"

## Shielded VM
- Enhanced VMs with a set of security controls to protect against rootkits and bootkits
- Protects against threats like remote attacks, privilege escalation, and malicious insiders

## Windows VM
- By default, authentication is done using username and password, not SSH
- If SSH is not enabled, new credentials must be generated before connecting to the VM
  - gcloud compute reset-windows-password VM_NAME

## Snapshots
- Snapshot Frequency
  - Hourly, daily, weekly
    - No monthly option
- Restrictions
  - Cannot delete a snapshot schedule associated with a disk
    - Must detach the schedule from all disks before deletion
  - To update a snapshot schedule, delete the existing schedule and create a new one
- Create Snapshot Schedule
  - gcloud compute resource-policies create snapshot-schedule

# Cloud Load Balancing
## Types
- Internal HTTP(S) Load Balancer
- External HTTP(S) Load Balancer
- Internal TCP/UDP Load Balancer
- External TCP/UDP Network Load Balancer
- External SSL Proxy Load Balancer
- External TCP Proxy Load Balancer
- Internal Regional TCP Proxy Load Balancer

The following LBs can configure both IPv4 and IPv6 external IP addresses.
- Global External HTTP(S) Load Balancer
- Global External HTTP(S) Load Balancer (Legacy)
- External SSL Proxy Load Balancer
- External TCP Proxy Load Balancer

## Choosing a Load Balancer
### External Load Balancing vs. Internal Load Balancing
- External Load Balancer
  - Distributes traffic from the internet to VPC networks
  - Use Premium Tier of Network Service Tiers for global load balancing
  - Standard Tier can be used for regional load balancing
- Internal Load Balancer
  - Distributes traffic to instances within GCP

### Global Load Balancing vs. Regional Load Balancing
- Global Load Balancing
  - Choose when backends are distributed across multiple regions, users need access to the same application and content, and access is provided via a single anycast IP address
  - Fault-tolerant to zone and region outages
- Regional Load Balancing
  - Choose when backends exist within the same region, only IPv4 termination is needed, or traffic needs to be confined to a specific region
  - Fault-tolerant to zone outages but not region outages

### Proxy Load Balancing vs. Pass-through Load Balancing
- Proxy Load Balancing
  - Terminates incoming client connections and opens new connections from the load balancer to backends
- Pass-through Load Balancing
  - Does not terminate client connections
  - Packets received by the load balancer are passed to backends without changing source, destination, or port information

# Cloud Scheduler
- Fully managed cron service that allows scheduling units of work to run at defined times or intervals
- Cron jobs created using Cloud Scheduler can be sent to one of the following targets
  - Public HTTP/S endpoints
  - Pub/Sub topics
  - App Engine HTTP/S applications
- Cost
  - Charged based on the number of jobs, not the number of job executions

# Cloud Batch
- Fully managed batch service that allows scheduling, queuing, and executing batch processing workloads
- Consists of the following components
  - Job
    - Scheduled program that completes a series of tasks without user intervention
    - Can be defined based on scripts, container images, or instance templates
  - Task
    - Programmatic action executed as part of a job
  - Resources
    - Infrastructure required to execute a job

# Cloud Run
- Connection Pool
  - Recommended to configure DB connection pools using client libraries
- Connection Limits
  - Cloud Run limits the number of connections to Cloud SQL to 100
    - Increase the number of instances to increase connections
- API Quota Limits
  - Cloud Run connects to Cloud SQL using Cloud SQL Auth Proxy, which uses Cloud SQL Admin API
    - API quotas apply to Cloud SQL Auth Proxy
      - Cloud SQL Admin API quota is approximately twice the number of Cloud SQL instances configured by the number of Cloud Run instances deployed for a specific service
      - To increase quota, limit or increase the number of Cloud Run instances
- CPU Allocation
  - CPU is allocated only during request processing
    - CPU is allocated only during request processing and released after response
  - Always on CPU
    - CPU is always allocated while the container instance exists
    - Autoscaling remains enabled
- Pricing Model
  - CPU is allocated only during request processing
    - Billed in 100ms increments for CPU and memory usage of container instances
    - Charges apply to idle instances if minimum instances are set to 1 or more
    - Billed per million requests for requests to deployed Cloud Run services
  - Always on CPU
    - No per-request billing, only billed for CPU and memory usage

# App Engine
- Build monolithic server-side rendered websites
- Regional
  - Infrastructure running the application is located in a specific region and redundantly available across all zones within the region
    - Region settings cannot be changed after project creation. To change, create a new project.

## app.yaml
- Manages app settings
- To configure a flexible environment, include the following in app.yaml
```yaml
env: flex
```

## Warmup Requests
- Reduces latency of requests and responses when application code is loaded into newly created instances
- To enable warmup, include the following in app.yaml

```yaml
inbound_services:
- warmup
```

## Deploy
- gcloud app deploy
- Use --no-promote flag to prevent automatic routing of traffic to the new version
  - gcloud app deploy --no-promote --version=[VERSION]

## Traffic
### Traffic Migration
- Switch traffic to a new version
  - gcloud app versions migrate [VERSION]
- Revert deployed version
  - If reverting to a previous version after deployment, traffic migration can be done to any version via Google Cloud Console
- Immediately migrate traffic to a specified version
  - gcloud app services set-traffic [MY_SERVICE] --splits [MY_VERSION]=1
- Gradually migrate traffic
  - gcloud app services set-traffic [MY_SERVICE] --splits [MY_VERSION]=1 --migrate

### Traffic Splitting
- Split traffic across multiple versions
  - gcloud app services set-traffic [MY_SERVICE] --splits [MY_VERSION1]=[VERSION1_WEIGHT],[MY_VERSION2]=[VERSION2_WEIGHT] --split-by [IP_OR_COOKIE]

## Scaling Types
- Autoscaling
  - Creates instances based on application metrics like request rate and response latency
  - Configure metrics, thresholds, and minimum number of always-on instances using the automatic_scaling element
- Basic Scaling
  - Instances are created when the application receives requests
  - Each instance shuts down when the application becomes idle
  - Suitable for intermittent processing or processing that operates based on user activity
- Manual Scaling
  - Specifies the number of instances that are always running regardless of load level
  - Allows running applications with complex initialization or those that depend on memory state over time

## Standard Environment
- Based on container instances running on Google's infrastructure

## Flexible Environment
- Based on Compute Engine

## Environment Comparison
|  | Standard Environment | Flexible Environment |
| --- | --- | --- |
| Instance Startup Time | Milliseconds | Minutes |
| Deployment Time | Seconds | Minutes |
| Background Processes | ✗ | ○ |
| SSH Access | ✗ | ○ |
| WebSocket | ✗ | ○ |
| Scaling | Manual, Basic, Auto | Manual, Auto |
| Scaling to Zero | ○ | ✗ (Minimum 1 instance) |
| Runtime Changes | ✗ | ○ (via Dockerfile) |
| Writing to Local Disk | ✗ | ○ |
| Third-party Binary Support | ✗ | ○ |
| Network Access | Via App Engine services | ○ |
| Pricing Model | Charges apply per instance class after exceeding daily free usage | Charges based on hourly resource (vCPU, memory, persistent disk) allocation |
| Automatic Shutdown | ○ | ✗ |

cf. https://cloud.google.com/appengine/docs/the-appengine-environments?hl=ja

Flexible environment offers more flexibility in features, configuration, and pricing. Instance startup and deployment times are shorter in the standard environment. The standard environment allows more flexible scaling.

# GKE
## Cluster Redundancy
- Zonal Cluster
  - Master: Single zone
  - Worker Nodes: Single zone
- Multi-Zonal Cluster
  - Master: Single
  - Worker Nodes: Multi-zone
- Regional Cluster
  - Master: Multi-zone
  - Worker Nodes: Multi-zone

All multi-zone is a regional cluster.

## Maintenance Window
- Mechanism to specify time periods for or against GKE upgrades

## Release Channels
- Mechanism to control versioning and upgrade policies in GKE's automatic upgrades
  - Rapid
    - Latest version available
  - Regular
    - Balances feature availability and release stability
  - Stable
    - Prioritizes stability over new features

Regular or Stable is recommended for production environments.

## Manual Node Pool Upgrades
- Surge Upgrade
  - Default
  - Rolls nodes to new versions sequentially
- Blue/Green Upgrade
  - Retains both old and new version nodes during upgrade
  - Faster rollback than surge upgrade but consumes more resources

Autopilot clusters use surge upgrades. Standard clusters use surge upgrades for automatic upgrades and can use either surge or blue/green upgrades for manual upgrades.

## Rolling Updates
  - Updates images, configurations, labels, annotations, resource limits/requests for workloads in a cluster
  - Gradually replaces pods with new ones and schedules pods on nodes with available resources
    - No downtime
  - Trigger workload rolling updates by updating Kubernetes workload Pod templates
    - Kubernetes workloads
      - DaemonSet
      - Deployment
      - StatefulSet

## Autoscaling
- Horizontal Pod Autoscaler (HPA)
  - Increases or decreases the number of pods based on CPU or memory consumption, custom metrics reported by K8S, or external metrics obtained from outside the cluster
- Vertical Pod Autoscaler (VPA)
  - Adjusts resource requests based on pod CPU or memory resources
    - Aims to improve stability and cost efficiency
- Multidimensional Pod Autoscaler (MPA)
  - Performs HPA and VPA simultaneously
- Cluster Autoscaler (CA)
  - Increases or decreases the number of nodes in a node pool
- Node Auto Provisioning (NAP)
  - Creates and deletes node pools

Generally based on K8S features or extensions of those features

# Cloud Functions
## Commands
- gcloud functions deploy
  - Deploys functions

## Supported Triggers (2nd Generation)
  - HTTP Trigger
  - Event Trigger
    - Pub/Sub Trigger
    - Cloud Storage Trigger
    - General Eventarc Trigger

1st Generation does not support Eventarc triggers and supports Firebase-related triggers

# Cloud Audit
## Types of Audit Logs
- Admin Activity Audit Logs
  - Log entries for API calls and other administrative actions that modify resource configurations or metadata
  - Always written
  - Cannot be configured, excluded, or disabled
- Data Access Audit Logs
  - Log entries for API calls that read resource configurations or metadata, or create, modify, or read user-provided resource data
  - Disabled by default
  - Audit logs are not generated for publicly accessible resources using allAuthenticatedUsers or allUsers
- System Event Audit Logs
  - Contains log entries for Google Cloud actions that modify resource configurations
  - Always written
  - Cannot be configured, excluded, or disabled
- Policy Denied Audit Logs
  - Recorded when Google Cloud services deny access to a user or service account due to security policy violations
  - Enabled by default
  - Cannot be disabled but can be limited with exclusion filters

# Cloud IDS
- Intrusion detection service
- Detects threats like intrusions, malware, spyware, and command & control attacks on the network
- Functions by creating a Google-managed peered network using mirrored VMs

# Google Cloud Directory Sync
- Service that synchronizes Google account data with Microsoft Active Directory or LDAP servers

# Cloud Storage
## Encryption
- By default, data is encrypted on the server before being written to disk
- Encryption Options
  - Server-Side Encryption
    - Encryption performed after Cloud Storage receives data
  - Customer-Managed Encryption Keys
    - Create and manage encryption keys using KMS
    - Customer-managed encryption keys can be stored as software keys in HSM clusters or externally
  - Customer-Supplied Encryption Keys
    - Create and manage your own encryption keys
    - Functions as additional encryption to enhance standard Cloud Storage encryption
  - Client-Side Encryption
    - Encrypts data before sending it to Cloud Storage
    - Data is sent encrypted to Cloud Storage, but server-side encryption is also performed
- Configure Website Settings for Buckets (Main Page or Error Page)
  - gsutil web set -m [<main_page_suffix>] [-e <error_page>] gs://<bucket_name>...
- Make All Objects in a Bucket Public
  - gsutil iam ch allUsers:objectViewer gs://<bucket_name>...
- Make a Specific Object Public
  - gsutil acl set public-read gs://<bucket_name>...

## Storage Classes

| Type | Feature | Availability SLA |
| ---- | ---- | ---- |
|  Standard | Used for most frequently accessed data. No minimum storage duration. No retrieval fees | 99.95% |
| Nearline | Accessed less than once a month, minimum storage duration of 30 days. Retrieval fees apply | 99.0% |
| Coldline | Accessed once a quarter, used for disaster recovery, etc. Minimum storage duration of 90 days. Retrieval fees apply | 99.0% |
| Archive | Accessed less than once a year, used for audit logs, archives, etc. Minimum storage duration of 365 days. Retrieval fees apply | 99.0% |

Storage Fees
Standard > Nearline > Coldline > Archive

Retrieval Fees
Standard < Nearline < Coldline < Archive

Storage class can be changed after bucket creation.

## Location Types
- Region
  - Specific region
- Dual-Region
  - Pair of regions (e.g., Tokyo and Osaka)
- Multi-Region
  - Two or more regions

Location type cannot be changed after bucket creation.
Data can be moved to a different location.

## Access Control
- Cloud IAM
- ACL
  - Access Control Lists that define users with access to buckets and objects and their access levels
  - Configuration Methods
    - Google Cloud Console, gsutil, client libraries, JSON API, XML API
      - Bucket ACLs cannot be changed in Google Cloud Console (objects can be)
- Signed URLs
  - Signed and time-limited encryption keys
  - Options for generating signed URLs
    - V4 signing with service account authentication
    - Signing with HMAC authentication
    - V2 signing with service account authentication
      - Legacy mechanism, not recommended
- Signed Policy Documents
  - Controls file upload policies

## Available Audit Logs
- Admin Activity Audit Logs
  - ADMIN_WRITE
    - Operations that modify the configuration or metadata of Google Cloud projects, buckets, or objects
- Data Access Audit Logs
  - ADMIN_READ
    - Operations that read the configuration or metadata of Google Cloud projects, buckets, or objects
  - DATA_READ
    - Operations that read objects are recorded
  - DATA_WRITE
    - Operations that create or modify objects are recorded

## Performance Optimization
- Upload Improvements
  - Upload small files quickly
    - Use the -m option
      - Performs parallel uploads
  - Efficiently upload large files
    - Use the parallel_composite_upload_threshold option
      - Splits into chunks and uploads in parallel
      - ex. gsutil -o GSUtil:parallel_composite_upload_threshold=150M cp ./localbigfile gs://your-bucket
  - Avoid sequential naming
    - Avoid sequential file naming conventions like YYYY/MM/DD/CUSTOMER/timestamp
- Download Improvements
  - Set optimal fetch size
  - Optimize large file reads with gsutil
    - Use HTTP Range GET requests to perform sliced downloads in parallel
    - ex. gsutil -o 'GSUtil:parallel_thread_count=1' -o
'GSUtil:sliced_object_download_max_components=8' cp gs://bukket/fileSRC.dat
./localDST.bin

## Consistency
- Global Strong Consistency
  - Read-after-write
  - Read-after-metadata-update
    - For buckets, metadata updates are strongly consistent, but resulting configuration reflection may take time (eventual consistency)
      - ex. When enabling object versioning for a bucket, wait at least 30 seconds before deleting or replacing objects
  - Read-after-delete
  - Bucket listing
  - Object listing
- Eventual Consistency
  - Granting or revoking access to resources
- Cache
  - Strong consistency may be compromised for publicly accessible, cached objects

## Logs
- Usage Logs
  - CSV file format
  - Information about all requests made to a specific bucket is created hourly
  - Can track access to resources using allUsers or allAuthenticatedUsers, which Cloud Audit Logs cannot
- Storage Logs
  - CSV file format
  - Information about bucket storage consumption for the previous day is created daily

# Storage Transfer Service
- Service that allows data transfer between file storage and objects across Google, Amazon, Azure, on-premises, etc.
- Predefined Roles
  - Editor (roles/editor)
  - Admin (admin)
  - User (user)
  - Viewer (viewer)

# GCP Database Services
Partial List

| Service Name | Type | Feature |
| ---- | ---- | ---- |
| BigTable | NoSQL (Columnar) | Low latency, high throughput |
| Firestore | NoSQL (Document) | Web, Native App, IoT, etc. |
| Firebase Realtime Database | NoSQL (Document) | Real-time synchronization |
| Memorystore | NoSQL (Key-Value) | Compatible with Redis/Memcached |
| BigQuery | Data Warehouse | Large datasets and queries |
| Bare Metal Solution | Hardware for RDB operation | Special requirements |

# Cloud SQL
- Machine Type Notation
  - db-n1<Machine Type>-standard<Usage>-1<CPU Count>
- Custom Instance Configuration
  - db-custom-1<vCPU Count>-3840<RAM(MB)>
- [Unsupported MySQL Statements in Cloud SQL](https://cloud.google.com/sql/docs/features?hl=ja#differences)
  - LOAD DATA INFILE
  - SELECT ... INTO OUTFILE
  - INSTALL PLUGIN ...
  - UNINSTALL PLUGIN ...
  - CREATE FUNCTION ... SONAME ...
    - CREATE PROCEDURE is supported
- Enabling/Disabling High Availability
  - Can be enabled when creating an instance or for existing instances
  - gcloud CLI
    - --availability-type=AVAILABILITY_TYPE
      - regional
        - High availability enabled.
      - zonal
        - Default value. High availability disabled.
- Cloud SQL Connection
  - Public IP
    - Configure authorized networks (whitelist of allowed source IPs) to allow internet access
  - Private IP
    - Access via VPC peering
      - Set up peering between Google's VPC (Cloud SQL instances belong to Google-managed VPC) and user-provided VPC
  - Cloud SQL Auth Proxy
    - Proxy feature for connecting to Cloud SQL
    - Automatically encrypts communication with SSL/TLS
    - Connection can be restricted with IAM permissions
    - IAM Database Authentication
      - Allows logging into databases with IAM

# Bigtable
- Distributed KVS
- Compatible with HBase, Hadoop
- Low latency within 10 milliseconds
- Scaling without restart
  - No downtime when adding or removing nodes
- Supports various data formats beyond structured and semi-structured
  - API Connections: VMs, Hbase, REST Server
  - Streaming Connections: Cloud Dataflow Streaming, Spark Streaming, Storm etc.
  - Batch Process: Hadoop, MapReduce, Dataflow, Spark etc.
- Optimal Use Cases
  - Time Series Data
  - Marketing Data
  - Financial Data
  - IoT Data
  - Graph Data

# Firestore
- Native Mode
  - Atomic Operations
    - Transactions
    - Batch Writes
      - Execute multiple write operations as a single batch
  - Firestore API mode
  - Document-oriented DB
  - Real-time Synchronization
    - Maintains persistence offline and synchronizes when online
  - Accessible from client and server
- Datastore Mode: Mode of the previously existing Datastore service
  - Atomic Operations
    - Transactions
  - Guarantees strong consistency
  - Data model with entity groups
  - Not compatible with Firestore

# Firebase Realtime Database
- Data stored as JSON
- Real-time Data Synchronization
  - Uses WebSocket
  - Maintains persistence offline and synchronizes when online
- Accessible from client without server
- Scaling

# Cloud Composer
- Workflow orchestration service based on Apache Airflow
- Configures GKE clusters

# Cloud Dataflow
- Fully managed data processing service supporting batch and streaming based on Apache Beam
- Batch processing is done with Dataflow alone, streaming processing is done in combination with Pub/Sub
- Dataflow SQL
  - Can execute queries on the following sources
    - Streaming data from Pub/Sub topics
    - Streaming or batch data from file sets in Cloud Storage
    - Batch data from BigQuery tables
  - Query results can be written to
    - Pub/Sub topics
    - BigQuery tables

# Cloud Dataprep by Trifacta
- Data cleansing service
- Configurable via GUI
- Integrated partner service operated by Trifacta, a Google partner
- Data sources are limited to BigQuery, Cloud Storage, and file uploads

# Cloud Dataproc
- Fully managed Map Reduce service supporting Apache Hadoop, Apache Spark, Apache Flink, Presto, etc.
- Enables data lake modernization, ETL, and secure data science in a scalable environment
- Master and worker instances are configured on Compute Engine, so it is not serverless

# Cloud Data Fusion
- GUI-based Dataproc
  - Configures pipelines using Dataproc clusters

# BigQuery
- Data warehouse service
- Columnar storage (column-oriented)
- Tree architecture

## Predefined Roles
- BigQuery Admin (roles/bigquery.admin)
- BigQuery Connection Admin (roles/bigquery.connectionAdmin)
- BigQuery Connection User (roles/bigquery.connectionUser)
- BigQuery Data Editor (roles/bigquery.dataEditor)
- BigQuery Data Owner (roles/bigquery.dataOwner)
- BigQuery Data Viewer (roles/bigquery.dataViewer)
- BigQuery Filtered Data Viewer (roles/bigquery.filteredDataViewer)
- BigQuery Job User (roles/bigquery.jobUser)
- BigQuery Metadata Viewer (roles/bigquery.metadataViewer)
- BigQuery Read Session User (roles/bigquery.readSessionUser)
- BigQuery Resource Admin (roles/bigquery.resourceAdmin)
- BigQuery Resource Editor (roles/bigquery.resourceEditor)
- BigQuery Resource Viewer (roles/bigquery.resourceViewer)
- BigQuery User (roles/bigquery.user)
- Masked Read (roles/bigquerydatapolicy.maskedReader)

## Query
- Standard SQL
  - Similar to SQL used in regular RDBs
- Legacy SQL
  - BigQuery SQL

## Cost
- Purchased capacity is a regional resource
  - Slot commitments purchased in a single region or multi-region cannot be used or moved to another region
- Analysis Fees
  - Costs for processing queries
  - On-demand fees or flat-rate fees
- Storage Fees
  - Costs for storing data loaded into BigQuery
  - Active Storage
    - Includes tables or table partitions modified in the past 90 days
  - Long-term Storage
    - Includes tables or table partitions not modified for 90 consecutive days
- No charges for queries that return errors or results from cache

## Cost Management
- Avoid SELECT *
- Use preview options to sample data
- Estimate costs before running queries
  - Use query validation tools
  - Perform dry runs
  - Use cost estimation tools
- Limit billed bytes to control query costs
  - Adjust settings for maximum billed bytes
- Use clustered or partitioned tables
- Do not use LIMIT for cost management in non-clustered tables
  - LIMIT does not affect read data volume in non-clustered tables
- Use dashboards to display costs and query audit logs
- Partition data by date
- Materialize query results incrementally
  - Cost of maintaining materialized results > Cost of processing large data
- Consider costs of large result sets
  - Maintaining large result sets incurs costs, so discard unnecessary data by applying table expiration

## Cost Estimation
  - Estimate costs for on-demand queries
    - Check query result bytes using one of the following and calculate estimates from bytes
      - Query validation tool in Google Cloud Console
      - --dry_run flag in bq command-line tool
      - Submit query jobs using API or client libraries with dryRun parameter

# BigQuery Data Transfer Service
- Managed service that automates data movement to BigQuery based on a schedule
- Supported Sources (Databases not supported)
  - Campaign Manager
  - Cloud Storage
  - Dataset Copy
  - Google Ad Manager
  - Google Ads (formerly AdWords)
  - Google Ads
  - Google Merchant Center
  - Google Play
  - Scheduled Queries
  - Search Ads 360
  - YouTube Channel
  - YouTube Content Owner
  - Amazon S3
  - Azure Blob Storage
  - Teradata
  - Amazon Redshift

# BI Tools
- Google Data Portal (Google Data Studio)
- BigQuery BI Engine
- Looker

# Machine Learning
- Cloud Vision API
  - Image analysis
- Speech-To-Text Text-To-Speech
  - Real-time transcription
- Cloud Natural Language API
  - Text analysis
- Cloud Translation API
  - Language translation
- Cloud Video Intelligence API
 - Video content analysis

# Transfer Appliance
- Large-capacity storage device
- Physical data transfer
- Uploads to Cloud Storage
- Criteria for Suitability
  - Using GCP
  - Data size of 10TB or more
  - Data located in a serviceable location
  - Data size that may take over a week to upload via network

# API
## Cloud Endpoints
- Usable when the backend is on GCP

## Apigee
- Usable even if the backend is not on GCP

# Cloud Logging
- Service for collecting, storing, and managing logs

## Log Types
- Logs generated by Google Cloud services
  - Most services like BigQuery, Cloud Run, Cloud Functions output logs to Cloud Logging
- Logs generated by users
  - Logs ingested via agents installed on VMs or API
- Security Logs
  - Cloud Audit Logs
  - Access Transparency Logs
    - Logs output when Google Support accesses user content
- Multi-cloud and Hybrid Cloud Logs
  - Logs ingested from other cloud services or on-premises

## Log Storage
- Log Buckets
  - Proprietary storage of Cloud Logging
- Cloud Storage Buckets
- BigQuery Datasets

## Pricing
- Ingestion Processing
  - Applies only when using log buckets
- Storage

## Sink
- Functionality for distributing logs ingested into Cloud Logging
  - Configuration values for distribution
    - Log destination, inclusion filter, exclusion filter
