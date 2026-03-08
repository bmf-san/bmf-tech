---
title: Go CodeReviewCommentsのまとめ
slug: go-code-review-comments-summary
date: 2020-09-15T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
  - コードレビュー　
translation_key: go-code-review-comments-summary
---


# 概要
[github.com - CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)を読んでメモしておきたいことをまとめる。

# Comment Sentences
- コメントの終わりはピリオドで終わるようにする。
- [golang.org - commentary](https://golang.org/doc/effective_go.html#commentary)

# Copying
- 別のパッケージから構造体をコピーするときは、予期しない参照に気を付ける。
- メソッドがポインタの値に関連付けられているならTではなく*Tを使うようにする。

# Crypt Rand
- 鍵の生成に`math/rand`を使ってはいけない。`crypto/rand`の`Reader`を使う。
  - 文字列として扱いたい場合は16進数かbase64エンコードする。

# Declaring Empty Slices
```go
// 長さ0のスライス
t := []string{}
```

よりも、

```go
// nilのスライス
var t []string
```

のほうを使うようにする。

JSONオブジェクトのエンコード時、`nil`は`null`に変換されるが、`[]string{}`は`[]`に変換される。

インターフェース設計では両者を区別しないほうがよい。分かりづらいミスを誘発する可能性があるため。

# Don't Panic
- 通常のエラーハンドリングで`panic`を使うのは避け、`error`型を含んだ複数の値を返すようにする。

# Goroutine Lifetimes
- goroutineの生成時は、いつ終了されるかを明確にする。
- goroutineはchannelの送受信をブロックによってメモリリークを起こす場合がある。
- ガベージコレクターはブロックされているchannelに到達できなくても、goroutineを停止させない。

# Import Blank
- `import _ "pkg"`ではパッケージをインポートした際の副作用を利用することができる。
- この方法はプログラムのメインパッケージもしくはテストのみで利用する。

# Import Dot
- .を使ったimportはimport先のパッケージをimportで指定したパッケージの一部のように扱うため、循環参照を回避することができる。

```go
package foo_test

import (
    "bar/testutil" // fooでもimportされている
    . "foo" // foo_testをfooの一部のように見せる
)
```

# Named Result Parameters
- 名前付き戻り値は、戻り値が何を意味するかわかりづらいときに使う。

# Naked Returns
- Named Result Parametersと同じ。

# Receiver Type
メソッドのレシーバをポインタにするか、値にするかの基準。
迷う場合はポインタにしておく。

### ポインタを避けるケース
- レシーバが、`map`、`func`、`channel`ならポインタは避ける。
- レシーバが`slice`で、メソッドが`slice`を作り直さない場合は、ポインタを避ける。
- レシーバが小さく、本来値型であったり（ex. `time.Time`）、変更するフィールドやポインタがない構造体や配列、あるいは`int`や`string`のような型の場合は、レシーバが値であるほうが良い場合もある。
  - もし値がメソッドに渡されるとヒープ領域にメモリを確保する代わりにスタックメモリのコピーが行われてしまう。

### ポインタにするケース
- メソッドが値を変更する必要がある場合、レシーバーはポインタにする。
- レシーバが`sync.Mutex`か、同期するようなフィールドを持つ構造体ならレシーバはポインタにする。
- レシーバが大きな構造体や配列の場合は、ポインタにする。
- 関数が同時実行されたり、メソッドが呼び出されたときにレシーバの値を変更するか？
  - 値渡しではメソッド実行時にレシーバのコピーを生成する。従ってメソッドの外ではレシーバの変更が適用されない。変更がオリジナルのレシーバに適用される必要があるならレシーバはポインタにする。
- レシーバが値を変更されるかもしれない構造体、配列、スライス、その他の要素である場合、レシーバをポインタにしたほうがコードリーディングしやすい。

# Useful Test Failures
テストが失敗したときに伝えるべきメッセージ。
- 何が悪かったか（≒errorの原因）
- どんな入力があったか（≒test cases）
- 実際にどんな値があったか（≒actual）
- どんな値が来ることを期待していたのか（≒expected）

# 参考
- [github.com - CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)
