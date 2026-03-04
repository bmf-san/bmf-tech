---
title: "アルゴリズムとデータ構造 - 挿入ソート"
slug: "algorithms-data-structures-insertion-sort"
date: 2020-02-01
author: bmf-san
categories:
  - "アルゴリズムとデータ構造"
tags:
  - "挿入ソート"
draft: false
---

# 概要
[アルゴリズム図鑑](https://www.shoeisha.co.jp/book/detail/9784798149776)を参考に、アルゴリズムとデータ構造を学ぶ。

実装は[github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master)にも置いてある。

# 挿入ソート
- データ列の先頭から順番にソートしていく
- ソート済みと未ソートでそれぞれ部分列に分けられる
    - 1回目：0番目をソート済みとするので何もしない
	- 2回目：0番目と1番目を比較して順序が逆なら入れ変える
	- 3回目：0番目から1番目までのデータ列と比較、順序入れ替え
	- 4回目：0番目から2番目までのデータ列と比較、順序入れ替え
	- 以下、未ソート部分がなくなるまで繰り返す


# 計算時間
- O(n²)

# 実装
```golang
package main

import "fmt"

func insertionSort(n []int) []int {
	for i := 1; i < len(n); i++ {
		for j := 0; j < i; j++ {
			if n[i-j-1] > n[i-j] {
				n[i-j-1], n[i-j] = n[i-j], n[i-j-1]
			} else {
				break
			}
		}
	}

	return n
}

func main() {
	n := []int{2, 1, 5, 7, 9}
	fmt.Println(insertionSort(n))
}
```

- 要素を順番に処理していって、スワップするだけなので単純。

# 参考
- [アルゴリズムとデータ構造](http://www-ikn.ist.hokudai.ac.jp/~arim/pub/algo/algo6.pdf)

