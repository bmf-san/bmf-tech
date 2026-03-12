---
title: Modern JS Talk──async function
description: An in-depth look at Modern JS Talk──async function, covering key concepts and practical insights.
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



※This article is a reprint from the [Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/).

# What is async function
An `async function` is a function that returns an **Async Function object**.

By using the keywords `async` and `await`, you can write asynchronous processing more concisely than with Promises.

The specification was defined in ES2017.

# How to Use

It's easy to use.

Simply add the `async` keyword at the beginning of the function definition.

If you define it to return a value other than a Promise, a Promise resolved with that value will be returned.

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

Incidentally, the above can be rewritten as follows.

```javascript
async function asyncFuncC() {
	return Promise.resolve('Wonderful!')
}

asyncFuncC().then((result) => {
	console.log(result);
});
```



Also, within an `async` function, you can use the `await` keyword.

The `await` keyword is an operator that can pause the execution until the Promise's result is returned.

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

# Impressions
By using the `async` keyword, you can create functions that return Promises concisely without having to write Promises repeatedly, making asynchronous processing easier to implement.


# References
- [MDN - async function](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/async_function)
