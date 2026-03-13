---
title: When AWS (Elastic Beanstalk) Latency Was Unusually High...
description: 'Diagnose and resolve high latency issues on AWS Elastic Beanstalk ELB nodes through systematic investigation and monitoring.'
slug: high-latency-aws-elasticbeanstalk
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Amazon Web Service
  - Elastic Beanstalk
translation_key: high-latency-aws-elasticbeanstalk
---



It was an issue on AWS's side.

While monitoring an instance (m4) launched on AWS (Elastic Beanstalk), I noticed that the latency was unusually high, and there seemed to be users timing out roughly once a minute. (The average was around 5 seconds... I think.)

I initially thought there might be a bottleneck on the application side, but the environment of the instance I previously launched for testing (almost the same environment) was clearly better. As an emergency measure, I decided to create a clone and operate from there.

Upon investigating the cause and contacting AWS, I received an apology from them.

The issue was due to AWS, specifically an anomaly with the ELB node. They said they would address it by replacing the ELB node.

That's the story of how such things can happen. (I wonder if issues originating from AWS are quite common...?)
