---
title: "dotEnvの環境変数をクライアントサイドでも使えるようにする - dotenv-webpack"
slug: "use-dotenv-variables-client-side"
date: 2017-09-26
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "JavaScript"
  - "npm"
  - "webpack"
draft: false
---

dotenv-webpackは`.env`ファイルに用意された環境変数をクライアントサイドでも利用したい時に便利なライブラリです。

# 導入
`npm install dotenv-webpack --save-dev`

`webpack.config.js`にpluginとして設定します。

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

`path`は`.env`ファイルへのパス、`safe`は`.env_example`を読み込むか否かを設定します。

# 使い方

```:.env
DOMAIN=hereisyourdomain
```

```js:hogehoge.js
config.log(process.env.DOMAIN) // hereisyourdomain
```
# 所感
便利だけどセキュリティ的なところは問題ないのだろうか？

# 参考
- [npm - dotenv-webpack](https://www.npmjs.com/package/dotenv-webpack)
- [github - dotEnv](https://github.com/bkeepers/dotenv)

