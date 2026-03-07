---
title: モダンなJSの話──Proxy
slug: modern-js-proxy
date: 2018-02-28T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - ES5
  - ES6
  - JavaScript
translation_key: modern-js-proxy
---


※この記事は[Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/)で掲載されている記事を転載したものです。

# Proxyとは
ProxyはECMAScript 2015から追加されたオブジェクトで、オブジェクトが持つ機能をラップすることにより、オブジェクトの機能をカスタマイズすることができます。

# Proxyの関連用語

Proxyを知る上で必要な用語です。

**handler**
・・・トラップを入れるためのオブジェクトで、プレースホルダ的な扱いをされます。

**trap**
・・・Proxyがプロパティへのアクセスを実装するためのメソッド。

**target**
・・・プロキシするオブジェクト。

**invariant**
・・・オブジェクトの機能をカスタマイズした時に変更しない不変的な条件のこと。

# Proxyの使い方

 基本的な構文は以下の通りです。

```javascript
let proxy = new Proxy(target, handler);
```

`target`にはラップしたいオブジェクトまたは関数を定義します。

`handler`には関数をプロパティとして持つオブジェクトを定義します。オブジェクトに定義された関数がProxyの操作が行われた時の挙動となります。
`handler`の中でラップ前に本来の挙動を呼び出したい時は`Reflect`というオブジェクトを呼び出します。

# 例

オブジェクトに渡された値にバリデーションをかけるような簡単な例です。

```javascript
const handler = {
  get: function(target, prop) {
    if (target[prop] === 'foo') {
      return target[prop];
    }

    return 'Default Value';
  }
};

const proxy = new Proxy({}, handler);
proxy.foo = 'foo';
proxy.bar = 'bar';

console.log(proxy.foo); // foo
console.log(proxy.bar); // Default Value
```

getトラップが実装されたhandlerオブジェクトを定義し、Proxyオブジェクトである`{}（空のオブジェクト）`がオブジェクトのプロパティを取得した際に、取得された値によって条件分岐するという処理になっています。

# 所感
JavaScriptの仕様で用意されていないオブジェクトを独自に実装したい時や本来のオブジェクトの挙動を変更したいときなどに便利そうですね。


#参考
- [MDN - Proxy](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/Proxy)
