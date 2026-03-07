---
title: Algorithms and Data Structures - Selection Sort
slug: algorithms-data-structures-selection-sort
date: 2020-02-01T00:00:00Z
author: bmf-san
categories:
  - Algorithms
  - Data Structures
tags:
  - Selection Sort
translation_key: algorithms-data-structures-selection-sort
---

# Overview
Referencing the [Algorithm Encyclopedia](https://www.shoeisha.co.jp/book/detail/9784798149776), we learn about algorithms and data structures.

The implementation is also available at [github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master).

# Selection Sort
- One of the sorting algorithms that rearranges data in ascending or descending order.
- It compares the first element with the smallest value among the subsequent elements, and if the order is reversed, it swaps them, repeating this operation until just before the last element of the data sequence.

# Time Complexity
- Best-case and average-case time complexity:
  - Similar to bubble sort, О(n²)

# Implementation
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

# References
- [Algorithms and Data Structures](http://www-ikn.ist.hokudai.ac.jp/~arim/pub/algo/algo6.pdf)