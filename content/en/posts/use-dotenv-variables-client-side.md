---
title: Using dotEnv Environment Variables on the Client Side - dotenv-webpack
slug: use-dotenv-variables-client-side
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - JavaScript
  - npm
  - webpack
translation_key: use-dotenv-variables-client-side
---

dotenv-webpack is a useful library when you want to use environment variables prepared in a `.env` file on the client side.

# Introduction
`npm install dotenv-webpack --save-dev`

Set it as a plugin in `webpack.config.js`.

```
const Dotenv = require('dotenv-webpack');

module.exports = [
  ~~~ゴニョゴニョゴニョ~~~
  {
    plugins: [new Dotenv({
        path: 'path/to/.env',
        safe: false
      })]
  }
  ~~~ゴニョゴニョゴニョ~~~
];

```

`path` sets the path to the `.env` file, and `safe` determines whether to load `.env_example`.

# Usage

```:.env
DOMAIN=hereisyourdomain
```

```js:hogehoge.js
config.log(process.env.DOMAIN) // hereisyourdomain
```
# Thoughts
It's convenient, but is there no security issue?

# References
- [npm - dotenv-webpack](https://www.npmjs.com/package/dotenv-webpack)
- [github - dotEnv](https://github.com/bkeepers/dotenv)
