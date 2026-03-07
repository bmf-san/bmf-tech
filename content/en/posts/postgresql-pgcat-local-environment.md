---
title: Created a Local Environment for Experimenting with PostgreSQL and PgCat
slug: postgresql-pgcat-local-environment
date: 2024-09-15T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Load Testing
  - PgCat
  - PostgreSQL
  - Grafana
  - Prometheus
translation_key: postgresql-pgcat-local-environment
---

I created a local environment to experiment with PostgreSQL and PgCat.

[bmf-san/postgresql-pgcat-example](https://github.com/bmf-san/postgresql-pgcat-example)

You can now tune PostgreSQL and PgCat parameters and perform load testing locally.

The setup consists of:

- Web
  - API server using Ruby
- PostgreSQL
  - Similar to MySQL, but I got stuck for about an hour because init.sql doesn't run if data exists in the volume. It's clearly stated in the documentation, so I'll be careful next time...
- PgCat
- Prometheus
  - Collecting metrics from postgres_exporter
  - Collecting metrics from PgCat
    - PgCat has an embedded exporter
- [postgres_exporter](https://github.com/prometheus-community/postgres_exporter)
- Grafana
  - Visualizing metrics
  - It might be sufficient to just pick up the necessary ones from Prometheus
- Locust
  - A tool to write scenarios in Python for load testing
  - You can increase the number of parallel requests by scaling the worker containers
  - It's convenient and easy to use. The UI is also nice. I liked it so much that I thought I would use it in personal projects for load testing.

I wanted to handle PgBouncer as well, but I gave up halfway because the configuration became cumbersome.

I had the motivation to observe the impact of connection pooling on performance, but I haven't been able to validate it well yet.

I also think it would be nice to observe the impact of tuning PostgreSQL and PgCat parameters on performance, but I haven't been able to do that satisfactorily yet. (I ran out of energy while adjusting the environment... orz)