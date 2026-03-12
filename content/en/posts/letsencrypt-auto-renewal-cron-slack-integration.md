---
title: Automating Let's Encrypt Certificate Renewal (cron) with a Touch of Slack Integration
description: 'An in-depth look at Automating Let''s Encrypt Certificate Renewal (cron) with a Touch of Slack Integration, covering key concepts and practical insights.'
slug: letsencrypt-auto-renewal-cron-slack-integration
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - cron
  - Let's Encrypt
  - shellscript
  - Slack
translation_key: letsencrypt-auto-renewal-cron-slack-integration
---

# Automating Let's Encrypt Certificate Renewal (cron) with a Touch of Slack Integration

# What We'll Do
I've long forgotten to automate the renewal of Let's Encrypt certificates, so I'll set it up using shell and cron.

# What We Won't Do
* Installing Let's Encrypt... I remember the Python errors being a bit troublesome...
* Explaining Let's Encrypt options... there are many...

# Environment
* Sakura VPS
* CentOS 6 series
* Let's Encrypt

# ShellScript
Let's write a shell script. Please set an appropriate location to save the shell script.

```
#!/bin/sh
service nginx stop
/root/letsencrypt/letsencrypt-auto certonly --standalone --renew-by-default -d DOMAIN_NAME
service nginx start
```

Specify the path up to /letsencrypt-auto as needed. The same goes for options. (This seems quite troublesome...)

# Registering cron
Register cron using `crontab -e`.

# Note: Certificate Issuance Limits
When testing cron execution, I kept updating the certificate, and I got an error saying "Too many renewal requests, please stop~".

[Let's Encrypt - Rate Limits](https://letsencrypt.org/docs/rate-limits/)

**It seems you have to wait until Monday if you hit the limit.**

**Up to 20 certificates per week** are allowed, so let's test without exceeding that.

# An Article About crontab That Caught My Attention
[You Should "Never" Use crontab -e](http://d.hatena.ne.jp/ozuma/20120711/1342014448)

# Thoughts
It doesn't seem to take much time to update, but since you need to stop nginx once, it might be necessary to consider this depending on the service environment. I thought it would be nice to hit the Slack API in the shell I created to send update notifications. However, I'm not sure how to catch errors ＿|￣|○

Found it → [I Wanted to Notify Slack When Automatically Updating with Let's Encrypt](http://qiita.com/kamemory/items/cfa6d9511ce831e90116)

It seems you can just write an if statement.

So, I hastily created a shell script. It only sends notifications of completion or failure of updates. It does not output update dates, etc.

letsencrypt.sh

```
#!/bin/sh

# Stop Nginx
service nginx stop

# WebHookUrl
WEBHOOKURL="SLACK_WEBHOOK_URL"

# Slack Channel
CHANNEL=${CHANNEL:-"#letsencrypt"}

# Slack Bot Name
BOTNAME=${BOTNAME:-"ssl-bot"}


if ! /root/letsencrypt/letsencrypt-auto certonly --standalone --renew-by-default -d DOMAIN_NAME ; then
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
        }" ${WEBHOOKURL} >/dev/null
else
        # Slack Title
        TITLE=${TITLE:-"Let's Encrypt Update Completion Notification"}

        # Slack Message
        MESSAGE=${MESSAGE:-"Certificate update completed!"}

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
        }" ${WEBHOOKURL} >/dev/null
fi

# Restart Nginx
service nginx start
```

# Thoughts
Thank you, Let's Encrypt

# References
* [Memo on the Procedure When Obtaining a Certificate with Let's Encrypt](http://qiita.com/TsutomuNakamura/items/4166423699061e38d296)
