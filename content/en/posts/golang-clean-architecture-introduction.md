---
title: "Clean Architecture in Go: A Practical Implementation Guide"
description: 'Learn how to implement Clean Architecture in Go with practical code. Covers layer separation, dependency rules, directory structure, and real-world trade-offs.'
slug: golang-clean-architecture-introduction
date: 2019-08-18T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Clean Architecture
  - Golang
  - DIP
translation_key: golang-clean-architecture-introduction
---

# Overview
I attempted to implement Clean Architecture in Golang, so I will organize my thoughts here.

The content generally follows the material from the slides.

**There may be parts that I do not fully understand, and I am writing my interpretations and thoughts, so there may be inaccuracies.**

# Slides
I had the opportunity to give a lightning talk, so I will share the slides.

[Dive to clean architecture with golang](https://speakerdeck.com/bmf_san/dive-to-clean-architecture-with-golang)

# Source
Here is the source code.

[bmf-san/go-clean-architecture-web-application-boilerplate](https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate)

The implementation using the MVC pattern is also tagged and preserved.

[1.0.0 - MVC](https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate/releases/tag/1.0.0)

# Background
There is a CMS application called [github - bmf-san/Rubel](https://github.com/bmf-san/Rubel) that I am currently not maintaining but was used to run this blog.

To replace this application, I chose Go as the language and decided to adopt Clean Architecture to review the architecture.

The reason I considered adopting Clean Architecture is that I believe an architecture pattern that does not depend on libraries or other technologies is the optimal solution for an application that I want to maintain for a long time.

Rubel uses frameworks like Laravel and React, but the implementation heavily relies on these frameworks, making it difficult to keep up with the rapid version updates of these relatively modern and changing frameworks.

Ideally, I would want to focus on adding and improving CMS functionalities, but I found it unreasonable to spend time on non-essential development for an application I want to operate for a long time.

I thought that if I could minimize dependencies on frameworks, libraries, and other technologies while sufficiently utilizing Go's standard library, I could create a highly maintainable application.

While I have a strong desire to start from scratch, I recognize that this is not an application that requires immediate responses to business requirements like service development, and there is a learning element in the development purpose, so it seems somewhat reasonable.

I feel that I am taking an optimal strategy for personal development, but there will likely be aspects that are not visible until I enter the operational phase.

Here is the ongoing development of the Rubel replacement.

[github - bmf-san/Gobel](https://github.com/bmf-san/Gobel)

# Table of Contents
- What is Clean Architecture?
- Implementation of Clean Architecture (implementation methods are not detailed)
- Thoughts

# What is Clean Architecture?
## History of System Architecture
Before the concept of Clean Architecture emerged, there were several architectural ideas.

- Hexagonal Architecture (Ports and Adapters)
- Onion Architecture
- Screaming Architecture
- DCI
- BCE
- etc...

These ideas share the common goal of "separation of concerns," which includes:

- Independence from frameworks
- Testability
- Independence from UI
- Independence from databases
- Independence from other technologies

They pursue the elimination of dependencies and testability concerning all these aspects.

## Clean Architecture
The well-known diagram when researching Clean Architecture refers to this original source.

[cleancoder.com - The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

Let’s explain each layer.

### Entities
- Entities encapsulate the most important business rules.
  - e.g., objects with defined methods or a set of data structures and functions.

### Use Cases
- Use cases contain specific business rules of the application.

### Interface Adapters
- Interface adapters are adapters that perform data transformation for entities and use cases.

### Frameworks and Drivers
- Frameworks and drivers consist of tools like frameworks and databases.

## Rules Between Layers
Regarding the constraints between the layers mentioned above:

- There are four layers, but it is not limited to that. You can increase or decrease layers as needed.
- Inner layers do not know about outer layers.
    - → The direction of dependencies should flow from the outer layers to the inner layers.

# Implementation of Clean Architecture

## Directory Structure
This is the same as the source introduced at the beginning, but I will reiterate it.

[bmf-san/go-clean-architecture-web-application-boilerplate](https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate)

```
./app/
├── database
│   ├── migrations
│   │   └── schema.sql
│   └── seeds
│       └── faker.sql
├── domain
│   ├── post.go
│   └── user.go
├── go_clean_architecture_web_application_boilerplate
├── infrastructure
│   ├── env.go
│   ├── logger.go
│   ├── router.go
│   └── sqlhandler.go
├── interfaces
│   ├── post_controller.go
│   ├── post_repository.go
│   ├── sqlhandler.go
│   ├── user_controller.go
│   └── user_repository.go
├── log
│   ├── access.log
│   └── error.log
├── main.go
└── usecases
    ├── logger.go
    ├── post_interactor.go
    ├── post_repository.go
    ├── user_interactor.go
    └── user_repository.go

8 directories, 22 files
```

The correspondence between layers and directories is as follows.

| Layer                | Directory      |
|----------------------|----------------|
| Frameworks & Drivers | infrastructure |
| Interface            | interfaces     |
| Usecases             | usecases       |
| Entities             | domain         |

## DIP
Before implementing Clean Architecture, it is necessary to understand the rule of DIP (Dependency Inversion Principle).

It is one of the SOLID principles, which is a rule about constraints between modules that states that abstractions should not depend on details.

I will skip the details of this rule, but in the context of Clean Architecture, this rule is upheld by utilizing interfaces to maintain the direction of dependencies from the outer layers to the inner layers while adhering to the constraints between layers.

If you strictly follow the rules of each layer during implementation, you may encounter situations where the direction of dependencies flows from the inner layers to the outer layers.

In such cases, defining interfaces and depending on abstractions helps maintain the direction of dependencies, which is a crucial part of the implementation.

## Accept interfaces, return structs
In Golang, there is a concept of "accept interfaces, return structs."

I think this is a compatible idea for implementing DIP.

```golang
package examples

// Logger is an interface which will be used for an argument of a function.
type Logger interface {
	Printf(string, ...interface{})
}

// FooController is a struct which will be returned by function.
type FooController struct {
	Logger Logger
}

// NewFooController is a function for an example, "Accept interfaces, return structs".
// Also, this style of a function takes on a role of constructor for struct.
func NewFooController(logger Logger) *FooController {
	return &FooController{
		Logger: logger,
	}
}
```

This is a basic implementation pattern commonly seen in Golang.

By depending on interfaces, you can write code that is resilient to changes and easy to test.

## DIP in Golang

An example of DIP in Golang.

Code that does not follow DIP.

```golang
package examples

// sqlHandler is a struct for handling sql.
type sqlHandler struct{}

// Execute is a function for executing sql.
func (sqlHandler *sqlHandler) Execute() {
	// do something...
}

// FooRepository is a struct depending on details.
type FooRepository struct {
	sqlHandler sqlHandler
}

// Find is a method depending on details.
func (ur *FooRepository) Find() {
	// do something
	ur.sqlHandler.Execute()
}
```

Code that considers DIP.

```golang
package examples

// SQLHandler is an interface for handling sql.
type SQLHandler interface {
	Execute()
}

// sqlHandler is a struct which will be returned by function.
type sqlHandler struct{}

// NewSQLHandler is a function for an example of DIP.
// This function depends on abstraction (interface).
// This pattern is an idiom of constructor in golang.
// You can do DI (Dependency Injection) by using nested struct.
func NewSQLHandler() SQLHandler {
	// do something ...

	// sqlHandler struct implements SQLHandler interface.
	return &sqlHandler{}
}

// Execute is a function for executing sql.
// A sqlHandler struct implements a SQLHandler interface by defining Execute().
func (s *sqlHandler) Execute() {
	// do something...
}

// FooRepository is a struct depending on an interface.
type FooRepository struct {
	SQLHandler SQLHandler
}

// Find is a method of FooRepository depending on an interface.
func (ur *FooRepository) Find() {
	// do something
	ur.SQLHandler.Execute()
}
```

By inserting an interface, the dependency relationship changes, resulting in the reversal of the direction of dependencies.

Before

```
SQLHandler
　　↑
FooRepository
```

After

```
SQLHandler
   ↓   
SQLHandler Interface
   ↑
FooRepository
``` 

In the practical example of Clean Architecture, the code in infrastructure and interfaces corresponds to this.
[bmf-san/go-clean-architecture-web-application-boilerplate](https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate)

## Code Reading
When actually tackling Clean Architecture, I think it is easier to understand the implementation by starting with code reading or transcription rather than jumping straight into implementation.

When reading code, I found it helpful to read from the outer layers to the inner layers.

main.go
  ↓
router.go...Infrastructure
　↓
user_controller.go...Interfaces
　↓
user_interactor.go...Use Cases
　↓
user_repository.go...Use Cases
　↓
user.go...Domain

# Thoughts
- Since my experience with Golang is limited, I had to relearn the language specifications regarding interfaces and structs several times.
- I think it is better for an architect to lead and decide on the ambiguous parts like "Where should this be written?"
    - I felt that there should be at least one person in the team who plays the role of architect when adopting Clean Architecture.
        - I think this is not limited to Clean Architecture...
- I felt that Clean Architecture is more of a way of thinking than an implementation pattern, so I need to study a wide range of architectural patterns.
- I felt that there is an assumption of fighting with a monolith.
    - If it were microservices, I think there would be a preference for architecture patterns that are easier to learn and discard.
- A framework is just a tool, not a way of life.
    - A phrase from the original text of "Clean Architecture: A Craftsman's Guide to Software Structure and Design."
    - Such a good saying.

# References
- [github - manuelkiessling/go-cleanarchitecture](https://github.com/manuelkiessling/go-cleanarchitecture)
- [github - rymccue/golang-standard-lib-rest-api](https://github.com/rymccue/golang-standard-lib-rest-api)
- [github - hirotakan/go-cleanarchitecture-sample](https://github.com/hirotakan/go-cleanarchitecture-sample)
- [Recruit Technologies - Go Language and Dependency Injection](https://recruit-tech.co.jp/blog/2017/12/11/go_dependency_injection/)
- [Building an API Server with Clean Architecture](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
- [github - ponzu-cms/ponzu](https://github.com/ponzu-cms/ponzu)
- [I Read a Book on Clean Architecture, So I Tried Implementing an API Server](https://qiita.com/yoshinori_hisakawa/items/f934178d4bd476c8da32)
- [Sample Implementation of Go × Clean Architecture](http://nakawatch.hatenablog.com/entry/2018/07/11/181453)
- [Uncle Bob – Payroll Case Study (A full implementation)](http://web.archive.org/web/20240106141740/http://cleancodejava.com/uncle-bob-payroll-case-study-full-implementation/)
