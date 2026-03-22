---
title: Modern JS Discussion ─ Proxy
description: 'Master JavaScript Proxy objects, trap handlers, target wrapping, and value validation in ES2015 development.'
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


*This article is a reprint of an article published on the [Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/).*

# What is Proxy?
Proxy is an object added in ECMAScript 2015 that allows you to customize the functionality of an object by wrapping its capabilities.

# Related Terms for Proxy

These are the terms necessary to understand Proxy.

**handler**
...An object used to insert traps, treated as a placeholder.

**trap**
...A method that implements access to properties in the Proxy.

**target**
...The object being proxied.

**invariant**
...An unchanging condition that remains unchanged when customizing the functionality of an object.

# How to Use Proxy

The basic syntax is as follows.

```javascript
let proxy = new Proxy(target, handler);
```

In `target`, define the object or function you want to wrap.

In `handler`, define an object that has functions as properties. The functions defined in the object will dictate the behavior when operations on the Proxy are performed. If you want to call the original behavior before wrapping in `handler`, you can call the `Reflect` object.

# Example

Here is a simple example that applies validation to values passed to an object.

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

This defines a handler object with an implemented get trap, and when the Proxy object `{}` (an empty object) retrieves the properties of the object, it processes conditionally based on the retrieved value.

# Thoughts
It seems useful when you want to implement objects that are not provided in the JavaScript specification or when you want to change the behavior of the original object.

# References
- [MDN - Proxy](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Proxy)
