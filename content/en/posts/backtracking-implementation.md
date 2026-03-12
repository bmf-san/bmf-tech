---
title: Implementing Backtracking
description: A step-by-step guide on Implementing Backtracking, with practical examples and configuration tips.
slug: backtracking-implementation
date: 2023-06-23T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Backtrack
translation_key: backtracking-implementation
---

# What is Backtracking
An algorithm that explores combinations that satisfy specified constraints.

It can be used when exploring all unique combinations (`nCr`).

# Implementation
The source code is as follows.

[backtrack](https://github.com/bmf-san/road-to-algorithm-master/tree/master/algorithm/backtrack)

This process retrieves N unique subsequences from a given array. Below is an example of 4C3.

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

```// Output
[[1 2 3] [1 2 4] [1 3 4] [2 3 4]]
```

In the recursive process, it may feel like the brain's memory is about to burst, but thinking of the data as a tree structure makes it clearer.

When data consisting of 1, 2, 3, and 4 is passed, consider what results can be obtained for k from 0 to 4.

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

In short, k represents the depth of the tree structure, and the 3 in 4C3 can be said to be the depth.

If you try to solve combinations naively without backtracking, it will result in nested loops in a for statement.

One approach to writing it well with recursion is backtracking.

[leetcode.com - subsets](https://leetcode.com/problems/subsets) is a problem with several solutions, but it can be solved using backtracking.

# References
- [ja.wikipedia.org - Backtracking](https://ja.wikipedia.org/wiki/%E3%83%90%E3%83%83%E3%82%AF%E3%83%88%E3%83%A9%E3%83%83%E3%82%AD%E3%83%B3%E3%82%B0)
- [p-www.iwate-pu.ac.jp - Algorithm Theory](http://p-www.iwate-pu.ac.jp/~k-yamada/lecture/algorithm2014/12th-week.pdf)