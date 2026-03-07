---
title: 'Modern JS: Import and Export'
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

※This article is a repost from [Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/).

# What is export
`export` is a statement that allows functions, variables, objects, and classes (classes are syntactic sugar for prototype-based inheritance and a type of function. For more details, see [Modern JS: Classes](http://tech.innovator.jp.net/entry/2017/09/27/164750)) to be used in any file by exporting them from a specified file.

There are mainly two types of usage for export.

#### Named exports
This method involves naming the elements you want to `export`.

```javascript
export { fooFunction };

export { fooFunction, barFunction, ... }; 

export const foo = 'bar'; 

export let foo, bar, ...;

export class foo{...};
```

You can export elements like this. You can use `var` and `let` for variable exports as well.

#### Default exports
This method uses the default keyword to `export` a default element you want to define.

```javascript
export default fooFunction() {}

export default class {}
```

Note that `var`, `let`, and `const` cannot be used with `export default`.

# What is import
`import` is a statement that allows you to load functions, variables, and objects exported from another file so that you can use them.

```javascript
import { foo } from "Foo";
import { foo, bar } from "FooBar";
import { foo as bar } "Foo";  // You can specify an alias
import { foo as bar, bar as foo, ... } "FooBar";
import "FooBar"; // Import everything
```

Regarding the scope of imported elements, it is generally the current scope (local scope).

#### How to import default elements defined with export default
To simply call the default, it looks like this:
```javascript
import fooDefault from "Bar";
```

If you want to `import` named elements together, define them after the `default import`.

```javascript
import fooDefault, { foo, bar } "FooBar";
```

# Example of exporting/importing classes
When exporting a class, remember to call `new` at the import or export destination.

An example of calling `new` at the import destination looks like this:

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

If you want to call `new` at the caller, it looks like this:

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

# Conclusion
Since it is commonly used in recent frameworks like Vue.js and React, it is a good idea to thoroughly understand the specifications again.

# References
- [MDN - Export](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/export)
- [MDN - Import](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/import)