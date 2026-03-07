---
title: How to Restore Hosts Removed from Monitoring in Mackerel
slug: mackerel-recover-excluded-hosts
date: 2019-09-17T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Mackerel
  - Tips
translation_key: mackerel-recover-excluded-hosts
---

# Overview
This is a note on how to restore hosts that were mistakenly retired from monitoring or unintentionally removed from monitoring in Mackerel.

# Steps
Since the host will not automatically return even after restarting the mackerel-agent, it is necessary to update the hostId assigned to the host.

```sh
service mackerel-agent stop
cd /var/lib/mackerel-agent
mv id /tmp/
service mackerel-agent start
```

# Thoughts
I need to be careful as it can be a bit stressful at times...