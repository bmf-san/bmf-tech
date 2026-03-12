---
title: URLルーティングをつくる　エピソード2
description: URLルーティングをつくる　エピソード2について、基本的な概念から実践的な知見まで詳しく解説します。
slug: creating-url-routing-episode-2
date: 2019-01-06T00:00:00Z
author: bmf-san
categories:
  - アルゴリズムとデータ構造
tags:
  - HTTP
  - URLルーティング
  - 木構造
  - router
translation_key: creating-url-routing-episode-2
---


# 概要
[URLルーティングをつくる　エピソード1](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%891)の続き。

とりあえず動く形のものを仕上げて[packagist - ahi-router](https://packagist.org/packages/bmf-san/ahi-router#v1.0)という名前でパッケージ公開した。

# エピソード1からの変更点
エピソード1では、データ構造に木構造を採用してルーティングを作ろうというと試みた。

パフォーマンスが考慮されているライブラリでは、木構造を生成するロジックを用意して、最適化された探索アルゴリズムを実装するような形になっているようだが、木構造を生成するロジックをかくのはめん(ry 時間がかかりそうだったので、探索部分だけ頑張る方向性でやってみることにした。

前回はルーティング定義のデータ構造を、

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

としていたが、

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

こんな感じに定義し直した。

変更点としては、

- ルートが2つある構造になっていたので、一つに統一して木構造として成立するようにした。
    - **根ノード (英: root node) とは、親ノードを持たないノードのこと。根ノードは木構造の最上位にあるノードであり、1つの木構造に高々1つしか存在しない** - [Wikipedia - 木構造（データ構造）](https://ja.wikipedia.org/wiki/%E6%9C%A8%E6%A7%8B%E9%80%A0_(%E3%83%87%E3%83%BC%E3%82%BF%E6%A7%8B%E9%80%A0))より引用
    - つまり、前回のやつは正確には木構造ではなく、木構造モドキだった
- END_POINTという識別子を用意した
    - END_POINTという名前が適切だとは思えないが、ルート（根ノード）と区別を明確につけるために用意することにした

前回は関数で頑張ろうとしたが色々辛かったのでオブジェクトで戦うことにしたらすんなり実装できた。
データ構造を変更したのも実装のしやすさに影響を与えたと思う。

# 実装
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
                // ルートの時
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

// こんな感じに使う
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

計算量はざっくりO(n)なっているので、n（ルート定義）が増えるほど計算量は比例して増えていく残念アルゴリズム。

# 所感
ちゃんとつくるならやっぱり木構造の探索のアルゴリズムはかじっておくべきだろう。察していたが反省した。

アルゴリズムの重要さが身に沁みたような気がする。（小並感）

日頃こんなにグルグルとしたコードを書かないので頭の体操にはなった。（不定期でこういう体操をしてアルゴリズムに慣れていくのは良いと思った）

割とメジャーなルーティングライブラリでも、正規表現を使用していたり、最適化されていないアルゴリズムで実装されていたりするっぽいので今後も色んな実装に目を通したり、アルゴリズムの勉強をしたりしてそのうちルーティングの実装に再挑戦してみたい。

# ソースとパッケージ
- [github - bmf-san/ahi-router](https://github.com/bmf-san/ahi-router)
- [packagist - ahi-router](https://packagist.org/packages/bmf-san/ahi-router#v1.0)
    - 雑だがパッケージ化しておいた

# 参考
- [pixiv inside - PHPで高速に動作するURLルーティングを自作してみた](https://devpixiv.hatenablog.com/entry/2015/12/13/145741)
- [github - devlibs/routing](https://github.com/devlibs/routing)
