---
title: Reading Sidekiq Code
slug: sidekiq-code-reading
date: 2024-09-21T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Ruby
  - Sidekiq
translation_key: sidekiq-code-reading
---

# Overview
Reading through the Sidekiq code lightly.

# Preparation
1. Clone Sidekiq
  - https://github.com/sidekiq/sidekiq
2. Start Redis
  - `docker run --name redis-server -p 6379:6379 -d redis`
3. Insert `binding.pry` where you want to debug
4. Start Sidekiq
  - `bundle exec sidekiq -r ./examples/blog.rb`
5. Push a job
  - `bundle exec irb -r ./examples/por.rb`

# Code Reading
## Job Submission
1. Call asynchronous processing starting from `perform_async`
  - [sidekiq - lib/sidekiq/job.rb#L205](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/job.rb#L205)
2. Push the job to the queue
  - `client_push`
    - [sidekiq - lib/sidekiq/job.rb#L368](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/job.rb#L368)
  - `push`
    - [sidekiq - lib/sidekiq/client.rb#L86](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/client.rb#L86)
  - `raw_push`
    - [sidekiq - lib/sidekiq/client.rb#L239](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/client.rb#L86)

## Job Execution
1. Starting Sidekiq
  - `run`
    - [lib/sidekiq/cli.rb#41](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/cli.rb#L41)
  - `launch`
    - [lib/sidekiq/cli.rb#116](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/cli.rb#L116)
  - `run`
    - [lib/sidekiq/cli.rb#38](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/cli.rb#L38)
2. Fetching the job
  - `run`
    - [lib/sidekiq/processor.rb#71](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/processor.rb#L71)
  - `process_one`
    - [lib/sidekiq/processor.rb#71](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/processor.rb#L71)
  - `process`
    - [lib/sidekiq/processor.rb#159](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/processor.rb#L159)
      - Process the job
3. Executing the job
  - `execute_job`
    - [lib/sidekiq/processor.rb#185](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/processor.rb#L185)
    - [lib/sidekiq/processor.rb#217](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/processor.rb#L271)