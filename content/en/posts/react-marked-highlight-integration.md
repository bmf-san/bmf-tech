---
title: React＋marked＋highlight
slug: react-marked-highlight-integration
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - ES6
  - React
  - highlightjs
  - markdown
  - marked
translation_key: react-marked-highlight-integration
---


I created a markdown editor using React instead of a WYSIWYG editor.

Most of the source code was referenced from [React入門](http://yusuke-aono.hatenablog.com/entry/20150503/1430661392).

Here's a rough gif sample _(:3」∠)_
![markdown.gif](/assets/images/posts/react-marked-highlight-integration/a60a6293-1345-ae00-942c-e544e6e526a6.gif)


# Environment
* React
* marked([github](https://github.com/chjj/marked)) - Markdown parser
* highlight.js([highlightjs.org](https://highlightjs.org/)) - Syntax highlighting
* bower - Used for managing all the above packages


# Preparation
Install marked and highlight.js using bower

`bower install marked`
`bower install highlightjs`

Please install each in your environment and set the paths.
It's highlightjs, not bower install highlight.
They seem to be different, and I got stuck for about an hour because I made this mistake... (cry)


# Implementation

The HTML looks like this:

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

I like the monokai color theme for syntax highlighting, so I set the monokai stylesheet.
Regarding babel, I'm using a CDN this time, but you can also install it with bower.

Now, let's create the React components. As mentioned at the beginning, most of it is referenced from [React入門](http://yusuke-aono.hatenablog.com/entry/20150503/1430661392), so it might be good to read through it.

I just added the highlight.js setting code to the reference source. (Not much work done lol)


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

The components are divided into three: the text input component, the markdown output component, and the component that integrates them.

Markdown parsing is done with a function called marked.
The options for this marked function are set to use highlight.js in componentDidUpdate.
The method for setting options is written in the highlight.js README.

dangerouslySetInnerHTML is a property that sanitizes data for XSS protection.


# Impressions
This was my first time creating an editor, and it can be done quickly with libraries~ _(:3」∠)_


# ES6 Version
I recently studied ES6, so I rewrote it. I wasn't sure how to handle propsType, so I omitted it lol

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
