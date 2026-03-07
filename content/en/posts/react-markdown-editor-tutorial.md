---
title: Quickly Create a Markdown Editor with React
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

※This article is a repost from the [Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/).

# Preparation
Setting up a build environment can be tedious, so this time we’ll use Facebook’s official tool, [create-react-app](https://github.com/facebookincubator/create-react-app).

`npm install -g create-react-app`

We’ll set up the environment with the app name `md-editor`.

`create-react-app md-editor`

Next, let’s install the libraries we’ll use this time.

`cd ./md-editor`

`npm install --save marked`

`npm install`

Finally, start the server, and the setup is complete.
`npm start`

# Implementation

## STEP1
Before starting the implementation, let’s delete unnecessary files that we won’t use this time.

- `App.css` 
- `App.test.js`
- `logo.svg`

Remove the import statements for the above files in `src/index.js` and `src/App.js`.

Then, in `src/App.js`, clear the content inside the return statement. (You’ll get a build error because the return statement is empty, but you can ignore it for now.)

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
Create a file called `Markdown.js` under the `src` directory. This file will contain the implementation of the Markdown component.

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

Just a few lines of code. This is enough to make it function as a Markdown editor. It’s almost pure JavaScript. The only React-specific part is the JSX.

## STEP3
Finally, import `Markdown.js` into `App.js`.

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

# Testing
If you want to highlight source code, you can customize `marked` using [isagalaev/highlight.js - github](https://github.com/isagalaev/highlight.js).

# References
- [chjj/marked](https://github.com/chjj/marked)
- [React.Component - React](https://reactjs.org/docs/react-component.html#constructor)
- [super - MDN](https://developer.mozilla.org/en/docs/Web/JavaScript/Reference/Operators/super#Description)

# Repository
The source code is available at [bmf-san/til/javascript/md-editor/ - github](https://github.com/bmf-san/til/tree/master/javascript/md-editor).

# Thoughts
I like React because it allows coding in a way that’s close to pure JavaScript, which reduces the risk of being locked into framework-specific knowledge.

I’ve skipped most of the code explanations, but you can check out the article [Modern JavaScript by @bmf_san](http://tech.innovator.jp.net/archive/category/%E3%83%A2%E3%83%80%E3%83%B3%E3%81%AAJS%E3%81%AE%E8%A9%B1) for more details.