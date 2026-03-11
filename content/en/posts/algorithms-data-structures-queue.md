---
title: Algorithms and Data Structures - Queue
slug: algorithms-data-structures-queue
date: 2019-11-17T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Queue
translation_key: algorithms-data-structures-queue
---



# Overview
Referencing the [Algorithm Encyclopedia](https://www.shoeisha.co.jp/book/detail/9784798149776), we learn about algorithms and data structures.

The implementation is also available on [github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master).

# Queue
- A structure where data is lined up in a row so that only the data added first can be accessed.
  - The direction of addition and removal is opposite to that of a stack.
  - FIFO (First In First Out)
    - First come, first served
- Also known as a waiting line.
- Adding data is called enqueue, and removing data is called dequeue.

# Computational Time
Depends on the implementation form such as arrays or linked lists.

# Implementation
```golang
package main

// Queue is a queue.
type Queue struct {
	nodes []*Node
}

// Node is a item of a stack.
type Node struct {
	value string
}

// newQueue create a Stack.
func newQueue() *Queue {
	return &Queue{}
}

// enqueue adds an node to the end of the queue.
func (s *Queue) enqueue(n *Node) {
	s.nodes = append(s.nodes, n)
}

// dequeue removes an node from the top of the queue.
func (s *Queue) dequeue() {
	s.nodes = s.nodes[1:len(s.nodes)]
}
```

- The implementation is almost the same as a stack. Only the slice index access is different.
- Note
  - ![Image from iOS](/assets/images/posts/algorithms-data-structures-queue/69003840-2b7d0700-094c-11ea-996e-c116235a1dbe.jpg)

# References
- [flaviocopes.com - Go Data Structures: Queue](https://flaviocopes.com/golang-data-structure-queue/)
- [Wikipedia - Queue](https://ja.wikipedia.org/wiki/%E3%82%AD%E3%83%A5%E3%83%BC_(%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF))
