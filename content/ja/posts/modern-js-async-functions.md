---
title: モダンなJSの話──async function
description: モダンなJSの話──async functionについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: modern-js-async-functions
date: 2018-01-29T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - ES5
  - ES6
  - JavaScript
translation_key: modern-js-async-functions
---


※この記事は[Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/)で掲載されている記事を転載したものです。

# async functionとは
`async function`は**Async Functionオブジェクト**を返す関数です。

`async`と`await`というキーワードを使って、Promiseよりも簡潔に非同期処理を書くことができます。

ES2017で仕様が定義されています。

# 使い方

使い方はカンタンです。

`async`関数を関数定義の時に頭につけるだけです。

Promise以外の値を返すように定義した場合は、その値で解決された形でPromiseが返されます。

```javascript
async function asyncFunc() {
	return 'すごい！';
}

asyncFunc().then((result) => {
	console.log(result);
});
```

```javascript
async function asyncFuncB(text) {
	return 'すごい！' + text;
}

asyncFuncB('さすが！').then((result) => {
	console.log(result);
});
```

もちろん、Promiseを返すこともできます。

```javascript
async function asyncFuncC() {
  return new Promise((resolve, reject) => {
    resolve('すばらしい！');
  });
}

asyncFuncC().then((result) => {
  console.log(result);
});
```

ちなみに、上は下のように書き換え可能です。

```javascript
async function asyncFuncC() {
	return Promise.resolve('すばらしい！')
}

asyncFuncC().then((result) => {
	console.log(result);
});
```



また、`async`関数内では、`await`キーワードを使うことができます。

`await`キーワードはPromiseの結果が返されるまで処理を止めることができる演算子です。

`await`キーワードを使うことで`Promise.then()~`の部分を省略して記述することができます。
　
```javascript
async function awaitFunc() {
	return 'ワンダフル！';
}

async function asyncFuncD() {
	let result = await awaitFunc();
  console.log(result);
}

asyncFuncD();
```

# 所感
Promiseを一々書かなくとも`async`キーワードを使うと簡潔にPromiseを返す関数が作れる上に、非同期処理がより実装しやすくなりましたね。


# 参考
- [MDN - async function](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/async_function)
