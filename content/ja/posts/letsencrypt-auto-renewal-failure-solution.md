---
title: Let'sEncryptの証明書自動更新に失敗し続けていたけど何とかした話
slug: letsencrypt-auto-renewal-failure-solution
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - cron
  - Let's Encrypt
  - Slack
  - shellscript
translation_key: letsencrypt-auto-renewal-failure-solution
---


[letsencryptの証明書自動更新(cron)とちょっとだけSlack連携](http://qiita.com/bmf_san/items/9a072023df9ca6fab354) でかいたスクリプトは、手動で実行する場合は更新成功するのですが、cronで実行すると何故か毎回失敗するという問題作でした。

そこでスクリプトを見直し、何とか正しく動作するように改良してみました。

筆者はnginx+apacheのサーバー構成です。基本的には--webrootオプションを使って証明書発行や更新を行っています。

注：letsencryptのオプションについては各自の環境に読み替えて下さい。

# スクリプト

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
    TITLE=${TITLE:-"Let's Encrypt更新エラー通知"}

    # Slack Message
    MESSAGE=${MESSAGE:-"証明書の更新に失敗しました。"}

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
        TITLE=${TITLE:-"Let's Encrypt更新完了通知"}

    # Slack Message
        MESSAGE=${MESSAGE:-"証明書を更新しました！"}

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

前回との違いは、`--force-renew`というオプションを採用したところでしょうか。証明書の残りの有効期限に関係なく更新するというものです。

それからsleepという動作を指定時間停止させる処理を追記しました。
証明書発行に時間がかかることを考慮し、slackやnginxの再起などが問題なく行なわれることよう配慮したものですが、効果の程はわかりません。。。（どこかのブログで見たので真似してみました）

# 所感
* オプションが多いけどちゃんと目を通す
* ログを見る
* 発行制限に気をつける→取得できるかのテスト用のコマンドがあった気がします

以上を注意深く行えばもう少し早く解決できたような気がします。


# 参考
[Let's Encrypt ユーザーガイド](https://letsencrypt.jp/docs/using.html)

