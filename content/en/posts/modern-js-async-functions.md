---
title: 'Modern JS: async function'
slug: modern-js-async-functions
date: 2018-01-29T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - ES5
  - ES6
  - JavaScript
translation_key: modern-js-async-functions
---

※This article is a reprint from [Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/).

# What is async function
`async function` is a function that returns an **Async Function object**.

By using the keywords `async` and `await`, you can write asynchronous code more concisely than with Promises.

It was defined in the ES2017 specification.

# How to use

Using it is simple.

Just prefix your function definition with `async`.

If you define it to return a value other than a Promise, a Promise that resolves with that value will be returned.

```javascript
async function asyncFunc() {
	return 'Amazing!';
}

asyncFunc().then((result) => {
	console.log(result);
});
```

```javascript
async function asyncFuncB(text) {
	return 'Amazing!' + text;
}

asyncFuncB('Indeed!').then((result) => {
	console.log(result);
});
```

Of course, you can also return a Promise.

```javascript
async function asyncFuncC() {
	return new Promise((resolve, reject) => {
		resolve('Wonderful!');
	});
}

asyncFuncC().then((result) => {
	console.log(result);
});
```

By the way, the above can be rewritten as follows:

```javascript
async function asyncFuncC() {
	return Promise.resolve('Wonderful!')
}

asyncFuncC().then((result) => {
	console.log(result);
});
```

Additionally, within an `async` function, you can use the `await` keyword.

The `await` keyword is an operator that can pause execution until the Promise resolves.

By using the `await` keyword, you can omit the `Promise.then()~` part.

```javascript
async function awaitFunc() {
	return 'Wonderful!';
}

async function asyncFuncD() {
	let result = await awaitFunc();
	console.log(result);
}

asyncFuncD();
```

# Thoughts
By using the `async` keyword, you can create functions that return Promises concisely without having to write out Promises each time, making asynchronous processing easier to implement.

# References
- [MDN - async function](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/async_function)