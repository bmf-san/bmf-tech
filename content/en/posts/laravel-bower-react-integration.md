---
title: Introducing React with Bower in Laravel
slug: laravel-bower-react-integration
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
  - React
translation_key: laravel-bower-react-integration
---

Recently, it seems that Laravel 5.3 has decided to adopt Vue.js as the default front-end framework.

I usually develop the front end using jQuery, but I decided to try React to keep up with recent trends. While I thought it might be safer to stick with Vue.js for Laravel, I chose React because it seems to be the most rapidly growing option right now. I was torn between AngularJS and React, but since I was looking for something to replace jQuery and handle only the view, I opted for React.

I can't really explain the technical value of each framework, so to be honest, I'm not very knowledgeable about it... lol

The official recommendation for installing React seems to be npm, but I feel more comfortable with Bower, so I will install it using Bower this time. (Is npm richer in packages than Bower??)

**Postscript:** After some research, it appears that npm has a more standard feel.

# Environment
* Laravel 5.2
* React... latest version for now (as of writing v15.3.0)
* Babel... a JavaScript compiler. **By introducing Babel, you can write React using JSX syntax.**
* Bower... [Refer to this setup guide for Bower in Laravel](hogehoge.com).
* Gulp (Elixir)

# Required Knowledge
* Knowledge of Bower setup.
* Knowledge of Gulp (Elixir)

You might get a grasp of things by looking at the official site tutorials or copying code found online, but it would be good to be aware of the recent chaotic state of the front-end world, including Babel, JSX, and Browserify.

# Setting Up React
`bower install react --save`
`bower install babel --save`

Use the following files:
* react.min.js (react)
* react-dom.js (react)
* browser.min.js (babel)

**If you want to use animations, please load react-with-addons.js instead of react.js.**

That's all you need to prepare to use React.

# Trying Out React
```html
<!DOCTYPE html>
<html>
  <head>
    <script src="path/to/react.min.js"></script>
    <script src="path/to/react-dom.min.js"></script>
    <script src="path/to/browser.min.js"></script>
    <script src="path/to/example.js" type="text/babel"></script>
  </head>

  <div id="example"></div>
```

```js
ReactDOM.render(
   <h1>Hello React Boy and Girl!</h1>,
   document.getElementById("example")
);
```

You can also write the JS before the closing body tag.

# Thoughts
I feel that there is less information available in Japanese about React compared to Laravel. (It seems that the amount of information about Laravel has increased rapidly this year...) However, it seems to be gaining attention, so I look forward to what’s next.

# Additional Note
If you want to use require() with React, you can use Browserify or Webpack. [What is require()?](http://qiita.com/uryyyyyyy/items/b10b012703b5396ded5a) Since Laravel has Browserify built-in by default, it might be easier to use that, but please choose according to your environment.

# References
* [How to Quickly Introduce React into a Laravel 5 Based Project](http://blog.mudatobunka.org/entry/2016/01/21/231546)
* [Let's Get Information from Laravel Using React.js](http://blog.comnect.jp.net/blog/98)
* [Introduction to React - Part 2: Using Browserify/Reactify/Gulp](http://qiita.com/masato/items/35b0900e3a7282b33bf8)
* [[Sy] How to Set Up a Development Environment for React Using Bower](https://utano.jp/entry/2016/07/react-js-install-use-bower/)