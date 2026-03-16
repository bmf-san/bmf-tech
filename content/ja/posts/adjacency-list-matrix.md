---
title: 隣接リストと隣接行列
description: 隣接リストと隣接行列
slug: adjacency-list-matrix
date: 2023-07-29T00:00:00Z
author: bmf-san
categories:
  - アルゴリズムとデータ構造
tags:
  - グラフ
  - 隣接行列
  - 隣接リスト
translation_key: adjacency-list-matrix
---


# 概要　
グラフを表現するためのデータ構造である隣接リストと隣接リストについてまとめる。

隣接リストや隣接行列は有向グラフでも無向グラフでも利用できる。

# 隣接リスト（Adjacency List）
各頂点（ノード）ごとに隣接する頂点をリストとして持つデータ構造。

```sh
// 無向グラフの例
A---B
|   / |
| /   |
C---D

// 隣接リストで表現すると、次のような形になる
A: [B, C]
B: [A, C, D]
C: [A, B, D]
D: [B, C]
```

## 計算量
空間計算量： O(V+E) ※Vは頂点数、Eは辺の数
特定の頂点の隣接する頂点を見つける： O(1)
特定の辺の存在を判定する: O(degree) ※degreeは隣接する辺の数
全ての頂点の隣接する頂点を列挙する: O(V+E)

隣接リストは辺の数が少ないグラフだと計算量が効率なデータ構造。

## 実装
ソースコードは[adjacency_list](https://github.com/bmf-san/road-to-algorithm-master/tree/master/data_structures/graph/adjacency_list)。

```go
// See: https://www.youtube.com/watch?v=JDP1OVgoa0Q
// See: https://www.youtube.com/watch?v=bSZ57h7GN2w
package main

import "fmt"

// Graph represents an adjacency list graph.
type graph struct {
	vertices []*vertex
}

// Vertex represents a graph vertex.
type vertex struct {
	key int
	adj []*vertex
}

// addVertex adds a vertext to the graph.
func (g *graph) addVertex(k int) {
	if contains(g.vertices, k) {
		fmt.Println(fmt.Errorf("Vertex %v not added because it is an existing key", k))
		return
	}
	g.vertices = append(g.vertices, &vertex{key: k})
}

// addEdge adds an edge to the graph.
func (g *graph) addEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	// check error
	if fromVertex == nil || toVertex == nil {
		fmt.Println("Invalid edge")
		return
	}

	// check if edge already exists
	if contains(fromVertex.adj, to) {
		fmt.Println("Existing edge")
		return
	}

	// add edge
	fromVertex.adj = append(fromVertex.adj, toVertex)
}

// getVertex returns a pointer to the vertex.
func (g *graph) getVertex(k int) *vertex {
	for _, v := range g.vertices {
		if k == v.key {
			return v
		}
	}
	return nil
}

// contains returns true if the key exists in the slice.
func contains(s []*vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// print prints the adjacency list.
func (g *graph) print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adj {
			fmt.Printf("%v ", v.key)
		}
	}
}

func main() {
	g := &graph{}
	for i := 0; i < 5; i++ {
		g.addVertex(i)
	}
	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(2, 3)
	g.addEdge(4, 1)
	g.addEdge(4, 2)
	g.addEdge(4, 3)
	g.print()
}
```

辺（エッジ）を足す処理が複雑である。

# 隣接行列（Adjacency Matrix）
頂点（ノード）間の接続関係を2次元の行列として表現するデータ構造。

頂点間の辺（エッジ）の有無は0または1の値が使われる。

```sh
// 無向グラフの例
A---B
|   / |
| /   |
C---D

// 隣接行列で表現すると次のようになる。
    A  B  C  D
A  0  1  1  0
B  1  0  1  1
C  1  1  0  1
D  0  1  1  0
```

## 計算量
空間計算量: O(V^2)
特定の頂点の隣接する頂点を見つける： O(V)
特定の辺の存在を判定する: O(1)
全ての頂点の隣接する頂点を列挙する: O(V^2)

隣接行列は辺の数が多いグラフや辺の存在を頻繁に判定するような必要がある場合に効率的なデータ構造。

## 実装
ソースコードは[adjacency_matrix](https://github.com/bmf-san/road-to-algorithm-master/tree/master/data_structures/graph/adjacency_matrix)。

```go
package main

import (
	"fmt"
)

// graph represents an adjacency matrix graph.
type graph struct {
	matrix [][]int
	size   int
}

// newGraph returns a new graph with the given size.
func newGraph(size int) *graph {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}
	return &graph{matrix: matrix, size: size}
}

// addEdge adds an edge to the graph from -> to.
func (g *graph) addEdge(from, to int) {
	if from < 0 || to < 0 || from >= g.size || to >= g.size {
		return
	}
	g.matrix[from][to] = 1
}

// print prints the adjacency matrix.
func (g *graph) print() {
	for _, row := range g.matrix {
		fmt.Println(row)
	}
}

func main() {
	size := 5
	graph := newGraph(size)

	graph.addEdge(0, 1)
	graph.addEdge(0, 3)
	graph.addEdge(1, 2)
	graph.addEdge(2, 4)
	graph.addEdge(3, 4)

	graph.print()
}
```

辺（エッジ）を足す部分の条件が少しややこしいかもしれない。

# 参考
- [mathwords.net - グラフを表すデータ構造（隣接行列と隣接リスト）](https://mathwords.net/gurahu)
- [Graph data structure and graph representation (Part 1 of 2)](https://www.youtube.com/watch?v=JDP1OVgoa0Q)
- [Graph data structure and graph representation in Golang (Part 2 of 2)](https://www.youtube.com/watch?v=bSZ57h7GN2w)
