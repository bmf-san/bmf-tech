---
title: About Template Method and Strategy Patterns
description: An in-depth exploration of About Template Method and Strategy Patterns, covering design principles, trade-offs, and practical applications.
slug: template-method-strategy-patterns
date: 2023-08-20T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Strategy
  - Template Method
translation_key: template-method-strategy-patterns
---


# Overview
This post summarizes the Template Method and Strategy patterns, which are behavioral patterns from the GoF.

# What is the Template Method Pattern?
The Template Method pattern is a design pattern where the overall process is defined in a higher-level class, and the specific flow of processing is delegated to lower-level classes.

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

# What is the Strategy Pattern?
The Strategy pattern is a design pattern that allows the selection of processing at runtime.

While similar to the Template pattern, the Strategy pattern is structured to switch all processing collectively, whereas the Template pattern has fixed specific processes with others being variable.

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

Both the Template Method and Strategy patterns allow for the use of inheritance or delegation, so there is no strict requirement to use one over the other.

# References
- [github.com - crazybber/awesome-patterns](https://github.com/crazybber/awesome-patterns)
- [ja.wikipedia.org - Template Method Pattern](https://ja.wikipedia.org/wiki/Template_Method_%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3)
- [ja.wikipedia.org - Strategy Pattern](https://ja.wikipedia.org/wiki/Strategy_%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3)
- [www.techscore.com - 3. Template Method Pattern](https://www.techscore.com/tech/DesignPattern/TemplateMethod)
- [www.techscore.com - 10. Strategy Pattern](https://www.techscore.com/tech/DesignPattern/Strategy)
