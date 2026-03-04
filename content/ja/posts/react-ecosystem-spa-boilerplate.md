---
title: "React+react-redux+react-router+ES6+webpackで作るSPAボイラープレート"
slug: "react-ecosystem-spa-boilerplate"
date: 2017-10-01
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "boilerplate"
  - "ES6"
  - "React"
  - "react-router"
  - "Redux"
  - "webpack"
draft: false
---

# 概要
Reactで作るSPAの簡易的なボイラープレートをつくりました。
最近のフロントエンドには何とかついていくだけで精一杯なため、ソースに自信はありませんが、一応形にはなっていると思いますと言い訳だけしておきます(*_*)

# 環境
## モジュール
+ axios
+ lodash
+ react
+ react-dom
+ react-router
+ react-redux
+ redux
+ redux-promise

## ビルド周り
+ babel-core
+ babel-loader
+ babel-preset-es-2015
+ babel-preset-react
+ babel-preset-stage-0
+ webpack
+ webpack-dev-server

※ほとんどが現時点での最新版を使っていると思いますが、react-routerなんかは一つ前のバージョンだと思います。

# 仕様
+ ルーティング
+ APIコール

これだけです(--)

# 動作
![redux.gif](https://qiita-image-store.s3.amazonaws.com/0/124495/409d4146-bc7e-7f54-5c9f-ad963b56a25b.gif)

# ソース
+ [bmf-san/react-redux-spa-boilerplate](https://github.com/bmf-san/react-redux-spa-boilerplate)

# 所感
実際にそこそこのSPAを構築すると、さらに外部のライブラリを追加したり、コンポーネントが複雑化したりしてカオスになりますが、そこは自分の設計力が足りないとこなんだと思います。

