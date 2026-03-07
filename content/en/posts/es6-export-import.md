---
title: About ES6 Export and Import
slug: es6-export-import
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - JavaScript
  - ES6
description: Exploring the usage of export and import in ES6.
translation_key: es6-export-import
---

I researched the parts of ES6 export and import that I hadn't fully understood.

## How to Use `export`
> The `export` statement is used to export functions, objects, or primitives from a given file (or module). Source: [MDN - Export](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/export)

Here, "export" can be thought of as defining something.

There are two ways to export.

### Named Export

```js
export { hogeFunction }; // Export a declared function
export const hoge = 1; // Export a constant. `let` or `var` can also be used.
```

You can also export using `from`.

```js
export * from 'Hoge'; // Wildcard
export { hoge, moge, huge } from 'hogemogehuge'; // Multiple exports
export { importHoge as hoge, importMoge as moge } from 'hogemoge'; // Alias
```

### Default Export

```js
export default function() {}

export default class() {}
```

`default` means "**if nothing specific is specified during import, this class or function will be called**." If you want to call a class or function other than the default during import, specify the class or file name in `{}`.

## How to Use `import`
> The `import` statement is used to import functions, objects, or primitives that have been exported from an external module or another script. Source: [MDN - Import](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/import)

There are also two ways to import.

### Named Import

```js
import * as hoge from "Hoge"; // Wildcard
import {hoge} from "Hoge"; // Import only one specified item
import {hoge, moge} from "HogeMoge"; // Import multiple items
import {hogeHoge as aliasHoge} from "HogeHoge"; // *1
import {hogeHoge as aliasHoge, mogeMoge as aliasMoge} from "MogeMoge"; // *2
import "Hoge"; // Import all modules

// *1 Import all modules and specify some members.
// *2 Specify member names to import.
```

Strictly speaking, scope-related topics are involved, but please refer to the reference sites for more details.

### Default Import

```js
import hoge from "Hoge"; // The default member is called.
```

### Note
**It will cause an error if you try to import a default-specified member with a named import.**

## Thoughts
I feel like I still have a lot to catch up on with modern JavaScript, so I need to study more... (´・ω・`)

## References
* [MDN - Export](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/export)
* [MDN - Import](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/import)
* [【Q&A】ReactJSのComponentをimport,exportする](http://qiita.com/HIGAX/items/28f3bec814928b7395da)