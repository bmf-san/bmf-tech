---
title: "Template MethodパターンとStarategyパターンについて"
slug: "template-method-strategy-patterns"
date: 2023-08-20
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Strategy"
  - "Template Method"
draft: false
---

# 概要
GoFの振る舞いに関するパターンであるTemplate MethodパターンとStrategyパターンについてまとめる。

# Template Methodパターンとは
大枠の処理を上位のクラスで決めておき、具体的な処理の流れを下位のクラスに任せる設計パターン。

```go
package main

type Game interface {
	Init()
	Start()
	End()
}

type BaseBall struct{}

func (b *BaseBall) Init() {
	println("BaseBall Init")
}

func (b *BaseBall) Start() {
	println("BaseBall Start")
}

func (b *BaseBall) End() {
	println("BaseBall End")
}

func (b *BaseBall) Play() {
	b.Init()
	b.Start()
	for i := 0; i < 9; i++ {
		println("Top & Bottom")
	}
	b.End()
}

func main() {
	b := &BaseBall{}
	b.Play()
}
```

# Strategyパターンとは
実行時に処理を選択することができるような設計パターン。

Templateパターンと似ているが、Strategyパターンは処理の全てをまとめて切り替えるような構成で、Templateパターンは特定の処理は固定で他が可変といったイメージ。

```go
package main

type PaymentStrategy interface {
	Pay(amount int)
}

type CreditCard struct{}

func (cc *CreditCard) Pay(amount int) {
	println("CreditCard Pay")
}

type Cash struct{}

func (c *Cash) Pay(amount int) {
	println("Cash Pay")
}

type Cart struct {
	paymentMethod PaymentStrategy
}

func (c *Cart) Checkout(amount int) {
	c.paymentMethod.Pay(amount)
}

func main() {
	cc := &CreditCard{}
	c := &Cash{}

	cart := &Cart{paymentMethod: cc}
	cart.Checkout(100)

	cart.paymentMethod = c
	cart.Checkout(100)
}
```

Template MethodパターンもStrategyパターンも継承を使うか委譲を使う方は実装に任させれているので、どちらを使わないといけないということはない。

# 参考
- [github.com - crazybber/awesome-patterns](https://github.com/crazybber/awesome-patterns)
- [ja.wikipedia.org - Template Method パターン](https://ja.wikipedia.org/wiki/Template_Method_%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3)
- [ja.wikipedia.org - Strategy パターン](https://ja.wikipedia.org/wiki/Strategy_%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3)
- [www.techscore.com - 3. Template Method パターン](https://www.techscore.com/tech/DesignPattern/TemplateMethod)
- [www.techscore.com - 10. Strategy パターン](https://www.techscore.com/tech/DesignPattern/Strategy)
