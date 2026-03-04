---
title: "Golangの開発環境を構築"
slug: "golang-development-environment"
date: 2018-04-07
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Golang"
draft: false
---

# 概要
Goの環境を構築します。

# Goの環境構築
## Goをインストールする
インストール手段は省略します。私は**anyenv**というツールでインストールしています。

## GOPATHを指定する
`.bashrc`または`.bash_profile`にGOPATHを指定します。

```bash
export GOPATH=$HOME/localdev/project/go_dev // 好きなように設定してください
PATH=$PATH:$GOPATH/bin
```

## Goのディレクトリ構成を確認する
ローカル環境でのディレクトリ構成は公式ドキュメントに従う形で構成していきます。

```
go_dev/
├── bin
├── pkg
└── src
```

**go_dev**というGoの開発用ディレクトリを用意し、その中に公式ドキュメントのディレクトリ構成に準拠した役割ごとの３つのディレクトリを用意する形になっています。
**bin**には実行可能なコマンドが、**pkg**にはパッケージが、**src**にはソースがそれぞれ配置されます。
git管理するのは**src/**です。

# パッケージをつくってみる
試しにここまでの設定が上手できているかの確認でパッケージを作ってみます。

`src/`に`test/`を用意し、`main.go`というファイルを以下の通り作成します。

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```

`go build main.go`でコンパイルし、バイナリファイルを作成、`go install`で`bin/`に`test`というバイナリファイルが生成されていればOKです。


# Dockerの開発環境構築
- [github - astaxie/build-web/application-with-golang](https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/01.0.md)
- [golang.org](https://golang.org/doc/code.html)
- [Go言語の開発環境を設定する](https://medium.com/@Akitsuyoshi/go%E8%A8%80%E8%AA%9E%E3%81%AE%E9%96%8B%E7%99%BA%E7%92%B0%E5%A2%83%E3%82%92%E8%A8%AD%E5%AE%9A%E3%81%99%E3%82%8B-77f272229a64)
- [他言語から来た人がGoを使い始めてすぐハマったこととその答え](https://qiita.com/mumoshu/items/0d2f2a13c6e9fc8da2a4)
