---
title: Technologies Supporting bmf-tech
description: "Explore Docker, Golang, Vue.js, Nginx, Prometheus, and Grafana powering modern blogging infrastructure and monitoring systems."
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
translation_key: bmf-tech-supporting-technologies
---

# Technologies Supporting bmf-tech
This post discusses the technology stack that supports this blog (bmf-tech.com).

# Previous bmf-tech Configuration
First, let's look at the previous generation configuration of bmf-tech.

<img style="width:757px;!important" alt="old_architecture" src="/assets/images/posts/bmf-tech-supporting-technologies/183280770-84280c0f-e9ab-4cce-9f2d-0ea775e96ea5.png">

- The application was a monolithic structure based on Laravel.
  - The API was built with PHP, and the admin panel was a SPA built with React.
    - Technologies were adopted somewhat randomly based on what I was using at the time.
- Hosted on Sakura VPS.
  - I didn't use any configuration management tools (like Ansible) and manually installed middleware and configured the setup.
    - A "warm" configuration.
- No monitoring tools were provided.
  - When I wanted to see logs, I would ssh into the server and check directly.
- Containers were not used.
  - It was an era when Vagrant was mainstream or trending.
- Deployment was done using git hooks to manage everything.

The application I was operating in the previous generation was the first self-made CMS called [Rubel](https://github.com/bmf-san/Rubel).

I don't remember how many years it was in operation, but I think it was around 3 to 5 years.

Before operating [Rubel](https://github.com/bmf-san/Rubel), I was running a blog with an original theme on Wordpress.

Wordpress (Original Theme 1) → Wordpress (Original Theme 2) → [Rubel](https://github.com/bmf-san/Rubel) → and here we are now.

When I checked the domain age of bmf-tech.com, it was registered on November 2, 2015.

I forgot when I started operating the blog, but based on the domain age, it has been running for nearly 7 years.

# Current bmf-tech Configuration
Now, let's discuss the current configuration of bmf-tech.

I have published sample code with the same configuration at [gobel-example](https://github.com/bmf-san/gobel-example).

## Design
There were several reasons I wanted to replace the system.

- I wanted to have the opportunity to touch new technologies.
  - At that time, I was mainly working with PHP, so I wanted a chance to work with other languages.
- I lacked the motivation to maintain [Rubel](https://github.com/bmf-san/Rubel).
  - Laravel's update cycle was fast, requiring frequent updates.
    - I wanted to focus on business logic as much as possible rather than framework updates.
  - React was beyond my capabilities.
    - I struggled with Redux.
      - The scale did not require FLUX.
  - I wanted to separate the API and make it easy to discard the front end.
  - I wanted to have more control over the source code.
    - I felt a vague sense of crisis about being too dependent on the framework.
  - It was difficult to trace system bugs.
  - etc...

Based on these reasons, I roughly thought about the design policy for the new system.

- Build a system that is easy to maintain in the long term.
  - Server configuration management.
    - Properly implement IaC.
      - Ensure idempotency.
      - Make it easy to switch servers if needed.
  - Application design.
    - Avoid dependency on frameworks.
    - Limit dependencies to standard libraries or custom libraries outside of the front end.
  - Build a monitoring environment.
    - Enable monitoring of system metrics and logs, and set alerts.

## Architecture Configuration
The architecture configuration built based on the design policy is as follows.

![Screenshot 2022-11-22 22 53 31](/assets/images/posts/bmf-tech-supporting-technologies/203331548-95daeea8-8108-400a-91ae-35f8cddf899a.png)

- Instead of Sakura VPS, I adopted ConoHa VPS.
  - Specs:
    - CPU: 2 Core
    - Memory: 1 GB
    - SSD: 100 GB
    - Image Type: Ubuntu
  - ConoHa supports OpenStack, making instance construction easy to manage with Terraform.
  - It's also affordable and user-friendly.
- I use Ansible for server configuration management.
- SSL is provided by Let's Encrypt.
  - Certificate acquisition and renewal are handled by [go-acme/lego](https://github.com/go-acme/lego).
    - cf. [Obtaining Let's Encrypt SSL Certificates via DNS-01 with lego](https://bmf-tech.com/posts/lego%E3%81%A7Let%27s%20encrypt%E3%81%AESSL%E8%A8%BC%E6%98%8E%E6%9B%B8%E3%82%92DNS-01%E6%96%B9%E5%BC%8F%E3%81%A7%E5%8F%96%E5%BE%97%E3%81%99%E3%82%8B)
  - Initially, I attempted to obtain DNS-01 certificates with my own script without using lego, but I gave up because it occasionally failed to acquire them.
    - cf. [k2snow/letsencrypt-dns-conoha](https://github.com/k2snow/letsencrypt-dns-conoha)
- I adopted Docker for virtualization.
  - I manage multiple containers using Docker Compose.
  - I also experimented with Kubernetes.
    - [Building a Kubernetes Environment with Terraform and Ansible](https://bmf-tech.com/posts/Terraform%E3%81%A8Ansible%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6Kubernetes%E7%92%B0%E5%A2%83%E6%A7%8B%E7%AF%89)
    - I wanted to operate my own Kubernetes cluster, but I gave up because I couldn't handle a bare-metal load balancer.
- Nginx serves as the web server for the admin panel (SPA), API (headless CMS, Go), and Client (user-side screen, Go).
  - The Nginx image is customized to include the admin panel image through multi-stage builds, so the built source code for the admin panel is included.
  - The API and Client are images that contain binaries built with Go.
- Monitoring is done using Grafana as the UI.
  - Various metrics collection for the server is handled by Prometheus, while container log collection is managed by Promtail and Loki.
    - Visualization of collected data is done with Grafana.
  - Initially, I built an EFK stack for container log collection, but due to the 1GB server specs, resources were tight, so I switched to a configuration using Loki and Promtail.
    - It seemed functionally sufficient and simple to use, which appeared to meet my requirements.

## Deployment
The deployment process is not particularly complex.

<img style="width:820px;!important" alt="deploy" src="/assets/images/posts/bmf-tech-supporting-technologies/183280768-78484c56-5775-4691-898b-f12b42d573e3.png">

- I have a private repository (bmf-tech) to manage the container configuration.
  - It uses [gobel-example](https://github.com/bmf-san/gobel-example) as a template.
    - [gobel-example](https://github.com/bmf-san/gobel-example) is an EFK stack, but the private repository (bmf-tech) uses a configuration with promtail and Loki.
  - I can run containers in the local environment with the same container configuration as the production environment.
  - The application source code is not included at all.
  - All environment variables are configured to be injected from outside.
- Deployment is done by running a script that sshs into the server, pulls bmf-tech, uploads environment variable files via rsync, and then runs docker-compose build & up.
  - There is no particular version control.
  - A drawback is that there is a momentary downtime, but considering the current traffic, it is not a significant issue.
    - It affects availability, though...
- Initially, I was testing a deployment method using docker context, but it was not adopted (I forgot the reason...)

## Source Code Management
The following diagram shows how source code management is structured based on the container configuration.

![NOTE - Source code management](https://github.com/bmf-san/bmf-tech-client/assets/13291041/1fb40523-cfc2-4030-82bd-10e7f38dafff)

- bmf-tech
  - A private repository managed based on the template of [gobel-example](https://github.com/bmf-san/gobel-example) for container configuration.
- bmf-tech-ops
  - A private repository managed based on the template of [gobel-ops-example](https://github.com/bmf-san/gobel-ops-example) for managing server configuration and deployment scripts, related to construction and operational operations.
- [bmf-tech-client](https://github.com/bmf-san/bmf-tech-client)
  - A repository managed based on the template of [gobel-client-example](https://github.com/bmf-san/gobel-client-example) for managing the front-end source code.
  - The images are pushed to DockerHub and managed as public images.
- gobel-api
  - [gobel-api](https://github.com/bmf-san/gobel-api)
  - Manages the source code for the headless CMS.
  - The images are pushed to Dockerhub and managed as public images.
- gobel-admin-client
  - Directly uses [gobel-admin-client-example](https://github.com/bmf-san/gobel-admin-client-example).
  - Manages the source code for the admin panel.
  - The images are pushed to Dockerhub and managed as public images.

## Application Design
This section discusses the applications for API, Client (user-side screen), and Admin (admin panel).

### API
- Built as a headless CMS.
- Uses Go.
  - The language specification is simple, backward compatibility is well maintained, compilation is fast, and it runs as a binary when built, making it compatible with containers. It has types and a rich standard library, etc.
- Adopted Clean Architecture.
  - There may be mixed opinions about the compatibility of Go and Clean Architecture, but I tried it out.
  - I believe that adhering to "separation of concerns" in Clean Architecture is one way to keep the application easy to maintain in the long term.
- There are few dependencies outside of the standard package.
  - cf. [go.mod](https://github.com/bmf-san/gobel-api/blob/master/app/go.mod)
- The API protocol is REST.
  - If I were to create it now, gRPC might have been a good option.

### Client
- A web server that serves the screen.
- Implements the API client in Go.
- Template files (html) are embedded in the binary.
- Uses a custom CSS framework for design.
  - cf. [sea.css](https://github.com/bmf-san/sea.css)

### Admin
- Admin panel.
- Authentication is done via JWT token authentication.
- Session management is handled by Redis.
- Adopted Vue.js.
  - It seems simpler to write than React and easier to work with.
    - I haven't touched React recently, but I've started using Vue in various places, so I adopted it to gain insights.
  - I tried to challenge Atomic Design.
    - However, I'm not sure if I did it well.
  - During development, I updated from version 2 to 3.
  - I considered introducing TypeScript, but I fell behind.
    - I assumed that the admin panel wouldn't require frequent feature additions, so I planned to only maintain library updates.
    - Considering that frontend trends may change over the years and new frameworks may emerge, I thought it might become complicated if I invested too much effort, so I kept it as simple as possible.
- SPA delivery is handled by Nginx.

## DB Design
Basically, I inherited the DB design from [Rubel](https://github.com/bmf-san/Rubel), but I reviewed and redesigned some parts regarding logical and physical deletions.
I also reviewed the data types and sizes of columns.

## Migration Work
I wrote my own data migration tool to handle data migration.

[migrate-rubel-to-gobel](https://github.com/bmf-san/migrate-rubel-to-gobel)

Since there were no significant differences in DB design, I was able to implement the migration tool in about 2 to 3 days.

Regarding server migration, I didn't do much.
I prepared the server environment for migration, obtained a verification domain for operation confirmation, and conducted various operational checks in the new server environment.
I repeatedly created and destroyed to check for issues with IaC.

During the release process, the migration to the new environment was completed simply by switching the DNS.

I deleted the old environment after confirming that there were no issues with the new environment for about a month and completed the contract termination process.

## Monitoring
Monitoring dashboards and alerts were created and configured using Grafana.

The monitoring dashboard is managed in JSON file format, allowing for provisioning, but alerts are set from the Grafana UI.
~(Alerts have not yet been supported for provisioning, cf. [github.com - grafana/issues/36153](https://github.com/grafana/grafana/issues/36153)~ → This has been addressed, and provisioning is now supported.)

## SLI/SLO
It feels futile to set this up with little traffic, but I want to set it up to observe whether I can maintain a certain level of availability stably, rather than focusing on traffic. However, it is still not addressed.

## Load Testing
I am considering doing some load testing, so it's under consideration.

## Summary of What I Created
Here’s a summary of what I created before releasing the new bmf-tech.

- Applications
  - [gobel-api](https://github.com/bmf-san/gobel-api)
  - [gobel-admin-client-example](https://github.com/bmf-san/gobel-admin-client-example)
  - [gobel-client-example](https://github.com/bmf-san/gobel-client-example)
  - [gobel-example](https://github.com/bmf-san/gobel-example)
  - [gobel-ops-example](https://github.com/bmf-san/gobel-ops-example)
- Libraries
  - [goblin](https://github.com/bmf-san/goblin)
    - Trie-based router.
    - I spent quite a bit of time on this, so I wrote various blog posts about it.
  - [golem](https://github.com/bmf-san/golem)
    - Simple JSON logger. Allows specifying log levels.
  - [goemon](https://github.com/bmf-san/goemon)
    - Go-based dotEnv.
    - I created it but ended up not using it.
- Tools
  - [migrate-rubel-to-gobel](https://github.com/bmf-san/migrate-rubel-to-gobel)
    - Data migration tool.
- Boilerplates (created during the verification process)
  - [go-clean-architecture-web-application-boilerplate](https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate)
    - A repository created to provide a base for doing Clean Architecture in Go.
  - [go-production-boilerplate](https://github.com/bmf-san/go-production-boilerplate)
    - A repository for testing deployment methods using docker context.
  - [docker-based-monitoring-stack-boilerplate](https://github.com/bmf-san/docker-based-monitoring-stack-boilerplate)
    - A repository created for reuse when building a monitoring environment.
  - [terraform-ansible-openstack-boilerplate](https://github.com/bmf-san/terraform-ansible-openstack-boilerplate)
    - A repository tested for using Terraform and Ansible together.
  - [setup-kubernetes-cluster-on-vps-boilerplate](https://github.com/bmf-san/setup-kubernetes-cluster-on-vps-boilerplate)
    - A repository tested for building a Kubernetes cluster on VPS. I gave up because I couldn't handle a bare-metal load balancer.
  - [vue-js-boilerplate](https://github.com/bmf-san/vue-js-boilerplate)
    - A repository created to catch up on how recent Vue.js looks.

While creating the above, I wrote blog posts, gave lightning talks, and did various other things, so I spent quite a bit of time on the release of the new bmf-tech.

## Future Prospects
I have occasionally paused development, strayed off course, and considered switching to Wordpress or another existing system several times, but I have successfully reached a manageable operational state.

There are various issues I want to address and things I want to do, so I plan to tackle them as a hobby on the side.

I want to refine not just what to create, but how to create it and how to operate it, so I plan to invest in this through the blog system.

The reason I operate my self-made blog is strongly tied to learning. It has provided me with many learning opportunities, and I feel I can continue to learn even more in the future.

- For my own learning.
  - Writing articles to organize my learning.
  - By creating and hosting my own blog system, I believe it will provide me with further learning opportunities.
- Gaining system construction and operation experience.
  - I believe that operating a system I fully understand will allow me to learn a lot.
- Earning some pocket money.
  - While I joke about wanting to live off advertising revenue, I do hope to earn at least enough to cover server operation costs.
  - I consider it a side benefit, so I am not focusing on monetization.
    - Over the past 1-2 years, I have been making some profit that covers a portion of the annual operating costs, so I optimistically think there may be some value in continuing this.

For the time being, I believe I can continue to operate the current system, so I plan to take my time with it.