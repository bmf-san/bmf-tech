---
title: Algorithms and Data Structures - Merge Sort
description: An in-depth look at Algorithms and Data Structures - Merge Sort, covering key concepts and practical insights.
slug: algorithms-data-structures-merge-sort
date: 2020-02-01T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Merge Sort
translation_key: algorithms-data-structures-merge-sort
---



# Overview
Learning algorithms and data structures with reference to [Algorithm Encyclopedia](https://www.shoeisha.co.jp/book/detail/9784798149776).

The implementation is also available at [github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master).

# Merge Sort
- Recursively divide the data sequence until it can no longer be divided (single element), and sort by repeatedly merging the divided data multiple times
- Sort based on the divide and conquer method
  - Divide a large problem into smaller problems

# Time Complexity
- Worst-case time complexity
  - O(n log n)

# Implementation
```golang
// cf. https://github.com/TheAlgorithms/Go/blob/master/sorts/merge_sort.go
package main

func merge(a []int, b []int) []int {
	r := make([]int, len(a)+len(b))
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}

	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r
}

func mergeSort(n []int) []int {
	if len(n) < 2 {
		return n
	}

	var middle = len(n) / 2
	a := mergeSort(n[:middle])
	b := mergeSort(n[middle:])
	return merge(a, b)
}
```

- cf. [Algorithms and Data Structures - Heap](https://bmf-tech.com/posts/%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0%E3%81%A8%E3%83%87%E3%83%BC%E3%82%BF%E6%A7%8B%E9%80%A0%20-%20%E3%83%92%E3%83%BC%E3%83%97)

# References
- [Algorithms and Data Structures](http://www-ikn.ist.hokudai.ac.jp/~arim/pub/algo/algo6.pdf)

