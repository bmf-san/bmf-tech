---
title: Search Patterns in Binary Search Trees
description: An in-depth exploration of Search Patterns in Binary Search Trees, covering design principles, trade-offs, and practical applications.
slug: binary-search-tree-search-patterns
date: 2023-08-04T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Binary Search Tree
  - DFS
  - BFS
translation_key: binary-search-tree-search-patterns
---

# Overview

# What is a Binary Search Tree
A tree where for any node, the left child node < parent node < right child node.

```sh
ex. 
    5
   / \
  3   8
 / \ / \
1  4 6  9
```

# Search Patterns
## Depth First Search (DFS)
The order of node traversal is easy to remember using the method described at [www.momoyama-usagi.com - Understanding Binary Search Trees for Beginners Part 2: Four Traversal Methods](https://www.momoyama-usagi.com/entry/info-algo-tree-traverse), so I have included the link.

### Preorder
Start traversal from the **root node**, recursively traversing the left subtree, then the right subtree.

Remember it as **root is pre (before), so root → left → right**. (Just remember the position of the root. Left always comes before right.)

cf.
[(2) The Magic of Correctly Determining Preorder](https://www.momoyama-usagi.com/entry/info-algo-tree-traverse#2)

When you trace the tree with a single stroke, the order in which you pass the left side of the nodes is preorder.

```sh
      5
     / \
    3   8
   / \ / \
  1  4 6  9

5 -> 3 -> 1 -> 4 -> 8 -> 6 -> 9
```

### Inorder
Recursively traverse from the **left subtree**, then the root node, and finally the right subtree.

Remember it as **root is in (between), so left → root → right**. (Just remember the position of the root. Left always comes before right.)

cf.
[(2) The Magic of Correctly Determining Inorder](https://www.momoyama-usagi.com/entry/info-algo-tree-traverse#2-3)

When you trace the tree with a single stroke, the order in which you pass the bottom side of the nodes is inorder.

```sh
    5
   / \
  3   8
 / \ / \
1  4 6  9

1 -> 3 -> 4 -> 5 -> 6 -> 8 -> 9
```

### Postorder
Recursively traverse from the **left subtree**, then the right subtree, and finally the root node.

Remember it as **root is post (after), so left → right → root**. (Just remember the position of the root. Left always comes before right.)

cf. [(2) The Magic of Correctly Determining Postorder](https://www.momoyama-usagi.com/entry/info-algo-tree-traverse#2-4)

As an aside, Reverse Polish Notation can be solved with a stack, but it can also be solved with postorder traversal of a binary tree.

When you trace the tree with a single stroke, the order in which you pass the right side of the nodes is postorder.

```sh
    5
   / \
  3   8
 / \ / \
1  4 6  9

1 -> 4 -> 3 -> 6 -> 9 -> 8 -> 5
```

## Breadth First Search (BFS)
Traverse by depth.

```sh
    5
   / \
  3   8
 / \ / \
1  4 6  9

5 -> 3 -> 8 -> 1 -> 4 -> 6 -> 9
```

# Implementation
The source code is available at [binary_search_tree](https://github.com/bmf-san/road-to-algorithm-master/tree/master/data_structures/tree/binary_search_tree).

```go
package main

import "fmt"

// a binary search tree.
type tree struct {
	root *node
}

// a node for binary search tree.
type node struct {
	val string
	l   *node
	r   *node
}

// insert a value to tree.
func (t *tree) insert(v string) {
	t.root = t.root.insertNode(v)
}

// insert a node to tree.
func (n *node) insertNode(v string) *node {
	if n == nil {
		return &node{val: v}
	}
	if v < n.val {
		n.l = n.l.insertNode(v)
	} else if v > n.val {
		n.r = n.r.insertNode(v)
	}
	return n
}

// search a value from tree.
func (t *tree) search(v string) bool {
	return t.root.searchNode(v)
}

// search a node from tree.
func (n *node) searchNode(v string) bool {
	if n == nil {
		return false
	}
	if v == n.val {
		return true
	} else if v < n.val {
		return n.l.searchNode(v)
	} else {
		return n.r.searchNode(v)
	}
}

// remove a value from tree.
func (t *tree) remove(v string) {
	t.root = t.root.removeNode(v)
}

// remove a node from tree.
func (n *node) removeNode(v string) *node {
	if n == nil {
		return nil
	}

	if v < n.val {
		n.l = n.l.removeNode(v)
		return n
	} else if v > n.val {
		n.r = n.r.removeNode(v)
		return n
	} else {
		// node has no children
		if n.l == nil && n.r == nil {
			return nil
		}

		// node has only right child
		if n.l == nil {
			return n.r
		}

		// node has only left child
		if n.r == nil {
			return n.l
		}

		// node has both children
		leftmostrightside := n.r
		for leftmostrightside.l != nil {
			leftmostrightside = leftmostrightside.l
		}
		n.val = leftmostrightside.val
		n.r = n.r.removeNode(n.val)
		return n
	}
}

// breadth first search - preorder
//
//    5
//   / \
//  3   8
// / \ / \
//1  4 6  9
//
//5 -> 3 -> 1 -> 4 -> 8 -> 6 -> 9
func (t *tree) preorder(n *node, f func(string)) {
	if n != nil {
		// root → left → right
		f(n.val)
		t.preorder(n.l, f)
		t.preorder(n.r, f)
	}
}

// breadth first search - inorder
//
//    5
//   / \
//  3   8
// / \ / \
//1  4 6  9
//
//1 -> 3 -> 4 -> 5 -> 6 -> 8 -> 9
func (t *tree) inorder(n *node, f func(string)) {
	if n != nil {
		// left → root → right
		t.inorder(n.l, f)
		f(n.val)
		t.inorder(n.r, f)
	}
}

// breadth first search - postorder
//
//    5
//   / \
//  3   8
// / \ / \
//1  4 6  9
//
//1 -> 4 -> 3 -> 6 -> 9 -> 8 -> 5
func (t *tree) postorder(n *node, f func(string)) {
	if n != nil {
		// left → right → root
		t.postorder(n.l, f)
		t.postorder(n.r, f)
		f(n.val)
	}
}

// depth first search
//
//    5
//   / \
//  3   8
// / \ / \
//1  4 6  9
//
//5 -> 3 -> 8 -> 1 -> 4 -> 6 -> 9
func (t *tree) dfs(n *node, f func(string)) {
	if n != nil {
		s := []*node{n}
		for len(s) > 0 {
			crtn := s[0]
			f(crtn.val)
			s = s[1:]
			if crtn.l != nil {
				s = append(s, crtn.l)
			}
			if crtn.r != nil {
				s = append(s, crtn.r)
			}
		}
	}
}

func main() {
	t := &tree{}
	t.insert("5")
	t.insert("3")
	t.insert("8")
	t.insert("1")
	t.insert("4")
	t.insert("6")
	t.insert("9")
	t.insert("11")
	fmt.Println(t.search("11")) // true
	t.remove("11")
	fmt.Println(t.search("11")) // false
	f := func(v string) {
		fmt.Println(v)
	}
	t.preorder(t.root, f) // 5314869
	fmt.Println("-----")
	t.inorder(t.root, f) // 1345689
	fmt.Println("-----")
	t.postorder(t.root, f) // 1436985
	fmt.Println("-----")
	t.dfs(t.root, f) // 5381469
}
```

Breadth-first search only requires remembering the starting position of the root (preorder is root → left → right), and recursive processing can be written accordingly.

Depth-first search is a bit more troublesome. Moreover, node deletion processing is cumbersome not only in binary search trees...

# Impressions
If you remember the search patterns this way, it seems like you can keep them in mind.

# References
- [www.momoyama-usagi.com - Understanding Binary Search Trees for Beginners Part 2: Four Traversal Methods](https://www.momoyama-usagi.com/entry/info-algo-tree-traverse)
  - The traversal methods are very easy to remember, and I used them as a reference.