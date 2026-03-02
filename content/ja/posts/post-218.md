---
title: "アルゴリズムとデータ構造 - ヒープ"
slug: "post-218"
date: 2020-01-14
author: bmf-san
categories:
  - "アルゴリズムとデータ構造"
tags:
  - "ヒープ"
draft: false
---

# 概要
[アルゴリズム図鑑](https://www.shoeisha.co.jp/book/detail/9784798149776)を参考に、アルゴリズムとデータ構造を学ぶ。

実装は[github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master)にも置いてある。

# ヒープ
- 優先度付きキュー（priority queue）の一種 
  - 優先度付きキューは、集合（set）を扱うデータ型
    - 集合に含まれる要素は優先度順に取り出される   
    - 集合を扱うデータ型の例：キュー、スタック
- ヒープの種類
  - 最小ヒープ（min heap）
    - 根が常に最小となっているヒープ。親ノードは子ノードより常に小さい。
  - 最大ヒープ（max heap）
    - 根が常に最大の要素となっているヒープ。親ノードは子ノードより常に大きい。

# 計算時間
- 追加・削除共にO(log n)

# 実装
```golang
package main

import "fmt"

// Heap is a heap.
type Heap struct {
	values  []int
	size    int
	maxsize int
}

// newHeap creates a heap.
func newHeap(maxsize int) *Heap {
	return &Heap{
		values:  []int{},
		size:    0,
		maxsize: maxsize,
	}
}

// leaf checks whether index is a leaf.
func (h *Heap) leaf(index int) bool {
	if index >= (h.size/2) && index <= h.size {
		return true
	}
	return false
}

// parent checks whether index is a parent.
func (h *Heap) parent(index int) int {
	return (index - 1) / 2
}

// leftchild checks whether index is a leftchild.
func (h *Heap) leftchild(index int) int {
	return 2*index + 1
}

// rightchild checks whether index is a rightchild.
func (h *Heap) rightchild(index int) int {
	return 2*index + 2
}

// insert inserts a item to a heap.
func (h *Heap) insert(item int) error {
	if h.size >= h.maxsize {
		return fmt.Errorf("Error!")
	}
	h.values = append(h.values, item)
	h.size++
	h.upHeapify(h.size - 1)
	return nil
}

// swap swaps two values.
func (h *Heap) swap(first, second int) {
	temp := h.values[first]
	h.values[first] = h.values[second]
	h.values[second] = temp
}

// upHeapify reconstruct a heap for up.
func (h *Heap) upHeapify(index int) {
	for h.values[index] < h.values[h.parent(index)] {
		h.swap(index, h.parent(index))
	}
}

// downHeapify reconstruct a heap for down.
func (h *Heap) downHeapify(current int) {
	if h.leaf(current) {
		return
	}

	smallest := current
	leftChildIndex := h.leftchild(current)
	rightRightIndex := h.rightchild(current)

	if leftChildIndex < h.size && h.values[leftChildIndex] < h.values[smallest] {
		smallest = leftChildIndex
	}
	if rightRightIndex < h.size && h.values[rightRightIndex] < h.values[smallest] {
		smallest = rightRightIndex
	}
	if smallest != current {
		h.swap(current, smallest)
		h.downHeapify(smallest)
	}
	return
}

// buildMinHeap builds a min heap.
func (h *Heap) buildMinHeap() {
	for index := ((h.size / 2) - 1); index >= 0; index-- {
		h.downHeapify(index)
	}
}

// remove removes a value.
func (h *Heap) remove() int {
	top := h.values[0]
	h.values[0] = h.values[h.size-1]
	h.values = h.values[:(h.size)-1]
	h.size--
	h.downHeapify(0)
	return top
}

func main() {
	inputArray := []int{6, 5, 3, 7, 2, 8}
	h := newHeap(len(inputArray))
	for i := 0; i < len(inputArray); i++ {
		h.insert(inputArray[i])
	}
	h.buildMinHeap()
	for i := 0; i < len(inputArray); i++ {
		fmt.Println(h.remove())
	}
	fmt.Scanln()
}
```

- 最小ヒープ（min heap）の実装
- ノードの追加は常に幅優先順で追加される
- ノード追加後、追加した値が根のノードより小さい場合、根ノードと追加したノードを入れ替える
- [Data Structure Visualizations - Min Heap](https://www.cs.usfca.edu/~galles/visualization/Heap.html)でイメージを確認するとわかりやすい
- 根、親、左の子ノード、右の子ノードの計算が特徴的
  - 配列のインデックスが優先度となるので、計算によってそれらのノードの値が基準となるノードから求めることができる
  - [ヒープをわかりやすく解説してみた](https://medium.com/@yasufumy/data-structure-heap-ecfd0989e5be)の記事がわかりやすい
- [Welcome To Golang By Example - Heap in Golang](https://golangbyexample.com/heap-in-golang/)を参考にした（ほとんど写経..）

# 参考
- [ヒープをわかりやすく解説してみた](https://medium.com/@yasufumy/data-structure-heap-ecfd0989e5be)
- [Welcome To Golang By Example - Heap in Golang](https://golangbyexample.com/heap-in-golang/)



