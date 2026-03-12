---
title: Laravel+React+ES6+Browserify
description: Laravel+React+ES6+Browserifyについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: laravel-react-es6-browserify
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - React
  - JavaScript
  - npm
  - webpack
  - ES6
  - browerify
  - ES5
translation_key: laravel-react-es6-browserify
---


ちょっと前の記事で、bowerでReactの環境をセットアップするという話をしたのですが、npmでReactのパッケージ管理をしたほうがスマートだし、ES6かけるようになっといた方がこの先お得よねということで環境を再構築しました。

ES5からES6で書き方が色々変わるのでその辺の改修がちょっと面倒でしたが、さほど難しいことはないので気負いしなくとも良さそうです。


# 環境
* Laravel5.2・・・5.1とか5.2ユーザーはelixirのバージョンを最新（5.3のそれと同じ）に上げとくといいかもです。
* Browerify(Elixirについてるやつ)
* React
* ES6

# Reactをnpmで用意
`npm i react react-dom -D`

# コンパイル
```gulpfile.js
elixir(function(mix) {
  mix
    .browserify('hoge.js', 'hogehoge.js')
});

```

# ES5をES6に改修
こちらがとても参考になります。
[ES5のReact.jsソースをES6ベースに書き換える](http://qiita.com/kuniken/items/2e850daa26a10b5098d6)

# ES6をES5に対応させたい
トランスパイラを利用。babelとか。

# 所感
ただのメモ書きでしたφ(..)

