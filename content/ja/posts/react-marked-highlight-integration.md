---
title: React＋marked＋highlight
description: React＋marked＋highlightについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: react-marked-highlight-integration
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - ES6
  - React
  - highlightjs
  - markdown
  - marked
translation_key: react-marked-highlight-integration
---


wysiwygエディタではなく、マークダウンエディタをReactでつくってみました。

ソースコードの大部分は[React入門](http://yusuke-aono.hatenablog.com/entry/20150503/1430661392)を参考にさせていただきました。


雑なgifサンプルはこちら_(:3」∠)_
![markdown.gif](/assets/images/posts/react-marked-highlight-integration/a60a6293-1345-ae00-942c-e544e6e526a6.gif)



# 環境
* React
* marked([github](https://github.com/chjj/marked))・・・マークダウンパーサー
* highlight.js([highlightjs.org](https://highlightjs.org/))・・・シンタックスハイライト
* bower・・・上記全てのパッケージ管理に使用


# 準備
markedとhighlight.jsをbowerでインストール

`bower install marked`
`bower install highlightjs`

それぞれご自分の環境にインストールしてパスの設定までしておいてください。
bower install highlightではなく、highlightjsです。
両者は別物ようで、私はこれを間違えていたせいて小一時間ハマりました・・・・（泣）


# 実装

htmlはこんな感じで↓

index.html

```html
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8" />
<title>Hello React!</title>
<link href="//netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.css" rel="stylesheet">
<link href="path/to/monokai.css" rel="stylesheet">
<link href="path/to/style.css" rel="stylesheet">
</head>
    <body>

    <div class="markdown-component">

        <h1>React Markdown Editor</h1>

        <div id="content"></div>

    </div><!-- .component -->


    <!-- scripts -->
    <script src="path/to/react.js"></script>
    <script src="path/to/react-dom.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/babel-core/5.8.34/browser.min.js"></script>
    <script src="path/to/marked.min.js"></script>
    <script src="path/to/highlight.pack.min.js"></script>
    <script type="text/babel" src="path/to/markdown.js"></script>

    </body>
</html>
```

シンタックスハイライトに使用するカラーテーマはmonokaiが好きなのでmonokaiのスタイルシートを設定しました。
babelについては今回CDNを使用していますが、bowerでインストールしてもOKです。


さて、Reactコンポーネントを作っていきますが、冒頭でも述べたように、大部分は[React入門](http://yusuke-aono.hatenablog.com/entry/20150503/1430661392)を参考にしているので、こちらを一読しておくとよろしいかと思います。

参考ソースにhighlight.jsの設定コードだけ追加した感じです。（全然仕事していないｗｗ）



markdown.js

```js
var App = React.createClass({
    getInitialState: function() {
        return {
            markdown: ""
        };
    },

    updateMarkdown: function(markdown) {
        this.setState({
            markdown: markdown
        });
    },

    render: function() {
        return (
            <div>
                <TextInput onChange = {this.updateMarkdown}/>
                <Markdown markdown = {this.state.markdown}/>
            </div>
        );
    }
});

var TextInput = React.createClass({
    propTypes: {
        onChange: React.PropTypes.func.isRequired
    },

    _onChange: function(e) {
        this.props.onChange(e.target.value);
    },

    render: function() {
        return (
            <textarea onChange = {this._onChange}></textarea>
        );
    }
});

var Markdown = React.createClass({
    componentDidUpdate: function() {
        marked.setOptions({
            highlight: function(code, lang) {
                return hljs.highlightAuto(code, [lang]).value;
            }
        });
    },

    propTypes: {
        markdown: React.PropTypes.string.isRequired
    },

    render: function() {
        var html = marked(this.props.markdown);

        return (
            <div dangerouslySetInnerHTML={{__html: html}}></div>
        );
    }
});

ReactDOM.render(
    <App />,
    document.getElementById("content")
);
```

テキストの入力部分のコンポーネント、　マークダウンの出力部分のコンポーネント、それらを統合するコンポーネントの3つに分割されています。

マークダウンのパースはmarkedという関数で行っています。
このmarked関数のオプションをcomponentDidUpdateのところでhighlight.jsを使うよう設定しています。
オプションの設定方法についてはhighlight.jsのREADMEにかいてあります。

dangerouslySetInnerHTMLというのはxss対策でデータをサニタイズするプロパティです。


# 所感
初めてエディタつくったのですが、ライブラリでパパっと出来てしまうのですね〜_(:3」∠)_


# ES6バージョン
先日、ES6を勉強したので書き換えてみました。propsTypeの対応方法はよくわからなかったので省略してしまいましたｗ

```markdown.js
/**
 *
 * Editor
 *
 */

import React from 'react';
import ReactDOM from 'react-dom';

export default class Editor extends React.Component{
  constructor(props) {
    super(props);

    this.state = {
      markdown: ''
    };

    this.updateMarkdown = this.updateMarkdown.bind(this);
  }

  updateMarkdown(markdown) {
    this.setState({
      markdown: markdown
    });
  }

  render() {
    return (
      <div>
        <TextInput onChange={this.updateMarkdown}/>
        <Markdown markdown={this.state.markdown}/>
      </div>
    );
  }
};

class TextInput extends React.Component{
  constructor(props) {
    super(props);

    this._onChange = this._onChange.bind(this);
  }

  _onChange(e) {
    this.props.onChange(e.target.value);
  }

  render() {
    return (
      <textarea onChange={this._onChange}></textarea>
    );
  }
};

class Markdown extends React.Component{
  constructor(props) {
    super(props);
  }

  componentDidUpdate() {
    marked.setOptions({
        highlight: function(code, lang) {
          return hljs.highlightAuto(code, [lang]).value;
        }
    });
  }

  render() {
    var html = marked(this.props.markdown);

    return (
      <div dangerouslySetInnerHTML={{__html: html}}></div>
    );
  }
};
```

