---
title: Letsencrypt Certificate Auto-Renewal Script
slug: letsencrypt-auto-renewal-script
image: /assets/images/posts/letsencrypt/26b2eef6-3d92-7a3b-38a0-c5ea40b3e22f.png
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
This post introduces a script for automatically renewing Let's Encrypt certificates. I had been working on it for a while, but due to various issues and changes in the server environment, I hadn't been able to finalize it, so I've compiled it again.

# Environment
- nginx v1.12.0

*Note: This post does not cover the installation of Let's Encrypt or how to execute shell scripts.*

# Script
This script updates the certificate once a month regardless of its expiration date (`--force-renew`) and sends the result (success or failure) as a Slack notification.

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
  TITLE=${TITLE:-"Let's Encrypt Update Error Notification"}

  # Slack Message
  MESSAGE=${MESSAGE:-"Failed to update the certificate."}

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
  TITLE=${TITLE:-"Let's Encrypt Update Complete Notification"}

  # Slack Message
  MESSAGE=${MESSAGE:-"The certificate has been updated!"}

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
![Screenshot 2017-07-01 15.46.02.png](https://qiita-image-store.s3.amazonaws.com/0/124495/90642f68-e9bc-f8cf-5164-e33897cf11dd.png)

If it fails,
![Screenshot 2017-07-01 15.45.56.png](https://qiita-image-store.s3.amazonaws.com/0/124495/26b2eef6-3d92-7a3b-38a0-c5ea40b3e22f.png)

It's nonsensical that both success and failure notifications are red...