---
title: Quickly Create a Markdown Editor with React
description: 'Create functional React markdown editor with marked library supporting live markdown preview and transformation.'
slug: react-markdown-editor-tutorial
date: 2017-12-25T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - ES5
  - ES6
  - JavaScript
  - React
translation_key: react-markdown-editor-tutorial
---

※This article is a repost from [Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/).

# Preparation
Setting up the build environment can be cumbersome, so this time we will use the official Facebook tool [create-react-app](https://github.com/facebookincubator/create-react-app).

`npm install -g create-react-app`

We will prepare the environment with the app name `md-editor`.

`create-react-app md-editor`

Next, let's install the libraries we will use this time.

`cd ./md-editor`

`npm install --save marked`

`npm install`

Finally, once the server is started, we are ready to go.
`npm start`

# Implementation

## STEP 1
Before we start implementing, let's delete the unnecessary files that we won't be using this time.

- `App.css` 
- `App.test.js`
- `logo.svg`

Make sure to remove the imports for these files from `src/index.js` and `src/App.js`.

Also, in `src/App.js`, let's leave the contents of the return statement empty. (We will ignore the warning during build about the return statement being empty for now.)

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

## STEP 2
Create a file named `Markdown.js` under the `src` directory. We will implement the markdown component in this file.

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

It's just a few lines. This will function as a markdown editor. It's almost raw JS. The only thing specific to React is JSX.

## STEP 3
Finally, let's import `Markdown.js` into `App.js`.

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

# Verification
If you want to highlight the source code, you can customize marked using [isagalaev/highlight.js - github](https://github.com/isagalaev/highlight.js) for a better experience.

# References
- [chjj/marked](https://github.com/chjj/marked)
- [React.Component - React](https://reactjs.org/docs/react-component.html#constructor)
- [super - MDN](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Operators/super#Description)

# Repository
The source code is available at ~~bmf-san/til/javascript/md-editor/ - github~~.

# Thoughts
I like React because it allows coding in a way that is close to raw JS, making it less likely to lock you into knowledge of the framework.

I have omitted most of the code explanations, but I think you can understand most of it by looking at the article [Modern JS Discussion by @bmf_san](http://tech.innovator.jp.net/archive/category/%E3%83%A2%E3%83%80%E3%83%B3%E3%81%AAJS%E3%81%AE%E8%A9%B1).