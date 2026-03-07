---
title: スライディングウィンドウの実装
slug: sliding-window-implementation
date: 2023-08-17T00:00:00Z
author: bmf-san
categories:
  - アルゴリズムとデータ構造
tags:
  - スライディングウィンドウ
translation_key: sliding-window-implementation
---


# スライディングウィンドウとは
配列のサブアレイを”ウィンドウ（サブセット）”をずらすしていくように探索するアルゴリズム。

ウィンドウサイズは固定または動的。

実例としては、レートリミッターで使われたりする。

# 実装
ソースコードは以下。

[sliding_window](https://github.com/bmf-san/road-to-algorithm-master/tree/master/algorithm/sliding_window)

与えられた配列から合計がN以上になるサブアレイを探索する関数。

```go
package main

import "fmt"

func SlidingWindow(s []int, sum int) [][]int {
	rslt := [][]int{}
	windowSum := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < len(s); windowEnd++ {
		windowSum += s[windowEnd]

		// サブアレイを見つける
		for windowSum >= sum {
			rslt = append(rslt, s[windowStart:windowEnd+1])
			windowSum -= s[windowStart]
			windowStart++
		}
	}

	return rslt
}

func main() {
	s := []int{1, 3, 2, 6, 4, 9, 9, 5}
	sum := 9
	subAry := SlidingWindow(s, sum)
	for _, sa := range subAry {
		fmt.Println(sa)
	}
}
```

```
// 出力
[1 3 2 6]
[3 2 6]
[2 6 4]
[6 4]
[4 9]
[9]
[9]
```

SlidingWindow関数の処理の流れとしてはこんな感じ。

1. データの先頭から開始位置までウィンドウを配置
2. ウィンドウ内の要素が条件を満たすかどうかを確認
3. 条件を満たす場合、ウィンドウ内のサブアレイを結果に追加
4. ウィンドウを右に1段階スライド
5. ウィンドウがデータの終わりに達するまで、手順 3 ～ 5 を繰り返す

ウィンドウがずれていくイメージが次の通り。
```
[1 3 2 6] ← 最初に見つかった条件を満たすウィンドウ
[3 2 6] ←条件を満たすウィンドウ内で更に探索

[2 6 4] 　←探索開始位置をズラして次のウィンドウを確保
[6 4]　 ←ウィンドウ内を更に探索

以下繰り返し・・・
[4 9]
[9]

[9]
```

配列を扱う問題で応用が効く。

最適解ではないが、[leetcode.com - two-sum](https://leetcode.com/problems/two-sum)はスライディングウィンドウを使って解くこともできる。

# 追記
この動画が分かりやすかった。

[youtube.com - Solve subarray problems FASTER (using Sliding Windows)](https://www.youtube.com/watch?v=GcW4mgmgSbw)

ウインドウサイズが固定の場合、動的な場合の解説がある。

# 参考
- [www.techinterviewhandbook.org - Array cheatsheet for coding interviews](https://www.techinterviewhandbook.org/algorithms/array/)
- [itnext.io - Sliding Window Algorithm Technique](https://itnext.io/sliding-window-algorithm-technique-6001d5fbe8b3)
