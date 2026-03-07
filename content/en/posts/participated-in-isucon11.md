---
title: Participated in ISUCON11
slug: participated-in-isucon11
date: 2021-10-21T00:00:00Z
author: bmf-san
categories:
  - Poetry
tags:
  - ISUCON
  - ISUCON10
translation_key: participated-in-isucon11
---

# This Year I Participated in ISUCON

Continuing from last year, I participated with the same members this year.

This is my third time participating in ISUCON.

[bmf-tech.com - Road to ISUcon](https://bmf-tech.com/posts/Road%20to%20ISUcon)

[bmf-tech.com - Participated in ISUCON10](https://bmf-tech.com/posts/ISUCON10%E3%81%AB%E5%8F%82%E5%8A%A0%E3%81%97%E3%81%A6%E3%81%8D%E3%81%9F)

Over the past year, I focused on studying the flow of investigating bottlenecks, practicing operations, and analyzing past ISUCON questions with my team based on last year's KPT, but we ended up failing in the preliminaries...

The things we studied and practiced were somewhat effective, and we were able to work smoothly during the actual event and follow the planned schedule, so I think we performed better than last year, but our approach to increasing the score ended up being half-hearted. We ran out of time.

# Differences from Last Year

Last year, we participated offline, but this year, considering the current situation, we participated online, so we needed to consider tools for information sharing.

There are various online whiteboard tools, but we decided to use Miro, which we often use at the company and are familiar with.

![Screenshot 2021-08-22 21 26 05](https://user-images.githubusercontent.com/13291041/130355082-819552fe-7bf3-4af2-a0b7-f8db76df5d2b.png)

We utilized it for tracking investigation status on the day and for reflections after the event.

We kept Discord connected for chat. Since AirPods run out of battery, I also prepared wired earphones to avoid wasting time on audio issues.

# What We Did

There were two members focusing on application work and one (myself) handling infrastructure and some application tasks, and we performed the following:

- Thoroughly read the regulations
- Thoroughly read the application manual
- Built the competition environment (just clicked cloud formation)
- Prepared the deployment environment
- Introduced pprof
- Introduced alp & analyzed access logs
- Checked processes
- Confirmed architecture structure, versions of various middleware, and settings
- Set slowquery & mysqldumpslow
- Checked system load
- Confirmed DB structure
- Checked screen structure
- Organized endpoint list
- Tuned Nginx
- Adjusted Cache-Control headers
- Resolved N+1 queries
- Various tasks to save image binaries stored in the DB as files (batch processing for initial data generation, adjustments to APIs for writing and referencing image data → we were scolded during the benchmarker's specification check and couldn't release)
- Split servers for app and db (1 app server, 1 db server. The third server was only used like a test environment, so we couldn't fully utilize resources...)
- Stopped log output for MySQL and Nginx (in the final stages)

We scheduled enough time for reading regulations and manuals together, and for preparation and investigation time to understand bottlenecks (strongly conscious of not spending too much time on areas with small improvement impacts). We started tuning work around 1 PM. (One of the homework items for next time is to start tuning work a bit earlier.)

# Thoughts

I felt frustrated that there were still many things we could do, but we couldn't finish our tasks in time.

Over the past year, we held online study sessions with the team every 2-3 weeks, but I didn't expect time to pass so quickly...

I want to participate again next year, so I will continue to improve my skills over the next year.

Next year, I hope to significantly increase our score...

P.S.
Thank you to the organizers. 

I had a great time this year too!