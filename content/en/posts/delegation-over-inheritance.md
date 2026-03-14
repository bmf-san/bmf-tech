---
title: Delegation Over Inheritance
description: 'Learn duck typing and Liskov Substitution Principle. Build resilient Go applications using delegation patterns over inheritance.'
slug: delegation-over-inheritance
date: 2025-10-20T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Duck Typing
  - Liskov Substitution Principle
  - Delegation
  - Inheritance
  - Golang
translation_key: delegation-over-inheritance
---

## Introduction

Object-oriented programming (OOP) is a way of thinking about "representing real-world objects in programs." However, directly bringing real-world classifications and definitions into programs can lead to unexpected breakdowns.

In this article, we will explain three important concepts through the specific example of "Rectangle and Square."

- **Duck Typing** - Determining types by behavior, not by name
- **Liskov Substitution Principle (LSP)** - Ensuring behavioral compatibility
- **Delegation Over Inheritance** - Achieving robust design

## Duck Typing - Types Determined by Behavior

"Duck typing" is a type concept based on the following philosophical metaphor.

> "If it quacks like a duck and walks like a duck, then it is a duck."

This means determining types not by "type name" or "inheritance relationship," but by how the object behaves.

### Duck Typing in Go

The Go language naturally realizes this concept despite being statically typed. As long as an object has the necessary methods, it satisfies the interface without explicitly writing "implements."

```go
package main

import "fmt"

type Greeter interface {
    Greet() string
}

type User struct {
    Name string
}

func (u User) Greet() string {
    return "Hello, " + u.Name
}

func SayHello(g Greeter) {
    fmt.Println(g.Greet())
}

func main() {
    user := User{Name: "Alice"}
    SayHello(user) // Output: Hello, Alice
}
```

`User` does not explicitly implement `Greeter`, but since it has the `Greet()` method, it can be treated as a `Greeter`. This kind of **behavior-based abstraction** is fundamental to Go's interface design.

## Liskov Substitution Principle - Maintaining Behavioral Compatibility

One of the principles of object-oriented programming is the **Liskov Substitution Principle (LSP)**, defined as follows.

> Derived classes must be substitutable for their base classes without affecting the correctness of the program.

In other words, child classes must behave the same as their parent classes. The key here is "behavioral consistency," not "structural similarity."

### What is "Behavior"?

The behavior of an object is its **dynamic nature** of responding to external operations (method calls).

For example, if a type offers a contract that "width and height can be set independently," a child class that breaks this contract cannot be substituted, even if structurally similar.

Behavioral consistency is fundamental to the reliability of programs.

## Square Inherited from Rectangle - A Typical LSP Violation

A typical example of LSP violation is "Square inherited from Rectangle."

### Problematic Implementation

```php
<?php

class Rectangle {
    protected int $width;
    protected int $height;

    public function setWidth(int $w): void {
        $this->width = $w;
    }

    public function setHeight(int $h): void {
        $this->height = $h;
    }

    public function area(): int {
        return $this->width * $this->height;
    }
}
```

Implementing a square by inheriting this.

```php
<?php

class Square extends Rectangle {
    public function setWidth(int $w): void {
        $this->width = $w;
        $this->height = $w;  // Keep width and height the same
    }

    public function setHeight(int $h): void {
        $this->width = $h;   // Keep width and height the same
        $this->height = $h;
    }
}
```

### Breaking Example

It seems correct at first glance, but it breaks with the following code.

```php
<?php

$r = new Square();
$r->setWidth(5);
$r->setHeight(10);
echo $r->area(); // Expected: 50, Actual: 100
```

`Square` cannot be substituted for `Rectangle`. This is because it breaks the parent class's expectation of "**being able to change width and height independently**."

Thus, even if structurally similar, the behavior does not match.

## Aristotelian Classification and OOP Discrepancy

In Aristotelian classification, things are classified by common properties. For example, "a square is a type of rectangle" seems natural.

However, this is a **structural classification (commonality in appearance or properties)**, which differs from the **behavioral classification (consistency in response to operations)** required in OOP.

Bringing real-world classification relationships directly into program inheritance structures risks violating the Liskov Substitution Principle.

## Composition Over Inheritance

Inheritance may seem like a convenient means of reuse, but it strongly depends on the internal structure and behavior of the parent class, making it susceptible to changes and prone to violating substitution principles.

To avoid this problem, **"Composition over Inheritance"** is advocated. This is a design approach that involves holding and utilizing necessary functions internally rather than inheriting them.

### Example of Delegation

```go
package main

import "fmt"

type Logger struct{}

func (l Logger) Log(msg string) {
    fmt.Println(msg)
}

type Server struct {
    Logger // Embedding Logger
}

func (s Server) Start() {
    s.Log("Starting server...")
}

func main() {
    server := Server{Logger: Logger{}}
    server.Start() // Output: Starting server...
}
```

In this design, `Server` uses `Logger` without inheriting it. This makes dependencies explicit and enhances maintainability.

## Solving with Delegation and Interfaces (In Go)

In Go, since inheritance does not exist, such problems can be naturally avoided. A square can "hold" a rectangle to achieve equivalent functionality.

### Implementation Using Delegation

```go
package main

type Rectangle struct {
    Width, Height int
}

func (r Rectangle) Area() int {
    return r.Width * r.Height
}

type Square struct {
    rect Rectangle
}

func (s *Square) SetSize(n int) {
    s.rect.Width = n
    s.rect.Height = n
}

func (s Square) Area() int {
    return s.rect.Area()
}
```

### Abstraction with Interfaces

Alternatively, define a common interface for abstraction.

```go
package main

type Shape interface {
    Area() int
}

type Rectangle struct {
    Width, Height int
}

func (r Rectangle) Area() int {
    return r.Width * r.Height
}

type Square struct {
    Size int
}

func (s Square) Area() int {
    return s.Size * s.Size
}

func PrintArea(shape Shape) {
    println("Area:", shape.Area())
}

func main() {
    rect := Rectangle{Width: 5, Height: 10}
    square := Square{Size: 5}

    PrintArea(rect)   // Area: 50
    PrintArea(square) // Area: 25
}
```

As long as both `Rectangle` and `Square` satisfy `Shape`, they can be treated as the same abstraction. This ensures reusability while preventing breakdowns due to inheritance.

## Summary

| Aspect | Content |
|------|------|
| **Duck Typing** | The concept of determining types by behavior |
| **Behavior** | Consistent response of an object to external operations |
| **Liskov Substitution Principle** | Child classes should be substitutable for parent classes |
| **Composition Over Inheritance** | Emphasizing behavior reuse over structural reuse |
| **Aristotelian Classification** | Structural classification can lead to breakdowns when directly applied to programs |
| **Rectangle and Square Problem** | Structurally correct but behaviorally inconsistent, leading to LSP violation |

## Conclusion

The "is-a" relationship in object-oriented programming is established **only through behavioral consistency**, unlike philosophical or linguistic classifications.

In the real world, "a square is a type of rectangle," but in programs, "a square that cannot behave like a rectangle" is not substitutable.

Therefore, in OOP design, **"Composition over Inheritance" should be the foundation, and abstraction should be based on behavior, not type names**, leading to robust design.

## References

- [Barbara Liskov, Data Abstraction and Hierarchy, 1987](https://www.cs.cmu.edu/~wing/publications/LiskovWing94.pdf)
- [Robert C. Martin, Design Principles and Design Patterns](https://fi.ort.edu.uy/innovaportal/file/2032/1/design_principles.pdf)
- [Effective Go - Embedding](https://go.dev/doc/effective_go#embedding)
