---
title: Creating URL Routing Episode 3 (Final Episode)
slug: creating-url-routing-episode-3
image: /assets/images/posts/url/70861219-30929d80-1f6e-11ea-8e86-114e8ba0942b.png
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
translation_key: creating-url-routing-episode-3
---

# Overview
I have documented the trial and error process of creating URL routing in [Creating URL Routing Episode 1](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%891) and [Creating URL Routing Episode 2](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%892), and I can finally wrap it up, so I want to conclude it in this final episode.

That said, there are always challenges, and I understand that if I delve deeper, I could spend a lot of time on this alone...

# Previous Discussions
In [Episode 1](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%891), I thought about the data structure for routing and tried to implement it to get a grasp of the concept. (I couldn't get it to work...)

In [Episode 2](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%892), I reviewed the data structure and searched for potentially useful repositories to bring it to a working state.

In this Episode 3, I completed the implementation of the parts I left unfinished in [Creating URL Routing Episode 2](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%892).

Specifically, in [Episode 2](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%892), I implemented routing without generating a routing map, relying on the client side to prepare the routing map in advance. This time, I implemented that part.

The content of this article also serves as a summary for my presentation at [phperkaigi2019](https://phperkaigi.jp/2019/), so it includes the contents of [Episode 1](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%891) and [Creating URL Routing Episode 2](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%892).

Proposal↓
[fortee - Creating URL Routing in PHP by bmf_san](https://fortee.jp/phperkaigi-2019/proposal/08d951da-29cb-4ee5-bf08-c88129c0bb3f)

Slides here↓
[Speakerdeck - Creating URL Routing](https://speakerdeck.com/bmf_san/urlruteinguwotukuru?slide=43)

**The content of the article may be easier to understand by looking at the slides as they serve as supplementary material.**

# Source Code
I have published the repository and package.

- [github - bmf-san/ahi-router](https://github.com/bmf-san/ahi-router)
- [Packagist - bmf-san/ahi-router](https://packagist.org/packages/bmf-san/ahi-router#2.0.0)

# What is URL Routing?
**It returns the process to execute for the requested URL.**

If you can implement logic that parses the path part of the URL and returns arbitrary values, you should meet the minimum functionality of URL routing.

For parsing paths (e.g., /foo/bar/1), you can use regular expressions or string search algorithms.

# PHP Routing Libraries
Here are some well-known ones:

- [FastRoute](https://github.com/nikic/FastRoute)
- [Pux](https://github.com/c9s/Pux)
- [Klein](https://github.com/klein/klein.php)

etc...

I believe FastRoute was adopted in Slim.

# Trying to Create My Own
## Prerequisites
As a minimum requirement, I considered ease of porting to other languages and aimed to avoid using PHP's standard functions as much as possible. (I want to rewrite it in Go later...) 

Also, I simply wanted to implement it from scratch with my own algorithm, so I aimed to write it with pure logic. (Thus, I won't use regular expressions.)

## Specifications
I set the specifications to meet the minimum conditions for routing.

- Support URLs containing multiple path parameters:
  - /foo/bar/:foo/:bar
  - Common routing patterns
- Return action and parameter information for matched routes.

## I/O
Before implementation, I will confirm the I/O.

The intention is to organize what kind of data the Router (the class that performs routing) receives and what form of data it returns.

- Input
  - Request URI
    - /foo/bar/1
  - HTTP Method
    - GET/POST/PUT/PATCH/DELETE
  - Routing Map
    - Data mapping I/O. To be discussed later.

- Output
  - Action
    - ex. PostController@getPosts
  - Parameter
    - Set of path parameters and values
       - ex. /foo/bar/1
         - /foo/bar/:id
           - → id 1

## Considering Data Structure
I will consider the data structure that the Router handles internally.

The data handled internally = Routing Map.

The term Routing Map does not seem to have a defined meaning, so let me explain:

**It is data that maps the URI to the desired process.**

It summarizes the rules that say, "If this path is requested, perform this process," and the Router generates this Routing Map from predefined route definitions and searches this Routing Map when performing routing to return the process.

The route definitions here refer to the settings written in the routing configuration files that describe the endpoints and processes handled by the application according to the library's API.

For example, in Laravel, it would be defined like this:

```php
<?php
Route::get('/home', 'HomeController@index);
```

The route definitions serve as information for creating the Routing Map.

Now, let's think about the data structure of this Routing Map.

Focusing on the hierarchical structure of the paths to be explored in routing, I will represent the route definitions as a tree structure.

There are various types of tree algorithms, but this time I referred to a radix tree (though it may not strictly be a radix tree; I haven't studied that thoroughly).

By representing the path part as a tree structure and treating the leaf nodes as actions, the result of exploring the tree structure will be the values that should be returned by routing (the values of the leaves).

This part may be difficult to understand through text, so I would like you to refer to the slides.

[Speakerdeck - Creating URL Routing](https://speakerdeck.com/bmf_san/urlruteinguwotukuru?slide=43)

The Routing Map using a tree structure is represented as a multidimensional array in PHP.

Here’s a rough example:

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

Once the data structure is considered, all that’s left is to implement it straightforwardly...

## Implementation
I will implement the Router class, which is responsible for the processes related to routing.

This Router class needs two processes:

- Process to generate the Routing Map
- Process to search from the Routing Map

The specifications are simple, so that’s all. (The implementation is somewhat tedious...)

For specific implementation, please refer to [github - bmf-san/ahi-router](https://github.com/bmf-san/ahi-router).

Here, I will only describe the key points.

The implementation on the client side using the Router looks like this:

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

### Process to Generate the Routing Map
Implement the process to update the Routing Map from the dataset of Path, Method, and Action.

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
        // Use recursion and references to add routing to the routing map
    }
```

I am writing logic to dynamically generate the multidimensional array using references.

A simplified example is as follows:

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

### Process to Search from the Routing Map
Implement the process to explore the routing map based on the dataset of Path, Method, and Parameter.

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
       // Process to explore the routing map
    }
```

I am implementing it straightforwardly with the following logic.

Since I am trying to avoid using PHP's standard functions as much as possible, it has become a bit of a power play...

```php
<?php

$request_uri = '/posts';
$routing_path = '/posts'; // Path defined in the routing map

// The following part is simplified for explanation
for ($i = 0; $i < str_length($routing_path); $i++) {
    if ($request_uri{$i} === $routing_path{$i}) {  // Comparing the paths character by character
        // something to do
    }
}
```

# Challenges Faced During Implementation
Error handling and execution speed considerations have not been addressed, making it somewhat disappointing as a library...

Regarding the former, the latter requires selecting string search algorithms, which seems to be a bit challenging. (I need to study...)

This time, I aimed to implement only simple functionality, but I think it would be convenient to handle named routes (grouping of route definitions) and use regular expressions for path parameters, as well as implement middleware integration, making it a more useful routing library.

# Thoughts
- It worked with just 158 lines.
- Recursion and power play can work.
- To establish it as a library, I want to learn about tree structure algorithms (trie, radix tree, Patricia tree...etc) and be able to select the appropriate one.

# Addendum
The performance of routing processes may vary depending on the expected number of elements, so even if the implementation is not smart, it might still be practical.
This time, the algorithm has become one where the computational complexity increases proportionally as the number of routes and parameter information increases.
Also, there are libraries implemented with regular expressions instead of tree structures, so it doesn't seem that tree structures are the standard implementation.
It would be better to follow the saying, "Don't guess, measure," and take benchmarks. (I slacked off this time...)

# Addendum 2
I found a good article, so here’s a note.

[Hatena Developers Blog - How to Learn String Algorithms](https://developer.hatenastaff.com/entry/2016/12/22/210006)