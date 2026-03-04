---
title: "モダンなJSの話──importとexport"
slug: "modern-js-import-export"
date: 2017-12-25
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "ES5"
  - "ES6"
  - "JavaScript"
draft: false
---

※この記事は[Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/)で掲載されている記事を転載したものです。

# exportとは
`export`は、指定のファイルから関数、変数、オブジェクト、クラス（クラスはプロトタイプベース継承の糖衣構文であり、関数の一種。詳しくは [モダンなJSの話──クラス](http://tech.innovator.jp.net/entry/2017/09/27/164750)）などを受け取り、任意のファイルでそれらを使えるようにするための文です。

exportには主に2種類の使い方があります。

#### Named exports
`export`したい要素の名前を付けて`export`する方法です。

```javascript
export { fooFunction };

export { fooFunction, barFunction, ... }; 

export const foo = 'bar'; 

export let foo, bar, ...;

export class foo{...};
```

こんな感じで要素を`export`することができます。
変数の`export`は`var`、`let`も使うことができます。

#### Default exports
`export`したいデフォルトの要素を決めておきたいときにdefaultキーワードを使って`export`する方法です。

```javascript
export default fooFunction() {}

export default class {}
```

`var`、`let`、`const`は`export default`で使うことができないので注意です。

# importとは
`import`は、別ファイルからexportされた関数や変数、オブジェクトを読み込み、それらを使えるようにするための文です。

```javascript
import { foo } from "Foo";
import { foo, bar } from "FooBar";
import { foo as bar } "Foo";  // エイリアスを指定することができる
import { foo as bar, bar as foo, ... } "FooBar";
import "FooBar"; // 全てインポート
```

`import`された要素のスコープについてですが、原則的には現在のスコープ（ローカルスコープ）になります。

#### export defaultで定義されたdefaultの要素をimportする方法

シンプルにdefaultを呼び出す場合はこんな感じです。
```javascript
import fooDefault from "Bar";
```

名前付きの要素を一緒に`import`したい場合は、`default import`の後に定義します。

```javascript
import fooDefault, { foo, bar } "FooBar";
```

# classをexport/importする例
クラスを`export`する場合は`import`先または`export`先で`new`呼び出しをするのを忘れないようにしましょう。

`import`先で`new`呼び出しする例はこんな感じです。

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
import { foo } from 'export'; // {}がないとdefaultのbarが呼ばれてしまう
import bar from 'export';

const objFoo = new foo;
const objBar = new bar;

console.log(objFoo.fooFunction()); // foo
console.log(objBar.barFunction()); // bar
```

<br/>

呼び出し元で`new`呼び出ししておく場合はこんな感じです。

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

# まとめ
Vue.jsやReactといった最近のフレームワークでは当たり前のように使われているので今一度仕様をしっかりと理解しておくと良いでしょう。 

# 参考
- [MDN - Export](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/export)
- [MDN - Import](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/import)
