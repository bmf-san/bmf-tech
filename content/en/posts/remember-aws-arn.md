---
title: Remembering AWS ARNs
description: 'Understand AWS ARN format structure with partition, service, region, account-id, and resource-id components for resource naming.'
slug: remember-aws-arn
date: 2021-06-23T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Amazon Web Service
translation_key: remember-aws-arn
---

Make a note to remember the format of AWS Resource Names (ARN).

```
arn:partition:service:region:account-id:resource-id
arn:partition:service:region:account-id:resource-type/resource-id
arn:partition:service:region:account-id:resource-type:resource-id
```

*Apsrairr*

The last 'r' is either `/` or `:`.

cf. [Amazon Resource Names (ARN)](https://docs.aws.amazon.com/en_us/general/latest/gr/aws-arns-and-namespaces.html)
