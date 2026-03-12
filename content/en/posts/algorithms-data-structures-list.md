---
title: Algorithms and Data Structures - Lists
description: An in-depth look at Algorithms and Data Structures - Lists, covering key concepts and practical insights.
slug: algorithms-data-structures-list
date: 2019-10-18T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Linked List
  - Singly Linked List
translation_key: algorithms-data-structures-list
---



# Overview
Learn algorithms and data structures with reference to [Algorithm Encyclopedia](https://www.shoeisha.co.jp/book/detail/9784798149776).

The implementation is also available on [github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master).

# List (Singly Linked List)
- A structure where data is arranged in a straight line
  - Each node has a pointer to the next node
- Adding and deleting data is easy, but accessing data takes time
- In a list, data does not need to be stored in contiguous memory areas
  - Generally stored in separate areas

# Computational Time
Let n be the number of data stored in the list.

## Accessing Data
- O(n)
  - Since sequential access from the beginning of the data is necessary, it takes linear time

## Adding Data
- O(1)
  - Assuming data access to the addition point is complete, it only requires swapping two pointers, so it takes constant time

## Deleting Data
- Same as adding data

# Implementation
## Linear List
### Singly Linked List

```golang
package main

import (
	"errors"
	"fmt"
)

// A node is a node of list.
type node struct {
	value string
	next  *node
}

// A list is a singly linked list.
type list struct {
	head *node
}

// add add a node to tail of a list.
func (l *list) add(newn *node) {
	if l.head == nil {
		l.head = newn
		newn.next = nil
		return
	}

	// sequential access
	for n := l.head; n != nil; n = n.next {
		if n.next == nil {
			n.next = newn
			return
		}
	}

	return
}

// insert a node before a particular node of a list.
func (l *list) insert(newn *node, v string) error {
	if l.head == nil {
		return errors.New("a target node is not exists")
	}

	// sequential access
	for n := l.head; n.next != nil; n = n.next {
		if n.next.value == v {
			newn.next = n.next
			n.next = newn
			return nil
		}
	}

	return errors.New("a target node is not exists")
}

// display display all nodes of a list.
func (l *list) display() {
	// sequential access
	for n := l.head; n != nil; n = n.next {
		fmt.Println(n.value, n.next)
	}
}

func main() {
	l := &list{}

	first := &node{"first", nil}
	second := &node{"second", nil}
	third := &node{"third", nil}

	l.add(first)
	l.add(second)
	l.add(third)

	between := &node{"between", nil}
	l.insert(between, "second")

	l.display()

	fmt.Printf("%#v\n", l)
}
```

- The List struct stores the head node of the list
  - Because data access in the list is done sequentially
- add
  - A method to add a node to the end of the list
- insert
  - A method to add a node before a specified node in the list
    - Identify the position to add → adjust the pointer of the node before the addition position and the pointer of the node to be added
    - When identifying the position to add, refer to the pointer of the next node in the loop and determine if the value of the next node matches the specified value
- Note
![singly_linked_list](/assets/images/posts/algorithms-data-structures-list/67028620-a0d79b00-f146-11e9-8d47-a5d0d933d15d.jpg)

# References
- [Naim Ibrahim - Golang singly linked list](https://www.naimibrahim.me/2019/06/08/golang-singly-linked-list/) 
    - The implementation was easy to understand

# Related
- [bmf-tech.com - O Notation and How to Determine Algorithm Complexity](https://bmf-tech.com/posts/O（オーダー）記法とアルゴリズムの計算量の求め方)
