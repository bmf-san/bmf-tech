---
title: "アルゴリズムとデータ構造 - スタック"
slug: "algorithms-data-structures-stack"
date: 2019-11-17
author: bmf-san
categories:
  - "アルゴリズムとデータ構造"
tags:
  - "スタック"
draft: false
---

# 概要
[アルゴリズム図鑑](https://www.shoeisha.co.jp/book/detail/9784798149776)を参考に、アルゴリズムとデータ構造を学ぶ。

実装は[github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master)にも置いてある。

# スタック
- 常に最新のデータからしかアクセスできないようにデータを一列に並べた構造
  - LIFO(Last In First Out)
    - 後入れ先出し
- 常に最新のデータへアクセスしたいときに便利な構造
- データの追加をPush、削除をPopという。
  - 他にDup、Peek、Swap(またはExchange)、Rotateといった操作がある。
    - cf. [Wikipedia - スタック](https://ja.wikipedia.org/wiki/%E3%82%B9%E3%82%BF%E3%83%83%E3%82%AF)

# 計算時間
配列や連結リストなど実装形式による。

# 実装
```golang
package main

// Stack is a stack.
type Stack struct {
	nodes []*Node
}

// Node is a item of a stack.
type Node struct {
	value string
}

// newStack create a Stack.
func newStack() *Stack {
	return &Stack{}
}

// push adds an node to the top of the stack.
func (s *Stack) push(n *Node) {
	s.nodes = append(s.nodes[:len(s.nodes)], n)
}

// pop removes an node from the top of the stack.
func (s *Stack) pop() {
	s.nodes = s.nodes[:len(s.nodes)-1]
}

```

- Goのスライスに慣れていれば難しいところは特にないはず
- ノート
  - ![Image from iOS (1)](https://user-images.githubusercontent.com/13291041/69003839-2ae47080-094c-11ea-8506-cb733abec36e.jpg)

# 参考
- [Golang program for implementation LIFO Stack and FIFO Queue](https://www.golangprograms.com/golang-program-for-implementation-lifo-stack-and-fifo-queue.html)
