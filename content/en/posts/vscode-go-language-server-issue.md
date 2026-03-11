---
title: Code Definition Jump Disabled After Enabling Go Language Server in VSCode
slug: vscode-go-language-server-issue
date: 2020-07-19T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - gocode
  - gopls
  - Language Server
  - vscode
  - Tips
translation_key: vscode-go-language-server-issue
---

# Overview
After enabling the Go Language Server settings in VSCode, I found that jumping to code definitions was no longer possible, so I investigated the cause.

settings.json
```json
"go.useLanguageServer": true,
```

# Conclusion
A `go.mod` file needs to exist at the root of the project.

cf. [stackoverflow - How to properly use go modules in vscode?](https://stackoverflow.com/questions/59732657/how-to-properly-use-go-modules-in-vscode)

When opening a folder in VSCode, instead of this,

```
.
├── app
    ├── go.mod
```

You need to open it like this, otherwise, the path won't be resolved correctly, preventing code jumps.
```
.
├── go.mod
```

For reference, here is the content of go.mod.
```go
module github.com/bmf-san/gobel-api/app

go 1.14

require (
	github.com/bmf-san/goblin v0.0.0-20200718124906-8b3133b538d6
	github.com/bmf-san/golem v0.0.0-20200718182453-066c8e70e46e
	github.com/go-redis/redis/v7 v7.4.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899
)
```

Since it says `module github.com/bmf-san/gobel-api/app`, I thought it would interpret it correctly, but apparently not.

# Investigation Method
Open the terminal in VSCode, select OUTPUT>gopls(server), and try jumping to code to check the error logs.

From the error logs, it seemed like the path was suspicious, and after some investigation, I found a relevant post on Stack Overflow.

# Solution
Here are some immediate countermeasures:

- Turn off the language server settings
- Place go.mod at the project root, or open the folder where go.mod exists

It might be possible to adjust with gopls or VSCode settings, but since I couldn't find a quick solution and it seemed time-consuming, I temporarily responded by turning off the language server settings...

I don't think this is a mature setting yet, so similar cases or the best solution might be found eventually...

I plan to update if I find anything.

# Related
- [Big Sky - Stopping gocode (and moving to Language Server)](https://mattn.kaoriya.net/software/lang/go/20181217000056.htm)