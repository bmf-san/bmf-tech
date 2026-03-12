---
title: Creating a mysqldump Tool with Go
description: A step-by-step guide on Creating a mysqldump Tool with Go, with practical examples and configuration tips.
slug: creating-mysqldump-tool-with-go
date: 2019-02-04T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - MySQL
  - SSH
  - mysqldump
translation_key: creating-mysqldump-tool-with-go
---

# Overview
Since I was manually backing up the database for this blog like a caveman, I decided to create a tool in Go that can back up the database from remote to local with a single command.

# Packages
- "net"
  - Provides interfaces for network I/O, TCP/IP, UDP, domain name resolution, Unix domain sockets, etc.
- "time"
  - Provides functions for time calculations and display.
- "io/ioutil"
  - Provides I/O utilities for file operations.
- "golang.org/x/crypto/ssh"
  - Provides implementations for SSH clients and servers.
- "github.com/BurntSushi/toml"
  - TOML parser
  - It seems to be designed similarly to Go's standard library parsers for JSON and XML.
  - **Burnt Sushi**

# Implementation
I implemented it roughly to a working state. Since I'm not familiar with Go, it feels a bit clumsy...
Also, I haven't written tests yet.

```go
package main

import (
	"net"
	"time"
	"io/ioutil"
	"golang.org/x/crypto/ssh"
	"github.com/BurntSushi/toml"
)

type Config struct {
	SSH SSH
	Mysql Mysql
}

type SSH struct {
	IP string
	Port string
	User string
	IdentityFile string
}

type Mysql struct {
	MysqlConf string
	Database string
	DumpDir string
	DumpFilePrefix string
}

func dump() {
	var config Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		panic(err)
	}
	
	buf, err := ioutil.ReadFile(config.SSH.IdentityFile)
	if err != nil {
		panic(err)
	}
	
	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		panic(err)
	}
	
	conn, err := ssh.Dial("tcp", config.SSH.IP+":"+config.SSH.Port, &ssh.ClientConfig{
		User: config.SSH.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	})
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	
	session, err := conn.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	
	byte, err := session.Output("sudo mysqldump --defaults-file="+config.Mysql.MysqlConf+" "+config.Mysql.Database+" --quick --single-transaction")
	if err != nil {
		panic(err)
	}
	
	ioutil.WriteFile(config.Mysql.DumpDir+config.Mysql.DumpFilePrefix+time.Now().Format("2006-01-02")+".sql", byte, 0644)
}

func main() {
	dump()
}
```

# GitHub
Left it here.

- [github - bmf-san/go-mysqldump](https://github.com/bmf-san/go-mysqldump)

# Thoughts
For now, I will look at various implementations in Go and accumulate knowledge...

# References
- [GoDoc - package mysql](https://godoc.org/github.com/go-sql-driver/mysql)
- [GoDoc - package ssh](https://godoc.org/golang.org/x/crypto/ssh#Request)
- [Golang.org - Package net](https://golang.org/pkg/net/)
- [Golang.org - Package time](https://golang.org/pkg/time/)
- [Golang.org - Package ioutil](https://golang.org/pkg/io/ioutil/)
- [golang.jp - net package](http://golang.jp/pkg/net)
- [Using SSH with Authentication Keys in Go](https://saitodev.co/article/Go%E8%A8%80%E8%AA%9E%E3%81%A7%E8%AA%8D%E8%A8%BC%E9%8D%B5%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6SSH%E3%81%AE%E6%8E%A5%E7%B6%9A%E3%82%92%E8%A1%8C%E3%81%86)
- [github.com - siddontang/go-mysql](https://github.com/siddontang/go-mysql/blob/master/dump/dump.go)
- [Mysqldump Through a HTTP Request with Golang](https://intelligentbee.com/2017/09/04/mysqldump-through-http-request-golang/)