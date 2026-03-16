---
title: 実用Go言語
description: 実用Go言語
slug: practical-go-language
date: 2023-08-05T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
  - 書評
translation_key: practical-go-language
books:
  - asin: "4873119693"
    title: "実用Go言語"
---


[実用Go言語](https://amzn.to/3KO6sr4)を読み終わったので読書メモを残す。

何年かGoを触っているが知らなかったり忘れていたことに気づけて大変勉強になった。

# 変数名
- Errorインターフェースを満たすError型の変数名の接尾辞はError
  - ex. NotImplementedError
- errors.Newで宣言されるエラーの変数名の接頭辞はErr
  - ex. `ErrNotImplemented := errors.New("Not Implemented!")`
- 変数名は短縮形が好まれるが、**宣言箇所から離れた場所で使われるような変数については説明的な変数名**が望ましい
- 変数は短縮形式を積極的に使う
  - 型を明示的にしたい場合はvarを使っても良い

# パッケージ名
- internalという名前のパッケージはモジュール外に公開されない（≒internalの配置されているパッケージとその配下のパッケージ以外から利用できない）
- パッケージ名と関数名の重複を避ける
  - ○http.Server　✕http.HTTPServer
  - ○zip.NewWriter ✕zip.NewZipWriter
- `.`と`_`と`testdata`という名前のフォルダはコンパイル対象から外れる

# インターフェース名
- 標準パッケージで、単一のメソッドのみを持つインターフェースは接尾辞がerとなっていることがある
  - ex. io.Reader, fmt.Stringer

# 定数
- constはコンパイル時に決定されるイミュータブルな値にしか利用できない
- iotaは意図しない定数値のずれに気をつける
  - 末尾以外に定数を追加する可能性があるときはiotaではなく、それぞれの個別に定数を定義する
  - 1つのプロセスだけではなく、他のプロセスでも利用されるような値にiotaを使うのは避ける
    - ex. HTTPサーバーのレスポンスでクライアントに返す値など注意
-  ログなどでiotaの整数値を文字列に変換するには、golang.org/x/tools/cmd/stringerが利用できる
  - 整数値だと意味を追わないといけないので面倒な時に便利

# データのマスキング
- StringerとGoStringerのインターフェースを使って拡張することで実現できる

# 構造体
- 構造体埋め込み≠継承
  - 継承は子が親に依存するという関係性だが、構造体埋め込みは親が子に緩い依存をする関係性（語弊あるかも・・）。Is-Aではなく、Has-A。
  - 継承ではなく、委譲。
- 構造体のフィールドで、ゼロ値と区別してomitemptyしたい場合はフィールドにポインター型を使う

# 関数
- log.Fatalとpanicを使っても良い場所
  - main関数、init関数、Must接頭辞の関数、その他のアプリケーション初期化処理

# テスト
- 順序依存の検知
  - `go test -shuffle=on`でテストの実行順がシャッフルされる

