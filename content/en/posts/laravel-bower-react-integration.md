---
title: Introducing React to Laravel with Bower
slug: laravel-bower-react-integration
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
  - React
description: A guide on integrating React into a Laravel project using Bower.
translation_key: laravel-bower-react-integration
---


Recently, it seems that Laravel 5.3 has decided to adopt Vue.js as the default frontend framework.

I usually develop the frontend with jQuery, but I decided to try React to keep up with the latest trends. Although I thought sticking with Vue.js would be safer with Laravel, I chose React because it seems to be growing the most right now. I was torn between AngularJS and React, but since I was looking for something to replace jQuery and only handle the View, I chose React.

I'm not a frontend expert who can explain the technical value of each framework, so honestly, I don't really understand... lol

The official recommendation for installing React seems to be npm, but I feel more familiar with Bower, so I'll install it with Bower this time. (Is npm richer in packages than Bower??)

After some research, it seems that npm is more standard.

# Environment
* Laravel 5.2
* React... the latest version at the time of writing (v15.3.0)
* babel... a JavaScript compiler. **With babel, you can write React in JSX syntax as it interprets jsx.**
* bower... Refer to [laravelでbowerのセットアップ](hogehoge.com) for setup.
* gulp (elixir)

# Required Knowledge
* Knowledge of setting up Bower.
* Knowledge of gulp (elixir).

By looking at tutorials on the official site or sources available online, you might get a sense of it, but it's good to be aware of the recent chaotic frontend landscape, including babel, jsx, and browserify.

# Setting Up React
`bower install react --save`
`bower install babel --save`

Use the following files:
* react.min.js (react)
* react-dom.js (react)
* browser.min.js (babel)

※ If you want to use animations, **load react-with-addons.js instead of react.js.**

This is all you need to prepare to use React.

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
   <h1>Hello React Boy and Girl!</h1>
   document.getElementId("example")
);
```

You can also write the js before the closing body tag.

# Impressions
It seems there is less information about React in Japanese compared to Laravel. (Laravel seems to have increased rapidly this year...)
However, since it seems to be gaining attention, I have high hopes for the future.

# Additional Notes
If you want to use require() with React, use browserify or webpack. [What is require()?](http://qiita.com/uryyyyyyy/items/b10b012703b5396ded5a)
Since Laravel comes with browserify by default, it might be easier to use that, but choose according to your environment.

# References
* [Laravel5 ベースのプロジェクトに React が爆速で導入できた話](http://blog.mudatobunka.org/entry/2016/01/21/231546)
* [React.jsでLaravelから情報をもらってみよう](http://blog.comnect.jp.net/blog/98)
* [React入門 - Part2: Browserify/Reactify/Gulpを使う](http://qiita.com/masato/items/35b0900e3a7282b33bf8)
* [[Sy] bowerを使ってReactの開発環境を構築する方法](https://utano.jp/entry/2016/07/react-js-install-use-bower/)
