---
title: Participated in ISUCON10
description: An in-depth look at Participated in ISUCON10, covering key concepts and practical insights.
slug: isucon10-participation
date: 2020-05-21T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - ISUCON
  - ISUCON10
translation_key: isucon10-participation
---

[ISUCON10](http://isucon.net/archives/54704557.html) was held, and I participated. It was my second time participating in ISUCON after two years.

Last time, I participated as a two-person team, but this time I was part of a three-person team.

# What We Did
- Reviewed the preliminary manual and regulations
- Confirmed SSH connection
  - I was able to connect using the key registered on GitHub. This time, a bastion server was provided.
- Checked processes
- Verified API endpoints and UI
- Checked the database
- Prepared deployment
  - Made it possible to deploy the application and manage the schema with a single make command.
- Introduced alp
- Adjusted MySQL general query log, slow query log, and various parameters
- Adjusted Nginx settings
- Executed the initial benchmark

## Tasks
- Query tuning
- Filtering bot access (blocking)
- Upgraded MySQL from 5.7 to 8
  - This was a trap... Just upgrading wasn't enough; fundamental adjustments were necessary.
    - Latitude and longitude were stored as double precision, but I needed to change them to geometry or point to enable spatial indexing. By the time I realized this, it was already too late...
    - I also needed to improve the N+1 query, but I ran out of time...
    - Later, while browsing various blogs, I noticed some teams took a competitive programming approach by optimizing the calculation logic on the application side to reduce computational load, which I found insightful.
- Configured a two-server setup for AP and DB
  - There were three servers, but due to an operational mistake, one was rendered unusable... such sadness.
  - The DB was nearly at 100% CPU usage, so I was trying to make significant improvements there, but I ran out of time towards the end.
  - Since there were only two tables, I realized afterward that I should have tried splitting the DB server between the table with location information and the other table. If I had set up the remaining two servers as APP+DB and DB, perhaps the score would have changed a bit...

# Results
We were eliminated in the preliminaries. Last time, we scored 0 points due to failure, but this time we managed to score some points, so perhaps I have grown a little.
I felt quite accustomed to the operations in the early stages, and although I couldn't resolve the bottlenecks (still struggling with them...), I felt some growth. However, there is still much to do and learn, so I need more training...

After doing a KPT with the team, we plan to start training for next year, so I hope to see you again next year!
Thank you to the organizers!!