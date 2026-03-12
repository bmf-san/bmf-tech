---
title: Sidekiqのコードリーディング
description: Sidekiqのコードリーディングについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: sidekiq-code-reading
date: 2024-09-21T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Ruby
  - Sidekiq
translation_key: sidekiq-code-reading
---


# 概要
Sidekiqのコードをさらっと読んでみる。

# 準備
1. sidekiqをクローン
  - https://github.com/sidekiq/sidekiq
2. redisを起動
  - `docker run --name redis-server -p 6379:6379 -d redis`
3. デバッグしたいところでbinding.pryを仕込む
4. sidekiqを起動
  - `bundle exec sidekiq -r ./examples/blog.rb`
5. ジョブを投入
  - `bundle exec irb -r ./examples/por.rb`

# コードリーディング
## ジョブの投入
1. `perform_async`を起点に非同期処理を呼び出す
  - [sidekiq - lib/sidekiq/job.rb#L205](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/job.rb#L205)
2. ジョブをキューに投入する
  - `client_push`
    - [sidekiq - lib/sidekiq/job.rb#L368](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/job.rb#L368)
  - `push`
    - [sidekiq - lib/sidekiq/client.rb#L86](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/client.rb#L86)
  - `raw_push`
    - [sidekiq - lib/sidekiq/client.rb#L239](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/client.rb#L86)

## ジョブの実行
1. sidekiqの起動
  - `run`
    - [lib/sidekiq/cli.rb#41](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/cli.rb#L41)
  - `launch`
    - [lib/sidekiq/cli.rb#116](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/cli.rb#L116)
  - `run`
    - [lib/sidekiq/cli.rb#38](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/cli.rb#L38)
2. ジョブの取得
  - `run`
    - [lib/sidekiq/processor.rb#71](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/processor.rb#L71)
  - `process_one`
    - [lib/sidekiq/processor.rb#71](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/processor.rb#L71)
  - `process`
    - [lib/sidekiq/processor.rb#159](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/processor.rb#L159)
      - jobを処理する
3. ジョブの実行
  - `execute_job`
    - [lib/sidekiq/processor.rb#185](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/processor.rb#L185)
    - [lib/sidekiq/processor.rb#217](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/processor.rb#L271)
