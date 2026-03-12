---
title: Implementation of Counting Sort
description: A step-by-step guide on Implementation of Counting Sort, with practical examples and configuration tips.
slug: counting-sort-implementation
date: 2023-06-24T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Counting Sort
translation_key: counting-sort-implementation
---



# What is Counting Sort
Counting Sort is an unusual sorting algorithm that sorts without using comparisons.

It can sort by counting elements instead of comparing them.

I found it fascinating that sorting can be done by counting, so I decided to look into it.

# Prerequisites
You need to know about prefix sums.

cf. [qiita.com - What is a Prefix Sum (For Beginners)](https://qiita.com/xryuseix/items/1059101a31107ba330d4)

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

	// Calculate prefix sums
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

```
// Output
[1 2 2 3 3 4 5]
```

The process flow is as follows:
1. Count the occurrences of each element in the array using the elements as indices
2. Calculate the prefix sums
3. Traverse the array from the end to get the sorted result

It feels very mathematical rather than intuitive.

# References
- [en.wikipedia.org - prefix sum](https://en.wikipedia.org/wiki/Prefix_sum)
- [www.techinterviewhandbook.org - sorting-searching](https://www.techinterviewhandbook.org/algorithms/sorting-searching/)
- [qiita.com - What is a Prefix Sum (For Beginners)](https://qiita.com/xryuseix/items/1059101a31107ba330d4)
