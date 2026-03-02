---
title: "連結リストのランナーテクニック"
slug: "post-353"
date: 2023-07-22
author: bmf-san
categories:
  - "アルゴリズムとデータ構造"
tags:
  - "アルゴリズム"
  - "データ構造"
  - "ランナーテクニック"
  - "Tips"
draft: false
---

連結リストの走査で役立つランナーテクニックについてまとめる。

[世界で闘うプログラミング力を鍛える本 ~コーディング面接189問とその解法](https://amzn.to/3q35TCw)で紹介されていて初めて知った。

# ランナーテクニックとは
連結リストの先頭から走査していくポインタと、そのポインタより先から走査していくポインタの2種類を用意して、同時に走査していく方法。

これが何に役立つかというと、例えば次のような例題を解くのに役立つ。

## 例題
単方向連結リストの末尾からn番目の要素を見つけるアルゴリズムを実装しなさい。

```go
package main

import "fmt"

type node struct {
	val  string
	next *node
}

type list struct {
	head *node
}

// 末尾からn番目のノードを探す
func (l *list) search(n int) *node {
	n1 := l.head
	n2 := l.head

       // n1はk番目のノードに設定
	for i := 0; i < n; i++ {
		if n1 == nil {
			return nil
		}
		n1 = n1.next
	}
        
　　　　　　　　// n2は先頭ノードから、n1はk番目のノードから走査する。
        // n1が末尾ノードに到達したらn2は末尾から数えてn番目のノードである。
	for n1 != nil {
		n1 = n1.next
		n2 = n2.next
	}

	return n2
}

func main() {
	l := &list{
		head: &node{
			val: "a",
			next: &node{
				val: "b",
				next: &node{
					val: "c",
					next: &node{
						val: "d",
					},
				},
			},
		},
	}
	fmt.Printf("%+v\n", l.search(1)) // d
	fmt.Printf("%+v\n", l.search(2)) // c
	fmt.Printf("%+v\n", l.search(3)) // b
	fmt.Printf("%+v\n", l.search(4)) // a
}
```

このようにランナーテクニックを使うと時間計算量はO（N）、空間計算量はO(1）で解くことができる。

連結リストのノード数が決まっていればわざわざポインタを２つ用意しなくとも、（全ノード数-n）が末尾からn番目となるので単純に解けるが、そうでない場合はこのように解くか、再帰で解くことになる。再帰の場合は計算量が増えるはず・・

# 所感
コーディングクイズで使えるシーンがありそうなので頭の片隅に留めておきたい。
