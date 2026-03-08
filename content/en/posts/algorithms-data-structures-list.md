---
title: Algorithms and Data Structures - Lists
slug: algorithms-data-structures-list
image: /assets/images/posts/post-201/67028620-a0d79b00-f146-11e9-8d47-a5d0d933d15d.jpg
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
Referencing the [Algorithm Encyclopedia](https://www.shoeisha.co.jp/book/detail/9784798149776), we learn about algorithms and data structures.

The implementation is also available at [github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master).

# List (Singly Linked List)
- A structure that arranges data in a straight line
  - Each node has a pointer to the next node
- Adding and removing data is easy, but accessing it takes time
- In a list, data does not need to be stored in contiguous memory areas
  - Generally, it is stored in separate areas

# Time Complexity
Let n be the number of data stored in the list.

## Accessing Data
- O(n)
  - It requires sequential access from the head of the data, resulting in linear time.

## Adding Data
- O(1)
  - Assuming access to the insertion point is complete, it only requires swapping two pointers, so it takes constant time.

## Deleting Data
- Similar to adding data.

# Implementation
## Singly Linked List
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
  - This is to perform sequential access for data in the list.
- add
  - A method to add a node to the tail of the list.
- insert
  - A method to add a node before a specified node in the list.
    - Identify the insertion position → Adjust the pointer of the node before the insertion position and the pointer of the new node.
    - When identifying the insertion position, check the value of the next node in the loop to see if it matches the specified value.
- Note
![singly_linked_list](https://user-images.githubusercontent.com/13291041/67028620-a0d79b00-f146-11e9-8d47-a5d0d933d15d.jpg)

# References
- [Naim Ibrahim - Golang singly linked list](https://www.naimibrahim.me/2019/06/08/golang-singly-linked-list/)
    - The implementation was easy to understand.

# Related
- [bmf-tech.com - Big O Notation and Algorithm Complexity Calculation](https://bmf-tech.com/posts/O（オーダー）記法とアルゴリズムの計算量の求め方)