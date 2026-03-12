---
title: New Relic実践入門 第2版 オブザーバビリティの基礎と実現
description: New Relic実践入門 第2版 オブザーバビリティの基礎と実現について調査・整理したメモ。基本概念と重要ポイントを解説します。
slug: new-relic-observability-introduction
date: 2024-08-10T00:00:00Z
author: bmf-san
categories:
  - ツール
tags:
  - New Relic
  - 書評
translation_key: new-relic-observability-introduction
books:
  - asin: "4798184500"
    title: "New Relic実践入門 第2版 オブザーバビリティの基礎と実現"
---


[New Relic実践入門 第2版 オブザーバビリティの基礎と実現](https://amzn.to/4cmAsFm)を読んだ。

New Relicの機能を一通り網羅している。活用方法についても半分近くページが割かれているので、実践方法について知ることができる。

NRQLの仕様とかも書いてあると良いなと思ったが、足りないところはドキュメントで補完したり、実際に触ってみるのが早そう。


- テレメトリデータについて次のように定義している
  - メトリクス
    - 定期的にグループ化または収集された測定値の集合
    - 特定期間のデータ集計
  - ログ
    - 特定のコードが実行されたときにシステムが生成するテキスト行
  - トレース
    - イベントやトランザクションの連携する状態   
