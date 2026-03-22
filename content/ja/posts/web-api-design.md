---
title: Web APIの設計
description: "Web APIの設計"
slug: web-api-design
date: 2024-07-28T00:00:00Z
author: bmf-san
categories:
  - アーキテクチャ
tags:
  - API
  - OpenAPI
  - 設計
  - 書評
translation_key: web-api-design
books:
  - asin: "4798167010"
    title: "Web APIの設計"
---


[Web APIの設計](https://amzn.to/3y7dqEG)を読んだ。

API設計の視点を与えてくれる本。

API設計のインターフェースだけでなく、もっと手前の入り口から書かれている。

- API設計は、仕組みではなくユーザーが何をできるかに焦点を当てて設計すると良い。仕組みに焦点を当てると複雑化する
- APIのゴールを洗いだすアプローチとしてゴールキャンパスというフレームワークの解説があったが、良い方法だと思った
- APIバージョニングのパターンにはパス、ドメイン、クエリパラメータ、カスタムヘッダー、コンテントネゴシエーション、コンシューマの設定（コンシューマごとの設定をDBに持つ）がある
