---
title: "ES6のExportとImportについて"
slug: "es6-export-import"
date: 2017-09-26
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "JavaScript"
  - "ES6"
draft: false
---

ES６のexportとimportについて把握しきれていないところがあったので調べてみました。


## exportの使い方
>export 文は、指定したファイル (またはモジュール) から関数、オブジェクト、プリミティブをエクスポートするために使用 引用元：[MDN - Export](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/export)

ここでいうエクスポートとは、何かを定義するという意味合いに近いかと思います。


エクスポートには２種類の方法があります。

### 名前付きエクスポート

```js
export { hogeFunction };　 // 宣言済みの関数をエクスポート
export const hoge = 1; // 定数をエクスポート letやvarも可。
```

fromを使ってエクスポートすることもできます。

```js
export * from 'Hoge'; // ワイルドカード
export { hoge, moge, huge } from 'hogemogehuge'; // 複数エクスポート
export { importHoge as hoge, importMoge as moge } from 'hogemoge';　// エイリアス
```

### デフォルトエクスポート

```js
export default function() {}

export default class() {}
```

defaultとは、　「**importする際に特に指定がなければそのクラスや関数を呼ぶ**」というものです。
importする際にdefault以外のクラスや関数を呼び出したいときは、{}でクラスやファイル名を指定して呼び出してあげます。


## importの使い方
>import 文 は、外部モジュールや他のスクリプトなどからエクスポートされた関数、オブジェクト、プリミティブをインポートするために使用します。 引用元：[MDN - Import](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/import)

インポートにも２種類の方法があります。

###　名前指定のインポート

```js
import * as hoge from "Hoge"; // ワイルドカード
import {hoge} from "Hoge"; // 指定の1つだけをインポート
import {hoge, moge} from "HogeMoge"; // 複数インポート
import {hogeHoge as aliasHoge} from "HogeHoge"; // ※1
import {hogeHoge as aliasHoge, mogeMoge as aliasMoge} from "MogeMoge"; // ※2
import "Hoge"; // 全モジュールをインポート

//　※1 全モジュールをインポートし、さらに一部のメンバーを指定する。
// ※2 メンバー名を指定してインポートする。
```

厳密にはスコープの話が関わってきますが、そちらについては参考サイトをご確認ください。

### デフォルトのインポート

```js
import hoge from "Hoge"; // defaultのメンバーが呼び出される。
```

### 注意
**デフォルト指定されたメンバーを名前指定でインポートするとエラーになります。**

## 所感
モダンなjavascriptはまだまだキャッチアップしきれていない気がするので勉強しなくては・・(´・ω・`)


## 参考
* [MDN - Export](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/export)
* [MDN - Import](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/import)
* [【Q&A】ReactJSのComponentをimport,exportする](http://qiita.com/HIGAX/items/28f3bec814928b7395da)

