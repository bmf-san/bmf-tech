---
title: "バックトラッキングの実装"
slug: "post-348"
date: 2023-06-23
author: bmf-san
categories:
  - "アルゴリズムとデータ構造"
tags:
  - "バックトラック"
draft: false
---

# バックトラッキングとは
指定された制約を満たすような組み合わせを探索するアルゴリズム。　

重複しない組み合わせ（`nCr`）を全て探索するようなときに使える。

# 実装
ソースコードは以下。

[backtrack](https://github.com/bmf-san/road-to-algorithm-master/tree/master/algorithm/backtrack)


与えられた配列からN個の重複しないサブシーケンスを取得する処理。
以下は4C3の例。

```go
package main

import "fmt"

func backtrack(rslt *[][]int, tmp []int, items []int, start int, k int) {
	if k == 0 {
		combination := make([]int, len(tmp))
		copy(combination, tmp)
		*rslt = append(*rslt, combination)
		return
	}

	for i := start; i < len(items); i++ {
		tmp = append(tmp, items[i])
		backtrack(rslt, tmp, items, i+1, k-1)
		tmp = tmp[:len(tmp)-1]
	}
}

func get(items []int, k int) [][]int {
	rslt := [][]int{}
	tmp := []int{}
	backtrack(&rslt, tmp, items, 0, k)
	return rslt
}

func main() {
	items := []int{1, 2, 3, 4}
	k := 3
	combination := get(items, k)
	fmt.Println(combination)
}
```

```
// 出力
[[1 2 3] [1 2 4] [1 3 4] [2 3 4]]
```

再帰処理のところで脳内メモリがパンクしそうになるが、データを木構造にして考えると分かりやすい。

1, 2, 3, 4からなるデータが渡されるとき、kが0~4でそれぞれどういう結果が得られるか考えてみる。

```
k = 0
N/A

k = 1
1  2  3  4

k = 2

	   1      2    3
	 / | \   / \   |
	2  3  4 3   4  4

k = 3

	    1     2
	   / \    |
	  2   3   3
	 / \  |   |
	3  4  4   4

k = 4
1
2
3
4
```

要するに、kは木構造の深さで、4C3の3は深さであるといえる。

組み合わせをバックトラッキングを使わずに愚直に解こうとすると、for文の多重ループになってしまう。

それを再帰処理で上手く書くアプローチの一つがバックトラッキングだったりする。

[leetcode.com - subsets](https://leetcode.com/problems/subsets)はいくつか解法がある問題だが、バックトラッキングで解くことができる。

# 参考
- [ja.wikipedia.org - バックトラッキング](https://ja.wikipedia.org/wiki/%E3%83%90%E3%83%83%E3%82%AF%E3%83%88%E3%83%A9%E3%83%83%E3%82%AD%E3%83%B3%E3%82%B0)
- [p-www.iwate-pu.ac.jp - アルゴリズム論](http://p-www.iwate-pu.ac.jp/~k-yamada/lecture/algorithm2014/12th-week.pdf)
