---
title: Implementing Sliding Window
description: "Implement sliding window algorithms for efficient subarray searching with optimized time complexity and rate limiter applications."
slug: sliding-window-implementation
date: 2023-08-17T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Sliding Window
translation_key: sliding-window-implementation
---



# What is Sliding Window
An algorithm that explores subarrays of an array by shifting a "window (subset)".

The window size can be fixed or dynamic.

A practical example is its use in rate limiters.

# Implementation
The source code is below.

[sliding_window](https://github.com/bmf-san/road-to-algorithm-master/tree/master/algorithm/sliding_window)

A function that searches for subarrays whose sum is greater than or equal to N from a given array.

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

```
// Output
[1 3 2 6]
[3 2 6]
[2 6 4]
[6 4]
[4 9]
[9]
[9]
```

The process flow of the SlidingWindow function is as follows.

1. Place the window from the start of the data to the starting position
2. Check if the elements within the window meet the condition
3. If the condition is met, add the subarray within the window to the result
4. Slide the window one step to the right
5. Repeat steps 3 to 5 until the window reaches the end of the data

The image of the window sliding is as follows.
```
[1 3 2 6] ← The first window found that meets the condition
[3 2 6] ← Further exploration within the window that meets the condition

[2 6 4]  ← Shift the starting position to secure the next window
[6 4]   ← Further exploration within the window

Repeat below...
[4 9]
[9]

[9]
```

It can be applied to problems dealing with arrays.

Although not the optimal solution, [leetcode.com - two-sum](https://leetcode.com/problems/two-sum) can also be solved using the sliding window.

# Additional Notes
This video was easy to understand.

[youtube.com - Solve subarray problems FASTER (using Sliding Windows)](https://www.youtube.com/watch?v=GcW4mgmgSbw)

There is an explanation for when the window size is fixed and when it is dynamic.

# References
- [www.techinterviewhandbook.org - Array cheatsheet for coding interviews](https://www.techinterviewhandbook.org/algorithms/array/)
- ~~itnext.io - Sliding Window Algorithm Technique~~
