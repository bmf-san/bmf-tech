---
title: "Goを学びなおす"
slug: "learning-go-again"
date: 2023-01-23
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Golang"
draft: false
---

# 概要
仕事でもプライベートでも何年かGoを触っているが、今一度このタイミングで学び直してみると効果的ではないかなと思って色々学び直した。
その際に読んだ記事をリストアップしておく。

# モチベーション
基本的なことの復習、仕様で拾い切れていなかった部分や新機能のキャッチアップ、tips周りを拾って、Goのコーディング力を上げるためのベースを鍛え直したい。

# 仕様理解
仕様理解に関連する記事をgo.devを中心に読み漁った。

- [go.dev - The Go Programming Language Specification](https://go.dev/ref/spec)
- [go.dev - The Go Memory Model](https://go.dev/ref/mem)
- [go.dev - Effective Go](https://go.dev/doc/effective_go)
- [go.dev - How to Write Go Code](https://go.dev/doc/code)
- [go.dev - Writing Web Applications](https://go.dev/doc/articles/wiki/)
- [go.dev - Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
- [go.dev - Case Studies](https://go.dev/solutions/case-studies)
- [go.dev - Use Cases](https://go.dev/solutions/use-cases)
- [go.dev - Policy](https://go.dev/security/policy)
- [go-tour-jp.appspot.com - Welcome to a tour of Go](https://go-tour-jp.appspot.com/list)
- [google.github.io - Go style](https://google.github.io/styleguide/go/)
- [github.com - CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)

# Generics
Go1.18で追加されたGenericsの仕様について今一度キャッチアップした。

- [go.dev - Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics)
- [go.dev - When To Use Generics](https://go.dev/blog/when-generics)

# GoのPros・Cons
Goの良いところや得意なところ、苦手なところや不得手なところってどこだろうというのを整理しようと思って読み漁った記事。

育ったきた土壌が違うと見方も変わるので、色んな人の意見を見ると為になる。　

自分自身は多くの言語に触れてきた身ではないので、言語の設計思想の深いところに触れた洞察はできないが、Goのシンプルさ（シンプルに見える、というほうが的確かもしれない）を保つ思想に特に好感を持っている。色んな書き方ができる機能性の高い言語を使っているときはどう書くべきかということに悩むこともあるが、Goの場合は素直に書けると感じていて、書いていて楽しさを感じる言語であると思っている。

- [zenn.dev - 改めて見直すGoの特徴](https://zenn.dev/nobonobo/scraps/cec2259ac330a8)
- [zenn.dev - Go言語が成功した理由](https://zenn.dev/takehiro0740/articles/b5ef4fc26e9ba6)
- [www.mobulous.com - GoLang- A Complete Details of All The Pros and Cons in Programming](https://www.mobulous.com/blog/golang-a-complete-details-of-all-the-pros-and-cons-in-programming/)
- [www.scalefocus.com - Why You Should Go with Go for Your Next Software Project](https://www.scalefocus.com/blog/why-you-should-go-with-go-for-your-next-software-project)
- [www.uptech.team - Best Practices: Why Use Golang For Your Project](https://www.uptech.team/blog/why-use-golang-for-your-project)
- [https://medium.com - Why Go: The benefits of Golang](https://medium.com/@julienetienne/why-go-the-benefits-of-golang-6c39ea6cff7e)
- [www.infoworld.com - What’s the Go programming language really good for?](https://www.infoworld.com/article/3198928/whats-the-go-programming-language-really-good-for.html)
- [builtin.com - Why Go? 8 Engineers Discuss Golang’s Advantages and How They Use It.
](https://builtin.com/software-engineering-perspectives/golang-advantages)
- [madappgang.com - When and Why Use Go in Software Development
](https://madappgang.com/blog/why-golang/)

# 本
3冊ほどピックアップして読んだ。他にも読もうかと思った本があるが、今回の目的に沿いそうな本を厳選した。（特に並行処理周りはいい加減履修しないと思っているが、それだけに集中する必要があると思ったので、別の機会とした。。。）

- [実用Go言語](https://amzn.to/3QUlWO3)
- [go言語による分散サービス](https://amzn.to/3qIwOEj)
- [Go言語100Tips 開発者にありがちな間違いへの対処法 (impress top gear)](https://amzn.to/3QXZt2F)

全部良い本だが、特におすすめしたいのは実用Go言語。

自分のようにある程度Goを触ってきたけど今一度知識を整理したいという場合に学びがある本だと思う。

cf. https://bmf-tech.com/posts/%e5%ae%9f%e7%94%a8Go%e8%a8%80%e8%aa%9e%e3%82%92%e8%aa%ad%e3%82%93%e3%81%a0

# 所感
記憶の引き出しに色々としまえたので、どこかで引き出すときがきたら役立つはず。

あとはまだ理解しきれていないことも多いので、またどこかの節目にでも学びなおす。



