---
title: What is Open Policy Agent (OPA)? A Practical Guide to Policy as Code
slug: open-policy-agent
date: 2025-05-13T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Open Policy Agent
  - Access Control
translation_key: about-open-policy-agent
---

I wanted to learn more about Open Policy Agent, so I referred to the official documentation.

# What is Open Policy Agent
Open Policy Agent (OPA, pronounced "Opa") is an open-source general-purpose policy engine for policy enforcement.

It allows you to write policies as code using a declarative language called Rego (pronounced "Ray-go").

Developed by [Styra](https://www.styra.com/), it is still being developed by Styra and the OSS community.

It is recognized as a Graduated project by the Cloud Native Computing Foundation (CNCF).

# Overview
OPA separates policy decision from policy enforcement.

It takes any structured data as input, applies policies, and outputs results.

The structured data can be in any format such as JSON, YAML, CSV, or XML.

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

OPA makes policy decisions by matching queries, policies, and data as input.

Since OPA and policies (Rego) are not domain-specific, they can describe various conditions in a general-purpose policy engine.

OPA can be used as a sidecar, a host-level daemon, or a library.

# Policy Decoupling
By separating policies into policy decisions and policy enforcement, OPA improves the reusability, testability, and deployment independence of policies.

As a result, it enhances adaptability to business requirements and increases the ability to detect policy violations and conflicts.

# Policies and Authorization
OPA defines policies and authorization as follows:

- A policy is a rule that controls the behavior of software.
- Authorization is a policy that defines which resources a user or system can perform actions on.

The concept of policy includes authorization. Authorization is a more generalized form of policy.

# OPA's Document Model

OPA handles three types of data for policy evaluation:

## Base Documents

- Structured data (in JSON format) loaded from external sources into OPA.
- Mainly referenced by `data` and `input`.
- Classified by loading method:

| Model     | Reference in Rego | Loading Method                    |
| ------- | -------- | ------------------------ |
| Asynchronous Push | `data`   | OPA API (e.g., PUT /v1/data) |
| Asynchronous Pull   | `data`   | Bundle feature                   |
| Synchronous Push  | `input`  | Passed directly during request (e.g., POST)     |
| Synchronous Pull    | Local Variable   | Retrieved using built-in functions like `http.send`  |

### Virtual Documents

- Evaluation results of rules defined in Rego policies.
- Not loaded from external sources but generated during evaluation.
- Can be referenced similarly to base documents, e.g., `data.foo.bar`.

### Unified Reference and Syntax

- Base documents and virtual documents can coexist under `data`.
- Both can have the same types such as numbers, strings, maps, and lists, and can be handled with the same syntax.
- `input` is considered a temporary base document specific to the request.

### Performance and Caching

- Asynchronously loaded `data` and policies are cached internally in OPA.
- Synchronously retrieved data via `http.send` is also treated as a local variable and can be cached as needed.

# Policy Language
OPA uses a declarative language called Rego for policy writing.

Rego allows you to focus on the logic of what results a query should return rather than how to execute the query.

An example of a policy that performs role-based access control (RBAC) is as follows:

This policy can achieve the following:
- Access control based on roles
- Permission for actions at the resource level

The following is sample code that can be executed in the [Rego Playground](https://play.openpolicyagent.org/).

```rego
// Policy: Policy Definition
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

OPA also provides a framework for testing, allowing you to test policies in isolation.

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

You can easily implement tests by simply showing input values and expected values.

It also supports parameterized tests, data-driven tests, mocks, benchmarks, and coverage.

For details on the Rego specification, refer to [Policy Language](https://www.openpolicyagent.org/docs/latest/policy-language/).

For what can be defined in policies, refer to [Policy Reference](https://www.openpolicyagent.org/docs/latest/policy-reference/).

The [Introduction to OPA/Rego](https://zenn.dev/mizutani/books/d2f1440cfbba94/viewer/chap-rego) provides a clear explanation of the basics of Rego.

# Performance
OPA is optimized internally using a rule indexing mechanism based on Trie.

cf. [Optimizing OPA: Rule Indexing](https://blog.openpolicyagent.org/optimizing-opa-rule-indexing-59f03f17caf3)

By writing policies with the Trie data structure in mind, you can prevent performance degradation (probably).

Actual points of caution are listed in [Policy Performance](https://www.openpolicyagent.org/docs/latest/policy-performance/).

Whether to use OPA as a library or to set up an OPA server affects how concurrency is handled.

> If you are embedding OPA as a library, it is your responsibility to dispatch concurrent queries to different Goroutines/threads. If you are running the OPA server, it will parallelize concurrent requests and use as many cores as possible.

cf. [policy-performance/#resource-utilization](https://www.openpolicyagent.org/docs/latest/policy-performance/#resource-utilization)

When using it as a library, you need to handle concurrency. When using the OPA server, concurrency is handled automatically. You need to choose based on performance requirements.

There is an option to specify an optimization level called `--optimize`, which allows you to adjust the time and resources spent on optimizing `opa bundle`. By optimizing the size of policies, you can reduce memory usage during policy loading and evaluation. This can be useful for large policies or when strict performance tuning is required.

# OPA Build Patterns
Here is a summary of patterns for handling external data.

| Approach                | Performance/Availability | Update Frequency※1               | Size Limit※2           | Recommended Data       | Security                         | Notes                                                                                                 |
| ------------------------- | --------------------- | ---------------------- | -------------------- | ---------------------- | ------------------------------------ | ---------------------------------------------------------------------------------------------------- |
| **JWT Token**           | High                  | When the user re-logs in | Limited             | User Attributes           | Token signature verification, TTL check        | Updated only during user authentication. Not suitable if attributes change frequently.                                           |
| **`input` Overload** | High                  | Depends on update frequency         | Few size restrictions | Dynamic local data | Security between service and OPA    | Suitable for dynamic and frequently updated data. Responsibility for incorporating external data lies with the developer.                     |
| **Bundle API**           | High                  | Low                   | Size limit applies       | Static medium-sized data     | API access control                    | Synchronizes external data as a bundle. Optimal if the data is static.                                     |
| **Push Data**        | High                  | High                   | Size limit applies       | Dynamic medium-sized data     | API access control                    | Allows fine-tuning of the update frequency for external data. Suitable for data that is frequently updated.                           |
| **Data Retrieval at Evaluation**    | Network-dependent      | Always up-to-date               | No size limit       | Large or dynamic data | Security for access to external services | Suitable for external data that is updated very frequently or when the data volume is large. Network availability is key. |

※1 Update refers to the frequency and ease with which changes to external data are reflected in OPA. When data changes frequently, real-time reflection and mechanisms become important.

※2 Size refers to the amount of external data that OPA can handle, influencing whether it can be kept in memory or needs to be passed with each request. For large data, on-demand retrieval becomes realistic.

Refer to [External Data](https://www.openpolicyagent.org/docs/latest/external-data/).

# Monitoring
Supports Open Telemetry and Prometheus.

Refer to [Monitoring](https://www.openpolicyagent.org/docs/latest/monitoring/).

# References
- [www.openpolicyagent.org](https://www.openpolicyagent.org/docs/latest/)