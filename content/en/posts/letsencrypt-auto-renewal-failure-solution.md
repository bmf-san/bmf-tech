---
title: How I Managed to Fix the Persistent Failure of Let's Encrypt Certificate Auto-Renewal
description: 'A troubleshooting guide for How I Managed to Fix the Persistent Failure of Let''s Encrypt Certificate Auto-Renewal, explaining the root cause and how to resolve it.'
slug: letsencrypt-auto-renewal-failure-solution
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - cron
  - Let's Encrypt
  - Slack
  - shell script
translation_key: letsencrypt-auto-renewal-failure-solution
---

[Script for Let's Encrypt certificate auto-renewal (cron) with a bit of Slack integration](http://qiita.com/bmf_san/items/9a072023df9ca6fab354) was successful when executed manually, but for some reason, it failed every time when run via cron.

Therefore, I reviewed the script and managed to modify it to work correctly.

The author uses an nginx+apache server configuration. Basically, I use the --webroot option for certificate issuance and renewal.

Note: Please adjust the Let's Encrypt options according to your environment.

# Script

```bash
#!/bin/sh

# WebHookUrl
WEBHOOKURL="*************************"

# Slack Channel
CHANNEL=${CHANNEL:-"#ChannelName"}

# Slack Bot Name
BOTNAME=${BOTNAME:-"BotName"}

if ! /path/to/certbot-auto renew --force-renew ; then
    sleep 15

    # Slack Title
    TITLE=${TITLE:-"Let's Encrypt Update Error Notification"}

    # Slack Message
    MESSAGE=${MESSAGE:-"Failed to update the certificate."}

    #POST
    curl -s -S -X POST --data-urlencode "payload={
             \"channel\": \"${CHANNEL}\",
             \"username\": \"${BOTNAME}\",
             \"attachments\": [{
             \"color\": \"danger\",
             \"fallback\": \"${TITLE}\",
             \"title\": \"${TITLE}\",
             \"text\": \"${MESSAGE}\"
        }]
    }" ${WEBHOOKURL} > /dev/null
else
    sleep 15

    # Slack Title
        TITLE=${TITLE:-"Let's Encrypt Update Completion Notification"}

    # Slack Message
        MESSAGE=${MESSAGE:-"Certificate updated!"}

    #POST
    curl -s -S -X POST --data-urlencode "payload={
            \"channel\": \"${CHANNEL}\",
            \"username\": \"${BOTNAME}\",
            \"attachments\": [{
            \"color\": \"danger\",
            \"fallback\": \"${TITLE}\",
            \"title\": \"${TITLE}\",
            \"text\": \"${MESSAGE}\"
        }]
    }" ${WEBHOOKURL} > /dev/null
fi
```

The difference from last time is the adoption of the `--force-renew` option. This renews the certificate regardless of the remaining validity period.

Additionally, I added a sleep command to pause the operation for a specified time. This was to consider the time it takes to issue the certificate and ensure that Slack and nginx restarts are performed without issues, although I'm not sure about its effectiveness... (I saw it on some blog and decided to imitate it)

# Thoughts
* There are many options, but make sure to read them carefully
* Check the logs
* Be cautious of issuance limits → I think there was a command for testing if it can be obtained

If these were done carefully, I feel like it could have been resolved a bit sooner.

# References
[Let's Encrypt User Guide](https://letsencrypt.jp/docs/using.html)