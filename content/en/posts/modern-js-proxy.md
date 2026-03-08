---
title: Modern JS Discussion ─ Proxy
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



※This article is a reprint from the [Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/).

# What is Proxy
Proxy is an object added in ECMAScript 2015 that allows you to customize the functionality of an object by wrapping its features.

# Related Terms for Proxy

These are the terms necessary to understand Proxy.

**handler**
・・・An object for inserting traps, treated as a placeholder.

**trap**
・・・A method for Proxy to implement access to properties.

**target**
・・・The object being proxied.

**invariant**
・・・An unchanging condition when customizing the functionality of an object.

# How to Use Proxy

 The basic syntax is as follows:

```javascript
let proxy = new Proxy(target, handler);
```

Define the object or function you want to wrap in `target`.

Define an object with functions as properties in `handler`. The functions defined in the object determine the behavior when operations on the Proxy are performed. If you want to call the original behavior before wrapping within `handler`, call an object called `Reflect`.

# Example

Here is a simple example of applying validation to values passed to an object.

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

A handler object with a get trap is defined, and when the Proxy object `{}` (an empty object) retrieves the properties of the object, it processes by branching conditions based on the retrieved value.

# Impressions
It seems useful when you want to implement an object not provided by JavaScript specifications or when you want to change the behavior of an original object.


# References
- [MDN - Proxy](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/Proxy)