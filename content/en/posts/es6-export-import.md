---
title: Understanding ES6 Export and Import
description: An in-depth look at Understanding ES6 Export and Import, covering key concepts and practical insights.
slug: es6-export-import
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - JavaScript
  - ES6
translation_key: es6-export-import
---

I realized that I didn't fully grasp the export and import in ES6, so I did some research.

## How to Use Export
> The export statement is used to export functions, objects, or primitives from a specified file (or module). Source: [MDN - Export](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/export)

Here, export is close to the meaning of defining something.

There are two types of exports.

### Named Exports

```js
export { hogeFunction }; // Exporting a declared function
export const hoge = 1; // Exporting a constant; let and var are also allowed.
```

You can also export using from.

```js
export * from 'Hoge'; // Wildcard
export { hoge, moge, huge } from 'hogemogehuge'; // Multiple exports
export { importHoge as hoge, importMoge as moge } from 'hogemoge'; // Alias
```

### Default Exports

```js
export default function() {}

export default class() {}
```

default means that "**if nothing is specified during import, that class or function will be called**." If you want to call a class or function other than default during import, you specify the class or file name in {}.

## How to Use Import
> The import statement is used to import functions, objects, or primitives that have been exported from external modules or other scripts. Source: [MDN - Import](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/import)

There are also two types of imports.

### Named Imports

```js
import * as hoge from "Hoge"; // Wildcard
import {hoge} from "Hoge"; // Importing only the specified one
import {hoge, moge} from "HogeMoge"; // Multiple imports
import {hogeHoge as aliasHoge} from "HogeHoge"; // *1
import {hogeHoge as aliasHoge, mogeMoge as aliasMoge} from "MogeMoge"; // *2
import "Hoge"; // Importing the entire module

// *1 Importing the entire module and specifying some members.
// *2 Specifying member names to import.
```

Strictly speaking, this involves scope, but please refer to the reference site for that.

### Default Import

```js
import hoge from "Hoge"; // The default member is called.
```

### Note
**Importing a default member with a named import will result in an error.**

## Thoughts
I feel like I still need to catch up on modern JavaScript, so I need to study more.. (´・ω・`)

## References
* [MDN - Export](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/export)
* [MDN - Import](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/import)
* [Q&A: Importing and Exporting ReactJS Components](http://qiita.com/HIGAX/items/28f3bec814928b7395da)