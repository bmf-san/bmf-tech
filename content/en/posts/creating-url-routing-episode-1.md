---
title: 'Creating URL Routing: Episode 1'
slug: creating-url-routing-episode-1
date: 2018-12-19T00:00:00Z
author: bmf-san
categories:
  - Algorithms
  - Data Structures
tags:
  - HTTP
  - URL Routing
  - Tree Structure
  - Router
description: Exploring the basics of URL routing and implementing a tree structure-based routing system.
translation_key: creating-url-routing-episode-1
---

# Creating URL Routing: Episode 1

## Overview
Previously, I created a very basic routing system in React (cf. [Creating a custom router using React and History API](https://bmf-tech.com/posts/React%E3%81%A8History%20API%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6router%E3%82%92%E8%87%AA%E4%BD%9C%E3%81%99%E3%82%8B)), but I decided to challenge myself to create a more robust routing system. The motivation came from working with Golang recently.

In Golang, the standard library allows for lightweight application implementation, but the routing functionality is somewhat underpowered, often requiring reliance on external libraries. This inspired me to learn how to create my own routing system, which could expand my capabilities not only in Golang but also in other languages.

## What is URL Routing?
URL routing determines the process to execute based on the requested URL. It also handles path parameters and query parameters as needed during execution.

## URL Routing Implementation Patterns
Broadly speaking, there are two patterns:

- Using regular expressions for URL matching
- Using tree structures for string searching

Although routing may not significantly impact application execution speed, optimizing memory usage and computational complexity with efficient algorithms is always preferable, regardless of the programming language.

For this implementation, I chose the tree structure pattern. While I haven't measured performance, I believe tree structure algorithms are computationally more efficient than regular expressions. Many libraries are implemented using tree structures.

## What is a Tree Structure?
A tree structure is a type of data structure defined in the mathematical field of graph theory. A tree in graph theory consists of multiple nodes (also called vertices) and edges.

```
   ○ ・・・Root
 / | ・・・Edge
◯  ◯ ・・・Node (Vertex)
    \
  　  ◯
       \
    　   ○ ・・・Leaf
```

There are various types of tree structures depending on node properties and tree height, but we'll omit those details here.

## Examples of Tree Structures
- Family trees
- File systems
- Domain names
  - cf. https://www.nic.ad.jp/ja/dom/system.html
- Syntax trees
  - Used in compilers
- DOM trees
- Hierarchical tags or category structures

## Creating URL Routing
What should be treated as a tree structure? Naturally, the list of route definitions will be treated as a tree structure.

The implementation process can be summarized as follows: given route definitions and the current URL (path) as input, generate a tree structure from the route definitions, search the tree structure using the current URL (path) as the target, and return the matched data.

When working with tree structures, operations like adding or deleting nodes are sometimes implemented. However, for URL routing, these operations are unnecessary for now.

### Defining the Data Structure
First, define the DSL (Domain-Specific Language) for routing. Many libraries provide simple DSLs, but this time, I’ll define a slightly more complex DSL with multiple levels.

```:php
$routes = [
    '/' => [
        'GET' => 'HomeController@get',
    ],
    '/users' => [
        '/' => [
            'GET' => 'UserController@get',
        ],
        '/:user_id' => [
            '/' => [
                'GET' => 'UserController@get',
                'POST' => 'UserController@post',
            ],
            '/events' =>  [
                '/' => [
                    'GET' => 'EventController@get',
                ],
                '/:id' => [
                    'GET' => 'EventController@get',
                    'POST' => 'EventController@post',
                ],
            ]
        ],
        '/support' => [
            '/' => [
                'GET' => 'SupportController@get',
            ],
        ]
    ],
];
```

Instead of generating a tree structure from route definitions, I decided to define the route definitions directly as a tree structure. This approach reduces unnecessary algorithms, potentially improving performance. While this DSL may seem straightforward, I suspect there are reasons why common routing libraries don’t use this format.

The terminal nodes (leaves) of the tree structure correspond to HTTP methods.

Prepare a list of HTTP methods separately. In Golang, these are predefined in `net/http`, which is convenient. For this example, I’ll use PHP.

```php
$methods = [
    'GET',
    'POST',
    // more...
];
```

### Implementation
Implement two functions: one to process the current URL (path) into an array for easier tree traversal, and another to compare the route definitions with the target route array and return the matched path data. Query parameters are not considered in this implementation.

To ensure portability across languages, I’ll minimize the use of built-in functions.

```:php
function createCurrentPathArray($routes) {
    $currentPath = '/users/1'; // Current path

    $currentPathLength = strlen($currentPath);

    $currentPathArray = [];

    for ($i=0; $i < $currentPathLength; $i++) {
        if ($currentPathLength == 1) {
            $currentPathArray[] = '/';
        } else {
            if ($currentPath{$i} == '/') {
                $currentPathArray[] = '/';
                $target = count($currentPathArray) - 1;
            } else {
                $currentPathArray[$target] .= $currentPath{$i};
            }
        }
    }

    return $currentPathArray;
}

// Search
// Compare route definitions with the target route array and return the matched data.
// The search ends when reaching a leaf node.
function urlMatch($routes, $currentPathArray) {
    // TODO Implementation in progress...
}

$currentPathArray = createCurrentPathArray($routes);
$result = urlMatch($routes, $currentPathArray);

var_dump($result); // Should return the matched path data...
```

This concludes Episode 1 as the implementation is still in progress.

## Thoughts
Starting directly with complex structures like Patricia trees or other advanced trees can be overwhelming. While I’ve looked at various implementations for reference, understanding each one is quite challenging. For now, I’ve focused on grasping the algorithm’s concept and working through it step by step. However, lacking mathematical knowledge can be tough.

Although the implementation is incomplete, I feel like I’m starting to see the goal. That said, I’m not confident this will evolve into a base suitable for practical use.

## Postscript
I presented this topic at the Makuake LT Party (an internal LT event).

[speaker-deck - Creating URL Routing: Episode 1](https://speakerdeck.com/bmf_san/urlruteinguwotukuruepisodo1)

### References
- [Algorithm visualization - Radix Tree](https://www.cs.usfca.edu/~galles/visualization/RadixTree.html)
- [github - [Japanese] Patricia Tree](https://github.com/ethereum/wiki/wiki/%5BJapanese%5D-Patricia-Tree)
- [WhiteDog@Blog](http://takao.blogspot.com/2012/03/patriciatrie.html)
- [404 Blog Not Found - algorithm - Patricia Trie (Radix Trie) in JavaScript](http://blog.livedoor.jp/dankogai/archives/51766842.html)
- [http request multiplexer and string matching](https://persol-pt.github.io/posts/tech-workshop1222/)
- [@IT - Choosing data structures can make a big difference (3/3)](http://www.atmarkit.co.jp/ait/articles/0809/01/news163_3.html)
- [Basic Data Structures: Traversing Tree Structures](http://www.sb.ecei.tohoku.ac.jp/lab/wp-content/uploads/2012/11/2012_d12.pdf)
- [pixiv inside - Creating a fast URL routing system in PHP](https://devpixiv.hatenablog.com/entry/2015/12/13/145741)
- [Doing good URL routing in PHP without frameworks](http://noranuk0.hatenablog.com/entry/2018/01/20/114933)
- [gist - neo-nanikaka/CommonPrefixTrieRouter.php](https://gist.github.com/neo-nanikaka/c2e2f7742b311696d50b)
- [github.com - nissy/bon](https://github.com/nissy/bon)
- [github.com - nissy/mux](https://github.com/nissy/mux)
- [github.com - ytakano/radix_tree](https://github.com/ytakano/radix_tree)
- [github.com - kkdai/radix](https://github.com/kkdai/radix)
- [github.com - MarkBaker/Tries](https://github.com/MarkBaker/Tries)
