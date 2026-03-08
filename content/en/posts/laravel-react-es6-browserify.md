---
title: Laravel+React+ES6+Browserify
slug: laravel-react-es6-browserify
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - React
  - JavaScript
  - npm
  - webpack
  - ES6
  - browserify
  - ES5
translation_key: laravel-react-es6-browserify
---


I previously wrote about setting up a React environment using bower, but managing React packages with npm is smarter, and it's better to be able to write in ES6 for future benefits. So, I rebuilt the environment.

Since the way of writing changes quite a bit from ES5 to ES6, making those modifications was a bit of a hassle, but it doesn't seem too difficult, so there's no need to feel overwhelmed.

# Environment
* Laravel 5.2 ... If you're using 5.1 or 5.2, it might be a good idea to upgrade the version of Elixir to the latest (the same as 5.3).
* Browserify (the one that comes with Elixir)
* React
* ES6

# Setting up React with npm
`npm i react react-dom -D`

# Compilation
```gulpfile.js
elixir(function(mix) {
  mix
    .browserify('hoge.js', 'hogehoge.js')
});
```

# Modifying ES5 to ES6
This is very helpful.
[Rewrite React.js source from ES5 to ES6](http://qiita.com/kuniken/items/2e850daa26a10b5098d6)

# Making ES6 compatible with ES5
Use a transpiler, like Babel.

# Thoughts
This was just a memo.