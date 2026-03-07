---
title: Summary of Go CodeReviewComments
slug: go-code-review-comments-summary
date: 2020-09-15T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Code Review
translation_key: go-code-review-comments-summary
---

# Overview
[github.com - CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments) summarizes the notes I want to keep from reading.

# Comment Sentences
- Comments should end with a period.
- [golang.org - commentary](https://golang.org/doc/effective_go.html#commentary)

# Copying
- Be careful of unexpected references when copying structs from another package.
- Use *T instead of T if the method is associated with a pointer value.

# Crypt Rand
- Do not use `math/rand` for key generation. Use `crypto/rand`'s `Reader`.
  - If you want to treat it as a string, encode it in hexadecimal or base64.

# Declaring Empty Slices
```go
// Length 0 slice
t := []string{}
```

Instead, use:

```go
// nil slice
var t []string
```

When encoding a JSON object, `nil` is converted to `null`, but `[]string{}` is converted to `[]`.

In interface design, it is better not to distinguish between the two, as it can lead to confusing mistakes.

# Don't Panic
- Avoid using `panic` for normal error handling; return multiple values including an `error` type instead.

# Goroutine Lifetimes
- Be clear about when a goroutine will finish when creating it.
- Goroutines can cause memory leaks due to blocking on channel sends and receives.
- The garbage collector will not stop a goroutine even if it cannot reach a blocked channel.

# Import Blank
- Using `import _ "pkg"` allows you to utilize side effects when importing a package.
- This method should only be used in the main package of the program or in tests.

# Import Dot
- Using . for import treats the imported package as part of the specified package, allowing you to avoid circular references.

```go
package foo_test

import (
    "bar/testutil" // also imported in foo
    . "foo" // makes foo_test appear as part of foo
)
```

# Named Result Parameters
- Use named return values when it is unclear what the return values mean.

# Naked Returns
- Same as Named Result Parameters.

# Receiver Type
Criteria for whether to use a pointer or value for a method receiver. If in doubt, use a pointer.

### Cases to Avoid Pointers
- Avoid pointers if the receiver is a `map`, `func`, or `channel`.
- Avoid pointers if the receiver is a `slice` and the method does not recreate the `slice`.
- If the receiver is small and is originally a value type (e.g., `time.Time`), or if it is a struct or array without fields that change or pointers, or types like `int` or `string`, it may be better for the receiver to be a value.
  - If a value is passed to a method, a copy is made in stack memory instead of allocating memory on the heap.

### Cases for Pointers
- If a method needs to change a value, the receiver should be a pointer.
- If the receiver is a `sync.Mutex` or has fields that require synchronization, the receiver should be a pointer.
- If the receiver is a large struct or array, it should be a pointer.
- Will the receiver's value be changed when functions are executed concurrently or methods are called?
  - Passing by value generates a copy of the receiver at method execution time. Therefore, changes to the receiver will not apply outside the method. If changes need to apply to the original receiver, the receiver should be a pointer.
- If the receiver may be a struct, array, slice, or other elements that could be changed, it is better to use a pointer for easier code reading.

# Useful Test Failures
Messages to convey when a test fails:
- What went wrong (≈ cause of error)
- What input was there (≈ test cases)
- What actual values were present (≈ actual)
- What values were expected (≈ expected)

# References
- [github.com - CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)