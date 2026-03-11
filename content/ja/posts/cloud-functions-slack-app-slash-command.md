---
title: Cloud Functionsを使ってSlack AppのSlash Commandを実装
slug: cloud-functions-slack-app-slash-command
image: /assets/images/posts/cloud-functions-slack-app-slash-command/188304723-637b0b8a-6253-45db-86c9-17b33131444b.png
date: 2022-09-19T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - slack-bot
  - Slack
  - Golang
  - Google Cloud Platform
  - Cloud Functions
translation_key: cloud-functions-slack-app-slash-command
---


# 概要
Cloud Functionsを使ってSlack AppのSlash Commandを実装する。

今回作ったボイラープレートはこちら。

[go-slack-app-on-cloud-functions-boilerplate](https://github.com/bmf-san/go-slack-app-on-cloud-functions-boilerplate)

Slash Commandの使えるSlack Appを作る方法は色々あるが、安く、簡単に、サーバーレスで作れるということでCloudFront Functionsを使ってみた。

# 前提
- Google Cloud Platformが利用できること
- gcloudコマンドのセットアップが完了していること
  - アプリケーションをデプロイするためにgcloudコマンドが利用できる必要がある
- Cloud Build APIが有効化されていること
  - Cloud Functionsに関数をデプロイするために、関数をビルドするのだが、そのために必要。

# Cloud Functionsで関数を作成
Cloud Functionsのコンソールにて、関数を作成しておく。
トリガータイプは**HTTP**、**認証は未認証の呼び出しを許可**を選択、**HTTPSが必須**をチェックする。

あとで関数をデプロイした後に、Cloud Functionsの関数の詳細 > トリガーに記載されているトリガーURLを使うので、メモしておく。

# Slack Appの準備
## Slack Appを作成
[Create an app](https://api.slack.com/apps?new_app=1)にて、**From scratch**を押下する。

<img width="721" alt="スクリーンショット 2022-09-04 17 29 50" src="/assets/images/posts/cloud-functions-slack-app-slash-command/188304723-637b0b8a-6253-45db-86c9-17b33131444b.png">

**App Name**を入力する。

**Pick a workspace to develop your app in:** にてワークスペースを選択して、
**Create App**を押下する。

<img width="714" alt="スクリーンショット 2022-09-04 17 32 07" src="/assets/images/posts/cloud-functions-slack-app-slash-command/188304790-225cfb1c-2a31-4627-8f57-35b856b4aed8.png">

## Slash Commandを設定
設定画面（ex. https://api.slack.com/app/****) にて、Slash Commandsを選択する。

<img width="720" alt="スクリーンショット 2022-09-04 17 33 21" src="/assets/images/posts/cloud-functions-slack-app-slash-command/188304841-db000433-2a20-4e4a-b303-9a2fac7e3e7b.png">

**Create New Command**を押下し、**Command**、**Short Description**、**Usage Hint**、**Escape channels, users, and links sent to your app**を任意で入力する。

**Request URL**は先程メモしたトリガーURLを入力する。
トリガーURLは**https://REGION-NAME-PROJECT-ID.cloudfunctions.net/FUNCTION_NAME**という形式になっている。

<img width="568" alt="スクリーンショット 2022-09-04 17 35 59" src="/assets/images/posts/cloud-functions-slack-app-slash-command/188304929-a17ccbf4-3194-490e-ad65-12c77c5f324a.png">

入力できたら**Save**を押下する。

## Slack Appをインストールする
設定画面（ex. https://api.slack.com/apps/****） にて、**Install App**を押下する。

<img width="738" alt="スクリーンショット 2022-09-04 17 37 16" src="/assets/images/posts/cloud-functions-slack-app-slash-command/188304972-f35057c1-7392-429e-90f0-b1ec02e096b0.png">

**Install to workspace**を押下して、任意のワークスペースにAppwoインストールする。

## Signing Secretを取得
設定画面（ex. https://api.slack.com/apps/****） にて、**Basic Infomation**を押下する。

App Credentialsという項目に**Signing Secret**があるので、値をメモしておく。

# 関数を実装
Cloud Functionsにデプロイする関数を実装する。

若干ハマりどころ（**go mod vendor**する部分とか）はあったりするが、実装詳細は割愛する。

ソースコードは以下参照。

[go-slack-app-on-cloud-functions-boilerplate](https://github.com/bmf-san/go-slack-app-on-cloud-functions-boilerplate)

# 関数をデプロイ
[go-slack-app-on-cloud-functions-boilerplate](https://github.com/bmf-san/go-slack-app-on-cloud-functions-boilerplate)のREADMEに従って、環境変数の準備とデプロイを実施する。

# 動作確認
作成したSlash CommandをSlackで使ってみる。

ex. 
**/hello Bob**

<img width="489" alt="スクリーンショット 2022-09-04 17 47 11" src="/assets/images/posts/cloud-functions-slack-app-slash-command/188305315-5fe063c2-971b-4b18-978a-719596c2bd87.png">

# 所感
Slack Appを作る部分をコード化できると嬉しい。

# 追記
仕事で勤怠管理のためのSlack Commandを作りたかったので[go-slack-app-on-cloud-functions-boilerplate](https://github.com/bmf-san/go-slack-app-on-cloud-functions-boilerplate)をベースに、[akashi-slack-slash-command](https://github.com/bmf-san/akashi-slack-slash-command)というのを作ってみた。

この実装では、ストレージをSpreadsheetにしているが、Google Workspaceを使っていると権限周りでSpreadsheetの共有設定が柔軟に調整できないという管理上の問題があって、職場ではストレージをSpreadsheetからCloud Storageに差し替えて実装を調整して使っている。

勤怠管理でAkashi、チャットツールでSlackを使っている組織があれば簡単に利用できるSlack Commandになっていると思う。

運用コストも大してかからず、スケーラビリティはちょっと怪しいかも。

数千人超えるような組織でなければ多分問題なく使えると思う。多分。
