---
title: "アルゴリズムとデータ構造 - キュー"
slug: "post-211"
date: 2019-11-17
author: bmf-san
categories:
  - "アルゴリズムとデータ構造"
tags:
  - "キュー"
draft: false
---

# 概要
[アルゴリズム図鑑](https://www.shoeisha.co.jp/book/detail/9784798149776)を参考に、アルゴリズムとデータ構造を学ぶ。

実装は[github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master)にも置いてある。

# キュー
- 常に先に追加されたデータからしかアクセスできないようにデータを一列に並べた構造
  - スタックとは追加と削除の方向が逆になる。
  - FIFO(First In First Out)
    - 先入れ先出し
- 待ち行列ともいう。
- データの追加をenqueue、削除をdequeueという。

# 計算時間
配列や連結リストなど実装形式による。

# 実装
```golang
package main

// Queue is a queue.
type Queue struct {
	nodes []*Node
}

// Node is a item of a stack.
type Node struct {
	value string
}

// newQueue create a Stack.
func newQueue() *Queue {
	return &Queue{}
}

// enqueue adds an node to the end of the queue.
func (s *Queue) enqueue(n *Node) {
	s.nodes = append(s.nodes, n)
}

// dequeue removes an node from the top of the queue.
func (s *Queue) dequeue() {
	s.nodes = s.nodes[1:len(s.nodes)]
}
```

- 実装内容的にはスタックとほぼ同じ。スライスの添字アクセスが違うだけ。
- ノート
  - ![Image from iOS](/assets/images/posts/post-211/69003840-2b7d0700-094c-11ea-996e-c116235a1dbe.jpg)

# 参考
- [flaviocopes.com - Go Data Structures: Queue](https://flaviocopes.com/golang-data-structure-queue/)
- [Wikipedia - キュー](https://ja.wikipedia.org/wiki/%E3%82%AD%E3%83%A5%E3%83%BC_(%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF))

