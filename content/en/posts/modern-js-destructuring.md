---
title: 'Modern JS: Destructuring Assignment'
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

※This article is a repost from [Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/).

# What is Destructuring Assignment
Destructuring assignment is an expression that assigns data from an array or object to separate variables. It might be hard to visualize in text, so let's look at some examples.

# Array Destructuring
```javascript
let a, b, c;
[a, b, c] = [1, 2, 3]
console.log(a, b, c) // 1 2 3

let color = [1, 2, 3]
const [red, green, yellow] = color
console.log(red, green, yellow) // 1 2 3
```

This should be intuitively understandable.

You can also set default values for elements extracted from an array that are undefined during destructuring.

```javascript
const [red=4, green=5, yellow=6] = [1, 2] // when yellow is undefined
console.log(red, green, yellow) // 1, 2, 6 
```

It's similar to specifying default values for function arguments.

# Object Destructuring
```javascript
({a, b} = {a:'foo', b:'bar'}) // Through destructuring, 'foo' is stored in variable a and 'bar' in variable b
console.log(a, b) // foo bar
```

Refer to the following quote regarding the parentheses around the assignment statement.

> The parentheses around the assignment statement are necessary syntax when using object literal destructuring without a declaration. <br>
> {a, b} = {a:1, b:2} is not valid standalone syntax because the left side {a, b} is considered a block, not an object literal. <br>
> However, the ({a, b} = {a:1, b:2}) format is valid because it can be interpreted as var {a, b} = {a:1, b:2}. <br>
> [Destructuring Assignment - JavaScript | MDN](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment)

Object destructuring is often used in frameworks like React.

```javascript
let state = {
    value: 'foo'
}
const {value} = state // Through destructuring, state.value is stored in variable value
console.log(value) // foo
```

Intuitively, it looks like this:

```javascript
const {value} = {value: 'foo'}
console.log(value) // foo
```

You can also specify default values for object destructuring.

```javascript
const {foo=3, bar=4} = {foo: 1} // when bar is undefined
console.log(foo, bar) // 1, 4
```

Additionally, you can assign values to variables with different names.
```javascript
const {value: value2} = {value: 'foo'} // Extracting value from variable value and assigning it to variable value2
console.log(value2) // foo
```

At first glance, `const {value} = state` might seem confusing, but understanding destructuring makes it clear! It's convenient and commonly used, so remembering it might bring you happiness.

# Conclusion
This post explained JavaScript destructuring assignment with a focus on code examples. It's a field that is easy to understand intuitively, so let's actively use it!

# Reference Links
[MDN - Destructuring Assignment](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment)