---
title: "PHP7.4からPHP8.1までの新機能をキャッチアップ"
slug: "php8-1-new-features"
date: 2022-03-21
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "PHP"
draft: false
---

# 概要
PHPの知識が7.3から止まっているので8.1までの新機能を駆け足でキャッチアップする。

# PHP 7.3.x - PHP 7.4.x
## 新機能
### 型付きプロパティ
```php
<?php
	class Person {
		public int $age; // 指定の型だけ代入できるように強制される
		public string $name;
	}

?>
```

### アロー関数
```php
<?php
// これが
$a = 1
$func_7_3 = function($b) use ($a) {
	return $a + $b;
}
echo $func_7_3(1); // 2

// 7.4からはこのように書ける
$func_7_4 = fn($b) => $a + $b; // 暗黙的な値スコープを持つ
echo $func_7_4(1); // 2

?>
```

### Null合体代入演算子
```php
<?php
// これが
$name = isset($name) $name : getName();

// 7.4からはこのように書ける
$name ??= getName();
?>
```

### 配列内でのアンパック
```php
<?php
$values = ['a', 'b'];
$all = [...$values, 'c']; // a, b, c
?>
```

### FFI（Foreign Function Interface）
PHPから別言語を呼び出せる。
Goのcgoみたいなやつ。

### preload
opcacheにスクリプトを事前ロードする機能が追加された。

```ini
// php.ini
opcache.preload=preload.php
```

# PHP 7.4.x - PHP 8.0.x
## 新機能
### 名前付き引数
```php
<?php
function namedFunc($foo, $bar, $baz) {
  echo $foo . $bar . $baz;
}

// 引数の順番は関係なく、名前付きで引数をセットできる
namedFunc(baz: "baz", foo: 'foo', bar: "bar"); // bazfoobar
?>
```

### アトリビュート
アノテーションの新機能。

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

// Reflection APIでアトリビュートにアクセスできる
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
//   'name' =>
//   string(4) "John"
// }
// class Person#3 (1) {
//   public $name =>
//   string(4) "John"
// }
?>
```

### union型
引数、戻り地、プロパティに対してunion型を宣言できるようになった。

```php
<?php
function unionFunc(int|string $value): int|string
  return $value;
}

unionFunc(1);
unionFunc("Hello World");
?>
```

### match式
```php
<?php
function matchFunc($value) {
	// switch文とは異なり、厳密な比較（===）となる
	return match($value) {
		"one" => 1,
		"two" => 2,
		"three" => 3,
	};
};

echo matchFunc("one"); // 1
?>
```

### nullsafe演算子
null安全なコードが書きやすくなる。

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
## 新機能
### 文字列をキーとして持つ配列のアンパック
```php
<?php
	$a1 = ["a" => 1];
	$a2 = ["a" => 2];
	var_dump(["a"=>0, ...$a1, ...$a2]); // ["a" => 2] キーは後勝ちになるっぽい。
?>
```

### 列挙型
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
完全なスタックを持つ、停止可能な関数。コールスタックのどこからでも停止・再開が可能。
非同期処理が書ける。

### 交差型
unionが型のORなら、交差型は型のAND。

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

### never型
戻り値にのみ指定できる型。関数がexit()するか例外を投げるか終了しないかを示す。
voidとは違う。

# 所感
どんどん型を意識したコードが書きやすくなってきている印象。
下位互換性は流し読みしかしていないのでアプデ対応のときでも再度読み直したい。

# 参考
- [PHP 7.3.x から PHP 7.4.x への移行](https://www.php.net/manual/ja/migration74.php)
- [PHP 7.4.x から PHP 8.0.x への移行](https://www.php.net/manual/ja/migration80.php)
- [PHP 8.0.x から PHP 8.1.x への移行](https://www.php.net/manual/ja/migration81.php)
