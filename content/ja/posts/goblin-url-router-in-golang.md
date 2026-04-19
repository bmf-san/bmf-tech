---
title: GolangでgoblinというURLルーターを自作した
description: GolangでgoblinというURLルーターを自作した
slug: goblin-url-router-in-golang
date: 2020-01-26T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
  - URLルーティング
  - router
translation_key: goblin-url-router-in-golang
---


# 概要
GolangでURLルーターを自作したので実装するまでの過程をメモしておく。

# 準備
URLルーターを実装する際に行った下準備をまとめる。

## データ構造とアルゴリズム
URLをどのようにマッチングさせるか、というロジックについて検討する。

多くのライブラリでは、データ構造として木構造がよく扱われているので、どんな種類の木構造を採用するかを考えてみた。

文字列探索に特化した木の中で、時間的・メモリ的計算量がよりベストなものを選定しようとすると、基数木というのが良さそうに見えるので最初はそれを採用しようとしていたのだが、実装が難し過ぎて挫折をした。

もう少し身近でシンプルなものをということでトライ木を採用することにした。

## `net/http`のコードリーディング
`net/http`が持つマルチプレクサの拡張として実装を行うため、内部の仕組みについてある程度理解しておく必要がある。

[GolangのHTTPサーバーのコードリーディング](https://bmf-tech.com/posts/Golang%E3%81%AEHTTP%E3%82%B5%E3%83%BC%E3%83%90%E3%83%BC%E3%81%AE%E3%82%B3%E3%83%BC%E3%83%89%E3%83%AA%E3%83%BC%E3%83%87%E3%82%A3%E3%83%B3%E3%82%B0)を参照。

## 色んなrouterの実装を読む
参考>リポジトリ参照。

## その他
過去URLルーティングについてまとめた記事。

[bmf-tech.com/posts/tags/URLルーティング](https://bmf-tech.com/posts/tags/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0)

# 実装

[github.com - goblin](https://github.com/bmf-san/goblin)を参照。

基本はトライ木を使いやすい形に変えていくだけではあるのだが、パスパラメータの扱いに何度か格闘した。
正規表現の対応はそれほど面倒ではなく、DSLを用意してあげるだけでどうにかなるので、DSLの扱いにセンスが問われる。

実装過程では、テストを並行して書いたり、ステップ実行のデバッグを繰り返したりして、データ構造が常にどう変化しているかキャッチしながらやっていたので、段々脳内デバッグ力が上がっていたような気がする。

普段あまり書かないようなロジックなので、コーディングの良いトレーニングになったのは間違いなさそう。

今後の課題はリポジトリに起票してあるissueの通りで、暇な時にでもブラッシュアップを重ねようかと思っている。

# 参考
## リポジトリ
設計とか実装の参考にさせてもらったリポジトリ

- [github - importcjj/trie-go](https://github.com/importcjj/trie-go)
- [github - julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
- [github - gorilla/mux](https://github.com/gorilla/mux)
- ~~github - xhallix/go-router~~
- [github - gowww/router](https://github.com/gowww/router)
- [github - go-chi/chi](https://github.com/go-chi/chi)
- [github - go-ozzo/ozzo-routing](https://github.com/go-ozzo/ozzo-routing)

## その他
実装時に参照した記事。

- ~~How to not use an http-router in go~~
- ~~Goのhttp serverの雰囲気を理解する~~
- [HTTPサーバとcontext.Context](https://medium.com/@agatan/http%E3%82%B5%E3%83%BC%E3%83%90%E3%81%A8context-context-7211433d11e6)
- ~~golangでhttptestを使ってテストする
golang~~
- [PHPで高速に動作するURLルーティングを自作してみた](https://devpixiv.hatenablog.com/entry/2015/12/13/145741)
- [世界最速だった URL ルーターをリリースしました](https://kuune.org/text/2014/06/12/denco/)
