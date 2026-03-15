---
title: Implementing a Trie in Golang
slug: golang-tries-implementation
date: 2019-09-24T00:00:00Z
lastmod: 2026-03-15
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Golang
  - Radix Tree
  - Trie
description: 'Learn how Trie (prefix tree) data structures work: O(m) search and insertion, applications in HTTP routing and IP lookups, and a Go implementation of insert and search operations.'
translation_key: golang-tries-implementation
---



# Overview
This post discusses the algorithm and implementation of a Trie.

[bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master)

# What is a Trie?
A Trie (also known as a prefix tree) is a type of tree structure that handles a collection of strings.

Each node can hold one or more strings or numbers (nodes do not necessarily need to hold values), and words are represented by connecting values as you traverse from the root node to the leaves.

In network layers, Tries can be applied to IP address lookups, in application layers to HTTP routing, and in machine learning contexts to morphological analysis.

Visuals are often easier to understand than verbal explanations.

[Algorithm Visualizations - Trie (Prefix Tree)](https://www.cs.usfca.edu/~galles/visualization/Trie.html)

Due to poor memory efficiency, if memory efficiency becomes a bottleneck, it might be better to consider tree structures like Radix Trees (Patricia Tries) that efficiently store string prefixes.

# Computational Complexity
When the length of the search key is m, the worst-case complexity is O(m). This applies to both search and insertion.

# Implementation

The implementation is available at [github.com/bmf-san/road-to-algorithm-master/tree/master/data_structures/tree/trie](https://github.com/bmf-san/road-to-algorithm-master/tree/master/data_structures/tree/trie).

Only insertion and search of search keys are implemented.

```go
package main

import "fmt"

// Node is a node of tree.
type Node struct {
	key      string
	children map[rune]*Node
}

// NewTrie is create a root node.
func NewTrie() *Node {
	return &Node{
		key:      "",
		children: make(map[rune]*Node),
	}
}

// Insert is insert a word to tree.
func (n *Node) Insert(word string) {
	runes := []rune(word)
	curNode := n

	for _, r := range runes {
		if nextNode, ok := curNode.children[r]; ok {
			curNode = nextNode
		} else {
			curNode.children[r] = &Node{
				key:      string(r),
				children: make(map[rune]*Node),
			}
		}
	}
}

// Search is search a word from a tree.
func (n *Node) Search(word string) bool {
	if len(n.key) == 0 && len(n.children) == 0 {
		return false
	}

	runes := []rune(word)
	curNode := n

	for _, r := range runes {
		if nextNode, ok := curNode.children[r]; ok {
			curNode = nextNode
		} else {
			return false
		}
	}

	return true
}

func main() {
	t := NewTrie()

	t.Insert("word")
	t.Insert("wheel")
	t.Insert("world")
	t.Insert("hospital")
	t.Insert("mode")

	fmt.Printf("%v", t.Search("mo")) // true
}
```

The `children` of the `Node` struct uses `rune` as the map key, but it seems like using `string` would also work.

The insertion algorithm is simple: loop through the characters of the string you want to insert, check if there is a matching key in the child nodes, and if not, add a node.

In Golang, you can check if there is a matching key in the child nodes using the idiom `v, ok := map[key]`. (I got a bit stuck because I was too much of a beginner to know this.)

The search algorithm can be understood if you can write the insertion algorithm. In fact, the concept is almost the same.

# Thoughts
I was working on implementing a tree structure called a Radix Tree, which is an application of a Trie, but I kept stumbling at the last step and gave up... I'll take a detour and try again.

I'm trying to create routing using a Trie in Golang.
[https://github.com/bmf-san/bmf-go-router](https://github.com/bmf-san/bmf-go-router)

I feel like Tries could also be used for implementing simple suggestion features, so I'd like to try implementing one in JavaScript.

## Related Posts

- [Order Notation and How to Determine Algorithm Complexity](/posts/big-o-notation-algorithms/)
- [Algorithms and Data Structures - Binary Search Tree](/posts/algorithms-data-structures-binary-search-tree/)
- [Reviewing the Basics of Algorithms and Data Structures](/posts/algorithms-data-structures-review/)
- [From Custom HTTP Router to New ServeMux](/posts/custom-http-router-to-new-servemux/)