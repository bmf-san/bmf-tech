---
title: "モダンなJSの話──Destructuring assignment（分割代入）"
slug: "js-destructuring-assignment"
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

# 分割代入とは
分割代入とは、配列またはオブジェクトのデータをそれぞれ別個の変数に代入する式のことです。
文章ではイメージがつきにくいかと思います。
それぞれの例を見て確認してみましょう。

# 配列の分割代入
```javascript
let a, b, c;
[a, b, c] = [1, 2, 3]
console.log(a, b, c) // 1 2 3

let color = [1, 2, 3]
const [red, green, yellow] = color
console.log(red, green, yellow) // 1 2 3
```

直感的に理解できるかと思います。

分割代入時に配列から取り出した要素がundefinedだった場合の既定値を設定することもできます。

```javascript
const [red=4, green=5, yellow=6] = [1, 2] // yellowがundefinedの場合
console.log(red, green, yellow) // 1, 2, 6 
```

引数のデフォルト値を指定するような感じですね。


# オブジェクトの分割代入
```javascript
({a, b} = {a:'foo', b:'bar'}) // 分割代入により、aという変数にaにfooが、bという変数にbarが格納される
console.log(a, b) // foo bar
```

代入文の周りの(..)については以下の引用文をご参照ください。


> 代入文の周りの ( .. ) は宣言のないオブジェクトリテラル分割代入を使用するときに必要な構文です。<br>
>{a, b} = {a:1, b:2} は有効なスタンドアロンの構文ではありません。というのも、左辺の {a, b} はブロックでありオブジェクトリテラルではないと考えられるからです。<br>
>しかしながら、({a, b} = {a:1, b:2}) 形式は有効です。var {a, b} = {a:1, b:2} と考えられるためです。<br>
> [分割代入 - JavaScript | MDN](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment)より引用

オブジェクトの分割代入は、Reactなどでよくこんな感じの使われ方をしているのを見ます。

```javascript
let state = {
    value: 'foo'
}
const {value} = state // 分割代入により、valueという変数にstate.valueが格納される
console.log(value) // foo
```

直感的にかくと、こんな感じです。

```javascript
const {value} = {value: 'foo'}
console.log(value) // foo
```

オブジェクトの分割代入も規定値を指定することができます。

```javascript
const {foo=3, bar=4} = {foo: 1} // barがundefinedの場合
console.log(foo, bar) // 1, 4
```

さらに、別の名前の変数へ値を代入することもできます。
```javascript
const {value: value2} = {value: 'foo'} // valueという変数から値を取り出してvalue2という変数に値を代入
console.log(value2) // foo
```

初見で見ると`const {value} = state`がなんのこっちゃという感じですが、分割代入を知っていると理解できますね！
便利でよく使うので覚えておくと幸せになれるかもしれません。

# まとめ
JavaScriptの分割代入について、コード例を中心に説明しました。
直感的に理解しやすい分野かと思いますので、積極的に使っていきたいですね！

# 参考リンク
[MDN - 分割代入](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment)
