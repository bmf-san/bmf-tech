---
title: "Goでmysqldumpツールをつくる"
slug: "creating-mysqldump-tool-with-go"
date: 2019-02-04
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Golang"
  - "MySQL"
  - "ssh"
  - "mysqldump"
draft: false
---

# 概要
このブログのDBバックアップを原始人のごとく手動でやっていたのでコマンド一発でバックアップをリモートからローカルにバックアップを取れるツールをgoでつくってみた。

# パッケージ
- "net"
  - ネットワークI/O、TCP/IP、UDP、ドメイン名前解決、Unixドメインソケットなどのインターフェースを提供してくれるやつ
- "time"
  - 時間の計算や表示のための機能を提供してくれるやつ
- "io/ioutil"
  - ファイル周りのI/Oユーティリティを提供してくれるやつ
- "golang.org/x/crypto/ssh"
  - sshのクライアント・サーバーの実装を提供してくれるやつ
- "github.com/BurntSushi/toml"
  - TOMLパーサー
  - goのjsonやxmlといったパーサーの標準ライブラリライクに作られているらしい
  - **焦げた寿司**さん

# 実装　
ざっくり動く形まで実装してみた。Goに不慣れなので愚直な感じになっている。。。
あとテストがかけていない。

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
	
  byte, err := session.Output("sudo mysqldump --defaults-file="+config.Mysql.MysqlConf+" "+config.Mysql.Database+" "+"--quick --single-transaction")
	if err != nil {
		panic(err)
	}
	
	ioutil.WriteFile(config.Mysql.DumpDir+config.Mysql.DumpFilePrefix+time.Now().Format("2006-01-02")+".sql", byte, 0644)
}

func main() {
	dump()
}
```

# github
置いといた。

- [github - bmf-san/go-mysqldump](https://github.com/bmf-san/go-mysqldump)

# 所感
とりあえずgoの色んな実装をみて知見を貯めていく...

# 参考
- [GoDoc - package mysql](https://godoc.org/github.com/go-sql-driver/mysql)
- [GoDoc - package ssh](https://godoc.org/golang.org/x/crypto/ssh#Request)
- [Golang.org - Package net](https://golang.org/pkg/net/)
- [Golang.org - Package time](https://golang.org/pkg/time/)
- [Golang.org - Package ioutil](https://golang.org/pkg/io/ioutil/)
- [golang.jp - netパッケージ](http://golang.jp/pkg/net)
- [Go言語で認証鍵を使ってSSHの接続を行う](https://saitodev.co/article/Go%E8%A8%80%E8%AA%9E%E3%81%A7%E8%AA%8D%E8%A8%BC%E9%8D%B5%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6SSH%E3%81%AE%E6%8E%A5%E7%B6%9A%E3%82%92%E8%A1%8C%E3%81%86)
- [github.com - siddontang/go-mysql](https://github.com/siddontang/go-mysql/blob/master/dump/dump.go)
- [Mysqldump Through a HTTP Request with Golang
](https://intelligentbee.com/2017/09/04/mysqldump-through-http-request-golang/)


