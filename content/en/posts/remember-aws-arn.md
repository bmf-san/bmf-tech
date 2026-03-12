---
title: Remembering AWS ARNs
description: An in-depth look at Remembering AWS ARNs, covering key concepts and practical insights.
slug: remember-aws-arn
date: 2021-06-23T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Amazon Web Service
translation_key: remember-aws-arn
---



AWSのリソースネーム（ARN)の形式を忘れないようにメモ。

```
arn:partition:service:region:account-id:resource-id
arn:partition:service:region:account-id:resource-type/resource-id
arn:partition:service:region:account-id:resource-type:resource-id
```

*Apsrairr*

最後のrは`/`か`:`。

cf. [Amazon リソースネーム (ARN)](https://docs.aws.amazon.com/ja_jp/general/latest/gr/aws-arns-and-namespaces.html)
