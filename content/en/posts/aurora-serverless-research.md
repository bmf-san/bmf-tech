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


I became interested in Aurora Serverless, so I decided to explore some articles.

# Articles Reviewed
- [aws.amazon.com - Amazon Aurora Serverless](https://aws.amazon.com/jp/rds/aurora/serverless/)
- [techblog.zozo.com - Aurora Serverless v2 Production Deployment: Considerations and Benefits](https://techblog.zozo.com/entry/aurora-serverless-v2)
- [www.plan-b.co.jp - Pros and Cons of Using Aurora Serverless](https://www.plan-b.co.jp/blog/tech/28232/)
- [www.stylez.co.jp - AWS Aurora Serverless Explained with Latest Features (v2)](https://www.stylez.co.jp/columns/aws_aurora_serverless_including_the_latest_features_v2/)
- [dev.classmethod.jp - Overview of Aurora Serverless](https://dev.classmethod.jp/articles/aurora-serverless-summary/)
- [qiita.com - Aurora Serverless v2 is GA: A Beginner-Friendly Summary](https://qiita.com/minorun365/items/2a548f6138b6869de51a)
- [engineering.dena.com - Evaluation of New Service Aurora Serverless v2 [DeNA Infrastructure SRE]](https://engineering.dena.com/blog/2022/06/aurora-serverless-v2/)
- [logmi.jp - Cost Comparison of Aurora Serverless v2 and Aurora Provisioned](https://logmi.jp/tech/articles/328375)
- [zenn.dev - Insights from Operating Aurora Serverless v2](https://zenn.dev/yumemi_inc/articles/7ad423906ec202)

# Features of Aurora Serverless v2
After a brief investigation, I found the following features:

- Multi-AZ support
- Auto-scaling (scale-up and scale-down, not scale-out)
- Accessible from outside the VPC (public IP assignment possible)
- Higher cost compared to Aurora Provisioned
- Suitable for use cases with predictable spikes at specific times
