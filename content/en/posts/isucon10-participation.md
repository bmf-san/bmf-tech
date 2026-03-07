---
title: Participated in ISUCON10
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

[ISUCON10](http://isucon.net/archives/54704557.html) participated.
This was my second time participating in ISUCON, after a two-year gap.

Last time, I participated as a two-person team, but this time, we joined as a three-person team.

# What We Did
- Reviewed the preliminary manual and regulations
- Checked SSH connections
  - We were able to connect using the key registered on GitHub. A bastion server was provided this time.
- Checked processes
- Verified API endpoints and UI
- Checked the database
- Improved deployment
  - Made it possible to deploy the application and manage the schema with a single `make` command.
- Introduced `alp`
- Adjusted MySQL general query logs, slow query logs, and various parameters
- Tuned Nginx settings
- Ran the initial benchmarker

## Tasks
- Query tuning
- Filtering (blocking) bot access
- Upgraded MySQL from 5.7 to 8
  - This turned out to be a trap... Simply upgrading wasn’t enough; fundamental adjustments were required.
    - Latitude and longitude were stored as `double precision`, but we needed to change them to `geometry` or `point` and enable spatial indexing. By the time we realized this, it was too late...
    - We also needed to address N+1 queries, but couldn’t finish in time...
    - Later, while reading various blogs, I found that some teams took a competitive programming-like approach by optimizing the calculation logic on the application side to reduce computational load. That was insightful.
- Configured a two-server setup for the application and database
  - Although we had three servers, we lost one due to an operational mistake... very unfortunate.
  - The database server’s CPU usage was close to 100%, so we focused on improving this, but ran out of time toward the end.
  - Since there were only two tables, we should have split the database server into two: one for the table with location data and the other for the remaining table. I realized this after the event. Even with the remaining two servers, a setup of APP+DB and DB might have slightly improved our score...

# Results
We didn’t make it past the preliminaries. Last time, we failed with a score of 0, but this time we managed to score some points, so perhaps we’ve grown a little.

I felt more comfortable with the initial operations (we had been practicing ISUCON as a team for about six months), and while we struggled with bottlenecks (though we couldn’t resolve them), I felt some growth. However, there’s still a lot more we can and should do, so more practice is needed...

After conducting a KPT session with the team, we plan to start preparing for next year. Looking forward to participating again next year!

Thank you to the organizers!!