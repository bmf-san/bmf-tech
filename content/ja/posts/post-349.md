---
title: "カウントソートの実装"
slug: "post-349"
date: 2023-06-24
author: bmf-san
categories:
  - "アルゴリズムとデータ構造"
tags:
  - "カウントソート"
draft: false
---

# カウントソートとは
ソートアルゴリズムの中でも比較を使わずにソートする変わった？アルゴリズム。

比較せずに要素をカウントすることでソートができる。

カウントしてソートすることができるというのは不思議！と思ったので調べてみた。

# 前提
累積和について知っておく必要がある。

cf. [qiita.com - 累積和とは（超初心者用）](https://qiita.com/xryuseix/items/1059101a31107ba330d4)

# 実装
ソースコードは以下。

[count_sort](https://github.com/bmf-san/road-to-algorithm-master/tree/master/algorithm/sort/count_sort)

```go
package main

import "fmt"

func countSort(s []int, maxVal int) []int {
	count := make([]int, maxVal+1)

	// カウント
	for _, num := range s {
		count[num]++
	}

	// 累積和計算
	// ex. [1, 2, 3, 4, 5] →　[1, 3, 6, 10, 15]
	for i := 1; i <= maxVal; i++ {
		count[i] += count[i-1]
	}

	sorted := make([]int, len(s))

	// 元の配列を末尾から走査し、要素をソート結果に配置
	for i := len(s) - 1; i >= 0; i-- {
		num := s[i]
		count[num]--
		sorted[count[num]] = num
	}

	return sorted
}

func main() {
	s := []int{5, 3, 2, 1, 2, 3, 4}
	maxVal := 5
	r := countSort(s, maxVal)
	fmt.Println(r)
}
```

```
// 出力
[1 2 2 3 3 4 5]
```

処理の流れとしては、
1. 配列の要素はインデックスとして配列内の要素の出現回数をカウント
2. 累積和を計算
3. 配列末尾から走査してソート結果を得る

直感的ではなく何というかすごく数学的な感じがする。

# 参考
- [en.wikipedia.org - prefix sum](https://en.wikipedia.org/wiki/Prefix_sum)
- [www.techinterviewhandbook.org - sorting-searching](https://www.techinterviewhandbook.org/algorithms/sorting-searching/)
- [qiita.com - 累積和とは（超初心者用）](https://qiita.com/xryuseix/items/1059101a31107ba330d4)
