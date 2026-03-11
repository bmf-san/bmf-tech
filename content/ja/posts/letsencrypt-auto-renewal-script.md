---
title: Letsencryptの証明書自動更新スクリプト
slug: letsencrypt-auto-renewal-script
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - cron
  - Let's Encrypt
  - shellscript
  - Slack
translation_key: letsencrypt-auto-renewal-script
---


# 概要
letsencryptの証明書を自動更新するスクリプトの紹介です。
以前から作ってはいたのですが、色々と問題があったり、サーバー環境が変わったりで完全版を残せていなかったので改めてまとめました。

# 環境
- nginx v1.12.0

※letsencryptのインストールやshell scriptの実行方法等には触れません。

# スクリプト
月に一回証明書の有効期限を問わず更新し（`--force-renew`）、更新結果（成功または失敗）をslack通知するスクリプトです。

slackの設定値は外部ファイルで管理しています。

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

# 結果
成功すると、
![スクリーンショット 2017-07-01 15.46.02.png](/assets/images/posts/letsencrypt-auto-renewal-script/90642f68-e9bc-f8cf-5164-e33897cf11dd.png)


失敗すると、
![スクリーンショット 2017-07-01 15.45.56.png](/assets/images/posts/letsencrypt-auto-renewal-script/26b2eef6-3d92-7a3b-38a0-c5ea40b3e22f.png)

成功しても失敗しても赤なのはナンセンスですね。。。

