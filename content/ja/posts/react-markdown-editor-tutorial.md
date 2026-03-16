---
title: Reactでマークダウンエディタをシャッとつくる
description: Reactでマークダウンエディタをシャッとつくる
slug: react-markdown-editor-tutorial
date: 2017-12-25T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - ES5
  - ES6
  - JavaScript
  - React
translation_key: react-markdown-editor-tutorial
---


※この記事は[Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/)で掲載されている記事を転載したものです。

# 準備
ビルド環境のセットアップが面倒なので、今回はFacebook公式の[create-react-app](https://github.com/facebookincubator/create-react-app)というツールを使います。

`npm install -g create-react-app`


md-editorというアプリ名で環境を用意することにします。

`create-react-app md-editor`

次に、今回使うライブラリのインストールをしておきます。

`cd ./md-editor`

`npm install --save marked`

`npm install`

最後にサーバーを起動したら準備OKです。
`npm start`

# 実装

## STEP1
実装に入る前に今回使用しない不要なファイルを削除しておきましょう。

- `App.css` 
- `App.test.js`
- `logo.svg`

`src/index.js`と``src/App.js`で上記ファイルをインポートしている部分を削除しておきます。

それから`src/App.js`のほうはreturn文の中身を空にしておきましょう。（ビルド時にreturn文が空で怒られますが一旦無視します。）


`src/index.js`
```
import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import registerServiceWorker from './registerServiceWorker';

ReactDOM.render(<App/>, document.getElementById('root'));
registerServiceWorker();
```

`src/App.js`
```javascript
import React, {Component} from 'react';

class App extends Component {
  render() {
    return ();
  }
}

export default App;
```

## STEP2
`src`以下に`Markdown.js`というファイルを作成します。
このファイルにはマークダウンのコンポーネントを実装していきます。

`src/Markdown.js`
```javascript
import React, {Component} from 'react';
import ReactDOM from 'react-dom';
import marked from 'marked';

class Markdown extends Component {
  constructor(props) {
    super(props);
    this.state = {
      html: ''
    };

    this.updateMarkdown = this.updateMarkdown.bind(this);
  }

  updateMarkdown(event) {
    this.setState({
      html: marked(event.target.value)
    });
  }

  render() {
    const html = this.state.html;

    return (<div>
      <h1>Markdown Input</h1>
      <textarea onChange={this.updateMarkdown}></textarea>
      <h1>Markdown Output</h1>
      <div dangerouslySetInnerHTML={{
          __html: html
        }}></div>
    </div>);
  }
}

export default Markdown;
```

ほんの数行です。
これだけとりあえずマークダウンとして機能します。
ほぼ生のJSですね。
React特有なのはJSXくらいでしょうか。

## STEP3
最後に`Mardown.js`を`App.js`内でインポートしましょう。

```javascript
import React, {Component} from 'react';
import Markdown from './Markdown';

class App extends Component {
  render() {
    return (<Markdown/>);
  }
}

export default App;
```

# 動作確認
ソースコードをハイライトしたい時には[isagalaev/highlight.js - github](https://github.com/isagalaev/highlight.js)を使ってmarkedをカスタマイズするといい感じになります。

# 参考
- [chjj/marked](https://github.com/chjj/marked)
- [React.Component - React](https://reactjs.org/docs/react-component.html#constructor)
- [super - MDN](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Operators/super#Description)

# リポジトリ
ソースコードは[bmf-san/til/javascript/md-editor/ - github](https://github.com/bmf-san/til/tree/master/javascript/md-editor)に置いてあります。

# 所感
Reactは素のJSに近い形でコーディングできるので、フレームワークに知識がロックインされづらいので好きです。

コードの説明はほとんど省きましたが、[モダンなJSの話 by @bmf_san](http://tech.innovator.jp.net/archive/category/%E3%83%A2%E3%83%80%E3%83%B3%E3%81%AAJS%E3%81%AE%E8%A9%B1)の記事を見て頂れば大体わかるのではないかと思います。
