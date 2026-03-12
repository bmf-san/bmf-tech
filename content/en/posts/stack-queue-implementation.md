---
title: Implementation of Stack and Queue
description: A step-by-step guide on Implementation of Stack and Queue, with practical examples and configuration tips.
slug: stack-queue-implementation
date: 2023-07-25T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Stack
  - Queue
translation_key: stack-queue-implementation
---


I tried implementing stack and queue in Go.

I implemented both patterns using slices and linked lists.

Personally, I find the pattern using slices easier to implement.

The time complexity for stack's push, pop, and queue's enqueue, dequeue can be implemented as O(1), but some parts are lazily implemented as O(N).

# Stack
Source code: [stack](https://github.com/bmf-san/road-to-algorithm-master/tree/master/data_structures/stack)

## Linked List
```go
package main

import "fmt"

// LIFO stack by using linked list.
type stack struct {
	top *node
}

type node struct {
	val  int
	next *node
}

// Remove data from the top of the stack.
func (s *stack) pop() {
	// LIFO
	s.top = s.top.next
}

// Add the item to the top of the stack.
func (s *stack) push(item int) {
	// LIFO
	s.top = &node{
		val:  item,
		next: s.top,
	}
}

// Returns the top item from the stack.
func (s *stack) peek() *node {
	return s.top
}

// Returns true if the stack is empty.
func (s *stack) isEmpty() bool {
	return s.top == nil
}

func (s *stack) traverse() {
	crt := s.top
	for {
		if crt == nil {
			break
		}
		fmt.Printf("%#v\n", crt)
		crt = crt.next
	}
}

func main() {
	s := &stack{
		top: &node{
			val: 1,
			next: &node{
				val: 2,
				next: &node{
					val: 3,
				},
			},
		},
	}
	s.pop()
	s.traverse()
	fmt.Println("----") // 2 3
	s.pop()
	s.traverse() // 3
	fmt.Println("----")
	s.pop()
	s.traverse() // nil

	s.push(1)
	s.traverse() // 1
	fmt.Println("----")
	s.push(2)
	s.push(3)
	s.traverse() // 3 2 1

	fmt.Println("----")
	fmt.Printf("%#v\n", s.peek()) // 3

	s2 := &stack{}
	fmt.Println(s2.isEmpty()) // true
}
```

It seems like a simple linked list without any particularly difficult parts.

## Slice
```go
package main

import "fmt"

type stack struct {
	nodes []*node
}

type node struct {
	val int
}

// Remove data from the top of the stack.
func (s *stack) pop() {
	// LIFO
	s.nodes = s.nodes[1:]
}

// Add the item to the top of the stack.
func (s *stack) push(item int) {
	// LIFO
	s.nodes = append(
		[]*node{
			&node{
				val: item,
			},
		},
		s.nodes...,
	)
}

// Returns the top item from the stack.
func (s *stack) peek() *node {
	return s.nodes[0]
}

// Returns true if the stack is empty.
func (s *stack) isEmpty() bool {
	return len(s.nodes) == 0
}

func (s *stack) traverse() {
	for _, n := range s.nodes {
		fmt.Printf("%#v\n", n)
	}
}

func main() {
	s := &stack{
		nodes: []*node{
			&node{
				val: 1,
			},
			&node{
				val: 2,
			},
			&node{
				val: 3,
			},
		},
	}
	s.pop()
	s.traverse()
	fmt.Println("----") // 2 3
	s.pop()
	s.traverse() // 3
	fmt.Println("----")
	s.pop()
	s.traverse() // nil

	s.push(1)
	s.traverse() // 1
	fmt.Println("----")
	s.push(2)
	s.push(3)
	s.traverse() // 3 2 1

	fmt.Println("----")
	fmt.Printf("%#v\n", s.peek()) // 3

	s2 := &stack{}
	fmt.Println(s2.isEmpty()) // true
}
```

You can implement it with slice operations. It might take some getting used to the way of adding elements to the beginning of a slice.

# Queue
Source code: [queue](https://github.com/bmf-san/road-to-algorithm-master/tree/master/data_structures/queue)

## Linked List
```go
package main

import "fmt"

// FIFO queue by using linked.
type queue struct {
	top *node
}

type node struct {
	val  int
	next *node
}

// Add the item to the end of the queue.
func (q *queue) enqueue(item int) {
	// FIFO
	if q.top == nil {
		q.top = &node{
			val: item,
		}
		return
	}

	crt := q.top
	for {
		if crt.next == nil {
			crt.next = &node{
				val: item,
			}
			break
		}
		crt = crt.next
	}
}

// Remove item from the front of the queue.
func (q *queue) dequeue() {
	// FIFO
	q.top = q.top.next
}

// Returns the front item from the queue.
func (q *queue) peek() *node {
	return q.top
}

// Returns true if the queue is empty.
func (q *queue) isEmpty() bool {
	return q.top == nil
}

func (q *queue) traverse() {
	crt := q.top
	for {
		if crt == nil {
			break
		}
		fmt.Printf("%#v\n", crt)
		crt = crt.next
	}
}

func main() {
	q := &queue{
		top: &node{
			val: 1,
			next: &node{
				val: 2,
				next: &node{
					val: 3,
				},
			},
		},
	}
	q.dequeue()
	q.traverse()
	fmt.Println("----") // 2 3
	q.dequeue()
	q.traverse() // 3
	fmt.Println("----")
	q.dequeue()
	q.traverse() // nil

	q.enqueue(1)
	q.traverse() // 1
	fmt.Println("----")
	q.enqueue(2)
	q.enqueue(3)
	q.traverse() // 1 2 3

	fmt.Println("----")
	fmt.Printf("%#v\n", q.peek()) // 3

	q2 := &queue{}
	fmt.Println(q2.isEmpty()) // true
}
```

The enqueue operation depends on the length of the queue, making it O(N).

If you implement it by having the tail of the queue or the length of the queue in the data structure (queue), it becomes O(1), which is more desirable.

## Slice
```go
package main

import "fmt"

// FIFO queue by using slice.
type queue struct {
	nodes []*node
}

type node struct {
	val int
}

// Add the item to the end of the queue.
func (q *queue) enqueue(item int) {
	// FIFO
	q.nodes = append(q.nodes, &node{
		val: item,
	})
}

// Remove item from the front of the queue.
func (q *queue) dequeue() {
	// FIFO
	q.nodes = q.nodes[1:]
}

// Returns the front item from the queue.
func (q *queue) peek() *node {
	return q.nodes[0]
}

// Returns true if the queue is empty.
func (q *queue) isEmpty() bool {
	return len(q.nodes) == 0
}

func (q *queue) traverse() {
	for _, n := range q.nodes {
		fmt.Printf("%#v\n", n)
	}
}

func main() {
	q := &queue{
		nodes: []*node{
			&node{
				val: 1,
			},
			&node{
				val: 2,
			},
			&node{
				val: 3,
			},
		},
	}
	q.dequeue()
	q.traverse()
	fmt.Println("----") // 2 3
	q.dequeue()
	q.traverse() // 3
	fmt.Println("----")
	q.dequeue()
	q.traverse() // nil

	q.enqueue(1)
	q.traverse() // 1
	fmt.Println("----")
	q.enqueue(2)
	q.enqueue(3)
	q.traverse() // 1 2 3

	fmt.Println("----")
	fmt.Printf("%#v\n", q.peek()) // 3

	q2 := &queue{}
	fmt.Println(q2.isEmpty()) // true
}
```

Simple slice operations. If there's no need for a linked list, this might be better, but be cautious about the memory efficiency of slices (allocation and copying).

# Thoughts
Implementing both in parallel can sometimes confuse which is LIFO and which is FIFO, haha.
