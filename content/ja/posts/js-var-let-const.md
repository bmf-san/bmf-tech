---
title: "モダンなJSの話──var/let/const"
slug: "js-var-let-const"
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

# スコープとは
本題に入る前に、スコープの定義について確認しておきましょう。


スコープとは、**変数名や関数名が参照可能な範囲のこと**です。


スコープの種類は色々ありますが、ここでは主に3つのスコープについて表で説明します。

|スコープ名|範囲|備考|
|:--|:--|:--|
|グローバル|関数の外側|どこからでもアクセスできる。|
|ローカル（関数）|関数の内側|ローカルスコープ内からでしかアクセスできない。|
|ブロック|ブロック（{ }）の内側|if, for, switchなど|


ブロックスコープはJavaScriptには元々ありませんでしたが、letとconstの登場によってブロックスコープが使えるようになりました。


# var, let, constについて

#### var
|再宣言|再代入|スコープ|
|:--:|:--:|:--:|
|○|○|ローカル|

varは再宣言も再代入もできます。

```javascript
var a = 1;
var a = 2;  // 再宣言可能

function sayNum()
{
  a = 100; // 再代入可能
  return a;
}

console.log(sayNum()); // 100
```

以前はvarしか変数宣言がありませんでしたが、後述するletやconstの登場により、varで変数を宣言する必要性はほとんどなくなりました。
一部の特殊なケースやパフォーマンスチューニング、ブラウザ対応等を考慮するような状況を除いて、varを使う機会はほぼ無いでしょう。

#####巻き上げ（hoisting）
varには**巻き上げ（hoisting）**という概念があり、**変数宣言はコードの実行よりも前に処理される**という特徴があります。

これが、

```javascript
a = 1;
var a;
```

こういうふうに処理されます。

```javascript
var a;
var a = 1;
```

ちなみに、letやconstで同様のコードを試してみると・・・

```javascript
let a;
let a = 1; // Uncaught ReferenceError: a is not defined
```

letの場合は、`ReferenceError`が投げられますが、巻き上げ自体は実は行われているようです。これについてはconstも同様です。

>ECMAScript 2015 では let は変数をブロックの先頭へ引き上げます。しかし、その変数を宣言より前で参照することは ReferenceError を引き起こします。ブロックの始めから変数宣言が実行されるまで、変数は "temporal dead zone" の中にいるのです。

[MDN - let](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/let)より引用

これについてはconstも同様です。

>let に対する "temporal dead zone" の懸念事項はすべて、const にも適用されます。

[MDN - const](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/const)より引用

```javascript
a = 1;
const a; // Uncaught SyntaxError: Missing initializer in const declaration
```

constの場合は、そもそも初期化（`const a = 1;`）をしないとシンタックスエラーが投げられます。

constの巻き上げを確認してみます。

```javascript
const a = 1;

function sayNum()
{
  console.log(a); // Uncaught ReferenceError: a is not defined
  const a = 100;
}

sayNum();
```

巻き上げを意識使うようなシーンは少ないような気がしますが、varとlet、constで巻き上げの挙動が異なることは覚えておいたほうが良さそうです。

#### let
|再宣言|再代入|スコープ|
|:--:|:--:|:--:|
|✗（**同スコープ内での再宣言**）|○|ブロック|

letは再宣言は不可ですが、再代入は可能です。

```javascript
let a = 1;
let a = 2; // Uncaught SyntaxError: Identifier 'a' has already been declared   => 同スコープ内での再宣言不可
```

```javascript
let a = 1;

function sayNum()
{
  a = 100 // 再代入可能
  return a;
}

console.log(sayNum()); // 100
```

letを使うシーンというのは、これまでのvarを使っていた部分になりますが、とりわけ**再代入する可能性がある部分**に限られます。

#### const
|再宣言|再代入|スコープ|
|:--:|:--:|:--:|
|✗（**同スコープ内での再宣言**）|✗|ブロック|

constは再宣言も再代入も不可能です。

```javascript
const a = 1;
const a = 2; // Uncaught SyntaxError: Identifier 'a' has already been declared => 同スコープ内での再宣言不可能
```

```javascript
const a = 1;

function sayNum()
{
  a = 100 // Uncaught TypeError: Assignment to constant variable. => 再代入不可能
  return a;
}

console.log(sayNum()); // 100
```

# まとめ
|宣言|再宣言|再代入|スコープ|
|:--:|:--:|:--:|:--:|
|var|○|○|ローカル|
|let|✗（**同スコープ内での再宣言**）|○|ブロック|
|const|✗（**同スコープ内での再宣言**）|✗|ブロック|

基本的にはconstを使って変数を宣言するようにして、再代入する可能性がある部分に関してはletを使う、それ以外の特殊なケースに関してはvarを検討するという方針で良さそうです。
変数のスコープ汚染はバグの元やコードリーディングの妨げになるので、しっかりと使い分けていきたいですね！

# 参考リンク

- [https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/var:title]
- [https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/let:title:title]
- [https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/const:title]






