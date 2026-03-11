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
I tried introducing a continuous profiling tool called [Pyroscope](https://pyroscope.io/).

For more about continuous profiling, refer to this article: [What is continuous profiling?](https://www.cncf.io/blog/2022/05/31/what-is-continuous-profiling/)

Apparently, Grafana acquired it earlier this year.

[Grafana Labs Acquires Pyroscope to Add Code Profiling Capability](https://devops.com/grafana-labs-acquires-pyroscope-to-add-code-profiling-capability/)

Since the acquisition, it seems the official name is Grafana Pyroscope?

There is a plugin available for Grafana, so integration is possible, but Pyroscope also has its own UI.

A [Demo](https://demo.pyroscope.io/?name=hotrod.python.frontend%7B%7D&query=) is available, which might make it easier to understand what you can see.

The code is open-source, so if you're interested in the implementation, you can check it out.

[grafana/pyroscope](https://github.com/grafana/pyroscope)

# Introduction
It's good to first look over the configuration.

[Pyroscope Agent](https://pyroscope.io/docs/agent-overview/)

## 1. Installing Pyroscope Server
An image is pushed to DockerHub, so you can use it.

[pyroscope/pyroscope](https://hub.docker.com/r/pyroscope/pyroscope)

Here's the Docker installation guide.

[Docker Guide](https://pyroscope.io/docs/docker-guide/)

There's also a Kubernetes installation guide.

[Kubernetes/Helm](https://pyroscope.io/docs/kubernetes-helm-chart/)

## 2. Enabling Profiling on the Application Side
Set up profiling and install the agent on the application side.

The basic approach is push-based, but for Go, there's a pull-based option, allowing target management on the Pyroscope server.

[Go Pull Mode](https://pyroscope.io/docs/golang-pull-mode/)

When implementing in pull mode with Go, you can manage targets like this:

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

Although not set here, it's advisable to configure data retention periods. By default, data is retained indefinitely.
cf. [Data retention](https://pyroscope.io/docs/data-retention/)

# Sample Code
Various examples are available in [examples](https://github.com/grafana/pyroscope/tree/main/examples).

Here's an example implemented in an application I manage:

[gobel-example](https://github.com/bmf-san/gobel-example/tree/master/pyroscope)

# Challenges
## Initial Password Authentication Setup
Pyroscope offers API KEY, OAuth2, and password authentication mechanisms.

I misunderstood the documentation when setting up initial authentication information with password authentication, which caused some confusion.

To set initial authentication information, include it in the configuration file like this:

```yml
auth:
  internal:
    admin:
      name: USERNAME
      password: PASSWORD
    enabled: true
```

If you read carefully, it's mentioned in [configuring-built-in-admin-user](https://pyroscope.io/docs/auth-internal/#configuring-built-in-admin-user), but I mistakenly thought it could only be changed via CLI, which led to unnecessary confusion...

## pprof
This is not about Pyroscope, but about Go applications. I encountered issues with pprof settings.

I wrote an article about it, so refer to [Using pprof with something other than DefaultServeMux](https://bmf-tech.com/posts/DefaultServeMux%e4%bb%a5%e5%a4%96%e3%81%a7pprof%e3%82%92%e4%bd%bf%e3%81%86%e6%96%b9%e6%b3%95).

# Impressions
I've been searching for a profiling tool available as OSS, and Pyroscope seems easy to implement with a user-friendly UI. It's also great that it supports pull mode for Go.