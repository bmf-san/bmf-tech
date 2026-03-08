---
title: 'Modern JS: Arrow Functions'
slug: modern-js-arrow-functions
date: 2017-12-25T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - ES5
  - ES6
  - JavaScript
translation_key: modern-js-arrow-functions
---

※This article is a repost from [Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/).

# What are Arrow Functions?
In summary:

- A new syntax introduced in ES2015
- Can be written shorter than regular function expressions
- Can lexically bind the value of `this` (i.e., makes it easier to understand the value of `this` from the context)
- Always anonymous functions

The key point is that arrow functions (=>) can **lexically bind the value of `this`**.

Using arrow functions, what used to be written like this...

```javascript
const foo = function() {
  console.log(this);
}

foo();
```

Can now be written like this:

```javascript
const foo = () => {
  console.log(this);
}

foo();
```

By the way, when there are no arguments, parentheses () are required, and when there is only one argument, parentheses are optional.

```javascript
// Parentheses are required when there are no arguments
const foo = () => {
  console.log(this);
}

foo();
```

```javascript
// Parentheses are optional when there is only one argument
const foo = (value) => {
  console.log(value);
}

foo('Hello!');
```

If you want to create an immediately invoked function, you can write it like this:

```javascript
(() => {
  console.log('Hello!');
})();
```

This might be a bit confusing...

# Where to Use
I think it's a good policy to actively replace where you can with arrow functions, but you should be aware of what `this` refers to.

For example, what about in the case below?

```javascript
const objA = {
  value: 'foo! foo!',
  sayHi: function() {
    console.log(this.value);
  }
}

objA.sayHi();

const objB = {
  value: 'bar! bar!',
  sayHi: () => {
    console.log(this.value);
  }
}

objB.sayHi();
```

In the first case, `this` refers to the value inside the object, while in the second case, `this` refers to the global object. Seeing cases like this makes me feel that there are situations where it's necessary to differentiate between function expressions and arrow function expressions.

For more details on JavaScript's `this`, please refer to [MDN - this](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Operators/this). Understanding what **`this` refers to** will deepen your understanding of arrow functions and JS.

# Conclusion
When using frameworks like Vue.js or React, the amount of code tends to increase, and `this` can get scattered around, leading to confusion. If you can use arrow functions to simplify the function parts, I believe it will improve the readability of your code.

# References
- [MDN - Arrow Functions](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/arrow_functions)