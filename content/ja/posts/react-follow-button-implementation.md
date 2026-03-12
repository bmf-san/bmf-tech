---
title: Reactでフォローボタンをつくってみる
description: Reactでフォローボタンをつくってみるについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: react-follow-button-implementation
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - React
translation_key: react-follow-button-implementation
---


![react_follow_button.gif](/assets/images/posts/react-follow-button-implementation/f532fbdd-45bb-93b4-aacf-d6220f58663a.gif)


# つくるもの
ツ◯ッターのフォローボタンを~~パクった~~リスペクトしたものをつくります。仕様はだいたい同じだと思いますが、仕組みは異なります。
クリックでフォロー／フォロー中とテキストが切り替わる、フォロー中の時にホバーした場合は解除というテキストを出す。これだけです。
やや装飾にこだわって全体に無駄なCSSが設定されていますが、その辺は適宜スタイルシートを調整してください。


# 必要な知識
* Reactの導入方法及び簡易的なコンポーネント作成方法
* jsxとbabelについての多少の知識と理解


# 環境
* React・・・v15.3.0
* babel・・・コンパイラ（jsxもコンパイルしてくれるそうです）


# htmlとcssを先に用意

**※パスは適宜調整してください！ (~~直すのめんどくさかった~~)**

```html
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8" />
<title>Hello React!</title>
<link href="//netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.css" rel="stylesheet">
<link href="style.css" rel="stylesheet">
</head>
    <body>

    <div class="component">

        <p><span><i class="fa fa-twitter fa-4x"></i></span></p>
        <h1>React Follow Button Component</h1>

        <div id="content"></div>

    </div><!-- .component -->


    <!-- scripts -->
    <script src="build/react.js"></script>
    <script src="build/react-dom.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/babel-core/5.8.34/browser.min.js"></script>
    <script type="text/babel" src="follow.js"></script>

    </body>
</html>
```

```css
html {
    height: 100%;
    margin: 0;
    height: 100vh;
}

body {
    color: #F1F1F1;
    font-family: 'Open Sans',Arial,Helvetica Neue,sans-serif;
    min-width: 100%;
    min-height: 100%;
    margin-top: 200px;
    background: linear-gradient(230deg, #a24bcf, #4b79cf, #4bc5cf);
    background-size: 300% 300%;
    /*-webkit-animation: bodyBg 60s ease infinite;
    -moz-animation: bodyBg 60s ease infinite;
    animation: bodyBg 60s ease infinite;*/
}

@-webkit-keyframes bodyBg {
    0%{background-position:0% 84%}
    50%{background-position:100% 16%}
    100%{background-position:0% 84%}
}
@-moz-keyframes bodyBg {
    0%{background-position:0% 84%}
    50%{background-position:100% 16%}
    100%{background-position:0% 84%}
}
@keyframes bodyBg {
    0%{background-position:0% 84%}
    50%{background-position:100% 16%}
    100%{background-position:0% 84%}
}

.component {
    text-align: center;
}

h1 {
    line-height: 0.8;
    letter-spacing: 3px;
    font-weight: 300;
    text-align: center;
    margin-bottom: 40px;
}


/*
Follow Button
 */
.follow-button {
    display: block;
    margin: 0 auto;
    width: 200px;
    color: white;
    font-size: 20px;
    padding: 10px 40px 10px 40px;
    border: solid white 1px;
    border-radius: 2px;
    cursor: pointer;
}

.follow-button:hover {
    transition: .3s;
    color: #43cea2;
    background-color: white;
}
```


id名contentのdiv内にフォローボタンのコンポーネントを生成していきます。
cssのfollow-buttonのクラスは生成するフォローボタンのスタイルです。
（挙動だけ確認したい方はcssはスルーしても問題ないでしょう）


# フォローボタンのコンポーネントをつくる

```js
var FollowButton = React.createClass({
    getInitialState: function () {
        return {
            value: "フォロー",
            toggle: false
        };
    },

    handleClick: function () {
        if (this.state.toggle) {
            this.setState({
                value: "フォロー",
                toggle: false
            });
        } else {
            this.setState({
                value: "フォロー中",
                toggle: true
            });
        };
    },

    handleMouseOver: function () {
        if (this.state.toggle) {
            this.setState({
                value: "解除",
            });
        };
    },

    handleMouseOut: function () {
        if (this.state.toggle) {
            this.setState({
                value: "フォロー中",
            });
        };
    },

    render: function () {
        return (
            <span className="follow-button" onClick={this.handleClick} onMouseOver={this.handleMouseOver} onMouseOut={this.handleMouseOut}>
                {this.state.value}
            </span>
        );
    },

});

ReactDOM.render(
    <FollowButton />,
    document.getElementById('content')
);
```

# 所感
Reactチュートリアル並のコンポーネントをつくれるようになるにはまだ場数が必要なようです_(:3」∠)_

# 参考
* [React で要素のクラスを動的に付け外しするなら JedWatson さんちの classnames が便利](http://qiita.com/taka1970/items/2b220b1c249a29797a08)
* [Reactのコード事例から学ぶ初心者向けReact入門と事例集](http://tango-ruby.hatenablog.com/entry/2016/04/30/090000)・・・段階を踏んでReactを学べます！おすすめ！

