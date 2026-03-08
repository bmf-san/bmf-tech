---
title: Scheduled Server Log Reports with Slack, hubot, cron (node-cron), shell, and logwatch
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



Previously, I wrote an article about executing ShellScript from Slack using slack and hubot.
[Creating a Slack Bot with Sakura VPS, hubot, and Slack](http://qiita.com/BMF_Engineer/items/1f04032b05c22de062af)

This time, instead of executing commands interactively, I aim to periodically run ShellScript with cron and report the output to Slack.

Surprisingly, I couldn't find any ~~copy-pasteable~~ respectable reference sources for this kind of task, so I almost gave up halfway.

There was a safe method of hitting the API and running it with crontab, but I wanted to make use of the hubot I had created, so I persevered.


**Note: At the time of writing this article, it was working, but there seems to be an issue with the output text of LogWatch, which might cause it to not execute correctly. I also tried using the Slack API's Attachment API, but it was the same. I don't really understand the cause...**

# Environment
* Sakura VPS  
* CentOS 6
* slack
* hubot
* hubot-slack 3.4.2 ... This is an adapter for using hubot with slack. When you specify the adapter as slack during hubot installation, it gets added to package.json and installed automatically. For more details, see [Creating a Slack Bot with Sakura VPS, hubot, and Slack](http://qiita.com/BMF_Engineer/items/1f04032b05c22de062af). By the way, the channel specification and the availability of Attachments differ between v2 and v3 series, so please check that on your own.
* pm2 ... For daemon persistence of hubot. For more details, see [Creating a Slack Bot with Sakura VPS, hubot, and Slack](http://qiita.com/BMF_Engineer/items/1f04032b05c22de062af).
* node-cron ... A package used to execute cron with CoffeeScript.
* logwatch ... It doesn't have to be logwatch. You can prepare any ShellScript you want to output to slack.

# Recommended Knowledge
* CoffeeScript

Note that the author seems to be ignorant about CoffeeScript _(:3」∠)_

# Preparation
Since we will use a package called node-cron, please install it with `npm install node-cron`.

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

I specified cron to run every minute because I wanted to quickly verify if cron was working properly. The time specification method for cron is the same as usual cron, so feel free to set it as you like.

Regarding the slack channel specification, it seems to execute even with a #, but it doesn't seem to execute for private channels. (Does # mean public channel in the first place??) The channel specification also seems to differ depending on the version of hubot-slack, but I didn't encounter any issues in my environment, so if you do, please investigate on your own. m(_ _)m


logwatch.sh

```
# !/bin/bash
logwatch --print
```

This is a simple script that just outputs logs with logwatch.

With this, we were able to assign the task of "posting the output result of `logwatch --print` to the Slack logwatch channel every minute" to hubot.

logwatch used to send output via email, but now that it can be sent to Slack, the hassle has been reduced \(^o^)/

Additional Note: If the text posted to Slack is too long, it seems to be split and posted. If you think it's being posted twice, it's because the output text of logwatch was too long.

Additional Note to the Additional Note: It seems that Attachments can be used from hubot-slack v3.3.0!
[Using attachments with hubot-slack](http://buty4649.hatenablog.com/entry/2015/12/04/011829)

# Bonus: Stopping logwatch Email Sending
[GMO Cloud Support Guide](https://help.gmocloud.com/app/answers/detail/a_id/148/~/logwatch%E3%81%AE%E3%83%A1%E3%83%BC%E3%83%AB%E9%85%8D%E9%80%81%E3%82%92%E6%AD%A2%E3%82%81%E3%81%9F%E3%81%84%E3%80%82) kindly explains this.

By setting Print=Yes in the logwatch configuration file, email sending is stopped. I saw an article somewhere that changing execution permissions is not very good because permissions might change unexpectedly, so I stopped emails using this method.


# Others
* [Monitoring URLs every minute with hubot and notifying Slack.](https://blog.yug1224.com/archives/563d9b67bf652a600632d060)

# Impressions
I kind of dislike CoffeeScript (cry)

In the previous article, I wrote about executing ShellScript with slack, but it seems to require some ingenuity to add conditions like "in a specific channel, only for specific users." I did a little research, but I still don't understand it well, so I plan to tackle it gradually.