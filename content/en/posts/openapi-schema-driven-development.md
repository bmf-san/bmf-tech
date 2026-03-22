---
title: Schema-Driven Development with OpenAPI
description: 'Learn OpenAPI schema-driven development, REST API specification, code generation, and mock server benefits.'
slug: openapi-schema-driven-development
date: 2024-01-19T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - OpenAPI
  - Schema-Driven
translation_key: openapi-schema-driven-development
---


# What is OpenAPI Specification?
A format for defining HTTP API specifications in a language-agnostic way. It is written in YAML or JSON.

The Swagger Specification is the predecessor of the OpenAPI Specification.

# Advantages and Disadvantages of Adopting OpenAPI
## Advantages
- It is a standardized format for REST APIs, which can reduce communication costs among developers.
- There is a rich ecosystem of tools (benefits of standardization).
- Automatic generation of API documentation is possible.
- Automation of code generation for API clients and servers can be achieved.
- Being text-based makes it easier to manage when combined with version control systems.
- You can start a mock server, allowing development to proceed without waiting for API implementation.

It enables the efficiency of the development process through schema-driven approaches.

## Disadvantages
- There is a learning cost for understanding the specifications when adopting it for the first time (which can also affect the completeness of the design).
- If the specifications and implementation are not kept in sync, discrepancies may arise between them.
- Since the OpenAPI specification is evolving, there is a cost associated with keeping up with the standards.
- Depending on the tools used, unexpected behaviors or differences in interpretation of the specifications may occur.

It seems that there are no critical disadvantages.

# Trying It Out
Since Docker is available, let's try it with Docker.

cf. [github.com - OpenAPITools/openapi-generator](https://github.com/OpenAPITools/openapi-generator?tab=readme-ov-file#openapi-generator-cli-docker-image)

```sh
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i https://raw.githubusercontent.com/openapitools/openapi-generator/master/modules/openapi-generator/src/test/resources/3_0/petstore.yaml \
    -g go \
    -o /local/out/go
```

For API specification documentation, the [OpenAPI (Swagger) Editor](https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi) seems to be a good choice if you are using VSCode.

# Thoughts
I haven't properly looked at the generation options, so I think I will try using it in a personal development project once I can confirm those.

I felt that effectively adjusting what to auto-generate and what not to auto-generate is a challenge during the introduction phase.

# References
- [www.openapis.org](https://www.openapis.org/)
- [zenn.dev - Mastering Schema-Driven Development with OpenAPI Generator in Go](https://zenn.dev/ysk1to/books/248fad8cb34abe)
- [medium.com - Generating Go code from OpenAPI Specification Document](https://medium.com/@MikeMwita/generating-go-code-from-openapi-specification-document-ae225e49e970)
