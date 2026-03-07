---
title: Technologies Behind bmf-tech
slug: bmf-tech-supporting-technologies
date: 2022-08-08T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Docker
  - Docker Compose
  - VPS
  - Golang
  - Vue.js
  - Prometheus
  - Promtail
  - Loki
  - Grafana
  - Nginx
description: An overview of the technology stack supporting the bmf-tech blog.
translation_key: bmf-tech-supporting-technologies
---

# Technologies Behind bmf-tech
This post describes the technology stack supporting this blog (bmf-tech.com).

# Previous bmf-tech Architecture
First, let's look at the architecture of the previous generation of bmf-tech.

<img style="width:757px;!important" alt="old_architecture" src="https://user-images.githubusercontent.com/13291041/183280770-84280c0f-e9ab-4cce-9f2d-0ea775e96ea5.png">

- The application was a monolithic architecture based on Laravel.
  - The API was built with PHP, and the admin panel was a React-based SPA.
    - These technologies were chosen somewhat arbitrarily at the time.
- Hosted on Sakura VPS.
  - No configuration management tools (like Ansible) were used; middleware was installed and configured manually.
    - A "warm" setup.
- No monitoring tools were in place.
  - Logs were checked by SSHing into the server directly.
- Containers were not used.
  - It was the era when Vagrant was mainstream or trending.
- Deployment was handled via Git hooks.

The previous application was the first custom CMS I built, called [Rubel](https://github.com/bmf-san/Rubel).

I don't remember exactly how long it was in operation, but probably around 3–5 years.

Before [Rubel](https://github.com/bmf-san/Rubel), I ran the blog using WordPress with custom themes.

The progression was: WordPress (custom theme 1) → WordPress (custom theme 2) → [Rubel](https://github.com/bmf-san/Rubel) → the current system.

The domain for bmf-tech.com was registered on November 2, 2015. Based on the domain age, the blog has been running for nearly seven years.

# Current bmf-tech Architecture
Now, let's discuss the current architecture of bmf-tech.

The same architecture is available as sample code in [gobel-example](https://github.com/bmf-san/gobel-example).

## Design
There were several reasons for wanting to replace the system:

- I wanted an opportunity to explore new technologies.
  - At the time, I was mainly working with PHP and wanted to try other languages.
- I lacked the motivation to maintain [Rubel](https://github.com/bmf-san/Rubel).
  - Laravel's update cycle was fast, requiring frequent updates.
    - I wanted to focus more on business logic rather than framework updates.
  - React felt overwhelming.
    - Redux was particularly challenging.
      - The scale didn't justify using FLUX.
  - I wanted a structure where the API could be separated, and the frontend could be easily discarded.
  - I wanted more control over the source code.
    - I felt uneasy about being overly dependent on frameworks.
  - Debugging system issues was difficult.
  - Etc.

Based on these reasons, I roughly outlined the design principles for the new system:

- Build a system that is easy to maintain in the long term.
  - Server configuration management:
    - Properly implement IaC.
      - Ensure idempotency.
      - Make it easy to switch servers if needed.
  - Application design:
    - Minimize dependency on frameworks.
    - Limit dependencies to standard libraries or custom libraries, except for the frontend.
  - Set up a monitoring environment:
    - Enable monitoring of system metrics, logs, and alerts.

## Architecture
The architecture built based on the design principles is as follows:

![スクリーンショット 2022-11-22 22 53 31](https://user-images.githubusercontent.com/13291041/203331548-95daeea8-8108-400a-91ae-35f8cddf899a.png)

- The server is hosted on ConoHa VPS instead of Sakura VPS.
  - Specs:
    - CPU: 2 cores
    - Memory: 1GB
    - SSD: 100GB
    - Image type: Ubuntu
  - ConoHa supports OpenStack, making it easier to manage instances with Terraform.
  - It's affordable and user-friendly.
- Server configuration management is handled with Ansible.
- SSL is provided by Let's Encrypt.
  - Certificates are obtained and renewed using [go-acme/lego](https://github.com/go-acme/lego).
    - cf. [legoでLet's encryptのSSL証明書をDNS-01方式で取得する](https://bmf-tech.com/posts/lego%E3%81%A7Let%27s%20encrypt%E3%81%AESSL%E8%A8%BC%E6%98%8E%E6%9B%B8%E3%82%92DNS-01%E6%96%B9%E5%BC%8F%E3%81%A7%E5%8F%96%E5%BE%97%E3%81%99%E3%82%8B)
  - Initially, I tried using a custom script for DNS-01 certificate acquisition but abandoned it due to occasional failures.
    - cf. [k2snow/letsencrypt-dns-conoha](https://github.com/k2snow/letsencrypt-dns-conoha)
- Docker is used for virtualization.
  - Docker Compose manages multiple containers.
  - I also experimented with Kubernetes.
    - [TerraformとAnsibleを使ってKubernetes環境構築](https://bmf-tech.com/posts/Terraform%E3%81%A8Ansible%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6Kubernetes%E7%92%B0%E5%A2%83%E6%A7%8B%E7%AF%89)
    - I wanted to manage a self-hosted Kubernetes cluster but struggled with bare-metal load balancers and gave up.
- Nginx serves as the web server for the admin panel (SPA), API (headless CMS in Go), and client (user-facing interface in Go).
  - The Nginx image is customized with a multi-stage build to include the admin panel's pre-built source code.
  - The API and client are built as Go binaries and included in the image.
- Monitoring is handled using Grafana as the UI.
  - Prometheus collects server metrics, while Promtail and Loki handle container log collection.
    - Grafana visualizes the collected data.
  - Initially, I used the EFK stack for log collection, but the 1GB server struggled with resource usage, so I switched to Loki and Promtail.
    - This setup is simpler and meets my requirements.

## Deployment
Deployment is straightforward.

<img style="width:820px;!important" alt="deploy" src="https://user-images.githubusercontent.com/13291041/183280768-78484c56-5775-4691-898b-f12b42d573e3.png">

- A private repository (bmf-tech) manages the container configuration.
  - [gobel-example](https://github.com/bmf-san/gobel-example) serves as the template.
    - While [gobel-example](https://github.com/bmf-san/gobel-example) uses the EFK stack, the private repository (bmf-tech) uses a Promtail × Loki setup.
  - The same container configuration can be run locally as in production.
  - Application source code is not included.
  - All environment variables are injected externally.
- Deployment involves running a script that:
  - SSHs into the server.
  - Pulls the bmf-tech repository.
  - Uploads environment variable files via rsync.
  - Runs `docker-compose build & up`.
  - There is no version management.
  - This approach causes brief downtime, but given the current traffic, it's not a significant issue.
    - It does affect availability, though...
- Initially, I tested deployment using Docker context but eventually abandoned it (I forgot why).

## Source Code Management
The following diagram illustrates how source code is managed based on the container configuration:

![NOTE - Source code management](https://github.com/bmf-san/bmf-tech-client/assets/13291041/1fb40523-cfc2-4030-82bd-10e7f38dafff)

- bmf-tech:
  - A private repository based on [gobel-example](https://github.com/bmf-san/gobel-example), managing the container configuration.
- bmf-tech-ops:
  - A private repository based on [gobel-ops-example](https://github.com/bmf-san/gobel-ops-example), managing server configuration, deployment scripts, and other operational code.
- [bmf-tech-client](https://github.com/bmf-san/bmf-tech-client):
  - A repository based on [gobel-client-example](https://github.com/bmf-san/gobel-client-example), managing the frontend source code.
  - Images are pushed to DockerHub as public images.
- gobel-api:
  - [gobel-api](https://github.com/bmf-san/gobel-api):
  - Manages the headless CMS source code.
  - Images are pushed to DockerHub as public images.
- gobel-admin-client:
  - [gobel-admin-client-example](https://github.com/bmf-san/gobel-admin-client-example):
  - Used as-is to manage the admin panel source code.
  - Images are pushed to DockerHub as public images.

## Application Design
Details of the API, Client (user-facing interface), and Admin (admin panel) applications:

### API
- Built as a headless CMS API.
- Developed using Go:
  - Simple language specifications, strong backward compatibility, fast compilation, portability with binaries, strong typing, and a rich standard library make it a good match for containers.
- Implements Clean Architecture:
  - While opinions on the compatibility of Go and Clean Architecture vary, I decided to try it.
  - I believe adhering to "separation of concerns" in Clean Architecture helps maintain long-term application maintainability.
- Minimal dependencies beyond standard packages.
  - cf. [go.mod](https://github.com/bmf-san/gobel-api/blob/master/app/go.mod)
- Uses REST as the API protocol.
  - If I were building it now, I might consider gRPC.

### Client
- A web server that responds with pages.
- Implements an API client in Go.
- Embeds template files (HTML) into the binary.
- Uses a custom CSS framework for design.
  - cf. [sea.css](https://github.com/bmf-san/sea.css)

### Admin
- Admin panel.
- Authentication uses JWT tokens.
- Session management uses Redis.
- Developed using Vue.js:
  - Chosen for its simplicity and ease of use compared to React.
    - I haven't used React recently, but I started using Vue more often and adopted it to gain experience.
  - Experimented with Atomic Design.
    - Not sure if it was implemented well.
  - Upgraded from Vue 2 to Vue 3 during development.
  - Considered introducing TypeScript but postponed it.
    - The admin panel is not expected to require frequent feature additions, so maintenance is assumed to be limited to library updates.
    - Considering the changing trends in frontend frameworks, I decided to keep it simple.
- SPA is served via Nginx.

## Database Design
The database design is mostly inherited from [Rubel](https://github.com/bmf-san/Rubel), but logical and physical deletions were reviewed and redesigned in some areas. Column data types and sizes were also reviewed.

## Migration
Data migration was handled using a custom migration tool I developed: [migrate-rubel-to-gobel](https://github.com/bmf-san/migrate-rubel-to-gobel).

Since there were no significant differences in database design, the migration tool was implemented in about 2–3 days.

Server migration involved setting up the new server environment, obtaining a test domain for verification, and conducting various checks. The release process was completed by simply switching DNS.

The old environment was deleted and the contract terminated after a month of stable operation in the new environment.

## Monitoring
Monitoring dashboards and alerts were created and configured in Grafana.

The monitoring dashboards are managed as JSON files for provisioning, while alerts were initially configured via the Grafana UI.
~(Alert provisioning was not supported at the time cf. [github.com - grafana/issues/36153](https://github.com/grafana/grafana/issues/36153)~ → Now supported and implemented provisioning for alerts.

## SLI & SLO
Although the traffic is low, I wanted to set these up to monitor and ensure consistent availability. This is still a work in progress.

## Load Testing
I am considering conducting load testing as a future task.

## Summary of Created Tools
Here is a summary of the tools and resources created during the development of the new bmf-tech:

- Applications:
  - [gobel-api](https://github.com/bmf-san/gobel-api)
  - [gobel-admin-client-example](https://github.com/bmf-san/gobel-admin-client-example)
  - [gobel-client-example](https://github.com/bmf-san/gobel-client-example)
  - [gobel-example](https://github.com/bmf-san/gobel-example)
  - [gobel-ops-example](https://github.com/bmf-san/gobel-ops-example)
- Libraries:
  - [goblin](https://github.com/bmf-san/goblin): Trie-based router.
  - [golem](https://github.com/bmf-san/golem): Simple JSON logger with log level support.
  - [goemon](https://github.com/bmf-san/goemon): Go-based dotEnv. Created but ultimately unused.
- Tools:
  - [migrate-rubel-to-gobel](https://github.com/bmf-san/migrate-rubel-to-gobel): Data migration tool.
- Boilerplates (created during testing):
  - [go-clean-architecture-web-application-boilerplate](https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate): A base repository for implementing Clean Architecture in Go.
  - [go-production-boilerplate](https://github.com/bmf-san/go-production-boilerplate): Repository for testing deployment methods using Docker context.
  - [docker-based-monitoring-stack-boilerplate](https://github.com/bmf-san/docker-based-monitoring-stack-boilerplate): Repository created for reusing the monitoring environment setup.
  - [terraform-ansible-openstack-boilerplate](https://github.com/bmf-san/terraform-ansible-openstack-boilerplate): Repository for testing Terraform and Ansible integration.
  - [setup-kubernetes-cluster-on-vps-boilerplate](https://github.com/bmf-san/setup-kubernetes-cluster-on-vps-boilerplate): Repository for testing Kubernetes cluster setup on VPS. Abandoned due to issues with bare-metal load balancers.
  - [vue-js-boilerplate](https://github.com/bmf-san/vue-js-boilerplate): Repository created to catch up with recent developments in Vue.js.

While creating these resources, I also wrote blog posts, gave lightning talks, and engaged in other activities, which significantly extended the time it took to release the new bmf-tech.

## Future Plans
There were times when I paused development or considered switching to WordPress or another existing system, but I am relieved to have finally completed a functional system.

There are still many unresolved issues and things I want to do, which I plan to address gradually as a hobby.

Rather than just building something, I want to focus on how to build and operate systems effectively. I see this blog system as an investment in that learning process.

The reasons I run my own blog are primarily for learning:

- For personal growth:
  - Writing articles helps organize my thoughts.
  - Building, hosting, and using my own blog system provides learning opportunities.
- To gain experience in system construction and operation:
  - Operating a system where I have full control offers valuable learning experiences.
- To earn a little extra money:
  - While living off ad revenue is a joke, it would be nice to cover server costs.
  - I don't focus on monetization, but in the past 1–2 years, I've managed to cover a portion of the annual operating costs, which makes me think this might be worthwhile in the long run.

For now, I plan to continue operating the current system at a relaxed pace.