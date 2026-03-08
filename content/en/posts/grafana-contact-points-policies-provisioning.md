---
title: Provisioning Contact Points and Policies in Grafana
slug: grafana-contact-points-policies-provisioning
date: 2023-12-27T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Grafana
  - Docker
translation_key: grafana-contact-points-policies-provisioning
---

# Overview
Grafana has recently enabled provisioning of Contact Points, so I will try to set up provisioning.

# Environment
- [grafana-oss 10.0.1](https://hub.docker.com/r/grafana/grafana-oss/tags)
- Docker
- Docker Compose

# Setup Method
This is a rough note, so I will skip the explanation of the build process using Docker Compose. Please refer to [bmf-san/gobel-example](https://github.com/bmf-san/gobel-example).

Add contact-points.yml and policies.yml under `provisioning/alerting`. The file names can be arbitrary.

```
└── provisioning
    └── alerting
            ├── alert-rules.yml
            ├── contact-points.yml
            └── policies.yml
```

## Contact Points
Below is an example configuration for notifying Slack via webhook.

```yml
apiVersion: 1

contactPoints:
  - orgId: 1
    name: Slack
    receivers:
      - uid: abc1234
        type: slack
        settings:
          recepient: alert-slack-channel-name
          url: [webhook url]
        disableResolveMessage: false
```

The uid can be any number you assign when provisioning.

recepient is the name of the Slack notification channel.

Set the url to the webhook url. It can also be configured to use a token.

cf. [grafana.com - #provision-contact-points](https://grafana.com/docs/grafana/latest/alerting/set-up/provision-alerting-resources/file-provisioning/#provision-contact-points)

## Policies
If you do not provision Policies, the default ones will be used, so to reflect the settings of the provisioned Contact Points and send Alert notifications, provisioning of Policies is also necessary. (I did not investigate if there are other methods.)

Below is an example configuration when a Contact Point named Slack is created.

```yml
apiVersion: 1

policies:
  - orgId: 1
    receiver: Slack
```

receiver specifies the name of the Contact Point.

The configuration is minimal.

cf. [grafana.com - #provision-notification-policies](https://grafana.com/docs/grafana/latest/alerting/set-up/provision-alerting-resources/file-provisioning/#provision-notification-policies)

By provisioning in this way, you can set up alerts to be sent to any Contact Points.

# Thoughts
It has become easier since I used to set it up manually.