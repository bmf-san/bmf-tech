---
title: Algorithms and Data Structures - Arrays
slug: algorithms-data-structures-array
date: 2019-10-31T00:00:00Z
author: bmf-san
categories:
  - Algorithms
  - Data Structures
tags:
  - Arrays
translation_key: algorithms-data-structures-array
---

# Overview
Referencing the [Algorithm Encyclopedia](https://www.shoeisha.co.jp/book/detail/9784798149776), we will learn about algorithms and data structures.

The implementation is also available at [github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master).

# Arrays
- A structure that arranges data in a single line.
- Accessing data is easy, but adding or deleting takes time.
- The data in an array is stored sequentially in a contiguous area of memory.
  - Fixed-length memory allocation.
    - Allocated at declaration (static allocation).
    - Allocated at runtime (dynamic allocation).

# Time Complexity
Let n be the number of data stored in the array.

## Accessing Data
- O(1)
  - Since the memory address can be calculated using the index, data can be accessed directly. (Random access)

## Adding Data
- O(n)
  - All data after the insertion point must be shifted one by one.

## Deleting Data
- Similar to adding data.

# Implementation
```golang
package main

import (
	"errors"
	"fmt"
)

// A Array is array implemented by slice.
type Array struct {
	data   []string
	length int // Keep a array memory size
}

// Insert is insert a data to array.
func (a *Array) insert(index int, value string) error {
	if a.length == int(cap(a.data)) {
		return errors.New("a array is full")
	}

	if index != a.length && index >= a.length {
		return errors.New("out of index range")
	}

	// shift data
	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}

	// insert a value to target index
a.data[index] = value

	// update the length
a.length++

	return nil
}

// delete is delete a target data by index.
func (a *Array) delete(index int) (string, error) {
	if index >= a.length {
		return "", errors.New("out of index range")
	}

	// target value for deleting
	v := a.data[index]

	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}

	// unset
a.data[a.length-1] = ""

	// update the length
a.length--

	return v, nil
}

// get is get a target data by index.
func (a *Array) get(index int) (string, error) {
	if index >= a.length {
		return "", errors.New("out of index range")
	}

	// random access
	return a.data[index], nil
}

func main() {
	a := &Array{
		data:   make([]string, 10, 10),
		length: 0,
	}

	cases := []struct {
		index int
		value string
	}{
		{
			index: 0,
			value: "foo",
		},
		{
			index: 1,
			value: "bar",
		},
		{
			index: 2,
			value: "foobar",
		},
	}

	for _, c := range cases {
		if err := a.insert(c.index, c.value); err != nil {
			fmt.Printf("index: %v value: %v is error. %v\n", c.index, c.value, err)
		}
	}

	if s, err := a.delete(2); err != nil {
		fmt.Printf("index: 0 is error. %v\n", err)
	} else {
		fmt.Printf("%v is deleted.", s)
	}

	if r, err := a.get(0); err != nil {
		fmt.Printf("index: 0 is error. %v", err)
	} else {
		fmt.Printf("%v", r)
	}
}
```
- The struct Array defines the data structure of the array.
  - Although Go has arrays, we implement the array data structure using slices here.
  - Since the array has a fixed length, we prepare a length variable.
- insert
  - Conditional branches:
    - When the array is full (i.e., the number of data equals length).
    - When out-of-bounds access occurs.
  - Shift data:
    - Loop to shift data from the end of the array to the specified index.
    - After shifting, add data to the array and update the length.
- delete
  - Conceptually similar to the reverse of insert.
  - Conditional branches:
    - Only consider out-of-bounds access.
    - Since data is being reduced, whether the array is full is not a condition.
  - Shift data:
    - Unset the data specified by the index.
    - Loop to shift data from the specified index to the end.
    - After shifting, update the length.
- get
  - Consider only out-of-bounds access for random access (referencing data by index).
- Note
![Image from iOS](https://user-images.githubusercontent.com/13291041/67912042-f3e03200-fbcb-11e9-8a42-34f28fd474f4.jpg)

# References
- [github - TomorrowWu/golang-algorithms](https://github.com/TomorrowWu/golang-algorithms/blob/master/data-structures/array/array.go)
  - I referred to this easy-to-understand code that implemented an array using slices.

# Related
- [bmf-tech.com - Big O Notation and How to Calculate Algorithm Complexity](https://bmf-tech.com/posts/O（オーダー）記法とアルゴリズムの計算量の求め方)