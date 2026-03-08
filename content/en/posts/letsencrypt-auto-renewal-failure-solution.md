---
title: How I Managed to Fix the Continuous Failures in Let's Encrypt Certificate Auto-Renewal
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

[The script for Let's Encrypt certificate auto-renewal (cron) and a bit of Slack integration](http://qiita.com/bmf_san/items/9a072023df9ca6fab354) works successfully when executed manually, but for some reason, it fails every time when run via cron.

So, I reviewed the script and tried to improve it to work correctly.

The author uses an nginx+apache server configuration. Basically, the certificate issuance and renewal are performed using the --webroot option.

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
    TITLE=${TITLE:-"Let's Encrypt Renewal Error Notification"}

    # Slack Message
    MESSAGE=${MESSAGE:-"Failed to renew the certificate."}

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
    TITLE=${TITLE:-"Let's Encrypt Renewal Complete Notification"}

    # Slack Message
    MESSAGE=${MESSAGE:-"The certificate has been renewed!"}

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

The difference from the last time is the adoption of the `--force-renew` option, which allows renewal regardless of the remaining validity period of the certificate.

Additionally, I added a sleep function to pause the process for a specified time. This was done considering that issuing a certificate can take time, ensuring that Slack and nginx restarts can proceed without issues, but I am not sure of its effectiveness... (I saw this on some blog and decided to mimic it)

# Thoughts
* There are many options, but it's important to review them carefully.
* Check the logs.
* Be careful of issuance limits → I think there was a command for testing if you can obtain it.

If I had done the above more carefully, I feel like I could have resolved it a bit faster.

# References
[Let's Encrypt User Guide](https://letsencrypt.jp/docs/using.html)