---
title: Creating a Follow Button with React
slug: react-follow-button-implementation
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - React
description: Learn how to create a Twitter-style follow button using React.
translation_key: react-follow-button-implementation
---

![react_follow_button.gif](https://qiita-image-store.s3.amazonaws.com/0/124495/f532fbdd-45bb-93b4-aacf-d6220f58663a.gif)

# What We'll Create
We'll create a follow button inspired by Twitter's design. The functionality is similar, but the implementation differs. The button toggles between "Follow" and "Following" text when clicked, and displays "Unfollow" text when hovered over while in the "Following" state. Feel free to adjust the CSS styling as needed.

# Required Knowledge
* Basic knowledge of React setup and component creation
* Some understanding of JSX and Babel

# Environment
* React: v15.3.0
* Babel: Compiler (also compiles JSX)

# Prepare HTML and CSS First

**Note: Adjust paths as needed!**

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

The `div` with the ID `content` will contain the follow button component. The `follow-button` CSS class defines the style of the button. If you only want to check the behavior, you can skip the CSS.

# Creating the Follow Button Component

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

# Thoughts
It seems I still need more practice to create components at the level of React tutorials _(:3」∠)_

# References
* [React: Dynamically Add/Remove Element Classes with JedWatson's classnames](http://qiita.com/taka1970/items/2b220b1c249a29797a08)
* [Beginner's Guide to React and Example Collection](http://tango-ruby.hatenablog.com/entry/2016/04/30/090000) - A great resource for learning React step by step!