---
title: '分散SQLクエリエンジンTrino徹底ガイド '
description: "『分散SQLクエリエンジンTrino徹底ガイド』の書評。分散SQLクエリエンジンTrinoを解説する一冊。ANSI SQL対応やフェデレーテッドクエリにより異なるシステムのデータを横断的に扱い、OLAP向けに分散してリソースを拡縮する仕組みなど、従来のデータウェアハウスとは異なるアプローチを学べる。"
slug: trino-sql-query-engine-guide
date: 2024-10-24T00:00:00Z
author: bmf-san
categories:
  - データベース
tags:
  - Trino
  - 書評
translation_key: trino-sql-query-engine-guide
books:
  - asin: "4798071676"
    title: "分散SQLクエリエンジンTrino徹底ガイド"
draft: false
---


[分散SQLクエリエンジンTrino徹底ガイド](https://amzn.to/3BXhPeQ)を読んだ。

- SQLクエリエンジン
- ANSI SQLをサポート
- フェデレーテッドクエリ
  - 同じSQLで異なるシステムからデータベースとスキーマを参照および使用する 
- データベースではない
- OLTPを処理するために設計されていない
  - OLAP向けなので、あくまで性能目標はOLAP水準なのではなかろうか 　
- 分散システムとしてコンピューティングリソースを拡大・縮小できる点はNew SQLの性質に似ている
- ビッグデータの従来型アプローチ（様々なクエリ言語やツール、運用や維持費などのコストの掛かるデータウェアハウスを必要とするアプローチ）とは異なる形でデータソースを扱う
