---
title: Adjacency List and Adjacency Matrix
description: 'Compare adjacency lists (O(V+E) space, efficient for sparse graphs) and adjacency matrices (O(V²) space, O(1) edge lookup). Includes Go implementations for directed and undirected graphs.'
slug: adjacency-list-matrix
date: 2023-07-29T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Graph
  - Adjacency Matrix
  - Adjacency List
translation_key: adjacency-list-matrix
---

# Overview
This post summarizes the data structures for representing graphs: adjacency lists and adjacency matrices.

Both adjacency lists and adjacency matrices can be used for directed and undirected graphs.

# Adjacency List
A data structure that holds a list of adjacent vertices for each vertex (node).

```sh
// Example of an undirected graph
A---B
|   / |
| /   |
C---D

// Represented as an adjacency list, it looks like this:
A: [B, C]
B: [A, C, D]
C: [A, B, D]
D: [B, C]
```

## Time Complexity
Space complexity: O(V+E) (where V is the number of vertices and E is the number of edges)
Finding adjacent vertices for a specific vertex: O(1)
Determining the existence of a specific edge: O(degree) (where degree is the number of adjacent edges)
Enumerating all adjacent vertices for all vertices: O(V+E)

The adjacency list is an efficient data structure for graphs with a small number of edges.

## Implementation
The source code can be found at [adjacency_list](https://github.com/bmf-san/road-to-algorithm-master/tree/master/data_structures/graph/adjacency_list).

```go
// See: ~~https://www.youtube.com/watch?v=JDP1OVgoa0Q~~
// See: ~~https://www.youtube.com/watch?v=bSZ57h7GN2w~~
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

// addVertex adds a vertex to the graph.
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

Adding edges can be complex.

# Adjacency Matrix
A data structure that represents the connection relationships between vertices (nodes) as a two-dimensional matrix.

The presence or absence of edges between vertices is represented using values of 0 or 1.

```sh
// Example of an undirected graph
A---B
|   / |
| /   |
C---D

// Represented as an adjacency matrix, it looks like this:
    A  B  C  D
A  0  1  1  0
B  1  0  1  1
C  1  1  0  1
D  0  1  1  0
```

## Time Complexity
Space complexity: O(V^2)
Finding adjacent vertices for a specific vertex: O(V)
Determining the existence of a specific edge: O(1)
Enumerating all adjacent vertices for all vertices: O(V^2)

The adjacency matrix is an efficient data structure for graphs with many edges or when frequent edge existence checks are required.

## Implementation
The source code can be found at [adjacency_matrix](https://github.com/bmf-san/road-to-algorithm-master/tree/master/data_structures/graph/adjacency_matrix).

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

The conditions for adding edges may be a bit complicated.

# References
- [mathwords.net - Data Structures for Graphs (Adjacency Matrix and Adjacency List)](https://mathwords.net/gurahu)
- ~~Graph data structure and graph representation (Part 1 of 2)~~
- ~~Graph data structure and graph representation in Golang (Part 2 of 2)~~