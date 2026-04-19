---
title: net／httpでつくるHTTPルーター自作入門
description: "Golangのnet/httpでHTTPルーター自作を解説。ルートマップ、URLパス解析、メソッド別ルーティング、パスパラメータの仕組みをステップバイステップで実装ガイドします。"
slug: building-http-router-with-net-http
date: 2021-10-24T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
  - HTTP
translation_key: building-http-router-with-net-http
---


# はじめに
本記事では、Golangの標準パッケージであるnet/httpを用いて、HTTPルーターを自作する方法について解説します。

標準パッケージはあまり多くのルーティングの機能を提供していません。

例えばHTTPメソッドごとのルーティングの定義ができなかったり、URLをパスパラメータとして利用できなかったり、正規表現を利用したルーティングの定義ができなかったりします。

その為、実際のアプリケーション開発ではより高機能なHTTPルーターを導入していることが少なくないのではないでしょうか。

そんなHTTPルーターですが、自作してみると次のようなメリットを享受できます。

- net/httpについて知ることができる
- アルゴリズムの面白さに触れることができる
- 自分が使いたい、使いやすいと思えるHTTPルーターを実装できる
- 実装を全て理解したHTTPルーターを自分のアプリケーションに組み込むことができる

本記事では次のような構成でHTTPルーターの自作について解説します。

- はじめに
- 目次
- 第1章HTTPルーターとは何か
- 第2章HTTPルーターのデータ構造
- 第3章HTTPサーバーのコードリーディング
- 第4章HTTPルーターの実装
- まとめ

各章について本筋から離れる内容についてはコラムとして記載しています。

本記事は以下のような読者にとって有意義な内容となることを想定しています。

- Golangの文法は理解したので何かを作ってみたい
- Golangの標準パッケージについて理解を深めるきっかけが欲しい
- 簡単なアルゴリズムの実装にGolangで取り組んでみたい
- 標準のルーティング機能が物足りないので拡張する方法が知りたい
- 普段使っているHTTPルーターの実装を理解したい

本記事を読むに当たっては、次のような前提知識があれば内容を十分に理解できます。

- Golangの基本的な文法の理解
- 何かしらのHTTPルーターを利用した経験

筆者は[bmf-san/goblin](https://github.com/bmf-san/goblin)というHTTPルーターのパッケージを公開しています。

ぜひコードを見たり、使ってみたりしてみてください。コントリビュートも大歓迎です。

# 目次
- はじめに
- 第1章HTTPルーターとは何か
- 第2章HTTPルーターのデータ構造
- 第3章HTTPサーバーのコードリーディング
- 第4章HTTPルーターの実装
- まとめ

# 第1章HTTPルーターとは何か
HTTPルーターはURLルーターと呼ばれたり、単にルーターと呼ばれたりもしますが、本記事ではHTTPルーターと呼称を統一することにします。

HTTPルーターは次の図のように、リクエストされたURLとレスポンスの処理を結びつけるアプリケーションです。

![route_in_client_and_server](/assets/images/posts/building-http-router-with-net-http/138551922-8fce2b9c-51d6-49e6-bea1-015e383cdb6e.png)

HTTPルーターはURLとレスポンスの処理がマッピングされたデータ（以下、ルートマップ）を元にすることで、ルーティングを行うことができます。

|       Request URL       |     Handler      |
| :---------------------- | :--------------- |
| GET  /                  | IndexHandler     |
| GET  /foo               | FooHandler       |
| POST /foo               | FooHandler       |
| GET  /foo/:id           | FooHandler       |
| POST /foo/:id           | FooHandler       |
| GET  /foo/:id/:name     | FooHandler       |
| POST /foo/:id/:name     | FooHandler       |
| GET  /foo/bar           | FooBarHandler    |
| GET  /foo/bar/:id       | FooBarHandler    |
| GET  /foo/bar/:id/:name | FooBarHandler    |
| GET  /foo/bar/baz       | FooBarBazHandler |
| GET  /bar               | BarHandler       |
| GET  /baz               | BazHandler       |

HTTPルーターの内部では、定義されたルートマップはルーティングに最適化されたデータ構造となります。

データ構造については、次の章で解説します。

本記事では、ルートマップを元に、リクエストのURLに応じたレスポンスの処理を探し出すことを「ルーティング」と定義します。

また、HTTPにおいてルーティングを行うアプリケーションのことを「HTTPルーター」と定義します。

---
### コラム：URLの仕様

URLは、インターネット上のページのアドレスを表し、Uniform ResourceLocatorの略語です。 

URL文字列の形式は次のように定義されています。 

```sh
<scheme>：<scheme-specific-part>
```

この部分にはhttp、https、ftpなどのプロトコル名がよく使用されますが、プロトコル名以外のスキーマ名も定義されています。 

[ユニフォームリソース識別子（URI）スキーム](https://www.iana.org/assignments/uri-schemes/uri-schemes.xhtml)

`<scheme-specific-part>`の部分では、スキーマに基づく文字列が定義されています。 

例えば、httpおよびhttpsスキームの場合、ドメイン名とパス名（またはディレクトリ名）が定義されるという規則があります。 

詳細なURL仕様については、RFC 1738を参照してください。 

[RFC 738-ユニフォームリソースロケーター（URL）](https://tools.ietf.org/html/rfc1738)

RFC 1738は、インターネット標準（STD1）として位置付けられています。

# 第2章HTTPルーターのデータ構造
## データ構造を考える
以下は第1章で例示したルートマップです。

|       Request URL       |     Handler      |
| :---------------------- | :--------------- |
| GET  /                  | IndexHandler     |
| GET  /foo               | FooHandler       |
| POST /foo               | FooHandler       |
| GET  /foo/:id           | FooHandler       |
| POST /foo/:id           | FooHandler       |
| GET  /foo/:id/:name     | FooHandler       |
| POST /foo/:id/:name     | FooHandler       |
| GET  /foo/bar           | FooBarHandler    |
| GET  /foo/bar/:id       | FooBarHandler    |
| GET  /foo/bar/:id/:name | FooBarHandler    |
| GET  /foo/bar/baz       | FooBarBazHandler |
| GET  /bar               | BarHandler       |
| GET  /baz               | BazHandler       |

URLに着目すると階層構造であることが見て取れます。

階層構造は木構造と相性が良いので、ルートマップを木構造で表現することを考えます。

## 木構造とは
グラフ理論における木の構造をしたデータ構造のことです。

木構造は階層構造を表現するのに適したデータ構造です。

木を構成する要素をノード（節）、一番上位に親のないノードをルート（根）、最下位にある子のないノードをリーフ（葉）と呼びます。ノードとノードの繋がりはエッジ（枝）と呼びます。

木にノードを追加することを挿入、木からノードを探し出すことを探索と言います。

![tree_structure](/assets/images/posts/building-http-router-with-net-http/138551925-b0544b07-1f2d-44bb-9724-495780908b2d.png)

木構造の中でも基本的な木である二分探索木の実装例を次に示します。

```golang
package main

import (
    "fmt"
)

// Node is a node of a tree.
type Node struct {
    Key   int
    Left  *Node
    Right *Node
}

// BST is a binary search tree.
type BST struct {
    Root *Node
}

// insert insert a node to tree.
func (b *BST) insert(key int) {
    if b.Root == nil {
        b.Root = &Node{
            Key:   key,
            Left:  nil,
            Right: nil,
        }
    } else {
        recursiveInsert(b.Root, &Node{
            Key:   key,
            Left:  nil,
            Right: nil,
        })
    }
}

// recursiveInsert insert a new node to targetNode recursively.
func recursiveInsert(targetNode *Node, newNode *Node) {
    // if a newNode is smaller than targetNode, insert a newNode to left child node.
    // if a newNode is a bigger than targetNode, insert a newNode to right childe node.
    if newNode.Key < targetNode.Key {
        if targetNode.Left == nil {
            targetNode.Left = newNode
        } else {
            recursiveInsert(targetNode.Left, newNode)
        }
    } else {
        if targetNode.Right == nil {
            targetNode.Right = newNode
        } else {
            recursiveInsert(targetNode.Right, newNode)
        }
    }
}

// remove remove a key from tree.
func (b *BST) remove(key int) {
    recursiveRemove(b.Root, key)
}

// recursiveRemove remove a key from tree recursively.
func recursiveRemove(targetNode *Node, key int) *Node {
    if targetNode == nil {
        return nil
    }

    if key < targetNode.Key {
        targetNode.Left = recursiveRemove(targetNode.Left, key)
        return targetNode
    }

    if key > targetNode.Key {
        targetNode.Right = recursiveRemove(targetNode.Right, key)
        return targetNode
    }

    if targetNode.Left == nil && targetNode.Right == nil {
        targetNode = nil
        return nil
    }

    if targetNode.Left == nil {
        targetNode = targetNode.Right
        return targetNode
    }

    if targetNode.Right == nil {
        targetNode = targetNode.Left
        return targetNode
    }

    leftNodeOfMostRightNode := targetNode.Right

    for {
        if leftNodeOfMostRightNode != nil && leftNodeOfMostRightNode.Left != nil {
            leftNodeOfMostRightNode = leftNodeOfMostRightNode.Left
        } else {
            break
        }
    }

    targetNode.Key = leftNodeOfMostRightNode.Key
    targetNode.Right = recursiveRemove(targetNode.Right, targetNode.Key)
    return targetNode
}

// search search a key from tree.
func (b *BST) search(key int) bool {
    result := recursiveSearch(b.Root, key)

    return result
}

// recursiveSearch search a key from tree recursively.
func recursiveSearch(targetNode *Node, key int) bool {
    if targetNode == nil {
        return false
    }

    if key < targetNode.Key {
        return recursiveSearch(targetNode.Left, key)
    }

    if key > targetNode.Key {
        return recursiveSearch(targetNode.Right, key)
    }

    // targetNode == key
    return true
}

// depth-first search
// inOrderTraverse traverse tree by in-order.
func (b *BST) inOrderTraverse() {
    recursiveInOrderTraverse(b.Root)
}

// recursiveInOrderTraverse traverse tree by in-order recursively.
func recursiveInOrderTraverse(n *Node) {
    if n != nil {
        recursiveInOrderTraverse(n.Left)
        fmt.Printf("%d\n", n.Key)
        recursiveInOrderTraverse(n.Right)
    }
}

// depth-first search
// preOrderTraverse traverse by pre-order.
func (b *BST) preOrderTraverse() {
    recursivePreOrderTraverse(b.Root)
}

// recursivePreOrderTraverse traverse by pre-order recursively.
func recursivePreOrderTraverse(n *Node) {
    if n != nil {
        fmt.Printf("%d\n", n.Key)
        recursivePreOrderTraverse(n.Left)
        recursivePreOrderTraverse(n.Right)
    }
}

// depth-first search
// postOrderTraverse traverse by post-order.
func (b *BST) postOrderTraverse() {
    recursivePostOrderTraverse(b.Root)
}

// recursivePostOrderTraverse traverse by post-order recursively.
func recursivePostOrderTraverse(n *Node) {
    if n != nil {
        recursivePostOrderTraverse(n.Left)
        recursivePostOrderTraverse(n.Right)
        fmt.Printf("%v\n", n.Key)
    }
}

// breadth-first search
// levelOrderTraverse traverse by level-order.
func (b *BST) levelOrderTraverse() {
    if b != nil {
        queue := []*Node{b.Root}

        for len(queue) > 0 {
            currentNode := queue[0]
            fmt.Printf("%d ", currentNode.Key)

            queue = queue[1:]

            if currentNode.Left != nil {
                queue = append(queue, currentNode.Left)
            }

            if currentNode.Right != nil {
                queue = append(queue, currentNode.Right)
            }
        }
    }
}

func main() {
    tree := &BST{}

    tree.insert(10)
    tree.insert(2)
    tree.insert(3)
    tree.insert(3)
    tree.insert(3)
    tree.insert(15)
    tree.insert(14)
    tree.insert(18)
    tree.insert(16)
    tree.insert(16)

    tree.remove(3)
    tree.remove(10)
    tree.remove(16)

    fmt.Println(tree.search(10))
    fmt.Println(tree.search(19))

    // Traverse
    tree.inOrderTraverse()
    tree.preOrderTraverse()
    tree.postOrderTraverse()
    tree.levelOrderTraverse()

    fmt.Printf("%#v\n", tree)
}
```

ここでは詳細に説明することは割愛しますが、二分探索木は木構造の基本的なアルゴリズムを学ぶにちょうど良い木です。

木構造には二分探索木の他にも様々な種類があります。その中でもトライ木（プレフィックス木ともいわれる。本記事ではトライ木と呼称します）と呼ばれる木構造は文字列の探索がしやすいという特徴があります。

トライ木を利用することによりルーティングで扱いやすいデータ構造を表現できます。

## トライ木とは
トライ木は、IPアドレス探索や形態素解析などでも利用されている木構造の一種です。

各ノードは単一または複数の文字列あるいは数値を持ち、根ノードから葉に向かって探索して値をつなげていくことで単語を表現します。

アルゴリズムを可視化して動的に理解できるサービスがあるので、そちらを見てみるとトライ木のデータ構造を理解しやすいです。

cf. [Algorithm Visualizations - Trie (Prefix Tree）](https://www.cs.usfca.edu/~galles/visualization/Trie.html)

トライ木は比較的簡単に実装できます。

次のコードは探索と挿入だけ実装されたトライ木のコード例です。

```golang
package main

import "fmt"

// Node is a node of tree.
type Node struct {
    key      string
    children map[rune]*Node
}

// NewTrie is create a root node.
func NewTrie() *Node {
    return &Node{
        key:      "",
        children: make(map[rune]*Node),
    }
}

// Insert is insert a word to tree.
func (n *Node) Insert(word string) {
    runes := []rune(word)
    curNode := n

    for _, r := range runes {
        if nextNode, ok := curNode.children[r]; ok {
            curNode = nextNode
        } else {
            curNode.children[r] = &Node{
                key:      string(r),
                children: make(map[rune]*Node),
            }
        }
    }
}

// Search is search a word from a tree.
func (n *Node) Search(word string) bool {
    if len(n.key) == 0 && len(n.children) == 0 {
        return false
    }

    runes := []rune(word)
    curNode := n

    for _, r := range runes {
        if nextNode, ok := curNode.children[r]; ok {
            curNode = nextNode
        } else {
            return false
        }
    }

    return true
}

func main() {
    t := NewTrie()

    t.Insert("word")
    t.Insert("wheel")
    t.Insert("world")
    t.Insert("hospital")
    t.Insert("mode")

    fmt.Printf("%v", t.Search("mo")) // true
}
```

このトライ木をベースにすることで、ルーティングに最適化したデータ構造を検討します。

## トライ木をベースにルートマップのデータ構造を考える
トライ木の考え方をベースにルートマップのデータ構造を考えます。

以下は筆者が開発している[bmf-san/goblin](https://github.com/bmf-san/goblin)で採用しているデータ構造となります。

goblinでは、ミドルウェアやパスパラメータをサポートしているため、それらに対応したデータ構造となっています。

![trie_based_tree_for_goblin](/assets/images/posts/building-http-router-with-net-http/138551926-666c7e6e-03f7-4a5a-8f18-ad1ba27b615a.png)

このデータ構造は次のようなルートマップを表現しています。

|   Request URL    |     Handler      | Middleware |
| :--------------- | :--------------- | :--------- |
| GET  /           | IndexHandler     | none       |
| GET  /foo        | FooHandler       | FooMws     |
| POST /foo        | FooHandler       | FooMws     |
| GET  /foo/bar    | FooBarHandler    | none       |
| GET /foo/bar/:id | FooBarHandler    | none       |
| GET /foo/baz     | FooBazHandler    | none       |
| GET /foo/bar/baz | FooBarBazHandler | none       |
| GET /baz         | BazHandler       | none       |

観点としては、以下の二点に集約されます。

- URLをどのようなルールで木構造として表現するか
- ノードに持たせる必要なデータは何か

前者はルーティングの性能を決める部分であり、処理時間やメモリ効率などを追求する場合はより高度な木構造の採用を検討する必要があります。

後者はHTTPルーターの機能に関わる部分なので、提供したい機能によって様々です。

今回紹介したトライ木をベースとした木構造は、あくまで筆者が考えた木構造に過ぎません。

HTTPルーターの実装要件によりデータ構造はそれぞれです。

次の章でこのデータ構造をHTTPルーターに組み込むため上で知っておきたいことについて説明します。

---
### コラム：基数木（パトリシア木）

文字列を格納するトライ木を更に発展させた木構造に基数木という木構造があります。

基数木はパフォーマンスに配慮したHTTPルーターでは良く使われているのを筆者は観測しています。

Golangのstringsパッケージの内部でも使われているようです。

[https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/strings/strings.go;l=924](https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/strings/strings.go;l=924)

# 第3章HTTPサーバーのコードリーディング
HTTPルーターの実装について説明する前に、次のようなnet/httpを利用したHTTPサーバーのコードを例に、HTTPルーターを実装する上で知っておきたいことについて説明します。

必要に応じて以下のリンクを参照してください。

[https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:](https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:)

```golang
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
```

このコードは単純であるものの、HTTPルーターを自作する上で示唆に富んだコードです。

このコードはマルチプレクサの呼び出し、ハンドラの登録、サーバーの起動という流れの処理になっています。

それぞれについて順番に見ていきます。

## マルチプレクサの呼び出し
まず最初のコードは、`http.ServeMux`という構造体を生成しています。

```golang
mux := http.NewServeMux()
```

net/httpのドキュメントでは、`http.ServeMux`はHTTPリクエストマルチプレクサ（以下、マルチプレクサ）であると説明がなされています。

[type ServeMux](https://pkg.go.dev/net/http#ServeMux)

このマルチプレクサは、リクエストのURLを登録済みのパターンと照らし合わせて、最もマッチするハンドラ（レスポンスを返却する関数）を呼び出すという役目を持っています。

`http.ServeMux`はつまり、ルーティングのための構造体であるということが言えます。

この`http.ServeMux`には`ServeHTTP`というメソッドが実装されています。

```golang
// ServeHTTP dispatches the request to the handler whose
// pattern most closely matches the request URL.
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(StatusBadRequest)
		return
	}
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}
```

[cs.opensource.google - go1.17.2:src/net/http/server.go;l=2415](https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/net/http/server.go;l=2415)

`ServeHTTP`の以下の部分を更に読み進めていくと、`ServeHTTP`のルーティングの処理が見えてきます。
```golang
h, _ := mux.Handler(r)
```

順々にコードジャンプしていくと、マッチするハンドラを探して返却する関数にたどり着きます。

```golang
// Find a handler on a handler map given a path string.
// Most-specific (longest) pattern wins.
func (mux *ServeMux) match(path string) (h Handler, pattern string) {
	// Check for exact match first.
	v, ok := mux.m[path]
	if ok {
		return v.h, v.pattern
	}

	// Check for longest valid match.  mux.es contains all patterns
	// that end in / sorted from longest to shortest.
	for _, e := range mux.es {
		if strings.HasPrefix(path, e.pattern) {
			return e.h, e.pattern
		}
	}
	return nil, ""
}
```

[https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/net/http/server.go;l=2287;drc=refs%2Ftags%2Fgo1.17.2](https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/net/http/server.go;l=2287;drc=refs%2Ftags%2Fgo1.17.2)

マッチするハンドラが見つかった場合は、そのハンドラの`ServeHTTP`を呼びだすことで、レスポンスのための処理を呼び出します。

それが`http.ServeMux`に実装された`ServeHTTP`メソッドの末尾にある処理です。

```golang
h.ServeHTTP(w, r)
```

HTTPルーターを自作するためには、標準のマルチプレクサに取って代われるように、`http.Handler`型を満たした（≒`ServeHTTP`を実装した）マルチプレクサを実装してあげる必要があります。

[type Handler](https://pkg.go.dev/net/http#Handler)

## ハンドラの登録
続いて次のコードでは、マルチプレクサにハンドラを登録しています。

```golang
mux.HandleFunc("/", handler)
```

マルチプレクサに登録されるハンドラは、`http.Handler`型を満たす（≒`ServeHTTP`が実装される）必要があります。

```golang
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handler := foo{}
	mux.Handle("/", handler)

	http.ListenAndServe(":8080", mux)
}

type foo struct{}

// Satisfy the http.Handler type by implementing ServeHTTP.
func (f foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
```

あるいは、`http.HandlerFunc`型を実装する形でもハンドラを作成できます。

[func (HandlerFunc) ServeHTTP](https://pkg.go.dev/net/http#HandlerFunc.ServeHTTP)

`http.HandlerFunc`型は`func(ResponseWriter, *Request)`を型として定義したもので、`ServeHTTP`メソッドを実装しています。

```golang
// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

[https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/net/http/server.go;l=2045](https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/net/http/server.go;l=2045)

従って`http.HandlerFunc`型を使う場合は次のようにハンドラを作成できます。

```golang
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handler))

	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
```

HTTPルーターを実装する上では、`http.Handler`型をサポートするように実装を意識すると、ハンドラの作成方法に柔軟性をもたせることができるため、扱いやすいパッケージになります。

## サーバーの起動
最後のコードでは、サーバーを起動するポート番号とマルチプレクサを関数に渡して、HTTPサーバーを起動しています。

```golang
http.ListenAndServe(":8080", mux)
```

[func ListenAndServe](https://pkg.go.dev/net/http#Server.ListenAndServe)

内部的には、`http.Server`型の`ListenAndServe`が呼び出されています。

[func (*Server) ListenAndServe](https://pkg.go.dev/net/http#Server.ListenAndServe)

この関数では、第2引数がnilのときは`http.DefaultServeMux`というデフォルトの`http.ServeMux`が利用されるようになっています。

つまり、マルチプレクサを拡張したいようなケース以外では、マルチプレクサをわざわざ生成しなくても良いということです。

HTTPルーターを実装していく上では、話の前触れとして必要だったため、マルチプレクサをわざわざ生成するようなコードを例として上げました。

HTTPルーターを実装する上での必要なコードリーディングができたので、次の章から実装の解説をします。

---
### コラム:末尾スラッシュについて

URLの末尾に付与される`/`はドメイン名末尾と、サブディレクトリ末尾のケースでそれぞれ違いがあります。

ドメイン名末尾の場合、一般的なブラウザでは`/`が無い場合は、`/`が有るURLにリクエストします。

- `https://bmf-tech.com` → `https://bmf-tech.com/` にリクエスト
- `https://bmf-tech.com/` → `https://bmf-tech.com` にリクエスト

ドメイン名末尾の場合は`/`の有無にあまり違いはありませんが、サブディレクトリ末尾場合は明確な違いがあります。

- `https://bmf-tech.com/posts` → ファイルへのリクエスト
- `https://bmf-tech.com/posts/` → ディレクトリへのリクエスト

より詳しく仕様について知りたい場合はRFCを参照してください。

[RFC 2616](https://www.w3.org/Protocols/rfc2616/rfc2616-sec5.html#sec5.1)
[RFC 3986](https://www.ietf.org/rfc/rfc3986.txt)

HTTPルーターを実装する上では、URLのパス部分をどう解釈するかという点で気にしておく必要がある部分です。

筆者が開発した[bmf-san/goblin](https://github.com/bmf-san/goblin)では、末尾`/`有無は同じルーティングの定義として扱う仕様としています。

# 第4章HTTPルーターの実装
HTTPルーターの実装をするための準備ができたので、実装の解説をします。

今回は標準パッケージよりも少しだけ高機能なルーターを実装します。

具体的には次の2つの特徴を備えたルーターになります。

- メソッドベースのルーティングをサポートしている
- トライ木をベースとしたアルゴリズムを実装している

標準パッケージの機能では、HTTPメソッド別にルーティングを登録できません。

HTTPメソッド別にルーティングをしたい場合はハンドラーの中でHTTPメソッドごとの条件分岐をするような実装が必要となります。

```golang
// ex.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			// do something...
		case http.MethodPost:
			// do something...

		...

		default:
```

ハンドラーにこのような条件分岐を定義せずとも、メソッドごとにルーティングを定義できる機能を実装します。

メソッドベースでルーティングを定義可能なHTTPルーターのアルゴリズムとしては、HTTPルーターのデータ構造で説明したトライ木をベースとした木構造を採用します。

## 準備
今回実装するHTTPルーターのソースコードは以下にあります。

[bmf-san/introduction-to-golang-http-router-made-with-net-http](https://github.com/bmf-san/introduction-to-golang-http-router-made-with-net-http)

テストコードは実装過程で書くことを推奨しますが、テストコードについて解説は行いません。

CIについても同様です。

Golangのバージョンは1.17を利用しています。

## 実装
実装手順としてはまず、トライ木をベースとしたルーティングのアルゴリズムを実装するところから始めます。

その後でメソッドベースのルーティングをサポートするための実装をします。

### トライ木をベースとしたルーティングのアルゴリズムの実装
それでは早速実装していきます。

ここで実装するコードは全て以下で参照できます。

[bmf-san/introduction-to-golang-http-router-made-with-net-http/blob/main/trie.go](https://github.com/bmf-san/introduction-to-golang-http-router-made-with-net-http/blob/main/trie.go)

今回は、goblinのデータ構造を簡素化した以下のような木構造を採用することにします。

![tree_for_implementation](/assets/images/posts/building-http-router-with-net-http/138551924-ca2e28f5-a0c1-4ae4-ac9f-76d6ee1c46d9.png)


この木構造で表現されるルートマップは以下の通りです。

|       Request URL       |     Handler      |
| :---------------------- | :--------------- |
| GET  /                  | IndexHandler     |
| GET  /foo               | FooHandler       |
| POST /foo               | FooHandler       |
| GET  /foo/bar           | FooBarHandler    |
| GET  /foo/bar/baz       | FooBarBazHandler |
| GET  /bar               | BarHandler       |
| GET  /baz               | BazHandler       |

上記の木構造を表現するために、まずは必要なデータを定義していくところから書き始めます。

`trie.go`というファイルを作成し、構造体を定義してください。

```golang
package myrouter

// tree is a trie tree.
type tree struct {
	node *node
}

type node struct {
	label    string
	actions  map[string]*action // key is method
	children map[string]*node   // key is a label o f next nodes
}

// action is an action.
type action struct {
	handler http.Handler
}

// result is a search result.
type result struct {
	actions *action
}
```

`tree`は木そのもの、`node`は木を構成する要素で、`label`、`actions`、`children`を持ちます。

`label`はURLのパス、`actions`はHTTPメソッドとハンドラーのマップを表現します。`children`は`label`と`node`のマップで、子ノードを表現します。

`result`は木からの探索結果を表現します。

続いてこれらの構造体を生成する関数を定義しておきます。

```golang
// newResult creates a new result.
func newResult() *result {
	return &result{}
}

// NewTree creates a new trie tree.
func NewTree() *tree {
	return &tree{
		node: &node{
			label:    pathRoot,
			actions:  make(map[string]*action),
			children: make(map[string]*node),
		},
	}
}
```

では、木へノードを追加する部分の処理を実装します。

`tree`をポインタレシーバとした`Insert`メソッドを定義します。

```golang
func (t *tree) Insert(methods []string, path string, handler http.Handler)) {
	// 
}
```

この関数の引数のポイントとしては、HTTPメソッドを複数渡せるように引数を定義している点です。

HTTPメソッドごとに単一のハンドラーを定義できるだけでなく、複数のメソッドに対して同一のハンドラーを定義できるようになっています。

```golang
// ex.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			// do something...
		case http.MethodPost:
			// do something...

		...

		default:
```

実装の方針によっては、ハンドラーの中でHTTPメソッドの条件分岐をしたいというケースもあるという可能性を考慮して、汎用性を持たせています。

続いて、`Insert`の中身ですが、最初にスタート地点となるノードを変数として定義しています。

```golang
func (t *tree) Insert(methods []string, path string, handler http.Handler)) {
	curNode := t.node
}
```

次に探索する対象が`/`（ルート）の場合の条件分岐を処理します。

```golang
const (
	pathRoot      string = "/"
)

func (t *tree) Insert(methods []string, path string, handler http.Handler)) {
	curNode := t.node
	if path == pathRoot {
		curNode.label = path
		for _, method := range methods {
			curNode.actions[method] = &action{
				handler: handler,
			}
		}
		return nil
	}
}
```

`/`の場合は、後続のループ処理をする必要がないので、ここで木へのノード追加して終了するように処理します。

`/`以外の場合は処理を継続します。

URLのパスを`/`で分解して、[]string型のスライスにパスの文字列を格納する処理を行います。

```golang
const (
	...
	pathDelimiter string = "/"
)

func (t *tree) Insert(methods []string, path string, handler http.Handler)) {
	...

	ep := explodePath(path)	
}

// explodePath removes an empty value in slice.
func explodePath(path string) []string {
	s := strings.Split(path, pathDelimiter)
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
```

[]string型のスライスは、ノード追加する位置を見つけるために、rangeで走査します。

ここでの処理はHTTPルーターのデータ構造で説明したトライ木の実装がベースとしています。

子ノードが見つからなかったときにノードを追加するようにします。

ルーティングの定義が重複するようなケースとなった場合は、後勝ちとなる仕様になるように処理しています。

```golang
// Insert inserts a route definition to tree.
func (t *tree) Insert(methods []string, path string, handler http.Handler) error {
	...

	for i, p := range ep {
		nextNode, ok := curNode.children[p]
		if ok {
			curNode = nextNode
		}
		// Create a new node.
		if !ok {
			curNode.children[p] = &node{
				label:    p,
				actions:  make(map[string]*action),
				children: make(map[string]*node),
			}
			curNode = curNode.children[p]
		}
		// last loop.
		// If there is already registered data, overwrite it.
		if i == len(ep)-1 {
			curNode.label = p
			for _, method := range methods {
				curNode.actions[method] = &action{
					handler: handler,
				}
			}
			break
		}
	}

	return nil
}
```

最終的な`Insert`の実装は次のようになります。

```golang
// Insert inserts a route definition to tree.
func (t *tree) Insert(methods []string, path string, handler http.Handler) error {
	curNode := t.node
	if path == pathRoot {
		curNode.label = path
		for _, method := range methods {
			curNode.actions[method] = &action{
				handler: handler,
			}
		}
		return nil
	}
	ep := explodePath(path)
	for i, p := range ep {
		nextNode, ok := curNode.children[p]
		if ok {
			curNode = nextNode
		}
		// Create a new node.
		if !ok {
			curNode.children[p] = &node{
				label:    p,
				actions:  make(map[string]*action),
				children: make(map[string]*node),
			}
			curNode = curNode.children[p]
		}
		// last loop.
		// If there is already registered data, overwrite it.
		if i == len(ep)-1 {
			curNode.label = p
			for _, method := range methods {
				curNode.actions[method] = &action{
					handler: handler,
				}
			}
			break
		}
	}

	return nil
}

// explodePath removes an empty value in slice.
func explodePath(path string) []string {
	s := strings.Split(path, pathDelimiter)
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
```

これで木への挿入の処理が実装できたので、次は木からの探索の処理を実装します。

挿入と比べて探索は比較的シンプルなので、一度で説明します。

```golang
func (t *tree) Search(method string, path string) (*result, error) {
	result := newResult()
	curNode := t.node
	if path != pathRoot {
		for _, p := range explodePath(path) {
			nextNode, ok := curNode.children[p]
			if !ok {
				if p == curNode.label {
					break
				} else {
					return nil, ErrNotFound
				}
			}
			curNode = nextNode
			continue
		}
	}
	result.actions = curNode.actions[method]
	if result.actions == nil {
		// no matching handler was found.
		return nil, ErrMethodNotAllowed
	}
	return result, nil
}
```

探索の場合も挿入と同じく、URLのパスが`/`か否かでループ処理に進むかどうか決まります。

ループ処理に進む場合は、子ノードを見ていき対象のノードが存在するか探索していくだけです。

対象のノードが存在する場合は、リクエストのHTTPメソッドとマッチするハンドラを探して、`result`を返します。

### メソッドベースのルーティングをサポートするための実装
ここで実装するコードの全体像は以下になります。

[bmf-san/introduction-to-golang-http-router-made-with-net-http/blob/main/router.go](https://github.com/bmf-san/introduction-to-golang-http-router-made-with-net-http/blob/main/router.go)

ここではHTTPルーターとしての機能を提供するための実装も合わせて行います。

まずは構造体の定義と生成用の関数です。

```golang
// Router represents the router which handles routing.
type Router struct {
	tree *tree
}

// route represents the route which has data for a routing.
type route struct {
	methods []string
	path    string
	handler http.Handler
}

func NewRouter() *Router {
	return &Router{
		tree: NewTree(),
	}
}
```

`Router`はnet/httpでいう`http.ServeMux`に当たります。

`route`はルーティングの定義のためのデータを持ちます。

次に、`Router`に次の3つのメソッドを実装します。

```golang
...

func (r *Router) Methods(methods ...string) *Router {
	tmpRoute.methods = append(tmpRoute.methods, methods...)
	return r
}

// Handler sets a handler.
func (r *Router) Handler(path string, handler http.Handler) {
	tmpRoute.handler = handler
	tmpRoute.path = path
	r.Handle()
}

// Handle handles a route.
func (r *Router) Handle() {
	r.tree.Insert(tmpRoute.methods, tmpRoute.path, tmpRoute.handler)
	tmpRoute = &route{}
}
```

`Methods`はHTTPメソッドのセッター、`Handler`はURLのパスとハンドラーのセッターで`Handle`を呼び出します。`Handle`では先程実装した木への挿入の処理を呼び出します。

`Methods`や`Handler`はHTTPルーターを利用する側の可読性を意識して、メソッドチェインとして実装しています。

メソッドベースのルーティングは木と組み合わせてこれで実現できます。

最後は、`Router`に`ServeHTTP`を実装させたら完成です。

```golang
...

// ServeHTTP dispatches the request to the handler whose
// pattern most closely matches the request URL.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path
	result, err := r.tree.Search(method, path)
	if err != nil {
		status := handleErr(err)
		w.WriteHeader(status)
		return
	}
	h := result.actions.handler
	h.ServeHTTP(w, req)
}

func handleErr(err error) int {
	var status int
	switch err {
	case ErrMethodNotAllowed:
		status = http.StatusMethodNotAllowed
	case ErrNotFound:
		status = http.StatusNotFound
	}
	return status
}
```

## 実装したHTTPルーター使ってみる
今回実装したHTTPルーターは次のように使うことができます。

サーバーを起動してそれぞれのエンドポイントにリクエストして動作確認してみてください。

```golang
package main

import (
	"fmt"
	"net/http"

	myroute "github.com/bmf-san/introduction-to-golang-http-router-made-with-net-http"
)

func indexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /")
	})
}

func fooHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "GET /foo")
		case http.MethodPost:
			fmt.Fprintf(w, "POST /foo")
		default:
			fmt.Fprintf(w, "Not Found")
		}
	})
}

func fooBarHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /foo/bar")
	})
}

func fooBarBazHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /foo/bar/baz")
	})
}

func barHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /bar")
	})
}

func bazHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /baz")
	})
}

func main() {
	r := myroute.NewRouter()

	r.Methods(http.MethodGet).Handler(`/`, indexHandler())
	r.Methods(http.MethodGet, http.MethodPost).Handler(`/foo`, fooHandler())
	r.Methods(http.MethodGet).Handler(`/foo/bar`, fooBarHandler())
	r.Methods(http.MethodGet).Handler(`/foo/bar/baz`, fooBarBazHandler())
	r.Methods(http.MethodGet).Handler(`/bar`, barHandler())
	r.Methods(http.MethodGet).Handler(`/baz`, bazHandler())

	http.ListenAndServe(":8080", r)
}
```

駆け足気味になってしまいましたが、実装の解説は以上です。

---
### コラム：HTTPルーターのパフォーマンス比較

HTTPルーターのパフォーマンス比較に興味があるのであれば、以下のリポジトリを見てみてください。

[julienschmidt/go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark)

筆者はこのリポジトリにgoblinのパフォーマンス比較のPRを出しました。

[Add a new router goblin #97](https://github.com/julienschmidt/go-http-routing-benchmark/pull/97)

# まとめ
本記事ではHTTPルーターを自作するまでのアプローチについて解説しました。

第1章では、HTTPルーターは何をするアプリケーションなのかについて整理しました。

第2章では、HTTPルーターにおけるデータ構造について、例を混じえながら解説しました。

第3章では、net/http使ったHTTPサーバーのコードについて深堀りしました。

そして、第4章ではHTTPルーターの実装方法についてコードとともに説明しました。

本記事を通じて、何か1つでも読者の役に立つことがあったり、興味を持ってもらえることがあれば幸いです。

また。拙作である[bmf-san/goblin](https://github.com/bmf-san/goblin)についてもコードを見てもらえるきっかけになれば嬉しいです。

質問や修正依頼、フィードバック等あればぜひ教えて下さい。

# あとがき
- [zenn.dev - 
net／httpでつくるHTTPルーター 自作入門](https://zenn.dev/bmf_san/books/3f41c5cd34ec3f)
  - この記事の内容を本にしてあります。 
- [dev.to - Introduction to Golang HTTP router made with net/http](http://web.archive.org/web/20250815231536/https://dev.to/bmf_san/introduction-to-golang-http-router-made-with-nethttp-3nmb)
  - 英語に翻訳したものです。
- ~~github.com - bmf-san/book-introduction-to-golang-http-router-made-with-net-http~~
  - 原文管理しているリポジトリです。
- [net/httpでつくるHTTPルーター自作入門](https://speakerdeck.com/bmf_san/httpdetukuruhttprutazi-zuo-ru-men)
  - Go Conference 2021に登壇してきました
