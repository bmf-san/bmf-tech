---
title: "Goでクリーンアーキテクチャを実装する実践ガイド"
slug: golang-clean-architecture-introduction
date: 2019-08-18T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Clean Architecture
  - Golang
  - DIP
translation_key: golang-clean-architecture-introduction
---


# 概要
GolangでClean Architectureの実装に挑戦したみたので整理しておく。

内容は概ねスライドの内容を踏襲している。

**理解しきれていないところがあったり、自分の解釈、考えを記述しているので、正しくない部分もあるかもしれない。**

# スライド
LTをする機会があったのでスライドを貼っておく。

[Dive to clean architecture with golang](https://speakerdeck.com/bmf_san/dive-to-clean-architecture-with-golang)

# ソース
ソースはこれ。

[bmf-san/go-clean-architecture-web-application-boilerplate](https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate)

MVCのパターンでの実装もtagを切って残してある。

[1.0.0 - MVC](https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate/releases/tag/1.0.0)

# 背景 
[github - bmf-san/Rubel](https://github.com/bmf-san/Rubel)という今はメンテはしていないが、このブログを運用しているCMSアプリケーションがある。

このアプリケーションをリプレースすべく、選定した言語がgoで、アーキテクチャも見直そうということでClean Architectureを採用する方針となった。

なぜClean Architectureを採用しようと考えたかというと、個人で長い付き合いをしていくことのできるアプリケーションのアーキテクチャとして、

ライブラリやその他技術に依存しないようなアーキテクチャパターンが最適解なのではないかと考えたからである。

RubelはLaravelやReactといったフレームワークを採用しているが、フレームワークにどっしりと乗っかる形で実装してしまっているため、比較的モダンで変化（バージョンアップ）の早いそれらのフレームワークのバージョンアップに追随していく時間が惜しく感じた。

本来、CMSの機能追加や機能改善に力を注ぎたいはずが、今後長く運用していていきたいアプリケーションで、本質的ではない部分の開発に時間を注ぐのは合理的だと考えることができなかった。

フレームワークやライブラリ、その他技術への依存を極力減らして、goの標準ライブラリを十分に使いつつ、開発していくことができれば、保守性の高いアプリケーションがつくれるのではないかと考えた。

スクラッチこそ大正義みたいな気概を持っていたりするが、サービス開発のようなビジネス要件に即対応が求められるようなそういったアプリケーションではないのと、開発目的が学習要素を含む部分もあるのである程度理にかなっているとは思う。

自分の目的としては個人開発で取るべき最適に近い戦略が取れているような気はしているが、運用フェーズに乗っかってからでないと見えないところはまぁあるだろうと思っている。

現在鋭意開発中のRubelリプレースはこちら。

[github - bmf-san/Gobel](https://github.com/bmf-san/Gobel)

# 目次
- Clean Architectureとは何か？
- Clean Architectureの実装（実装方法は詳細に書いていない）
- 所感

# Clean Architectureとは何か？
## システムアーキテクチャの歴史
Clean Architectureの考えが生まれる前まで、過去いくつかのアーキテクチャのアイデアが存在していた。

- Hexagonal Architecture(Ports and Adapters)
- Onion Architecture
- Screaming Architecture
- DCI
- BCE
- etc...

これらのアイデアは共通して「関心事の分離」という目的を持っていて、

- フレームワークからの独立
- テスト可能
- UIからの独立
- データベースからの独立
- その他の技術からの独立

といったあらゆるものへの依存性脱却とテスタビリティを追求していく。

## Clean Architecture
Clean Architectureで調べるとよく見るあの図はこちらの元ネタを参照。

[cleancoder.com - The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

各レイヤーについて説明していく。

### Entities
- エンティティは最も重要なビジネスルールをカプセル化する
  - ex. メソッドが用意されたオブジェクトや一連のデータ構造と関数など

### Use Cases
- ユースケースはアプリケーションの特定のビジネスルールを含む

### Interface Adapters
- インターフェースアダプターはエンティティとユースケースのためにデータ変換をするアダプター。

### Frameworks and Drivers
- フレームワークとドライバーはフレームワークやデータベースといったツールで構成される。

## レイヤー間のルール
上記のレイヤー間の制約について。

- 4つのレイヤーだけどは限らない。必要であればレイヤーを増減してよい。
- 内側のレイヤーは外側のレイヤーのことを知らない。
    - →依存の方向性は外側から内側に向かうようにする。

# Clean Architectureの実装

## ディレクトリ構成
冒頭で紹介したソースと同じだが再掲。

[bmf-san/go-clean-architecture-web-application-boilerplate](https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate)

```
./app/
├── database
│   ├── migrations
│   │   └── schema.sql
│   └── seeds
│       └── faker.sql
├── domain
│   ├── post.go
│   └── user.go
├── go_clean_architecture_web_application_boilerplate
├── infrastructure
│   ├── env.go
│   ├── logger.go
│   ├── router.go
│   └── sqlhandler.go
├── interfaces
│   ├── post_controller.go
│   ├── post_repository.go
│   ├── sqlhandler.go
│   ├── user_controller.go
│   └── user_repository.go
├── log
│   ├── access.log
│   └── error.log
├── main.go
└── usecases
    ├── logger.go
    ├── post_interactor.go
    ├── post_repository.go
    ├── user_interactor.go
    └── user_repository.go

8 directories, 22 files
```

レイヤーとディレクトリの対応は以下ような形になっている。

| Layer                | Directory      |
|----------------------|----------------|
| Frameworks & Drivers | infrastructure |
| Interface            | interfaces     |
| Usecases             | usecases       |
| Entities             | domain         |


## DIP
Clean Architectureを実装する前にDIP（依存関係逆転の原則）というルールを知っておく必要がある。

SOLID原則の一つで、抽象は詳細に依存すべきではないというモジュール間の制約についてルールである。

このルールの詳細については割愛するが、Clean Architectureの文脈では、このルールは依存方向を外側から内側に保つため、インターフェースを活用することでDIPを守り、レイヤー間の制約も守る。　

愚直に各レイヤーのルールに従って実装すると依存方向が内側から外側に向いてしまうような事態が発生する。

その事態のときにインターフェースを定義し、抽象への依存をすることで依存方向を守っていく、というのは実装の肝になる部分である。

## Accept interfaces, return structs
Golangには「インターフェースを受け入れて、構造体を返す」という考え方がある。

これはDIPの実装に親和性がある考え方だと思う。

```golang
package examples

// Logger is an interface which will be used for an argument of a function.
type Logger interface {
	Printf(string, ...interface{})
}

// FooController is a struct which will be returned by function.
type FooController struct {
	Logger Logger
}

// NewFooController is a function for an example, "Accept interfaces, return structs".
// Also, this style of a function take on a role of constructor for struct.
func NewFooController(logger Logger) *FooController {
	return &FooController{
		Logger: logger,
	}
}
```

Golangではよく見かける基本的な実装パターンかと思う。

インターフェースに依存させることで変更に強く、テストの書きやすいコードを書くことができる（はず）

## DIP in golang

GolangでのDIP例。

DIPではないコード。

```golang
package examples

// sqlHandler is a struct for handling sql.
type sqlHandler struct{}

// Execute is a function for executing sql.
func (sqlHandler *sqlHandler) Execute() {
	// do something...
}

// FooRepository is a struct depending on details.
type FooRepository struct {
	sqlHandler sqlHandler
}

// Find is a method depending on details.
func (ur *FooRepository) Find() {
	// do something
	ur.sqlHandler.Execute()
}
```

DIPを考慮したコード。

```golang
package examples

// SQLHandler is an interface for handling sql.
type SQLHandler interface {
	Execute()
}

// sqlHandler is a struct which will be returned by function.
type sqlHandler struct{}

// NewSQLHandler is a function for an example of DIP.
// This function depend on abstruction(interface).
// This pattern is an idiom of constructor in golang.
// You can do DI(Dependency Injection) by using nested struct.
func NewSQLHandler() SQLHandler {
	// do something ...

	// sqlHandler struct implments SQLHandler interface.
	return &sqlHandler{}
}

// Execute is a function for executing sql.
// A sqlHanlder struct implments a SQLHandler interface by defining Execute().
func (s *sqlHandler) Execute() {
	// do something...
}

// FooRepository is a struct depending on an interface.
type FooRepository struct {
	SQLHandler SQLHandler
}

// Find is a method of FooRepository depending on an interface.
func (ur *FooRepository) Find() {
	// do something
	ur.SQLHandler.Execute()
}
```

インターフェースを間に挟むことで依存関係が変化し、結果として依存の方向性が逆転するような形になっている。

Before

```
SQLHandler
　　↑
FooRepository
```

After

```
SQLHandler
   ↓   
SQLHandler Interface
   ↑
FooRepository
``` 

Clean Architectureの実例では、infrastructureとinterfacesのコードがそれに当たる。
[bmf-san/go-clean-architecture-web-application-boilerplate](https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate)

## コードリーディング
実際にClean Architectureに取り組む際はいきなり実装しないでまずはコードリーディングや写経から入ると実装を理解しやすいのではないかと思う。

コードリーディングする際は自分は外側から内側に向かってコードを読んでいくのがわかりやすいと感じた。

main.go
  ↓
router.go・・・Infrastructure
　↓
user_controller.go・・・Interfaces
　↓
user_interactor.go・・・Use Cases
　↓
user_repository.go・・・Use Cases
　↓
user.go・・・Domain

# 所感
- Golangの経験が浅いのでinterfaceやstructといった言語仕様について何度か学び直しを繰り返した。
- 「これはどこにかくの？」みたいな迷いの部分はアーキテクトがリードして決めていくのが良いのかなと思った。
    - Clean Architectureを採用する際はアーキテクトという役割を果たす人がチームに一人はいるべきなんじゃないかと感じた。
        - Clean Architectureに限った話ではないかと思うが、、、
- Clean Architectureは実装パターンというより考え方だと思うので、幅広く色んなアーキテクチャのパターンについても研究する必要があると感じた。
- モノリスで戦う前提なところがあるのかなとちょっと感じた。
    - マイクロサービスだったら学習コストをもっと下げた捨てやすいアーキテクチャのパターンが好まれのではないかなと思った。
- A framework is just a tool, not a way of life.
    - 「Clean Architecture 達人に学ぶソフトウェアの構造と設計」の原著にある一節。
    - いい言葉ダナァ

# 参考
- [github - manuelkiessling/go-cleanarchitecture](https://github.com/manuelkiessling/go-cleanarchitecture)
- [github - rymccue/golang-standard-lib-rest-api](https://github.com/rymccue/golang-standard-lib-rest-api)
- [github - hirotakan/go-cleanarchitecture-sample](https://github.com/hirotakan/go-cleanarchitecture-sample)
- [Recruit Technologies - Go言語とDependency Injection](https://recruit-tech.co.jp/blog/2017/12/11/go_dependency_injection/)
- [Clean ArchitectureでAPI Serverを構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
- [github - ponzu-cms/ponzu](https://github.com/ponzu-cms/ponzu)
- [クリーンアーキテクチャの書籍を読んだのでAPIサーバを実装してみた](https://qiita.com/yoshinori_hisakawa/items/f934178d4bd476c8da32)
- [Go × Clean Architectureのサンプル実装](http://nakawatch.hatenablog.com/entry/2018/07/11/181453)
- [Uncle Bob – Payroll Case Study (A full implementation)](http://cleancodejava.com/uncle-bob-payroll-case-study-full-implementation/)



