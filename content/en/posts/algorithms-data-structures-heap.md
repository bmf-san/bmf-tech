---
title: Algorithms and Data Structures - Heap
slug: algorithms-data-structures-heap
date: 2020-01-14T00:00:00Z
author: bmf-san
categories:
  - Algorithms
  - Data Structures
tags:
  - Heap
translation_key: algorithms-data-structures-heap
---

# Overview
Referencing the [Algorithm Encyclopedia](https://www.shoeisha.co.jp/book/detail/9784798149776), we learn about algorithms and data structures.

The implementation is also available at [github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master).

# Heap
- A type of priority queue
  - A priority queue is a data type that handles a set
    - Elements in the set are retrieved in order of priority
    - Examples of data types that handle sets: queue, stack
- Types of heaps
  - Min Heap
    - A heap where the root is always the minimum. The parent node is always smaller than the child nodes.
  - Max Heap
    - A heap where the root is always the maximum element. The parent node is always larger than the child nodes.

# Time Complexity
- Both insertion and deletion are O(log n)

# Implementation
```golang
package main

import "fmt"

// Heap is a heap.
type Heap struct {
	values  []int
	size    int
	maxsize int
}

// newHeap creates a heap.
func newHeap(maxsize int) *Heap {
	return &Heap{
		values:  []int{},
		size:    0,
		maxsize: maxsize,
	}
}

// leaf checks whether index is a leaf.
func (h *Heap) leaf(index int) bool {
	if index >= (h.size/2) && index <= h.size {
		return true
	}
	return false
}

// parent checks whether index is a parent.
func (h *Heap) parent(index int) int {
	return (index - 1) / 2
}

// leftchild checks whether index is a leftchild.
func (h *Heap) leftchild(index int) int {
	return 2*index + 1
}

// rightchild checks whether index is a rightchild.
func (h *Heap) rightchild(index int) int {
	return 2*index + 2
}

// insert inserts a item to a heap.
func (h *Heap) insert(item int) error {
	if h.size >= h.maxsize {
		return fmt.Errorf("Error!")
	}
	h.values = append(h.values, item)
	h.size++
	h.upHeapify(h.size - 1)
	return nil
}

// swap swaps two values.
func (h *Heap) swap(first, second int) {
	temp := h.values[first]
	h.values[first] = h.values[second]
	h.values[second] = temp
}

// upHeapify reconstruct a heap for up.
func (h *Heap) upHeapify(index int) {
	for h.values[index] < h.values[h.parent(index)] {
		h.swap(index, h.parent(index))
	}
}

// downHeapify reconstruct a heap for down.
func (h *Heap) downHeapify(current int) {
	if h.leaf(current) {
		return
	}

	smallest := current
	leftChildIndex := h.leftchild(current)
	rightRightIndex := h.rightchild(current)

	if leftChildIndex < h.size && h.values[leftChildIndex] < h.values[smallest] {
		smallest = leftChildIndex
	}
	if rightRightIndex < h.size && h.values[rightRightIndex] < h.values[smallest] {
		smallest = rightRightIndex
	}
	if smallest != current {
		h.swap(current, smallest)
		h.downHeapify(smallest)
	}
	return
}

// buildMinHeap builds a min heap.
func (h *Heap) buildMinHeap() {
	for index := ((h.size / 2) - 1); index >= 0; index-- {
		h.downHeapify(index)
	}
}

// remove removes a value.
func (h *Heap) remove() int {
	top := h.values[0]
	h.values[0] = h.values[h.size-1]
	h.values = h.values[:(h.size)-1]
	h.size--
	h.downHeapify(0)
	return top
}

func main() {
	inputArray := []int{6, 5, 3, 7, 2, 8}
	h := newHeap(len(inputArray))
	for i := 0; i < len(inputArray); i++ {
		h.insert(inputArray[i])
	}
	h.buildMinHeap()
	for i := 0; i < len(inputArray); i++ {
		fmt.Println(h.remove())
	}
	fmt.Scanln()
}
```

- Implementation of Min Heap
- Nodes are always added in breadth-first order
- After adding a node, if the added value is smaller than the root node, the root node and the added node are swapped
- It is easy to understand by checking the image at [Data Structure Visualizations - Min Heap](https://www.cs.usfca.edu/~galles/visualization/Heap.html)
- The calculations for root, parent, left child, and right child are characteristic
  - Since the array index serves as the priority, the values of those nodes can be derived from the reference node through calculations
  - The article [Explaining Heaps Clearly](https://medium.com/@yasufumy/data-structure-heap-ecfd0989e5be) is easy to understand
- Referenced [Welcome To Golang By Example - Heap in Golang](https://golangbyexample.com/heap-in-golang/) (mostly transcribed..)

# References
- [Explaining Heaps Clearly](https://medium.com/@yasufumy/data-structure-heap-ecfd0989e5be)
- [Welcome To Golang By Example - Heap in Golang](https://golangbyexample.com/heap-in-golang/)