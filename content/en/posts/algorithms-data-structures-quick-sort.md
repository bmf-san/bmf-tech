---
title: Algorithms and Data Structures - Quick Sort
description: 'Understand Quick Sort: O(n log n) average, O(n²) worst-case time complexity. Covers pivot selection, partitioning into low/high ranges, and a randomized Go implementation.'
slug: algorithms-data-structures-quick-sort
date: 2020-02-01T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Quick Sort
translation_key: algorithms-data-structures-quick-sort
---



# Overview
Referencing [Algorithm Encyclopedia](https://www.shoeisha.co.jp/book/detail/9784798149776), we learn about algorithms and data structures.

The implementation is also available on [github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master).

# Quick Sort
- Select an appropriate data (pivot) from the data sequence, and move data smaller than the pivot to the front and data larger than the pivot to the back.
- Sort each divided data
- A type of divide and conquer method

# Computational Time
- Worst-case time complexity
  - O(n²)
- Best-case and average-case time complexity
  - O(n log n)

# Implementation
```golang
package main

import (
	"fmt"
	"math/rand"
)

func quickSort(n []int) []int {
	if len(n) <= 1 {
		return n
	}

	pivot := n[rand.Intn(len(n))]

	low := make([]int, 0, len(n))
	high := make([]int, 0, len(n))
	middle := make([]int, 0, len(n))

	for _, i := range n {
		switch {
		case i < pivot:
			low = append(low, i)
		case i == pivot:
			middle = append(middle, i)
		case i > pivot:
			high = append(high, i)
		}
	}

	low = quickSort(low)
	high = quickSort(high)

	low = append(low, middle...)
	low = append(low, high...)

	return low
}

func main() {
	n := []int{2, 5, 7, 1, 3, 9}
	fmt.Println(quickSort(n))
}

```

- cf. [Algorithms and Data Structures - Heap](https://bmf-tech.com/posts/%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0%E3%81%A8%E3%83%87%E3%83%BC%E3%82%BF%E6%A7%8B%E9%80%A0%20-%20%E3%83%92%E3%83%BC%E3%83%97)

# References
- [Algorithms and Data Structures](http://www-ikn.ist.hokudai.ac.jp/~arim/pub/algo/algo6.pdf)
