---
title: 'レガシーコード改善ガイド: 保守開発のためのリファクタリング'
slug: legacy-code-improvement-guide
date: 2017-03-12T00:00:00Z
author: bmf-san
categories:
  - アーキテクチャ
tags:
  - レガシーコード
  - 書評
translation_key: legacy-code-improvement-guide
books:
  - asin: "4798116831"
    title: "レガシーコード改善ガイド: 保守開発のためのリファクタリング"
---


[レガシーコード改善ガイド: 保守開発のためのリファクタリング](https://amzn.to/4adL0FR)を読んだ。

レガシーコードと対峙するなら読んでおきたい本。

前提として、テストがある程度書ける一定の能力が求められる気がする。

  - 接合部(seam)とは、その場所を直接編集できなくても、プログラムの振る舞いを変えることができる場所
  - どの接合部にも許容点(enabling point)を持つ。許容点では、どの振る舞いを使うかを決定できる
