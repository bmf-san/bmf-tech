---
title: bitflyerのprivate apiを使ってビットコイン資産状況をslackに通知する
slug: notify-bitcoin-status-slack-bitflyer-api
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - API
  - bitcoin
  - Node.js
  - bitflyer
translation_key: notify-bitcoin-status-slack-bitflyer-api
---


# 背景
bitflyerが用意しているAPIをちょっと使ってみたかったので資産状況を返すAPIをたたいてみました。

# ソース
nodejsでAPIをたたきます。
レスポンスデータは恥ずかしいので見せません。////

なおソースコードの大部分はドキュメントの例を参考にしています（）

```bitflyer.js
var request = require('request');
var crypto = require('crypto');

var key = 'your_bitflyer_api';
var secret = 'your_bitflyer_secret';

var timestamp = Date.now().toString();
var method = 'GET';
var path = '/v1/me/getbalance';

var text = timestamp + method + path;
var sign = crypto.createHmac('sha256', secret).update(text).digest('hex');

var options = {
        url: 'https://api.bitflyer.jp' + path,
        method: method,
        headers: {
        'ACCESS-KEY': key,
        'ACCESS-TIMESTAMP': timestamp,
        'ACCESS-SIGN': sign,
        'Content-Type': 'application/json'
    }
};

request(options, function (err, response, payload) {
    var hoge = 'hoge';

    var data = JSON.stringify({"text": payload, "username": "your_bot_name", "icon_url": "your_icon_url","channel": "#channel_name"});

    var options = {
        url: 'your_slack_webhookurl',
        form: 'payload=' + data,
        json: true
    };

    request.post(options, function(error, response, body){
        if (!error && response.statusCode == 200) {
            console.log(body.name);
        } else {
            console.log('error: '+ response.statusCode + body);
        }
    });
});
```

jsonレスポンスは好きなように整形してください。（怠惰）

# github
ソースおいてあります。[bmf-san/bitflyer-private-api-and-slack-api-sample](https://github.com/bmf-san/bitflyer-private-api-and-slack-api-sample)

# 所感
WebSocketとか組み合わせてリアルタイムなアプリケーションを構築してみたいのですが、WebSocketの実装がどうにも腰が上がりません。。。

