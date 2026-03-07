---
title: Got Stuck with ReadOnlyTransaction in Go's Spanner Client
slug: go-spanner-client-read-only-transaction-issue
date: 2021-02-08T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Google Cloud Platform
  - Spanner
  - Tips
description: Notes on issues encountered while using ReadOnlyTransaction in the Go Spanner client.
translation_key: go-spanner-client-read-only-transaction-issue
---

# Overview
[pkg.go.dev - cloud.google.com/go/spanner](https://pkg.go.dev/cloud.google.com/go/spanner) Notes on issues encountered while using `ReadOnlyTransaction`.

# What happened?
I was writing batch processing code to handle tens of thousands of data records split across multiple requests. The processing using `ReadOnlyTransaction` was written as follows:

```golang
for {
    // ~omitted~

    // c is *spanner.Client
    iter := c.ReadOnlyTransaction().Query(ctx, stmt)
    defer iter.Stop()

    // ~omitted~
}
```

At first glance, the code seemed fine, so I ran the batch process. However, the process stopped after exceeding a certain number of records.

# Cause
The Go Spanner client has a session management mechanism. Due to missing transaction termination handling, the session pool was exhausted, causing requests to time out. Internally, it seems the session management mechanism was blocking the execution of `ReadOnlyTransaction`.

# Solution
Modify the code to call the transaction termination process:

```golang
for {
    // ~omitted~

    // c is *spanner.Client
    tx := c.ReadOnlyTransaction()
    defer tx.Close()
    iter := tx.Query(ctx, stmt)
    defer iter.Stop()

    // ~omitted~
}
```

Without transaction termination handling, a new session is created for each transaction execution, exhausting the session pool. The process consistently stopped at the same number of records due to hitting the `SPANNER_SESSION_POOL_MAX_OPEND` limit. Calculations confirmed this.

# Countermeasures
Besides thoroughly reading the documentation, tools can help resolve this issue.
[github.com - gcpug/zagane](https://github.com/gcpug/zagane)

Additionally, monitoring the session count of Cloud Spanner via GCP monitoring might be a viable option.

# References
- [cloud.google.com - Sessions](https://cloud.google.com/spanner/docs/sessions?hl=ja)
- [chidakiyo.hatenablog.com - Go and Spanner: Experiments for Better Integration](https://chidakiyo.hatenablog.com/entry/2020/12/14/go-spanner-tools)
- [medium.com - Detailed Explanation of google-cloud-go/spanner — Session Management](https://medium.com/google-cloud-jp/%E8%A9%B3%E8%A7%A3-google-cloud-go-spanner-%E3%82%BB%E3%83%83%E3%82%B7%E3%83%A7%E3%83%B3%E7%AE%A1%E7%90%86%E7%B7%A8-d805750edc75)
- [medium.com - Detailed Explanation of google-cloud-go/spanner — Transactions](https://medium.com/google-cloud-jp/%E8%A9%B3%E8%A7%A3-google-cloud-go-spanner-%E3%83%88%E3%83%A9%E3%83%B3%E3%82%B6%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3%E7%B7%A8-6b63099bd7fe)