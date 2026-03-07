---
title: AWSサービスの種類ーAZ・リージョン・グローバル
slug: types-of-aws-services-az-region-global
date: 2021-06-23T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Amazon Web Service
translation_key: types-of-aws-services-az-region-global
---


# 概要
AWSの代表的なサービスの分類をメモ。

# AZサービス
- サブネット（複数AZにはまたがらない）
- NAT
- ネットワークACL
- EC2インスタンス
- RDSインスタンス
- Elastic Cache
- Elastic File System
- Elastic Load Balancing
- EBS Volume
- Redshift

# リージョンサービス
- VPC（複数のAZにまたがることができる）
- セキュリティグループ
- VPC Endpoints
- VPS Peering
- Elastic IP
- Auto Scaling
- S3
- Glacier
- Snapshot（EBS/RDS）
- AMI
- DynamoDB
- SQS
- SNS
- CloudSearch
- CloudWatch
- Lamda
- API Gateway

# グローバルサービス
- IAM（ユーザー、グループ、ポリシー、ロール）
- Route53
- CloudFront
- CloudTrail
- WAF
- STS
