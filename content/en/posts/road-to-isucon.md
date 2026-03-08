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
translation_key: road-to-isucon
---

# Overview
This post outlines the preparations made for participating in ISUcon.

# Goals and Objectives
- Goal
    - Make full use of the preliminary round time for tuning
        - Winning is the goal, but since this is my first participation, I set a realistic target.
- Objectives
    - Enhance knowledge around infrastructure
    - Enhance knowledge for building applications with performance considerations
    - Aim to be involved in performance tuning tasks within the company (I will do my best)

# Team Recruitment
I invited colleagues from the company and formed a two-person team to participate.

# Confirmed Performance Tuning Procedures
## Preparation
- Regulation check
- Environment setup
- SSH configuration
- Backup acquisition
- Automated deployment setup
 - Simple scripts
  - SSH connection to execute deployment tasks, restart middleware, etc.
- Monitoring tool introduction
  - netdata
- Profiling tool introduction
  - alp
    - Access log profiler
    - https://github.com/tkuchiki/alp
  - php
    - xhprof
- Benchmarking

## Configuration Check
- Process check
  - Stop unnecessary processes
- Hardware resources
  - top
- Database
  - MySQL
    - Check data size
      - Confirm size and row count for each table
```
mysql> use database;
mysql> SELECT table_name, engine, table_rows, avg_row_length, floor((data_length+index_length)/1024/1024) as allMB, floor((data_length)/1024/1024) as dMB, floor((index_length)/1024/1024) as iMB FROM information_schema.tables WHERE table_schema=database() ORDER BY (data_length+index_length) DESC;
```
- Application
  - Identify URLs
  - Code reading

## Tuning
- Performance measurement → analysis → tuning → measurement
 - Is it CPU, IO, application, or DB?

# Thoughts
I lost my dignity in my first participation. While I lacked practice, I also failed to do basic things, so I want to prepare and grow to compete properly next year. It was a good opportunity to reflect on my challenges and it will greatly influence my motivation moving forward. I am very grateful to the organizers for such a great competition that feels like it should cost money.

# Reference Repository
- [Vagrantfile collection for building past ISUCON problems](https://github.com/matsuu/vagrant-isucon)

# Reference Links
- [ISUcon Official Site](http://isucon.net/)
- [Presentation materials on tips for improving web application performance](http://blog.nomadscafe.jp/2014/08/isucon-2014-ami.html)
- [Practical tips for improving web application performance](https://www.slideshare.net/kazeburo/isucon-summerclass2014action2)
- [ISUCON7 preliminary countermeasures for beginners](http://isucon.net/archives/50697356.html)
- [Video notes on how to win ISUCON](https://wiki.infra-workshop.tech/user/kameneko/ISUCON8/ISUCON%E3%81%AE%E5%8B%9D%E3%81%A1%E6%96%B9%E5%8B%95%E7%94%BB%E3%83%A1%E3%83%A2)
- [How to set up the ISUCON7 (preliminary) past exam environment on ConoHa](https://nishinatoshiharu.com/how-to-create-isucon-conoha/)
- [YAPC::Asia Tokyo 2015 notes on how to win ISUCON](http://kobtea.net/posts/2015/08/22/yapc-isucon/)
- [Notes for finishing the ISUcon preliminary next year without crying (1)](https://cloudpack.media/25281)

# Books Read for ISUcon
- [Detailed Explanation of System Performance](https://www.oreilly.co.jp/books/9784873117904/)
    - I haven't finished reading it yet....
- [Understanding System Performance Mechanisms Visually](https://www.shoeisha.co.jp/book/detail/9784798134604)
- [[24/7] Technologies Supporting Servers/Infrastructure... Scalability, High Performance, and Labor-saving Operations](https://gihyo.jp/magazine/wdpress/plus/978-4-7741-3566-3)