---
title: アルゴリズムとデータ構造 - 選択ソート
slug: algorithms-data-structures-selection-sort
date: 2020-02-01T00:00:00Z
author: bmf-san
categories:
  - アルゴリズムとデータ構造
tags:
  - 選択ソート
translation_key: algorithms-data-structures-selection-sort
---


# 概要
[アルゴリズム図鑑](https://www.shoeisha.co.jp/book/detail/9784798149776)を参考に、アルゴリズムとデータ構造を学ぶ。

実装は[github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master)にも置いてある。

# 選択ソート
- データを昇順または降順に並べ変えるソートの一つ
- １番目の要素と2番目以降の要素の中で最小の値を比較して、順序が逆なら入れ変えを行う、という操作をデータ列の最後の一つ手前まで繰り返す

# 計算時間
- 最良計算時間、平均計算時間
  - バブルソートと同じくО(n²)

# 実装
```golang
package main

func selectionSort(n []int) []int {
	for i := 0; i < len(n); i++ {
		min := i

		// Compare the smallest value in the data with the first value
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[min] {
				min = j
			}
		}

		// Swap
		n[i], n[min] = n[min], n[i]
	}

	return n
}

func main() {
	n := []int{2, 1, 5, 7, 9}
	fmt.Println(selectionSort(n))
}
```

# 参考
- [アルゴリズムとデータ構造](http://www-ikn.ist.hokudai.ac.jp/~arim/pub/algo/algo6.pdf)

