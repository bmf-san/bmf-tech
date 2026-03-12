---
title: アルゴリズムとデータ構造 - バブルソート
description: アルゴリズムとデータ構造 - バブルソートについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: algorithms-data-structures-bubble-sort
date: 2020-02-01T00:00:00Z
author: bmf-san
categories:
  - アルゴリズムとデータ構造
tags:
  - バブルソート
translation_key: algorithms-data-structures-bubble-sort
---


# 概要
[アルゴリズム図鑑](https://www.shoeisha.co.jp/book/detail/9784798149776)を参考に、アルゴリズムとデータ構造を学ぶ。

実装は[github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master)にも置いてある。

# バブルソート
- データを昇順または降順に並べ変えるソートの一つ
- 全要素に対して、隣合う要素同士を比較し、順序が逆なら入れ替えを行う、という操作を要素数-1回繰り返す

# 計算時間
- 最悪計算時間、最良計算時間、平均計算時間
  - O(n²)

# 実装
```golang
package main

import "fmt"

func bubbleSort(n []int) []int {
	for i := 0; i < len(n)-1; i++ {
		for j := 0; j < len(n)-i-1; j++ {
			// Compare adjacent values
			if n[j] > n[j+1] {
				// Swap adjacent values
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}

	return n
}

func main() {
	n := []int{2, 1, 5, 7, 9}
	fmt.Println(bubbleSort(n))
}
```

- 全要素をループさせて、その中で隣合う要素同士をループ、比較する

# 参考
- [アルゴリズムとデータ構造](http://www-ikn.ist.hokudai.ac.jp/~arim/pub/algo/algo6.pdf)

