---
title: Runner Technique for Linked Lists
slug: linked-list-runner-technique
date: 2023-07-22T00:00:00Z
author: bmf-san
categories:
  - Algorithms
  - Data Structures
tags:
  - Algorithm
  - Data Structure
  - Runner Technique
  - Tips
translation_key: linked-list-runner-technique
---

This post summarizes the runner technique useful for traversing linked lists.

I first learned about it in the book [Improve Your Programming Skills to Compete Globally ~ 189 Coding Interview Questions and Their Solutions](https://amzn.to/3q35TCw).

# What is the Runner Technique?
The runner technique involves using two pointers: one that traverses from the head of the linked list and another that traverses from a position ahead of the first pointer, allowing for simultaneous traversal.

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

	// Set n1 to the nth node
	for i := 0; i < n; i++ {
		if n1 == nil {
			return nil
		}
		n1 = n1.next
	}

	// n2 traverses from the head, while n1 traverses from the nth node.
	// When n1 reaches the end node, n2 will be the nth node from the end.
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

If the number of nodes in the linked list is known, you could simply solve it without needing two pointers, as (total nodes - n) would give you the nth node from the end. However, if that is not the case, you would either solve it this way or use recursion, which would increase the computational complexity...

# Thoughts
I think this technique could be useful in coding quizzes, so I want to keep it in mind.