---
title: アルゴリズムとデータ構造 - クイックソート
description: アルゴリズムとデータ構造 - クイックソートについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: algorithms-data-structures-quick-sort
date: 2020-02-01T00:00:00Z
author: bmf-san
categories:
  - アルゴリズムとデータ構造
tags:
  - クイックソート
translation_key: algorithms-data-structures-quick-sort
---


# 概要
[アルゴリズム図鑑](https://www.shoeisha.co.jp/book/detail/9784798149776)を参考に、アルゴリズムとデータ構造を学ぶ。

実装は[github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master)にも置いてある。

# クイックソート
- データ列の中から適当なデータ（ピボット）を選択し、ピボットより小さいデータを前方、大きいデータを後方に移動させる。
- 分割されたデータをそれぞれソートする
- 分割統治法の一種

# 計算時間
- 最悪計算時間
  - O(n²)
- 最良計算時間、平均計算時間
  - O(n log n)

# 実装
```golang
package main

import (
	"fmt"
	"math/rand"
)

func quickSort(n []int) []int {
	if len(n) <= 1 {
		return n
	}

	pivot := n[rand.Intn(len(n))]

	low := make([]int, 0, len(n))
	high := make([]int, 0, len(n))
	middle := make([]int, 0, len(n))

	for _, i := range n {
		switch {
		case i < pivot:
			low = append(low, i)
		case i == pivot:
			middle = append(middle, i)
		case i > pivot:
			high = append(high, i)
		}
	}

	low = quickSort(low)
	high = quickSort(high)

	low = append(low, middle...)
	low = append(low, high...)

	return low
}

func main() {
	n := []int{2, 5, 7, 1, 3, 9}
	fmt.Println(quickSort(n))
}

```

- cf. [アルゴリズムとデータ構造 - ヒープ](https://bmf-tech.com/posts/%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0%E3%81%A8%E3%83%87%E3%83%BC%E3%82%BF%E6%A7%8B%E9%80%A0%20-%20%E3%83%92%E3%83%BC%E3%83%97)

# 参考
- [アルゴリズムとデータ構造](http://www-ikn.ist.hokudai.ac.jp/~arim/pub/algo/algo6.pdf)



