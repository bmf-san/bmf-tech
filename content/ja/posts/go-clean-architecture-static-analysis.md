---
title: "GoでClean Architectureのレイヤーを静的解析する"
slug: "go-clean-architecture-static-analysis"
date: 2022-09-04
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Golang"
  - "Clean Architecture"
draft: false
---

# 概要
GoでClean Architectureのレイヤーを静的解析する方法についてのメモ。

# go-cleanarchを使う
静的解析のツールを自作しても良かったが、簡単に導入できるツールがあったのでこちらを使ってみた。

[roblaszczak/go-cleanarch](https://github.com/roblaszczak/go-cleanarch) 

自作CMSの[gobel-api](https://github.com/bmf-san/gobel-api)に導入してみた。
cf. [PR](https://github.com/bmf-san/gobel-api/pull/74/files)

`go install github.com/roblaszczak/go-cleanarch@latest`でインストール。

レイヤーのネーミングがデフォルトと異なるので、オプションをつけてチェック実行。
`go-cleanarch -application usecase`

レイヤーの依存関係に違反するとチェックに引っかかってエラーになる。

エラーになると、`Uncle Bob is not happy.`と怒られる。

# 所感
そのうち自作したいと思っているが当分はこのツールにお世話になりたい。

アプリケーションの設計、構造を保守していくためにこういった静的解析ツールは早い段階から導入すべきだと思った。

Goなら静的解析のツールが実装しやすいと思うので、Clean Architectureに限らず、特定のレイヤーの構造を維持していきたいのであれば積極的にこういったツールを自作していくのが良さそう。

