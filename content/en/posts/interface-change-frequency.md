---
title: Why Interfaces Change Less Frequently Than Implementations
description: A step-by-step guide on Why Interfaces Change Less Frequently Than Implementations, with practical examples and configuration tips.
slug: interface-change-frequency
date: 2025-10-18T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Interface
translation_key: interface-change-frequency
---



Reading [Clean Code Cookbook: Recipes for Improving Code Design and Quality](https://amzn.to/47uvc3g), I was intrigued by the claim that interfaces change less frequently than implementations, so I decided to articulate it.

## Interfaces as "Contracts" and "Abstractions"

Interfaces represent a "contract" that states:

> "This feature can be used in this way."

In contrast, implementations represent:

> "The specific method of how it operates."

Both have different roles and varying degrees of resistance to change.

| Layer                | Role      | Susceptibility to Change |
| -------------------- | --------- | ------------------------ |
| Abstraction (interface) | Purpose/Promise | More stable              |
| Concrete (implementation) | Method/Means   | More changeable          |

## Contracts Are Shared Externally, So They Can't Be Changed Recklessly

Changing an interface breaks **all calling code** that uses it.

Example:

```go
type UserRepository interface {
    Find(id int) (*User, error)
}
```

Changing it to:

```go
type UserRepository interface {
    FindByID(ctx context.Context, id int) (*User, error)
}
```

Requires all calling locations to be updated.

```go
// Before change
user, err := repo.Find(123)

// After change
ctx := context.Background()
user, err := repo.FindByID(ctx, 123)
```

Interface changes have a wide impact (i.e., they are fragile), so they are approached cautiously.

As a result, they are **designed to change rarely**.

## Implementations Are Behind the Scenes, So They Can Be Changed Freely

Implementations are not called directly from outside.

Internal logic, caching methods, algorithms, and storage can be changed **without affecting users**.

```go
type userRepository struct {
    db    *sql.DB
    cache map[int]*User // Added cache
    mu    sync.RWMutex  // For concurrency safety
}

func (r *userRepository) Find(id int) (*User, error) {
    // The interface remains unchanged, but internal implementation can be freely modified

    // Version 1: Direct DB search
    // return r.findFromDB(id)

    // Version 2: Search with cache
    r.mu.RLock()
    user, exists := r.cache[id]
    r.mu.RUnlock()

    if exists {
        return user, nil
    }

    user, err := r.findFromDB(id)
    if err == nil {
        r.mu.Lock()
        r.cache[id] = user
        r.mu.Unlock()
    }
    return user, err
}

func (r *userRepository) findFromDB(id int) (*User, error) {
    // Database access logic
    // Changing from PostgreSQL to MySQL does not affect the outside
    var user User
    err := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)
    return &user, err
}
```

Implementations are **targets for internal improvement, optimization, and refactoring**.

Thus, they are "**layers that can change frequently without breaking**."

## The Higher the Level of Abstraction, the More Resistant to Change

Abstraction expresses "requirements (what to do)."

Implementation expresses "means (how to do it)."

**Means change, but purposes are less likely to change.**

```go
// Abstract layer (stable)
type NotificationService interface {
    Send(message string, recipient string) error
}

// Implementation layer (prone to change)
type emailNotifier struct{}
func (e *emailNotifier) Send(message, recipient string) error {
    // SMTP → SendGrid → AWS SES, etc., implementations change
}

type slackNotifier struct{}
func (s *slackNotifier) Send(message, recipient string) error {
    // Slack API implementation
}

type smsNotifier struct{}
func (s *smsNotifier) Send(message, recipient string) error {
    // Twilio → AWS SNS, etc., implementations change
}
```

- "Want to send notifications" (abstraction) is less likely to change
- "Send via Email/Slack/SMS" (implementation) changes frequently

Therefore, **interfaces as abstractions are more stable**.

## Relation to Go's Design Philosophy

In Go, it is customary to **keep interfaces small and define them on the consumer side**.

This means:

> Usage (contract) is stable,
> but implementation (internal logic) can be freely changed.

```go
// Small interface (stable)
// The standard library's io.Reader has only one method
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Various implementations (prone to change)
type fileReader struct { /* File reading */ }
type networkReader struct { /* Network reading */ }
type compressedReader struct { /* Compressed file reading */ }
```

Since Go lacks explicit `implements` declarations (structural subtyping), implementers can implement multiple interfaces without awareness, and consumers can define only the contracts they need.

This characteristic naturally leads to designing with **dependency direction from "stable → unstable"**.

## Summary

| Aspect   | Interface        | Implementation       |
| -------- | ---------------- | -------------------- |
| Role     | Function promise (contract) | Actual operation (means) |
| Usage    | Exposed externally | Internal only         |
| Change Impact | Large (fragile) | Small (self-contained) |
| Result   | Hard to change (stable) | Easy to change (frequent) |
| Essence  | "Purpose" doesn't change | "Method" changes        |

### Conclusion

Interfaces are "contracts with users," and once a contract is established, it cannot be easily changed.

On the other hand, implementations can be freely changed as long as they adhere to the contract.

Therefore, **interfaces change less frequently**.