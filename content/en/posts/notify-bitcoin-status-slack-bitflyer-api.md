---
title: Notify Bitcoin Asset Status on Slack Using bitFlyer's Private API
description: 'Learn bitFlyer API integration with Node.js, HMAC-SHA256 signing, and Slack webhook message formatting.'
slug: notify-bitcoin-status-slack-bitflyer-api
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - API
  - Bitcoin
  - Node.js
  - bitFlyer
translation_key: notify-bitcoin-status-slack-bitflyer-api
---


# Background
I wanted to try out the API provided by bitFlyer, so I decided to call the API that returns asset status.

# Source
We will call the API using Node.js. The response data is too embarrassing to show.////

Most of the source code is based on the examples in the documentation.

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

Feel free to format the JSON response as you like. (Lazy)

# GitHub
The source code is available here: [bmf-san/bitflyer-private-api-and-slack-api-sample](https://github.com/bmf-san/bitflyer-private-api-and-slack-api-sample)

# Thoughts
I would like to build a real-time application by combining it with WebSocket, but I just can't get myself to implement WebSocket...