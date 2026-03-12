---
title: GrafanaでContact PointsとPoliciesをプロビジョニングするようにする
description: GrafanaでContact PointsとPoliciesをプロビジョニングするようにするについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: grafana-contact-points-policies-provisioning
date: 2023-12-27T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Grafana
  - Docker
translation_key: grafana-contact-points-policies-provisioning
---


# 概要
GrafanaでContact Pointsがプロビジョニングがいつの間にかできるようになっていたので、プロビジョニングを設定してみる。

# 環境
- [grafana-oss 10.0.1](https://hub.docker.com/r/grafana/grafana-oss/tags)
- Docker
- Docker Compose

# 設定方法
雑メモなのでDocker Composeを使ったビルド周りの説明は割愛する。[bmf-san/gobel-example](https://github.com/bmf-san/gobel-example)を参照。

`provisioning/alerting`配下にcontact-points.ymlとpolilcies.ymlを追加する。ファイル名は任意で良い。

```
└── provisioning
    └──alerting
            ├── alert-rules.yml
            ├── contact-points.yml
            └── policies.yml
```

## Contact Points
下記はwebhookでSlackに通知する設定例。

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

uidはプロビジョニングする場合は自分で採番可能であるので何でもOK。

recepientはSlack通知先チャンネル名。

urlはwebhook urlを設定。tokenを使う形にもできる。

cf. [grafana.com - #provision-contact-points](https://grafana.com/docs/grafana/latest/alerting/set-up/provision-alerting-resources/file-provisioning/#provision-contact-points)

## Policies
Policiesもプロビジョニングするようにしないとデフォルトで用意されているものが利用されてしまうため、プロビジョニングしたContact Pointsの設定を反映し、Alertの通知を行うためにはPolliciesのプロビジョニングも必要になる。（他にもやりようがあるのかは詳しく
調べなかった。）

下記はSlackという名前でContact Pointsを作った場合の設定例。

```yml
apiVersion: 1

policies:
  - orgId: 1
    receiver: Slack
```

receiverはContact Points名を指定する。

設定内容は最低限のみ。

cf. [grafana.com - #provision-notification-policies](https://grafana.com/docs/grafana/latest/alerting/set-up/provision-alerting-resources/file-provisioning/#provision-notification-policies)




このような形でプロビジョニングすれば任意のContact PointsにAlertの通知が飛ぶように設定ができる。

# 所感
今まで手動で設定していたので楽になった。
