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
description: 'Set up a local PostgreSQL and PgCat connection pooler environment with Prometheus, Grafana monitoring, and Locust load testing to observe the performance impact of connection pooling.'
translation_key: postgresql-pgcat-local-environment
---



Set up a local environment to experiment with PostgreSQL and PgCat.

[bmf-san/postgresql-pgcat-example](https://github.com/bmf-san/postgresql-pgcat-example)

You can tune parameters of PostgreSQL and PgCat and perform load testing locally.

The setup includes:

- Web
  - A simple API server using Ruby
- PostgreSQL
  - Similar to MySQL, there's a trap where init.sql doesn't run if data exists in the volume. I got stuck on this for a while. It's clearly mentioned in the documentation, so I'll be careful next time...
- PgCat
- Prometheus
  - Collecting metrics with postgres_exporter
  - Collecting metrics for PgCat
    - PgCat includes an exporter
- [postgres_exporter](https://github.com/prometheus-community/postgres_exporter)
- Grafana
  - Visualizing metrics
  - Picking only the necessary ones with Prometheus might be sufficient
- Locust
  - A tool for writing scenarios in Python to perform load testing
  - You can increase the number of parallel requests by scaling the Worker container
  - Convenient and easy to use. The UI is nice. If you're doing load testing with OSS, this seems good, and I liked it, so I thought of using it for personal projects as well

I wanted to also handle PgBouncer, but the configuration became cumbersome, and I gave up halfway.

I was motivated by the desire to observe the impact of connection pooling on performance effectively, but I haven't been able to verify it satisfactorily yet.

I also think it would be good to observe the impact of parameter tuning of PostgreSQL and PgCat on performance, but I haven't been able to do it satisfactorily yet. (I got exhausted while adjusting the environment...orz)



