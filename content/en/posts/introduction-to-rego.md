---
title: Introduction to Rego
slug: introduction-to-rego
date: 2025-07-31T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Open Policy Agent
  - Rego
translation_key: introduction-to-rego
---

## Table of Contents

- [Chapter 1: What is Rego](#chapter-1-what-is-rego)
- [Chapter 2: Basic Syntax and Data Structures](#chapter-2-basic-syntax-and-data-structures)
- [Chapter 3: Types of Rules and How to Write Them](#chapter-3-types-of-rules-and-how-to-write-them)
- [Chapter 4: Control Structures and Operators](#chapter-4-control-structures-and-operators)
- [Chapter 5: Comprehensions and Data Manipulation](#chapter-5-comprehensions-and-data-manipulation)
- [Chapter 6: Built-in Functions](#chapter-6-built-in-functions)
- [Chapter 7: Testing and Debugging](#chapter-7-testing-and-debugging)

## Chapter 1: What is Rego

### 1.1 Philosophy and Features of Rego

Rego is a **declarative** policy language. It describes "what should be satisfied" rather than "how to process it."

#### Comparison of Imperative vs Declarative

| Imperative (Go, JavaScript, etc.) | Declarative (Rego) |
| --------------------------------- | ------------------ |
| Loop with for and condition check | `some i; input.groups[i] == "admin"` |
| Branching with if/else            | Condition description with multiple rules |
| Procedural execution flow         | Determination of condition satisfaction |

#### Separation of Input and Data

Rego handles two types of data:

| Type    | Description                       | Example                |
| ------- | --------------------------------- | ---------------------- |
| `input` | Dynamic data provided at runtime  | Request content        |
| `data`  | Static, shared data for reference | User information, role definitions |

This separation makes policies generic and reusable.

### 1.2 Rule-based Evaluation Model

The basic unit in Rego is a **rule**:

- Multiple rules are evaluated, and if conditions are met, it returns `true`
- If there is no explicit `default`, unmatched cases result in `undefined` (not evaluated)
- Decision description: "Is it allowed?" "Which fields are visible?"

## Chapter 2: Basic Syntax and Data Structures

### 2.1 Basic Syntax Elements

#### Package Statement
```rego
package authz  # Define namespace
```

#### Import Statement
```rego
import data.roles      # Reference data
import input.user as u # Use alias
```

#### Comments
```rego
# Single line comment
allow {  # Inline comment
    input.user.role == "admin"
}
```

### 2.2 Data Types

| Data Type  | Example                        | Description       |
| ---------- | ------------------------------ | ----------------- |
| String     | `"admin"`, `"user"`           | Double quotes     |
| Number     | `42`, `3.14`                  | Integer, float    |
| Boolean    | `true`, `false`               | Boolean values    |
| Array      | `[1, 2, 3]`                   | Ordered           |
| Object     | `{"name": "Alice", "age": 30}` | Key-value pairs  |
| Set        | `{"admin", "user"}`           | No duplicates     |

### 2.3 input and data

#### input - Dynamic Data
**[Try on Playground](https://play.openpolicyagent.org/)**

```rego
package input_example

import rego.v1

# Request information, etc.
allow if {
    input.method == "GET"
    input.user.role == "admin"
}

# Sample execution for testing
# Input: {"method": "GET", "user": {"role": "admin"}}
# Result: allow = true

# Input: {"method": "POST", "user": {"role": "admin"}}
# Result: allow = false
```

#### data - Static Data
**[Try on Playground](https://play.openpolicyagent.org/)**

```rego
package data_example

import rego.v1

# Configuration or master data
allow if {
    data.roles[input.user.role].permissions[_] == "read"
}

# Sample execution for testing
# Input: {"user": {"role": "admin"}}
# Result: allow = true

# Input: {"user": {"role": "guest"}}
# Result: allow = false
```

**JSON to paste in Data section**:
```json
{
    "roles": {
        "admin": {
            "permissions": ["read", "write", "delete"]
        },
        "user": {
            "permissions": ["read"]
        }
    }
}
```

## Chapter 3: Types of Rules and How to Write Them

### 3.1 Rule Syntax Patterns

Rego rules have several patterns:

| Pattern                           | Syntax Example                       | Result Type | Description          |
| --------------------------------- | ------------------------------------ | ----------- | -------------------- |
| `<name> := <value>`               | `pi := 3.14`                        | Single value | Complete rule        |
| `<name> if <body>`                | `valid if input.age >= 18`          | Boolean     | Conditional complete rule |
| `<name> contains <key> if <body>` | `users contains name if ...`         | Set         | Set rule             |
| `<name>[<key>] := <value> if <body>` | `scores[user] := score if ...`     | Object      | Object rule          |

### 3.2 Complete Rules

Rules that return a single value.

**[Try on Playground](https://play.openpolicyagent.org/)**

```rego
package quota

import rego.v1

# Constant value
default request_quota := 100

# Conditional value
request_quota := 1000 if input.user.internal

request_quota := 50 if input.user.plan == "trial"

# Sample execution for testing
# Input: {"user": {"plan": "basic"}}
# Result: request_quota = 100

# Input: {"user": {"internal": true}}
# Result: request_quota = 1000

# Input: {"user": {"plan": "trial"}}
# Result: request_quota = 50
```

**Evaluation Order**: If multiple rules match, the value of the last defined rule is adopted.

### 3.3 Partial Set Rules

Rules that generate sets.

**[Try on Playground](https://play.openpolicyagent.org/)**

```rego
package paths

import rego.v1

# Basic form
allowed_paths contains "/public"

# Conditional set generation
allowed_paths contains path if {
    some team in input.user.teams
    path := sprintf("/teams/%v/*", [team])
}

# Sample execution for testing
# Input: {"user": {"teams": ["engineering", "design"]}}
# Result: allowed_paths = {"/public", "/teams/engineering/*", "/teams/design/*"}

# Input: {"user": {}}
# Result: allowed_paths = {"/public"}
```

### 3.4 Partial Object Rules

Rules that generate objects (maps).

**[Try on Playground](https://play.openpolicyagent.org/)**

```rego
package object_rules

import rego.v1

# Group paths by prefix
paths_by_prefix[prefix] := paths if {
    some path in input.paths
    parts := split(path, "/")
    count(parts) > 0
    prefix := parts[0]  # Bind prefix to a specific value

    paths := [p |
        some p in input.paths
        p_parts := split(p, "/")
        count(p_parts) > 0
        p_parts[0] == prefix
    ]
}

# Sample execution for testing
# Input: {"paths": ["admin/users", "admin/settings", "public/docs", "public/help"]}
# Result: paths_by_prefix = {"admin": ["admin/users", "admin/settings"], "public": ["public/docs", "public/help"]}

# Input: {"paths": ["home/dashboard"]}
# Result: paths_by_prefix = {"home": ["home/dashboard"]}
```

## Chapter 4: Control Structures and Operators

### 4.1 Comparison Operators

| Operator | Description | Example                      |
| -------- | ----------- | ---------------------------- |
| `==`     | Equal       | `input.role == "admin"`    |
| `!=`     | Not equal   | `input.status != "banned"` |
| `<`      | Less than   | `input.age < 65`            |
| `<=`     | Less or equal | `input.score <= 100`     |
| `>`      | Greater than | `input.salary > 50000`    |
| `>=`     | Greater or equal | `input.age >= 18`     |

### 4.2 Logical Operators

#### AND (Logical Conjunction)
**[Try on Playground](https://play.openpolicyagent.org/)**

```rego
package logical_and

import rego.v1

default allow := false

allow if {
    input.user.role == "admin"  # AND
    input.user.active == true   # AND
    input.method == "GET"       # AND
}

# Sample execution for testing
# Input: {"user": {"role": "admin", "active": true}, "method": "GET"}
# Result: allow = true

# Input: {"user": {"role": "admin", "active": false}, "method": "GET"}
# Result: allow = false
```

#### OR (Logical Disjunction)
**[Try on Playground](https://play.openpolicyagent.org/)**

```rego
package logical_or

import rego.v1

default allow := false

# Express OR with multiple rules
allow if {
    input.user.role == "admin"
}

allow if {
    input.user.role == "manager"
    input.action == "read"
}

# Sample execution for testing
# Input: {"user": {"role": "admin"}}
# Result: allow = true

# Input: {"user": {"role": "manager"}, "action": "read"}
# Result: allow = true

# Input: {"user": {"role": "user"}, "action": "read"}
# Result: allow = false
```

#### NOT (Negation)
**[Try on Playground](https://play.openpolicyagent.org/)**

```rego
package logical_not

import rego.v1

default allow := false

allow if {
    input.user.role == "user"
    not input.user.banned  # If not banned
}

# Sample execution for testing
# Input: {"user": {"role": "user", "banned": false}}
# Result: allow = true

# Input: {"user": {"role": "user", "banned": true}}
# Result: allow = false

# Input: {"user": {"role": "user"}}
# Result: allow = true
```

### 4.3 Quantification

#### some - Existential Quantification
Returns true if "at least one satisfies the condition."

**[Try on Playground](https://play.openpolicyagent.org/)**

```rego
package quantification_some

import rego.v1

# Basic form
default has_admin_role := false

has_admin_role if {
    some i
    input.user.roles[i] == "admin"
}

# Simplified form
default has_admin_role_short := false

has_admin_role_short if {
    input.user.roles[_] == "admin"
}

# Combined with variable binding
default allowed_action := false

allowed_action if {
    some role in input.user.roles
    role == "admin"  # Simplified
}

# Sample execution for testing
# Input: {"user": {"roles": ["user", "admin"]}}
# Result: has_admin_role = true, has_admin_role_short = true, allowed_action = true

# Input: {"user": {"roles": ["user", "guest"]}}
# Result: has_admin_role = false, has_admin_role_short = false, allowed_action = false
```

#### every - Universal Quantification
Returns true if "all satisfy the condition."

**[Try on Playground](https://play.openpolicyagent.org/)**

```rego
package quantification_every

import rego.v1

default all_files_owned := false

all_files_owned if {
    # Check existence of input data
    input.paths
    input.user.id
    is_array(input.paths)

    every path in input.paths {
        startswith(path, sprintf("/users/%v/", [input.user.id]))
    }
}

# Sample execution for testing
# Input: {"user": {"id": "u123"}, "paths": ["/users/u123/file1.txt", "/users/u123/file2.txt"]}
# Result: all_files_owned = true

# Input: {"user": {"id": "u123"}, "paths": ["/users/u123/file1.txt", "/public/file2.txt"]}
# Result: all_files_owned = false
```

### 4.4 in Operator
**[Try on Playground](https://play.openpolicyagent.org/)**

```rego
package in_operator

import rego.v1

# Membership in a set
valid_method if {
    input.method in {"GET", "POST", "PUT"}
}

# Membership in an array
valid_role if {
    input.user.role in ["admin", "manager", "user"]
}

# Sample execution for testing
# Input: {"method": "GET", "user": {"role": "admin"}}
# Result: valid_method = true, valid_role = true

# Input: {"method": "DELETE", "user": {"role": "guest"}}
# Result: valid_method = false, valid_role = false
```

## Chapter 5: Comprehensions and Data Manipulation

### 5.1 Types of Comprehensions

Comprehensions are syntax for generating new collections from existing data.

| Type           | Syntax                   | Features           | Use Case                |
| -------------- | ------------------------ | ------------------ | ----------------------- |
| Array          | `[term \| body]`        | Ordered, duplicates allowed | List processing         |
| Set            | `{term \| body}`        | Unordered, duplicates removed | Unique value collection |
| Object         | `{key: value \| body}`  | Key-value pairs    | Structured data transformation |

### 5.2 Array Comprehension

**[Try on Playground](https://play.openpolicyagent.org/)**

```rego
package array_comprehension

import rego.v1

# Basic form: Double the numbers
doubled := [result |
    some x in input.numbers
    result := x * 2
]

# Conditional: Double only even numbers
doubled_evens := [result |
    some x in input.numbers
    x % 2 == 0
    result := x * 2
]

# Complex transformation: Extract user names
user_names := [name |
    some user in input.users
    user.active == true
    name := user.name
]

# Sample execution for testing
# Input: {"numbers": [1, 2, 3, 4, 5], "users": [{"name": "Alice", "active": true}, {"name": "Bob", "active": false}, {"name": "Carol", "active": true}]}
# Result:
# doubled = [2, 4, 6, 8, 10]
# doubled_evens = [4, 8]
# user_names = ["Alice", "Carol"]
```

### 5.3 Set Comprehension

**[Try on Playground](https://play.openpolicyagent.org/)**

```rego
package set_comprehension

import rego.v1

# Automatically remove duplicates
unique_roles := {role |
    some user in input.users
    some role in user.roles
}

# Conditional filter
admin_users := {user.name |
    user := input.users[_]
    user.role == "admin"
}

# Sample execution for testing
# Input: {"users": [{"name": "Alice", "role": "admin", "roles": ["admin", "user"]}, {"name": "Bob", "role": "user", "roles": ["user"]}, {"name": "Carol", "role": "admin", "roles": ["admin"]}]}
# Result:
# unique_roles = {"admin", "user"}
# admin_users = {"Alice", "Carol"}
```

### 5.4 Object Comprehension

**[Try on Playground](https://play.openpolicyagent.org/)**

```rego
package object_comprehension

import rego.v1

# Create a map with user IDs as keys
users_by_id := {user.id: user |
    user := input.users[_]
}

# Count number of people per department
dept_counts := {dept: count(members) |
    some dept in {user.department | user := input.users[_]}
    members := [user |
        user := input.users[_]
        user.department == dept
    ]
    count(members) > 0
}

# Sample execution for testing
# Input: {"users": [{"id": "u1", "name": "Alice", "department": "eng"}, {"id": "u2", "name": "Bob", "department": "eng"}, {"id": "u3", "name": "Carol", "department": "sales"}]}
# Result:
# users_by_id = {"u1": {"id": "u1", "name": "Alice", "department": "eng"}, "u2": {"id": "u2", "name": "Bob", "department": "eng"}, "u3": {"id": "u3", "name": "Carol", "department": "sales"}}
# dept_counts = {"eng": 2, "sales": 1}
```

### 5.5 Nested Comprehensions

**[Try on Playground](https://play.openpolicyagent.org/)**

```rego
package nested_comprehension

import rego.v1

# Active users per department
active_by_dept := {dept: active_users |
    some dept in {user.department | user := input.users[_]}
    active_users := [user.name |
        user := input.users[_]
        user.department == dept
        user.active == true
    ]
}

# Sample execution for testing
# Input: {"users": [{"name": "Alice", "department": "eng", "active": true}, {"name": "Bob", "department": "eng", "active": false}, {"name": "Carol", "department": "sales", "active": true}]}
# Result:
# active_by_dept = {"eng": ["Alice"], "sales": ["Carol"]}
```

---

## Chapter 6: Built-in Functions

Rego provides a rich set of built-in functions. The availability of functions depends on the execution environment (Go SDK, WebAssembly, etc.), so refer to the official reference for details.

- **[Built-in Functions Reference](https://www.openpolicyagent.org/docs/policy-reference/builtins)**
  - Detailed specifications of all built-in functions
  - Environment-specific support (Wasm / SDK-dependent)
  - Usage examples and parameter descriptions

## Chapter 7: Testing and Debugging

### 7.1 Basics of Testing

Rego tests are written as functions starting with `test_`.

#### Basic Test Syntax
**[Try on Playground](https://play.openpolicyagent.org/)**

```rego
package example

import rego.v1

# Policy to be tested
default allow := false

allow if {
    input.user.role == "admin"
}

# === Test Code ===

# Successful test
test_admin_is_allowed if {
    allow with input as {"user": {"role": "admin"}}
}

# Failing test (using not)
test_guest_is_denied if {
    not allow with input as {"user": {"role": "guest"}}
}

# Test results
# allow = false
# test_admin_is_allowed = true
# test_guest_is_denied = true
```

### 7.2 Overriding Values with with

Use the `with` keyword to override `input` or `data` during testing.

**[Try on Playground](https://play.openpolicyagent.org/)**

```rego
package with_example

import rego.v1

# Policy to be tested
allow if {
    input.user.role == "admin"
    data.config.strict_mode == false
}

# === Test Code ===

# Override input
test_with_input if {
    allow with input as {"user": {"role": "admin"}}
          with data.config.strict_mode as false
}

# Override data
test_with_data if {
    allow with input as {"user": {"role": "admin"}}
          with data as {"config": {"strict_mode": false}}
}

# Override multiple values simultaneously
test_with_multiple if {
    allow with input as {"user": {"role": "admin"}}
          with data.config.strict_mode as false
}

# Test results
# allow = false
# test_with_input = true
# test_with_data = true
# test_with_multiple = true
```

### 7.3 Comprehensive Test Cases

```rego
package authz_test

import data.authz

# Normal test

test_admin_user_allowed {
    authz.allow with input as {
        "user": {"role": "admin", "active": true},
        "action": "read"
    }
}

test_user_with_permission_allowed {
    authz.allow with input as {
        "user": {"role": "user", "active": true},
        "action": "read"
    } with data as {
        "permissions": {
            "user": ["read"]
        }
    }
}

# Abnormal test

test_inactive_user_denied {
    not authz.allow with input as {
        "user": {"role": "admin", "active": false},
        "action": "read"
    }
}

test_insufficient_permission_denied {
    not authz.allow with input as {
        "user": {"role": "user", "active": true},
        "action": "delete"
    } with data as {
        "permissions": {
            "user": ["read"]
        }
    }
}

# Edge case test

test_empty_input_denied {
    not authz.allow with input as {}
}

test_missing_role_denied {
    not authz.allow with input as {
        "user": {"active": true},
        "action": "read"
    }
}
```

### 7.4 Debugging Techniques

#### Debugging with print
```rego
allow if {
    user_role := input.user.role
    print("User role:", user_role)  # Debug output

    permissions := data.roles[user_role].permissions
    print("Permissions:", permissions)  # Debug output

    permissions[_] == input.action
}
```

#### trace Function (More Detailed Output)
```rego
allow if {
    trace(sprintf("Evaluating access for user: %v", [input.user.name]))
    input.user.role == "admin"
}
```

### 7.5 Running Tests

```bash
# Run all tests
opa test .

# Verbose output
opa test . -v

# Test specific file
opa test policy_test.rego

# With coverage information
opa test . --coverage
```

## References

- [OPA Official Site](https://www.openpolicyagent.org/)
- [Rego Language Reference](https://www.openpolicyagent.org/docs/latest/policy-language/)
- [Rego Cheat Sheet](https://docs.styra.com/opa/rego-cheat-sheet)
- [Introduction to OPA/Rego (Zenn)](https://zenn.dev/mizutani/books/d2f1440cfbba94)