---
title: Stuck with ReadOnlyTransaction in Go's Spanner Client
description: 'Fix session pool exhaustion caused by missing tx.Close() on Cloud Spanner ReadOnlyTransaction in Go. Learn proper transaction lifecycle management and issue detection with zagane.'
slug: go-spanner-client-read-only-transaction-issue
date: 2021-02-08T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Google Cloud Platform
  - Spanner
  - Tips
translation_key: go-spanner-client-read-only-transaction-issue
---

# Overview
I made some notes about the issues I encountered while using `ReadOnlyTransaction` from [pkg.go.dev - cloud.google.com/go/spanner](https://pkg.go.dev/cloud.google.com/go/spanner).

# What Happened?
I was writing batch processing code that handles tens of thousands of records split across multiple requests. I wrote the processing using `ReadOnlyTransaction` as follows:

```golang
for {
    // 〜略〜

    // c is *spanner.Client
    iter := c.ReadOnlyTransaction().Query(ctx, stmt)
    defer iter.Stop()

    // 〜略〜
}
```

At first glance, it seemed fine, so I ran the batch process, but a problem occurred where processing would stop after exceeding a certain number of records.

# Cause
The Spanner Go client has a session management mechanism, but due to a missing transaction termination process, the session pool was exhausted, causing requests to time out. Internally, it seemed that the session management mechanism was blocking the execution of ReadOnlyTransaction.

# Solution
I changed it to call the transaction termination process:

```golang
for {
    // 〜略〜

    // c is *spanner.Client
    tx := c.ReadOnlyTransaction()
    defer tx.Close()
    iter := tx.Query(ctx, stmt)
    defer iter.Stop()

    // 〜略〜
}
```

Without the transaction termination process, a new session is generated every time a transaction is executed, leading to exhaustion of the session pool. The reason processing stopped at the same record count each time was likely due to hitting the limit of `SPANNER_SESSION_POOL_MAX_OPEND`. When calculated, it adds up.

# Countermeasures
Aside from reading the documentation properly, there are also tool-based solutions available.
[github.com - gcpug/zagane](https://github.com/gcpug/zagane)

Additionally, monitoring the cloudspanner session count in GCP might also be a good idea.

# References
- [cloud.google.com - Sessions](https://cloud.google.com/spanner/docs/sessions?hl=ja)
- [chidakiyo.hatenablog.com - A story about fiddling with Spanner in Go](https://chidakiyo.hatenablog.com/entry/2020/12/14/go-spanner-tools)
- [medium.com - Detailed explanation of google-cloud-go/spanner — Session Management Edition](https://medium.com/google-cloud-jp/%E8%A9%B3%E8%A7%A3-google-cloud-go-spanner-%E3%82%BB%E3%83%83%E3%82%B7%E3%83%A7%E3%83%B3%E7%AE%A1%E7%90%86%E7%B7%A8-d805750edc75)
- [medium.com - Detailed explanation of google-cloud-go/spanner — Transaction Edition](https://medium.com/google-cloud-jp/%E8%A9%B3%E8%A7%A3-google-cloud-go-spanner-%E3%83%88%E3%83%A9%E3%83%B3%E3%82%B8%E3%82%B7%E3%83%A7%E3%83%B3%E7%B7%A8-6b63099bd7fe)