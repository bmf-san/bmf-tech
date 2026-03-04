---
title: "vscodeでgoのLanguage Serverを有効にしたらコード定義元ジャンプができなくなった"
slug: "vscode-go-language-server-issue"
date: 2020-07-19
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Golang"
  - "gocode"
  - "gopls"
  - "Language Server"
  - "vscode"
  - "Tips"
draft: false
---

# 概要
vscodeでgoのLanguage Serverの設定を有効にしたらコード定義元へのジャンプができなくなってしまったので原因を調査した。

settings.json
```json
"go.useLanguageServer": true,
```

# 結論
`go.mod`がプロジェクトのルートに存在している必要がある。

cf. [stackoverflow - How to properly use go modules in vscode?
](https://stackoverflow.com/questions/59732657/how-to-properly-use-go-modules-in-vscode)


vscodeでフォルダを開くときに、こうではなく、

```
.
├── app
    ├── go.mod
```

こう開くようにしないとパスが良しなに解決されないせいか、コードジャンプできなかった。
```
.
├── go.mod
```

参考までに、go.modの内容。
```go
module github.com/bmf-san/gobel-api/app

go 1.14

require (
	github.com/bmf-san/goblin v0.0.0-20200718124906-8b3133b538d6
	github.com/bmf-san/golem v0.0.0-20200718182453-066c8e70e46e
	github.com/go-redis/redis/v7 v7.4.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899
)
```

`module github.com/bmf-san/gobel-api/app`と書いているのでこれを読み取って良しなに解釈してくれるのかと思ったらそうではないらしい。

# 調査方法
vscodeのターミナルを開いて、OUTPUT>gopls(server)を選択。コードジャンプしてみると、エラーログを確認することができる。

エラーログから推測するに、パスが怪しかったので色々調べてみたら上述のstackoverflowに当たりがあった。

# 解決策
パッと思いつく限りの対策は以下。

- language serverの設定をオフにする
- go.modをプロジェクトルートに置くようにする、あるいはgo.modが存在しているフォルダを開く

goplsやvscodeの設定等で調整できるのかもしれないのが、すぐ見つけられなくて時間かかりそうだったのでlanguage serverの設定をオフにする方向でとりあえず対応した...

まだ枯れた設定ではないと思うので、そのうち類似の事例やベストな解決策が見つかると思う多分...

何かわかったら追記予定。

# 関連
- [Big Sky - gocode やめます(そして Language Server へ)](https://mattn.kaoriya.net/software/lang/go/20181217000056.htm)
