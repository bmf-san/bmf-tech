---
title: "ReactとHistory APIを使ってrouterを自作する"
slug: "react-history-api-router"
date: 2018-01-03
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "ES6"
  - "JavaScript"
  - "React"
  - "history-api"
  - "router"
draft: false
---

# 概要

# 準備
まずはHistory APIを理解しておきます。GO TO MDN。

- [MDN - History](https://developer.mozilla.org/ja/docs/Web/API/History)
- [MDN - ブラウザの履歴を操作する](https://developer.mozilla.org/ja/docs/Web/Guide/DOM/Manipulating_the_browser_history)

忙しい人は`pushState`と`window.popstate`だけ理解しておけばなんとかなるはず。

# 仕様
このrouterでは、以下のようなURLに対応します。

- `/post`
- `/post/:id`
- `/post/:id/:title`

クエリパラメータには対応しません。

# 使用するパッケージ
React周りは省略します。

React以外で使うパッケージは1つだけです。

[pillarjs/path-to-regexp](https://github.com/pillarjs/path-to-regexp)

URL部分の正規表現を良しなにやってくれるパッケージです。

そのうち自分で正規表現書きたいですが、今回はパッケージに頼っちゃいます。

# 実装

## ナビゲーションとページに対応するコンポーネントを作成
ナビゲーション、ナビゲーションにそれぞれ対応するコンポーネントを用意しておきます。

```
src/
├── App.js
├── Dashboard.js
├── Home.js
├── Post.js
└── Profile.js
```


## ルーティングの実装
ルーティングを実装していきます。

コンポーネントは、`Router`と`Route`という2つを用意します。

`Router`はURLに応じて描画切り替えを行うコンポーネントです。

`Route`はaタグをラップしただけのコンポーネントです。

それからルーティング規約を記述するファイルとして、`routes.js`を用意します。

`routes.js`はパスと、パスに対応するコンポーネントの対応をオブジェクトの配列で記述したものです。

ここまででおおよそ察しがつくかと思いますが、ルーティングの一連の処理としては、

**初期状態（ファーストビュー）**
①現在のURL情報を取得
②現在のURL情報に一致するコンポーネントを描画

URL情報をStateとして持ちます。

**遷移**
①クリックされたリンクのパスを取得
②History APIの`pushState`で履歴を追加・遷移
③コンポーネントを再描画

Stateが更新され、コンポーネントが再描画されます。


各コンポーネントの実装はこんな感じです。

`Route.js`

```javascript
import React, {Component} from 'react';
const history = window.history;

class Route extends Component {
  constructor(props) {
    super(props);

    this.handleClick = this.handleClick.bind(this);
  }

  handleClick(event) {
    event.preventDefault();

    const info = {
      'url': event.target.href,
      'path': event.target.pathname
    };

    this.handlePush(info.url);
    this.props.handleRoute(info);
  }

  handlePush(url) {
    // Create a history, and transition to next url
    history.pushState(null, null, url);
  }

  render() {
    return (<React.Fragment>
      <a href={this.props.path} onClick={this.handleClick}>{this.props.text}</a>
    </React.Fragment>);
  }
}

export default Route;
```

`Router.js`

```javascript
import React, {Component} from 'react';
import toRegex from 'path-to-regexp';

class Router extends Component {
  handleComponent() {
    const routes = this.props.routes;
    const info = this.props.info;

    for (const route of routes) {
      const keys = [];
      const string = new String(route.path);
      const pattern = toRegex(string, keys);
      const match = pattern.exec(info.path);

      if (!match) {
        continue;
      }

      const params = Object.create(null);
      for (let i = 1; i < match.length; i++) {
        params[keys[i - 1].name] = match[i] !== undefined
          ? match[i]
          : undefined;
      }

      if (match) {
        return route.action(Object.assign(info, {"params": params}));
      }
    }

    return 'Not Found';
  }

  render() {
    return (this.handleComponent());
  }
}

export default Router;
```

`routes.js`

```javascript
import React, {Component} from "react";
import Home from "./Home";
import Dashboard from "./Dashboard";
import Profile from "./Profile";
import Post from "./Post";

const HomeComponent = (params) => (<Home {...params}/>);
const DashboardComponent = (params) => (<Dashboard {...params}/>);
const ProfileComponent = (params) => (<Profile {...params}/>);
const PostComponent = (params) => (<Post {...params}/>);

export const routes = [
  {
    path: "/",
    action: HomeComponent
  }, {
    path: "/dashboard",
    action: DashboardComponent
  }, {
    path: "/profile",
    action: ProfileComponent
  }, {
    path: "/post/:id",
    action: PostComponent
  }
];
```

`App.js`

```javascript
import React, {Component} from 'react';
import Router from './Router';
import Route from './Route';
import {routes} from './routes';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      'url': '', // current url
      'path': '' // current path
    };

    this.handleRoute = this.handleRoute.bind(this);
  }

  handleRoute(info) {
    // Update url info
    this.setState(info);
  }

  render() {
    return (<React.Fragment>
      <p>Current URL: {this.state.url}</p>
      <p>Current Path: {this.state.path}</p>
      {/* Navigation */}
      <ul>
        <li>
          <Route path="/" text="Top" handleRoute={this.handleRoute}/>
        </li>
        <li>
          <Route path="/dashboard" text="Dashboard" handleRoute={this.handleRoute}/>
        </li>
        <li>
          <Route path="/profile" text="Profile" handleRoute={this.handleRoute}/>
        </li>
        <li>
          <Route path="/post/9" text="Post-Id" handleRoute={this.handleRoute}/>
        </li>
      </ul>
      {/* Router Component */}
      <Router routes={routes} info={this.state}/>
    </React.Fragment>);
  }
}

export default App;
```

※jsxの改行がなんか変なのは多分eslintをちゃんと設定していないからだと思います...

[You might not need React Router](https://medium.freecodecamp.org/you-might-not-need-react-router-38673620f3d)
を結構参考にしました。

実装する上で厄介だった部分は、「パラメータ（:id）の情報をどうやって取得するか、保持するか」という点でしたが、`path-to-regexp`というawesomeなライブラリのおかげで、その点は克服できました。

# Github
今回のソース置いておきます。

[bmf-san/rubel-router](https://github.com/bmf-san/rubel-router)

npmにも公開しています。

[rubel-router](https://www.npmjs.com/package/rubel-router)

# 所感
EventEmitterやObserverをつかったらもっと綺麗になる気が・・（勉強不足）

# 参考
## 参考記事
- [You might not need React Router](https://medium.freecodecamp.org/you-might-not-need-react-router-38673620f3d)
- [Building a React-based Application](https://reactjsnews.com/building-a-react-based-application)
- [Routing in React, the uncomplicated way](https://hackernoon.com/routing-in-react-the-uncomplicated-way-b2c5ffaee997)
- [MDN - History](https://developer.mozilla.org/ja/docs/Web/API/History)
- [MDN - ブラウザの履歴を操作する](https://developer.mozilla.org/ja/docs/Web/Guide/DOM/Manipulating_the_browser_history)
- [History API を使ってみる](http://www.allinthemind.biz/markup/javascript/history_api.html)
- [JavaScriptでURLを操作するメモ](https://qiita.com/PianoScoreJP/items/fa66f357419fece0e531)

## 参考ソース
- [jsfiddle - frenzzy](https://jsfiddle.net/frenzzy/4ota5fag/2/)
- [jsfiddle - janfoeh](http://jsfiddle.net/janfoeh/2SCbv/)
- [jsfiddle - rgrove](http://jsfiddle.net/rgrove/WsHXm/)
- [jsfiddle - Swiftaxe](http://jsfiddle.net/Swiftaxe/zx9a17gg/)
