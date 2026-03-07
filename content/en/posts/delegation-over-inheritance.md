---
title: About Delegation Over Inheritance
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

Object-oriented programming (OOP) is a way of thinking about "representing real-world objects in programs." However, directly applying real-world classifications and definitions into programs can lead to unexpected breakdowns.

In this article, we will explain three important concepts through the concrete example of "rectangle and square."

- **Duck Typing** - Determining type by behavior rather than name
- **Liskov Substitution Principle (LSP)** - Ensuring behavioral compatibility
- **Delegation Over Inheritance** - Achieving robust design

## Duck Typing - Type Determined by Behavior, Not Name

"Duck typing" is a way of thinking about types based on the following philosophical metaphor.

> "If it quacks like a duck and walks like a duck, then it is a duck."

This philosophy suggests that types are not determined by "type names" or "inheritance relationships," but rather by how the object behaves.

### Duck Typing in Go

The Go language is statically typed, yet it naturally realizes this philosophy. You don't need to explicitly write "implements"; as long as it has the necessary methods, it satisfies the interface.

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

`User` does not declaratively implement `Greeter`, but it can be treated as `Greeter` because it has the `Greet()` method. Thus, **abstraction by behavior** is the foundation of Go's interface design.

## Liskov Substitution Principle - Maintaining Behavioral Compatibility

One of the principles of object-oriented programming is the **Liskov Substitution Principle (LSP)**. This principle is defined as follows.

> Derived classes must be substitutable for their base classes without affecting the correctness of the program.

In other words, a subclass must behave the same way as its parent class. The key point here is not "structural matching" but rather "behavioral consistency."

### What is "Behavior"?

The behavior of an object refers to its **dynamic nature** in how it responds to external operations (method calls).

For example, if a type offers a contract that "width and height can be set independently," a subclass that breaks this contract, even if structurally similar, cannot be substituted.

Behavioral consistency is fundamental to the reliability of a program.

## Square Inheriting Rectangle - A Typical LSP Violation

A classic example of LSP violation is the case of "square (Square) inheriting rectangle (Rectangle)."

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

Now, let's implement square by inheriting this.

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

### Breakdown Example

At first glance, this seems correct, but it breaks down with the following code.

```php
<?php

$r = new Square();
$r->setWidth(5);
$r->setHeight(10);
echo $r->area(); // Expected: 50, Actual: 100
```

`Square` cannot be substituted as `Rectangle`. This is because it breaks the contract expected by the parent class, which is "**width and height can be changed independently**."

Thus, even if they are structurally similar, their behaviors do not match.

## Aristotelian Classification and OOP Misalignment

In Aristotelian classification, things are classified by common properties. For example, the classification "a square is a type of rectangle" seems natural.

However, this is a **structural classification (commonality in appearance or properties)**, which differs from the **behavioral classification (consistency in response to operations)** required in OOP.

Bringing real-world classification relationships directly into the program's inheritance structure risks violating the Liskov Substitution Principle.

## Delegation Over Inheritance - Composition over Inheritance

Inheritance may seem like a convenient means of reuse, but it strongly depends on the internal structure and behavior of the parent class, making it fragile to changes and prone to breaking the substitution principle.

To avoid this problem, the concept of **"delegation over inheritance"** is proposed. This design approach retains and utilizes the necessary functionality internally instead of inheriting.

### Example of Delegation

```go
package main

import "fmt"

type Logger struct{}

func (l Logger) Log(msg string) {
    fmt.Println(msg)
}

type Server struct {
    Logger // "has" a Logger (embedding)
}

func (s Server) Start() {
    s.Log("Starting server...")
}

func main() {
    server := Server{Logger: Logger{}}
    server.Start() // Output: Starting server...
}
```

In this design, `Server` uses `Logger` without inheriting it. This makes the dependencies explicit and increases maintainability.

## Solving with Delegation and Interfaces (in Go)

In Go, since inheritance does not exist, such problems can be naturally avoided. A square can "have" a rectangle to achieve equivalent functionality.

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

### Abstraction Using Interfaces

Alternatively, you can define a common interface for abstraction.

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

As long as both `Rectangle` and `Square` satisfy `Shape`, they can be treated as the same abstraction. This allows for ensuring reusability while preventing breakdowns due to inheritance.

## Conclusion

| Perspective | Content |
|------|------|
| **Duck Typing** | A way of determining type by behavior |
| **Behavior** | Consistent response of an object to external operations |
| **Liskov Substitution Principle** | Subclasses should be substitutable for their parent classes |
| **Delegation Over Inheritance** | Emphasizing reuse of behavior rather than structure |
| **Aristotelian Classification** | Bringing structural classification directly into programs leads to breakdowns |
| **Rectangle and Square Issue** | Structurally correct but behaviorally inconsistent, leading to LSP violation |

## Conclusion

The "is-a" relationship in object-oriented programming is established **only by behavioral consistency**, differing from philosophical and linguistic classifications.

Even if in the real world "a square is a type of rectangle," in programming, a "square that cannot behave like a rectangle" is not substitutable.

Therefore, in OOP design, it is essential to **base designs on "delegation over inheritance" and perform abstraction based on behavior rather than type names** to achieve robust design.

## References

- [Barbara Liskov, Data Abstraction and Hierarchy, 1987](https://www.cs.cmu.edu/~wing/publications/LiskovWing94.pdf)
- [Robert C. Martin, Design Principles and Design Patterns](https://fi.ort.edu.uy/innovaportal/file/2032/1/design_principles.pdf)
- [Effective Go - Embedding](https://go.dev/doc/effective_go#embedding)