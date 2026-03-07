---
title: Implementing Backtracking
slug: backtracking-implementation
date: 2023-06-23T00:00:00Z
author: bmf-san
categories:
  - Algorithms
tags:
  - Backtracking
description: Exploring combinations that satisfy given constraints using backtracking.
translation_key: backtracking-implementation
---

# What is Backtracking
An algorithm for exploring combinations that satisfy specified constraints.

It can be used to explore all unique combinations (`nCr`).

# Implementation
Source code is as follows.

[backtrack](https://github.com/bmf-san/road-to-algorithm-master/tree/master/algorithm/backtrack)

The process retrieves N unique subsequences from a given array. Below is an example of 4C3.

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
// Output
[[1 2 3] [1 2 4] [1 3 4] [2 3 4]]
```

The recursive process can feel overwhelming, but thinking of the data as a tree structure makes it easier to understand.

When data consisting of 1, 2, 3, 4 is provided, consider what results are obtained for k values ranging from 0 to 4.

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

In short, k represents the depth of the tree structure, and the 3 in 4C3 can be considered as the depth.

Attempting to solve combinations without backtracking would result in nested loops of `for` statements.

One approach to write this efficiently using recursion is backtracking.

[leetcode.com - subsets](https://leetcode.com/problems/subsets) is a problem with multiple solutions, one of which can be solved using backtracking.

# References
- [ja.wikipedia.org - Backtracking](https://ja.wikipedia.org/wiki/%E3%83%90%E3%83%83%E3%82%AF%E3%83%88%E3%83%A9%E3%83%83%E3%82%AD%E3%83%B3%E3%82%B0)
- [p-www.iwate-pu.ac.jp - Algorithm Theory](http://p-www.iwate-pu.ac.jp/~k-yamada/lecture/algorithm2014/12th-week.pdf)