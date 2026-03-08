---
title: 'Modern JS Talk: Destructuring Assignment'
slug: modern-js-destructuring
date: 2017-12-25T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - ES5
  - ES6
  - JavaScript
translation_key: modern-js-destructuring
---



※This article is a reprint from the [Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/).

# What is Destructuring Assignment
Destructuring assignment is an expression that assigns data from arrays or objects to separate variables. It might be hard to visualize just from the text. Let's look at some examples to understand better.

# Array Destructuring
```javascript
let a, b, c;
[a, b, c] = [1, 2, 3]
console.log(a, b, c) // 1 2 3

let color = [1, 2, 3]
const [red, green, yellow] = color
console.log(red, green, yellow) // 1 2 3
```

You can intuitively understand it.

You can also set default values for elements extracted from the array that are undefined during destructuring.

```javascript
const [red=4, green=5, yellow=6] = [1, 2] // when yellow is undefined
console.log(red, green, yellow) // 1, 2, 6 
```

It's like specifying default values for function arguments.


# Object Destructuring
```javascript
({a, b} = {a:'foo', b:'bar'}) // Destructuring assigns 'foo' to variable a and 'bar' to variable b
console.log(a, b) // foo bar
```

Please refer to the following quote regarding the surrounding (..) in the assignment statement.


> The ( .. ) around the assignment statement is necessary syntax when using object literal destructuring without a declaration.<br>
>{a, b} = {a:1, b:2} is not valid standalone syntax because the left-hand {a, b} is considered a block, not an object literal.<br>
>However, ({a, b} = {a:1, b:2}) is valid because it can be considered as var {a, b} = {a:1, b:2}.<br>
> [Destructuring Assignment - JavaScript | MDN](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment)

Object destructuring is often used like this in React and other frameworks.

```javascript
let state = {
    value: 'foo'
}
const {value} = state // Destructuring assigns state.value to the variable value
console.log(value) // foo
```

Intuitively, it looks like this.

```javascript
const {value} = {value: 'foo'}
console.log(value) // foo
```

You can also specify default values for object destructuring.

```javascript
const {foo=3, bar=4} = {foo: 1} // when bar is undefined
console.log(foo, bar) // 1, 4
```

Furthermore, you can assign values to variables with different names.
```javascript
const {value: value2} = {value: 'foo'} // Extracts value and assigns it to value2
console.log(value2) // foo
```

At first glance, `const {value} = state` might seem confusing, but once you know destructuring, it makes sense! It's convenient and frequently used, so it's worth remembering.

# Conclusion
We explained JavaScript's destructuring assignment with code examples. It's an area that's easy to understand intuitively, so let's actively use it!

# Reference Links
[MDN - Destructuring Assignment](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment)
