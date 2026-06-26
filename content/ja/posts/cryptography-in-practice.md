---
title: "暗号技術の応用：TLS・JWT・SSH"
slug: cryptography-in-practice
date: 2026-06-25
author: bmf-san
categories:
  - アプリケーション
tags:
  - TLS
  - JWT
  - SSH
  - 認証
  - セキュリティ
description: "TLS・JWT/JWS・SSHという身近なプロトコルが、共通鍵・公開鍵・署名・鍵交換・PKIをどう組み合わせているかを、公開鍵の3用途（署名・暗号化・鍵交換）の視点でRFCを引きながら解説する。暗号シリーズ全3回の最終回。"
translation_key: cryptography-in-practice
draft: false
---

# はじめに
本記事は暗号シリーズの第3回（最終回）である。

第1回で基礎部品（共通鍵・公開鍵・一方向性関数・ハッシュ・署名）を、第2回で鍵交換とPKIを扱った。

最終回では、これらが実際のプロトコルでどう組み合わさるかを見る。題材はTLS・JWT・SSHである。

鍵となる視点は、第1回で示した公開鍵の3用途（①署名・②暗号化・③鍵交換）である。

- 第1回: [暗号の基礎](https://bmf-tech.com/ja/posts/cryptography-fundamentals)
- 第2回: [鍵交換とPKI](https://bmf-tech.com/ja/posts/cryptography-key-exchange-and-pki)
- 第3回（本記事）: 暗号技術の応用：TLS・JWT・SSH
- 実践編: [公開鍵の3つの使い道](https://bmf-tech.com/ja/posts/public-key-three-uses)

# TLS（HTTPSの土台）
TLSは、HTTPSをはじめとする通信を暗号化するプロトコルである。

その仕組みは、これまでの部品の組み合わせそのものである。

## ハンドシェイクの3要素
TLS 1.3の接続確立（ハンドシェイク）は、次の要素を組み合わせる。

- ③鍵交換（ECDHE）: 共通鍵を確立する
- ①署名＋PKI: サーバが証明書を提示し、ハンドシェイクに署名する。クライアントは証明書チェーンを検証する
- 共通鍵＋ハッシュ: 確立した鍵で、以降のデータを暗号化し完全性も守る

```
1. 鍵交換（ECDHE）で共通鍵を確立
2. サーバ証明書（CAの署名）＋ハンドシェイク署名でサーバを認証
3. 以降は共通鍵で高速に暗号化通信
```

これは第1回で述べた「公開鍵で認証・鍵共有し、共通鍵で高速通信する」というハイブリッドの典型である。

## 前方秘匿性とTLS 1.3
TLS 1.2までは、公開鍵で共通鍵を暗号化して送る方式（RSA鍵輸送）もあった。

しかしこの方式は前方秘匿性をもたない。TLS 1.3では廃止され、ECDHEなどの鍵交換に一本化された。

# JWT / JWS
JWTは、APIのアクセストークンなどで広く使われるトークン形式である。

## 構造
JWTは、クレーム（発行者や有効期限などの主張）をJSONで表し、JWSで署名したものである。

```
header.payload.signature   （各部はbase64urlエンコード）
```

署名はpayloadそのものではなく、第1回で述べたとおりハッシュを介して行う。

## algで見分ける
署名方式はヘッダーの`alg`で指定する。ここに重要な区別がある。

- `HS256`など（`HS*`）: HMAC。これは共通鍵であり、公開鍵暗号ではない
- `RS256`・`ES256`・`PS256`・`EdDSA`: 公開鍵による署名（①の用途）

`HS*`は送受信者が同じ秘密の鍵を共有する点に注意したい。鍵を渡した相手も署名を偽造できる。

## alg:noneの罠
JWSには`alg`を`none`にした「無署名」の形式もある。

検証側がこれをうっかり受理すると、誰でもトークンを偽造できる。検証時は許可する`alg`を固定するべきである。

# SSH
SSHは、安全なリモートログインなどに使うプロトコルである。

ここでも鍵交換と署名が組み合わさる。

## トランスポート層
接続時、SSHは鍵交換（DH）で共通鍵を確立する。

同時に、サーバのホスト鍵で署名してサーバを認証する。RFC 4253は「鍵交換はホスト鍵による署名と組み合わされ、サーバ認証を提供する」と述べている。

以降の通信は共通鍵で暗号化し、HMACで完全性を守る。

## TOFUという信頼モデル
SSHはPKIを使わないことが多い。

初回接続時、クライアントはサーバのホスト鍵を`known_hosts`に記録する。次回以降はその鍵と照合する。

この「最初の接続を信じて以後は固定する」方式をTOFU（Trust On First Use）と呼ぶ。PKIとは別の信頼モデルである。

## 公開鍵認証
ユーザー認証では、クライアントが自分の秘密鍵で署名する。

サーバは`authorized_keys`に登録された公開鍵で署名を検証する。これがSSH公開鍵認証であり、第1回の署名（①の用途）の代表例である。

# 3用途×応用の地図
これまでの応用を、公開鍵の3用途で整理すると次のようになる。

| | ①署名 | ③鍵交換 | 共通鍵 |
| --- | --- | --- | --- |
| TLS 1.3 | 証明書・ハンドシェイク署名 | ECDHE | 本文をAEADで暗号化 |
| JWT/JWS | RS256・ES256など（HS256はHMAC＝共通鍵） | ECDH-ESなど | HMAC・AES |
| SSH | ホスト鍵・ユーザー鍵の署名 | DH | 暗号化＋HMAC |

注意したいのは、HMAC方式（`HS256`やAWS SigV4など）やclient_secretは共通鍵であり、公開鍵暗号ではない点である。

# まとめ
3回にわたり、暗号技術を基礎から応用まで見てきた。

実際のプロトコルは、いずれも第1回・第2回で扱った部品の組み合わせである。

- TLS: ①署名（証明書）＋③鍵交換（ECDHE）＋共通鍵
- JWT/JWS: ①署名（ただし`HS*`は共通鍵）
- SSH: ①署名（ホスト鍵・ユーザー鍵）＋③鍵交換＋共通鍵

公開鍵の用途を「署名・暗号化・鍵交換」の3つで捉えると、複雑に見える応用も整理して理解できる。

# 参考
- RFC 8446: The Transport Layer Security (TLS) Protocol Version 1.3. https://www.rfc-editor.org/rfc/rfc8446
- RFC 7519: JSON Web Token (JWT). https://www.rfc-editor.org/rfc/rfc7519
- RFC 7515: JSON Web Signature (JWS). https://www.rfc-editor.org/rfc/rfc7515
- RFC 4251: The Secure Shell (SSH) Protocol Architecture. https://www.rfc-editor.org/rfc/rfc4251
- RFC 4252: The Secure Shell (SSH) Authentication Protocol. https://www.rfc-editor.org/rfc/rfc4252
- RFC 4253: The Secure Shell (SSH) Transport Layer Protocol. https://www.rfc-editor.org/rfc/rfc4253

