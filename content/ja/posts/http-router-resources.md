---
title: HTTP Routerの自作で参考にした資料
description: "自作のHTTPルーター「goblin」の開発で参考にした、各種の資料や自分のブログ記事をまとめた記事です。radix木やトライ木といったデータ構造、ベンチマーク比較など、GoでHTTPルーターを自作する際に役立つ情報を、分散していたものを集約して紹介します。"
slug: http-router-resources
date: 2023-10-30T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - router
  - HTTP
  - URLルーティング
  - リンク集
translation_key: http-router-resources
draft: false
---


[goblin](https://github.com/bmf-san/goblin)の開発で参考にした各種ソースや自分の記事をリストアップする。

いくつかの自分のブログ記事に参考リンクを貼ったりしていたせいで分散していたので集約した。

# 参考
HTTP Routerの開発で参考にした資料リスト。

## GitHub
- [jba/muxpatterns](https://github.com/jba/muxpatterns)
- [importcjj/trie-go](https://github.com/importcjj/trie-go)
- [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
- [gorilla/mux](https://github.com/gorilla/mux)
- [gowww/router](https://github.com/gowww/router)
- [go-chi/chi](https://github.com/go-chi/chi)
- [go-ozzo/ozzo-routing](https://github.com/go-ozzo/ozzo-routing)
- [nissy/bon](https://github.com/nissy/bon)
- [nissy/mux](https://github.com/nissy/mux)
- [ytakano/radix_tree](https://github.com/ytakano/radix_tree)
- [kkdai/radix](https://github.com/kkdai/radix)
- [MarkBaker/Tries](https://github.com/MarkBaker/Tries)
- [razonyang/routing](https://github.com/razonyang/routing)
- ~~ethereum/wiki - [Japanese] Patricia Tree~~
- [neo-nanikaka - CommonPrefixTrieRouter.php](https://gist.github.com/neo-nanikaka/c2e2f7742b311696d50b)
- [golang/go - proposal: net/http: enhanced ServeMux routing](https://github.com/golang/go/issues/61410)

## ブログ
- [blog.merovius.de - How to not use an http-router in go](https://blog.merovius.de/posts/2017-06-18-how-not-to-use-an-http-router/)
- [medium.com/@agatan - HTTPサーバとcontext.Context](https://medium.com/@agatan/http%E3%82%B5%E3%83%BC%E3%83%90%E3%81%A8context-context-7211433d11e6)
- [devpixiv.hatenablog.com - 
PHPで高速に動作するURLルーティングを自作してみた](https://devpixiv.hatenablog.com/entry/2015/12/13/145741)
- [devpixiv.hatenablog.com - PHPで高速に動作するURLルーティングを自作してみた](https://devpixiv.hatenablog.com/entry/2015/12/13/145741)
- [kuune.org - 
世界最速だった URL ルーターをリリースしました](https://kuune.org/text/2014/06/12/denco/)
- [takao.blogspot.com - JavaでPatriciaTrieを実装してみた](https://takao.blogspot.com/2012/03/patriciatrie.html)
- [dankogai.livedoor.blog - algorithm - Patricia Trie (Radix Trie) を JavaScript で](https://dankogai.livedoor.blog/archives/51766842.html)
- [persol-pt.github.io - 勉強会[http request multiplexerと文字列マッチング]](http://web.archive.org/web/20201109013417/https://persol-pt.github.io/posts/tech-workshop1222/)
- [atmarkit.itmedia.co.jp - データ構造の選択次第で天国と地獄の差](https://atmarkit.itmedia.co.jp/ait/articles/0809/01/news163_3.html)
- [www.sb.ecei.tohoku.ac.jp - 基本データ構造：木構造のたどり方（アーカイブ）](https://web.archive.org/web/20250811134848/http://www.sb.ecei.tohoku.ac.jp/lab/wp-content/uploads/2012/11/2012_d12.pdf)
- [noranuk0.hatenablog.com - PHPでフレームワークを使わずURLのルーティングをいい感じにやる](https://noranuk0.hatenablog.com/entry/2018/01/20/114933)
- [reiki4040.hatenablog.com - golangのHTTPサーバを構成しているもの](https://reiki4040.hatenablog.com/entry/2017/03/01/212647)
- [qiita.com/immrshc - 【Go】net/httpパッケージを読んでhttp.HandleFuncが実行される仕組み](https://qiita.com/immrshc/items/1d1c64d05f7e72e31a98)

## ドキュメント
- [urlpattern.spec.whatwg.org](https://urlpattern.spec.whatwg.org/)
  - whatwgが提唱するURLPatternの標準
- [developer.mozilla.org - URL Pattern API](https://developer.mozilla.org/en-US/docs/Web/API/URL_Pattern_API)
  - MDNで実験的に実装されているURL Pattern APIの仕様

## ツール
- [www.cs.usfca.edu - Radix Tree](https://www.cs.usfca.edu/~galles/visualization/RadixTree.html)

# 執筆した記事
[bmf-tech.com](https://bmf-tech.com/)に投稿している記事。

- [URLルーティングをつくる　エピソード1](/ja/posts/creating-url-routing-episode-1/)
- [URLルーティングをつくる　エピソード2](/ja/posts/creating-url-routing-episode-2/)
- [URLルーティング自作入門　エピソード１](/ja/posts/creating-url-routing-episode-3/)
- [URLルーティング自作入門　エピソード２](/ja/posts/url-routing-introduction-episode-2/)
- [GolangでgoblinというURLルーターを自作した](/ja/posts/goblin-url-router-in-golang/)
- [GoのHTTP Routerを比較するベンチマーカーを実装した](/ja/posts/http-router-benchmark-go/)
- [Goで始めるコードのパフォーマンス改善](/ja/posts/go-performance-improvement/)
- [net／httpでつくるHTTPルーター自作入門](/ja/posts/building-http-router-with-net-http/)
- [自作ルーティングをアップデートした](/ja/posts/update-custom-routing/)
- [GolangでgoblinというURLルーターを自作した](/ja/posts/goblin-url-router-in-golang/)
- [GolangのHTTPサーバーのコードリーディング](/ja/posts/golang-http-server-code-reading-2019/)
- [Golangでトライ木を実装する](/ja/posts/golang-tries-implementation/)

# 執筆した本
- [net／httpでつくるHTTPルーター 自作入門
](https://zenn.dev/bmf_san/books/3f41c5cd34ec3f)

