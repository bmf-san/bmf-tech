---
title: "モダンなJSの話──Promise"
slug: "modern-js-promises"
date: 2017-12-29
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "ES5"
  - "ES6"
  - "JavaScript"
draft: false
---

※この記事は[Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/)で掲載されている記事を転載したものです。

# Promiseとは
Promiseとは・・

> The Promise object represents the eventual completion (or failure) of an asynchronous operation, and its resulting value.   [MDN - Promise](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise)

だそうです。

ざっくりまとめると、Promiseとは、**非同期処理やその結果をいい感じにしてくれるオブジェクトのこと**です

Promiseを使うと主に以下のようなメリットが得られます。

- ネストが減る
- 可読性が上がる
- 処理の結果を次の処理に渡すことができる
- 例外を補足できるようになる

Promiseについて例を上げて確認してみます。

# Promiseを使わない場合の非同期処理

Promiseを使わないコールバックを使った非同期処理の例です。

```javascript
// Promiseを使わない高階関数の例
const asyncSayHi = (greet, callback) => {
    setTimeout(function () {
    callback(greet);
  }, 1000);
};

asyncSayHi('Hello', (value) => {
	console.log(value);
});
// 出力：Hello
```

`asyncSayHi`を連続して呼び出したい時は、こんな感じのいわゆる**コールバック地獄**になってしまいます。

```javascript
// callback地獄
asyncSayHi('Hello', (value) => {
	console.log(value);
    asyncSayHi('こんにちは', (value) => {
        console.log(value);
        asyncSayHi('你好', (value) => {
    	    console.log(value);
            // callback loop is forever...
        });
    });
});
// 出力：Hello こんにちは 你好
```

# Promiseを使う場合の非同期処理
先程のコールバックを使った非同期処理の例をPromiseを使って書き直すとこんな感じになります。

```javascript
// Promiseの実装
const asyncPromiseSayHi = function (greet) {
	return new Promise((resolve, reject) => {
  	if (greet) {
    	resolve(greet);
    } else {
    	reject('挨拶してください');
    }
  })
};

// 非同期処理の実行
asyncPromiseSayHi('Hello').then((value) => {
	console.log(value);
}).catch((error) => {
	console.log(error);
});
// Hello

// 連続して非同期処理を実行
asyncPromiseSayHi('Hello').then((value) => {
	console.log(value);
  return new asyncPromiseSayHi(value);
}).then((value) => {
	console.log(value);
  return new asyncPromiseSayHi(value);
}).then((value) => {
	console.log(value);
	return new asyncPromiseSayHi(value);
}).catch((error) => {
	console.log(error);
});
// Hello Hello Hello
```

並列で処理を複数実行したい場合は、`Promise.all`というメソッドが用意されています。

```javascript
const asyncPromiseSayHi = function (greet) {
	return new Promise((resolve, reject) => {
  	if (greet) {
    	resolve(greet);
    } else {
    	reject('挨拶してください');
    }
  })
};

const asyncPromiseSalute = function (salute) {
	return new Promise((resolve, reject) => {
  	if (salute) {
    	resolve(salute);
    } else {
    	reject('敬礼してください');
    }
  });
};

// 連続して非同期処理を実行
Promise.all(['asyncPromiseSayHi', 'asyncPromiseSalute']).then((value) => {
   asyncPromiseSayHi('Hello').then((value) => {
       console.log(value);
   });
   asyncPromiseSalute('Attention').then((value) => {
       console.log(value);
   });
});
// Hello Attention
```

# 所感
コールバックのことを理解していればPromiseはさほど難しく感じないかと思います。

ここで紹介していないPromiseのメソッドはMDNで確認してみてください。

# 参考
- [MDN - Promise](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise)

