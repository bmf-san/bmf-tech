---
title: Obtain Let's Encrypt SSL Certificates with lego Using DNS-01 Method
slug: lego-lets-encrypt-dns-01-ssl
date: 2022-01-18T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Let's Encrypt
  - Tips
description: A guide on using lego to acquire Let's Encrypt SSL certificates via the DNS-01 method.
translation_key: lego-lets-encrypt-dns-01-ssl
---

Conoha VPSでAnsibleを使ってLet's EncryptのSSL証明書の取得を行おうとしていた。

証明書の取得は[DNS-01](https://datatracker.ietf.org/doc/html/draft-ietf-acme-acme-03#section-7.4)方式（ドメインのTXTレコードに認証局が発行したワンタイムトークンを登録して検証する）で取得したかったので、ConohaのAPIを使って、TXTレコードを登録、削除するようなスクリプトを組んで対応（cf. [github.com - k2snow/letsencrypt-dns-conoha](https://github.com/k2snow/letsencrypt-dns-conoha)）していたが、スクリプトの管理が面倒だったので、もっと単純なやり方を模索していたところ、[go-acme/lego](https://github.com/go-acme/lego)というLets' Encryptのクライアントツールを見つけたので使ってみた。

# go-acme/lego
legoはGo製のLets't Encryptクライアント&ACMEのライブラリ。

Conoha以外にも様々な[DNS Providers](https://go-acme.github.io/lego/dns/)が用意されている。

インストールはdockerでもパッケージマネージャーでもGoでも良い。

# legoでLets' EncryptのSSL証明書を取得
dockerを使う場合のコマンドはこんな感じ（Ansibleのコードそのまま持ってきた）。

`docker run --rm -e CONOHA_POLLING_INTERVAL=30 -e CONOHA_PROPAGATION_TIMEOUT=3600 -e CONOHA_TTL=3600 -e CONOHA_REGION={{ conoha_region }} -e CONOHA_TENANT_ID={{ conoha_tenant_id }} -e CONOHA_API_USERNAME={{ conoha_username }} -e CONOHA_API_PASSWORD={{ conoha_password }} -v /home/{{ ssh_user_name }}/lego:/lego goacme/lego --path /lego --email {{ email }} --dns conoha --domains *.{{ domain }} --domains {{ domain }} --accept-tos run`

ConohaのDNSはTXTレコードの反映が遅い？か何かあるらしく、デフォルトの設定ではtimeoutのエラーを吐くので、`CONOHA_PROPAGATION_TIMEOUT`、`CONOHA_PROPAGATION_TIMEOUT`、`CONOHA_TTL`は設定値を上記のようにセットしたほうが良い。

スクリプトで対応していたときもDNSの挙動にハマって上手く行かないことが多かった。なぜだろう..

# Thoughts
Super convenient. You can also use lego for certificate renewal.