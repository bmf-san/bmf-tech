---
title: Remembering AWS ARN
slug: remember-aws-arn
date: 2021-06-23T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Amazon Web Service
translation_key: remember-aws-arn
---

Here's a note to help remember the format of AWS Resource Names (ARN).

```
arn:partition:service:region:account-id:resource-id
arn:partition:service:region:account-id:resource-type/resource-id
arn:partition:service:region:account-id:resource-type:resource-id
```

*Apsrairr*

The last 'r' can be either `/` or `:`.

cf. [Amazon Resource Names (ARN)](https://docs.aws.amazon.com/ja_jp/general/latest/gr/aws-arns-and-namespaces.html)