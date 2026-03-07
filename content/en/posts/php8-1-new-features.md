---
title: Catching Up on New Features from PHP 7.4 to PHP 8.1
slug: php8-1-new-features
date: 2022-03-21T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - PHP
translation_key: php8-1-new-features
---

# Overview
I will quickly catch up on the new features from PHP 7.4 to 8.1 since my knowledge has been stuck at 7.3.

# PHP 7.3.x - PHP 7.4.x
## New Features
### Typed Properties
```php
<?php
	class Person {
		public int $age; // Enforced to only allow specified type
		public string $name;
	}

?>
```

### Arrow Functions
```php
<?php
// This is
$a = 1
$func_7_3 = function($b) use ($a) {
	return $a + $b;
}
echo $func_7_3(1); // 2

// From 7.4, you can write like this
$func_7_4 = fn($b) => $a + $b; // Implicit value scope

echo $func_7_4(1); // 2

?>
```

### Null Coalescing Assignment Operator
```php
<?php
// This is
$name = isset($name) ? $name : getName();

// From 7.4, you can write like this
$name ??= getName();
?>
```

### Array Unpacking
```php
<?php
$values = ['a', 'b'];
$all = [...$values, 'c']; // a, b, c
?>
```

### FFI (Foreign Function Interface)
You can call other languages from PHP. Similar to Go's cgo.

### Preload
A feature has been added to preload scripts into opcache.

```ini
// php.ini
opcache.preload=preload.php
```

# PHP 7.4.x - PHP 8.0.x
## New Features
### Named Arguments
```php
<?php
function namedFunc($foo, $bar, $baz) {
	echo $foo . $bar . $baz;
}

// You can set arguments by name regardless of their order
namedFunc(baz: "baz", foo: 'foo', bar: "bar"); // bazfoobar
?>
```

### Attributes
A new feature for annotations.

```php
<?php
#[Attribute]
class Person
{
	public $name;
	
	public function __construct($name)
	{
		$this->name = $name;
	}
}

#[Person(name: "John")]
class Man
{
}

// You can access attributes via the Reflection API
function output($reflection) {
	foreach ($reflection->getAttributes() as $attribute) {
		var_dump($attribute->getName());
		var_dump($attribute->getArguments());
		var_dump($attribute->newInstance());
	}
}

output(new ReflectionClass(Man::class));
// string(6) "Person"
// array(1) {
// 	'name' =>
// 	string(4) "John"
// }
// class Person#3 (1) {
// 	public $name =>
// 	string(4) "John"
// }
?>
```

### Union Types
You can now declare union types for arguments, return types, and properties.

```php
<?php
function unionFunc(int|string $value): int|string
	return $value;

unionFunc(1);
unionFunc("Hello World");
?>
```

### Match Expression
```php
<?php
function matchFunc($value) {
	// Unlike switch statements, it uses strict comparison (===)
	return match($value) {
		"one" => 1,
		"two" => 2,
		"three" => 3,
	};
};

echo matchFunc("one"); // 1
?>
```

### Nullsafe Operator
Makes it easier to write null-safe code.

```php
<?php
class User
{
	public $name;
	
	public function __construct()
	{
		$this->name = $name;
	}

	public function getName()
	{
		return $this->name;
	}
}

class Account
{
	public User|null $user = null;
}

$account = new Account();
// $account->user->getName(); // PHP Fatal error:  Uncaught Error: Call to a member function getName() on null
$account->user?->getName();
?>
```

# PHP 8.0.x - PHP 8.1.x
## New Features
### Array Unpacking with String Keys
```php
<?php
	$a1 = ["a" => 1];
	$a2 = ["a" => 2];
	var_dump(["a"=>0, ...$a1, ...$a2]); // ["a" => 2] The last key wins.
?>
```

### Enumerations
```php
<?php
enum ColorCode: string
{
	case BLUE = "#0000ff";
	case YELLOW = "#ffff00";
	case RED = "#ff0000";
}

function enumFunc(ColorCode $color) {
	echo $color->name; 
	echo $color->value; 
}

enumFunc(ColorCode::BLUE); // BLUE blue
// enumFunc("green"); // PHP Fatal error:  Uncaught TypeError: enumFunc(): Argument #1 ($color) must be of type ColorCode, string given
?>
```

### Fibers
Functions that can be paused and resumed from anywhere in the call stack. You can write asynchronous code.

### Intersection Types
If union types are OR types, intersection types are AND types.

```php
<?php
function check(Foo&Bar $intersection)
{
	return;
}

class Foo{}
class Bar{}
class Baz{}


$foo = new Foo();
$foo->check(new Foo());
$foo->check(new Bar());
// $foo->check(new Baz()); // PHP Fatal error:  Uncaught Error: Call to undefined method Foo::check() 
?>
```

### Never Type
A type that can only be specified for return values. Indicates whether a function exits or throws an exception or does not terminate. Different from void.

# Thoughts
It feels like writing code with a focus on types is becoming easier and easier. I have only skimmed the backward compatibility, so I would like to read it again when updating.

# References
- [Migration from PHP 7.3.x to PHP 7.4.x](https://www.php.net/manual/ja/migration74.php)
- [Migration from PHP 7.4.x to PHP 8.0.x](https://www.php.net/manual/ja/migration80.php)
- [Migration from PHP 8.0.x to PHP 8.1.x](https://www.php.net/manual/ja/migration81.php)