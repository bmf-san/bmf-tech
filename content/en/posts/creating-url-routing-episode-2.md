---
title: Creating URL Routing Episode 2
description: "Implement optimized URL routing with refined tree structure patterns, endpoint mapping, and scalable web application routing logic."
slug: creating-url-routing-episode-2
date: 2019-01-06T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - HTTP
  - URL Routing
  - Tree Structure
  - Router
translation_key: creating-url-routing-episode-2
---

# Overview
Continuing from [Creating URL Routing Episode 1](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%891).

I finished a working version and published it as a package named [packagist - ahi-router](https://packagist.org/packages/bmf-san/ahi-router#v1.0).

# Changes from Episode 1
In Episode 1, I attempted to create routing using a tree structure for the data structure.

While libraries that consider performance seem to prepare logic to generate tree structures and implement optimized search algorithms, writing the logic to generate the tree structure seemed to take too much time, so I decided to focus on just the search part.

Previously, the data structure for routing definitions was defined as follows:

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

However, I redefined it as follows:

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

- The structure had two roots, so I unified it to form a valid tree structure.
    - **A root node is a node that has no parent node. The root node is the topmost node in a tree structure and can exist only once in a single tree structure.** - Quoted from [Wikipedia - Tree Structure (Data Structure)](https://ja.wikipedia.org/wiki/%E6%9C%A8%E6%A7%8B%E9%80%A0_(%E3%83%87%E3%83%BC%E3%82%BF%E6%A7%8B%E9%80%A0))
    - In other words, the previous version was not exactly a tree structure, but rather a tree-like structure.
- An identifier called END_POINT was introduced.
    - Although I don't think the name END_POINT is appropriate, I decided to use it to clearly distinguish it from the root node.

In the previous attempt, I struggled with functions, but switching to an object-oriented approach made implementation much smoother. I believe changing the data structure also contributed to the ease of implementation.

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
                // When at root
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

The time complexity is roughly O(n), so as n (the number of route definitions) increases, the computational complexity increases proportionally, which is unfortunate for the algorithm.

# Thoughts
If I want to do it properly, I should definitely study tree structure search algorithms. I had a feeling about this and I regret it.

I feel like I have come to understand the importance of algorithms more deeply. (Just a casual thought)

Since I don't usually write such convoluted code, it was a good mental exercise. (I think it's good to do such exercises irregularly to get accustomed to algorithms)

Even relatively major routing libraries seem to use regular expressions or implement unoptimized algorithms, so I want to continue looking at various implementations and studying algorithms, and eventually challenge routing implementation again.

# Source and Package
- [github - bmf-san/ahi-router](https://github.com/bmf-san/ahi-router)
- [packagist - ahi-router](https://packagist.org/packages/bmf-san/ahi-router#v1.0)
    - It's rough, but I've packaged it.

# References
- [pixiv inside - I tried creating a fast URL routing in PHP](https://devpixiv.hatenablog.com/entry/2015/12/13/145741)
- [github - devlibs/routing](https://github.com/devlibs/routing)