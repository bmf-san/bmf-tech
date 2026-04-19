---
title: 自作ルーティングをアップデートした
description: 自作ルーティングをアップデートした
slug: update-custom-routing
date: 2021-06-18T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
  - OSS
  - router
translation_key: update-custom-routing
---


# 概要
最近、自作ルーティングの[goblin](https://github.com/bmf-san/goblin)をアップデートしたのでその記録を書き残しておく。

以下は過去ルーティングについて書き残した記事。他にも実装検討フェーズの記事があるが、内容があまり良くないので割愛。

- [URLルーティング自作入門　エピソード１](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E8%87%AA%E4%BD%9C%E5%85%A5%E9%96%80%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%89%EF%BC%91)
- [URLルーティング自作入門　エピソード２](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E8%87%AA%E4%BD%9C%E5%85%A5%E9%96%80%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%89%EF%BC%92)
- [GolangのHTTPサーバーのコードリーディング](https://bmf-tech.com/posts/Golang%E3%81%AEHTTP%E3%82%B5%E3%83%BC%E3%83%90%E3%83%BC%E3%81%AE%E3%82%B3%E3%83%BC%E3%83%89%E3%83%AA%E3%83%BC%E3%83%87%E3%82%A3%E3%83%B3%E3%82%B0)
- [Introduction to URL router from scratch with Golang](http://web.archive.org/web/20251207231313/https://dev.to/bmf_san/introduction-to-url-router-from-scratch-with-golang-3p8j)

# 何が変わったのか？
基本的な動作をするバージョンを[1.0.0](https://github.com/bmf-san/goblin/releases/tag/1.0.0)としてリリースしていた。
実際に自分で利用する中で、バグを見つけたり、機能の物足りなさを感じて、後方互換性のない変更を何回か経て（行き当たりばったりな実装をしていたツケが回った）、現在は[5.0.1](https://github.com/bmf-san/goblin/releases/tag/5.0.1)が最新のバージョンになっている。

具体的には、ミドルウェアの機能をサポートするようになったことが一番の変更点で、それに伴い内部のデータ構造を見直したり、DSLを見直したり、バグを改修したりした。

# なぜミドルウェアをサポートしたのか？
ミドルウェアはルーターを利用する側で自由に対応できると考えていたが、実際には制約があった。

利用側でミドルウェアを実装しても、ルーティングのマッチングの処理（≒パスとHTTP
メソッドが登録済みのルーティングにマッチするかどうかの処理）が完了した後にミドルウェアの処理が実行されるような形になるため、HTTPメソッドのマッチング前にミドルウェアを適用したいようなケースに対応できないという制約を設けてしまっていた。

これはPreflightのリクエスト（CORS対応など）を捌きたいときに不便なため、根本的に解決するために、ミドルウェアをサポートする判断をした。

そのようなケースを考慮する上で厄介だったのが、ルーティングが内部的に持つデータ構造で、パスとHTTPメソッドの一致を前提とするようなデータ構造になってしまっていたため、そこの見直しからする必要があった。

なので、データ構造を以下のように変更をして、ミドルウェアのサポートを実装した。

Before
![Based on trie tree](/assets/images/posts/update-custom-routing/70862745-7148e180-1f83-11ea-85d3-2cd8fb4db0d3.png "Based on trie tree")

After
![after](/assets/images/posts/update-custom-routing/117675761-d4c25780-b1e7-11eb-9ec7-e78ac0ce142b.png)


# ベンチマーク
静的なルーティングだけ対応したベンチマークを書いていたが、他のライブラリとの比較を動的なルーティングのテストと合わせて比較したみたかったので一番充実していそうな[github.com - julienschmidt/go-http-routijng-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark)を使ってベンチマークテストを実施してみた。


以下は最新のスコア。
```
#GithubAPI Routes: 203
   Goblin: 80864 Bytes

#GPlusAPI Routes: 13
   Goblin: 7856 Bytes

#ParseAPI Routes: 26
   Goblin: 8688 Bytes

#Static Routes: 157
   Goblin: 34488 Bytes

goos: darwin
goarch: amd64
pkg: github.com/julienschmidt/go-http-routing-benchmark
cpu: Intel(R) Core(TM) i5-8210Y CPU @ 1.60GHz
BenchmarkGoblin_Param             738289              1964 ns/op             128 B/op          4 allocs/op
BenchmarkGoblin_Param5            754988              1920 ns/op             368 B/op          6 allocs/op
BenchmarkGoblin_Param20            56145             23260 ns/op            3168 B/op         58 allocs/op
BenchmarkGoblinWeb_ParamWrite     304082              4610 ns/op             648 B/op         11 allocs/op
BenchmarkGoblin_GithubStatic     1156518              2745 ns/op             128 B/op          4 allocs/op
BenchmarkGoblin_GithubParam       125570              9985 ns/op             816 B/op         15 allocs/op
BenchmarkGoblin_GithubAll           2232            622376 ns/op           49424 B/op       1018 allocs/op
BenchmarkGoblin_GPlusStatic      1000000              1298 ns/op              80 B/op          3 allocs/op
BenchmarkGoblin_GPlusParam        417717              2893 ns/op             664 B/op         11 allocs/op
BenchmarkGoblin_GPlus2Params      274990              4551 ns/op             824 B/op         15 allocs/op
BenchmarkGoblin_GPlusAll           95580             14536 ns/op            2208 B/op         57 allocs/op
BenchmarkGoblin_ParseStatic      1651083               707.0 ns/op           128 B/op          4 allocs/op
BenchmarkGoblin_ParseParam        413840              2876 ns/op             728 B/op         12 allocs/op
BenchmarkGoblin_Parse2Params      260120              4119 ns/op             808 B/op         15 allocs/op
BenchmarkGoblin_ParseAll           54518             21692 ns/op            4656 B/op        120 allocs/op
BenchmarkGoblin_StaticAll          26689             46104 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/julienschmidt/go-http-routing-benchmark      37.270s
```

goblinのベンチマーク対応を追加したPRを投げてみている。
[github.com - julienschmidt/go-http-routing-benchmark Add a new router goblin #97](https://github.com/julienschmidt/go-http-routing-benchmark/pull/97)

# 所感
ようやく人並みのルーター？になったような気がする。
まだまだ改善点はあるので今後も継続的にメンテしていく。


