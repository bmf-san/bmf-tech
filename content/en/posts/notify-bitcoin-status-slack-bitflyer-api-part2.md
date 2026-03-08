---
title: Using bitflyer's private API to notify Bitcoin asset status on Slack Part 2
slug: notify-bitcoin-status-slack-bitflyer-api-part2
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - bitflyer
  - API
  - async
  - bitcoin
  - Node.js
translation_key: notify-bitcoin-status-slack-bitflyer-api-part2
---

[Previous post](http://qiita.com/bmf_san/items/7ae9fc2c83d563291671) covered just hitting the bitflyer API, so this time we'll format the response data to send asset information to Slack.

# Asynchronous Processing with Async

Since we needed to hit multiple APIs, we used async. The error handling part is inspired by a reference site (which I forgot). The code ended up looking a bit convoluted...

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
                "fallback": "bitflyer asset information",
                "color": "danger",
                "title": "bitflyer Asset Information",
                "text": "Currently holding BTC: " + btc_available + "",
                "fields": [{
                    "title": "Total BTC Assets",
                    "value": total_assets,
                    "short": true
                }, {
                    "title": "JPY (Balance)",
                    "value": jpy_available + ' yen',
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

# Questions

The total assets are calculated as `owned Bitcoin × recent transaction price (yen)`, but is there no way to get total assets in one go?

# Issues

Due to sloppy calculations of the fractional part, there are discrepancies. I want to address that properly....

# Source
+ [bitflyer-private-api-and-slack-api-sample](https://github.com/bmf-san/bitflyer-private-api-and-slack-api-sample)