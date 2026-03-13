---
title: Cache Write Methods
description: 'Learn the three cache write methods: Write-Through (sync writes to cache and memory), Write-Back (write to cache first, flush later), and Write-Around (bypass cache, write direct to memory).'
slug: cache-write-methods
date: 2023-06-03T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Cache
translation_key: cache-write-methods
---

# Overview
This post summarizes cache write methods.

# Write-Through
A method where data is written to both the cache and main memory simultaneously. Write operations occur on both the cache and main memory. It is easier to maintain data consistency, but there may be delays in writing.

# Write-Back
A method where data is held in the cache after being written, and is only written to main memory later. Write operations occur only in the cache, and writing to main memory happens as needed. This can hide write delays, but requires measures to maintain data consistency.

# Write-Around
A method where data is written directly to main memory, bypassing the cache. This reduces the write load on the cache. The cache is used only for reading.

# References
- [Let's investigate the differences between Write-Through and Write-Back!](https://itmanabi.com/write-through-back/)
- [Where and how? Latest trends in cache technology/products](https://techtarget.itmedia.co.jp/tt/news/1407/28/news01.html)
- [Storage Magazine translated article](https://www.jdsf.gr.jp/sms/stm/201610.html)