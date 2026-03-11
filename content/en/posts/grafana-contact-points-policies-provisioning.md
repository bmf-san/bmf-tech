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
It has become possible to provision Contact Points in Grafana, so I tried setting up provisioning.

# Environment
- [grafana-oss 10.0.1](https://hub.docker.com/r/grafana/grafana-oss/tags)
- Docker
- Docker Compose

# Configuration Method
This is a rough note, so I will omit the explanation around building with Docker Compose. Refer to [bmf-san/gobel-example](https://github.com/bmf-san/gobel-example).

Add `contact-points.yml` and `policies.yml` under `provisioning/alerting`. The file names can be arbitrary.

```
└── provisioning
    └──alerting
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

You can assign any uid when provisioning.

`recepient` is the Slack notification channel name.

`url` is set to the webhook url. It can also be configured to use a token.

cf. [grafana.com - #provision-contact-points](https://grafana.com/docs/grafana/latest/alerting/set-up/provision-alerting-resources/file-provisioning/#provision-contact-points)

## Policies
Provisioning Policies is also necessary to reflect the settings of provisioned Contact Points and send Alert notifications, otherwise, the default ones will be used. (I didn't investigate if there are other ways.)

Below is an example configuration when creating a Contact Point named Slack.

```yml
apiVersion: 1

policies:
  - orgId: 1
    receiver: Slack
```

`receiver` specifies the Contact Points name.

The configuration is minimal.

cf. [grafana.com - #provision-notification-policies](https://grafana.com/docs/grafana/latest/alerting/set-up/provision-alerting-resources/file-provisioning/#provision-notification-policies)



By provisioning in this way, you can configure alerts to be sent to any Contact Points.

# Impressions
It has become easier since I used to configure it manually.