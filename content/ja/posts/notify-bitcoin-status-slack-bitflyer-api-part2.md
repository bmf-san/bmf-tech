---
title: "bitflyerのprivate apiを使ってビットコイン資産状況をslackに通知する Part2"
slug: "notify-bitcoin-status-slack-bitflyer-api-part2"
date: 2017-09-26
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "bitflyer"
  - "API"
  - "async"
  - "bitcoin"
  - "Node.js"
draft: false
---

[前回](http://qiita.com/bmf_san/items/7ae9fc2c83d563291671)、bitflyerのAPIを叩くだけ叩いてポイーしてたので、今回はレスポンスデータを整形して資産情報データとしてslackに投げれるようにします。

# とりあえずasyncで非同期処理

APIを複数叩く必要があったのでasyncを使いました。
エラー拾っているところは参考サイト（失念しました）を真似ています。
なんだか見通しの悪いコードになってしまいました・・・

```js
var request = require('request');
var crypto = require('crypto');
var async = require('async');

var key = 'YOUR_KEY';
var secret = 'YOUR_SECRET';

var timestamp = Date.now().toString();
var sign = crypto.createHmac('sha256', secret).update(timestamp + 'GET' + '/v1/me/getbalance').digest('hex');

var requests = [{
    url: 'https://api.bitflyer.jp/v1/me/getbalance',
    headers: {
        'ACCESS-KEY': key,
        'ACCESS-TIMESTAMP': timestamp,
        'ACCESS-SIGN': sign,
        'Content-Type': 'application/json'
    }
}, {
    url: 'https://api.bitflyer.jp/v1/getexecutions'
}];

async.map(requests, function(obj, callback) {
    request(obj, function(error, response, body) {
        if (!error && response.statusCode == 200) {
            var body = JSON.parse(body);
            callback(null, body);
        } else {
            callback(error || response.statusCode);
        }
    });
}, function(err, results) {
    if (err) {
        console.log(err);
    } else {
        var jpy_available = parseInt(results[0][0]['available']);
        var btc_available = results[0][1]['available'];
        var price = parseInt(results[1][1]['price']);
        var total_assets = Math.floor(btc_available*price);

        var data = JSON.stringify({
            "attachments": [{
                "fallback": "bitflyer資産情報",
                "color": "danger",
                "title": "bitflyer資産情報",
                "text": "現在保有しているBTCは" + btc_available + "です",
                "fields": [{
                    "title": "BTC総資産",
                    "value": total_assets,
                    "short": true
                }, {
                    "title": "JPY(残高)",
                    "value": jpy_available + '円',
                    "short": true
                }]
            }],
            "username": "your-bot-name",
            "icon_url": "/path/to/img",
            "channel": "#yourslackchannel"
        });

        var options = {
            url: 'https://hooks.slack.com/services/WEBHOOK_TOKEN',
            form: 'payload=' + data,
            json: true
        };

        request.post(options, function(error, response, body) {
            if (!error && response.statusCode == 200) {
                console.log(body.name);
            } else {
                console.log('error: ' + response.statusCode + body);
            }
        });
    }
});
```

# 疑問
総資産は`所有するビットコイン×直近の取引価格（円）`で計算しているのですが、総資産って一発データで取れないんですかね・・？

# 課題
浮動小数部分の数値の計算が雑すぎるためか、誤差が生じていますｗ
なのでその辺しっかりやりたいです。。。。

# ソース
+ [bitflyer-private-api-and-slack-api-sample](https://github.com/bmf-san/bitflyer-private-api-and-slack-api-sample)

