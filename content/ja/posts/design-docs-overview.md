---
title: "Design Docsについて"
slug: "design-docs-overview"
date: 2022-10-07
author: bmf-san
categories:
  - "アーキテクチャ"
tags:
  - "Design Docs"
draft: false
---

# 概要
Design Docsについて調べてみた。

# Design Docsとは
Design Docsはソフトウェア設計のためのドキュメント。

決まった形式を持たず、プロジェクトにとって意味ある形で書くことをルールとしている

Design Docsは開発プロセスにおいて、以下のようないくつかのメリットを持つ。

- 設計上の課題洗い出し、手戻りの軽減
- 設計について合意形成
- 横断的な関心事についての整理・確認
- シニアのエンジニアの知見共有

特定の形式は持たないが、設計のコンテキストやスコープ、目標や非目標を明確にすることが推奨される。

ドキュメントの長さについては、忙しい人でも手短に読める程度の長さが推奨される。

Design Docsを書くべきかどうかは、Design Docsを書くメリットがDesign Docsを運用するコストよりも上回るかどうかが基準となる。

Design Docsは次のようなライフサイクルを持っている。

- 作成と繰り返し（ドキュメントの再編集）
- レビュー
- 実装と繰り返し（ドキュメントの更新）
- メンテナンス（ドキュメントの更新）と学習（システムを触ろうとする人へのシステムの理解の補助）

Design Docsの例には例えば次のようなものがある。

- [docs.google.com - WebKit WebSocket design doc
](https://docs.google.com/document/d/1s1ryja1V8dDotMK2WBGT2wnwchZ_x7Tag2L3OZfn5Po/preview)
- [www.chromium.org - Extensions](https://www.chromium.org/developers/design-documents/extensions/)

# 所感
- チームやプロジェクトにとって価値あるドキュメントの形を定義、理解し合う必要がありそう
  - なぜ書くか（Why）、誰が読むか（Who）、どのように書くか（How）
  - 特に設計書なのか、議事録なのか、人によって見え方が違うドキュメントになるとと書き手にとっても読み手にとっても辛い
- ドキュメントの型を決めてしまうとその型を意識して型にはめ込んだ考え方をしてしまったり、型を出た考え方がしづらくなってしまうというリスクもありそう
  - 型よりも目的の意識、理解を重要だと感じる
- ドキュメントの運用コスト（レビュー・更新などのメンテナンス）がどれくらいか。そのコストを払う価値があるかよく検討する必要がある
  - メルカリshopみたいに更新しないというのありかも。あるいは気づいたときは更新するかくらいの感じでも
  - システムと同じく運用面をどうするかもちゃんと検討する必要がある
- Design Docsが持つライフサイクルはアジャイルな開発プロセスの中に上手く組み込むことできるはず
- 外部公開（APIドキュメントなど）、内部公開など閲覧範囲によっても運用の形が変わりそう
- レビュープロセスを手厚く行いたいのであればgithub上でリポジトリ管理する形もありかも

# 参考
- [www.industrialempathy.com - Design Docs at Google](https://www.industrialempathy.com/posts/design-docs-at-google/)
- [messagepassing.github.io - Design Docs のいけすかなさ](https://messagepassing.github.io/011-designdocs/01-morrita/)
- [please-sleep.cou929.nu - Google の Design Doc について](https://please-sleep.cou929.nu/20091116.html)
- [medium.com - Writing Technical Design Docs](https://medium.com/machine-words/writing-technical-design-docs-71f446e42f2e)
- [tkybpp.hatenablog.com - 【翻訳】Googleのエンジニアがソフトウェア開発する時に必ず書くドキュメント「Design Docs at Google」](https://tkybpp.hatenablog.com/entry/2020/08/03/090000)
  - [www.industrialempathy.com - Design Docs at Google](https://www.industrialempathy.com/posts/design-docs-at-google/)の翻訳
- [myenigma.hatenablog.com - Googleなどで利用されているDesign Doc入門](https://myenigma.hatenablog.com/entry/2021/07/25/140308)
  - コードに現れない部分（なぜ別の方法を選択しなかったのか？）を議論する場としてのDesign Docs
- [atmarkit.itmedia.co.jp - 残業も減らせる!? 上級エンジニアになるためのDesign Doc超入門](https://atmarkit.itmedia.co.jp/ait/articles/1606/21/news016.html)
  - Design Docsはアジャイルのプロセスにマッチする
- [www.flywheel.jp - デザインドックで学ぶデザインドック](https://www.flywheel.jp/topics/design-doc-of-design-doc/)
  - 設計についてのフィードバックを早期に得られる
- [engineering.mercari.com - メルカリShopsでのDesign Docs運用について](https://engineering.mercari.com/blog/entry/20220225-design-docs-by-mercari-shops/)
  - Design Docsは原則更新、メンテナンスを行っていない
    - 最新のDesignでないことを認識できるというメリット
- [nhiroki.jp - Design Docsへの思い](https://nhiroki.jp/2021/03/31/design-docs)
  - Design Docsを更新しないことによる陳腐化はあまり問題ではない、スナップショットとしての価値がある


