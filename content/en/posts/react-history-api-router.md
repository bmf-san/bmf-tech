---
title: Creating a Custom Router with React and the History API
description: 'Implement custom React router using History API with pushState and popstate for SPA navigation without libraries.'
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
translation_key: react-history-api-router
---

# Overview

# Preparation
First, let's understand the History API. GO TO MDN.

- [MDN - History](https://developer.mozilla.org/ja/docs/Web/API/History)
- [MDN - Manipulating the Browser History](https://developer.mozilla.org/ja/docs/Web/Guide/DOM/Manipulating_the_browser_history)

For those in a hurry, just understanding `pushState` and `window.popstate` should suffice.

# Specifications
This router will support the following URLs:

- `/post`
- `/post/:id`
- `/post/:id/:title`

It does not support query parameters.

# Packages Used
We will skip the React-related packages.

There is only one package used besides React:

[pillarjs/path-to-regexp](https://github.com/pillarjs/path-to-regexp)

This package helps with regular expressions for the URL part.

I would like to write my own regular expressions eventually, but for now, I will rely on this package.

# Implementation

## Create Components for Navigation and Pages
Prepare components corresponding to navigation and the pages.

```
src/
├── App.js
├── Dashboard.js
├── Home.js
├── Post.js
└── Profile.js
```

## Implement Routing
Now, let's implement the routing.

We will prepare two components: `Router` and `Route`.

`Router` is the component that handles rendering based on the URL.

`Route` is just a component that wraps an anchor tag.

We will also prepare a file called `routes.js` to describe the routing conventions.

`routes.js` is an array of objects that describes the paths and their corresponding components.

At this point, you might have an idea of the overall routing process:

**Initial State (First View)**
1. Get the current URL information.
2. Render the component that matches the current URL information.

The URL information is held as State.

**Transition**
1. Get the path of the clicked link.
2. Add to history and transition using the History API's `pushState`.
3. Re-render the component.

The State is updated, and the component is re-rendered.

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

*The strange line breaks in jsx are probably due to not properly configuring eslint...*

I referred to [You might not need React Router](https://medium.freecodecamp.org/you-might-not-need-react-router-38673620f3d) quite a bit.

The challenging part of the implementation was figuring out how to retrieve and maintain the information of parameters (like :id), but thanks to the awesome library `path-to-regexp`, I was able to overcome that.

# Github
Here is the source code for this project.

[bmf-san/rubel-router](https://github.com/bmf-san/rubel-router)

It is also published on npm.

[rubel-router](https://www.npmjs.com/package/rubel-router)

# Thoughts
I feel like it could be cleaner if I used EventEmitter or Observer... (lack of study)

# References
## Reference Articles
- [You might not need React Router](https://medium.freecodecamp.org/you-might-not-need-react-router-38673620f3d)
- [Building a React-based Application](https://reactjsnews.com/building-a-react-based-application)
- [Routing in React, the uncomplicated way](https://hackernoon.com/routing-in-react-the-uncomplicated-way-b2c5ffaee997)
- [MDN - History](https://developer.mozilla.org/ja/docs/Web/API/History)
- [MDN - Manipulating the Browser History](https://developer.mozilla.org/ja/docs/Web/Guide/DOM/Manipulating_the_browser_history)
- [Trying out the History API](http://www.allinthemind.biz/markup/javascript/history_api.html)
- [Notes on manipulating URLs with JavaScript](https://qiita.com/PianoScoreJP/items/fa66f357419fece0e531)

## Reference Sources
- [jsfiddle - frenzzy](https://jsfiddle.net/frenzzy/4ota5fag/2/)
- [jsfiddle - janfoeh](http://jsfiddle.net/janfoeh/2SCbv/)
- [jsfiddle - rgrove](http://jsfiddle.net/rgrove/WsHXm/)
- [jsfiddle - Swiftaxe](http://jsfiddle.net/Swiftaxe/zx9a17gg/)