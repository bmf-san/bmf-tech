---
title: Exploring Aurora Serverless
slug: aurora-serverless-research
date: 2023-12-27T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - Amazon Web Service
translation_key: aurora-serverless-research
---

I became interested in Aurora Serverless, so I looked into some articles.

# Articles Reviewed
- [aws.amazon.com - Amazon Aurora Serverless](https://aws.amazon.com/jp/rds/aurora/serverless/)
- [techblog.zozo.com - The Story of Implementing Aurora Serverless v2 in Production - Points to Consider and Effects Obtained](https://techblog.zozo.com/entry/aurora-serverless-v2)
- [www.plan-b.co.jp - Advantages and Disadvantages of Using Aurora Serverless](https://www.plan-b.co.jp/blog/tech/28232/)
- [www.stylez.co.jp - Explanation of AWS Aurora Serverless Including Latest Features (v2)](https://www.stylez.co.jp/columns/aws_aurora_serverless_including_the_latest_features_v2/)
- [dev.classmethod.jp - Summary of Aurora Serverless](https://dev.classmethod.jp/articles/aurora-serverless-summary/)
- [qiita.com - Finally! The Rumored Aurora Serverless v2 is GA. A Beginner-Friendly Summary](https://qiita.com/minorun365/items/2a548f6138b6869de51a)
- [engineering.dena.com - Verification and Evaluation of New Service Aurora Serverless v2 [DeNA Infrastructure SRE]](https://engineering.dena.com/blog/2022/06/aurora-serverless-v2/)
- [logmi.jp - Cost Comparison Between Aurora Serverless v2 and Aurora Provisioned - Exploring Advantageous Boundaries with Three Variables as a Compass](https://logmi.jp/tech/articles/328375)
- [zenn.dev - Insights Gained from Operating Aurora Serverless v2](https://zenn.dev/yumemi_inc/articles/7ad423906ec202)

# Features of Aurora Serverless v2
From my rough research, I found the following features:

- Multi-AZ support
- Auto-scaling support (not scale-out, but scale-up and down)
- Access from outside the VPC is possible (public IP assignment is possible)
- Higher costs compared to Aurora Provisioned
- Use cases seem to fit scenarios where spikes are expected during specific time periods.