---
title: 'Creating URL Routing: Episode 2'
slug: creating-url-routing-episode-2
date: 2019-01-06T00:00:00Z
author: bmf-san
categories:
  - Algorithms
  - Data Structures
tags:
  - HTTP
  - URL Routing
  - Tree Structure
  - Router
description: Continuing from Episode 1, this post discusses updates and improvements made to the URL routing implementation, including changes to the data structure and exploration of tree structures.
translation_key: creating-url-routing-episode-2
---

# Overview
[Creating URL Routing: Episode 1](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%891) continuation.

I managed to create a working version and published it as a package named [packagist - ahi-router](https://packagist.org/packages/bmf-san/ahi-router#v1.0).

# Changes from Episode 1
In Episode 1, I attempted to create routing by adopting a tree structure for the data structure.

In libraries optimized for performance, it seems common to prepare logic for generating tree structures and implement optimized search algorithms. However, writing the logic to generate a tree structure seemed time-consuming, so I decided to focus on improving the search part instead.

Previously, the routing definition data structure was:

```php
<?php
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

But I redefined it as:

```php
<?php

$routes = [
    '/' => [
        'END_POINT' => [
            'GET' => 'IndexController@getIndex',
        ],
        'posts' => [
            'END_POINT' => [
                'GET' => 'PostController@getPosts',
            ],
            ':title' => [
                'END_POINT' => [
                    'GET' => 'PostController@getPostByPostTitle',
                    'POST' => 'PostController@postPostByPostTitle',
                ],
                ':token' =>  [
                    'END_POINT' => [
                        'GET' => 'PostController@getPostByToken',
                    ],
                ],
            ],
            ':category_name' => [
                'END_POINT' => [
                    'GET' => 'PostController@getPostsByCategoryName',
                ],
            ],
        ],
    ],
];
```

The changes include:

- Unified the structure into a single root to form a proper tree structure.
    - **A root node is a node without a parent node. It is the topmost node in a tree structure, and there can be at most one root node in a tree structure.** - [Wikipedia - Tree Structure (Data Structure)](https://ja.wikipedia.org/wiki/%E6%9C%A8%E6%A7%8B%E9%80%A0_(%E3%83%87%E3%83%BC%E3%82%BF%E6%A7%8B%E9%80%A0))
    - In other words, the previous version was not technically a tree structure but a pseudo-tree structure.
- Introduced an identifier called `END_POINT`.
    - While the name `END_POINT` may not be ideal, it was introduced to clearly distinguish it from the root node.

Previously, I tried to manage everything with functions, but it was challenging. Switching to an object-oriented approach made the implementation smoother. The changes to the data structure also contributed to easier implementation.

# Implementation
```php
<?php

namespace bmfsan\AhiRouter;

class Router
{
    /**
     * Path parameters
     * @var array
     */
    private $params = [];

    /**
     * Create array for search path from current path
     *
     * @param  string $currentPath
     * @return array
     */
    public function createArrayFromCurrentPath($currentPath): array
    {
        $currentPathLength = strlen($currentPath);

        $arrayFromCurrentPath = [];

        for ($i=0; $i < $currentPathLength; $i++) {
            if ($currentPathLength == 1) {
                // Root case
                if ($currentPath{$i} == '/') {
                    $arrayFromCurrentPath[] = '/';
                }
            } else {
                if ($currentPath{$i} == '/') {
                    $arrayFromCurrentPath[] = '';
                    $target = count($arrayFromCurrentPath) - 1;
                } else {
                    $arrayFromCurrentPath[$target] .= $currentPath{$i};
                }
            }
        }

        return $arrayFromCurrentPath;
    }

    /**
     * Search a path and return action and parameters
     *
     * @param  array $routes
     * @param  array $arrayFromCurrentPath
     * @param  string $requestMethod
     * @param  array  $targetParams
     * @return array
     */
    public function search($routes, $arrayFromCurrentPath, $requestMethod, $targetParams = []): array
    {
        $i = 0;
        while ($i < count($arrayFromCurrentPath)) {
            if ($i == 0) {
                $targetArrayDimension = $routes['/'];
            }

            // Condition for root
            if ($arrayFromCurrentPath[$i] == '/') {
                $result = $targetArrayDimension['END_POINT'];
                break;
            }

            foreach ($targetArrayDimension as $key => $value) {
                if (isset($arrayFromCurrentPath[$i])) {
                    if (isset($targetArrayDimension[$arrayFromCurrentPath[$i]])) {
                        $targetArrayDimension = $targetArrayDimension[$arrayFromCurrentPath[$i]];
                    } else {
                        // Condition for parameters
                        $targetArrayDimension = $this->createParams($targetParams, $targetArrayDimension, $arrayFromCurrentPath[$i]);
                    }
                }

                // Condition for last loop
                if ($i == count($arrayFromCurrentPath) - 1) {
                    $result = $targetArrayDimension['END_POINT'];
                }

                $i++;
            }
        }

        return [
            'action' => $result[$requestMethod],
            'params' => $this->params,
        ];
    }

    /**
     * Create parameter data
     *
     * @param  array $targetParams
     * @param  array $targetArrayDimension
     * @param  string $targetPath
     * @return array
     */
    private function createParams($targetParams, $targetArrayDimension, $targetPath)
    {
        for ($i=0; $i < count($targetParams); $i++) {
            if (isset($targetArrayDimension[$targetParams[$i]])) {
                $this->params[$targetParams[$i]] = $targetPath;
                
                return $targetArrayDimension[$targetParams[$i]];
            }
        }
    }
}

// Example usage
$currentPath = '/posts/1/abc123!@#';
$currentMethod = 'GET';
$currentParams = [
    ':title',
    ':token',
];
$router = new Router();
$currentPathArray = $router->createArrayFromCurrentPath($currentPath);
$router->search($routes, $currentPathArray, $currentMethod, $currentParams);
```

The time complexity is roughly O(n), so as `n` (route definitions) increases, the computational complexity increases proportionally, making it a suboptimal algorithm.

# Thoughts
If I were to build this properly, I should definitely study tree traversal algorithms. I had a hunch about this, but I’ve learned my lesson.

I feel like I’ve come to appreciate the importance of algorithms. (Just my two cents.)

I don’t usually write such convoluted code, so this was a good mental exercise. (I think doing such exercises occasionally to get used to algorithms is a good idea.)

Even some major routing libraries seem to use regular expressions or non-optimized algorithms, so I’d like to continue studying various implementations and algorithms and eventually try implementing routing again.

# Source Code and Package
- [github - bmf-san/ahi-router](https://github.com/bmf-san/ahi-router)
- [packagist - ahi-router](https://packagist.org/packages/bmf-san/ahi-router#v1.0)
    - It’s a bit rough, but I packaged it.

# References
- [pixiv inside - Creating a High-Performance URL Router in PHP](https://devpixiv.hatenablog.com/entry/2015/12/13/145741)
- [github - devlibs/routing](https://github.com/devlibs/routing)