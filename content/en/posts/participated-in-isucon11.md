---
title: Participated in ISUCON11
slug: participated-in-isucon11
date: 2021-10-21T00:00:00Z
author: bmf-san
categories:
  - Poem
tags:
  - ISUCON
  - ISUCON10
translation_key: participated-in-isucon11
---



# Participated in ISUCON Again This Year
Following last year, we participated with the same members again this year.

This marks our third time participating in ISUCON.

[bmf-tech.com - Road to ISUcon](https://bmf-tech.com/posts/Road%20to%20ISUcon)

[bmf-tech.com - Participated in ISUCON10](https://bmf-tech.com/posts/ISUCON10%E3%81%AB%E5%8F%82%E5%8A%A0%E3%81%97%E3%81%A6%E3%81%8D%E3%81%9F)

Over the past year, based on last year's KPT with the team, we focused on practicing the flow and operations up to bottleneck investigation, studying all past ISUCON problems to understand the trends and solution patterns, but ended up being eliminated in the preliminaries...

The things we studied and practiced were somewhat fruitful, allowing us to work smoothly during the event and follow the planned schedule, so I think we performed better than last year. However, our approach to increasing the score ended up being half-baked. We ran out of time orz

# Differences from Last Year
Last year we participated offline, but this year, considering the circumstances, we participated online, so we needed to consider tools for information sharing.

There are various online whiteboard tools, but since we often use miro at work and are familiar with it, we decided to use it. ![Screenshot 2021-08-22 21 26 05](/assets/images/posts/participated-in-isucon11/130355082-819552fe-7bf3-4af2-a0b7-f8db76df5d2b.png) It was utilized for checking the investigation status on the day and for retrospectives afterward.

We kept Discord connected for chat. Since AirPods run out of battery, we also prepared wired earphones to avoid wasting time on audio issues.

# What We Did
We worked with two members focusing on the application side and one (myself) on infrastructure + a bit of application, doing the following tasks:

- Thoroughly read the regulations
- Thoroughly read the application manual
- Set up the competition environment (just clicked cloud formation)
- Prepared the deployment environment
- Introduced pprof
- Introduced alp & analyzed access logs
- Checked processes
- Checked architecture configuration, middleware versions, and settings
- Set slowquery & mysqldumpslow
- Checked system load
- Checked DB configuration
- Checked screen configuration
- Organized endpoint list
- Tuned Nginx
- Adjusted Cache-Control headers
- Resolved N+1 queries
- Various tasks to save image binaries stored in the DB as files (batch processing for initial data generation, adjustments to APIs with image data writing/referencing, etc. → couldn't release due to benchmarker's specification check)
- Split servers for app and db (1 app, 1 db. Couldn't fully utilize the third as a test environment...)
- Stopped MySQL and Nginx log output (towards the end)

We scheduled enough time for reading the regulations and manual together, and for preparation and investigation to understand bottlenecks (strongly conscious of not spending too much time on areas with small improvement impact). We started tuning work around 1 PM. (Starting tuning work earlier is one of the tasks for next time)

# Impressions
I felt frustrated that there was still so much to do, yet we couldn't finish in time.

Over the past year, we held online study sessions with the team every 2-3 weeks, but time flew by so quickly...

I want to participate again next year, so I'll work hard for another year to gain strength and participate.

Next year, I hope to significantly increase our score...


P.S.
Thank you to the organizers.

It was a lot of fun this year too!

