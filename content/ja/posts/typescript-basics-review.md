---
title: TypeScriptの基本を学び直す
slug: typescript-basics-review
date: 2024-07-20T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - TypeScript
translation_key: typescript-basics-review
---


# 概要
TypeScriptの基本について学び直す。

# JavaScriptの復習
## 変数のスコープ（scope）
### グローバルスコープ
windowオブジェクトのプロパティとして定義されるスコープ。

```javascript
const a = "Hello";
console.log(window.a); // Hello
```

###　ローカルスコープ
#### 関数スコープ
関数内で定義された変数はその関数内でのみ有効。

```javascript
function func() {
	const a = "Hello";
	console.log(a); // Hello
}
```

#### レキシカルスコープ
関数内で関数が定義された場合、内側の関数から外側の関数の変数にアクセスできる。

```javascript
function outer() {
	const a = "Hello";
	function inner() {
		console.log(a); // Hello
	}
	inner();
}
```

#### ブロックスコープ
if文やfor文などのブロック内で定義された変数はそのブロック内でのみ有効。

```javascript
if (true) {
	const a = "Hello";
	console.log(a); // Hello
}
```

## constは再代入付加だがイミュータブルではない
```typescript
const obj = { key: "value"};
obj = { key: "newValue" }; // Error
obj.key = "newValue"; // OK
```

## varの問題点
- 同名で変数宣言できてしまう
- グローバル変数を上書きしてしまう可能性がある
- 変数の巻き上げによるバグ混入のリスクがある
- スコープが広い
  - ブロックではなく関数スコープ

## ボックス化（boxing）
プリミティブ型をオブジェクト型に変換すること。

```typescript
const a = "Hello";
const aobj = new String(a);
aobj.length; // 5
```

プリミティブ型はフィールドやメソッドを持たないためボックス化が必要であるが、JavaScriptでは暗黙的にボックス化される。これを自動ボックス化（auto-boxing）という。

```typescript
const a = "Hello";
a.length; // 5
```

自動ボックス化の変換先のオブジェクトのことをラッパーオブジェクト（wrapper object）と呼ぶ。例えば、booleanならBooleanがラッパーオブジェクトとなる。undefinedとnullにはラッパーオブジェクトはない。

## オブジェクト
### プリミティブ以外は全てオブジェクト
```typescript
// プリミティブ
const num = 1;
const str = "Hello";
// etc...

// オブジェクト
const obj = { key: "value" };
const arr = [1, 2, 3];
const func = function() { return "Hello"; };
// etc...
```

## ジェネレーター
ジェネレーターは関数内でyieldを使って値を返すことができる。

```typescript
function* gen() {
	yield 1;
	yield 2;
	yield 3;
}

const g = gen();
console.log(g.next()); // { value: 1, done: false }
```

# TypeScriptの基本
## 変数宣言の型注釈（type annotation）
変数に型を付与することができる。

```typescript
const a: string = "Hello";
```

ラッパーオブジェクトを使うこともできるが、ラッパーオブジェクト型はプリミティブ型に代入できない。

```typescript
const a: Number = 0;
const b: number = a; // Type 'Number' is not assignable to type 'number'.'number' is a primitive, but 'Number' is a wrapper object. Prefer using 'number' when possible.
```

また、ラッパーオブジェクト型には演算子を使うことができない。

```typescript
const a: Number = 0;
const b = a + 1; // Operator '+' cannot be applied to types 'Number' and '1'.
```

ラッパーオブジェクト型は原則利用せず、プリミティブ型を使うことが推奨される。

## 変数宣言の型推論（type inference）
型を推論してくれる。

```typescript
let a = "Hello"; // a: string
a = 1; // Type 'number' is not assignable to type 'string'
```

## 型強制（type coercion）
型が異なる演算であってもエラーにならない場合がある。

```typescript
"10" - 1; // 9
```

型強制は暗黙的別の型へ変換する仕組みのこと。

## リテラル型
特定の値のみを取る型のこと。

```typescript
let a: "Hello" = "Hello";
a = "World"; // Type '"World"' is not assignable to type '"Hello"'.
```

リテラル型として利用できるプリミティブ型は次のとおり。
- string型
- number型
- boolean型

## any型
どんな型でも代入できる型。

```typescript
let a: any = "Hello";
a = 1; // OK
```

文脈から型推論ができない場合（ex. 型注釈を省略したとき）は暗黙的にany型として扱われる。

## オブジェクト
### オブジェクトの型注釈（type annotation）
```typescript
const obj: { key: string } = { key: "value" };
```

メソッドの型注釈も可能。

```typescript
const obj: { key: () => string } = { key: () => "value" };
```

object型もあるが、object型はプリミティブ型を除く全てのオブジェクトを表す型であるため、object型を使うことは推奨されない。また、object型は型安全が保証されない。

```typescript
const obj: object = { key: "value" };
obj.key; // Property 'key' does not exist on type 'object'.
```

### オブジェクト型のreadonly
プロパティを読み取り専用するための修飾子。

```typescript
const obj: { readonly key: string } = { key: "value" };
obj.key = "newValue"; // Cannot assign to 'key' because it is a read-only property.
```

ひとまとめで書くこともできる。

```typescript
const obj: Readonly<{
	foo: string;
	bar: number;
};>
```

### オブジェクト型のオプションプロパティ（optional property）
オブジェクトのプロパティをオプショナルにするための修飾子。

```typescript
let obj: { key?: string } = {};
obj = {} // OK
```

## nerver型
値を持たない型。

```typescript
function error(message: string): never {
	throw new Error(message);
}
```

## unknown型
any型と同じくどんな型でも代入できる型だが、unknown型は型安全が保証される。型が不明なときに利用する。

```typescript
let a: unknown = "Hello";
a = 1; // OK
const b: string = a; // Type 'unknown' is not assignable to type 'string'.
```

unknown型を使うときは、型アサーション（type assertion）やtypeof、instanceofを使って型を明示的に指定する。

```typescript
const a: unknown = "hello";

const b: string = a as string;

if (typeof a === "string") {
  const c: string = a;
}

if (a instanceof String) {
  const d: string = a as string;
}
```

## 関数
### 関数宣言の型注釈（type annotation）
```typescript
function func(a: string, b: number): string {
	return a + b;
}
```

### 関数式の型注釈
```typescript
const sayHi = function(name: string): string {
	return "Hi, " + name;
}
```

### アロー関数の型注釈
```typescript
const sayHi = (name: string): string => {
	return "Hi, " + name;
}
```

### 関数の型宣言（function type declaration）
関数の実装を省略して型だけを宣言することができる。

```typescript
type SayHi = (name: string) => string;
const sayHi: SayHi = (name) => {
	return "Hi, " + name;
}
```

メソッド構文も可能。

```typescript
type Obj = {
	sayHi: (name: string) => string;
}
```

### 型ガード関数（type guard function）
型が不明なときに型を特定する関数。

```typescript
// a is string部分はtype predicateと呼ばれる
function isString(a: unknown): a is string {
	return typeof a === "string";
}

const a: unknown = "Hello";
if (isString(a)) {
	const b: string = a;
}
```

### アサーション関数（assertion functions）
型アサーションを行う関数。

```typescript
function isString(a: unknown): asserts a is string {
	if (typeof a !== "string") {
		throw new Error("Type assertion failed.");
	}
}

const a: unknown = "Hello";
isString(a);
```

### オーバーロード関数（overload function）
同じ名前の関数に異なる型の引数を取る関数を複数定義すること。

```typescript
function add(a: number, b: number): number;
function add(a: string, b: string): string;
function add(a: any, b: any): any {
	return a + b;
}

add(1, 2); // 3
add("Hello", "World"); // HelloWorld
```

## クラス
### クラスの型注釈
```typescript
const animal: Animal = new Animal();
```

### クラスのコンストラクターの型注釈
```typescript
class Animal {
	constructor(name: string) {
		this.name = name;
	}
}
```

### クラスのメソッドの型注釈
```typescript
class Animal {
	sayHi(name: string): string {
		return "Hi, " + name;
	}
}
```

### 公称型（nominal type）
型の名前が同じでも別の型として扱うこと。

TypeScriptでは公称型をサポートしていないため、構造型（structural type）として扱われる。

```typescript
type Animal = {
	name: string;
}

type Person = {
	name: string;
}

const animal: Animal = { name: "Taro" };
```

公称型を実現するためには、構造を変える（ex. プロパティを追加する）必要がある。

### オープンエンド（open-ended）と宣言マージ（declaration merging）
オープンエンドとは、同名のインターフェースを複数定義しても重複エラーにならない性質のこと。

宣言マージとは、同名のインターフェースを複数定義すると、それらがマージされる性質のこと。

```typescript
interface Animal {
	name: string;
}

interface Animal {
	age: number;
}

const animal: Animal = { name: "Taro", age: 3 };
```

これらの性質の何が嬉しいかというと、例えばライブラリの型定義を拡張するときに有用となる。

型定義ファイルを分割することで必要な型だけをインポートすることができる。

## 型の再利用
### typeof
変数の型を取得する。

```typescript
const a = "Hello";
type A = typeof a; // type A = string;
```

### keyof
オブジェクトの型からプロパティ名を型として取得する。

```typescript
type Animal = {
  name: string;
  age: number;
};

type AnimalKey = keyof Animal; // type AnimalKey = "name" | "age";に同じ
```

### ユーティリティ型
#### Required<T>
全てのプロパティを必須にする（≒オプショナルを取り除く）。

```typescript
type Animal = {
  name?: string;
  age?: number;
};

type RequiredAnimal = Required<Animal>; // type RequiredAnimal = { name: string; age: number; };
```

#### Readonly<T>
全てのプロパティを読み取り専用にする。

```typescript
type Animal = {
  name: string;
  age: number;
};

type ReadonlyAnimal = Readonly<Animal>; // type ReadonlyAnimal = { readonly name: string; readonly age: number; };
```
#### Partial<T>
全てのプロパティをオプショナルにする。

```typescript
type Animal = {
  name: string;
  age: number;
};

type PartialAnimal = Partial<Animal>; // type PartialAnimal = { name?: string; age?: number; };
```

#### Record<Keys, Type>
プロパティのキーと値がそれぞれKeysとTypeであるオブジェクト型を生成する。

```typescript
type Name = string
type Age = number
type AnimalRecord = Record<Name, Age>; // type AnimalRecord = { [key: string]: number; };
```

#### Pick<T, Keys>
型TからKeysのプロパティを抽出する。

```typescript
type Animal = {
  name: string;
  age: number;
};

type Name = Pick<Animal, "name">; // type Name = { name: string; };
```

#### Omit<T, Keys>
型TからKeysのプロパティを除外する。

```typescript
type Animal = {
  name: string;
  age: number;
};

type Name = Omit<Animal, "age">; // type Name = { name: string; };
```

#### Exclude<T, U>
型Tから型Uで指定した型を除外したユニオン型を生成する。

```typescript
type Animal = "dog" | "cat" | "rabbit";

type ExcludeAnimal = Exclude<Animal, "dog">; // type ExcludeAnimal = "cat" | "rabbit";
```

#### Extract<T, U>
型Tから型Uで指定した型を抽出したユニオン型を生成する。

```typescript
type Animal = "dog" | "cat" | "rabbit";

type ExtractAnimal = Extract<Animal, "dog">; // type ExtractAnimal = "dog";
```

#### NoInfer<T>
型Tを推論させない。

```typescript
type Animal = {
  name: string;
  age: number;
};

function getAnimal<T>(animal: T): T {
  return animal;
}

const animal = getAnimal<NoInfer<Animal>>({ name: "Taro", age: 3 });
```

### Mapped Types
指定した型を元に新しい型を生成する。

```typescript
type Animal = {
  name: string;
  age: number;
};

type ReadonlyAnimal = {
  readonly [K in keyof Animal]: Animal[K];
};
```

### インデックスアクセス型（indexed access types）
プロパティの型や配列の要素の型を取得する。

```typescript
type Animal = {
  name: string;
  age: number;
};

type Name = Animal["name"]; // type Name = string;

type ArrayType = string[];
type ElementType = ArrayType[number]; // type ElementType = string;
```

### 条件付き型（Conditional Types）
条件に応じて型を変更する。

```typescript
type IsString<T> = T extends string ? "string" : "not string";

type A = IsString<string>; // type A = "string";
```

### infer
条件付き型の中で使われる型演算子で、型変数を取得する。

```typescript
// 型の戻り値部分を抽出するユーティリティ型
type MyReturnType<T> = T extends (...args: any[]) => infer R ? R : never;

// 関数の例
function exampleFunction(): string {
  return "Hello, World!";
}

// 関数の戻り値の型を取得
type ExampleFunctionReturnType = MyReturnType<typeof exampleFunction>;

// ExampleFunctionReturnTypeはstring型になる
const exampleReturnValue: ExampleFunctionReturnType = "This is a string";

console.log(exampleReturnValue); // This is a string
```

### ユニオン分配
ユニオン型を分配して、それぞれの型に適用する。

```typescript
type A = "a" | "b";

type B = A extends "a" ? "c" : "d"; // type B = "c" | "d";
```

## ジェネリクス
型を引数として受け取る型。

```typescript
// ジェネリクス
function identity<T>(arg: T): T {
  return arg;
}

const a = identity<string>("Hello");
const b = identity<number>(1);

// 型引数
type Identity<T> = T;
type A = Identity<string>; // type A = string;
```

# 所感
これまでJavaScriptについて何度か学んだことがあるが、TypeScriptいぜんにJavaScriptこんなに難しかったっけ...という気持ちになった。

# 参考
- [typescriptbook.jp - サバイバルTypeScript](https://typescriptbook.jp/)
  - とてもわかりやすくまとまっている
