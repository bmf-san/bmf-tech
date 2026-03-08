---
title: Enabling Go Language Server in VSCode Prevented Code Definition Jumping
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
After enabling the Go Language Server settings in VSCode, I was unable to jump to code definitions, so I investigated the cause.

settings.json
```json
"go.useLanguageServer": true,
```

# Conclusion
It is necessary for `go.mod` to exist at the root of the project.

cf. [stackoverflow - How to properly use go modules in vscode?](https://stackoverflow.com/questions/59732657/how-to-properly-use-go-modules-in-vscode)

When opening a folder in VSCode, it should not be like this,

```
.
├── app
    ├── go.mod
```

Instead, it should be opened like this, otherwise the paths may not resolve correctly, preventing code jumping.
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

I thought it would interpret this correctly since it says `module github.com/bmf-san/gobel-api/app`, but it seems that is not the case.

# Investigation Method
I opened the terminal in VSCode and selected OUTPUT>gopls(server). When trying to jump to code, I was able to check the error logs.

From the error logs, it seemed the paths were suspicious, and after some investigation, I found the aforementioned StackOverflow post.

# Solution
The immediate countermeasures I could think of are as follows:

- Turn off the language server settings
- Place go.mod at the project root, or open the folder where go.mod exists

It might be possible to adjust settings in gopls or VSCode, but I couldn't find it quickly, so I temporarily addressed it by turning off the language server settings...

I don't think this is a settled configuration, so I believe similar cases or the best solutions will be found eventually...

I plan to update if I find out anything.

# Related
- [Big Sky - Stopping gocode (and moving to Language Server)](https://mattn.kaoriya.net/software/lang/go/20181217000056.htm)