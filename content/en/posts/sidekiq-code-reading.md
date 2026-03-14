---
title: Code Reading of Sidekiq
description: "Understand Sidekiq architecture for Redis-based asynchronous job queuing, job enqueuing, and execution flow in Ruby applications."
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
A quick read through the Sidekiq code.

# Preparation
1. Clone sidekiq
   - https://github.com/sidekiq/sidekiq
2. Start redis
   - `docker run --name redis-server -p 6379:6379 -d redis`
3. Insert `binding.pry` where you want to debug
4. Start sidekiq
   - `bundle exec sidekiq -r ./examples/blog.rb`
5. Enqueue a job
   - `bundle exec irb -r ./examples/por.rb`

# Code Reading
## Enqueuing Jobs
1. Call asynchronous processing starting from `perform_async`
   - [sidekiq - lib/sidekiq/job.rb#L205](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/job.rb#L205)
2. Enqueue the job
   - `client_push`
     - [sidekiq - lib/sidekiq/job.rb#L368](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/job.rb#L368)
   - `push`
     - [sidekiq - lib/sidekiq/client.rb#L86](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/client.rb#L86)
   - `raw_push`
     - [sidekiq - lib/sidekiq/client.rb#L239](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/client.rb#L86)

## Executing Jobs
1. Start sidekiq
   - `run`
     - [lib/sidekiq/cli.rb#41](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/cli.rb#L41)
   - `launch`
     - [lib/sidekiq/cli.rb#116](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/cli.rb#L116)
   - `run`
     - [lib/sidekiq/cli.rb#38](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/cli.rb#L38)
2. Fetch jobs
   - `run`
     - [lib/sidekiq/processor.rb#71](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/processor.rb#L71)
   - `process_one`
     - [lib/sidekiq/processor.rb#71](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/processor.rb#L71)
   - `process`
     - [lib/sidekiq/processor.rb#159](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/processor.rb#L159)
       - Process the job
3. Execute the job
   - `execute_job`
     - [lib/sidekiq/processor.rb#185](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/processor.rb#L185)
     - [lib/sidekiq/processor.rb#217](https://github.com/sidekiq/sidekiq/blob/main/lib/sidekiq/processor.rb#L271)
