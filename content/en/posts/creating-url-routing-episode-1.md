---
title: Creating URL Routing Episode 1
description: A step-by-step guide on Creating URL Routing Episode 1, with practical examples and configuration tips.
slug: creating-url-routing-episode-1
date: 2018-12-19T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - HTTP
  - URL Routing
  - Tree Structure
  - Router
translation_key: creating-url-routing-episode-1
---

# Creating URL Routing Episode 1

## Overview
Previously, I created a very basic routing system using React (cf. [Creating a Custom Router with React and History API](https://bmf-tech.com/posts/React%E3%81%A8History%20API%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6router%E3%82%92%E8%87%AA%E4%BD%9C%E3%81%99%E3%82%8B)), but I wanted to challenge myself to create a more proper routing system. The trigger for this was my recent experience with Golang. It seems that by utilizing the standard library in Golang, applications can be implemented quite thinly, but the routing aspect often lacks power in the standard library, leading to a reliance on external libraries. Because of this, I felt that being able to create my own routing system would expand my capabilities in Golang and beyond, so I decided to take the plunge.

## What Does URL Routing Do?
It determines what processing should be executed in response to a requested URL. If necessary, it allows handling of path parameters and query parameters during the execution of the processing.

## Implementation Patterns for URL Routing
There are roughly two patterns:

- A pattern that matches URLs using regular expressions.
- A pattern that uses a tree structure for string searching.

While the impact of routing on application execution speed may not be significant, it is always better to be as fast as possible. Regardless of the language, it should be implemented with optimized algorithms for memory usage and computational complexity.

This time, I will choose the pattern implemented with a tree structure. Although I haven't measured performance, I feel that using a tree structure algorithm is likely to be more efficient than regular expressions, so I will go with that. In fact, there are many libraries implemented with tree structures.

## What is a Tree Structure?
A data structure that has a tree structure defined in the field of graph theory in mathematics. A tree defined in graph theory consists of multiple points (nodes or vertices) and multiple edges.

```
   ○ ・・・Root
 / | ・・・Edge
◯  ◯ ・・・Node
    \
  　  ◯
       \
    　   ○ ・・・Leaf
```

There are various types of tree structures depending on the properties of the nodes and the height of the tree, but I will omit those details here.

## Examples of Tree Structures
- Family Trees
- File Systems
- Domain Names
  - cf. https://www.nic.ad.jp/ja/dom/system.html
- Syntax Trees
  - Compilers, etc.
- DOM Trees
- Hierarchical structures for tags or categories

## Creating URL Routing
What will we treat as a tree structure? Of course, we will treat the list of route definitions as a tree structure.

To roughly explain the implementation flow, when given the route definitions and the current URL (path) as input, we will generate a tree structure from the route definitions, explore the tree structure targeting the current URL (path), and return the matched data.

When handling the tree structure, there may be cases where we implement processes such as adding or deleting nodes, but for URL routing, we will not implement those for now.

### Deciding on the Data Structure
We will first decide on the DSL for the routing definitions. Many libraries provide a simple DSL, but this time we will define a slightly complex DSL with multiple layers.

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

I mentioned earlier that we would generate a tree structure from the route definitions, but we will define the route definitions in a way that they are already structured as a tree. The reason for this approach is simply that writing an algorithm to generate a tree structure seems cumbersome, but conversely, it might reduce unnecessary algorithms and improve performance. I believe the route definitions are not particularly difficult to understand, but I think there must be a reason why common routing libraries do not adopt this structure.

The terminal nodes of the tree structure (leaves) correspond to the HTTP methods.

In addition to the tree structure, we will prepare a list of HTTP methods. In Golang, they are already defined in net/http, which is convenient. This time, we will do it in PHP...

```php
$methods = [
    'GET',
    'POST',
    // more...
];
```

### Implementation
We will implement two functions: one to process the current URL (path) into an array for easier exploration of the tree structure, and another that takes the array and the route definitions array as arguments to return the matched path data. Note that we are not particularly considering query parameters this time.

As an implementation policy, we will avoid using built-in functions as much as possible to consider portability to other languages.

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

// Exploration
// Compare the route definitions and the target route array to return the corresponding data.
// End exploration when reaching a leaf.
function urlMatch($routes, $currentPathArray) {
    // TODO Implementing...
}

$currentPathArray = createCurrentPathArray($routes);
$result = urlMatch($routes, $currentPathArray);

var_dump($result); // Should return the matched path data...
```

As this is still a work in progress, Episode 1 comes to a close here.

## Thoughts
If you try to start with Patricia trees or other complex structures from the beginning, you might get burned badly. I looked at various implementations that might be helpful, but understanding each one is quite hard, so I started by grasping the image of the algorithm and working with my hands. However, it can be tough if you lack mathematical background. Although it is still in progress, I feel like I can see the goal somewhat. However, I am not confident that I can bring it to a level that can be used in actual operations.

## Postscript
I gave a talk at the Makuake LT Party (internal LT competition).

[speaker-deck - Creating URL Routing Episode 1](https://speakerdeck.com/bmf_san/urlruteinguwotukuruepisodo1)

### References
- [Algorithm visualization - Radix Tree](https://www.cs.usfca.edu/~galles/visualization/RadixTree.html)
- [github - [Japanese] Patricia Tree](https://github.com/ethereum/wiki/wiki/%5BJapanese%5D-Patricia-Tree)
- [WhiteDog@Blog](http://takao.blogspot.com/2012/03/patriciatrie.html)
- [404 Blog Not Found - algorithm - Patricia Trie (Radix Trie) in JavaScript](http://blog.livedoor.jp/dankogai/archives/51766842.html)
- [http request multiplexer and string matching](https://persol-pt.github.io/posts/tech-workshop1222/)
- [@IT - The difference between heaven and hell depends on the choice of data structure (3/3)](http://www.atmarkit.co.jp/ait/articles/0809/01/news163_3.html)
- [Basic Data Structures: How to Traverse Tree Structures](http://www.sb.ecei.tohoku.ac.jp/lab/wp-content/uploads/2012/11/2012_d12.pdf)
- [pixiv inside - Trying to create a fast URL routing system in PHP](https://devpixiv.hatenablog.com/entry/2015/12/13/145741)
- [Doing URL routing nicely in PHP without using a framework](http://noranuk0.hatenablog.com/entry/2018/01/20/114933)
- [gist - neo-nanikaka/CommonPrefixTrieRouter.php](https://gist.github.com/neo-nanikaka/c2e2f7742b311696d50b)
- [github.com - nissy/bon](https://github.com/nissy/bon)
- [github.com - nissy/mux](https://github.com/nissy/mux)
- [github.com - ytakano/radix_tree](https://github.com/ytakano/radix_tree)
- [github.com - kkdai/radix](https://github.com/kkdai/radix)
- [github.com - MarkBaker/Tries](https://github.com/MarkBaker/Tries)