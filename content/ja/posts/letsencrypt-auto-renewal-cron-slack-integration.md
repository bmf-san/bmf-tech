---
title: letsencryptの証明書自動更新(cron)とちょっとだけSlack連携
description: letsencryptの証明書自動更新(cron)とちょっとだけSlack連携
slug: letsencrypt-auto-renewal-cron-slack-integration
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - cron
  - Let's Encrypt
  - shellscript
  - Slack
translation_key: letsencrypt-auto-renewal-cron-slack-integration
---


# letsencryptの証明書自動更新(cron)とちょっとだけSlack連携　


# やること
letsencryptの証明書を自動更新させるのを長らく忘れていたのでshellとcronで設定します。


# やらないこと
* letsencryptのインストール・・・pythonエラーがちょっと面倒だった記憶が。。。
* letsencryptのオブションの説明・・・色々あるので。。。

# 環境
* さくらvps
* CentOS6系
* letsencrypt


# ShellScript
shellをかきます。
shellの保存場所は適宜設けてください。


```
#!/bin/sh
service nginx stop
/root/letsencrypt/letsencrypt-auto certonly --standalone --renew-by-default -d DOMAIN_NAME
service nginx start
```

/letencrypt-autoまでのパスは適宜指定。
オプションについて同様です。（これが結構面倒な気がします。。。）


# cron登録
`crontab -e`でcron登録します。


# 注意：証明書の発行数制限
cronの実行確認のテストで証明書更新しまくっていたら、「更新リクエスト多すぎやめちくり〜」というエラーがでました。


[Let's Encrypt - Rate Limits](https://letsencrypt.org/docs/rate-limits/)


**「制限に引っかかったら月曜日まで待たないといけないようです。**

**週に20個の証明書**までOKだそうなので、それを超えないようにテストしましょう。



# crontabでちょっと気になった記事
[crontab -eは「絶対に」使ってはいけない](http://d.hatena.ne.jp/ozuma/20120711/1342014448)


# 所感
更新には時間がかからないようですが、nginxを一度停止する必要があるのでサービス環境によってその辺り考慮する必要はある気がします。
作成したshellにSlack APIたたいて更新通知投げるのもいいと思いました。
ただエラーをどうやってキャッチすればいいかわからないです＿|￣|○

あった→[Let's Encryptで自動更新したらSlackに通知したかった](http://qiita.com/kamemory/items/cfa6d9511ce831e90116)

普通にif文でかけばいいみたいです。


というわけでやっつけでshellをつくりました。
更新完了か失敗か通知を送るだけです。更新日時等の出力はしません。

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
        TITLE=${TITLE:-"Let's Encrypt更新エラー通知"}

        # Slack Message
        MESSAGE=${MESSAGE:-"証明書の更新に失敗しました。"}

        #POST
        curl -s -S -X POST --data-urlencode "payload={
                \\"channel\\": \\"${CHANNEL}\\",
                \\"username\\": \\"${BOTNAME}\\",
                \\"attachments\\": [{
                \\"color\\": \\"danger\\",
                \\"fallback\\": \\"${TITLE}\\",
                \\"title\\": \\"${TITLE}\\",
                \\"text\\": \\"${MESSAGE}\\"
                }]
        }" ${WEBHOOKURL} >/dev/null
else
        # Slack Title
        TITLE=${TITLE:-"Let's Encrypt更新完了通知"}

        # Slack Message
        MESSAGE=${MESSAGE:-"証明書の更新が完了しました！"}

        #POST
        curl -s -S -X POST --data-urlencode "payload={
                \\"channel\\": \\"${CHANNEL}\\",
                \\"username\\": \\"${BOTNAME}\\",
                \\"attachments\\": [{
                \\"color\\": \\"danger\\",
                \\"fallback\\": \\"${TITLE}\\",
                \\"title\\": \\"${TITLE}\\",
                \\"text\\": \\"${MESSAGE}\\"
                }]
        }" ${WEBHOOKURL} >/dev/null
fi

# Restart Nginx
service nginx start
```

# 所感
letsencryptありがたや


# 参考
* [Let's Encrypt で証明書を取得した時の手順備忘録](http://qiita.com/TsutomuNakamura/items/4166423699061e38d296)

