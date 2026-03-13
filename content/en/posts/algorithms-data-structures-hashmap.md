---
title: Algorithms and Data Structures - HashMap
description: 'Understand how HashMaps work: O(1) average access, open addressing vs. chaining for hash collision handling, and a basic Go HashMap implementation.'
slug: algorithms-data-structures-hashmap
date: 2023-07-31T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - HashMap
translation_key: algorithms-data-structures-hashmap
---

# Overview
Referencing [Algorithm Encyclopedia](https://www.shoeisha.co.jp/book/detail/9784798149776), we learn about algorithms and data structures.

The implementation is also available on [github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master).

# HashMap
- An array indexed by hash values
- Hash collision handling
  - Open Addressing
    - A method to find another address using a different function when a collision occurs.
  - Chaining
    - A method where, instead of finding a new address when a collision occurs, a linked list is stored at the collided address, connecting the collided keys with pointers.

# Computational Time
## Data Access
- O(1)
  - Random access is possible using indices.

## Data Addition
- O(1)
  - In the case of arrays, a linear search is needed to find the addition location, making it O(n), but hash tables determine the addition location using a hash, so it is O(1).
  - This does not apply if a hash collision occurs.

# Implementation
Below is a rudimentary HashMap that does not consider hash collisions.

```golang
package main

import "fmt"

// A HashMap is hash map.
type HashMap struct {
	data map[int]string
}

// hash is create a hash key.
func hash(key int) int {
	return key % 5
}

// put is add key to hash map.
func (h HashMap) put(key int, value string) {
	hash := hash(key)
	if h.data == nil {
		h.data = make(map[int]string)
	}
	h.data[hash] = value
}

// get is get a value from hash map.
func (h HashMap) get(key int) string {
	var hash int = hash(key)

	return h.data[hash]
}

func main() {
	h := &HashMap{
		data: make(map[int]string),
	}

	h.put(1, "foo")
	h.put(2, "bar")

	fmt.Printf("%#v\n", h.get(1))
	fmt.Printf("%#v\n", h.get(2))
}
```

- There are various algorithms for hash functions.
  - cf. [wikipedia - Hash Function](https://ja.wikipedia.org/wiki/%E3%83%8F%E3%83%83%E3%82%B7%E3%83%A5%E9%96%A2%E6%95%B0#:~:targetText=%E3%83%8F%E3%83%83%E3%82%B7%E3%83%A5%E9%96%A2%E6%95%B0%20(%E3%83%8F%E3%83%83%E3%82%B7%E3%83%A5%E3%81%8B%E3%82%93%E3%81%99%E3%81%86,%E5%80%A4%E3%81%BE%E3%81%9F%E3%81%AF%E5%8D%98%E3%81%AB%E3%83%8F%E3%83%83%E3%82%B7%E3%83%A5%E3%81%A8%E3%81%84%E3%81%86%E3%80%82)

# References
- [Wikipedia - Hash Table](https://ja.wikipedia.org/wiki/%E3%83%8F%E3%83%83%E3%82%B7%E3%83%A5%E3%83%86%E3%83%BC%E3%83%96%E3%83%AB)
- [github.com - flaviocopes/datastructures/blob/master/hashtable/hashtable.go](https://github.com/flaviocopes/datastructures/blob/master/hashtable/hashtable.go)
