---
title: Modern JS Talk──Classes
description: 'Explore ES6 class syntax, constructor definitions, getters/setters, and strict mode behavior in modern JavaScript.'
slug: modern-js-classes
date: 2017-12-25T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - ES5
  - ES6
  - JavaScript
translation_key: modern-js-classes
---

※This article is a reprint from the [Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/).

# Class Definition from ECMAScript6
Before ECMAScript6, we used the `new` operator and `prototype` property to achieve class-like functionality, but from ECMAScript6, we can define classes using the `class` keyword.<br>
The `class` keyword is syntactic sugar for class definitions using the `new` operator and `prototype` property.

There are two ways to define a class using the `class` keyword: class declarations and class expressions.

Example of class definition using class declaration:
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

Example of class definition using class expression:
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

In the case of class expressions, the class name is optional.

[MDN - Strict Mode](https://developer.mozilla.org/ja/docs/Web/JavaScript/Strict_mode)

##### Hoisting
Function declarations are hoisted (but not function expressions), and class declarations and expressions are not hoisted.<br>
Therefore, when using class declarations or expressions, you need to declare the class before calling it.

※About the concept of hoisting→[Modern JS Talk──var/let/const](http://tech.innovator.jp.net/entry/2017/08/30/112324)

##### All Classes are in Strict Mode
All classes defined with class declarations or expressions are in strict mode.
For more details on strict mode, please refer to [MDN - Strict Mode](https://developer.mozilla.org/ja/docs/Web/JavaScript/Strict_mode).

##### Constructor Definition is Once Only
The constructor, which is a method for initialization, can only be defined once.<br>
If defined more than once, a Syntax Error is returned.

# Class Method Definitions
#### Get and Set
Getters and setters can be defined with the `get` and `set` keywords, respectively.<br>
A getter is a method executed when a property is accessed, and a setter is a method executed when a value is assigned to a property.<br>

Example of get and set:
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

#### Static Methods
Static methods can be defined within a class using the `static` keyword.<br>
Static methods can be called with `ClassName.staticMethodName`.

Example of static methods:
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

You can make methods static, but you cannot make properties static.
(It seems that in TypeScript, properties can be made static.)

#### Inheritance
Inheritance can be defined with the `extends` keyword.
If you want to call a method from the parent class, you can do so using the `super` keyword.

Example of inheritance:
```javascript
class Gorilla {
	constructor(iq) {
  	this.iq = 2; // A slightly smart gorilla. The average IQ of a gorilla is 1.5~1.8.
  }
  
  speak() {
  	return 'ウホ！';
  }
}

class Human extends Gorilla {
  constructor(age, name, iq) {
  	super(); // By specification, you need to call super() before using this.
    this.age = age;
    this.name = name;
    this.iq = 100; // Overwriting the parent value. By the way, 100 is the average IQ of a human.
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
		return `Gorilla says ${super.speak()}, Human says hello!`; 
  }
}

const humanInstance = new Human(24, "Bob");
console.log(humanInstance.sayAge()); // 24
console.log(humanInstance.sayName()); // Bob
console.log(humanInstance.sayIq());　// 100
console.log(humanInstance.speak()); // Gorilla says ウホ！, Human says hello!
```

By the way, if `super()` is not called, an `Uncaught ReferenceError: Must call super constructor in derived class before accessing 'this' or returning from derived constructor` error will occur.

# Conclusion
With the introduction of the `class` keyword, the way OOP is done in JavaScript has changed.<br>
However, since access modifiers are not implemented and the features are not as rich as OOP in other class-based languages, it still seems to be in development.

# Side Note
Having knowledge about JavaScript's object model (cf. [Details of the Object Model](https://developer.mozilla.org/ja/docs/Web/JavaScript/Guide/Details_of_the_Object_Model)) can lead to a deeper understanding.

# References
- [MDN - Classes](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Classes)