---
title: Implementation of Sliding Window
slug: sliding-window-implementation
date: 2023-08-17T00:00:00Z
author: bmf-san
categories:
  - Algorithms
  - Data Structures
tags:
  - Sliding Window
translation_key: sliding-window-implementation
---

# What is Sliding Window
An algorithm that explores subarrays of an array by shifting a "window (subset)".

The window size can be fixed or dynamic.

It is used in examples such as rate limiters.

# Implementation
The source code is as follows.

[sliding_window](https://github.com/bmf-san/road-to-algorithm-master/tree/master/algorithm/sliding_window)

A function that searches for subarrays with a sum greater than or equal to N from the given array.

```go
package main

import "fmt"

func SlidingWindow(s []int, sum int) [][]int {
	rslt := [][]int{}
	windowSum := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < len(s); windowEnd++ {
		windowSum += s[windowEnd]

		// Find subarray
		for windowSum >= sum {
			rslt = append(rslt, s[windowStart:windowEnd+1])
			windowSum -= s[windowStart]
			windowStart++
		}
	}

	return rslt
}

func main() {
	s := []int{1, 3, 2, 6, 4, 9, 9, 5}
	sum := 9
	subAry := SlidingWindow(s, sum)
	for _, sa := range subAry {
		fmt.Println(sa)
	}
}
```

```// Output
[1 3 2 6]
[3 2 6]
[2 6 4]
[6 4]
[4 9]
[9]
[9]
```

The flow of the SlidingWindow function is as follows:

1. Place the window from the start of the data to the starting position.
2. Check if the elements within the window meet the conditions.
3. If the conditions are met, add the subarray within the window to the results.
4. Slide the window one step to the right.
5. Repeat steps 3 to 5 until the window reaches the end of the data.

The image of the window shifting is as follows:
```
[1 3 2 6] ← The first window that meets the condition
[3 2 6] ← Further exploration within the window that meets the condition

[2 6 4]  ← Shift the starting position to secure the next window
[6 4] ← Further exploration within the window

Repeat...
[4 9]
[9]

[9]
```

This can be applied to problems involving arrays.

Although not the optimal solution, [leetcode.com - two-sum](https://leetcode.com/problems/two-sum) can also be solved using the sliding window.

# Postscript
This video was easy to understand.

[youtube.com - Solve subarray problems FASTER (using Sliding Windows)](https://www.youtube.com/watch?v=GcW4mgmgSbw)

It explains both fixed and dynamic window sizes.

# References
- [www.techinterviewhandbook.org - Array cheatsheet for coding interviews](https://www.techinterviewhandbook.org/algorithms/array/)
- [itnext.io - Sliding Window Algorithm Technique](https://itnext.io/sliding-window-algorithm-technique-6001d5fbe8b3)