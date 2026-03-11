---
title: モジュラモノリスについて調べたことをメモ
slug: modular-monolith-research-notes
date: 2023-09-25T00:00:00Z
author: bmf-san
categories:
  - アーキテクチャ
tags:
  - モジュラモノリス
translation_key: modular-monolith-research-notes
---


# 概要　
モジュラモノリスについて調べたことをメモする。

# モジュラモノリスとは
- モジュール分割をしたモノリス
  - モジュール分割はドメインによる分割が一般的に見えるが、機能分割や技術分割など様々なパターンを検討することができる
- モノリスと同じく単一のデプロイメントパイプラインを持つ
- メリット
  - モジュールが分割されているのでモジュール単位で開発を独立して行うことができる
  - マイクロサービスへの移行が容易
    - モジュール単位でマイクロサービス化を進めやすい
      - モノリスからマイクロサービスへ移行する際、ストラングラーパターンとしての位置づけができる（但し楽に移行できるとは限らない）
    - モジュールの境界の見直しがしやすいので、マイクロサービスよりも境界の変化に柔軟に対応できる
- デメリット
  - モジュール間の境界を超えやすい
    - マイクロサービスはネットワークを介して通信するため、モジュール間の境界を超えることができないが、モジュラモノリスはそうではないため境界違反の扱いに注意する必要がある
  - モノリスと同じだが、デプロイメントのパイプラインは単一なので、肥大化したりモジュール間の依存関係が複雑になったりすると運用が難しくなる
  - 各モジュールが単一のDBを共有している場合、マイクロサービス移行時のコストが高くなる


# Service Weaver
モジュラモノリスとして開発し、マイクロサービスとしてデプロイするためのツールとして、Googleが[Service Weaver](https://serviceweaver.dev/)というものをリリースしている。

# 所感
モジュラモノリスというものに少しの夢を見ていたのだが、所感としてはポエムを書きたくなったので綴った。

組織の拡大に合わせてアーキテクチャも進化が必要になってくるのが筋だと思うが、組織のスケーラビリティに柔軟に対応できる、あるいはコストがかかり過ぎない銀の弾丸のようなアーキテクチャがないかなぁと思ったりした。（ない。）

組織は拡大したり、縮小したり、変化しなかったりするような時もあるとは思うが、会社は成長を前提とするので組織のスケーラビリティには前のり気味で投資していくのが良いのかなぁなどと思った。

このへんの話に関して引用したい文章があったので、そちらを記載して締めとする。

> しかし、ここで注目しなければいけないのは、両者のライフサイクルの違いです。 組織やチーム配置は、その気になれば会社の方針次第で翌日から変更できます。 しかしアーキテクチャやシステムは、組織のようにすぐ変更することが困難です。

cf. [eh-career.com - モジュラモノリスに移行する理由 ─ マイクロサービスの自律性とモノリスの一貫性を両立させるアソビューの取り組み大規模](https://eh-career.com/engineerhub/entry/2022/07/25/093000)

# 参考
- [microservices.io - How modular can your monolith go? Part 1 - the basics](https://microservices.io/post/architecture/2023/07/31/how-modular-can-your-monolith-go-part-1.html)
- [microservices.io - How modular can your monolith go? Part 2 - a first look at how subdomains collaborate](https://microservices.io/post/architecture/2023/08/20/how-modular-can-your-monolith-go-part-2.html)
- [microservices.io - How modular can your monolith go? Part 3 - encapsulating a subdomain behind a facade](https://microservices.io/post/architecture/2023/08/28/how-modular-can-your-monolith-go-part-3.html)
- [microservices.io - How modular can your monolith go? Part 4 - physical design principles for faster builds](https://microservices.io/post/architecture/2023/09/12/how-modular-can-your-monolith-go-part-4-physical-design.html)
- [techblog.hacomono.jp - モノリスなRailsにモジュラーモノリスを導入した話](https://techblog.hacomono.jp/entry/2023/08/22/110000#:~:text=%E3%83%A2%E3%82%B8%E3%83%A5%E3%83%A9%E3%83%BC%E3%83%A2%E3%83%8E%E3%83%AA%E3%82%B9%E3%81%A8%E3%81%AF%E3%80%81%E3%83%A2%E3%83%8E%E3%83%AA%E3%82%B9,%E6%B4%BB%E7%94%A8%E3%81%A7%E3%81%8D%E3%82%8B%E7%89%B9%E5%BE%B4%E3%81%8C%E3%81%82%E3%82%8A%E3%81%BE%E3%81%99%E3%80%82)
- [tech-blog.rakus.co.jp - サービス分割に備えたモノリス（モジュラーモノリスとかアグリゲートとか）](https://tech-blog.rakus.co.jp/entry/20201026/microservice)
- [r-kaga.com - Modular Monolith(モジュラーモノリス)についてまとめる](https://r-kaga.com/blog/what-is-modular-monolith)
- [speakerdeck.com - モジュラモノリスで表現する複雑なドメイン領域と境界](https://speakerdeck.com/showmant/expressing-complex-domain-regions-and-boundaries-with-modular-monoliths)
- [shopify.engineering - Deconstructing the Monolith: Designing Software that Maximizes Developer Productivity](https://shopify.engineering/deconstructing-monolith-designing-software-maximizes-developer-productivity)
- [www.infoq.com - モノリスの分解において、マイクロサービスは必然ではない - QCon LondonにおけるSam Newman氏の講演より](https://www.infoq.com/jp/news/2020/06/monolith-decomposition-newman/?utm_campaign=infoq_content&utm_source=infoq&utm_medium=feed&utm_term=global)
- [www.infoq.com - Shopifyはいかにしてモジュラモノリスへ移行したか](https://www.infoq.com/jp/news/2019/10/shopify-modular-monolith/)
- [www.publickey1.jp - Google、モノリスとマイクロサービスのいいとこ取りをする「Service Weaver」フレームワークをオープンソースで公開](https://www.publickey1.jp/blog/23/googleservice_weaver.html)
- [モジュラモノリスにおけるトランザクション設計の考え方](https://speakerdeck.com/nazonohito51/transaction-design-on-modular-monolith)
- [dev.classmethod.jp - [レポート] モノリスかマイクロサービスか、その選択に迷っている人へ届けたい話 #devio_day1 #main](https://dev.classmethod.jp/articles/20230411-developersio-day-one-monolithandmicroservices/)
- [engineering.mercari.com - メルカリの取引ドメインにおけるモジュラーモノリス化の取り組み](https://engineering.mercari.com/blog/entry/20220913-modular-monolithization-in-mercari-transaction-domain/)
- [eh-career.com - モジュラモノリスに移行する理由 ─ マイクロサービスの自律性とモノリスの一貫性を両立させるアソビューの取り組み大規模](https://eh-career.com/engineerhub/entry/2022/07/25/093000)
- [logmi.jp - 重要なのは「先を見据えた柔軟なアーキテクチャ構成」3チームの並列開発を実現したモジュラモノリスの採用](https://logmi.jp/tech/articles/328130)
- [medium.com - The Modular Monolith: Rails Architecture](https://medium.com/@dan_manges/the-modular-monolith-rails-architecture-fb1023826fc4)

