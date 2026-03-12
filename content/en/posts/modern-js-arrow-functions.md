---
title: 'Modern JS Talk: Arrow Functions'
description: 'An in-depth look at Modern JS Talk: Arrow Functions, covering key concepts and practical insights.'
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



※This article is a reprint from the [Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/).

# What are Arrow Functions?
In summary,

- A new syntax added in ES2015
- Shorter than regular function expressions
- Lexically binds the value of `this` (making it easier to understand the context of `this`)
- Always anonymous functions

The big point of arrow function expressions, written with `=>`, is that they "**lexically bind the value of `this`**".

With arrow functions, what used to be written like this...

```javascript
const foo = function() {
  console.log(this);
}

foo();
```

Can now be written like this.

```javascript
const foo = () => {
  console.log(this);
}

foo();
```

By the way, if there are no arguments, parentheses `()` are required, and if there is only one argument, parentheses are optional.

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

If you want to make it an immediately invoked function, you can write it like this.

```javascript
(() => {
  console.log('Hello!');
})();
```

This might be a bit confusing...

# When to Use
I think it's a good approach to actively replace where you can with arrow functions, but you should be aware of what `this` refers to.

For example, how about in the following case?


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

The first `this` returns the value within the object, while the second `this` returns the global object.
Looking at cases like this, there seem to be some situations where you need to differentiate between function expressions and arrow function expressions.

For more details on JavaScript's `this`, please refer to [MDN - this](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Operators/this).
Understanding "**what the value of `this` refers to**" will deepen your understanding of arrow functions and JS.

# Summary
When using frameworks like Vue.js or React, the code tends to get lengthy, and `this` can be scattered all over, making it hard to understand.
If you can use arrow functions to simplify the function parts, the code will be more readable.

# References
- [MDN - Arrow Functions](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/arrow_functions)
