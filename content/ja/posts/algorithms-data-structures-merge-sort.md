---
title: アルゴリズムとデータ構造 - マージソート
description: アルゴリズムとデータ構造 - マージソート
slug: algorithms-data-structures-merge-sort
date: 2020-02-01T00:00:00Z
author: bmf-san
categories:
  - アルゴリズムとデータ構造
tags:
  - マージソート
translation_key: algorithms-data-structures-merge-sort
---


# 概要
[アルゴリズム図鑑](https://www.shoeisha.co.jp/book/detail/9784798149776)を参考に、アルゴリズムとデータ構造を学ぶ。

実装は[github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master)にも置いてある。

# マージソート
- データ列が分割できなくなるまで（要素が1つ）再帰的に分割を行い、分割されたデータを複数回マージを繰り返していくことによってソートする
- 分割統治法に基づくソート
  - 大きな問題を小さな問題に分割する

# 計算時間
- 最悪計算時間
  - O(n log n)

# 実装
```golang
// cf. https://github.com/TheAlgorithms/Go/blob/master/sorts/merge_sort.go
package main

func merge(a []int, b []int) []int {
	r := make([]int, len(a)+len(b))
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}

	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r
}

func mergeSort(n []int) []int {
	if len(n) < 2 {
		return n
	}

	var middle = len(n) / 2
	a := mergeSort(n[:middle])
	b := mergeSort(n[middle:])
	return merge(a, b)
}
```

- cf. [アルゴリズムとデータ構造 - ヒープ](https://bmf-tech.com/posts/%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0%E3%81%A8%E3%83%87%E3%83%BC%E3%82%BF%E6%A7%8B%E9%80%A0%20-%20%E3%83%92%E3%83%BC%E3%83%97)

# 参考
- [アルゴリズムとデータ構造](http://www-ikn.ist.hokudai.ac.jp/~arim/pub/algo/algo6.pdf)



