---
title: Introduction to Building an HTTP Router with net/http
slug: building-http-router-with-net-http
date: 2021-10-24T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - HTTP
description: This article explains how to create a custom HTTP router using Golang's standard package net/http.
translation_key: building-http-router-with-net-http
---



# Introduction
In this article, we will explain how to create a custom HTTP router using Golang's standard package, net/http.

The standard package does not offer many routing features.

For example, it does not allow defining routing for each HTTP method, using URLs as path parameters, or defining routing using regular expressions.

Therefore, in actual application development, it is not uncommon to introduce more functional HTTP routers.

However, by creating your own HTTP router, you can enjoy the following benefits:

- Learn about net/http
- Experience the fun of algorithms
- Implement an HTTP router that you find easy to use
- Integrate an HTTP router that you fully understand into your application

This article will explain how to create a custom HTTP router with the following structure:

- Introduction
- Table of Contents
- Chapter 1: What is an HTTP Router?
- Chapter 2: Data Structure of an HTTP Router
- Chapter 3: Code Reading of an HTTP Server
- Chapter 4: Implementation of an HTTP Router
- Conclusion

Each chapter includes columns for content that deviates from the main topic.

This article is intended to be meaningful for readers who:

- Understand Golang syntax and want to create something
- Want to deepen their understanding of Golang's standard packages
- Want to try implementing simple algorithms in Golang
- Find the standard routing features insufficient and want to know how to extend them
- Want to understand the implementation of the HTTP router they usually use

To fully understand the content of this article, it is helpful to have the following prerequisite knowledge:

- Basic understanding of Golang syntax
- Experience using some HTTP router

I have published an HTTP router package called [bmf-san/goblin](https://github.com/bmf-san/goblin).

Please take a look at the code and try using it. Contributions are also welcome.

# Table of Contents
- Introduction
- Chapter 1: What is an HTTP Router?
- Chapter 2: Data Structure of an HTTP Router
- Chapter 3: Code Reading of an HTTP Server
- Chapter 4: Implementation of an HTTP Router
- Conclusion

# Chapter 1: What is an HTTP Router?
An HTTP router is sometimes called a URL router or simply a router, but in this article, we will unify the term as HTTP router.

An HTTP router is an application that connects the requested URL with the response processing, as shown in the following diagram.

![route_in_client_and_server](/assets/images/posts/building-http-router-with-net-http/138551922-8fce2b9c-51d6-49e6-bea1-015e383cdb6e.png)

An HTTP router can perform routing based on data (hereafter referred to as a route map) that maps URLs to response processing.

|       Request URL       |     Handler      |
| :---------------------- | :--------------- |
| GET  /                  | IndexHandler     |
| GET  /foo               | FooHandler       |
| POST /foo               | FooHandler       |
| GET  /foo/:id           | FooHandler       |
| POST /foo/:id           | FooHandler       |
| GET  /foo/:id/:name     | FooHandler       |
| POST /foo/:id/:name     | FooHandler       |
| GET  /foo/bar           | FooBarHandler    |
| GET  /foo/bar/:id       | FooBarHandler    |
| GET  /foo/bar/:id/:name | FooBarHandler    |
| GET  /foo/bar/baz       | FooBarBazHandler |
| GET  /bar               | BarHandler       |
| GET  /baz               | BazHandler       |

Internally, the defined route map becomes a data structure optimized for routing.

The data structure will be explained in the next chapter.

In this article, we define "routing" as finding the response processing corresponding to the request URL based on the route map.

Also, we define an application that performs routing in HTTP as an "HTTP router."

---
### Column: URL Specifications

A URL represents the address of a page on the internet and stands for Uniform Resource Locator.

The format of a URL string is defined as follows:

```sh
<scheme>:<scheme-specific-part>
```

This part often uses protocol names such as http, https, ftp, but non-protocol schema names are also defined.

[Uniform Resource Identifier (URI) Schemes](https://www.iana.org/assignments/uri-schemes/uri-schemes.xhtml)

In the `<scheme-specific-part>`, a string based on the schema is defined.

For example, in the case of http and https schemes, there is a rule that defines domain names and path names (or directory names).

For detailed URL specifications, refer to RFC 1738.

[RFC 738 - Uniform Resource Locators (URL)](https://tools.ietf.org/html/rfc1738)

RFC 1738 is positioned as an internet standard (STD1).

# Chapter 2: Data Structure of an HTTP Router
## Considering the Data Structure
Below is the route map exemplified in Chapter 1.

|       Request URL       |     Handler      |
| :---------------------- | :--------------- |
| GET  /                  | IndexHandler     |
| GET  /foo               | FooHandler       |
| POST /foo               | FooHandler       |
| GET  /foo/:id           | FooHandler       |
| POST /foo/:id           | FooHandler       |
| GET  /foo/:id/:name     | FooHandler       |
| POST /foo/:id/:name     | FooHandler       |
| GET  /foo/bar           | FooBarHandler    |
| GET  /foo/bar/:id       | FooBarHandler    |
| GET  /foo/bar/:id/:name | FooBarHandler    |
| GET  /foo/bar/baz       | FooBarBazHandler |
| GET  /bar               | BarHandler       |
| GET  /baz               | BazHandler       |

Focusing on the URL, you can see that it has a hierarchical structure.

A hierarchical structure is well-suited to a tree structure, so we consider representing the route map as a tree structure.

## What is a Tree Structure?
A tree structure in graph theory is a data structure that represents a hierarchical structure.

A tree structure is suitable for representing hierarchical structures.

The elements that make up a tree are called nodes, the topmost node with no parent is called the root, and the lowest node with no children is called a leaf. The connections between nodes are called edges.

Adding a node to a tree is called insertion, and finding a node in a tree is called searching.

![tree_structure](/assets/images/posts/building-http-router-with-net-http/138551925-b0544b07-1f2d-44bb-9724-495780908b2d.png)

An example implementation of a basic tree, a binary search tree, is shown below.

```golang
package main

import (
    "fmt"
)

// Node is a node of a tree.
type Node struct {
    Key   int
    Left  *Node
    Right *Node
}

// BST is a binary search tree.
type BST struct {
    Root *Node
}

// insert insert a node to tree.
func (b *BST) insert(key int) {
    if b.Root == nil {
        b.Root = &Node{
            Key:   key,
            Left:  nil,
            Right: nil,
        }
    } else {
        recursiveInsert(b.Root, &Node{
            Key:   key,
            Left:  nil,
            Right: nil,
        })
    }
}

// recursiveInsert insert a new node to targetNode recursively.
func recursiveInsert(targetNode *Node, newNode *Node) {
    // if a newNode is smaller than targetNode, insert a newNode to left child node.
    // if a newNode is a bigger than targetNode, insert a newNode to right childe node.
    if newNode.Key < targetNode.Key {
        if targetNode.Left == nil {
            targetNode.Left = newNode
        } else {
            recursiveInsert(targetNode.Left, newNode)
        }
    } else {
        if targetNode.Right == nil {
            targetNode.Right = newNode
        } else {
            recursiveInsert(targetNode.Right, newNode)
        }
    }
}

// remove remove a key from tree.
func (b *BST) remove(key int) {
    recursiveRemove(b.Root, key)
}

// recursiveRemove remove a key from tree recursively.
func recursiveRemove(targetNode *Node, key int) *Node {
    if targetNode == nil {
        return nil
    }

    if key < targetNode.Key {
        targetNode.Left = recursiveRemove(targetNode.Left, key)
        return targetNode
    }

    if key > targetNode.Key {
        targetNode.Right = recursiveRemove(targetNode.Right, key)
        return targetNode
    }

    if targetNode.Left == nil && targetNode.Right == nil {
        targetNode = nil
        return nil
    }

    if targetNode.Left == nil {
        targetNode = targetNode.Right
        return targetNode
    }

    if targetNode.Right == nil {
        targetNode = targetNode.Left
        return targetNode
    }

    leftNodeOfMostRightNode := targetNode.Right

    for {
        if leftNodeOfMostRightNode != nil && leftNodeOfMostRightNode.Left != nil {
            leftNodeOfMostRightNode = leftNodeOfMostRightNode.Left
        } else {
            break
        }
    }

    targetNode.Key = leftNodeOfMostRightNode.Key
    targetNode.Right = recursiveRemove(targetNode.Right, targetNode.Key)
    return targetNode
}

// search search a key from tree.
func (b *BST) search(key int) bool {
    result := recursiveSearch(b.Root, key)

    return result
}

// recursiveSearch search a key from tree recursively.
func recursiveSearch(targetNode *Node, key int) bool {
    if targetNode == nil {
        return false
    }

    if key < targetNode.Key {
        return recursiveSearch(targetNode.Left, key)
    }

    if key > targetNode.Key {
        return recursiveSearch(targetNode.Right, key)
    }

    // targetNode == key
    return true
}

// depth-first search
// inOrderTraverse traverse tree by in-order.
func (b *BST) inOrderTraverse() {
    recursiveInOrderTraverse(b.Root)
}

// recursiveInOrderTraverse traverse tree by in-order recursively.
func recursiveInOrderTraverse(n *Node) {
    if n != nil {
        recursiveInOrderTraverse(n.Left)
        fmt.Printf("%d\n", n.Key)
        recursiveInOrderTraverse(n.Right)
    }
}

// depth-first search
// preOrderTraverse traverse by pre-order.
func (b *BST) preOrderTraverse() {
    recursivePreOrderTraverse(b.Root)
}

// recursivePreOrderTraverse traverse by pre-order recursively.
func recursivePreOrderTraverse(n *Node) {
    if n != nil {
        fmt.Printf("%d\n", n.Key)
        recursivePreOrderTraverse(n.Left)
        recursivePreOrderTraverse(n.Right)
    }
}

// depth-first search
// postOrderTraverse traverse by post-order.
func (b *BST) postOrderTraverse() {
    recursivePostOrderTraverse(b.Root)
}

// recursivePostOrderTraverse traverse by post-order recursively.
func recursivePostOrderTraverse(n *Node) {
    if n != nil {
        recursivePostOrderTraverse(n.Left)
        recursivePostOrderTraverse(n.Right)
        fmt.Printf("%v\n", n.Key)
    }
}

// breadth-first search
// levelOrderTraverse traverse by level-order.
func (b *BST) levelOrderTraverse() {
    if b != nil {
        queue := []*Node{b.Root}

        for len(queue) > 0 {
            currentNode := queue[0]
            fmt.Printf("%d ", currentNode.Key)

            queue = queue[1:]

            if currentNode.Left != nil {
                queue = append(queue, currentNode.Left)
            }

            if currentNode.Right != nil {
                queue = append(queue, currentNode.Right)
            }
        }
    }
}

func main() {
    tree := &BST{}

    tree.insert(10)
    tree.insert(2)
    tree.insert(3)
    tree.insert(3)
    tree.insert(3)
    tree.insert(15)
    tree.insert(14)
    tree.insert(18)
    tree.insert(16)
    tree.insert(16)

    tree.remove(3)
    tree.remove(10)
    tree.remove(16)

    fmt.Println(tree.search(10))
    fmt.Println(tree.search(19))

    // Traverse
    tree.inOrderTraverse()
    tree.preOrderTraverse()
    tree.postOrderTraverse()
    tree.levelOrderTraverse()

    fmt.Printf("%#v\n", tree)
}
```

While we will not explain in detail here, a binary search tree is a good tree to learn the basic algorithms of tree structures.

There are various types of tree structures besides binary search trees. Among them, a tree structure called a trie (also known as a prefix tree, but we will call it a trie in this article) has the characteristic of being easy to search for strings.

By using a trie, you can express a data structure that is easy to handle in routing.

## What is a Trie?
A trie is a type of tree structure used in IP address searches and morphological analysis.

Each node holds one or more strings or numbers, and by searching from the root node to the leaves, you can express words.

There is a service that allows you to visualize algorithms and understand them dynamically, which makes it easier to understand the data structure of a trie.

cf. [Algorithm Visualizations - Trie (Prefix Tree)](https://www.cs.usfca.edu/~galles/visualization/Trie.html)

A trie can be implemented relatively easily.

The following code is an example of a trie with only search and insert implemented.

```golang
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

By basing it on this trie, we will consider a data structure optimized for routing.

## Considering the Data Structure of a Route Map Based on a Trie
Based on the idea of a trie, we will consider the data structure of a route map.

Below is the data structure adopted in [bmf-san/goblin](https://github.com/bmf-san/goblin), which I am developing.

Since goblin supports middleware and path parameters, it has a data structure that corresponds to them.

![trie_based_tree_for_goblin](/assets/images/posts/building-http-router-with-net-http/138551926-666c7e6e-03f7-4a5a-8f18-ad1ba27b615a.png)

This data structure represents the following route map.

|   Request URL    |     Handler      | Middleware |
| :--------------- | :--------------- | :--------- |
| GET  /           | IndexHandler     | none       |
| GET  /foo        | FooHandler       | FooMws     |
| POST /foo        | FooHandler       | FooMws     |
| GET  /foo/bar    | FooBarHandler    | none       |
| GET /foo/bar/:id | FooBarHandler    | none       |
| GET /foo/baz     | FooBazHandler    | none       |
| GET /foo/bar/baz | FooBarBazHandler | none       |
| GET /baz         | BazHandler       | none       |

The considerations are summarized in the following two points:

- How to express the URL as a tree structure
- What data is necessary for nodes to hold

The former determines the performance of routing, and if you want to pursue processing time and memory efficiency, you need to consider adopting a more advanced tree structure.

The latter is related to the functionality of the HTTP router, so it varies depending on the features you want to provide.

The tree structure based on the trie introduced this time is just a tree structure I thought of.

The data structure varies depending on the implementation requirements of the HTTP router.

In the next chapter, we will explain what you need to know to incorporate this data structure into an HTTP router.

---
### Column: Radix Tree (Patricia Tree)

A radix tree is a tree structure that further develops the trie used to store strings.

The author has observed that radix trees are often used in HTTP routers that consider performance.

It seems to be used internally in Golang's strings package as well.

[https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/strings/strings.go;l=924](https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/strings/strings.go;l=924)

# Chapter 3: Code Reading of an HTTP Server
Before explaining the implementation of an HTTP router, let's look at the following code example of an HTTP server using net/http to understand what you need to know to implement an HTTP router.

Refer to the following link as needed.

[https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:](https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:)

```golang
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
```

Although this code is simple, it is insightful for creating your own HTTP router.

This code follows the flow of calling a multiplexer, registering a handler, and starting the server.

Let's look at each one in order.

## Calling the Multiplexer
The first code generates a structure called `http.ServeMux`.

```golang
mux := http.NewServeMux()
```

The net/http documentation explains that `http.ServeMux` is an HTTP request multiplexer (hereafter referred to as a multiplexer).

[type ServeMux](https://pkg.go.dev/net/http#ServeMux)

This multiplexer has the role of comparing the request URL with registered patterns and calling the handler (a function that returns a response) that matches the most.

In other words, `http.ServeMux` is a structure for routing.

This `http.ServeMux` has a method called `ServeHTTP` implemented.

```golang
// ServeHTTP dispatches the request to the handler whose
// pattern most closely matches the request URL.
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(StatusBadRequest)
		return
	}
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}
```

[cs.opensource.google - go1.17.2:src/net/http/server.go;l=2415](https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/net/http/server.go;l=2415)

If you read further into the following part of `ServeHTTP`, you can see the routing process of `ServeHTTP`.
```golang
h, _ := mux.Handler(r)
```

By jumping through the code in order, you will reach a function that finds and returns the matching handler.

```golang
// Find a handler on a handler map given a path string.
// Most-specific (longest) pattern wins.
func (mux *ServeMux) match(path string) (h Handler, pattern string) {
	// Check for exact match first.
	v, ok := mux.m[path]
	if ok {
		return v.h, v.pattern
	}

	// Check for longest valid match.  mux.es contains all patterns
	// that end in / sorted from longest to shortest.
	for _, e := range mux.es {
		if strings.HasPrefix(path, e.pattern) {
			return e.h, e.pattern
		}
	}
	return nil, ""
}
```

[https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/net/http/server.go;l=2287;drc=refs%2Ftags%2Fgo1.17.2](https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/net/http/server.go;l=2287;drc=refs%2Ftags%2Fgo1.17.2)

If a matching handler is found, it calls the `ServeHTTP` of that handler to invoke the processing for the response.

This is the process at the end of the `ServeHTTP` method implemented in `http.ServeMux`.

```golang
h.ServeHTTP(w, r)
```

To create your own HTTP router, you need to implement a multiplexer that can replace the standard multiplexer by satisfying the `http.Handler` type (≒ implementing `ServeHTTP`).

[type Handler](https://pkg.go.dev/net/http#Handler)

## Registering the Handler
Next, the following code registers a handler with the multiplexer.

```golang
mux.HandleFunc("/", handler)
```

The handler registered with the multiplexer must satisfy the `http.Handler` type (≒ `ServeHTTP` must be implemented).

```golang
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handler := foo{}
	mux.Handle("/", handler)

	http.ListenAndServe(":8080", mux)
}

type foo struct{}

// Satisfy the http.Handler type by implementing ServeHTTP.
func (f foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
```

Alternatively, you can create a handler by implementing the `http.HandlerFunc` type.

[func (HandlerFunc) ServeHTTP](https://pkg.go.dev/net/http#HandlerFunc.ServeHTTP)

The `http.HandlerFunc` type is defined as `func(ResponseWriter, *Request)` and implements the `ServeHTTP` method.

```golang
// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

[https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/net/http/server.go;l=2045](https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/net/http/server.go;l=2045)

Therefore, when using the `http.HandlerFunc` type, you can create a handler as follows:

```golang
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handler))

	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
```

When implementing an HTTP router, if you implement it to support the `http.Handler` type, you can provide flexibility in how handlers are created, making it an easy-to-use package.

## Starting the Server
The last code starts the HTTP server by passing the port number and the multiplexer to the function.

```golang
http.ListenAndServe(":8080", mux)
```

[func ListenAndServe](https://pkg.go.dev/net/http#Server.ListenAndServe)

Internally, the `ListenAndServe` of the `http.Server` type is called.

[func (*Server) ListenAndServe](https://pkg.go.dev/net/http#Server.ListenAndServe)

In this function, if the second argument is nil, the default `http.ServeMux`, `http.DefaultServeMux`, is used.

In other words, unless you want to extend the multiplexer, you don't have to generate a multiplexer.

In the process of implementing an HTTP router, it was necessary to generate a multiplexer as a prelude to the story, so I raised the code as an example.

Now that we have done the necessary code reading for implementing an HTTP router, we will explain the implementation in the next chapter.

---
### Column: About Trailing Slashes

The trailing `/` in a URL has different meanings for the end of a domain name and the end of a subdirectory.

In the case of a domain name, most browsers request a URL with `/` if it is missing.

- `https://bmf-tech.com` → Request to `https://bmf-tech.com/`
- `https://bmf-tech.com/` → Request to `https://bmf-tech.com`

In the case of a domain name, there is not much difference in the presence or absence of `/`, but in the case of a subdirectory, there is a clear difference.

- `https://bmf-tech.com/posts` → Request to a file
- `https://bmf-tech.com/posts/` → Request to a directory

If you want to know more about the specifications, refer to the RFC.

[RFC 2616](https://www.w3.org/Protocols/rfc2616/rfc2616-sec5.html#sec5.1)
[RFC 3986](https://www.ietf.org/rfc/rfc3986.txt)

When implementing an HTTP router, it is necessary to be aware of how to interpret the path part of the URL.

In [bmf-san/goblin](https://github.com/bmf-san/goblin), which I developed, the presence or absence of a trailing `/` is treated as the same routing definition.

# Chapter 4: Implementation of an HTTP Router
Now that we are ready to implement an HTTP router, let's explain the implementation.

This time, we will implement a router that is slightly more functional than the standard package.

Specifically, it will be a router with the following two features:

- Supports method-based routing
- Implements an algorithm based on a trie

The standard package does not allow routing registration by HTTP method.

If you want to route by HTTP method, you need to implement conditional branching for each HTTP method in the handler.

```golang
// ex.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			// do something...
		case http.MethodPost:
			// do something...

		...

		default:
```

We will implement a feature that allows routing to be defined for each method without defining such conditional branching in the handler.

The algorithm for an HTTP router that allows routing to be defined based on methods adopts a tree structure based on the trie explained in the data structure of the HTTP router.

## Preparation
The source code for the HTTP router implemented this time is available below.

[bmf-san/introduction-to-golang-http-router-made-with-net-http](https://github.com/bmf-san/introduction-to-golang-http-router-made-with-net-http)

It is recommended to write test code during the implementation process, but we will not explain the test code.

The same applies to CI.

We are using Golang version 1.17.

## Implementation
As an implementation procedure, we will start by implementing the routing algorithm based on a trie.

After that, we will implement support for method-based routing.

### Implementation of the Routing Algorithm Based on a Trie
Let's start implementing it right away.

All the code implemented here can be referenced below.

[bmf-san/introduction-to-golang-http-router-made-with-net-http/blob/main/trie.go](https://github.com/bmf-san/introduction-to-golang-http-router-made-with-net-http/blob/main/trie.go)

This time, we will adopt the following tree structure, which is a simplified version of the data structure of goblin.

![tree_for_implementation](/assets/images/posts/building-http-router-with-net-http/138551924-ca2e28f5-a0c1-4ae4-ac9f-76d6ee1c46d9.png)

The route map represented by this tree structure is as follows.

|       Request URL       |     Handler      |
| :---------------------- | :--------------- |
| GET  /                  | IndexHandler     |
| GET  /foo               | FooHandler       |
| POST /foo               | FooHandler       |
| GET  /foo/bar           | FooBarHandler    |
| GET  /foo/bar/baz       | FooBarBazHandler |
| GET  /bar               | BarHandler       |
| GET  /baz               | BazHandler       |

To express the above tree structure, let's start by defining the necessary data.

Create a file called `trie.go` and define the structure.

```golang
package myrouter

// tree is a trie tree.
type tree struct {
	node *node
}

type node struct {
	label    string
	actions  map[string]*action // key is method
	children map[string]*node   // key is a label o f next nodes
}

// action is an action.
type action struct {
	handler http.Handler
}

// result is a search result.
type result struct {
	actions *action
}
```

`tree` is the tree itself, `node` is the element that makes up the tree, and it has `label`, `actions`, and `children`.

`label` represents the path of the URL, `actions` represents the map of HTTP methods and handlers, and `children` represents the map of `label` and `node` for child nodes.

`result` represents the search result from the tree.

Next, let's define the functions to generate these structures.

```golang
// newResult creates a new result.
func newResult() *result {
	return &result{}
}

// NewTree creates a new trie tree.
func NewTree() *tree {
	return &tree{
		node: &node{
			label:    pathRoot,
			actions:  make(map[string]*action),
			children: make(map[string]*node),
		},
	}
}
```

Now, let's implement the process of adding nodes to the tree.

Define the `Insert` method with `tree` as the pointer receiver.

```golang
func (t *tree) Insert(methods []string, path string, handler http.Handler)) {
	// 
}
```

The point of the arguments in this function is that it is defined to pass multiple HTTP methods.

It allows you to define a single handler for each HTTP method as well as define the same handler for multiple methods.

```golang
// ex.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			// do something...
		case http.MethodPost:
			// do something...

		...

		default:
```

Considering the possibility that there may be cases where you want to implement conditional branching for HTTP methods in the handler, we have made it versatile.

Next, in the `Insert` method, first define the node that will be the starting point as a variable.

```golang
func (t *tree) Insert(methods []string, path string, handler http.Handler)) {
	curNode := t.node
}
```

Next, process the conditional branching for when the target is `/` (root).

```golang
const (
	pathRoot      string = "/"
)

func (t *tree) Insert(methods []string, path string, handler http.Handler)) {
	curNode := t.node
	if path == pathRoot {
		curNode.label = path
		for _, method := range methods {
			curNode.actions[method] = &action{
				handler: handler,
			}
		}
		return nil
	}
}
```

In the case of `/`, there is no need to perform the subsequent loop processing, so we add a node to the tree here and end the process.

If it is not `/`, continue the process.

Process the URL path by splitting it with `/` and storing the string in a []string type slice.

```golang
const (
	...
	pathDelimiter string = "/"
)

func (t *tree) Insert(methods []string, path string, handler http.Handler)) {
	...

	ep := explodePath(path)	
}

// explodePath removes an empty value in slice.
func explodePath(path string) []string {
	s := strings.Split(path, pathDelimiter)
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
```

The []string type slice is used to find the position to add the node by running a range loop.

The process here is based on the implementation of the trie explained in the data structure of the HTTP router.

Add a node when a child node is not found.

If there is a case where the routing definition overlaps, the specification is to overwrite it.

```golang
// Insert inserts a route definition to tree.
func (t *tree) Insert(methods []string, path string, handler http.Handler) error {
	...

	for i, p := range ep {
		nextNode, ok := curNode.children[p]
		if ok {
			curNode = nextNode
		}
		// Create a new node.
		if !ok {
			curNode.children[p] = &node{
				label:    p,
				actions:  make(map[string]*action),
				children: make(map[string]*node),
			}
			curNode = curNode.children[p]
		}
		// last loop.
		// If there is already registered data, overwrite it.
		if i == len(ep)-1 {
			curNode.label = p
			for _, method := range methods {
				curNode.actions[method] = &action{
					handler: handler,
				}
			}
			break
		}
	}

	return nil
}
```

The final implementation of `Insert` is as follows.

```golang
// Insert inserts a route definition to tree.
func (t *tree) Insert(methods []string, path string, handler http.Handler) error {
	curNode := t.node
	if path == pathRoot {
		curNode.label = path
		for _, method := range methods {
			curNode.actions[method] = &action{
				handler: handler,
			}
		}
		return nil
	}
	ep := explodePath(path)
	for i, p := range ep {
		nextNode, ok := curNode.children[p]
		if ok {
			curNode = nextNode
		}
		// Create a new node.
		if !ok {
			curNode.children[p] = &node{
				label:    p,
				actions:  make(map[string]*action),
				children: make(map[string]*node),
			}
			curNode = curNode.children[p]
		}
		// last loop.
		// If there is already registered data, overwrite it.
		if i == len(ep)-1 {
			curNode.label = p
			for _, method := range methods {
				curNode.actions[method] = &action{
					handler: handler,
				}
			}
			break
		}
	}

	return nil
}

// explodePath removes an empty value in slice.
func explodePath(path string) []string {
	s := strings.Split(path, pathDelimiter)
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
```

Now that we have implemented the process of inserting into the tree, let's implement the process of searching from the tree next.

Compared to insertion, searching is relatively simple, so let's explain it all at once.

```golang
func (t *tree) Search(method string, path string) (*result, error) {
	result := newResult()
	curNode := t.node
	if path != pathRoot {
		for _, p := range explodePath(path) {
			nextNode, ok := curNode.children[p]
			if !ok {
				if p == curNode.label {
					break
				} else {
					return nil, ErrNotFound
				}
			}
			curNode = nextNode
			continue
		}
	}
	result.actions = curNode.actions[method]
	if result.actions == nil {
		// no matching handler was found.
		return nil, ErrMethodNotAllowed
	}
	return result, nil
}
```

In the case of searching, as with insertion, whether the URL path is `/` or not determines whether to proceed with the loop process.

If you proceed with the loop process, you look at the child nodes and search for the target node.

If the target node exists, find the handler that matches the request's HTTP method and return `result`.

### Implementing Support for Method-Based Routing
The overall code implemented here is as follows.

[bmf-san/introduction-to-golang-http-router-made-with-net-http/blob/main/router.go](https://github.com/bmf-san/introduction-to-golang-http-router-made-with-net-http/blob/main/router.go)

Here, we will also implement the functionality to provide as an HTTP router.

First, define the structure and the function for generation.

```golang
// Router represents the router which handles routing.
type Router struct {
	tree *tree
}

// route represents the route which has data for a routing.
type route struct {
	methods []string
	path    string
	handler http.Handler
}

func NewRouter() *Router {
	return &Router{
		tree: NewTree(),
	}
}
```

`Router` corresponds to `http.ServeMux` in net/http.

`route` holds data for defining routing.

Next, implement the following three methods for `Router`.

```golang
...

func (r *Router) Methods(methods ...string) *Router {
	tmpRoute.methods = append(tmpRoute.methods, methods...)
	return r
}

// Handler sets a handler.
func (r *Router) Handler(path string, handler http.Handler) {
	tmpRoute.handler = handler
	tmpRoute.path = path
	r.Handle()
}

// Handle handles a route.
func (r *Router) Handle() {
	r.tree.Insert(tmpRoute.methods, tmpRoute.path, tmpRoute.handler)
	tmpRoute = &route{}
}
```

`Methods` is a setter for HTTP methods, `Handler` is a setter for the URL path and handler, and calls `Handle`. `Handle` calls the insertion process into the tree implemented earlier.

`Methods` and `Handler` are implemented as method chains, considering the readability for those using the HTTP router.

Method-based routing can be realized by combining it with the tree.

Finally, implement `ServeHTTP` for `Router` to complete it.

```golang
...

// ServeHTTP dispatches the request to the handler whose
// pattern most closely matches the request URL.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path
	result, err := r.tree.Search(method, path)
	if err != nil {
		status := handleErr(err)
		w.WriteHeader(status)
		return
	}
	h := result.actions.handler
	h.ServeHTTP(w, req)
}

func handleErr(err error) int {
	var status int
	switch err {
	case ErrMethodNotAllowed:
		status = http.StatusMethodNotAllowed
	case ErrNotFound:
		status = http.StatusNotFound
	}
	return status
}
```

## Using the Implemented HTTP Router
The HTTP router implemented this time can be used as follows.

Start the server and make requests to each endpoint to check the operation.

```golang
package main

import (
	"fmt"
	"net/http"

	myroute "github.com/bmf-san/introduction-to-golang-http-router-made-with-net-http"
)

func indexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /")
	})
}

func fooHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "GET /foo")
		case http.MethodPost:
			fmt.Fprintf(w, "POST /foo")
		default:
			fmt.Fprintf(w, "Not Found")
		}
	})
}

func fooBarHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /foo/bar")
	})
}

func fooBarBazHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /foo/bar/baz")
	})
}

func barHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /bar")
	})
}

func bazHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /baz")
	})
}

func main() {
	r := myroute.NewRouter()

	r.Methods(http.MethodGet).Handler(`/`, indexHandler())
	r.Methods(http.MethodGet, http.MethodPost).Handler(`/foo`, fooHandler())
	r.Methods(http.MethodGet).Handler(`/foo/bar`, fooBarHandler())
	r.Methods(http.MethodGet).Handler(`/foo/bar/baz`, fooBarBazHandler())
	r.Methods(http.MethodGet).Handler(`/bar`, barHandler())
	r.Methods(http.MethodGet).Handler(`/baz`, bazHandler())

	http.ListenAndServe(":8080", r)
}
```

Although it became a bit rushed, the explanation of the implementation is now complete.

---
### Column: Performance Comparison of HTTP Routers

If you are interested in performance comparisons of HTTP routers, check out the following repository.

[julienschmidt/go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark)

I submitted a PR for goblin's performance comparison to this repository.

[Add a new router goblin #97](https://github.com/julienschmidt/go-http-routing-benchmark/pull/97)

# Conclusion
In this article, we explained the approach to creating a custom HTTP router.

In Chapter 1, we organized what an HTTP router is.

In Chapter 2, we explained the data structure of an HTTP router with examples.

In Chapter 3, we delved into the code of an HTTP server using net/http.

And in Chapter 4, we explained the implementation method of an HTTP router with code.

I hope that this article has been helpful or sparked interest in something for the readers.

Also, I would be happy if it becomes an opportunity to look at the code of my work, [bmf-san/goblin](https://github.com/bmf-san/goblin).

If you have any questions, requests for corrections, or feedback, please let me know.

# Afterword
- [zenn.dev - Introduction to Building an HTTP Router with net/http](https://zenn.dev/bmf_san/books/3f41c5cd34ec3f)
  - The content of this article is made into a book.
- [dev.to - Introduction to Golang HTTP router made with net/http](http://web.archive.org/web/20250815231536/https://dev.to/bmf_san/introduction-to-golang-http-router-made-with-nethttp-3nmb)
  - Translated into English.
- ~~github.com - bmf-san/book-introduction-to-golang-http-router-made-with-net-http~~
  - Repository for managing the original text.
- [Introduction to Building an HTTP Router with net/http](https://speakerdeck.com/bmf_san/httpdetukuruhttprutazi-zuo-ru-men)
  - Presented at Go Conference 2021
