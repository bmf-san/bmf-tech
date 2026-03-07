---
title: Algorithms and Data Structures - Stack
slug: algorithms-data-structures-stack
date: 2019-11-17T00:00:00Z
author: bmf-san
categories:
  - Algorithms
  - Data Structures
tags:
  - Stack
translation_key: algorithms-data-structures-stack
---

# Overview
Referencing the [Algorithm Encyclopedia](https://www.shoeisha.co.jp/book/detail/9784798149776), we learn about algorithms and data structures.

The implementation is also available at [github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master).

# Stack
- A structure where data is arranged in a line so that only the most recent data can be accessed.
  - LIFO (Last In First Out)
    - Last in, first out
- A convenient structure when you want to access the most recent data.
- Adding data is called Push, and removing data is called Pop.
  - Other operations include Dup, Peek, Swap (or Exchange), and Rotate.
    - cf. [Wikipedia - Stack](https://ja.wikipedia.org/wiki/%E3%82%B9%E3%82%BF%E3%83%83%E3%82%AF)

# Time Complexity
Depends on the implementation form, such as arrays or linked lists.

# Implementation
```golang
package main

// Stack is a stack.
type Stack struct {
	nodes []*Node
}

// Node is a item of a stack.
type Node struct {
	value string
}

// newStack create a Stack.
func newStack() *Stack {
	return &Stack{}
}

// push adds an node to the top of the stack.
func (s *Stack) push(n *Node) {
	s.nodes = append(s.nodes[:len(s.nodes)], n)
}

// pop removes an node from the top of the stack.
func (s *Stack) pop() {
	s.nodes = s.nodes[:len(s.nodes)-1]
}
```

- There shouldn't be any particularly difficult parts if you're familiar with Go slices.
- Note
  - ![Image from iOS (1)](https://user-images.githubusercontent.com/13291041/69003839-2ae47080-094c-11ea-8506-cb733abec36e.jpg)

# References
- [Golang program for implementation LIFO Stack and FIFO Queue](https://www.golangprograms.com/golang-program-for-implementation-lifo-stack-and-fifo-queue.html)