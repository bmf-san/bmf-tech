---
title: Continuous Profiling with Pyroscope
slug: continuous-profiling-with-pyroscope
date: 2023-05-07T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Profiling
translation_key: continuous-profiling-with-pyroscope
---

# Overview
I tried implementing a tool for Continuous Profiling called [Pyroscope](https://pyroscope.io/).

For more information on Continuous Profiling, refer to this article: [What is continuous profiling?](https://www.cncf.io/blog/2022/05/31/what-is-continuous-profiling/)

It seems that Grafana acquired it earlier this year.

[Grafana Labs Acquires Pyroscope to Add Code Profiling Capability](https://devops.com/grafana-labs-acquires-pyroscope-to-add-code-profiling-capability/)

Since the acquisition, it appears that the official name is Grafana Pyroscope?

There is a plugin available for Grafana, allowing integration, but Pyroscope also has its own UI.

You can try the [Demo](https://demo.pyroscope.io/?name=hotrod.python.frontend%7B%7D&query=) to see what you can view.

The code is publicly available as OSS, so if you're curious about the implementation, you can check it out here: [grafana/pyroscope](https://github.com/grafana/pyroscope)

# Installation
It's good to review the configuration beforehand.

[Pyroscope Agent](https://pyroscope.io/docs/agent-overview/)

## 1. Installing Pyroscope Server
You can use the image pushed to DockerHub.

[pyroscope/pyroscope](https://hub.docker.com/r/pyroscope/pyroscope)

Here is the Docker installation guide: [Docker Guide](https://pyroscope.io/docs/docker-guide/)

There is also a Kubernetes installation guide available: [Kubernetes/Helm](https://pyroscope.io/docs/kubernetes-helm-chart/)

## 2. Enabling Profiling on the Application Side
Set up profiling and install the agent on the application side.

The basic approach is Push-based, but for Go, there is a Pull-based option, allowing management of targets on the Pyroscope server.

[Go Pull Mode](https://pyroscope.io/docs/golang-pull-mode/)

When implementing in Pull mode with Go, you can manage targets as follows:

```yml
---
scrape-configs:
  - job-name: pyroscope
    scrape-interval: 60s
    enabled-profiles: [cpu, mem, goroutines, mutex, block]
    static-configs:
      - application: foo
        spy-name: gospy
        targets:
          - foo:80
      - application: bar
        spy-name: gospy
        targets:
          - bar:81
```

Although not configured here, it seems advisable to set a data retention period. By default, it retains data indefinitely.
cf. [Data retention](https://pyroscope.io/docs/data-retention/)

# Sample Code
Various examples are available in the [examples](https://github.com/grafana/pyroscope/tree/main/examples) repository.

An example implemented in my managed application is as follows:

[gobel-example](https://github.com/bmf-san/gobel-example/tree/master/pyroscope)

# Challenges Faced
## Initial Password Authentication Setup
Pyroscope has mechanisms for API KEY, OAuth2, and password authentication.

I misread the documentation while setting the initial authentication information for password authentication and got stuck for a bit.

To set the initial authentication information, you need to include it in the configuration file like this:

```yml
auth:
  internal:
    admin:
      name: USERNAME
      password: PASSWORD
    enabled: true
```

Upon reading carefully, I found the instructions in [configuring-built-in-admin-user](https://pyroscope.io/docs/auth-internal/#configuring-built-in-admin-user), but I mistakenly thought it could only be changed via CLI, which caused unnecessary confusion...

## pprof
This issue is not related to Pyroscope but rather to the Go application side, where I encountered difficulties with pprof configuration.

I wrote an article about it, which you can refer to: [How to Use pprof Outside of DefaultServeMux](https://bmf-tech.com/posts/DefaultServeMux%e4%bb%a5%e5%a4%96%e3%81%a7pprof%e3%82%92%e4%bd%bf%e3%81%86%e6%96%b9%e6%b3%95)

# Thoughts
I have been searching for an OSS profiling tool for a while, and Pyroscope seems easy to implement with a user-friendly UI. The Pull mode support for Go is also a plus.