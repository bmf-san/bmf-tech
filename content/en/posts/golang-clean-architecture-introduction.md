---
title: Getting Started with Clean Architecture in Golang
slug: golang-clean-architecture-introduction
date: 2019-08-18T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Clean Architecture
  - Golang
  - DIP
description: An introduction to implementing Clean Architecture in Golang, including concepts, examples, and personal insights.
translation_key: golang-clean-architecture-introduction
---

# Overview
I tried implementing Clean Architecture in Golang and decided to organize my thoughts here.

The content largely follows the slides I created.

**There may be parts I haven't fully understood, along with my interpretations and thoughts, so some parts might not be entirely accurate.**

# Slides
Since I had the opportunity to give a lightning talk, here are the slides I used:

[Dive to clean architecture with golang](https://speakerdeck.com/bmf_san/dive-to-clean-architecture-with-golang)

# Source Code
Here is the source code:

[bmf-san/go-clean-architecture-web-application-boilerplate](https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate)

The implementation using the MVC pattern is also tagged and preserved:

[1.0.0 - MVC](https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate/releases/tag/1.0.0)

# Background
I have a CMS application called [github - bmf-san/Rubel](https://github.com/bmf-san/Rubel), which is no longer maintained but is used to run this blog.

To replace this application, I chose Go as the language and decided to adopt Clean Architecture to rethink the architecture.

The reason I considered adopting Clean Architecture is that I believe it is the optimal pattern for an application architecture that can be maintained for a long time without depending on libraries or other technologies.

Rubel uses frameworks like Laravel and React, but since it was implemented heavily relying on these frameworks, I felt it was a waste of time to keep up with the frequent updates of these relatively modern and fast-evolving frameworks.

Ideally, I wanted to focus on adding and improving CMS features. However, for an application I plan to maintain for a long time, spending time on non-essential development seemed unreasonable.

By minimizing dependencies on frameworks, libraries, and other technologies, and leveraging Go's standard library as much as possible, I thought I could create a highly maintainable application.

While I have a mindset that "scratch development is king," this application is not one that requires immediate adaptation to business requirements like service development. Since the development also includes a learning aspect, I think this approach is somewhat reasonable.

Although I feel I am taking a strategy close to optimal for personal development, I understand there will be aspects I won't see until the operational phase.

The replacement for Rubel, currently under development, is here:

[github - bmf-san/Gobel](https://github.com/bmf-san/Gobel)

# Table of Contents
- What is Clean Architecture?
- Implementing Clean Architecture (implementation details are not deeply covered)
- Reflections

# What is Clean Architecture?
## History of System Architectures
Before the idea of Clean Architecture emerged, several architectural ideas existed:

- Hexagonal Architecture (Ports and Adapters)
- Onion Architecture
- Screaming Architecture
- DCI
- BCE
- etc...

These ideas share the common goal of "separation of concerns," aiming for:

- Independence from frameworks
- Testability
- Independence from UI
- Independence from databases
- Independence from other technologies

In essence, they pursue decoupling from dependencies and enhancing testability.

## Clean Architecture
The famous diagram often associated with Clean Architecture originates from this source:

[cleancoder.com - The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

Here is an explanation of each layer:

### Entities
- Entities encapsulate the most important business rules.
  - e.g., Objects with methods or a set of data structures and functions.

### Use Cases
- Use cases contain specific business rules of the application.

### Interface Adapters
- Interface adapters transform data for entities and use cases.

### Frameworks and Drivers
- Frameworks and drivers consist of tools like frameworks and databases.

## Rules Between Layers
The constraints between the layers are as follows:

- There are four layers, but you can add or remove layers as needed.
- Inner layers do not know about outer layers.
    - → Dependencies should flow from outer to inner layers.

# Implementing Clean Architecture

## Directory Structure
The source code introduced earlier is repeated here:

[bmf-san/go-clean-architecture-web-application-boilerplate](https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate)

```
./app/
├── database
│   ├── migrations
│   │   └── schema.sql
│   └── seeds
│       └── faker.sql
├── domain
│   ├── post.go
│   └── user.go
├── go_clean_architecture_web_application_boilerplate
├── infrastructure
│   ├── env.go
│   ├── logger.go
│   ├── router.go
│   └── sqlhandler.go
├── interfaces
│   ├── post_controller.go
│   ├── post_repository.go
│   ├── sqlhandler.go
│   ├── user_controller.go
│   └── user_repository.go
├── log
│   ├── access.log
│   └── error.log
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
Before implementing Clean Architecture, you need to understand the Dependency Inversion Principle (DIP).

DIP is one of the SOLID principles and is a rule about module dependencies, stating that abstractions should not depend on details.

While I won't go into detail about this rule, in the context of Clean Architecture, this rule is followed by using interfaces to maintain dependency direction from outer to inner layers, adhering to the constraints between layers.

When implementing each layer's rules straightforwardly, situations may arise where dependencies point from inner to outer layers. In such cases, defining interfaces and depending on abstractions ensures the dependency direction is maintained. This is a key aspect of implementation.

## Accept Interfaces, Return Structs
Golang has a concept of "accept interfaces, return structs," which aligns well with implementing DIP.

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

This is a common implementation pattern in Golang. By depending on interfaces, you can write code that is resilient to changes and easier to test.

## DIP in Golang

Here is an example of DIP in Golang.

### Code Without DIP

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

### Code Considering DIP

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

	// sqlHandler struct implments SQLHandler interface.
	return &sqlHandler{}
}

// Execute is a function for executing sql.
// A sqlHanlder struct implments a SQLHandler interface by defining Execute().
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

By introducing an interface, the dependency relationship changes, effectively reversing the dependency direction.

Before:

```
SQLHandler
　　↑
FooRepository
```

After:

```
SQLHandler
   ↓   
SQLHandler Interface
   ↑
FooRepository
```

In the Clean Architecture example, the code in `infrastructure` and `interfaces` corresponds to this.
[bmf-san/go-clean-architecture-web-application-boilerplate](https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate)

## Code Reading
When tackling Clean Architecture, I found it easier to start with code reading or copying rather than jumping straight into implementation.

When reading code, I personally found it helpful to read from the outer layers inward:

main.go
  ↓
router.go (Infrastructure)
　↓
user_controller.go (Interfaces)
　↓
user_interactor.go (Use Cases)
　↓
user_repository.go (Use Cases)
　↓
user.go (Domain)

# Reflections
- Since I am relatively new to Golang, I had to repeatedly revisit language features like interfaces and structs.
- When faced with questions like "Where should I write this?", I felt it would be beneficial to have an architect lead and make decisions.
    - When adopting Clean Architecture, having someone in the team to act as an architect seems essential.
        - Though this might not be limited to Clean Architecture...
- I felt that Clean Architecture is more of a mindset than an implementation pattern, and I realized the need to study various architectural patterns more broadly.
- It seems to be more suited for monolithic applications.
    - For microservices, simpler and more disposable architectural patterns with lower learning costs might be preferred.
- A framework is just a tool, not a way of life.
    - A quote from the original text of "Clean Architecture: A Craftsman's Guide to Software Structure and Design."
    - What a great phrase.

# References
- [github - manuelkiessling/go-cleanarchitecture](https://github.com/manuelkiessling/go-cleanarchitecture)
- [github - rymccue/golang-standard-lib-rest-api](https://github.com/rymccue/golang-standard-lib-rest-api)
- [github - hirotakan/go-cleanarchitecture-sample](https://github.com/hirotakan/go-cleanarchitecture-sample)
- [Recruit Technologies - Go言語とDependency Injection](https://recruit-tech.co.jp/blog/2017/12/11/go_dependency_injection/)
- [Clean ArchitectureでAPI Serverを構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
- [github - ponzu-cms/ponzu](https://github.com/ponzu-cms/ponzu)
- [クリーンアーキテクチャの書籍を読んだのでAPIサーバを実装してみた](https://qiita.com/yoshinori_hisakawa/items/f934178d4bd476c8da32)
- [Go × Clean Architectureのサンプル実装](http://nakawatch.hatenablog.com/entry/2018/07/11/181453)
- [Uncle Bob – Payroll Case Study (A full implementation)](http://cleancodejava.com/uncle-bob-payroll-case-study-full-implementation/)
