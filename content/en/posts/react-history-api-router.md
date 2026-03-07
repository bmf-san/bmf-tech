---
title: Creating a Custom Router with React and the History API
slug: react-history-api-router
date: 2018-01-03T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - ES6
  - JavaScript
  - React
  - History API
  - Router
description: Learn how to build a custom router using React and the History API.
translation_key: react-history-api-router
---

# Overview

# Preparation
First, understand the History API. GO TO MDN.

- [MDN - History](https://developer.mozilla.org/ja/docs/Web/API/History)
- [MDN - Manipulating the browser history](https://developer.mozilla.org/ja/docs/Web/Guide/DOM/Manipulating_the_browser_history)

If you're short on time, understanding just `pushState` and `window.popstate` should suffice.

# Specifications
This router will support the following URLs:

- `/post`
- `/post/:id`
- `/post/:id/:title`

Query parameters are not supported.

# Packages Used
We'll skip over the React-related packages.

Other than React, we'll use just one package:

[pillarjs/path-to-regexp](https://github.com/pillarjs/path-to-regexp)

This package handles regular expressions for URLs efficiently.

Someday, I want to write my own regular expressions, but for now, I'll rely on this package.

# Implementation

## Create Components for Navigation and Pages
Prepare components for navigation and the pages corresponding to each navigation link.

```
src/
├── App.js
├── Dashboard.js
├── Home.js
├── Post.js
└── Profile.js
```

## Implement Routing
Let's implement the routing.

We'll prepare two components: `Router` and `Route`.

`Router` is a component that switches rendering based on the URL.

`Route` is a component that simply wraps an anchor (`a`) tag.

Additionally, we'll create a file called `routes.js` to define the routing rules.

`routes.js` will contain an array of objects that map paths to their corresponding components.

By now, you might have guessed the sequence of routing operations:

**Initial State (First View)**
1. Retrieve the current URL information.
2. Render the component that matches the current URL information.

The URL information is stored in the state.

**Navigation**
1. Retrieve the path of the clicked link.
2. Use the History API's `pushState` to add to the history and navigate.
3. Re-render the component.

The state is updated, and the component is re-rendered.

The implementation of each component looks like this:

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

*Note: The strange line breaks in the JSX code might be due to improper ESLint settings.*

I referred heavily to [You might not need React Router](https://medium.freecodecamp.org/you-might-not-need-react-router-38673620f3d).

The most challenging part of the implementation was figuring out how to retrieve and manage parameter information (e.g., `:id`). Thanks to the awesome library `path-to-regexp`, I was able to overcome this issue.

# Github
Here is the source code for this implementation:

[bmf-san/rubel-router](https://github.com/bmf-san/rubel-router)

It is also published on npm:

[rubel-router](https://www.npmjs.com/package/rubel-router)

# Thoughts
Using EventEmitter or Observer might make the implementation cleaner... (I need to study more).

# References
## Articles
- [You might not need React Router](https://medium.freecodecamp.org/you-might-not-need-react-router-38673620f3d)
- [Building a React-based Application](https://reactjsnews.com/building-a-react-based-application)
- [Routing in React, the uncomplicated way](https://hackernoon.com/routing-in-react-the-uncomplicated-way-b2c5ffaee997)
- [MDN - History](https://developer.mozilla.org/ja/docs/Web/API/History)
- [MDN - Manipulating the browser history](https://developer.mozilla.org/ja/docs/Web/Guide/DOM/Manipulating_the_browser_history)
- [Using the History API](http://www.allinthemind.biz/markup/javascript/history_api.html)
- [Memo on manipulating URLs with JavaScript](https://qiita.com/PianoScoreJP/items/fa66f357419fece0e531)

## Source Code
- [jsfiddle - frenzzy](https://jsfiddle.net/frenzzy/4ota5fag/2/)
- [jsfiddle - janfoeh](http://jsfiddle.net/janfoeh/2SCbv/)
- [jsfiddle - rgrove](http://jsfiddle.net/rgrove/WsHXm/)
- [jsfiddle - Swiftaxe](http://jsfiddle.net/Swiftaxe/zx9a17gg/)