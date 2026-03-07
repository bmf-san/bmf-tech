---
title: Review of Basics of Algorithms and Data Structures
slug: algorithms-data-structures-review
date: 2023-06-28T00:00:00Z
author: bmf-san
categories:
  - Algorithms
  - Data Structures
tags:
  - Algorithm
  - Data Structure
translation_key: algorithms-data-structures-review
---

# Overview
As I resume my daily routine of solving coding quizzes, I will review the basics of algorithms and data structures as a form of rehabilitation.

# Arrays
- Memory is allocated in contiguous locations
- Single data type
- Size is generally fixed, but can be dynamic (variable-length arrays) in some languages
- Subarray
  - An array extracted in a contiguous manner from within an array
    - [1, 2, 3, 4, 5]
      - →[2, 3, 4]
- Subsequence
  - An array extracted with some elements removed without changing the order of elements
    - [1, 2, 3, 4, 5]
      - →[1, 3, 4]
  - A subarray can also be a subsequence

## Time Complexity
|        Operation        |  Complexity   |
| ----------------------- | ------------- |
| Access                  | O(1)         |
| Search                  | O(n)         |
| Search (sorted)        | O(log(n))    |
| Insert                  | O(n)         |
| Insert (at end)        | O(1)         |
| Delete                  | O(n)         |
| Delete (at end)        | O(1)         |

## Considerations
- How to handle duplicates within the array?
- Be careful not to go out of bounds when accessing by index

# Strings
- Depending on the language, they can be arrays, variable-length arrays, or objects
- Common tree structures for string searching
  - Trie (prefix tree)
  - Suffix tree

## Time Complexity
Omitted as it depends on the implementation

## Considerations
- Case sensitivity
- Can we utilize the ASCII Table?

# Hash Table (Hash Map)
- Associative array
- A structure that maps keys to values

|   Operation   | Complexity |
| ------------- | ---------- |
| Access        | N/A        |
| Search        | O(1)*     |
| Insert        | O(1)*     |
| Delete        | O(1)*     |

*Average case

## Considerations
- Is there a possibility of hash collisions?

# Recursion
- A process where a function calls itself within its own body

## Time Complexity
Omitted as it depends on the implementation

## Considerations
- Define the base case for recursion
- Be cautious of stack overflow with deep recursion
  - Tail Call Optimization (TCO) can help avoid this in supported languages

# Sorting and Searching
## Time Complexity of Sorting Algorithms
|  Algorithm       | Time Complexity | Space Complexity |
| ---------------- | ---------------- | ---------------- |
| Bubble Sort      | O(n^2)          | O(1)             |
| Insertion Sort   | O(n^2)          | O(1)             |
| Selection Sort   | O(n^2)          | O(1)             |
| Quick Sort       | O(nlog(n))      | O(log(n))        |
| Merge Sort       | O(nlog(n))      | O(n)             |
| Heap Sort        | O(nlog(n))      | O(1)             |
| Counting Sort    | O(n+k)          | O(k)             |
| Radix Sort       | O(nk)           | O(n+k)           |

## Time Complexity of Searching Algorithms
| Algorithm        |  Complexity   |
| ---------------- | ------------- |
| Binary Search    | O(log(n))     |

# Two-Dimensional Arrays
- An array that contains arrays
- A concept similar to vectors or matrices in mathematics
- A one-dimensional array is called a linear array, while a multi-dimensional array is called a multidimensional array

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
- Are there empty rows or columns? Are there arrays of length 0?
- How many dimensions are needed?

# Linked List
- A sequential data structure where each data element has a link (or pointer) to the next data element
- Data insertion is O(1)
- Data deletion is O(1) only if the position is specified
- Accessing data is linear
- Singly Linked List
  - Each data element only holds a reference to the next data element
- Doubly Linked List
  - Each data element holds references to both the previous and next elements
- Circular Linked List
  - The last data element holds a reference to the first data element

## Time Complexity
|   Operation   | Complexity |
| ------------- | ---------- |
| Access        | O(n)      |
| Search        | O(n)      |
| Insert        | O(1)      |
| Delete        | O(1)      |

# Queue
- A data structure generally maintained in a FIFO list structure
- Waiting line
- Enqueue
  - Adding to the queue
- Dequeue
  - Removing from the queue

## Time Complexity
|    Operation    | Complexity |
| --------------- | ---------- |
| Enqueue         | O(1)      |
| Dequeue         | O(1)      |

# Stack
- A data structure maintained in a LIFO or FIFO structure
- Push
  - Adding data
- Pop
  - Removing data

## Time Complexity
|   Operation   | Complexity |
| ------------- | ---------- |
| Push         | O(1)      |
| Pop          | O(1)      |

# Tree
- A data structure representing a hierarchical structure
- An undirected acyclic graph
- Binary Search Tree
  - cf. [Algorithms and Data Structures - Binary Search Tree](https://bmf-tech.com/posts/%e3%82%a2%e3%83%ab%e3%82%b4%e3%83%aa%e3%82%ba%e3%83%a0%e3%81%a8%e3%83%87%e3%83%bc%e3%82%bf%e6%a7%8b%e9%80%a0%20-%20%e4%ba%8c%e5%88%86%e6%8e%a2%e7%b4%a2%e6%9c%a8)

# Graph
- A data structure representing relationships between nodes
- Undirected or directed

## Time Complexity
Let V be the number of vertices and E be the number of edges.

|    Algorithm        | Complexity |
| ------------------- | ---------- |
| Depth First Search   | O(V+E)    |
| Breadth First Search | O(V+E)    |
| Topological Sort     | O(V+E)    |

## Heap
- A tree structure with the constraint that children are always greater (or smaller) than or equal to their parent

## Time Complexity
| Operation |  Complexity   |
| --------- | ------------- |
| Insert    | O(log(n))     |
| Delete    | O(log(n))     |

# Trie
- A tree structure specialized for efficient string searching and storage
- cf. [Implementing a Trie in Golang](https://bmf-tech.com/posts/Golang%e3%81%a7%e3%83%88%e3%83%a9%e3%82%a4%e6%9c%a8%e3%82%92%e5%ae%9f%e8%a3%85%e3%81%99%e3%82%8b)

## Time Complexity
Let m be the length of the string.

| Operation | Complexity |
| --------- | ---------- |
| Search    | O(m)      |
| Insert    | O(m)      |
| Delete    | O(m)      |

# References
- [www.techinterviewhandbook.org - study-cheatsheet](https://www.techinterviewhandbook.org/algorithms/study-cheatsheet/)
- [wikipedia.org - Data Structure](https://ja.wikipedia.org/wiki/%E3%83%87%E3%83%BC%E3%82%BF%E6%A7%8B%E9%80%A0)

# Others
Articles referenced while studying data structures and algorithms.

- [snamiki1212.com - Recommended Study Methods After Relearning Data Structures and Algorithms](https://snamiki1212.com/relearning-data-structure-and-algo/)