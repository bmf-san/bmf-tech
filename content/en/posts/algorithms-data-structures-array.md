---
title: Algorithms and Data Structures - Arrays
slug: algorithms-data-structures-array
date: 2019-10-31T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Arrays
translation_key: algorithms-data-structures-array
---

# Overview
Learning algorithms and data structures with reference to [Algorithm Encyclopedia](https://www.shoeisha.co.jp/book/detail/9784798149776).

The implementation is also available on [github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master).

# Arrays
- Data arranged in a single row
- Easy access to data, but adding or deleting takes time
- Array data is stored sequentially in contiguous memory regions
  - Fixed-length memory allocation
    - Allocated at declaration (static allocation)
    - Allocated at runtime (dynamic allocation)

# Computational Time
Let n be the number of data stored in the array.

## Accessing Data
- O(1)
  - Direct access to data is possible because the memory address can be calculated using the index (random access).

## Adding Data
- O(n)
  - All data after the insertion point needs to be shifted one by one.

## Deleting Data
- Same as adding data

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

- The `Array` struct defines the data structure of an array.
  - Although Go has arrays, here we implement the array data structure using slices.
  - Arrays are fixed-length, so we prepare a length.
- insert
  - Conditional branching
    - When the array is full (i.e., the number of data equals the length)
    - When an out-of-range access occurs
  - Shifting data
    - Loop to shift data from the end of the array to the desired index
    - After shifting, add data to the array and update the length
- delete
  - Conceptually similar to the reverse of insert
  - Conditional branching
    - Only need to consider out-of-range access
    - Since data is reduced, whether the array is full is not a condition
  - Shifting data
    - Unset the data specified by the index
    - Loop to shift data from the desired index to the end
    - After shifting, update the length
- get
  - Only consider out-of-range access and perform random access (referencing data by index)
- Note
![Image from iOS](https://user-images.githubusercontent.com/13291041/67912042-f3e03200-fbcb-11e9-8a42-34f28fd474f4.jpg)

# References
- [github - TomorrowWu/golang-algorithms](https://github.com/TomorrowWu/golang-algorithms/blob/master/data-structures/array/array.go)
  - Referenced as it was an easy-to-understand code implementing array with slice.

# Related
- [bmf-tech.com - O Notation and How to Calculate Algorithm Complexity](https://bmf-tech.com/posts/O（オーダー）記法とアルゴリズムの計算量の求め方)
