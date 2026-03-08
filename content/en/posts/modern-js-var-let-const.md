---
title: 'Modern JS: var/let/const'
slug: modern-js-var-let-const
date: 2017-12-25T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - ES5
  - ES6
  - JavaScript
translation_key: modern-js-var-let-const
---

※This article is a repost from [Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/).

# What is Scope
Before diving into the main topic, let's confirm the definition of scope.

Scope refers to the **range in which variable names and function names can be referenced**.

There are various types of scopes, but here we will mainly explain three types in a table.

|Scope Name|Range|Notes|
|:--|:--|:--|
|Global|Outside of functions|Accessible from anywhere.| 
|Local (Function)|Inside functions|Accessible only from within the local scope.| 
|Block|Inside blocks ({ })|if, for, switch, etc.|

Block scope was not originally available in JavaScript, but with the introduction of let and const, block scope can now be used.

# About var, let, and const

#### var
|Re-declaration|Re-assignment|Scope|
|:--:|:--:|:--:|
|○|○|Local|

var can be both re-declared and re-assigned.

```javascript
var a = 1;
var a = 2;  // Re-declaration allowed

function sayNum()
{
  a = 100; // Re-assignment allowed
  return a;
}

console.log(sayNum()); // 100
```

Previously, var was the only way to declare variables, but with the advent of let and const, the necessity to declare variables with var has almost disappeared. Except for some special cases, performance tuning, or browser compatibility considerations, there will be almost no opportunities to use var.

##### Hoisting
var has the concept of **hoisting**, which means that **variable declarations are processed before the execution of the code**.

This means that,

```javascript
a = 1;
var a;
```

is processed like this:

```javascript
var a;
var a = 1;
```

By the way, if you try similar code with let or const...

```javascript
let a;
let a = 1; // Uncaught ReferenceError: a is not defined
```

In the case of let, a `ReferenceError` is thrown, but hoisting actually occurs. The same applies to const.

> In ECMAScript 2015, let hoists the variable to the top of the block. However, referencing that variable before its declaration will cause a ReferenceError. The variable remains in a "temporal dead zone" until the declaration is executed from the beginning of the block.

Quoted from [MDN - let](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/let)

The same applies to const.

> All concerns regarding the "temporal dead zone" for let also apply to const.

Quoted from [MDN - const](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/const)

```javascript
a = 1;
const a; // Uncaught SyntaxError: Missing initializer in const declaration
```

In the case of const, a syntax error is thrown if it is not initialized (`const a = 1;`).

Let's check hoisting for const.

```javascript
const a = 1;

function sayNum()
{
  console.log(a); // Uncaught ReferenceError: a is not defined
  const a = 100;
}

sayNum();
```

It seems that there are few scenes where hoisting is consciously used, but it is good to remember that the behavior of hoisting differs between var, let, and const.

#### let
|Re-declaration|Re-assignment|Scope|
|:--:|:--:|:--:|
|✗ (Re-declaration in the same scope)|○|Block|

let cannot be re-declared but can be re-assigned.

```javascript
let a = 1;
let a = 2; // Uncaught SyntaxError: Identifier 'a' has already been declared   => Re-declaration in the same scope is not allowed
```

```javascript
let a = 1;

function sayNum()
{
  a = 100 // Re-assignment allowed
  return a;
}

console.log(sayNum()); // 100
```

The scenes where let is used are those that previously used var, particularly in areas where **re-assignment is possible**.

#### const
|Re-declaration|Re-assignment|Scope|
|:--:|:--:|:--:|
|✗ (Re-declaration in the same scope)|✗|Block|

const cannot be re-declared or re-assigned.

```javascript
const a = 1;
const a = 2; // Uncaught SyntaxError: Identifier 'a' has already been declared => Re-declaration in the same scope is not allowed
```

```javascript
const a = 1;

function sayNum()
{
  a = 100 // Uncaught TypeError: Assignment to constant variable. => Re-assignment not allowed
  return a;
}

console.log(sayNum()); // 100
```

# Summary
|Declaration|Re-declaration|Re-assignment|Scope|
|:--:|:--:|:--:|:--:|
|var|○|○|Local|
|let|✗ (Re-declaration in the same scope)|○|Block|
|const|✗ (Re-declaration in the same scope)|✗|Block|

Basically, it seems good to declare variables using const, use let for parts where re-assignment is possible, and consider var for other special cases. Variable scope pollution can lead to bugs and hinder code readability, so we want to use them properly!

# Reference Links

- [https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/var:title]
- [https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/let:title:title]
- [https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/const:title]