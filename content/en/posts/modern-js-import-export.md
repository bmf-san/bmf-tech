---
title: 'Modern JS: import and export'
description: 'An in-depth look at Modern JS: import and export, covering key concepts and practical insights.'
slug: modern-js-import-export
date: 2017-12-25T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - ES5
  - ES6
  - JavaScript
translation_key: modern-js-import-export
---



※This article is a reprint from the [Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/).

# What is export
`export` is a statement that allows you to receive functions, variables, objects, classes (classes are syntactic sugar for prototype-based inheritance and are a type of function. For more details, see [Modern JS: Classes](http://tech.innovator.jp.net/entry/2017/09/27/164750)), etc., from a specified file and use them in any file.

There are mainly two ways to use export.

#### Named exports
This method involves naming the elements you want to `export`.

```javascript
export { fooFunction };

export { fooFunction, barFunction, ... }; 

export const foo = 'bar'; 

export let foo, bar, ...;

export class foo{...};
```

You can `export` elements like this. You can also use `var` and `let` for exporting variables.

#### Default exports
This method uses the default keyword to `export` when you want to set a default element to export.

```javascript
export default fooFunction() {}

export default class {}
```

Note that `var`, `let`, and `const` cannot be used with `export default`.

# What is import
`import` is a statement that allows you to load functions, variables, and objects exported from another file and use them.

```javascript
import { foo } from "Foo";
import { foo, bar } from "FooBar";
import { foo as bar } "Foo";  // You can specify an alias
import { foo as bar, bar as foo, ... } "FooBar";
import "FooBar"; // Import everything
```

Regarding the scope of `import`ed elements, they are generally in the current scope (local scope).

#### How to import default elements defined with export default

To simply call the default, do it like this.
```javascript
import fooDefault from "Bar";
```

If you want to `import` named elements together, define them after the `default import`.

```javascript
import fooDefault, { foo, bar } "FooBar";
```

# Example of exporting/importing a class
When `export`ing a class, don't forget to call `new` in the `import` destination or `export` destination.

An example of calling `new` in the `import` destination is like this.

export.js
```javascript
export class foo {
  fooFunction() {
     return 'foo'; 
  }
}


export default class bar {
  barFunction() {
     return 'bar';
  }
}
```

import.js
```javascript
import { foo } from 'export'; // Without {}, the default bar will be called
import bar from 'export';

const objFoo = new foo;
const objBar = new bar;

console.log(objFoo.fooFunction()); // foo
console.log(objBar.barFunction()); // bar
```

<br/>

If you call `new` in the calling source, it looks like this.

export.js
```javascript
class foo {
  fooFunction() {
     return 'foo';
  }
}

function createFoo() {
  return new foo();
}

export default createFoo;
```

import.js
```javascript
import createFoo from 'export';

console.log(createFoo.fooFunction()); // foo
```

# Summary
It is good to thoroughly understand the specifications once more, as they are commonly used in recent frameworks like Vue.js and React.

# References
- [MDN - Export](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/export)
- [MDN - Import](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/import)