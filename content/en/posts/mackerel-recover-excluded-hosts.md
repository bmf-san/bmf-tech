---
title: How to Restore Hosts Removed from Monitoring in Mackerel
description: 'Recover accidentally excluded monitoring hosts in Mackerel by updating hostId after host migration or removal from monitoring system.'
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
This is a note on how to restore hosts that were mistakenly retired or unintentionally removed from monitoring in Mackerel.

# Procedure
Simply restarting the mackerel-agent will not automatically restore the host, so you need to update the hostId assigned to the host.

```sh
service mackerel-agent stop
cd /var/lib/mackerel-agent
mv id /tmp/
service mackerel-agent start
```

# Thoughts
I sometimes panic, so I want to be careful...
