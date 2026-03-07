---
title: Algorithms and Data Structures - HashMap
slug: algorithms-data-structures-hashmap
date: 2023-07-31T00:00:00Z
author: bmf-san
categories:
  - Algorithms
  - Data Structures
tags:
  - HashMap
translation_key: algorithms-data-structures-hashmap
---

# Overview
Referencing the [Algorithm Encyclopedia](https://www.shoeisha.co.jp/book/detail/9784798149776), we learn about algorithms and data structures.

The implementation is also available at [github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master).

# HashMap
- An array indexed by hash values
- Hash collision handling
  - Open addressing
    - A method to find a different address using a function other than the hash function when a collision occurs.
  - Chaining
    - A method to handle collisions by storing a linked list of colliding keys at the colliding address without seeking a new address.

# Time Complexity
## Accessing Data
- O(1)
 - Random access is possible using indices.

## Adding Data
- O(1)
 - In the case of an array, it requires linear search to find the insertion point, resulting in O(n), but a hash table can determine the insertion point using the hash, resulting in O(1).
 - This does not apply if there is a hash collision.

# Implementation
Below is a rough HashMap that does not consider hash collisions.

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