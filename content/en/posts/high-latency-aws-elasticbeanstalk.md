---
title: High Latency in AWS (Elastic Beanstalk) Was Due to AWS Issues...
slug: high-latency-aws-elasticbeanstalk
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Amazon Web Services
  - Elastic Beanstalk
translation_key: high-latency-aws-elasticbeanstalk
---

It was an issue on the AWS side.

While monitoring an instance (m4) launched on AWS (Elastic Beanstalk), I noticed that the latency was unusually high, with users timing out approximately once a minute. (The average was around 5 seconds... I think)

I initially thought there might be a bottleneck on the application side, but it was clearly worse than the environment of a previously launched instance (almost the same environment). As a temporary measure, I decided to create a clone and operate it instead.

To investigate the cause, I contacted AWS, and... I received an apology from AWS.

The cause was attributed to AWS, as there was an anomaly with the ELB node. They mentioned that they would address it by replacing the ELB node.

So, this was a reminder that such things can happen. (I wonder if issues originating from AWS are quite common...?)