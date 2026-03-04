---
title: "モダンなJSの話──アロー関数"
slug: "modern-js-arrow-functions"
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

# アロー関数ってなに？
ざっとまとめると、

- ES2015から追加された新しい構文
- 通常のfunction式よりも短くかける
- thisの値を語彙的に束縛することができる（＝文脈からthisの値を把握しやすい）
- 常に匿名関数

アロー（=>）を使ってかく**関数式**で、"**thisの値を語彙的に束縛することができる**"という点が大きなポイントです。

アロー関数を使うと、今までこう書いていたものが・・・

```javascript
const foo = function() {
  console.log(this);
}

foo();
```

こんな感じでかけます。

```javascript
const foo = () => {
  console.log(this);
}

foo();
```

ちなみに、引数を取らない場合は丸括弧（）が必要で、引数を1個しか取らない場合は丸括弧は任意です。

```javascript
// 引数がないときは丸括弧必須
const foo = () => {
  console.log(this);
}

foo();
```

```javascript
// 引数が一つしかないときは丸括弧は任意
const foo = (value) => {
  console.log(value);
}

foo('Hello!');
```


即時関数にしたい場合はこんな感じでかけます。

```javascript
(() => {
  console.log('Hello!');
})();
```

これはちょっと混乱しそうですね・・・

# 使いドコロ
使えるところは積極的にアロー関数に置き換えていく方針で良いかと思いますが、thisが何を指すのかだけは意識しておいたほうがいいです。

例えば、下のようなケースの場合はどうでしょうか。


```javascript
const objA = {
  value: 'foo! foo!',
  sayHi: function() {
    console.log(this.value);
  }
}

objA.sayHi();

const objB = {
  value: 'bar! bar!',
  sayHi: () => {
    console.log(this.value);
  }
}

objB.sayHi();
```

1つ目のthisはオブジェクト内のvalueを、2つ目のthisはグローバルオブジェクトを返します。
このようなケースを見ると、function式とアロー関数式を使い分ける必要性があるケースもいくつかあるような気がします。

JavaScriptのthisの詳しい話は[MDN - this](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Operators/this)をご参考ください。
"**thisの値が何を指すのか**"を理解しておくとよりアロー関数への理解、JSへの理解が深まります。

# まとめ
Vue.jsやらReactやらフレームワークを使っているとコード量が多くなりがちだったり、thisがあっちらこっちら散らばって何がなんやらという状態になりやすい気がします。
アロー関数を使って関数部分をシンプルに記述することができれば、コードの見通しも良くなると思います。

# 参考
- [MDN - アロー関数](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/arrow_functions)
