---
title: アルゴリズムとデータ構造 - リスト
description: アルゴリズムとデータ構造 - リスト
slug: algorithms-data-structures-list
date: 2019-10-18T00:00:00Z
author: bmf-san
categories:
  - アルゴリズムとデータ構造
tags:
  - 連結リスト
  - 片方向リスト
translation_key: algorithms-data-structures-list
---


# 概要
[アルゴリズム図鑑](https://www.shoeisha.co.jp/book/detail/9784798149776)を参考に、アルゴリズムとデータ構造を学ぶ。

実装は[github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master)にも置いてある。

# リスト（線形リストの片方向リスト）
- データを一直線上に並べた構造
  - 各ノードは次のノードへのポインタを持つ
- データの追加や削除は容易だが、アクセスには時間がかかる
- リストでは、データは連続したメモリ領域に格納される必要はない
  - 一般的には離れた領域に格納される

# 計算時間
リストに格納されているデータ数をnとする。

## データへのアクセス
- O(n)
  - データの先頭から順次アクセス（シーケンシャルアクセス）する必要があるため、線形時間となる

## データの追加
- O(₁)
  - 追加箇所へのデータアクセスは完了しているという前提で、ポインタを2つ差し替えるだけなので定数時間で済む

## データの削除
- データの追加と同様

# 実装
## 線形リスト
### 片方向リスト

```golang
package main

import (
	"errors"
	"fmt"
)

// A node is a node of list.
type node struct {
	value string
	next  *node
}

// A list is a singly linked list.
type list struct {
	head *node
}

// add add a node to tail of a list.
func (l *list) add(newn *node) {
	if l.head == nil {
		l.head = newn
		newn.next = nil
		return
	}

	// sequential access
	for n := l.head; n != nil; n = n.next {
		if n.next == nil {
			n.next = newn
			return
		}
	}

	return
}

// insert a node before a particular node of a list.
func (l *list) insert(newn *node, v string) error {
	if l.head == nil {
		return errors.New("a target node is not exists")
	}

	// sequential access
	for n := l.head; n.next != nil; n = n.next {
		if n.next.value == v {
			newn.next = n.next
			n.next = newn
			return nil
		}
	}

	return errors.New("a target node is not exists")
}

// display display all nodes of a list.
func (l *list) display() {
	// sequential access
	for n := l.head; n != nil; n = n.next {
		fmt.Println(n.value, n.next)
	}
}

func main() {
	l := &list{}

	first := &node{"first", nil}
	second := &node{"second", nil}
	third := &node{"third", nil}

	l.add(first)
	l.add(second)
	l.add(third)

	between := &node{"between", nil}
	l.insert(between, "second")

	l.display()

	fmt.Printf("%#v\n", l)
}
```

- 構造体Listにはリストの先頭ノードを格納しておく
  - リストのデータアクセスを順次アクセスで行うため
- add
  - リストの末尾にノードを追加するメソッド
- insert
  - リストの指定したノードの手前にノードを追加追加するメソッド
    - 追加する位置を特定する→追加位置の一つ手前のポインタと追加するノードのポインタを調整
    - 追加する位置を特定する時は、ループ内で次のノードのポインタを参照して次のノードの値が指定した値とマッチするか判定する
- ノート
![singly_linked_list](/assets/images/posts/algorithms-data-structures-list/67028620-a0d79b00-f146-11e9-8d47-a5d0d933d15d.jpg)

# 参考
- [Naim Ibrahim - Golang singly linked list](https://www.naimibrahim.me/2019/06/08/golang-singly-linked-list/) 
    - 実装が理解しやすかった

# 関連
- [bmf-tech.com - O（オーダー）記法とアルゴリズムの計算量の求め方](https://bmf-tech.com/posts/O（オーダー）記法とアルゴリズムの計算量の求め方)
