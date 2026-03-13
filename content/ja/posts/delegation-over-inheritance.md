---
title: 継承より委譲について
description: '継承より委譲が推奨される理由を解説。ダックタイピング・リスコフの置換原則（LSP）・Goのインターフェースを通じて、四角形と正方形の例から設計の落とし穴と委譲の利点を示します。'
slug: delegation-over-inheritance
date: 2025-10-20T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - ダック・タイピング
  - リスコフの置換原則
  - 委譲
  - 継承
  - Golang
translation_key: delegation-over-inheritance
---


## はじめに

オブジェクト指向プログラミング（OOP）は「現実のモノをプログラムで表現する」ための考え方である。しかし、現実の分類や言葉の定義をそのままプログラムへ持ち込むと、思わぬ破綻を招く。

本記事では、「四角形と正方形」という具体例を通じて、以下の3つの重要な概念を解説する。

- **ダックタイピング** - 名前ではなく振る舞いで型を判断する
- **リスコフの置換原則（LSP）** - 振る舞いの互換性を保証する
- **継承より委譲** - 堅牢な設計を実現する

## ダックタイピング - 名前ではなく振る舞いで決まる型

「ダックタイピング（duck typing）」とは、次の哲学的比喩に基づく型の考え方である。

> "それがアヒルのように鳴き、アヒルのように歩くなら、それはアヒルである"

これは、「型名」や「継承関係」によって型を決めるのではなく、そのオブジェクトがどのように振る舞うかによって型を判断するという思想である。

### Goにおけるダックタイピング

Go言語は静的型付きでありながら、この思想を自然に実現している。明示的に「implements」と書かずとも、必要なメソッドを持っていればインターフェースを満たす。

```go
package main

import "fmt"

type Greeter interface {
    Greet() string
}

type User struct {
    Name string
}

func (u User) Greet() string {
    return "Hello, " + u.Name
}

func SayHello(g Greeter) {
    fmt.Println(g.Greet())
}

func main() {
    user := User{Name: "Alice"}
    SayHello(user) // Output: Hello, Alice
}
```

`User`は`Greeter`を宣言的に実装していないが、`Greet()`メソッドを持っているため、`Greeter`として扱える。このように、**振る舞いによる抽象化**がGoのインターフェース設計の基本である。

## リスコフの置換原則 - 振る舞いの互換性を守る

オブジェクト指向の原則のひとつに、**リスコフの置換原則（Liskov Substitution Principle: LSP）**がある。この原則は次のように定義される。

> 派生クラスは、その基底クラスとして置き換えても、プログラムの正しさを損なってはならない。

つまり、子クラスは親クラスとして同じように振る舞えなければならないということだ。ここで重要なのは、「構造的な一致」ではなく「振る舞いの一貫性」である。

### 「振る舞い」とは何か

オブジェクトの振る舞いとは、外部からの操作（メソッド呼び出し）に対してどのような応答をするか、という**動的な性質**である。

たとえば、ある型が「幅と高さを独立に設定できる」ことを契約として提供しているなら、それを破る子クラスはたとえ構造が似ていても置き換え可能ではない。

振る舞いの一貫性は、プログラムの信頼性の根幹に関わる。

## 四角形を継承した正方形 - 典型的なLSP違反

典型的なLSP違反の例として、「四角形（Rectangle）を継承した正方形（Square）」を例に挙げる。

### 問題のある実装

```php
<?php

class Rectangle {
    protected int $width;
    protected int $height;

    public function setWidth(int $w): void {
        $this->width = $w;
    }

    public function setHeight(int $h): void {
        $this->height = $h;
    }

    public function area(): int {
        return $this->width * $this->height;
    }
}
```

これを継承して正方形を実装する。

```php
<?php

class Square extends Rectangle {
    public function setWidth(int $w): void {
        $this->width = $w;
        $this->height = $w;  // 幅と高さを同じに保つ
    }

    public function setHeight(int $h): void {
        $this->width = $h;   // 幅と高さを同じに保つ
        $this->height = $h;
    }
}
```

### 破綻する例

一見正しく見えるが、次のコードで破綻する。

```php
<?php

$r = new Square();
$r->setWidth(5);
$r->setHeight(10);
echo $r->area(); // 期待値: 50、実際: 100
```

`Square`は`Rectangle`として置き換え可能ではない。これは、親クラスが期待する「**幅と高さを独立に変更できる**」という契約を破っているためである。

つまり、構造的には似ていても、振る舞いが一致していない。

## アリストテレス的分類法とOOPのズレ

アリストテレス的分類法では、事物を共通の性質によって分類する。たとえば、「正方形は四角形の一種である」という分類は自然に思える。

しかし、これは**構造的な分類（見た目や性質の共通性）**であり、OOPで必要とされる**振る舞いによる分類（操作に対する応答の一致）**とは異なる。

現実世界での分類関係をそのままプログラムの継承構造に持ち込むと、リスコフの置換原則を破る危険がある。

## 継承より委譲 - Composition over Inheritance

継承は一見便利な再利用手段に見えるが、親クラスの内部構造や振る舞いに強く依存するため、変更に弱く、置換原則を破りやすい。

この問題を回避するために提唱されているのが**「継承より委譲（Composition over Inheritance）」**である。継承ではなく、必要な機能を内部に保持して利用するという設計手法である。

### 委譲の例

```go
package main

import "fmt"

type Logger struct{}

func (l Logger) Log(msg string) {
    fmt.Println(msg)
}

type Server struct {
    Logger // Loggerを「持っている」（埋め込み）
}

func (s Server) Start() {
    s.Log("Starting server...")
}

func main() {
    server := Server{Logger: Logger{}}
    server.Start() // Output: Starting server...
}
```

この設計では、`Server`は`Logger`を継承せずに利用している。これにより依存関係が明示的になり、保守性が高まる。

## 委譲とインターフェースによる解決（Goの場合）

Goでは、継承が存在しないため、このような問題を自然に回避できる。正方形は四角形を「持つ」ことで、同等の機能を実現する。

### 委譲を使った実装

```go
package main

type Rectangle struct {
    Width, Height int
}

func (r Rectangle) Area() int {
    return r.Width * r.Height
}

type Square struct {
    rect Rectangle
}

func (s *Square) SetSize(n int) {
    s.rect.Width = n
    s.rect.Height = n
}

func (s Square) Area() int {
    return s.rect.Area()
}
```

### インターフェースによる抽象化

または、共通のインターフェースを定義して抽象化する。

```go
package main

type Shape interface {
    Area() int
}

type Rectangle struct {
    Width, Height int
}

func (r Rectangle) Area() int {
    return r.Width * r.Height
}

type Square struct {
    Size int
}

func (s Square) Area() int {
    return s.Size * s.Size
}

func PrintArea(shape Shape) {
    println("Area:", shape.Area())
}

func main() {
    rect := Rectangle{Width: 5, Height: 10}
    square := Square{Size: 5}

    PrintArea(rect)   // Area: 50
    PrintArea(square) // Area: 25
}
```

`Rectangle`も`Square`も`Shape`を満たす限り、同一の抽象として扱える。これにより、継承による破綻を防ぎつつ再利用性を確保できる。

## まとめ

| 観点 | 内容 |
|------|------|
| **ダックタイピング** | 振る舞いによって型を判断する考え方 |
| **振る舞い** | オブジェクトが外部操作に対して示す一貫した応答 |
| **リスコフの置換原則** | 子クラスは親クラスとして置き換え可能であるべき |
| **継承より委譲** | 構造ではなく振る舞いの再利用を重視する設計 |
| **アリストテレス的分類** | 構造的な分類をそのままプログラムへ持ち込むと破綻する |
| **四角形と正方形の問題** | 構造的には正しくても振る舞いが一致せずLSP違反となる |

## 結論

オブジェクト指向における「is-a」関係は、哲学的・言語的な分類と異なり、**振る舞いの一貫性によってのみ成立する**。

現実世界では「正方形は四角形の一種」であっても、プログラム上では「四角形のように振る舞えない正方形」は置き換え不可能である。

したがって、OOP設計においては**「継承より委譲」を基本とし、型名ではなく振る舞いに基づく抽象化を行うこと**が堅牢な設計につながる。

## 参考

- [Barbara Liskov, Data Abstraction and Hierarchy, 1987](https://www.cs.cmu.edu/~wing/publications/LiskovWing94.pdf)
- [Robert C. Martin, Design Principles and Design Patterns](https://fi.ort.edu.uy/innovaportal/file/2032/1/design_principles.pdf)
- [Effective Go - Embedding](https://go.dev/doc/effective_go#embedding)


