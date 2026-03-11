---
title: URLルーティングをつくる　エピソード3（完結編）
slug: creating-url-routing-episode-3
date: 2019-03-17T00:00:00Z
author: bmf-san
categories:
  - アルゴリズムとデータ構造
tags:
  - PHP
  - URLルーティング
  - HTTP
  - 木構造
  - router
translation_key: creating-url-routing-episode-3
---


# 概要
[URLルーティングをつくる　エピソード1](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%891) と[URLルーティングをつくる　エピソード2
](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%892)でURLルーティングの自作について試行錯誤の過程を記してきたが、ようやく一段落させることができたので完結編という形で締括くりたい。

完結、といっても課題はいくらでもあるし突き詰めるとこればっかりに時間をかけることができるようなモノであるということは承知している。。。

# 前回までの話し
[エピソード1](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%891) では、ルーティングのデータ構造を考えたり、とりあえず手を動かして実装のイメージを掴もうとした。（動くところまで持っていけなかった。。。）

[エピソード2](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%892)では、データ構造を見直したり、参考になりそうなリポジトリを漁って動く形まで持っていった。

そして今回のエピソード3では、[URLルーティングをつくる　エピソード2
](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%892)でやり残した部分の実装を完了させた。

具体的にいうと、[URLルーティングをつくる　エピソード2
](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%892)までは、ルーティングマップを生成する処理を実装せずに、予めクライアント側でルーティングマップを用意する手を抜いた形でルーティングを実現していたが、今回はその部分の実装をした。

今回の内容はちょうど[phperkaigi2019](https://phperkaigi.jp/2019/)で登壇するのでそのためのまとめ記事という側面もあるので、[エピソード1](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%891) と[URLルーティングをつくる　エピソード2
](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%892)の内容も含んでいる。

プロポーザル↓
[fortee - PHPでURLルーティングをつくる by bmf_san](https://fortee.jp/phperkaigi-2019/proposal/08d951da-29cb-4ee5-bf08-c88129c0bb3f)

スライドはこちら↓
[Speakerdeck ー URLルーティングをつくる](https://speakerdeck.com/bmf_san/urlruteinguwotukuru?slide=43)


**記事の内容はスライドの補足のような感じなのでスライドを見るほうがわかりやすいかもしれない..**

# 今回のソース
リポジトリとパッケージを公開している。

- [github - bmf-san/ahi-router](https://github.com/bmf-san/ahi-router)
- [Packagist - bmf-san/ahi-router](https://packagist.org/packages/bmf-san/ahi-router#2.0.0)

#  URLルーティングとは

**リクエストされたURLに対して、実行したい処理を返すもの**

URLのパス部分をパースして、任意の値を返せるようなロジックが実装できればURLルーティングとしての最低限の機能を満たせるはず。

パス（/foo/bar/1）のパースには正規表現や文字列探索のアルゴリズムを使ったりする。

# PHPのルーティングライブラリ
有名所だとこんな感じ・・？

- [FastRoute](https://github.com/nikic/FastRoute)
- [Pux](https://github.com/c9s/Pux)
- [Klein](https://github.com/klein/klein.php)

etc...

FastRouteは確かSlimで採用されていた気がする。

# 自作してみる
## 前提
とりあえず動くモノにする、というのは最低限の条件として他の言語への移植しやすさを考慮してPHPの標準関数を極力さけるような実装を検討した。（あとでGo書き直したいので...

あとは単純にオレオレアルゴリズムでゼロから考えながら実装してみたかったので純粋なロジックでかくことを前提とした。（なので正規表現も使わない）

## 仕様
ルーティングとしての最低限の条件を満たせるであろう仕様とした。

- 複数のパスパラメーターを含むURLに対応すること
  - /foo/bar/:foo/:bar
  - よくあるルーティングのパターン
- マッチしたルートで、アクションとパラメーター情報を返却できること

## I/O
実装に入る前にI/Oを確認しておく。

Router（ルーティングを行うクラス）がどんなデータを受け取って、どういう形のデータを返すのか整理する意図。

- Input
  - Request URI
    - /foo/bar/1
  - HTTP Method
    - GET/POST/PUT/PATCH/DELETE
  - Routing Map
    - I/Oをマッピングしたデータ。後述。

- Output
  - Action
    - ex. PostController@getPosts
  - Parameter
    - パスパラメーターと値のセット
       - ex. /foo/bar/1
         - /foo/bar/:id
           - → id 1

## データ構造を考える
Routerが内部的に扱うデータ構造を検討する。

内部的に扱うデータ＝Routing Map

Routing Mapというワードは定義のある言葉ではなさそうなので説明しておくと、

**URIと返却したい処理をマッピングしたデータ**のことである。

このパスにリクエストされたらこの処理を行う、というルールをまとめておくもので、Routerは事前に定義されたルート定義からこのRouting Mapを生成して、ルーティングを行う際にこのRouting Mapの探索を行い、処理を返す。

ここでいうルート定義とは、ルーティングの設定ファイル等でアプリケーションが扱うエンドポイントと処理をライブラリのAPIに従って記述している設定のことをルート定義と呼んでいる。

例えばLaravelだったらこういう感じで定義するやつ。

```php
<?php
Route::get('/home', 'HomeController@index);
```

ルート定義はRouting Mapをつくるための情報となる。

このRouting Mapのデータ構造について考える。

ルーティングで探索したい対象であるパスの階層構造に着目して、ルート定義を木構造で表現する。

木構造のアルゴリズムには色々な種類があるが、今回は基数木という木構造を参考にしてみた。（厳密に基数木であるとは言えないかもしれない。そのへんはちゃんと勉強できていない。）

パス部分を木構造で表現し、Leafとなる部分をActionとして扱うことで、木構造を探索した結果がルーティングの返却すべき値（Leafの値）となるような構造にしてみた。

この辺はテキストではわかりづらいので、スライドを参照してもらいたい。

[Speakerdeck ー URLルーティングをつくる](https://speakerdeck.com/bmf_san/urlruteinguwotukuru?slide=43)

木構造を採用したRouting Mapは、PHPでは多次元配列で表現する。

ざっとこんな感じ。

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

データ構造が検討できたらあとは愚直に実装するのみ・・・！

## 実装
ルーティングに関わる処理に責務を持つRouterクラスを実装する。

このRouterクラスに必要な処理は2つで、

- Routing Mapを生成する処理
- Routing Mapから探索する処理

仕様が単純なのでこれだけ。（実装はやや面倒だが・・）

具体的な実装は[github - bmf-san/ahi-router](https://github.com/bmf-san/ahi-router)を参照。

ここでは要所だけ記載する。

Routerを利用するclient側の実装はこんな感じ。

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

### Routing Mapを生成する処理
Path、Method、ActionのデータセットからRouting Mapを更新していく処理を実装する。

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
        // 再帰処理と参照（＆）を駆使してルーティングマップにルーティングを追加していく
    }
```

参照を駆使して多次元配列を動的に生成するようなロジックをかいている。

雑に簡略化した例は下記の通り。

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

### Routing Mapから探索する処理
Path、Medthod、Parameterのデータセットを元にルーティングマップから該当するLeafを探索する処理を実装する。

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
       // ルーティングマップを探索していく処理
    }
```

下記のような処理で愚直に実装している。

極力PHPの標準関数を避けるのを前提しているのでパワープレイとなっている。。。

```php
<?php

$request_uri = '/posts';
$routing_path = '/posts'; // ルーティングマップに定義されたパス

// 以下は説明上簡略している部分がある
for ($i = 0; $i < str_length($routing_path); $i++) {
    if ($request_uri{$i} === $routing_path{$i}) {  // 一文字ずつパスを比較している
        // something to do
    }
}
```

# 実装してみて感じた課題
エラーハンドリングと実行速度の考慮がなされていないのでライブラリとしてはちょっと残念。。。

前者についてはともかく、後者については文字列探索のアルゴリズムの選定が必要なのでやや難易度が高いように思う。（勉強せねば・・・）

今回は単純な機能のみで実装することを目指したが、ネームルート（ルート定義のグルーピング）やパスパラーメーターに正規表現を使えるにしたり、ミドルウェアとのつなぎ込みを実装できたりすると便利なルーティングライブラリとして扱えるんじゃないかと思った。

# 所感
- たったの158行で動いた。
- 再帰処理とパワープレイでいける。
- ライブラリとして成立させるには木構造のアルゴリズム（トライ木、基数木、パトリシア木...etc）を学んで適したものを選定できるようにしたい）

# 追記
ルーティングの処理のパフォーマンスについては、N数となる部分の想定次第で許容できる閾値が変わってくるはずなので、スマートな実装でなくても実用レベルでは耐えうる可能性もあるはず。。。
今回はルーティング数やパラーメーター情報が増えると計算量が比例して増えていくようなアルゴリズムになってしまっている。
あとは木構造でなくとも正規表現で実装されているライブラリもあるので木構造がスタンダートな実装というわけでもなさそうではある。
「推測するな、計測せよ」という言葉に従ってベンチマークをとったほうが良い。（今回はサボったが...）

# 追記2
良い記事を見つけたのでメモ。

[Hatena Developpers Blog - 文字列アルゴリズムの学びかた](https://developer.hatenastaff.com/entry/2016/12/22/210006)


