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
description: Rebuilding the environment for React with npm and ES6 for better future compatibility.
translation_key: laravel-react-es6-browserify
---

A while ago, I wrote an article about setting up a React environment using bower. However, managing React packages with npm is smarter, and it's more beneficial to start writing in ES6 for future compatibility. So, I rebuilt the environment.

Switching from ES5 to ES6 involves some changes in syntax, which was a bit tedious, but it's not particularly difficult, so there's no need to worry too much.

# Environment
* Laravel 5.2: If you're using Laravel 5.1 or 5.2, it might be a good idea to update the elixir version to the latest (the same as 5.3).
* Browserify (included with Elixir)
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

# Refactoring ES5 to ES6
This resource is very helpful:
[Refactoring React.js source from ES5 to ES6](http://qiita.com/kuniken/items/2e850daa26a10b5098d6)

# Making ES6 compatible with ES5
Use a transpiler, such as babel.

# Thoughts
This was just a quick memo φ(..)