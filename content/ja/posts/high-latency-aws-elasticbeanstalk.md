---
title: AWS(Elasticbeanstalk)のレイテンシーがやたら高いと思ったら・・・
description: AWS(Elasticbeanstalk)のレイテンシーがやたら高いと思ったら・・・について、基本的な概念から実践的な知見まで詳しく解説します。
slug: high-latency-aws-elasticbeanstalk
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Amazon Web Service
  - Elastic Beanstalk
translation_key: high-latency-aws-elasticbeanstalk
---


AWS側の問題でした。

AWS(Elasticbeanstalk)で立ち上げたインスタンス(m4)のモニタリングをしていたら、レイテンシーがやたら高く、1分に一回くらいの頻度でタイムアウトしているユーザーがいるような状況でした。（アベレージは5秒くらいだった・・かな）

アプリケーション側にネックがあるのかなぁと思ったのですが、以前テストで立ち上げたインスタンスの環境（ほぼほぼ同じ環境）よりも明らかに悪かったので、応急処置としてクローンを作成してそちらで運用することにしました。

原因究明のため、AWSに問い合わせたところ・・・・AWSから謝罪が来ました。

原因はAWS側に起因するもので、ELBノードに異常があったからだそうです。
ELBノードを入れ替えることで対応するとのことでした。

以上、こんなこともあるんだなぁという話でした。（AWS側に起因するような問題って結構あるのでしょうか・・・？）

