---
title: Algorithms and Data Structures - Selection Sort
slug: algorithms-data-structures-selection-sort
date: 2020-02-01T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Selection Sort
translation_key: algorithms-data-structures-selection-sort
---



# Overview
Learn algorithms and data structures with reference to [Algorithm Encyclopedia](https://www.shoeisha.co.jp/book/detail/9784798149776).

The implementation is also available on [github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master).

# Selection Sort
- A sorting method that arranges data in ascending or descending order
- Compares the smallest value among the elements from the second onward with the first element, and if the order is reversed, swaps them. This operation is repeated until the second to last element of the data sequence.

# Time Complexity
- Best and average time complexity
  - Same as bubble sort, О(n²)

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
