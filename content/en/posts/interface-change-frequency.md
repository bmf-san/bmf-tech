---
title: Why Interfaces Change Less Frequently Than Implementations
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

[Clean Code Cookbook - A Collection of Recipes to Improve Code Design and Quality](https://amzn.to/47uvc3g) caught my attention with its claim that interfaces change less frequently than their implementations, so I decided to articulate my thoughts on this.

## Interfaces as "Contracts" and "Abstractions"

An interface is,

> A "contract" that represents "how this functionality can be used."

On the other hand, an implementation is,

> A "concrete means" of "how to actually make it work."

The two have different roles and varying degrees of susceptibility to change.

| Layer                | Role         | Susceptibility to Change |
| -------------------- | ------------ | ------------------------ |
| Abstraction (interface) | Purpose/Promise | More Stable              |
| Concrete (implementation) | Method/Means   | More Changeable          |

## Contracts Cannot Be Changed Arbitrarily Because They Are Shared Externally

Changing an interface will break **all the calling code that uses it**.

Example:

```go
type UserRepository interface {
    Find(id int) (*User, error)
}
```

If we change it to:

```go
type UserRepository interface {
    FindByID(ctx context.Context, id int) (*User, error)
}
```

All the places that call it will need to be modified.

```go
// Before modification
user, err := repo.Find(123)

// After modification
ctx := context.Background()
user, err := repo.FindByID(ctx, 123)
```

Changes to interfaces have a wide impact (i.e., they are fragile), so we must be cautious.

As a result, we design them to **rarely change**.

## Implementations Are Internal and Can Be Changed Freely

Implementations are not called directly from the outside.

Internal logic, caching methods, algorithms, and storage can be changed **as long as they do not affect the users**.

```go
type userRepository struct {
    db    *sql.DB
    cache map[int]*User // Added caching
    mu    sync.RWMutex  // For concurrency safety
}

func (r *userRepository) Find(id int) (*User, error) {
    // The interface remains unchanged, but the internal implementation can be freely modified

    // Version 1: Direct DB search
    // return r.findFromDB(id)

    // Version 2: Caching search
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
    // Changing from PostgreSQL to MySQL does not affect the external interface
    var user User
    err := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)
    return &user, err
}
```

Implementations are subject to **internal improvements, optimizations, and refactoring**.

In other words, they are a **layer that can change frequently without breaking**.

## Higher Levels of Abstraction Are More Resilient to Change

Abstraction expresses "requirements (what to do)".

Implementation expresses "means (how to do it)".

**Means can change, but purposes are less likely to change.**

```go
// Abstraction layer (stable)
type NotificationService interface {
    Send(message string, recipient string) error
}

// Implementation layer (changeable)
type emailNotifier struct{}
func (e *emailNotifier) Send(message, recipient string) error {
    // Implementation can change from SMTP to SendGrid to AWS SES
}

type slackNotifier struct{}
func (s *slackNotifier) Send(message, recipient string) error {
    // Slack API implementation
}

type smsNotifier struct{}
func (s *smsNotifier) Send(message, recipient string) error {
    // Implementation can change from Twilio to AWS SNS
}
```

- "I want to send a notification" (abstraction) is less likely to change.
- "Send via Email/Slack/SMS" (implementation) changes frequently.

Thus, **interfaces as abstractions are more stable**.

## Relation to Go's Design Philosophy

In Go, it is customary to **keep interfaces small and define them on the consumer side**.

This means,

> The way they are used (i.e., the contract) is stable, but
> the implementation (i.e., internal logic) can change freely.

```go
// Small interface (stable)
// The standard library's io.Reader has only one method
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Various implementations (changeable)
type fileReader struct { /* File reading */ }
type networkReader struct { /* Network reading */ }
type compressedReader struct { /* Compressed file reading */ }
```

Since Go does not have explicit `implements` declarations (structural subtyping), the implementation side can implement multiple interfaces without being aware of them, allowing the consumer to define only the necessary contracts.

This characteristic naturally leads to designing the dependency direction as **"stable → unstable"**.

## Conclusion

| Perspective | Interface         | Implementation       |
| ----------- | ------------------ | -------------------- |
| Role        | Function promise (contract) | Actual behavior (means) |
| Scope       | Exposed externally  | Internal only        |
| Impact of Change | Large (fragile)     | Small (self-contained) |
| Result      | Hard to change (i.e., stable) | Easy to change (i.e., frequent) |
| Essence     | "Purpose" does not change | "Method" changes    |

### Conclusion

Interfaces are a "contract with users," and once a contract is established, it cannot be easily changed.

On the other hand, implementations can be freely changed as long as they adhere to the contract.

Therefore, **interfaces change less frequently**.