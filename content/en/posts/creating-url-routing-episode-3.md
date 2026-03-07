---
title: 'Creating URL Routing: Episode 3 (Final Chapter)'
slug: creating-url-routing-episode-3
date: 2019-03-17T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - PHP
  - URL Routing
  - HTTP
  - Tree Structure
  - Router
description: The final chapter in the series on building a custom URL routing system.
translation_key: creating-url-routing-episode-3
---

# Overview

In [Episode 1](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%891) and [Episode 2](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%892), I documented the trial-and-error process of creating a custom URL routing system. Finally, I’ve reached a point where I can wrap things up, so I’d like to conclude this series with the final chapter.

That said, there are still plenty of challenges and areas to refine. If I were to pursue this further, I could easily spend an endless amount of time on it...

# Recap of Previous Episodes

In [Episode 1](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%891), I explored the data structure for routing and tried to get a feel for the implementation by writing some code (though I didn’t get it fully functional).

In [Episode 2](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%892), I revisited the data structure, explored some relevant repositories, and managed to create a working implementation.

In this final episode, I completed the parts that were left unfinished in [Episode 2](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%892).

Specifically, in [Episode 2](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%892), I skipped implementing the process to generate the routing map and instead relied on pre-defined routing maps on the client side. This time, I implemented that part.

This article also serves as a summary for my presentation at [phperkaigi2019](https://phperkaigi.jp/2019/), so it includes content from [Episode 1](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%891) and [Episode 2](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%892).

Proposal: [fortee - Creating URL Routing in PHP by bmf_san](https://fortee.jp/phperkaigi-2019/proposal/08d951da-29cb-4ee5-bf08-c88129c0bb3f)

Slides: [Speakerdeck - Creating URL Routing](https://speakerdeck.com/bmf_san/urlruteinguwotukuru?slide=43)

**This article serves as a supplement to the slides, so it might be easier to understand by referring to the slides.**

# Source Code
The repository and package are publicly available:

- [GitHub - bmf-san/ahi-router](https://github.com/bmf-san/ahi-router)
- [Packagist - bmf-san/ahi-router](https://packagist.org/packages/bmf-san/ahi-router#2.0.0)

# What is URL Routing?

**It’s a mechanism that returns the desired process for a requested URL.**

By parsing the path part of a URL, you can implement logic that returns arbitrary values, fulfilling the minimum requirements for URL routing.

For parsing paths (e.g., `/foo/bar/1`), you can use regular expressions or string search algorithms.

# PHP Routing Libraries
Some well-known libraries include:

- [FastRoute](https://github.com/nikic/FastRoute)
- [Pux](https://github.com/c9s/Pux)
- [Klein](https://github.com/klein/klein.php)

FastRoute, for instance, is used in Slim.

# Building My Own

## Prerequisites

The minimum requirement was to create something functional. Additionally, I wanted to make it easier to port to other languages, so I avoided PHP-specific functions as much as possible (I plan to rewrite it in Go later).

I also wanted to design and implement my own algorithm from scratch, so I focused on writing pure logic (avoiding even regular expressions).

## Specifications

I aimed to meet the minimum requirements for routing:

- Support URLs with multiple path parameters:
  - `/foo/bar/:foo/:bar`
  - A common routing pattern
- Return the action and parameter information for matched routes

## Input/Output

Before implementation, I clarified the input and output of the Router class.

### Input
- **Request URI**: `/foo/bar/1`
- **HTTP Method**: `GET/POST/PUT/PATCH/DELETE`
- **Routing Map**: Data mapping inputs to outputs (explained below)

### Output
- **Action**: e.g., `PostController@getPosts`
- **Parameter**: Path parameters and their values
  - e.g., `/foo/bar/1`
    - `/foo/bar/:id` → `id: 1`

## Designing the Data Structure

The data structure managed internally by the Router is the Routing Map.

Routing Map refers to **data mapping URIs to the desired actions**. It defines the rules for which process to execute when a specific path is requested. The Router generates this Routing Map from predefined route definitions and uses it to perform routing by searching for the appropriate process.

For example, in Laravel, route definitions look like this:

```php
<?php
Route::get('/home', 'HomeController@index');
```

The route definition serves as the information to create the Routing Map.

To design the Routing Map, I focused on the hierarchical structure of paths in routing and represented the route definitions as a tree structure. I referred to a radix tree for this structure (though it might not strictly qualify as a radix tree).

The path is represented as a tree structure, with the leaves representing actions. The result of traversing the tree is the value (action) that the routing should return.

This might be hard to grasp in text, so please refer to the slides:

[Speakerdeck - Creating URL Routing](https://speakerdeck.com/bmf_san/urlruteinguwotukuru?slide=43)

In PHP, the tree-structured Routing Map is represented as a multidimensional array, like this:

```php
<?php
$routeMap = [
    '/' => [
        'END_POINT' => [
            'GET' => 'IndexController@index',
        ],
        'posts' => [
            'END_POINT' => [
                'GET' => 'PostController@getPosts',
            ],
            ':id' => [
                'END_POINT' => [
                    'GET' => 'PostController@edit',
                    'POST' => 'PostController@update',
                ],
                ':token' =>  [
                    'END_POINT' => [
                        'GET' => 'PostController@preview',
                    ],
                ],
            ],
            ':category' => [
                'END_POINT' => [
                    'GET' => 'PostController@getPostsByCategory',
                ],
            ],
        ],
        'profile' => [
            'END_POINT' => [
                'GET' => 'ProfileController@getProfile',
            ],
        ],
    ],
];
```

Once the data structure is designed, the next step is to implement it.

## Implementation

The Router class, responsible for routing, needs to handle two main processes:

- Generating the Routing Map
- Searching the Routing Map

The specifications are simple, but the implementation can be a bit tricky.

Here’s an example of how the client-side implementation might look:

```php
<?php

require_once("../src/Router.php");

$router = new bmfsan\AhiRouter\Router();

$router->add('/', [
    'GET' => 'IndexController@index',
]);

$router->add('/posts', [
    'GET' => 'PostController@getPosts',
]);

$router->add('/posts/:id', [
    'GET' => 'PostController@edit',
    'POST' => 'PostController@update',
]);

$router->add('/posts/:id/:token', [
    'GET' => 'PostController@preview',
]);

$router->add('/posts/:category', [
    'GET' => 'PostController@getPostsByCategory',
]);

$router->add('/profile', [
    'GET' => 'ProfileController@getProfile',
]);

$result = $router->search('/posts/1/token', 'GET', [':id', ':token']);

var_dump($result);
// array(2) {
//     'action' =>
//     string(22) "PostController@preview"
//     'params' =>
//     array(2) {
//         ':id' =>
//         string(1) "1"
//         ':token' =>
//         string(5) "token"
//     }
// }
```

### Generating the Routing Map

This process updates the Routing Map with a dataset of Path, Method, and Action.

```php
/**
 * Add routing to route map
 *
 * @param string $route
 * @param array $handler
 * @return void
 */
public function add($route, $handler)
{
    // Use recursion and references to dynamically generate the routing map
}
```

Here’s a simplified example:

```php
<?php

$routeMap = [
    '/'
];

$ref = &$routeMap['/'];

$ref = [
    '/posts' => [
        'END_POINT' => [
            'GET' => 'PostController@getPosts'
        ]
    ]
];

var_dump($routeMap);

// array(2) {
//  [0] =>
//  string( 1) "/"
//  '/' =>
//  array( 1) {
//    '/posts' =>
//    array( 1) {
//      'END_POINT' =>
//       array( 1) {
//         'GET' =>
//         string( 23) "PostController@getPosts"
//       }
//    }
//  }
// } 
```

### Searching the Routing Map

This process searches the Routing Map for the corresponding leaf based on the dataset of Path, Method, and Parameters.

```php
/**
 * Search a path and return action and parameters
 *
 * @param  string $requestUri
 * @param  string $requestMethod
 * @param  array  $targetParams
 * @return array
 */
public function search($requestUri, $requestMethod, $targetParams = []): array
{
   // Implement logic to search the routing map
}
```

Here’s an example of the logic:

```php
<?php

$request_uri = '/posts';
$routing_path = '/posts'; // Defined path in the routing map

// Simplified for explanation purposes
for ($i = 0; $i < str_length($routing_path); $i++) {
    if ($request_uri{$i} === $routing_path{$i}) {  // Compare paths character by character
        // something to do
    }
}
```

# Challenges Encountered

- **Error Handling and Performance**: The library lacks proper error handling and performance optimization, which makes it less robust as a library.
- **Algorithm Selection**: Choosing the right string search algorithm for better performance is a challenge that requires further study.

Although the goal was to implement only basic functionality, adding features like named routes (grouping route definitions), support for regular expressions in path parameters, and middleware integration would make it a more useful routing library.

# Reflections

- It worked with just 158 lines of code.
- Recursive processing and brute force can work.
- To make it a proper library, I need to study tree structure algorithms (trie, radix tree, Patricia tree, etc.) and choose the most suitable one.

# Additional Notes

The performance of routing depends on the expected number of routes, so even a less elegant implementation might be sufficient for practical use. However, the current algorithm’s complexity increases proportionally with the number of routes and parameters, which could become a bottleneck.

It’s worth noting that some libraries implement routing using regular expressions instead of tree structures, so a tree structure isn’t necessarily the standard approach. Following the principle of "don’t guess, measure," it would be better to benchmark the implementation (though I skipped this step for now).

# Additional Resources

I found a great article worth noting:

[Hatena Developers Blog - How to Learn String Algorithms](https://developer.hatenastaff.com/entry/2016/12/22/210006)
