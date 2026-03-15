---
title: Setting Up Resource Monitoring with Prometheus and Grafana
description: 'Implement resource monitoring using Prometheus and Grafana with node_exporter for comprehensive infrastructure dashboards.'
slug: prometheus-grafana-resource-monitoring
date: 2018-09-15T00:00:00Z
lastmod: 2026-03-15
author: bmf-san
categories:
  - Infrastructure
tags:
  - Prometheus
  - Grafana
translation_key: prometheus-grafana-resource-monitoring
---

# Overview
I built a resource monitoring environment using Prometheus and Grafana, so here are my notes.

# Environment
- Sakura VPS
- CentOS 7 series

# Installing Prometheus
Follow the instructions on [Prometheus - Getting Started](https://prometheus.io/docs/prometheus/latest/getting_started/) to download Prometheus.

Once downloaded, start Prometheus and confirm that you can access the dashboard.

If the port is not open, make sure to open it.

# Installing node_exporter
Download node_exporter from [Prometheus - node_exporter](https://prometheus.io/download/#node_exporter).

Make sure to start it as well.

Check the port similarly.

Add the following to the `static_configs` section of `Prometheus.yml`:

```
- targets: ['localhost:9100']
```

Run `killall prometheus` and then restart Prometheus.

# Installing Grafana
Download Grafana following the instructions on [Grafana - Installing on RPM-based Linux (CentOS, Fedora, OpenSuse, RedHat)](http://docs.grafana.org/installation/debian/).
I added the Yum Repository and installed it.

After installation, start Grafana and confirm that you can access it.

Again, if the port is not open, make sure to open it.

# Integrating Prometheus and Grafana
## Login
Once Grafana is installed, access Grafana and log in first.

The initial login information is username: admin, password: admin.

You can change the login information after logging in, so adjust it as needed.

## Setting Up Data Source
There is a Windows-like icon in the left menu, click it and then click Dashboards → Home.

To set up the Data Source, select Add data source.

Refer to [Sakura Knowledge - Creating a Monitoring Dashboard by Combining Prometheus and Grafana](https://knowledge.sakura.ad.jp/12058/) for the setup method.

I mistakenly thought the URL in the HTTP settings was the default due to the placeholder design, and continued working without setting it, which resulted in the graphs not being generated properly. Don't forget to set it up.

## Preparing Dashboard Template
Prepare a dashboard template for Prometheus from Grafana Labs.

Click `Copy ID to Clipboard` on [Prometheus system by Thomas Cheronneau](https://grafana.com/dashboards/159).

Select the + icon in the left menu → Dashboards → Import.

Paste the ID into the Grafana.com Dashboard → Click Load.

Select prometheus (Prometheus Data source) in the Options Data source.

# Thoughts
I summarized this roughly, but it should allow for monitoring.
It seems that alerts can also be set up properly, so I would like to try that eventually.

# References
- [Qiita - Prometheus Environment Setup Procedure](https://qiita.com/tSU_RooT/items/fec5b9217417758988ae)