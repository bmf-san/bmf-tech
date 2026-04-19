---
title: URLルーティングをつくる　エピソード1
description: URLルーティングをつくる　エピソード1
slug: creating-url-routing-episode-1
date: 2018-12-19T00:00:00Z
author: bmf-san
categories:
  - アルゴリズムとデータ構造
tags:
  - HTTP
  - URLルーティング
  - 木構造
  - router
translation_key: creating-url-routing-episode-1
---


# URLルーティングをつくる　エピソード1

## 概要
以前、Reactで非常に軟弱なルーティング（cf. [ReactとHistory APIを使ってrouterを自作する](https://bmf-tech.com/posts/React%E3%81%A8History%20API%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6router%E3%82%92%E8%87%AA%E4%BD%9C%E3%81%99%E3%82%8B)）を作ったが、改めてそこそこにちゃんとしたルーティングを自作したいと思い、挑戦することにした。
きっかけは、最近触っているGolangだ。
Golangでは標準ライブラリを駆使することでアプリーケーションをうすーく実装できるようだが、ルーティング周りは標準ライブラリがパワー不足なのものあって、外部のライブラリに依存するケースが多いらしい。
そんなこともあってルーティングを自作できるようになるとGolangでもそれ以外でもルーティングを自前で用意できて世界が広がる気がしたので重い腰を上げてやってみることにした。

## URLルーティングとはそもそも何をするのか？
リクエストされるURLに対して、実行したい処理は何か判定させるもの。
必要に応じて、パスパラメータやクエリパラーメータのデータを処理の実行時に扱えるようにする。

## URLルーティングの実装パターン
大まかに2パターン。

- 正規表現でURLのマッチングを行うパターン
- 木構造を用いて文字列探索を行うパターン

ルーティングがアプリケーションの実行速度に与える影響の割合はそこまでないかもしれないが、なるべく速いに越したことはないはず。
言語問わずメモリ使用量、計算量の最適化されたアルゴリズムで実装すべし。

今回は木構造で実装するパターンを選択する。
パフォーマンスを測定したわけではないが、正規表現よりも計算量が最適か木構造のアルゴリズムを用いるほうがパフォーマンス的にはよろしい気がするので、木構造にする。
実際、木構造で実装されているライブラリは多い。

## 木構造とは？
グラフ理論という数学の分野で定義されている木の構造を持つデータ構造のこと。
グラフ理論で定義されている木とは、複数の点（nodeまたはvertex）と複数の辺（edge）で構成されたグラフのことである。

```
   ○ ・・・根（root）
 / | ・・・枝（edge）
◯  ◯ ・・・節点（noteまたはvertex）
    \
  　  ◯
       \
    　   ○ ・・・葉（leaf）
```

ノードの性質や木の高さなどによって色々な種類の木構造があるが、ここでは割愛。

## 木構造の例
- 家系図
- ファイルシステム
- ドメイン名
  - cf. https://www.nic.ad.jp/ja/dom/system.html
- 構文木
  - コンパイラとか
- DOMツリー
- 階層構造を持つタグとかカテゴリの機能

## URLルーティングをつくる
何を木構造として扱うか？
これはもちろん、ルート定義のリストを木構造として扱う。

実装の流れをざっくり説明すると、ルート定義と現在のURL（パス）をインプットとして与えられたときに、
ルート定義から木構造を生成し、現在のURL（パス）をターゲットとして木構造を探索し、マッチしたデータを返すというだけ。

木構造を扱う際はノードの追加や削除等の処理も実装する場合があるが、URLルーティングの場合はとりあえず不要なので実装しない。

### データ構造を決める
ルーティング定義のDSLを先に決める。
多くのライブラリではシンプルなDSLが提供されているが、今回は複数階層あるちょっと複雑なDSLを定義する。

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

先程ルート定義から木構造を生成すると書いたが、ルート定義そのものを最初から木構造となるような形で定義することにする。
なぜこのような形をとったかというと単純に木構造を生成するアルゴリズムを書くのが面倒そうだったからであるが、逆に考えてみるとむしろ余計なアルゴリズムが減ってパーフォマンス的に良いではという気がしているがいかに・・
さほどわかりにくくないルート定義だと思うが、一般的なルーティングライブラリのDSLがこのようになっていないのは何か理由があるはずだとは思っている。

木構造の終端ノードとなる部分（葉）がちょうどHTTPメソッドになる。

木構造とは別にHTTPメソッドのリストを用意しておく。
Golangだとnet/httpに最初から定義されていて楽ですね。今回はPHPでやりますが・・

```php
$methods = [
    'GET',
    'POST',
    // more...
];
```

### 実装
インプットして与えられる現在のURL（パス）を木構造の探索の際に使いやすいように配列に加工する関数とその配列とルート定義の配列を引数としてマッチングしたパスのデータを返す関数の2つを実装する。
なお今回はクエリパラーメータは特に考慮していない。

実装方針として、他の言語への移植性を考慮し、ビルトイン関数の使用を極力避けて実装する。

```:php
function createCurrentPathArray($routes) {
    $currentPath = '/users/1'; // 現在のパス

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

// 探索　
// ルート定義と検索対象であるルートの配列を比較して該当するデータを返す。
// リーフに到達したら探索終了
function urlMatch($routes, $currentPathArray) {
    // TODO 実装中・・・
}

$currentPathArray = createCurrentPathArray($routes);
$result = urlMatch($routes, $currentPathArray);

var_dump($result); // マッチしたパスのデータが返るはず・・・
```

実装途中というわけでエピソード1はこれにて終幕。　

## 所感
最初からパトリシア木とかなんとか木とかからやろうとすると大やけどする。
参考になりそうな実装も色々見てみたが、一つ一つを理解するのは中々ハードなので、まずはアルゴリズムのイメージを掴むことと考えながら手を動かすことから始めてみたが、数学的素養が乏しいと辛いところはある。
実装途中ではあるが、割とゴールが見えるような気がしないでもない。
が、こんな感じで実運用に使えそうなベースまで持っていけるか自信はない。

##  追記
Makuake LT Party（社内LT大会）にてLTをした。

[speaker-deck - URLルーティングをつくるエピソード１](https://speakerdeck.com/bmf_san/urlruteinguwotukuruepisodo1)

### 参考
- [Algorithm visualization - Radix Tree](https://www.cs.usfca.edu/~galles/visualization/RadixTree.html)
- [github - [Japanese] Patricia Tree](~~https://github.com/ethereum/wiki/wiki/%5BJapanese%5D-Patricia-Tree~~)
- [WhiteDog@Blog](http://takao.blogspot.com/2012/03/patriciatrie.html)
- [404 Blog Not Found - algorithm - Patricia Trie (Radix Trie) を JavaScript で](http://blog.livedoor.jp/dankogai/archives/51766842.html)
- [http request multiplexerと文字列マッチング](http://web.archive.org/web/20201109013417/https://persol-pt.github.io/posts/tech-workshop1222/)
- [@IT - データ構造の選択次第で天国と地獄の差 (3/3)](http://www.atmarkit.co.jp/ait/articles/0809/01/news163_3.html)
- [基本データ構造：木構造のたどり方](http://www.sb.ecei.tohoku.ac.jp/lab/wp-content/uploads/2012/11/2012_d12.pdf)
- [pixiv inside - PHPで高速に動作するURLルーティングを自作してみた](https://devpixiv.hatenablog.com/entry/2015/12/13/145741)
- [PHPでフレームワークを使わずURLのルーティングをいい感じにやる](http://noranuk0.hatenablog.com/entry/2018/01/20/114933)
- [gist - neo-nanikaka/CommonPrefixTrieRouter.php](https://gist.github.com/neo-nanikaka/c2e2f7742b311696d50b)
- [github.com - nissy/bon](https://github.com/nissy/bon)
- [github.com - nissy/mux](https://github.com/nissy/mux)
- [github.com - ytakano/radix_tree](https://github.com/ytakano/radix_tree)
- [github.com - kkdai/radix](https://github.com/kkdai/radix)
- [github.com - MarkBaker/Tries](https://github.com/MarkBaker/Tries)

