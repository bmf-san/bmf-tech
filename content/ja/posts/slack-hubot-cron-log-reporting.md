---
title: Slack＋hubot＋cron(node-cron)＋shell＋logwatchでサーバーログを定時報告
description: Slack＋hubot＋cron(node-cron)＋shell＋logwatchでサーバーログを定時報告
slug: slack-hubot-cron-log-reporting
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - cron
  - CentOS
  - Slack
  - さくらのVPS
  - CoffeScript
  - hubot
  - logwatch
  - shell
translation_key: slack-hubot-cron-log-reporting
---


以前、slackとhubotでShellScriptをSlackから実行しようという記事を書きました。
[さくらvps＋hubot＋Slackでslack botをつくる](http://qiita.com/BMF_Engineer/items/1f04032b05c22de062af)

今回は対話形式のコマンド実行ではなく、cronで定期的にShellScriptを実行、Slackに出力結果を報告しようという試みです。

意外とこの類いの~~コピペでパクれる~~リスペクトできる参考ソースが調べても出てこなかったので、途中で挫折仕掛けましたｗ

API叩いてcrontabで回すという無難な方法もありましたが、せっかくつくったhubotを活躍させたかったので頑張りました。


**注意:　この記事を執筆している時は確かに動いていたのですが、LogWatchの出力テキストに何かしらの問題があるらしく、正しく実行されない可能性があるようです。Slack APIのAttachment APIを利用した場合もやってみたのですが、同様でした。原因はよくわかっていません。。**

# 環境
* さくらVPS  
* CentOS６
* slack
* hubot
* hubot-slack3.4.2・・・hubotをslackで使うadapterです。hubotnのインストール時にadapterをslackに指定するとpackage.jsonに追加されて勝手にインストールされるようです。詳しくは[さくらvps＋hubot＋Slackでslack botをつくる](http://qiita.com/BMF_Engineer/items/1f04032b05c22de062af)。ちなみにv2とv3系でchannel指定やAttachments利用可否が異なりますが、その辺は各自調べてください。。
* pm2・・・hubotのデーモン永続化。詳しくは[さくらvps＋hubot＋Slackでslack botをつくる](http://qiita.com/BMF_Engineer/items/1f04032b05c22de062af)。
* node-cron・・・CoffeeScriptでcronを実行させるために利用するパッケージです。
* logwatch・・・logwatchでなくともかまいません。slackに出力したいShellScriptを各自用意すればOKです。

# あると望ましい知識
* CoffeeScript

なお筆者はCoffeeScriptに関して無知な模様_(:3」∠)_

# 準備
node-cronというパッケージを使用するので、`npm install node-cron`でインストールしておいてください。

# ソースコード

logwatch.coffee

```coffeescript
cron = require('cron').CronJob

module.exports = (robot) ->
  new cron '*/1 * * * *', () =>
    @exec = require('child_process').exec
    command = "/YourHubotName/scripts/shell/logwatch.sh"
    @exec command, (error, stdout, stderr) ->
      robot.send {room: "logwatch"}, stdout
  , null, true, "Asia/Tokyo"
```

CoffeeScriptは、cronがちゃんと動作するかすぐに確認したかったので1分ごとにcronを回すよう指定しています。
cronの時間指定方法は通常のcronと同じなので各自自由に設定してください。

slackのchannel指定についてですが、# をつけても実行されるようですが、private channelの場合は実行されないようです。（# はそもそもpublic　channelという意味？？）
hubot-slackのバージョンによってもchannel指定が異なるようですが、その辺は私の環境ではハマらなかったのでハマった方は各自調べてください。m(_ _)m


logwatch.sh

```
# !/bin/bash
logwatch --print
```

logwatchでログを出力するだけのカンタンなスクリプトです。


これで“一分ごとに`logwatch --print`の出力結果をSlackのlogwatchチャンネルに投稿する”というタスクをhubotに担当させることができました。


logwatchは今までメールで出力を送信していたのですが、Slackに送信できるようになったので煩わしさが減りました＼(^o^)／


追記：Slackに投稿されるテキストが長すぎると分割されて投稿されるようです。ダブって投稿されているなあ〜と思ったらlogwatchの出力テキストが長すぎたためでした。

追記の追記：Attachmentsがhubot-slack v3.3.0から使えるらしいです！
[hubot-slackでattachmentsを使う](http://buty4649.hatenablog.com/entry/2015/12/04/011829)

# おまけ logwatchのメール送信を停止する
[GMOクラウドサポートガイド](https://help.gmocloud.com/app/answers/detail/a_id/148/~/logwatch%E3%81%AE%E3%83%A1%E3%83%BC%E3%83%AB%E9%85%8D%E9%80%81%E3%82%92%E6%AD%A2%E3%82%81%E3%81%9F%E3%81%84%E3%80%82) さんが親切に説明されています。

logwatchの設定ファイルでPrint=Yesを設定するとメール送信が停止されます。
実行権限を変更するやり方だと何らかの拍子に権限が変わるからあんまり良くないという記事をどっかで見かけたので、私はこのやり方でメール停止しました。


#  その他
* [hubot で毎分URL監視をして Slack に通知する。
](https://blog.yug1224.com/archives/563d9b67bf652a600632d060)

# 所感
CoffeeScriptなんか嫌いです（泣）

slackでShellScriptを実行するという記事を前回書きましたが、“特定のチャンネルで、特定のユーザーのみ”という条件を加えるには一工夫いるようです。
少し調べてみたのですが、まだ良くわかっていないので追々手をつけていきたいと思います。

