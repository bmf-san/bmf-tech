---
title: Implementation of Count Sort
slug: counting-sort-implementation
date: 2023-06-24T00:00:00Z
author: bmf-san
categories:
  - Algorithms
  - Data Structures
tags:
  - Count Sort
translation_key: counting-sort-implementation
---

# What is Count Sort
A unique sorting algorithm that sorts without comparisons.

It can sort by counting the occurrences of elements.

I found it fascinating that sorting can be done by counting, so I researched it.

# Prerequisites
You need to know about cumulative sums.

cf. [qiita.com - What is Cumulative Sum (For Beginners)](https://qiita.com/xryuseix/items/1059101a31107ba330d4)

# Implementation
The source code is as follows.

[count_sort](https://github.com/bmf-san/road-to-algorithm-master/tree/master/algorithm/sort/count_sort)

```go
package main

import "fmt"

func countSort(s []int, maxVal int) []int {
	count := make([]int, maxVal+1)

	// Count
	for _, num := range s {
		count[num]++
	}

	// Calculate cumulative sum
	// ex. [1, 2, 3, 4, 5] → [1, 3, 6, 10, 15]
	for i := 1; i <= maxVal; i++ {
		count[i] += count[i-1]
	}

	sorted := make([]int, len(s))

	// Traverse the original array from the end and place elements in the sorted result
	for i := len(s) - 1; i >= 0; i-- {
		num := s[i]
		count[num]--
		sorted[count[num]] = num
	}

	return sorted
}

func main() {
	s := []int{5, 3, 2, 1, 2, 3, 4}
	maxVal := 5
	r := countSort(s, maxVal)
	fmt.Println(r)
}
```

```go
// Output
[1 2 2 3 3 4 5]
```

The flow of the process is as follows:
1. Count the occurrences of elements in the array using the indices of the array.
2. Calculate the cumulative sum.
3. Traverse from the end of the array to obtain the sorted result.

It feels very mathematical and not intuitive at all.

# References
- [en.wikipedia.org - prefix sum](https://en.wikipedia.org/wiki/Prefix_sum)
- [www.techinterviewhandbook.org - sorting-searching](https://www.techinterviewhandbook.org/algorithms/sorting-searching/)
- [qiita.com - What is Cumulative Sum (For Beginners)](https://qiita.com/xryuseix/items/1059101a31107ba330d4)