---
title: ReactとReduxでSPAを構築するために参考にした記事
description: "参考にするReact・Redux資源。SPA構築で役立つリポジトリ、記事、ReducerやAction、middleware、Redux Formの実装例集。"
slug: react-redux-spa-resources
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - React
  - Redux
  - SPA
translation_key: react-redux-spa-resources
---


雑メモです。
ReactとReduxでSPAで作ろうとした時に参考にしたリポジトリや記事です。

Reduxというアーキテクチャの概念を理解するには、日本語情報がそこそこ充実していました。
実際につくりたいモノの参考になりそうなソースを探すには少し手間がかかりました。
js仕様の違い、コンパイラやタスクランナーなどビルド環境の違い、Reactそのもののバージョンの違い、関連パッケージのバージョンの違いなどペチパー一筋だった自分には中々しんどかったです。。。

色々調べてみてようやく手を動かせるようになったので、その時参考になったリポジトリや記事をリストアップしておきます。

※都度更新していきますー

# 参考URLリスト
github - onerciller/react-redux-laravel
https://github.com/onerciller/react-redux-laravel

github - rajaraodv/react-redux-blog
https://github.com/rajaraodv/react-redux-blog

github - mustafawm/blogapp
~~(dead link)~~

github - mzabriskie/axios
https://github.com/mzabriskie/axios#handling-errors

React-Redux をわかりやすく解説しつつ実践的に一部実装してみる
http://ma3tk.hateblo.jp/entry/2016/06/20/182232

React + ReduxのプロジェクトでRedux Formを使ったので使い方のまとめと注意点
http://ichimaruni-design.com/2016/10/react-redux-form/

Reduxでコンポーネントを再利用する
http://qiita.com/kuy/items/869aeb7b403ea7a8fd8a

【Redux入門】 React + Redux の考え方を理解する
~~(dead link)~~

Redux入門 6日目 ReduxとReactの連携(公式ドキュメント和訳)
http://qiita.com/kiita312/items/d769c85f446994349b52

Reduxでのクライアントサイドvalidationをどこでやるべきか？
http://qiita.com/inuscript/items/5bed7812b3c1447b7b60

Reduxの実装とReactとの連携を超シンプルなサンプルを使って解説
http://mae.chab.in/archives/2885

React-Redux をわかりやすく解説しつつ実践的に一部実装してみる
http://ma3tk.hateblo.jp/entry/2016/06/20/182232

Redux入門 3日目 Reduxの基本・Reducers(公式ドキュメント和訳)
http://qiita.com/kiita312/items/7fdce94912d6d9c801f8

【React/Redux】わたしもみている | Container Components
http://kenjimorita.jp/read1/

Reduxで非同期処理をしたいときに、なぜMiddlewareを使わないといけないのか
http://qiita.com/enshi/items/557dcd7df60e6128249e

react+reduxで非同期処理を含むtodoアプリを作ってみる
http://qiita.com/halhide/items/a45c7a1d5f949596e17d

もうはじめよう、ES6~ECMAScript6の基本構文まとめ(JavaScript)~
http://qiita.com/takeharu/items/cbbe017bbdd120015ca0

React-Reduxを使った開発でのディレクトリ構成をどうしたらいいのか的なことから、こうやって組んだらいいんじゃないか的なお話
http://watanabeyu.blogspot.jp/2016/08/react-redux.html

Connecting Mapdisptachtoprops in V6 reduxForm()
~~(dead link)~~

react-routerでURLパラメータを指定した際、URL直打ちだと404になります
https://teratail.com/questions/26245

redux-form Multiple field errors?
~~(dead link)~~

Redux Form -Submit Validation Example
http://redux-form.com/6.0.0-alpha.4/examples/submitValidation/

throw new SubmissionError() causing Uncaught (in promise) error
~~(dead link)~~

redux-form textarea value not updating
http://stackoverflow.com/questions/40970691/redux-form-textarea-value-not-updating

# 所感
もうちょっとgithubのコードをあさってみたいところです。

