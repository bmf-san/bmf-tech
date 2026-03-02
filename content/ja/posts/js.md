---
title: "モダンなJSの話──クラス"
slug: "js"
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

# ECMAScript6からのクラス定義
ECMAScript6以前ではnew演算子やprototypeプロパティを使ってクラスに近い機能を実現していましたが、ECMAScript6からはclassキーワードでクラスを定義できるようになりました。<br>
classキーワードはこれまでのnew演算子やprototypeプロパティによるクラス定義のシンタックスシュガーです。

classキーワードを使ったクラス定義の方法には、クラス宣言とクラス式の2種類があります。

クラス宣言によるクラス定義の例：
```javascript
class Human {
  constructor (age, name) {
      this.age = age;
      this.name = name;
  }
  
  sayAge() {
      return this.age;
  }
  
  sayName() {
      return this.name;
  }
}

const humanInstance = new Human(24, "Bob");
console.log(humanInstance.sayAge()); // 24
console.log(humanInstance.sayName()); // Bob
```

クラス式によるクラス定義の例：
```javascript
const Human = class Human {
  constructor(age, name) {
  	this.age = age;
  	this.name = name;
  }
  
  sayAge() {
  	return this.age;
  }
  
  sayName() {
  	return this.name;
  }
}

const humanInstance = new Human(24, "Bob");
console.log(humanInstance.sayAge());
console.log(humanInstance.sayName());
```

クラス式の場合、クラス名はあってもなくてもOKです。

[MDN - Strict モード](https://developer.mozilla.org/ja/docs/Web/JavaScript/Strict_mode)

##### ホイスティング（巻き上げ）
関数宣言ではホイスティングが行われますが（関数式ではされません）、class宣言やクラス式ではホイスティングされません。<br>
したがって、クラス宣言やクラス式を使う場合はクラスを呼び出す前に、呼び出すクラスを先に宣言する必要があります。

※ホイスティングの概念について→[モダンなJSの話──var/let/const](http://tech.innovator.jp.net/entry/2017/08/30/112324)

##### クラスは全てstrictモード
クラス宣言やクラス式で定義されたクラスは全てstrictモードになります。
strictモードの詳細については[MDN - Strictモード](https://developer.mozilla.org/ja/docs/Web/JavaScript/Strict_mode)をご参照ください。

##### constructorの定義は一度だけ
初期化を行うメソッドであるconstructorの定義は一度だけしか定義できません。<br>
2回以上定義された場合はSyntax Errorが返されます。

#クラスのメソッド定義
#### getとset
ゲッターとセッターはそれぞれgetキーワードとsetキーワードで定義することができます。。<br>
ゲッターはプロパティアクセスされたときに実行されるメソッドで、セッターはプロパティに値が代入された時に実行されるメソッドです。<br>

getとsetの例：
```javascript
class Human {
  constructor(age, name) {
    this.age = age;
    this.name = name;
  }
  
  get echoProp() {
  	return `Age: ${this.age} Name: ${this.name}`;
  }
  
  set prop(prop) {
  	this.age = prop.age;
    this.name = prop.name;
  }
}

const humanInstance = new Human(24, "Bob");
console.log(humanInstance.echoProp); // Age: 24 Name: Bob
humanInstance.prop = {age: 30, name: "John"};
console.log(humanInstance.echoProp); // Age: 30 Name: John
```

#### 静的メソッド
staticキーワードを用いることで静的メソッドをクラス内に定義することができます。<br>
staticメソッドは、`クラス名.静的メソッド名` で呼び出すことができます。


静的メソッドの例：
```javascript
class Human {
  constructor(age, name) {
  	this.age = age;
  	this.name = name;
  }
  
  static sayAge(humanInstance) {
  	return `I'm ${humanInstance.age} years old.`;
  }
  
  static sayName(humanInstance) {
  	return `I'm ${humanInstance.name}.`;
  }
}

const humanInstance = new Human(24, "Bob");
console.log(Human.sayAge(humanInstance)); // I'm 24 years old.
console.log(Human.sayName(humanInstance)); // I'm Bob.
```

メソッドを静的メソッドにすることはできますが、プロパティをメソッドにすることはできません。
（TypeScriptではプロパティをstaticにすることができるらしいです。）

#### 継承
継承はextendsキーワードで定義することができます。
親クラスのメソッドを呼び出したい場合は、superキーワードを使うことで呼び出し可能です。

継承の例：
```javascript
class Gorilla {
	constructor(iq) {
  	this.iq = 2; // ちょっと優秀なゴリラ。ゴリラのIQの平均は1.5~1.8。
  }
  
  speak() {
  	return 'ウホ！';
  }
}

class Human extends Gorilla {
  constructor(age, name, iq) {
  	super(); // 仕様上super()をthisを使う前にsuper()を呼び出しておく必要がある。
    this.age = age;
    this.name = name;
    this.iq = 100; // 親の値を上書き。ちなみに100は人間の平均的なIQ。
  }
  
  sayAge() {
    return this.age;
  }
  
  sayName() {
    return this.name;
  }
  
  sayIq() {
  	return this.iq;
  }
  
  speak() {
		return `ゴリラは${super.speak()}, 人間はこんにちは！`; 
  }
}

const humanInstance = new Human(24, "Bob");
console.log(humanInstance.sayAge()); // 24
console.log(humanInstance.sayName()); // Bob
console.log(humanInstance.sayIq());　// 100
console.log(humanInstance.speak()); // ゴリラはウホ！, 人間はこんにちは！
```

ちなみに、`super()`がないと`Uncaught ReferenceError: Must call super constructor in derived class before accessing 'this' or returning from derived constructor`とエラーがでます。

# まとめ
classキーワードの登場によってJavaScriptでのOOPのあり方が変わったのではないでしょうか。<br>
とはいえアクセス修飾子が実装されていなかったり、他のクラスベースの言語のOOPほど機能が充実していないため、まだ発展途上といった印象があります。

# 余談
JavaScriptのオブジェクトモデルについて知識（cf. [オブジェクトモデルの詳細](https://developer.mozilla.org/ja/docs/Web/JavaScript/Guide/Details_of_the_Object_Model)）があるとより深い理解が得られます。


# 参考
- [MDN - クラス](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Classes)
