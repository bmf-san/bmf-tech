---
title: Let's Encrypt Certificate Auto-Renewal Script
slug: letsencrypt-auto-renewal-script
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - cron
  - Let's Encrypt
  - shell script
  - Slack
translation_key: letsencrypt-auto-renewal-script
---

# Overview
This post introduces a script for automatically renewing Let's Encrypt certificates. I had created it before, but due to various issues and changes in server environments, I couldn't leave a complete version, so I have summarized it again.

# Environment
- nginx v1.12.0

※ This post does not cover the installation of Let's Encrypt or how to execute shell scripts.

# Script
This script renews the certificate once a month regardless of its expiration (`--force-renew`) and sends a Slack notification of the renewal result (success or failure).

The Slack configuration values are managed in an external file.

```shell-session:letsencrypt.sh
#!/bin/sh

# Import config
. /home/bmf/scripts/conf/slack.conf

# Stop Nginx
/usr/sbin/service nginx stop

# POST
if ! /home/bmf/certbot/certbot-auto renew --force-renew ; then
  sleep 15

  # Slack Title
  TITLE=${TITLE:-"Let's Encrypt更新エラー通知"}

  # Slack Message
  MESSAGE=${MESSAGE:-"証明書の更新に失敗しました。"}

  #POST
  curl -s -S -X POST --data-urlencode "payload={
                \"channel\": \"${SL_CH_LETSENCRYPT}\",
                \"username\": \"${SL_BOTNAME}\",
                \"attachments\": [{
                \"color\": \"danger\",
                \"fallback\": \"${TITLE}\",
                \"title\": \"${TITLE}\",
                \"text\": \"${MESSAGE}\"
                }]
  }" ${SL_WEBHOOKURL} > /dev/null
else
  sleep 15

  # Slack Title
  TITLE=${TITLE:-"Let's Encrypt更新完了通知"}

  # Slack Message
  MESSAGE=${MESSAGE:-"証明書を更新しました！"}

  #POST
  curl -s -S -X POST --data-urlencode "payload={
                \"channel\": \"${SL_CH_LETSENCRYPT}\",
                \"username\": \"${SL_BOTNAME}\",
                \"attachments\": [{
                \"color\": \"danger\",
                \"fallback\": \"${TITLE}\",
                \"title\": \"${TITLE}\",
                \"text\": \"${MESSAGE}\"
                }]
  }" ${SL_WEBHOOKURL} > /dev/null
fi

# Start nginx
/usr/sbin/service nginx start

```

# Results
If successful,
![Screenshot 2017-07-01 15.46.02.png](/assets/images/posts/letsencrypt-auto-renewal-script/90642f68-e9bc-f8cf-5164-e33897cf11dd.png)

If failed,
![Screenshot 2017-07-01 15.45.56.png](/assets/images/posts/letsencrypt-auto-renewal-script/26b2eef6-3d92-7a3b-38a0-c5ea40b3e22f.png)

It's nonsensical that it's red whether it succeeds or fails...