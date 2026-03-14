---
title: API Design Patterns
description: 'Explore API design patterns, resource-oriented vs RPC approaches, and design principles for building executable, predictable, and maintainable APIs.'
slug: api-design-patterns
date: 2024-08-26T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Book
  - API
  - Design
translation_key: api-design-patterns
---



[APIデザイン・パターン](https://amzn.to/3AD3Vhr)を読んだ。

APIの設計原則に始まり、豊富なAPIのデザイン・パターンを紹介している。

- リソース指向
  - 単一のリソースに対して、標準メソッド（ex, Create/Get/List/Delete/Update）をかけ合わせて処理を決める
- RPC指向
  - メソッドを指定して特定の手続きを呼び出して処理を決める
- API設計を良くするため観点
  - 実行可能であること
  - 表現力があること
  - シンプルであること
  - 予測可能であること

APIのデザイン・パターンは、APIの設計や運用をしていく中で直面しそうな課題に対するアプローチとして参考になりそうなものが多い。設計に迷ったら参照したい。
