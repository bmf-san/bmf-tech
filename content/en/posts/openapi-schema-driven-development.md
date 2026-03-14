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

# What is the OpenAPI Specification
A format for defining HTTP API specifications in a language-agnostic way. It is written in YAML or JSON.

Swagger Specification is the predecessor of the OpenAPI Specification.

# Advantages and Disadvantages of Adopting OpenAPI
## Advantages
- It is a standardized format for REST APIs, which can reduce communication costs among developers
- The tools and ecosystem are well-developed (benefits of standardization)
- Automatic generation of API documentation is possible
- Automation of client and server code generation for APIs is possible
- Being text-based, it becomes easier to manage when combined with version control systems
- A mock server can be launched, allowing development to proceed without waiting for API implementation

Schema-driven development can streamline the development process.

## Disadvantages
- There is a learning cost to understand the specifications when introducing it for the first time (also affects the degree of design completeness)
- If the specification and implementation are not kept in sync, discrepancies may occur
- As the OpenAPI specification evolves, there is a cost to keep up with the standard
- Depending on the tools used, unexpected behavior or differences in interpretation of the specifications may occur

There do not seem to be any critical disadvantages.

# Trying It Out
Since Docker can be used, let's try it with Docker.

cf. [github.com - OpenAPITools/openapi-generator](https://github.com/OpenAPITools/openapi-generator?tab=readme-ov-file#openapi-generator-cli-docker-image)

```sh
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i https://raw.githubusercontent.com/openapitools/openapi-generator/master/modules/openapi-generator/src/test/resources/3_0/petstore.yaml \
    -g go \
    -o /local/out/go
```

For API specification documents, [OpenAPI (Swagger) Editor](https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi) seemed good for VSCode.

# Impressions
I haven't fully explored the generation options, so once I can confirm those, I plan to use it in a personal development project.

I felt that the challenge during introduction is to skillfully adjust what to auto-generate and what not to.

# References
- [www.openapis.org](https://www.openapis.org/)
- [zenn.dev - 【Go言語】OpenAPI Generatorを使いこなすスキーマ駆動開発](https://zenn.dev/ysk1to/books/248fad8cb34abe)
- [medium.com - Generating Go code from OpenAPI Specification Document](https://medium.com/@MikeMwita/generating-go-code-from-openapi-specification-document-ae225e49e970)