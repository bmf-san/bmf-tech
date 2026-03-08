---
title: 'Modern JS: Proxy'
slug: modern-js-proxy
date: 2018-02-28T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - ES5
  - ES6
  - JavaScript
translation_key: modern-js-proxy
---

※This article is a repost from [Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/).

# What is Proxy
Proxy is an object added in ECMAScript 2015 that allows you to customize the functionality of an object by wrapping its features.

# Related Terms for Proxy

These are terms necessary to understand Proxy.

**handler**
… An object used to insert traps, treated as a placeholder.

**trap**
… A method that implements access to properties in the Proxy.

**target**
… The object being proxied.

**invariant**
… An immutable condition that does not change when customizing the functionality of an object.

# How to Use Proxy

The basic syntax is as follows.

```javascript
let proxy = new Proxy(target, handler);
```

Define the object or function you want to wrap in `target`.

Define an object with functions as properties in `handler`. The functions defined in the object will determine the behavior when operations on the Proxy are performed. If you want to call the original behavior before wrapping in `handler`, you can call the `Reflect` object.

# Example

Here’s a simple example that applies validation to values passed to an object.

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

In this example, we define a handler object with a get trap, and the Proxy object `{}` (an empty object) processes property retrieval based on the value obtained.

# Thoughts
It seems useful when you want to implement objects that are not provided by JavaScript specifications or when you want to change the behavior of the original object.

# References
- [MDN - Proxy](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/Proxy)