---
title: ADR（Architecture Decision Record）について
slug: adr-architecture-decision-record
date: 2022-10-10T00:00:00Z
author: bmf-san
categories:
  - アーキテクチャ
tags:
  - Architecture Decision Record
translation_key: adr-architecture-decision-record
---


# 概要
ADR（Architecture Decision Record）について調べた。

# ADRとは
2011年にMichael Nygardによって紹介されたアーキテクチャに関する決定事項を記録したドキュメントのこと。
cf. [cognitect.com - DOCUMENTING ARCHITECTURE DECISIONS](https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions)

# フォーマット
Michael Nygardは提案するフォーマットは次の通り。
cf. [cognitect.com - DOCUMENTING ARCHITECTURE DECISIONS](https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions)

- タイトル
- コンテキスト
  - 文脈。どういうシーンにおける決定なのか。
- 決定
  - 決定事項は1つのADRにつき1つが推奨
- ステータス
  - 提案、承認、非推奨など
- 結果
  - 決定事項を適用した後にどういう状態になるか、結果のコンテキストについての説明

ドキュメントは1~2ページ程度の読みやすい長さにする。

# ADRを採用する利点
- アーキテクチャに関する設計判断についてキャッチアップしやすくなる
- 明記された決定事項に関する議論頻度の減少
- チーム内外におけるアーキテクチャの透明性への貢献

# ADRを採用するかどうかの判断
- チームで繰り返し議論を行うような内容であるか？
- チームで意思決定した設計について決定事項であるか？
- ソフトウェア全体に影響がある決定事項であるか？

などアーキテクチャの何かしらの判断・決定をする際はADRを書く機会がある。

# 所感
- [ソフトウェアアーキテクチャの基礎](https://amzn.to/3SQV0ge)では、
- ADRを導入しようと思った際は、ADRの導入についてのADRをまずは書いてみるのが良さそう
- Design Docsと少し雰囲気が違って、ADRの管理はリポジトリでソースコードとともに管理されるのが好まれる雰囲気を感じた
- ADRを書くか書かないかの判断基準はシンプルにするのが良いと思った
  - チームで議論し、決定したことについては全て書く、など
- ADRの検索性
  - 運用上ADRを検索したいときにちゃんと検索しやすい工夫があると良いのかもしれない（数が多くなるのであれば）
  - カテゴライズ、タグなど工夫したほうが良さそう

# 参考
- [github.com - Architecture decision record (ADR)](https://github.com/joelparkerhenderson/architecture-decision-record/blob/main/README.md)
- [cognitect.com - DOCUMENTING ARCHITECTURE DECISIONS](https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions)
- [docs.aws.amazon.com - ADR process](https://docs.aws.amazon.com/prescriptive-guidance/latest/architectural-decision-records/adr-process.html)
- [betterprogramming.pub - The Ultimate Guide to Architectural Decision Records
Introduction](https://betterprogramming.pub/the-ultimate-guide-to-architectural-decision-records-6d74fd3850ee)
- [engineering.atspotify.com - When Should I Write an Architecture Decision Record](https://engineering.atspotify.com/2020/04/when-should-i-write-an-architecture-decision-record/)
- [www.redhat.com - Why you should be using architecture decision records to document your project
](https://www.redhat.com/architect/architecture-decision-records)
- [www.thoughtworks.com - Lightweight Architecture Decision Records](https://www.thoughtworks.com/radar/techniques/lightweight-architecture-decision-records)
- [developer.mamezou-tech.com - アーキテクチャ・デシジョン・レコードの勧め](https://developer.mamezou-tech.com/blogs/2022/04/28/adr/)
- [cloud.google.com - アーキテクチャ決定レコードの概要](https://cloud.google.com/architecture/architecture-decision-records)
- [qiita.com - アーキテクチャの「なぜ？」を記録する！ADRってなんぞや？](https://qiita.com/fuubit/items/dbb22435202acbe48849)
- [qiita.com - アーキテクチャに関するドキュメントの残しかた - ADRとARCHITECTURE.md テンプレート](https://qiita.com/e99h2121/items/f508ef4c9743b8fc9f5b)
- [blog.studysapuri.jp - 〜その意思決定を刻め〜「アーキテクチャ・デシジョン・レコード(ADR)」を利用した設計の記録](https://blog.studysapuri.jp/entry/architecture_decision_records)
- [fintan.jp - Architecture Decision Records導入事例](https://fintan.jp/page/1636/)
