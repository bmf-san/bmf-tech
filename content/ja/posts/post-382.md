---
title: "サービスメッシュについて"
slug: "post-382"
date: 2023-10-29
author: bmf-san
categories:
  - "インフラストラクチャ"
tags:
  - "サービスメッシュ"
draft: false
---

# 概要
サービスメッシュについて調べたことをまとめる。

# サービスメッシュとは
サービス間（分散システム）の通信を管理するためのネットワークインフラストラクチャのこと。

一般的にはサービスにプロキシをサイドカーとして追加することで構成する。

# サービスメッシュが解決すること
- 可観測性の向上
- サービス間の通信の管理・制御
- サービス間の通信のセキュリティ向上

# サービスメッシュの機能
- サービスディスカバリ
  - DNSやロードバランサーによるサービスの接続情報管理ではデリバリーの遅延が発生する
  - サービスの接続情報管理をサービスメッシュに任せることでデリバリーの遅延を抑える
- サービスルーティング
  - サービス間のトラフィックのルーティングをサービスメッシュに任せることで、サービス間のトラフィックのルーティングを柔軟に変更できる
- 障害の分離
  - ex. サーキットブレーカー
- 負荷分散
- 認証と認可
  - 各サービスごとに仕組みを用意するのではなく、サービスメッシュに任せることで、各サービスの負担を軽減できる
- オブザーバリティ
  - 複数のサービスに跨るトラフィックの可観測性を向上させる
    - トレーシングを可能にする

# サービスメッシュのデメリット
以下は一般的なプロキシを必要とするサービスメッシュにおけるデメリット。
プロキシレスなサービスメッシュ（ex. Traffic Director）はその限りではない。

- 通信のパフォーマンス低下
- リソース使用量の増加

# 参考
- [speakerdeck.com - Service Meshがっつり入門/Get-Started-Service-Mesh](https://speakerdeck.com/oracle4engineer/get-started-service-mesh)
- [cloud.google.com - マイクロサービス アーキテクチャのサービス メッシュ](https://cloud.google.com/architecture/service-meshes-in-microservices-architecture?hl=ja)
- [www.netone.co.jp - サービスメッシュ入門](https://www.netone.co.jp/media/detail/20200715-1/)
- [aws.amazon.com - サービスメッシュとは何ですか?](https://aws.amazon.com/jp/what-is/service-mesh/)
- [www.alpha.co.jp - サービスメッシュ導入の前に知っておくべきこと](https://www.alpha.co.jp/blog/202205_01)
- [www.redhat.com - サービスメッシュとは](https://www.redhat.com/ja/topics/microservices/what-is-a-service-mesh)
- [qiita.com - サービスメッシュについて調査してみた件](https://qiita.com/mamomamo/items/92085e0e508e18bc8532)
- [www.infoq.com - サービスメッシュ必読ガイド - マイクロサービス時代のサービス間通信管理](https://www.infoq.com/jp/articles/service-mesh-ultimate-guide/)
- [https://dev.classmethod.jp -サービスメッシュについて理解する ](https://dev.classmethod.jp/articles/servicemesh/)
