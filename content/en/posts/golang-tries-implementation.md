---
title: Implementing a Trie in Golang
slug: golang-tries-implementation
date: 2019-09-24T00:00:00Z
author: bmf-san
categories:
  - Algorithms
  - Data Structures
tags:
  - Golang
  - Radix Tree
  - Trie
translation_key: golang-tries-implementation
---

# Overview
This post discusses the algorithm and implementation of a Trie.

[bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master)  

# What is a Trie?
A Trie (also known as a prefix tree) is a type of tree structure that handles a collection of strings.

Each node can hold one or more strings or values (nodes do not necessarily need to hold a value), and by traversing from the root node to the leaves, it represents words by connecting values.

Applications of Tries can be seen in network layers for IP address lookups, in application layers for HTTP routing, and in contexts like machine learning for morphological analysis.

Visuals are often easier to understand than verbal explanations.

[Algorithm Visualizations - Trie (Prefix Tree)](https://www.cs.usfca.edu/~galles/visualization/Trie.html)

Since Tries are not memory efficient, it may be better to consider data structures like Radix Trees (Patricia Tries) when memory efficiency is a bottleneck.

# Time Complexity
When the length of the search key is m, the worst-case time complexity is O(m). This applies to both search and insertion operations.

# Implementation

The implementation can be found at [github.com/bmf-san/road-to-algorithm-master/tree/master/data_structures/tree/trie](https://github.com/bmf-san/road-to-algorithm-master/tree/master/data_structures/tree/trie).

I have implemented only the insertion and search of the search key.

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

The children of the Node struct use rune as the key in the map, but it seems like using string would also work fine.

The algorithm for Insert is simple; it loops through the number of characters in the string to be inserted, checks if there is a matching key in the child nodes, and adds a node if there isn't.

In Golang, you can use the idiom `v, ok:=map[key]` to check for matching keys in child nodes. (I got stuck on this because I was too much of a beginner to know it.)

The algorithm for Search can be understood if you can write the Insert algorithm. In fact, the thought process is almost the same.

# Thoughts
I was working on implementing a Radix Tree, which is an application of Tries, but I stumbled multiple times just when I was about to finish and ended up giving up...
I will take a detour and try again later.

I am trying to create routing using Trie in Golang.
[https://github.com/bmf-san/bmf-go-router](https://github.com/bmf-san/bmf-go-router)

I feel that Tries can also be used for implementing simple suggestion features, so I would like to try implementing it in JavaScript.