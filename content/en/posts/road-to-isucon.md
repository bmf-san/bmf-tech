---
title: Road to ISUcon
slug: road-to-isucon
date: 2023-06-07T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Performance Tuning
  - ISUCON
  - ISUCON8
description: Preparation for participating in ISUcon.
translation_key: road-to-isucon
---


# Overview
This post details the preparations made for participating in ISUcon.

# Goals and Objectives
- **Goals**
  - Use the full duration of the ISUcon preliminaries for tuning.
    - While winning is the goal, as a first-time participant, a realistic target was set.
- **Objectives**
  - Enhance knowledge around infrastructure.
  - Improve understanding of building applications with performance considerations.
  - Aim to contribute to performance tuning tasks within the company (work hard).

# Recruited Members
Invited a colleague from the company to form a two-person team.

# Reviewed Performance Tuning Procedures
## Preparation
- Checked regulations
- Set up the environment
- Configured SSH
- Took backups
- Built automated deployment
  - Simple scripts
    - Execute deployment tasks via SSH connection, restart middleware, etc.
- Implemented monitoring tools
  - netdata
- Implemented profiling tools
  - alp
    - Access log profiler
    - https://github.com/tkuchiki/alp
  - php
    - xhprof
- Benchmarking

## Configuration Check
- Process check
  - Stop unnecessary ones
- Hardware resources
  - top
- Database
  - MySQL
    - Check data size
      - Verify size and row count per table
```
mysql> use database;
mysql> SELECT table_name, engine, table_rows, avg_row_length, floor((data_length+index_length)/1024/1024) as allMB, floor((data_length)/1024/1024) as dMB, floor((index_length)/1024/1024) as iMB FROM information_schema.tables WHERE table_schema=database() ORDER BY (data_length+index_length) DESC;
```
- Application
  - Identify URLs
  - Code reading

## Tuning
- Measure performance → Analyze → Tune → Measure
  - Determine if it's CPU, IO, application, or DB

# Impressions
Participating for the first time was a humbling experience. Despite the lack of practice, basic skills were lacking, so I aim to prepare and improve to compete properly next year. It was a valuable opportunity to reassess my challenges and significantly influence my future motivation. I am grateful to the organizers for such an incredible event, especially since it's free.

# Reference Repository
- [Collection of Vagrantfiles for Building Past ISUCON Problems](https://github.com/matsuu/vagrant-isucon)

# Reference Links
- [ISUcon Official Site](http://isucon.net/)
- [Presentation Materials on "Tips for Improving Web Application Performance"](http://blog.nomadscafe.jp/2014/08/isucon-2014-ami.html)
- [Practical Tips for Improving Web Application Performance](https://www.slideshare.net/kazeburo/isucon-summerclass2014action2)
- [ISUCON7 Preliminary Preparation for Beginners](http://isucon.net/archives/50697356.html)
- [ISUCON Winning Strategy Video Notes](https://wiki.infra-workshop.tech/user/kameneko/ISUCON8/ISUCON%E3%81%AE%E5%8B%9D%E3%81%A1%E6%96%B9%E5%8B%95%E7%94%BB%E3%83%A1%E3%83%A2)
- [Steps to Build ISUCON7 (Preliminary) Environment on ConoHa](https://nishinatoshiharu.com/how-to-create-isucon-conoha/)
- [YAPC::Asia Tokyo 2015 "ISUCON Winning Strategy" Notes](http://kobtea.net/posts/2015/08/22/yapc-isucon/)
- ~~Notes for Completing ISUCON Preliminaries Without Tears Next Year (1)~~

# Books Read for ISUcon
- [Systems Performance: Enterprise and the Cloud](https://www.oreilly.co.jp/books/9784873117904/)
  - Still not finished...
- [Understanding System Performance Through Illustrations](https://www.shoeisha.co.jp/book/detail/9784798134604)
- [Technology Supporting Servers/Infrastructure 24/7: Scalability, High Performance, Efficient Operation](https://gihyo.jp/magazine/wdpress/plus/978-4-7741-3566-3)
