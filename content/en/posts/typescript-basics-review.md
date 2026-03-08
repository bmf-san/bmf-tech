---
title: Revisiting the Basics of TypeScript
slug: typescript-basics-review
date: 2024-07-20T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - TypeScript
translation_key: typescript-basics-review
---

# Overview
Revisiting the basics of TypeScript.

# JavaScript Review
## Variable Scope
### Global Scope
Scope defined as a property of the window object.

```javascript
const a = "Hello";
console.log(window.a); // Hello
```

### Local Scope
#### Function Scope
Variables defined within a function are only valid within that function.

```javascript
function func() {
	const a = "Hello";
	console.log(a); // Hello
}
```

#### Lexical Scope
If a function is defined within another function, the inner function can access the variables of the outer function.

```javascript
function outer() {
	const a = "Hello";
	function inner() {
		console.log(a); // Hello
	}
	inner();
}
```

#### Block Scope
Variables defined within a block, such as an if statement or for loop, are only valid within that block.

```javascript
if (true) {
	const a = "Hello";
	console.log(a); // Hello
}
```

## const is Reassignable but Not Immutable
```typescript
const obj = { key: "value"};
obj = { key: "newValue" }; // Error
obj.key = "newValue"; // OK
```

## Issues with var
- Allows declaration of variables with the same name
- Can overwrite global variables
- Risk of bugs due to variable hoisting
- Broad scope
  - Function scope instead of block scope

## Boxing
Converting a primitive type to an object type.

```typescript
const a = "Hello";
const aobj = new String(a);
aobj.length; // 5
```

Since primitive types do not have fields or methods, boxing is necessary, but in JavaScript, boxing happens implicitly. This is called auto-boxing.

```typescript
const a = "Hello";
a.length; // 5
```

The object resulting from auto-boxing is called a wrapper object. For example, for boolean, the wrapper object is Boolean. There are no wrapper objects for undefined and null.

## Objects
### Everything Except Primitives is an Object
```typescript
// Primitive
const num = 1;
const str = "Hello";
// etc...

// Object
const obj = { key: "value" };
const arr = [1, 2, 3];
const func = function() { return "Hello"; };
// etc...
```

## Generators
Generators can return values using yield within a function.

```typescript
function* gen() {
	yield 1;
	yield 2;
	yield 3;
}

const g = gen();
console.log(g.next()); // { value: 1, done: false }
```

# Basics of TypeScript
## Type Annotation for Variable Declaration
You can assign a type to a variable.

```typescript
const a: string = "Hello";
```

You can also use wrapper objects, but wrapper object types cannot be assigned to primitive types.

```typescript
const a: Number = 0;
const b: number = a; // Type 'Number' is not assignable to type 'number'.'number' is a primitive, but 'Number' is a wrapper object. Prefer using 'number' when possible.
```

Additionally, you cannot use operators with wrapper object types.

```typescript
const a: Number = 0;
const b = a + 1; // Operator '+' cannot be applied to types 'Number' and '1'.
```

It is recommended to use primitive types instead of wrapper object types.

## Type Inference for Variable Declaration
Type inference is provided.

```typescript
let a = "Hello"; // a: string
a = 1; // Type 'number' is not assignable to type 'string'
```

## Type Coercion
There are cases where operations with different types do not result in an error.

```typescript
"10" - 1; // 9
```

Type coercion is the mechanism that implicitly converts to another type.

## Literal Types
A type that can only take specific values.

```typescript
let a: "Hello" = "Hello";
a = "World"; // Type '"World"' is not assignable to type '"Hello"'.
```

The primitive types that can be used as literal types are:
- string
- number
- boolean

## any Type
A type that can accept any value.

```typescript
let a: any = "Hello";
a = 1; // OK
```

When type inference cannot be made from context (e.g., when type annotation is omitted), it is treated as any type implicitly.

## Objects
### Type Annotation for Objects
```typescript
const obj: { key: string } = { key: "value" };
```

Method type annotations are also possible.

```typescript
const obj: { key: () => string } = { key: () => "value" };
```

There is also an object type, but it represents all objects excluding primitive types, so using the object type is not recommended. Additionally, the object type does not guarantee type safety.

```typescript
const obj: object = { key: "value" };
obj.key; // Property 'key' does not exist on type 'object'.
```

### Readonly for Object Types
A modifier to make properties read-only.

```typescript
const obj: { readonly key: string } = { key: "value" };
obj.key = "newValue"; // Cannot assign to 'key' because it is a read-only property.
```

You can also write it all together.

```typescript
const obj: Readonly<{
	foo: string;
	bar: number;
};>
```

### Optional Properties for Object Types
A modifier to make object properties optional.

```typescript
let obj: { key?: string } = {};
obj = {} // OK
```

## never Type
A type that has no value.

```typescript
function error(message: string): never {
	throw new Error(message);
}
```

## unknown Type
Like any type, it can accept any value, but the unknown type guarantees type safety. It is used when the type is unknown.

```typescript
let a: unknown = "Hello";
a = 1; // OK
const b: string = a; // Type 'unknown' is not assignable to type 'string'.
```

When using the unknown type, you should explicitly specify the type using type assertion, typeof, or instanceof.

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

## Functions
### Type Annotation for Function Declarations
```typescript
function func(a: string, b: number): string {
	return a + b;
}
```

### Type Annotation for Function Expressions
```typescript
const sayHi = function(name: string): string {
	return "Hi, " + name;
}
```

### Type Annotation for Arrow Functions
```typescript
const sayHi = (name: string): string => {
	return "Hi, " + name;
}
```

### Function Type Declaration
You can declare only the type without implementing the function.

```typescript
type SayHi = (name: string) => string;
const sayHi: SayHi = (name) => {
	return "Hi, " + name;
}
```

Method syntax is also possible.

```typescript
type Obj = {
	sayHi: (name: string) => string;
}
```

### Type Guard Functions
Functions that specify the type when it is unknown.

```typescript
// a is string part is called type predicate
function isString(a: unknown): a is string {
	return typeof a === "string";
}

const a: unknown = "Hello";
if (isString(a)) {
	const b: string = a;
}
```

### Assertion Functions
Functions that perform type assertions.

```typescript
function isString(a: unknown): asserts a is string {
	if (typeof a !== "string") {
		throw new Error("Type assertion failed.");
	}
}

const a: unknown = "Hello";
isString(a);
```

### Overload Functions
Defining multiple functions with the same name that take different types of arguments.

```typescript
function add(a: number, b: number): number;
function add(a: string, b: string): string;
function add(a: any, b: any): any {
	return a + b;
}

add(1, 2); // 3
add("Hello", "World"); // HelloWorld
```

## Classes
### Type Annotation for Classes
```typescript
const animal: Animal = new Animal();
```

### Type Annotation for Class Constructors
```typescript
class Animal {
	constructor(name: string) {
		this.name = name;
	}
}
```

### Type Annotation for Class Methods
```typescript
class Animal {
	sayHi(name: string): string {
		return "Hi, " + name;
	}
}
```

### Nominal Types
Types with the same name are treated as different types.

TypeScript does not support nominal types, so they are treated as structural types.

```typescript
type Animal = {
	name: string;
}

type Person = {
	name: string;
}

const animal: Animal = { name: "Taro" };
```

To achieve nominal types, you need to change the structure (e.g., add properties).

### Open-ended and Declaration Merging
Open-ended means that defining multiple interfaces with the same name does not result in duplication errors.

Declaration merging means that defining multiple interfaces with the same name merges them.

```typescript
interface Animal {
	name: string;
}

interface Animal {
	age: number;
}

const animal: Animal = { name: "Taro", age: 3 };
```

The benefit of these properties is that they are useful when extending type definitions for libraries.

By splitting type definition files, you can import only the necessary types.

## Type Reuse
### typeof
Get the type of a variable.

```typescript
const a = "Hello";
type A = typeof a; // type A = string;
```

### keyof
Get property names as types from the type of an object.

```typescript
type Animal = {
  name: string;
  age: number;
};

type AnimalKey = keyof Animal; // type AnimalKey = "name" | "age";
```

### Utility Types
#### Required<T>
Makes all properties required (removes optional).

```typescript
 type Animal = {
  name?: string;
  age?: number;
};

type RequiredAnimal = Required<Animal>; // type RequiredAnimal = { name: string; age: number; };
```

#### Readonly<T>
Makes all properties read-only.

```typescript
 type Animal = {
  name: string;
  age: number;
};

type ReadonlyAnimal = Readonly<Animal>; // type ReadonlyAnimal = { readonly name: string; readonly age: number; };
```

#### Partial<T>
Makes all properties optional.

```typescript
 type Animal = {
  name: string;
  age: number;
};

type PartialAnimal = Partial<Animal>; // type PartialAnimal = { name?: string; age?: number; };
```

#### Record<Keys, Type>
Generates an object type where the property keys and values are of Keys and Type respectively.

```typescript
 type Name = string
 type Age = number
 type AnimalRecord = Record<Name, Age>; // type AnimalRecord = { [key: string]: number; };
```

#### Pick<T, Keys>
Extracts properties from type T based on Keys.

```typescript
 type Animal = {
  name: string;
  age: number;
};

type Name = Pick<Animal, "name">; // type Name = { name: string; };
```

#### Omit<T, Keys>
Excludes properties from type T based on Keys.

```typescript
 type Animal = {
  name: string;
  age: number;
};

type Name = Omit<Animal, "age">; // type Name = { name: string; };
```

#### Exclude<T, U>
Generates a union type excluding the type specified by U from type T.

```typescript
 type Animal = "dog" | "cat" | "rabbit";

type ExcludeAnimal = Exclude<Animal, "dog">; // type ExcludeAnimal = "cat" | "rabbit";
```

#### Extract<T, U>
Generates a union type extracting the type specified by U from type T.

```typescript
 type Animal = "dog" | "cat" | "rabbit";

type ExtractAnimal = Extract<Animal, "dog">; // type ExtractAnimal = "dog";
```

#### NoInfer<T>
Prevents type inference for type T.

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
Generates a new type based on the specified type.

```typescript
 type Animal = {
  name: string;
  age: number;
};

type ReadonlyAnimal = {
  readonly [K in keyof Animal]: Animal[K];
};
```

### Indexed Access Types
Obtains the type of properties or elements of an array.

```typescript
 type Animal = {
  name: string;
  age: number;
};

type Name = Animal["name"]; // type Name = string;

type ArrayType = string[];
type ElementType = ArrayType[number]; // type ElementType = string;
```

### Conditional Types
Changes the type based on conditions.

```typescript
 type IsString<T> = T extends string ? "string" : "not string";

 type A = IsString<string>; // type A = "string";
```

### infer
A type operator used within conditional types to obtain a type variable.

```typescript
// Utility type to extract the return type of a function
 type MyReturnType<T> = T extends (...args: any[]) => infer R ? R : never;

// Example function
 function exampleFunction(): string {
   return "Hello, World!";
}

// Get the return type of the function
 type ExampleFunctionReturnType = MyReturnType<typeof exampleFunction>;

// ExampleFunctionReturnType will be of type string
 const exampleReturnValue: ExampleFunctionReturnType = "This is a string";

console.log(exampleReturnValue); // This is a string
```

### Union Distribution
Distributes union types and applies to each type.

```typescript
 type A = "a" | "b";

 type B = A extends "a" ? "c" : "d"; // type B = "c" | "d";
```

## Generics
Types that accept type arguments.

```typescript
// Generics
 function identity<T>(arg: T): T {
   return arg;
}

 const a = identity<string>("Hello");
 const b = identity<number>(1);

// Type arguments
 type Identity<T> = T;
 type A = Identity<string>; // type A = string;
```

# Thoughts
I have learned about JavaScript several times, but I feel like TypeScript makes JavaScript seem so much more difficult than I remembered...

# References
- [typescriptbook.jp - Survival TypeScript](https://typescriptbook.jp/)
  - Very well organized and easy to understand.