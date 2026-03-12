---
title: About the Two Pointer Technique
description: An in-depth look at About the Two Pointer Technique, covering key concepts and practical insights.
slug: sliding-window-technique
date: 2023-08-01T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Two-Pointer Technique
  - Two-Pointer Approach
translation_key: sliding-window-technique
---



# Overview
This post summarizes the Two Pointer Technique.

In English, it is called the Two Pointer Approach or Two Pointer Technique.

# What is the Two Pointer Technique?
It is an algorithm that maintains the indices of the right and left ends of a dataset (such as a sequence or string) and moves the indices based on certain conditions to search for data that meets those conditions.

It is useful when you want to search for data that meets specific conditions within a range.

# Example Problem
Write a function to find how many pairs of numbers in array `n` are less than a specified number `m`.

For example, if `n = [1, 3, -1, 2]` and `m = 4`, there are four pairs that meet the criteria: (1, -1), (1, 2), (3, -1), and (-1, 2), so the function should return 4.

The code for a straightforward implementation of the function is as follows:

```go
func findPairs(n []int, m int) int {
    rslt := 0
    for i := 0; i < len(n); i++ {
        for j := i+1; j < len(n); j++ {
            if (n[i] + n[j]) < m {
                rslt++
            }
        }
    }
    return rslt
}
```

This results in a time complexity of O(N^2), so we use the Two Pointer Technique to reduce it to O(N log N).

```go
func findPairs(n []int, m int) int {
    sort.Ints(n)
    cnt, l, r := 0, 0, len(n)-1
    for l < r {
        if n[l] + n[r] < m {
            cnt += r-l
            l++
            continue
        }
        r--
    }
    return cnt
}
```

By repeating until the left and right indices overlap, you can explore all pairs that meet the conditions.

Sorting is done to make it easier to efficiently find pairs.

Although `r-l` might seem counterintuitive, since the code is for finding pairs, it focuses on counting pairs by indices rather than the values themselves in the array, resulting in this kind of code.

# References
- [algodaily.com - The Two Pointer Technique](https://algodaily.com/lessons/using-the-two-pointer-technique/go)
- [www.geekforgeeks.org - Two Pointers Technique](https://www.geeksforgeeks.org/two-pointers-technique/)
- [www.codingninjas.com - Two Pointer Approach](https://www.codingninjas.com/studio/library/what-is-a-two-pointer-technique)
- [jaigotec.com - [Algorithm(Ruby)] Explanation of the Two Pointer Technique](https://jaigotec.com/algorithm_two-pointer-techinique/)
- [qiita.com - Explanation of the Two Pointer Technique and Problems Using It](https://qiita.com/drken/items/ecd1a472d3a0e7db8dce)
- [na.fuis.u-fukui.ac.jp - Two Pointer Technique](https://na.fuis.u-fukui.ac.jp/~hirota/course/2022_Exp2_Programming/03-1_Shakutori.pdf)
- [paiza.hatenablog.com - [Cumulative Sum, Two Pointer Technique] Algorithm Diagram for Beginners](https://paiza.hatenablog.com/entry/2015/01/21/%E3%80%90%E7%B4%AF%E7%A9%8D%E5%92%8C%E3%80%81%E3%81%97%E3%82%83%E3%81%8F%E3%81%A8%E3%82%8A%E6%B3%95%E3%80%91%E5%88%9D%E7%B4%9A%E8%80%85%E3%81%A7%E3%82%82%E8%A7%A3%E3%82%8B%E3%82%A2%E3%83%AB%E3%82%B4)
- [www.kumilog.net - Two Pointer Technique](https://www.kumilog.net/entry/two-pointers)
- [scrapbox.io - Two Pointer Technique](https://scrapbox.io/pocala-kyopro/%E3%81%97%E3%82%83%E3%81%8F%E3%81%A8%E3%82%8A%E6%B3%95)
