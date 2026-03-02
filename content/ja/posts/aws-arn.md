---
title: "AWSのARNを覚える"
slug: "aws-arn"
date: 2021-06-23
author: bmf-san
categories:
  - "インフラストラクチャ"
tags:
  - "Amazon Web Service"
draft: false
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
