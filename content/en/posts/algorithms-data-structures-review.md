---
title: Reviewing the Basics of Algorithms and Data Structures
slug: algorithms-data-structures-review
date: 2023-06-28T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Algorithm
  - Data Structure
description: 'A comprehensive review of algorithms and data structures basics: arrays, strings, hash tables, linked lists, trees, stacks, queues, sorting, and time complexity for coding problems.'
translation_key: algorithms-data-structures-review
---

# Overview
As I resume the daily routine of solving coding quizzes, I am reviewing the basics of algorithms and data structures as a form of rehabilitation.

# Arrays
- Memory is allocated in a contiguous arrangement
- Single data type
- Size is generally fixed, but can be dynamic (variable-length arrays) depending on the language
- Subarray
  - An array extracted in a continuous form from within an array
    - [1, 2, 3, 4, 5]
      - →[2, 3, 4]
- Subsequence
  - An array extracted with some elements removed without changing the order of elements
    - [1, 2, 3, 4, 5]
      - →[1, 3, 4]
  - A subarray can also be a subsequence

## Time Complexity
| Operation        | Complexity |
| ---------------- | ---------- |
| Access           | O(1)       |
| Search           | O(n)       |
| Search (sorted)  | O(log(n))  |
| Insertion        | O(n)       |
| Insertion (end)  | O(1)       |
| Deletion         | O(n)       |
| Deletion (end)   | O(1)       |

## Considerations
- How to handle duplicates within the array?
- Be careful not to go out of range with index access

# Strings
- Depending on the language, it can be an array, a variable-length array, or an object
- Common tree structures for string search
  - Trie (prefix tree)
  - Suffix tree

## Time Complexity
Omitted due to implementation

## Considerations
- Case sensitivity
- Can the ASCII Table be utilized?

# Hash Tables (Hash Maps)
- Associative array
- Structure mapping keys to values

| Operation | Complexity |
| --------- | ---------- |
| Access    | N/A        |
| Search    | O(1)*      |
| Insertion | O(1)*      |
| Deletion  | O(1)*      |

*Average case

## Considerations
- Is there a possibility of hash collisions?

# Recursion
- A process where a function calls itself within its own definition

## Time Complexity
Omitted due to implementation

## Considerations
- Define the termination condition for recursion
- Be cautious of stack overflow with deep recursion
  - Easier to avoid in languages that support Tail Call Optimization (TCO)

# Sorting and Searching
## Sorting Algorithm Complexity
| Algorithm     | Time Complexity | Space Complexity |
| ------------- | --------------- | ---------------- |
| Bubble Sort   | O(n^2)          | O(1)             |
| Insertion Sort| O(n^2)          | O(1)             |
| Selection Sort| O(n^2)          | O(1)             |
| Quick Sort    | O(nlog(n))      | O(log(n))        |
| Merge Sort    | O(nlog(n))      | O(n)             |
| Heap Sort     | O(nlog(n))      | O(1)             |
| Counting Sort | O(n+k)          | O(k)             |
| Radix Sort    | O(nk)           | O(n+k)           |

## Search Algorithm Complexity
| Algorithm | Complexity |
| --------- | ---------- |
| Binary Search | O(log(n)) |

# Two-Dimensional Arrays
- A structure where an array holds arrays
- A concept similar to vectors or matrices in mathematics
- A 1D array is a linear array, and a multidimensional array is a multidimensional array

```go
package main

import "fmt"

func main() {
    matrix := [][]int{[]int{0, 1}, []int{2, 3}, []int{4, 5}}
    for i, v := range matrix {
        for j, _ := range v {
            // i is row and j is column
            matrix[i][j] = 0
        }
    }
    fmt.Println(matrix)
}
```

## Considerations
- Are there empty rows or columns? Is there an array of length 0?
- How many dimensions are needed?

# Linked Lists
- A sequential data structure where each data holds a link (or pointer) to the next data
- Data insertion is O(1)
- Data deletion is O(1) only if the position is specified
- Access to data is linear
- Singly linked list
  - Each data holds a reference to the next data only
- Doubly linked list
  - Each data holds references to both the previous and next data
- Circular list
  - The last data holds a reference to the first data

## Time Complexity
| Operation | Complexity |
| --------- | ---------- |
| Access    | O(n)       |
| Search    | O(n)       |
| Insertion | O(1)       |
| Deletion  | O(1)       |

# Queue
- A data structure generally held in a FIFO list structure
- Waiting line
- Enqueue
  - Adding to the queue
- Dequeue
  - Removing from the queue

## Time Complexity
| Operation | Complexity |
| --------- | ---------- |
| Enqueue   | O(1)       |
| Dequeue   | O(1)       |

# Stack
- A data structure held in a LIFO or FIFO structure
- Push
  - Adding data
- Pop
  - Removing data

## Time Complexity
| Operation | Complexity |
| --------- | ---------- |
| Push      | O(1)       |
| Pop       | O(1)       |

# Trees
- A data structure representing a hierarchical structure
- An undirected acyclic graph
- Binary Search Tree
  - cf. [Algorithms and Data Structures - Binary Search Tree](https://bmf-tech.com/posts/%e3%82%a2%e3%83%ab%e3%82%b4%e3%83%aa%e3%82%ba%e3%83%a0%e3%81%a8%e3%83%87%e3%83%bc%e3%82%bf%e6%a7%8b%e9%80%a0%20-%20%e4%ba%8c%e5%88%86%e6%8e%a2%e7%b4%a2%e6%9c%a8)

# Graphs
- A data structure representing relationships between nodes
- Undirected or directed

## Complexity
Let V be vertices and E be the number of edges.

| Algorithm          | Complexity |
| ------------------ | ---------- |
| Depth-First Search | O(V+E)     |
| Breadth-First Search | O(V+E)   |
| Topological Sort   | O(V+E)     |

## Heap
- A tree structure with the constraint that children are always greater (or smaller) or equal to the parent

## Time Complexity
| Operation | Complexity |
| --------- | ---------- |
| Insertion | O(log(n))  |
| Deletion  | O(log(n))  |

# Trie
- A tree structure specialized for efficient string search and storage
- cf. [Implementing Trie in Golang](https://bmf-tech.com/posts/Golang%e3%81%a7%e3%83%88%e3%83%a9%e3%82%a4%e6%9c%a8%e3%82%92%e5%ae%9f%e8%a3%85%e3%81%99%e3%82%8b)

## Time Complexity
Let m be the length of the string.

| Operation | Complexity |
| --------- | ---------- |
| Search    | O(m)       |
| Insertion | O(m)       |
| Deletion  | O(m)       |

# References
- [www.techinterviewhandbook.org - study-cheatsheet](https://www.techinterviewhandbook.org/algorithms/study-cheatsheet/)
- [wikipedia.org - Data Structure](https://ja.wikipedia.org/wiki/%E3%83%87%E3%83%BC%E3%82%BF%E6%A7%8B%E9%80%A0)

# Others
Articles referenced while studying data structures and algorithms.

- [snamiki1212.com - Relearning Data Structures and Algorithms: Recommended Study Methods](https://snamiki1212.com/relearning-data-structure-and-algo/)