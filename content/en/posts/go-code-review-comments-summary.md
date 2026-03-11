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
I have summarized the points I want to note after reading [github.com - CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments).

# Comment Sentences
- End comments with a period.
- [golang.org - commentary](https://golang.org/doc/effective_go.html#commentary)

# Copying
- Be cautious of unexpected references when copying structures from other packages.
- If a method is associated with a pointer value, use *T instead of T.

# Crypt Rand
- Do not use `math/rand` for key generation. Use `crypto/rand`'s `Reader`.
  - If you want to treat it as a string, encode it in hexadecimal or base64.

# Declaring Empty Slices
```go
// Slice of length 0
t := []string{}
```
Instead, use:

```go
// Nil slice
var t []string
```
When encoding a JSON object, `nil` is converted to `null`, while `[]string{}` is converted to `[]`.

In interface design, it's better not to distinguish between the two, as it may lead to confusing errors.

# Don't Panic
- Avoid using `panic` for regular error handling; instead, return multiple values including an `error` type.

# Goroutine Lifetimes
- Clearly define when a goroutine will end when creating it.
- Goroutines can cause memory leaks if blocked by channel send/receive.
- The garbage collector will not stop a goroutine even if it cannot reach a blocked channel.

# Import Blank
- Using `import _ "pkg"` allows you to utilize the side effects of importing a package.
- This method should only be used in the main package of a program or in tests.

# Import Dot
- Importing with a dot treats the imported package as part of the specified package, helping to avoid circular references.

```go
package foo_test

import (
    "bar/testutil" // Also imported in foo
    . "foo" // Makes foo_test appear as part of foo
)
```

# Named Result Parameters
- Use named result parameters when the meaning of the return values is unclear.

# Naked Returns
- Same as Named Result Parameters.

# Receiver Type
Guidelines for choosing whether to make a method's receiver a pointer or a value.
If unsure, use a pointer.

### Cases to Avoid Pointers
- Avoid pointers if the receiver is a `map`, `func`, or `channel`.
- If the receiver is a `slice` and the method does not recreate the `slice`, avoid pointers.
- If the receiver is small, inherently a value type (e.g., `time.Time`), or a structure or array without fields or pointers to modify, or a type like `int` or `string`, it may be better as a value.
  - If the value is passed to a method, it will be copied to stack memory instead of allocating memory in the heap.

### Cases to Use Pointers
- If a method needs to modify the value, the receiver should be a pointer.
- If the receiver is a `sync.Mutex` or a structure with fields that require synchronization, use a pointer.
- If the receiver is a large structure or array, use a pointer.
- If the function is executed concurrently or the method modifies the receiver's value when called:
  - Value passing generates a copy of the receiver when the method is executed. Therefore, changes to the receiver are not applied outside the method. If changes need to be applied to the original receiver, use a pointer.
- If the receiver might be a structure, array, slice, or other element that may be modified, using a pointer makes the code easier to read.

# Useful Test Failures
Messages to convey when a test fails.
- What went wrong (≒cause of error)
- What input was given (≒test cases)
- What the actual values were (≒actual)
- What values were expected (≒expected)

# References
- [github.com - CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)
