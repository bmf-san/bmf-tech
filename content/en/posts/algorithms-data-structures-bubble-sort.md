---
title: Algorithms and Data Structures - Bubble Sort
description: 'Learn Bubble Sort, a comparison-based sorting algorithm with O(n²) time complexity. Covers the adjacent-element swap logic and a complete Go implementation.'
slug: algorithms-data-structures-bubble-sort
date: 2020-02-01T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Bubble Sort
translation_key: algorithms-data-structures-bubble-sort
---



# Overview
Learn algorithms and data structures with reference to the [Algorithm Encyclopedia](https://www.shoeisha.co.jp/book/detail/9784798149776).

The implementation is also available on [github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master).

# Bubble Sort
- A sorting method that arranges data in ascending or descending order
- For all elements, compare adjacent elements and swap them if they are in the wrong order, repeating this operation `number of elements - 1` times

# Computational Time
- Worst-case, best-case, and average-case computational time
  - O(n²)

# Implementation
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

- Loop through all elements, and within that, loop and compare adjacent elements

# References
- ~~Algorithms and Data Structures~~
