---
title: About Open Policy Agent
slug: about-open-policy-agent
date: 2025-05-13T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Open Policy Agent
  - Authorization Management
description: A detailed exploration of Open Policy Agent based on its official documentation.
translation_key: about-open-policy-agent
---

I wanted to learn more about Open Policy Agent, so I referred to its official documentation.

# What is Open Policy Agent?
Open Policy Agent (OPA, pronounced "Opa") is an open-source, general-purpose policy engine for policy enforcement.

Using a declarative language called Rego (pronounced "Rego"), you can write policies as code.

It is developed by [Styra](https://www.styra.com/) and continues to be maintained by Styra and the open-source community.

OPA is recognized as a Graduated project by the Cloud Native Computing Foundation (CNCF).

# Overview
OPA separates policy decision-making from policy enforcement.

It accepts any structured data as input, applies policies, and outputs results.

Structured data can be in any format, such as JSON, YAML, CSV, or XML.

```text
                +---------------------+
                |      Service        |
                +---------------------+
                     ↑           ↑
      Request, Event |           | Decision (any JSON value)
                     |           |
                     v           |
                +---------------------+
                |         OPA         |
                +---------------------+
                     ↑           ↑
         Query (any  |           |
         JSON value) |           |
                     |           |
          +----------+-----------+----------+
          |                                 |
  +---------------+              +------------------+
  |   Policy       |              |      Data        |
  |   (Rego)       |              |    (JSON)        |
  +---------------+              +------------------+
```

OPA makes policy decisions by matching queries, policies, and data.

Since OPA and its policies (Rego) are domain-agnostic, it can serve as a general-purpose policy engine capable of describing various conditions.

OPA can be used as a sidecar, a host-level daemon, or a library.

# Policy Decoupling
By separating policy decision-making from policy enforcement, OPA improves policy reusability, testability, and deployment independence.

As a result, it enhances adaptability to business requirements and improves the ability to detect policy violations or conflicts.

# Policies and Authorization
OPA defines policies and authorization as follows:

- A policy is a rule that controls the behavior of software.
- Authorization is a type of policy that defines which actions a user or system can perform on specific resources.

Authorization is a subset of the broader concept of policies.

# OPA's Document Model

OPA handles three types of data for policy evaluation:

## Base Documents

- Structured data (in JSON format) loaded into OPA from external sources.
- Primarily referenced as `data` and `input`.
- Categorized by loading method:

| Model          | Referenced in Rego | Loading Method                |
| -------------- | ------------------ | ----------------------------- |
| Async Push     | `data`             | OPA API (e.g., PUT /v1/data)  |
| Async Pull     | `data`             | Bundle feature                |
| Sync Push      | `input`            | Directly passed during request (e.g., POST) |
| Sync Pull      | Local variables    | Retrieved via built-in functions like `http.send` |

### Virtual Documents

- Evaluation results of rules defined in Rego policies.
- Not loaded externally but generated during evaluation.
- Referenced similarly to base documents, e.g., `data.foo.bar`.

### Unified References and Syntax

- Base and virtual documents coexist under `data`.
- Both share the same types (e.g., numbers, strings, maps, lists) and syntax.
- `input` is treated as a temporary base document specific to a request.

### Performance and Caching

- Asynchronously loaded `data` and policies are cached within OPA.
- Synchronously retrieved data via `http.send` can also be cached as local variables.

# Policy Language
OPA uses a declarative language called Rego to write policies.

Rego focuses on the results of queries rather than how the queries are executed, allowing for concise logic.

Here’s an example of a policy for role-based access control (RBAC):

```rego
// Policy: Define policies
package example

default allow = false

allow if {
    perms := data.role_permissions[input.user.role][input.resource]
    perms[_] == input.action
}

// INPUT: Input data referenced during policy execution
{
  "user": {
    "id": "u001",
    "role": "editor"
  },
  "resource": "projects",
  "action": "write"
}

// DATA: Data referenced during policy execution
{
  "role_permissions": {
    "viewer": {
      "projects": ["read"]
    },
    "editor": {
      "projects": ["read", "write"]
    },
    "admin": {
      "projects": ["read", "write", "delete"]
    }
  }
}

// OUTPUT: Result of policy execution
{
    "allow": true
}
```

OPA also provides a framework for testing policies, enabling unit tests for individual policies.

```rego
package example

test_allow_editor_write {
    input := {
        "user": {
            "id": "u001",
            "role": "editor"
        },
        "resource": "projects",
        "action": "write"
    }
    data := {
        "role_permissions": {
            "viewer": {
                "projects": ["read"]
            },
            "editor": {
                "projects": ["read", "write"]
            },
            "admin": {
                "projects": ["read", "write", "delete"]
            }
        }
    }
    allow with input as input with data as data
}
```

Testing is simple, requiring only input values and expected results.

Rego supports parameterized tests, data-driven tests, mocking, benchmarking, and coverage analysis.

For more details on Rego, refer to [Policy Language](https://www.openpolicyagent.org/docs/latest/policy-language/).

For what can be defined in policies, see [Policy Reference](https://www.openpolicyagent.org/docs/latest/policy-reference/).

For a beginner-friendly guide to Rego, check out [OPA/Rego Introduction](https://zenn.dev/mizutani/books/d2f1440cfbba94/viewer/chap-rego).

# Performance
OPA optimizes rules internally using a Trie-based index.

cf. [Optimizing OPA: Rule Indexing](https://blog.openpolicyagent.org/optimizing-opa-rule-indexing-59f03f17caf3)

Writing policies with Trie structures in mind can help prevent performance degradation.

For practical considerations, see [Policy Performance](https://www.openpolicyagent.org/docs/latest/policy-performance/).

Whether to use OPA as a library or as a server depends on parallelization needs:

> If you are embedding OPA as a library, it is your responsibility to dispatch concurrent queries to different Goroutines/threads. If you are running the OPA server, it will parallelize concurrent requests and use as many cores as possible.

cf. [policy-performance/#resource-utilization](https://www.openpolicyagent.org/docs/latest/policy-performance/#resource-utilization)

The `--optimize` option allows you to adjust the time and resources spent optimizing `opa bundle`. Optimized policies reduce memory usage during loading and evaluation, making it useful for large policies or performance-critical scenarios.

# OPA Deployment Patterns
Here’s a summary of patterns for handling external data:

| Approach               | Performance/Availability | Update Frequency*1      | Size Limit*2          | Recommended Data      | Security                          | Notes                                                                                               |
| ---------------------- | ------------------------ | ----------------------- | --------------------- | --------------------- | --------------------------------- | --------------------------------------------------------------------------------------------------- |
| **JWT Token**          | High                    | User re-login           | Limited               | User attributes       | Token signature verification, TTL | Updated only during user authentication. Unsuitable for frequently changing attributes.            |
| **Input Overload**     | High                    | Depends on update freq. | Minimal size limits   | Local dynamic data    | Secure service-to-OPA connection | Suitable for dynamic, frequently updated data. External data handling is the developer's responsibility. |
| **Bundle API**         | High                    | Low                     | Limited size          | Static medium data    | API access control               | Synchronizes external data as a bundle. Ideal for static data.                                     |
| **Push Data**          | High                    | High                    | Limited size          | Dynamic medium data   | API access control               | Allows fine-grained control over external data updates. Suitable for frequently updated data.       |
| **Data Fetch on Eval** | Network-dependent       | Always up-to-date       | No size limit         | Large or dynamic data | Secure external service access   | Suitable for frequently updated or large data. Network availability is critical.                   |

*1 Update frequency refers to how often or easily external data changes are reflected in OPA. Real-time updates or mechanisms to reflect changes are important for frequently changing data.

*2 Size refers to the amount of external data OPA handles, influencing whether it can be stored in memory or needs to be fetched per request. On-demand fetching is more practical for large datasets.

For more details, see [External Data](https://www.openpolicyagent.org/docs/latest/external-data/).

# Monitoring
OPA supports OpenTelemetry and Prometheus.

Refer to [Monitoring](https://www.openpolicyagent.org/docs/latest/monitoring/).

# References
- [www.openpolicyagent.org](https://www.openpolicyagent.org/docs/latest/)