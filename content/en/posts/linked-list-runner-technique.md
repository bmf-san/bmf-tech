---
title: Runner Technique for Linked Lists
slug: linked-list-runner-technique
date: 2023-07-22T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Algorithm
  - Data Structure
  - Runner Technique
  - Tips
description: Summarizing the runner technique useful for traversing linked lists.
translation_key: linked-list-runner-technique
---

Summarizing the runner technique useful for traversing linked lists.

I first learned about it from [Cracking the Coding Interview: 189 Programming Questions and Solutions](https://amzn.to/3q35TCw).

# What is the Runner Technique?
This method involves using two pointers: one that traverses from the head of the linked list and another that starts ahead of the first pointer, traversing simultaneously.

This technique is useful for solving problems like the following example.

## Example Problem
Implement an algorithm to find the nth element from the end of a singly linked list.

```go
package main

import "fmt"

type node struct {
	val  string
	next *node
}

type list struct {
	head *node
}

// Find the nth node from the end
func (l *list) search(n int) *node {
	n1 := l.head
	n2 := l.head

       // Set n1 to the kth node
	for i := 0; i < n; i++ {
		if n1 == nil {
			return nil
		}
		n1 = n1.next
	}
        
　　　　　　　　// Traverse n2 from the head node, and n1 from the kth node.
        // When n1 reaches the end node, n2 is the nth node from the end.
	for n1 != nil {
		n1 = n1.next
		n2 = n2.next
	}

	return n2
}

func main() {
	l := &list{
		head: &node{
			val: "a",
			next: &node{
				val: "b",
				next: &node{
					val: "c",
					next: &node{
						val: "d",
					},
				},
			},
		},
	}
	fmt.Printf("%+v\n", l.search(1)) // d
	fmt.Printf("%+v\n", l.search(2)) // c
	fmt.Printf("%+v\n", l.search(3)) // b
	fmt.Printf("%+v\n", l.search(4)) // a
}
```

Using the runner technique, the time complexity is O(N) and the space complexity is O(1).

If the number of nodes in the linked list is known, you can solve it simply by noting that (total nodes - n) is the nth node from the end, but if not, you can solve it this way or use recursion. In the case of recursion, the complexity should increase...

# Thoughts
I want to keep this in mind as it might be useful in coding quizzes.