---
title: 2分探索木の探索パターンについて
description: "2分探索木の探索パターンを解説。DFS（先行順・中間順・後行順）、BFS、一筆書き法による走査で木構造の走査順序をマスターする実践ガイドです。"
slug: binary-search-tree-search-patterns
date: 2023-08-04T00:00:00Z
author: bmf-san
categories:
  - アルゴリズムとデータ構造
tags:
  - 二分探索木
  - DFS
  - BFS
translation_key: binary-search-tree-search-patterns
---


# 概要

# 二分探索木とは
どのノードにおいても、左の子ノード＜親ノード＜右の子ノードとなるような木。

```sh
ex. 
    5
   / \
  3   8
 / \ / \
1  4 6  9
```

# 探索パターン
## 深さ優先探索（DFS: Depth first search）
それぞれのノード探索順は[www.momoyama-usagi.com - うさぎでもわかる2分探索木　後編　2分探索木における4つの走査方法](https://www.momoyama-usagi.com/entry/info-algo-tree-traverse)に記載されている方法が覚えやすいのでそちらのリンクを記載する。

### preorder(先行順/前順/行きがけ順)
**ルートノードから**走査を開始し、左部分木、右部分木の順で再帰的に走査する。

**rootがpre（事前）なので根→左→右**と覚えておく良い。（根の位置だけ覚えておけば良い。左、右は必ず左が先になる。）

cf.
[(2) 行きがけ順を正しく求める魔法の一筆書き](https://www.momoyama-usagi.com/entry/info-algo-tree-traverse#2)

木を一筆書きで括ったときに、ノードの左側を通った順がpreorder。


```sh
      5
     / \
    3   8
   / \ / \
  1  4 6  9

5 -> 3 -> 1 -> 4 -> 8 -> 6 -> 9
```

### inorder(中間順/間順/通りがけ順)
**左部分木から**再帰的に走査し、ルートノード、右部分木の順に走査する。

**rootがin（間）なので左→根→右**と覚えておく良い。（根の位置だけ覚えておけば良い。左、右は必ず左が先になる。）

cf.
[(2) 通りがけ順を正しく求める魔法の一筆書き](https://www.momoyama-usagi.com/entry/info-algo-tree-traverse#2-3)

木を一筆書きで括ったときに、ノードの下側を通った順がinorder。

```sh
    5
   / \
  3   8
 / \ / \
1  4 6  9

1 -> 3 -> 4 -> 5 -> 6 -> 8 -> 9
```

### postorder(後行順/後順/帰りがけ順)
**左部分木から**再帰的に走査し、右部分木、ルートノードの順で走査する。

**rootがpost（事後）なので左→右→根**と覚えておく良い。（根の位置だけ覚えておけば良い。左、右は必ず左が先になる。）

cf. [(2) 帰りがけ順を正しく求める魔法の一筆書き](https://www.momoyama-usagi.com/entry/info-algo-tree-traverse#2-4)

余談だが、逆ポーランド記法はスタックで解くこともできるが、二分木をpostorderでも解くことができる。

木を一筆書きで括ったときに、ノードの右側を通った順がpostorder。

```sh
    5
   / \
  3   8
 / \ / \
1  4 6  9

1 -> 4 -> 3 -> 6 -> 9 -> 8 -> 5
```

## 幅優先探索（BFS: Breadth first search）
深さごとに走査する。

```sh
    5
   / \
  3   8
 / \ / \
1  4 6  9

5 -> 3 -> 8 -> 1 -> 4 -> 6 -> 9
```

# 実装
ソースコードは[binary_search_tree](https://github.com/bmf-san/road-to-algorithm-master/tree/master/data_structures/tree/binary_search_tree)にある。

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
//	    5
//	   / \
//	  3   8
//	 / \ / \
//	1  4 6  9
//
//	5 -> 3 -> 1 -> 4 -> 8 -> 6 -> 9
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
//	    5
//	   / \
//	  3   8
//	 / \ / \
//	1  4 6  9
//
//	1 -> 3 -> 4 -> 5 -> 6 -> 8 -> 9
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
//	    5
//	   / \
//	  3   8
//	 / \ / \
//	1  4 6  9
//
//	1 -> 4 -> 3 -> 6 -> 9 -> 8 -> 5
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
//	    5
//	   / \
//	  3   8
//	 / \ / \
//	1  4 6  9
//
//	5 -> 3 -> 8 -> 1 -> 4 -> 6 -> 9
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

幅優先探索は、根の探索開始位置（preorderなら根→左→右）だけ覚えておけば再帰処理はそれに従って書ける。

深さ優先探索は少し面倒。後は二分探索木に限らずノードの削除処理はもっと面倒...。

# 所感
探索パターンはとりあえずこうやって覚えておけば頭には留められそう。

# 参考
- [www.momoyama-usagi.com - うさぎでもわかる2分探索木　後編　2分探索木における4つの走査方法](https://www.momoyama-usagi.com/entry/info-algo-tree-traverse)
  - 走査方法が大変覚えやすく、参考にさせて頂いた
