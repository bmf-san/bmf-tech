---
title: Modern JS Talk──Promise
description: 'Master Promises for asynchronous operations, chaining with .then(), error handling, and avoiding callback hell.'
slug: modern-js-promises
date: 2017-12-29T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - ES5
  - ES6
  - JavaScript
translation_key: modern-js-promises
---



※This article is a reprint from the [Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/).

# What is a Promise
A Promise is...

> The Promise object represents the eventual completion (or failure) of an asynchronous operation, and its resulting value.   [MDN - Promise](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise)

That's what it is.

In summary, a Promise is an **object that nicely handles asynchronous operations and their results**.

Using Promises provides the following main benefits:

- Reduces nesting
- Improves readability
- Allows passing the result of one operation to the next
- Enables catching exceptions

Let's look at some examples of using Promises.

# Asynchronous Processing Without Promises

Here is an example of asynchronous processing using callbacks without Promises.

```javascript
// Example of a higher-order function without Promises
const asyncSayHi = (greet, callback) => {
    setTimeout(function () {
    callback(greet);
  }, 1000);
};

asyncSayHi('Hello', (value) => {
	console.log(value);
});
// Output: Hello
```

When you want to call `asyncSayHi` consecutively, it becomes a so-called **callback hell** like this.

```javascript
// Callback hell
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
// Output: Hello こんにちは 你好
```

# Asynchronous Processing With Promises
Rewriting the previous example of asynchronous processing using callbacks with Promises looks like this.

```javascript
// Promise implementation
const asyncPromiseSayHi = function (greet) {
	return new Promise((resolve, reject) => {
  	if (greet) {
    	resolve(greet);
    } else {
    	reject('Please greet');
    }
  })
};

// Execute asynchronous processing
asyncPromiseSayHi('Hello').then((value) => {
	console.log(value);
}).catch((error) => {
	console.log(error);
});
// Hello

// Execute asynchronous processing consecutively
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

If you want to execute multiple processes in parallel, a method called `Promise.all` is available.

```javascript
const asyncPromiseSayHi = function (greet) {
	return new Promise((resolve, reject) => {
  	if (greet) {
    	resolve(greet);
    } else {
    	reject('Please greet');
    }
  })
};

const asyncPromiseSalute = function (salute) {
	return new Promise((resolve, reject) => {
  	if (salute) {
    	resolve(salute);
    } else {
    	reject('Please salute');
    }
  });
};

// Execute asynchronous processing consecutively
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

# Thoughts
If you understand callbacks, Promises shouldn't feel too difficult.

For Promise methods not introduced here, please check MDN.

# References
- [MDN - Promise](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise)
