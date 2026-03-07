---
title: Setting Up a Resource Monitoring Environment with Prometheus and Grafana
slug: prometheus-grafana-resource-monitoring
date: 2018-09-15T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Prometheus
  - Grafana
description: A guide to building a resource monitoring environment using Prometheus and Grafana.
translation_key: prometheus-grafana-resource-monitoring
---

# Overview
This is a memo on setting up a resource monitoring environment using Prometheus and Grafana.

# Environment
- Sakura VPS
- CentOS 7 series

# Installing Prometheus
Download Prometheus following the instructions on [Prometheus - Getting Started](https://prometheus.io/docs/prometheus/latest/getting_started/).

After downloading, start Prometheus and confirm that you can access the dashboard.

If the port is not open, make sure to open it.

# Installing node_exporter
Download node_exporter from [Prometheus - node_exporter](https://prometheus.io/download/#node_exporter).

Start node_exporter as well.

Similarly, check the port.

Add the following to the `static_configs` section in `Prometheus.yml`:

```
- targets: ['localhost:9100']
```

Run `killall prometheus` and then restart Prometheus.

# Installing Grafana
Download Grafana following the instructions on [Grafana - Installing on RPM-based Linux (CentOS, Fedora, OpenSuse, RedHat)](http://docs.grafana.org/installation/debian/).

Add the Yum Repository and install it.

After installation, start Grafana and confirm that you can access it.

If the port is not open, make sure to open it.

# Integrating Prometheus and Grafana
## Login
Once Grafana is installed, access Grafana and log in.

The default login credentials are username: `admin` and password: `admin`.

After logging in, you can change the login credentials as needed.

## Setting the Data Source
Click the Windows-like icon in the left menu, then click Dashboards → Home.

To set up the Data Source, select Add data source.

Refer to [Sakura Knowledge - Creating Monitoring Dashboards with Prometheus and Grafana](https://knowledge.sakura.ad.jp/12058/) for configuration details.

Due to the placeholder design in the HTTP settings URL, I mistakenly thought the default value was pre-configured. If you proceed without setting it, the graphs will not generate properly. Don’t forget to configure it.

## Preparing a Dashboard Template
Prepare a dashboard template for Prometheus from Grafana Labs.

Use [Prometheus system by Thomas Cheronneau](https://grafana.com/dashboards/159) and click `Copy ID to Clipboard`.

In the left menu, click the + icon → Dashboards → Import.

Paste the ID into the Grafana.com Dashboard field and click Load.

In the Options section, select `prometheus` (the Prometheus Data Source) for the Data source.

# Impressions
This is a rough summary, but this setup should allow for monitoring. It seems you can also configure alerts, so I’d like to try that at some point.

# References
- [Qiita - Prometheus Setup Guide](https://qiita.com/tSU_RooT/items/fec5b9217417758988ae)
