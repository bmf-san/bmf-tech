---
title: "LaravelにbowerでReactを導入してみる"
slug: "laravel-bower-react-integration"
date: 2017-10-01
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Laravel"
  - "React"
draft: false
---

先日、Laravel5.3ではフロントエンドフレームワークとしてVue.jsをデフォルトとして採用するのが決定したようです。

普段、フロントエンドはjQueryで開発しているのですが、最近の流行りに乗じてReactを使ってみることにしました。
LaravelならVue.jsにしとくのが無難かなと考えたのですが、Reactが今一番伸びがある（らしい）のでReactにしました。
AngularJSとで迷ったのですが、あくまでjQueryの代わりになるかつViewだけ担当するものを考えていたのでReactを選択しました。

それぞれのFWの技術的な利用価値について説明できるほどフロントエンドマンではないので、ぶっちゃけよくわかっていませんが・・・ｗ


Reactのインストールは公式ではnpmが推奨？されているようですが、bowerの方が何となく親しみがあるので、今回はbowerでインストールします。（bowerよりnpmの方がパッケージが豊富？？）


後日談：色々調べてみるとnpmのほうがスタンダードな感じがありますね。

# 環境
* Laravel5.2・
* React・・・とりあえず最新バージョン（執筆時 v15.3.0）
* babel・・・JavaScriptのコンパイラ。**jsxを解釈してくれるのでbabel導入でjsxのシンタックスでReactをかくことができます。**
* bower・・・[laravelでbowerのセットアップ](hogehoge.com)をご参考に。
* gulp（elixir）

# 必要な知識
* bowerのセットアップの知識。
* gulp(elixir)の知識

公式サイトのチュートリアルやネットに転がっているソースを眺めたり写経するうちに何となくわかってくるかと思いますが、babelやjsx,browerifyなど最近のカオスなフロントエンド界隈事情を知っておくと良いでしょう。


# Reactのセットアップ
`bower install react --save`
`bower install babel --save`

以下のファイルを使います。
* react.min.js (react)
* react-dom.js (react)
* browser.min.js (babel)


※animationとか使いたい場合は、**react-with-addons.jsをreact.jsの代わりに読み込んでください。**


Reactを使う準備はこれだけです。


# Reactやってみる
```html
<!DOCTYPE html>
<html>
  <head>
    <script src="path/to/react.min.js"></script>
    <script src="path/to/react-dom.min.js"></script>
    <script src="path/to/browser.min.js"></script>
    <script src="path/to/example.js" type="text/babel"></script>
  </head>

  <div id="example"></div>
```

```js
ReactDOM.render(
   <h1>Hello React Boy and Girl!</h1>
   document.getElementId("example")
);
```


jsはbodyタグの終了前に記述してもOKです。


# 所感
Reactの日本語情報はLaravelよりも少ないように感じました。（Laravelは今年に入って急激に増えた気がしますが。。）
ただ注目度は高いようなので今後に期待です。


# 追記
Reactでrequire()を使用したい場合はbrowserifyやwebpackを使用します。　[require()ってなに？](http://qiita.com/uryyyyyyy/items/b10b012703b5396ded5a)
laravelにはデフォルトでbrowerserifyが組み込まれているのでそちらを使用するのが楽かと思いますが、環境に合わせて選択してください。

# 参考
* [Laravel5 ベースのプロジェクトに React が爆速で導入できた話](http://blog.mudatobunka.org/entry/2016/01/21/231546)
* [React.jsでLaravelから情報をもらってみよう](http://blog.comnect.jp.net/blog/98)
* [React入門 - Part2: Browserify/Reactify/Gulpを使う](http://qiita.com/masato/items/35b0900e3a7282b33bf8)
* [[Sy] bowerを使ってReactの開発環境を構築する方法](https://utano.jp/entry/2016/07/react-js-install-use-bower/)

