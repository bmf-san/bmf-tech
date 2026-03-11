---
title: 'Modern JS Discussion: var/let/const'
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



※This article is a reprint from the [Innovator Japan Engineers’ Blog](http://tech.innovator.jp.net/).

# What is Scope
Before diving into the main topic, let's review the definition of scope.

Scope refers to the **range within which variable names and function names can be referenced**.

There are various types of scopes, but here we will explain mainly three types of scopes in a table.

|Scope Name|Range|Remarks|
|:--|:--|:--|
|Global|Outside of functions|Accessible from anywhere.|
|Local (Function)|Inside of functions|Accessible only from within the local scope.|
|Block|Inside of blocks ({ })|if, for, switch, etc.|

Block scope was not originally available in JavaScript, but with the introduction of let and const, block scope can now be used.

# About var, let, const

#### var
|Redeclaration|Reassignment|Scope|
|:--:|:--:|:--:|
|○|○|Local|

var allows both redeclaration and reassignment.

```javascript
var a = 1;
var a = 2;  // Redeclaration possible

function sayNum()
{
  a = 100; // Reassignment possible
  return a;
}

console.log(sayNum()); // 100
```

Previously, var was the only way to declare variables, but with the introduction of let and const, there is almost no need to declare variables with var anymore. Except for some special cases, performance tuning, or browser compatibility considerations, the opportunity to use var is almost nonexistent.

##### Hoisting
var has a concept called **hoisting**, which means that **variable declarations are processed before the code execution**.

This is processed as follows:

```javascript
a = 1;
var a;
```

It is processed like this:

```javascript
var a;
var a = 1;
```

Incidentally, if you try the same code with let or const...

```javascript
let a;
let a = 1; // Uncaught ReferenceError: a is not defined
```

In the case of let, a `ReferenceError` is thrown, but hoisting itself actually occurs. This is the same for const.

>In ECMAScript 2015, let lifts variables to the top of the block. However, referencing the variable before the declaration causes a ReferenceError. The variable is in a "temporal dead zone" from the start of the block until the variable declaration is executed.

Quoted from [MDN - let](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/let)

This is the same for const.

>All concerns about the "temporal dead zone" for let apply to const as well.

Quoted from [MDN - const](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/const)

```javascript
a = 1;
const a; // Uncaught SyntaxError: Missing initializer in const declaration
```

In the case of const, a syntax error is thrown unless initialization (`const a = 1;`) is performed.

Let's check the hoisting of const.

```javascript
const a = 1;

function sayNum()
{
  console.log(a); // Uncaught ReferenceError: a is not defined
  const a = 100;
}

sayNum();
```

There seem to be few scenes where hoisting is consciously used, but it seems better to remember that the hoisting behavior differs between var, let, and const.

#### let
|Redeclaration|Reassignment|Scope|
|:--:|:--:|:--:|
|✗ (**Redeclaration in the same scope**)|○|Block|

let does not allow redeclaration but allows reassignment.

```javascript
let a = 1;
let a = 2; // Uncaught SyntaxError: Identifier 'a' has already been declared   => Redeclaration in the same scope not allowed
```

```javascript
let a = 1;

function sayNum()
{
  a = 100 // Reassignment possible
  return a;
}

console.log(sayNum()); // 100
```

The scene to use let is where var was used before, but it is particularly limited to **parts where reassignment is possible**.

#### const
|Redeclaration|Reassignment|Scope|
|:--:|:--:|:--:|
|✗ (**Redeclaration in the same scope**)|✗|Block|

const does not allow either redeclaration or reassignment.

```javascript
const a = 1;
const a = 2; // Uncaught SyntaxError: Identifier 'a' has already been declared => Redeclaration in the same scope not allowed
```

```javascript
const a = 1;

function sayNum()
{
  a = 100 // Uncaught TypeError: Assignment to constant variable. => Reassignment not allowed
  return a;
}

console.log(sayNum()); // 100
```

# Summary
|Declaration|Redeclaration|Reassignment|Scope|
|:--:|:--:|:--:|:--:|
|var|○|○|Local|
|let|✗ (**Redeclaration in the same scope**)|○|Block|
|const|✗ (**Redeclaration in the same scope**)|✗|Block|

Basically, it seems good to declare variables using const and use let for parts where reassignment is possible, and consider var for other special cases. Variable scope pollution can lead to bugs and hinder code reading, so we want to use them properly!

# Reference Links

- [https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/var:title]
- [https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/let:title:title]
- [https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/const:title]

