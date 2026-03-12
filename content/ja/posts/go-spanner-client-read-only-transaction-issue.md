---
title: goのspannerクライアントのReadOnlyTransactionでハマった
description: goのspannerクライアントのReadOnlyTransactionでハマったについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: go-spanner-client-read-only-transaction-issue
date: 2021-02-08T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Google Cloud Platform
  - Spanner
  - Tips
translation_key: go-spanner-client-read-only-transaction-issue
---


# 概要
[pkg.go.dev - cloud.google.com/go/spanner](https://pkg.go.dev/cloud.google.com/go/spanner)で`ReadOnlyTransaction`を使ったときにハマったところについてメモ。

# 何が起きたか？
数万件のデータを複数回のリクエストに分けて処理するようなバッチ処理のコードを書いていた。
`ReadOnlyTransaction`を使った処理を以下のように書いていた。　

```golang
for {
    // 〜略〜

    // cは*spanner.Client
    iter := c.ReadOnlyTransaction().Query(ctx, stmt)
    defer iter.Stop()

    // 〜略〜
}
```

一見問題なさそうに見えたのでバッチ処理を走らせていたのだが、特定件数を超えると処理が止まる問題が発生した。

# 原因
spannerのgoクライアントにはセッション管理の仕組みがあるのだが、トランザクションの終了処理が漏れていたことにより、セッションプールが枯渇、リクエストがタイムアウトしていたらしい。
内部的にはセッション管理の仕組みがReadOnlyTransactionの実行をブロックしているような形になってしまっていたらしい。

# 対応
トランザクションの終了処理を呼び出すように変更する。

```golang
for {
    // 〜略〜

    // cは*spanner.Client
    tx := c.ReadOnlyTransaction()
    defer tx.Close()
    iter := tx.Query(ctx, stmt)
    defer iter.Stop()

    // 〜略〜
}
```

トランザクションの終了処理がないとトランザクション実行の度に新規セッションが生成され、セッションプールが枯渇してしまう。
処理が毎回同じ件数で止まっていたのは、`SPANNER_SESSION_POOL_MAX_OPEND`の値制限に引っかかったからの模様。計算してみると帳尻が合う。

# 対策
ドキュメントをちゃんと読むということ以外には、ツールを使った解決方法もある。
[github.com - gcpug/zagane](https://github.com/gcpug/zagane)

後はGCPのモニタリングでcloudspannerのsession countをウォッチするというのも有りかもしれない。

# 参考
- [cloud.google.com - セッション](https://cloud.google.com/spanner/docs/sessions?hl=ja)
- [chidakiyo.hatenablog.com - Go で Spanner とよろしくやるためにガチャガチャやっている話](https://chidakiyo.hatenablog.com/entry/2020/12/14/go-spanner-tools)
- [medium.com - 詳解 google-cloud-go/spanner — セッション管理編](https://medium.com/google-cloud-jp/%E8%A9%B3%E8%A7%A3-google-cloud-go-spanner-%E3%82%BB%E3%83%83%E3%82%B7%E3%83%A7%E3%83%B3%E7%AE%A1%E7%90%86%E7%B7%A8-d805750edc75)
- [medium.com - 詳解 google-cloud-go/spanner — トランザクション編](https://medium.com/google-cloud-jp/%E8%A9%B3%E8%A7%A3-google-cloud-go-spanner-%E3%83%88%E3%83%A9%E3%83%B3%E3%82%B6%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3%E7%B7%A8-6b63099bd7fe)
