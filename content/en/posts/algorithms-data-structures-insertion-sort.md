---
title: Algorithm and Data Structures - Insertion Sort
description: An in-depth look at Algorithm and Data Structures - Insertion Sort, covering key concepts and practical insights.
slug: algorithms-data-structures-insertion-sort
date: 2020-02-01T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Insertion Sort
translation_key: algorithms-data-structures-insertion-sort
---



# Overview
Referencing the [Algorithm Encyclopedia](https://www.shoeisha.co.jp/book/detail/9784798149776), we learn about algorithms and data structures.

The implementation is also available on [github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master).

# Insertion Sort
- Sorts sequentially from the beginning of the data series
- Divides into sorted and unsorted subsequences
    - 1st pass: Treat the 0th element as sorted, do nothing
    - 2nd pass: Compare the 0th and 1st elements, swap if in reverse order
    - 3rd pass: Compare with the data series from 0th to 1st, swap if necessary
    - 4th pass: Compare with the data series from 0th to 2nd, swap if necessary
    - Repeat until there are no unsorted parts


# Time Complexity
- O(n²)

# Implementation
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

- Simply processes elements in order and swaps them.

# References
- [Algorithms and Data Structures](http://www-ikn.ist.hokudai.ac.jp/~arim/pub/algo/algo6.pdf)
