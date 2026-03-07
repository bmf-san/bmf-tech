---
title: 自作HTTPルーターから新しいServeMuxへ
slug: custom-http-router-to-new-servemux
date: 2024-04-27T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
  - router
translation_key: custom-http-router-to-new-servemux
---


# 概要
これまで[goblin](https://github.com/bmf-san/goblin)という自作HTTPルーターを自分のアプリケーションで使っていたが、Go1.22でServeMuxの機能が拡張されてからはServeMuxを使うようになった。

この記事では、Go1.22で追加されたServeMuxの機能や性能について整理し、これからのGoのHTTPルーター選定について考えてみる。

# Go1.22で追加されたServeMuxの機能
Go1.22rcの公開時にServeMuxの新機能について調べていたが、もう少し詳細を追ってみようと思う。

cf. [Go1.22rcで変更されるServeMuxの仕様](https://bmf-tech.com/posts/Go1.22rc%e3%81%a7%e5%a4%89%e6%9b%b4%e3%81%95%e3%82%8c%e3%82%8bServeMux%e3%81%ae%e4%bb%95%e6%a7%98)

以下の参考情報を元にServeMuxの新機能について整理する。

- リリースノート
  - [Go 1.22 Release Notes - Enhanced routing patterns](https://tip.golang.org/doc/go1.22#enhanced_routing_patterns)
    - ServeMuxのパターンが拡張され、メソッドとワイルドカード（動的なパスパラメータ。ex. /items/{id}）を受け入れるようになったと記載されている
- pkg.go.dev
  - [pkg.go.dev - net/http](https://pkg.go.dev/net/http)
- go.dev
  - [go.dev - Routing Enhancements for Go 1.22](https://go.dev/blog/routing-enhancements)
    - ServeMuxの新機能について仕様が記載されている
- ディスカッション
  - [net/http: add methods and path variables to ServeMux patterns](https://github.com/golang/go/discussions/60227)
    - ServeMuxの機能拡張についてディスカッション
- プロポーザル
  - [net/http: enhanced ServeMux routing](https://github.com/golang/go/issues/61410)
    - ServeMuxの機能拡張についてプロポーザル

## HTTPメソッドによるルーティングの定義
HTTPメソッドを含んだパスを指定することにより、HTTPメソッドによるルーティングの定義が可能になった。ServeMuxを使うときはhandler内でHTTPメソッドの条件分岐を書く必要がなくなった。

```go
http.HandleFunc("GET /items", handleItems)
```

HTTPメソッドの定数（ex. http.MethodGet）が用意されているので、それを使う形が良いのではと思ったが、おそらく下位互換性を考慮して既存メソッドのシグネチャを崩さないようにするためにこのような形なのではないかと思っている。

あるいは単にHTTPのリクエスト形式に合わせただけなのかもしれない。

## ワイルドカードによるルーティングの定義
ワイルドカード（`{pathVal}`）を使ったパスを指定することにより、ワイルドカードによるルーティングの定義が可能になった。

```go
// GET /items/1や/items/fooなどにマッチする
http.HandleFunc("GET /items/{id}", handleItems)
```

パスパターンとしては以下のような形式で指定することができる。Go1.22以前であってもホスト名も指定することができる。（割と最近知った...）

```
[METHOD ][HOST]/[PATH]
```

ワイルドカードにマッチした値は、http.Requestの[PathValueメソッド](https://pkg.go.dev/net/http#Request.PathValue)を使うことで取得することができる。

```go
func handleItems(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	// idにはワイルドカードにマッチした値が入る
}
```

また、マルチワイルドカードによるルーティングの定義も可能である。

```go
// GET /items/1や/items/1/2、/itema/1/2/3などにマッチする
http.HandleFunc("GET /items/{id...}", handleItems)
```

パスを完全一致でマッチさせたい場合は`{$}`を使う。ルート（`/`）の定義をしたいときなどに留意しておきたい。

```go
http.HandleFunc("GET /{$}", handleIndex)
```

余談だが、サードパーティのルーターでは`*`をワイルドカードとしてパターンに使っていることがある。そのイメージに引きづられてしまうので個人的にはパスパラメータと呼称したいと思ったり...（気にするほどのことではないが。）

# 新しいServeMuxでの注意点
## HTTPメソッドを使ったパターンの定義
多くのサードパーティのHTTPルーターでは、HTTPメソッドがメソッドであることが多い。

```go
// ServeMux
http.HandleFunc("GET /items/{id}", handleItems)

// サードパーティの例
mux.Get("/items", handleItems)
```

ServeMuxではHTTPメソッドはパターンに内包されるためタイポなどが検知し辛いという些細な問題があるかもしれない。linterでチェックできればいいがどうだろう。自分で静的解析ツールを作るにはちょうど良い題材かもしれない。そのうち何かしら手が打たれそうな気はする。

## 優先順位ルール
ワイルドカードを使ったルーティングの定義においては、優先順位に留意したい。

cf.
- [pkg.go.dev - Precedence](https://pkg.go.dev/net/http#hdr-Precedence)
- [go.dev - Routing Enhancements for Go 1.22 precedence](https://go.dev/blog/routing-enhancements#precedence)

ServeMuxでは次のようなルーティングのパターンでも定義することができる。

```
// どちらにもマッチする場合は前者が優先される
/items/new
/items/{id}

// どちらにもマッチする場合は後者が優先される
/items/{id...}
/items/{id}/category/{name}
```

このようなパターンの場合、想定するリクエストがどのパターンにマッチするのか注意しておく必要がある。

世の中のHTTPルーターにはこのようなパターンでの重複を許さないものもあれば、許すものもあるが、ServeMuxは後者に該当する。

一方で、次のような競合が生じるケースについては、ServeMuxが競合を検知しpanicを発生させる。

```
// どちらにもマッチするケースもあれば、片方のみマッチするケースもある
/items/{id}
/{category}/items

// どちらにもマッチする場合
/items/{id}
/items/{name}
```

競合の場合、エラーは早い段階（テストやサーバー起動時等）で検知できると思われるため、重複よりは厄介ではないと思う。

ちなみに自作している[goblin](https://github.com/bmf-san/goblin)では、先に登録されたパターンが優先されるような先勝ちな仕様になっている。（ちゃんと設計できていないので実は結構煩雑...）

ServeMuxの競合検知については以下の記事が詳しい。

cf. [rhumie.github.io - ServeMuxの競合検知と性能](https://rhumie.github.io/go122party/1)

サードパーティのHTTPルーターであっても、競合検知は考慮されており、例えば[httprouter](https://github.com/julienschmidt/httprouter)だと、1つのパターンに一致するかしないかのどちらかになるという仕様になっている。

> Only explicit matches: With other routers, like http.ServeMux, a requested URL path could match multiple patterns. Therefore they have some awkward pattern priority rules, like longest match or first registered, first matched. By design of this router, a request can only match exactly one or no route. As a result, there are also no unintended matches, which makes it great for SEO and improves the user experience.

cf. https://github.com/julienschmidt/httprouter?tab=readme-ov-file#features

HTTPルーターのパターンのマッチにおける優先順位の仕様はHTTPルーターの品質を左右する重要なポイントであるので、これがちゃんと設計されているのは安心感がある。

## 後方互換性
Go1.22とGo1.21において一部後方互換性が保たれないケースがある。

cf. [pkg.go.dev - Compatibility](https://pkg.go.dev/net/http#hdr-Compatibility)

その場合は、GODEBUG環境変数で`httpmuxgo121=1`を設定することでGo1.21の挙動に戻すことができる。

# 内部実装コードリーディング
## ルーティングのパターン登録処理
- [register](https://cs.opensource.google/go/go/+/master:src/net/http/server.go;l=2735;drc=960fa9bf66139e535d89934f56ae20a0e679e203;bpv=1;bpt=1)
- [parsePattern](https://cs.opensource.google/go/go/+/master:src/net/http/pattern.go;drc=960fa9bf66139e535d89934f56ae20a0e679e203;l=84)

下記はServeMux構造体の定義である。

```go
// see: https://cs.opensource.google/go/go/+/master:src/net/http/server.go;l=2439;drc=960fa9bf66139e535d89934f56ae20a0e679e203;bpv=1;bpt=1

type ServeMux struct {
	mu       sync.RWMutex
	tree     routingNode
	index    routingIndex
	patterns []*pattern  // TODO(jba): remove if possible
	mux121   serveMux121 // used only when GODEBUG=httpmuxgo121=1
}
```

パスがノードになるように木構造（tree）を生成しており、HTTPルーターでは一般的なデータ構造になっていると思われる。

indexとpatternsはパターンの競合検知に使われるデータとなっている。

## ルーティングのマッチング処理
- [ServeHTTP](https://cs.opensource.google/go/go/+/master:src/net/http/server.go;l=3132;drc=960fa9bf66139e535d89934f56ae20a0e679e203)
- [findHandler](https://cs.opensource.google/go/go/+/master:src/net/http/server.go;drc=960fa9bf66139e535d89934f56ae20a0e679e203;bpv=1;bpt=1;l=2516?q=ServeHTTP&ss=go%2Fgo)
- [matchOrRedirect](https://cs.opensource.google/go/go/+/master:src/net/http/server.go;drc=960fa9bf66139e535d89934f56ae20a0e679e203;bpv=1;bpt=1;l=2578?q=ServeHTTP&ss=go%2Fgo)
- [match](https://cs.opensource.google/go/go/+/master:src/net/http/routing_tree.go;drc=960fa9bf66139e535d89934f56ae20a0e679e203;bpv=1;bpt=1;l=115)
- [matchMethodAndPath](https://cs.opensource.google/go/go/+/master:src/net/http/routing_tree.go;l=130;drc=960fa9bf66139e535d89934f56ae20a0e679e203)
- [matchPath](https://cs.opensource.google/go/go/+/master:src/net/http/routing_tree.go;l=152;drc=960fa9bf66139e535d89934f56ae20a0e679e203)

読み慣れていない場合は、前提としてHTTPサーバーのコードリーディングをしておくと良いかもしれない。

cf. [GolangのHTTPサーバーのコードリーディング](https://bmf-tech.com/posts/Golang%E3%81%AEHTTP%E3%82%B5%E3%83%BC%E3%83%90%E3%83%BC%E3%81%AE%E3%82%B3%E3%83%BC%E3%83%89%E3%83%AA%E3%83%BC%E3%83%87%E3%82%A3%E3%83%B3%E3%82%B0)

# ServeMuxとサードパーティとの比較
ServeMuxとサードパーティとの比較を自作ベンチマーカーにて行った。

自作ベンチマーカーは以前実装したものを利用している。詳細については以下のリンクを参照されたい。

- [bmf-san/go-router-benchmark](https://github.com/bmf-san/go-router-benchmark)
- [GoのHTTP Routerを比較するベンチマーカーを実装した](https://bmf-tech.com/posts/Go%e3%81%aeHTTP%20Router%e3%82%92%e6%af%94%e8%bc%83%e3%81%99%e3%82%8b%e3%83%99%e3%83%b3%e3%83%81%e3%83%9e%e3%83%bc%e3%82%ab%e3%83%bc%e3%82%92%e5%ae%9f%e8%a3%85%e3%81%97%e3%81%9f)

ベンチマーク結果は、[<Public>go-router-benchmark](https://docs.google.com/spreadsheets/d/1DrDNGJXfquw_PED3-eMqWqh7qbCTVBWZtqaKPXtngxg/edit#gid=1913830192)にて公開している。

上記のベンチマーカーでは、パスパターンのマッチングのみ計測しており、パスパターンの登録は計測対象外としている。

## ベンチマーク結果
ベンチマーク結果からパフォーマンスが良いと思われるEchoやgin、httprouterと比較すると差を感じる部分があるものの、全体的には平均以上のパフォーマンスを持っているように見える。

顕著だったのは、パスパラメータの数が増えるについて性能劣化が見られた点である。

上位のHTTPルーターはこの劣化を抑える工夫がされているように感じる。

パスパラメータをいくつも使うようなケースはあまり多くないと思うので、この点はあまり気にする必要はないかとは思う。

## goblinとServeMuxの比較
静的なルーティングのテストケースにおいてはgoblinが勝っているが、動的なルーティングについてはServeMuxが勝っている。

パスパラメータの数が増えることによる性能劣化はgoblinのほうが抑え気味で頑張っている部分が多く見受けられる。

繰り返しにはなるがパスパラメータをいくつも使うようなケースはあまり多くはないので、大した実用性のある性能差というほどではないかと思う。

# ServeMuxのパフォーマンスについて
ServeMux（の実装者）はHTTPルーターのパフォーマンスを次のように考えているらしい。

> Implementation is out of scope for this discussion/proposal. I think we'd be happy to have a more complex implementation if it could be demonstrated that the current one actually affects latency or CPU usage. For typical servers, that usually access some storage backend over the network, I'd guess the matching time is negligible. Happy to be proven wrong.

cf. https://github.com/golang/go/discussions/60227#discussioncomment-5932822 より引用

その他関連コメントも記載。

cf. https://github.com/golang/go/issues/61410#issuecomment-1867191476
cf. https://github.com/golang/go/issues/61410#issuecomment-1867485864
cf. https://github.com/golang/go/issues/61410#issuecomment-1868615273

性能を一つの特徴としているサードパーティのHTTPルーターでは、複雑な木構造のアルゴリズム（ex. メモリ効率が最適化されたRadix Tree）を採用しているケースが多い。

ServeMuxの実装においては、余程最悪なデータ構造を採用しない限り、レイテンシやCPU使用率に大きな影響を及ぼすことはないので、複雑なデータ構造やアルゴリズムを採用しないという思想らしい。

反証というわけではないが、[gorilla/mux](https://github.com/gorilla/mux)は人気のある（Starが多い）サードパーティのHTTPルーターの中では比較的ベンチマーク結果が見劣りするが、多くのユーザーに使われている。

パスパターンのマッチングだけではなく、パスパターンの登録についての性能について言及しているコメントもある。

> Registration time is potentially more of an issue. With the precedence rules described here, checking a new pattern for conflicts seems to require looking at all existing patterns in the worst case. (Algorithm lovers, you are hereby nerd-sniped.) That means registering n patterns takes O(n2) time in the worst case. With the naive algorithm that loops through all existing patterns, that "worst case" is in fact every (successful) case: if there are no conflicts it will check every pattern against every other, for n(n-1)/2 checks. To see if this matters in practice, I collected all the methods from 260 Google Cloud APIs described by discovery docs, resulting in about 5000 patterns. In reality, no one server would serve all these patterns—more likely there are 260 separate servers—so I think this is a reasonable worst-case scenario. (Please correct me if I'm wrong.) Using naive conflict checking, it took about a second to register all the patterns—not too shabby for server startup, but not ideal. I then implemented a simple indexing scheme to weed out patterns that could not conflict, which reduced the time 20-fold, to 50 milliseconds. There are still sets of patterns that would trigger quadratic behavior, but I don't believe they would arise naturally; they would have to be carefully (maliciously?) constructed. And if you are being malicious, you are probably only hurting yourself: one writes patterns for one's own server, not the servers of others. If we do encounter real performance issues, we can index more aggressively.

cf. https://github.com/golang/go/discussions/60227#discussioncomment-6204048 より引用

自分はHTTPルーターの性能について同様の考えを持っていたので、ServeMuxの性能に対する考えには同意できる。

[goblin](https://github.com/bmf-san/goblin)では、Trie Treeをベースにしたデータ構造を採用している。

Trie Treeよりもメモリ効率が良いRadix Treeを採用しなかったのは、単に複雑で理解や保守が大変そうと感じたという理由もあるが、複雑なデータ構造を採用するほど性能面でのメリットを享受できるだろうか？ということに疑問を持っていたからという理由もある。

goは言語思想からもシンプルさを追求すると思われるので、今後も複雑なデータ構造やアルゴリズムをServeMuxに採用するよりかは、現行のシンプルさを最適化していくような流れになるのではないだろうかと思う。（たぶん。）

ベンチマーク結果を見た感じではチューニングポイントはありそうなので、今後もっとスコアを伸ばしていくのではないかと思う。

その時は第2回天下一HTTPRouter武闘会を開催したい。

cf. [天下一HTTPRouter武闘会](https://speakerdeck.com/bmf_san/tian-xia-httprouterwu-dou-hui)

# ServeMuxと自作HTTPルーターを比較した結果から得られた学び
Go1.22のServeMuxの実装を眺めていて感じたことは、ルーティングのアルゴリズムがシンプルでありながらも性能が高いということだった。

余程選択を間違いない限り、ある程度性能が担保されるのだなぁと思った。（アルゴリズム素人の見解。）

[routing_tree_test.go](https://cs.opensource.google/go/go/+/refs/tags/go1.22.2:src/net/http/routing_tree_test.go)とか[pattern_test.go](https://cs.opensource.google/go/go/+/refs/tags/go1.22.2:src/net/http/pattern_test.go)のテストケースがスッキリしているのはデータ構造がシンプルだからだと思った。

[goblin](https://github.com/bmf-san/goblin)は結構悲惨で気にしているポイントだったりする。

[goblin](https://github.com/bmf-san/goblin)が採用しているTrie Tree自体はシンプルではあるが、その活用の仕方がシンプルじゃないので改善の余地はあるかもしれない。

ディスカッションやプロポーザルを眺めていて、データ構造やアルゴリズムの選択には期待する性能がどれほどなのか？という視点が大事だと改めて感じた。

性能を追求すればするほどシンプルさからかけ離れてしまうので、バランスが大事なのだろう。（要はバランスおじさん。）

# GoのHTTPルーター選定についての私見
これから新規にアプリケーションを開発する場合は、まずServeMuxを検討して、不足があればサードパーティを検討するという形が基本になるのではないかと思う。

一方で、既存のアプリケーションで利用しているHTTPルーターからServeMuxへ移行を検討するべきかどうかという点では、以下のような観点があると思う。

- サードパーティへの依存を減らしてできるだけ標準ライブラリに寄せたいか？
  - ワイルドカードが使えなかったためにサードパーティを渋々使っていたケースであれば積極的に乗り換えを検討したくなるのではないだろうか
- 使っているHTTPルーターのnet/httpとの互換性はどの程度か？
  - 独自のHandler定義やリクエストパラメータの取得方法などを提供しているものを利用している場合、移行に手間が掛かるかもしれない
- ServeMuxにはない機能や性能が必要か？
  - ミドルウェア周りや正規表現を使ったルーティング、グループ化などが必要であれば引き続きサードパーティを使うのが理にかなうかもしれない
- ルーティングの優先順位に独自のロジックがあるか？
  - テストで担保できれば問題ないが、移行する際の障壁の1つになるかもしれない

ServeMuxを使うか、サードパーティを使うか簡易的なフローチャートをつくってみた。

![フローチャート](https://github.com/bmf-san/bmf-san/assets/13291041/4bc81581-cdab-4fde-bb87-69f73511732f)

# 追記
[Go Conference 2024](https://gocon.jp/2024/)にて登壇してきた。

[speadkerdeck.com - 自作HTTPルーターから新しいServeMuxへ](https://speakerdeck.com/bmf_san/zi-zuo-httprutakaraxin-siiservemuxhe)

# 参考
- [bmf-tech.com/posts/tags/router](https://bmf-tech.com/posts/tags/router)
  - 過去投稿したrouter関連の全ての記事
- [net/http: add methods and path variables to ServeMux patterns](https://github.com/golang/go/discussions/60227)
- [net/http: enhanced ServeMux routing](https://github.com/golang/go/issues/61410)
- [Go 1.22 Release Notes - Enhanced routing patterns](https://tip.golang.org/doc/go1.22#enhanced_routing_patterns)
- [go.dev - Go 1.22 is released!](https://go.dev/blog/go1.22)
- [go.dev - Routing Enhancements for Go 1.22](https://go.dev/blog/routing-enhancements)
- [zenn.dev - Go 1.22の新しいrouterについて](https://zenn.dev/catatsuy/scraps/37e3b52bca7d13)
- [zenn.dev - Go 1.22のEnhanced ServeMuxに合わせて設計されたルーティングライブラリmichi](https://zenn.dev/sonatard/articles/831b761a27b230)
- [future-architect.github.io - Go1.22リリースパーティに「ServeMuxの競合検知と性能」というタイトルで登壇しました](https://future-architect.github.io/articles/20240408b/)
- [rhumie.github.io - ServeMuxの競合検知と性能](https://rhumie.github.io/go122party/1)
- [eli.thegreenplace.net - Better HTTP server routing in Go 1.22](https://eli.thegreenplace.net/2023/better-http-server-routing-in-go-122/)
- [shijuvar.medium.com - Building REST APIs With Go 1.22 http.ServeMux](https://shijuvar.medium.com/building-rest-apis-with-go-1-22-http-servemux-2115f242f02b)
- [www.calhoun.io - Go's 1.22+ ServeMux vs Chi Router](https://www.calhoun.io/go-servemux-vs-chi/)
- [www.alexedwards.net - Which Go router should I use? (with flowchart)](https://www.alexedwards.net/blog/which-go-router-should-i-use)
- [www.youtube.com - Why The Golang 1.22 HTTP Router Is Not Great](https://www.youtube.com/watch?v=agX6Ba2ODlw)
- [www.reddit.com - The proposal to enhance Go's HTTP router](https://www.reddit.com/r/golang/comments/15dvauk/the_proposal_to_enhance_gos_http_router/)
