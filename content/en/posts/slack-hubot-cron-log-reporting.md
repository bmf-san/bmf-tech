---
title: Scheduled Server Log Reports with Slack, Hubot, Cron (node-cron), Shell, and Logwatch
slug: slack-hubot-cron-log-reporting
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - cron
  - CentOS
  - Slack
  - Sakura VPS
  - CoffeeScript
  - hubot
  - logwatch
  - shell
translation_key: slack-hubot-cron-log-reporting
---

Previously, I wrote an article about executing ShellScript from Slack using Slack and Hubot.
[Creating a Slack Bot with Sakura VPS, Hubot, and Slack](http://qiita.com/BMF_Engineer/items/1f04032b05c22de062af)

This time, instead of executing commands interactively, I am attempting to run ShellScript periodically with cron and report the output results to Slack.

Surprisingly, I couldn't find many ~~copy-pasteable~~ respectable reference sources for this kind of task, so I almost gave up halfway through. 😂

While there was a safe method of hitting the API and running it via crontab, I wanted to make good use of the Hubot I had created, so I put in the effort.

**Note: At the time of writing this article, it was indeed working, but there seems to be an issue with the output text from LogWatch, which may cause it not to execute correctly. I also tried using the Slack API's Attachment API, but it was the same. The cause is still unclear..**

# Environment
* Sakura VPS  
* CentOS 6  
* Slack  
* Hubot  
* hubot-slack 3.4.2 ... This is the adapter for using Hubot with Slack. It seems that when you specify the adapter as Slack during Hubot installation, it gets added to package.json and installed automatically. For more details, see [Creating a Slack Bot with Sakura VPS, Hubot, and Slack](http://qiita.com/BMF_Engineer/items/1f04032b05c22de062af). By the way, the channel specification and the use of Attachments differ between v2 and v3, so please check that yourself.
* pm2 ... This is for daemonizing Hubot. For more details, see [Creating a Slack Bot with Sakura VPS, Hubot, and Slack](http://qiita.com/BMF_Engineer/items/1f04032b05c22de062af).
* node-cron ... This package is used to run cron jobs in CoffeeScript.
* logwatch ... It doesn't have to be logwatch. You can prepare any ShellScript that you want to output to Slack.

# Desired Knowledge
* CoffeeScript

The author seems to be quite ignorant about CoffeeScript _(:3」∠)_

# Preparation
Since we will be using the node-cron package, please install it with `npm install node-cron`.

# Source Code

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

In CoffeeScript, I specified to run cron every minute because I wanted to quickly check if it was working properly.
The time specification method for cron is the same as regular cron, so feel free to set it as you like.

Regarding the channel specification in Slack, it seems that it can be executed even with a #, but it does not execute for private channels. (Does # mean public channel in the first place??)
It seems that the channel specification differs depending on the version of hubot-slack, but I didn't encounter any issues in my environment, so those who did should check it themselves. m(_ _)m

logwatch.sh

```
# !/bin/bash
logwatch --print
```

This is a simple script that just outputs logs using logwatch.

With this, I was able to assign the task of "posting the output of `logwatch --print` to the Slack logwatch channel every minute" to Hubot.

I used to send the output via email with logwatch, but now that I can send it to Slack, the hassle has been reduced.＼(^o^)／

**Addendum:** It seems that if the text posted to Slack is too long, it gets split into multiple posts. If you notice duplicate posts, it was because the output text from logwatch was too long.

**Further Addendum:** It seems that Attachments can be used from hubot-slack v3.3.0!
[Using Attachments with hubot-slack](http://buty4649.hatenablog.com/entry/2015/12/04/011829)

# Bonus: Stopping Logwatch Email Sending
[GMO Cloud Support Guide](https://help.gmocloud.com/app/answers/detail/a_id/148/~/logwatch%E3%81%AE%E3%83%A1%E3%83%BC%E3%83%AB%E9%85%8D%E9%80%81%E3%82%92%E6%AD%A2%E3%82%81%E3%81%9F%E3%81%84%E3%80%82) provides a kind explanation.

By setting Print=Yes in the logwatch configuration file, email sending will be stopped.
I read somewhere that changing execution permissions is not very good because they might change unexpectedly, so I stopped the emails this way.

# Others
* [Monitoring URLs Every Minute with Hubot and Notifying Slack](https://blog.yug1224.com/archives/563d9b67bf652a600632d060)

# Thoughts
I dislike CoffeeScript (cry)

In the previous article, I wrote about executing ShellScript in Slack, but it seems that adding conditions like "only for specific users in a specific channel" requires some extra work.
I did a little research, but I still don't quite understand it, so I plan to tackle it eventually.