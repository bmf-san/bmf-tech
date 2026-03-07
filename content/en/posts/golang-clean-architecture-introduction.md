---
title: Introduction to Clean Architecture with Golang
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

The content generally follows the slides.

**There may be parts that I do not fully understand, and I may have written my own interpretations and thoughts, so some parts may not be correct.**

# Slides
I had the opportunity to give a lightning talk, so I will attach the slides.

[Dive to clean architecture with golang](https://speakerdeck.com/bmf_san/dive-to-clean-architecture-with-golang)

# Source
Here is the source code.

[bmf-san/go-clean-architecture-web-application-boilerplate](https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate)

The implementation using the MVC pattern is also tagged and preserved.

[1.0.0 - MVC](https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate/releases/tag/1.0.0)

# Background
There is a CMS application called [github - bmf-san/Rubel](https://github.com/bmf-san/Rubel) that I am currently not maintaining, which operates this blog.

To replace this application, I chose Go as the language and decided to review the architecture, adopting Clean Architecture as the policy.

The reason for considering Clean Architecture is that I believe an architecture pattern that does not depend on libraries or other technologies is the optimal solution for applications that I can maintain for a long time as an individual.

Rubel uses frameworks like Laravel and React, but since it is implemented in a way that heavily relies on the frameworks, I felt that I was wasting time trying to keep up with the relatively modern and rapidly changing versions of those frameworks.

Originally, I should be focusing on adding and improving CMS features, but I couldn't rationalize spending time on non-essential parts of development for an application I want to operate for a long time in the future.

I thought that if I could minimize dependencies on frameworks, libraries, and other technologies while sufficiently utilizing Go's standard library, I could create a highly maintainable application.

While I have a strong desire to start from scratch, I recognize that this is not an application that requires immediate responses to business requirements like service development, and there are learning elements involved in the development purpose, so it seems somewhat reasonable.

I feel that I am taking an optimal strategy for individual development, but there are certainly aspects that will only become clear after entering the operational phase.

The ongoing development of the Rubel replacement can be found here.

[github - bmf-san/Gobel](https://github.com/bmf-san/Gobel)

# Table of Contents
- What is Clean Architecture?
- Implementation of Clean Architecture (not detailed on implementation methods)
- Thoughts

# What is Clean Architecture?
## History of System Architecture
Before the concept of Clean Architecture emerged, several architectural ideas existed in the past.

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

They pursue the elimination of dependencies on everything and testability.

## Clean Architecture
The well-known diagram when researching Clean Architecture refers to this original source.

[cleancoder.com - The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

I will explain each layer.

### Entities
- Entities encapsulate the most important business rules.
  - e.g., objects with methods or a series of data structures and functions.

### Use Cases
- Use cases contain specific business rules of the application.

### Interface Adapters
- Interface adapters are adapters that perform data transformation for entities and use cases.

### Frameworks and Drivers
- Frameworks and drivers consist of tools such as frameworks and databases.

## Rules Between Layers
Regarding the constraints between the layers mentioned above:

- There are four layers, but they are not limited to that. You can increase or decrease layers as needed.
- Inner layers do not know about outer layers.
    - → The direction of dependencies should flow from the outside to the inside.

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

The correspondence between layers and directories is as follows:

| Layer                | Directory      |
|----------------------|----------------|
| Frameworks & Drivers | infrastructure |
| Interface            | interfaces     |
| Usecases             | usecases       |
| Entities             | domain         |


## DIP
Before implementing Clean Architecture, it is necessary to understand the rule of DIP (Dependency Inversion Principle).

It is one of the SOLID principles and is a rule about the constraints between modules that states that abstractions should not depend on details.

I will omit the details of this rule, but in the context of Clean Architecture, this rule is maintained by using interfaces to keep the direction of dependencies flowing from the outside to the inside, while also adhering to the constraints between layers.

If you implement strictly according to the rules of each layer, it can lead to a situation where the direction of dependencies flows from the inside to the outside.

In such cases, defining interfaces and depending on abstractions is crucial to maintaining the direction of dependencies.

## Accept Interfaces, Return Structs
In Golang, there is a concept of "accepting interfaces and returning structs."

I think this is a concept that is compatible with the implementation of DIP.

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
// Also, this style of a function take on a role of constructor for struct.
func NewFooController(logger Logger) *FooController {
	return &FooController{
		Logger: logger,
	}
}
```

This is a basic implementation pattern often seen in Golang.

By depending on interfaces, you can write code that is resilient to changes and easy to test (or so it should be).

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

Code considering DIP.

```golang
package examples

// SQLHandler is an interface for handling sql.
type SQLHandler interface {
	Execute()
}

// sqlHandler is a struct which will be returned by function.
type sqlHandler struct{}

// NewSQLHandler is a function for an example of DIP.
// This function depend on abstruction(interface).
// This pattern is an idiom of constructor in golang.
// You can do DI(Dependency Injection) by using nested struct.
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

By inserting an interface, the dependency relationship changes, resulting in the direction of dependencies being inverted.

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
When actually tackling Clean Architecture, I think it is easier to understand the implementation by starting with code reading or copying rather than jumping straight into implementation.

When reading code, I found it easier to read from the outside in.

main.go
  ↓
router.go・・・Infrastructure
　↓
user_controller.go・・・Interfaces
　↓
user_interactor.go・・・Use Cases
　↓
user_repository.go・・・Use Cases
　↓
user.go・・・Domain

# Thoughts
- Since my experience with Golang is shallow, I had to relearn the language specifications such as interfaces and structs several times.
- I think it is better for an architect to lead and decide on the parts where there is confusion like "Where should this go?"
    - I felt that there should be at least one person in the team who plays the role of an architect when adopting Clean Architecture.
        - I think this is not limited to Clean Architecture...
- I believe that Clean Architecture is more of a way of thinking than an implementation pattern, so I felt the need to study a wide range of architectural patterns.
- I felt that there is a premise of fighting with monoliths.
    - If it were microservices, I think more easily discardable architectural patterns with lower learning costs would be preferred.
- A framework is just a tool, not a way of life.
    - A phrase from the original text of "Learning Software Structure and Design from Clean Architecture Experts."
    - What a great saying.

# References
- [github - manuelkiessling/go-cleanarchitecture](https://github.com/manuelkiessling/go-cleanarchitecture)
- [github - rymccue/golang-standard-lib-rest-api](https://github.com/rymccue/golang-standard-lib-rest-api)
- [github - hirotakan/go-cleanarchitecture-sample](https://github.com/hirotakan/go-cleanarchitecture-sample)
- [Recruit Technologies - Go Language and Dependency Injection](https://recruit-tech.co.jp/blog/2017/12/11/go_dependency_injection/)
- [Building an API Server with Clean Architecture](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
- [github - ponzu-cms/ponzu](https://github.com/ponzu-cms/ponzu)
- [Implemented an API Server after reading a book on Clean Architecture](https://qiita.com/yoshinori_hisakawa/items/f934178d4bd476c8da32)
- [Sample Implementation of Go × Clean Architecture](http://nakawatch.hatenablog.com/entry/2018/07/11/181453)
- [Uncle Bob – Payroll Case Study (A full implementation)](http://cleancodejava.com/uncle-bob-payroll-case-study-full-implementation/)