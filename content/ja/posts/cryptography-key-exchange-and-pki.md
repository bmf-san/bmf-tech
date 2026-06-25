---
title: "鍵交換とPKI：Diffie-Hellmanと証明書・認証局"
slug: cryptography-key-exchange-and-pki
date: 2026-06-25
author: bmf-san
categories:
  - アプリケーション
tags:
  - 公開鍵暗号
  - PKI
  - 証明書
  - TLS
  - セキュリティ
description: "Diffie-Hellman鍵交換による共通鍵の安全な共有と、前方秘匿性をもたらすECDHE、そして証明書・認証局・信頼チェーンからなるPKIの仕組みを、RFCなどの一次ソース付きで解説する。暗号シリーズ全3回の第2回。"
translation_key: cryptography-key-exchange-and-pki
draft: false
---

# はじめに
本記事は暗号シリーズの第2回である。

第1回では、共通鍵暗号・公開鍵暗号・一方向性関数・ハッシュ・デジタル署名という基礎部品を扱った。

第1回（[暗号の基礎](https://bmf-tech.com/ja/posts/cryptography-fundamentals)）を読んでいない場合は、先に目を通すと理解しやすい。

今回は、公開鍵がもつ3つの用途のうち「鍵交換」と、公開鍵の正しさを保証する「PKI」を解説する。

- 鍵交換: 共通鍵を安全に共有する仕組み（Diffie-Hellman）
- PKI: その公開鍵が本人のものかを保証する仕組み

# 鍵交換とは
第1回で触れたとおり、共通鍵暗号には「鍵配送問題」がある。

盗聴されうる通信路で、どうやって両者が同じ共通鍵を共有するかという問題である。

公開鍵暗号で共通鍵を暗号化して送る方法もあるが、より洗練された解が鍵交換である。

鍵交換では、共通鍵そのものを送らない。両者が公開してよい値を交換し、それぞれの手元で同じ共通鍵を計算する。

# Diffie-Hellman鍵交換
Diffie-Hellman（DH）は、1976年に発表された最初の公開鍵的なアイデアである。

## 仕組み
DHは離散対数問題の困難性を利用する。第1回の楕円曲線暗号と同じく、「べき乗は簡単、逆算は困難」という性質である。

```
公開パラメータ: 素数 p と生成元 g
アリス: 秘密 a を選び、A = g^a mod p を送る
ボブ:   秘密 b を選び、B = g^b mod p を送る

共有鍵:
  アリス: K = B^a mod p
  ボブ:   K = A^b mod p
  --> どちらも K = g^(ab) mod p （同じ値）
```

盗聴者は`p`、`g`、`A`、`B`を見られるが、そこから秘密の`a`や`b`を求められない。だから共有鍵`K`も分からない。

## 中間者攻撃と認証
DH単体には弱点がある。通信相手が本物か確認していない点である。

攻撃者が中間に割り込み、アリスとボブそれぞれと別々に鍵交換すれば、通信を傍受できる。これが中間者攻撃である。

そのため、実際には鍵交換と「相手の認証」を組み合わせる。認証にはデジタル署名と、後述するPKIを使う。

## ECDHEと前方秘匿性
現在の主流は、楕円曲線版のDHを使う`ECDHE`である。

末尾の`E`はエフェメラル（ephemeral、一時的）を意味する。接続ごとに使い捨ての秘密値を生成する。

使い捨てにすると、ある時点の秘密鍵が漏れても、過去の通信は復号されない。この性質を前方秘匿性（forward secrecy）と呼ぶ。

# PKI（公開鍵基盤）
鍵交換や署名は、相手の公開鍵が「本物」であることを前提にする。

しかし、受け取った公開鍵が本当に通信相手のものか、どうやって信じればよいか。この問題を解くのがPKI（Public Key Infrastructure）である。

## 証明書＝認証局の署名
PKIの中心は証明書（certificate）である。

証明書は、「この公開鍵は確かにこの主体（ドメインなど）のものだ」という内容に、認証局（CA: Certificate Authority）がデジタル署名したものである。

```
証明書 = 署名( CAの秘密鍵, { 主体名, 主体の公開鍵, 有効期限, ... } )
```

RFC 5280はこの仕組みを「この結びつきは、信頼されたCAが各証明書に署名することで保証される」と述べている。

つまりPKIは、第1回のデジタル署名（①の用途）の応用である。

## 信頼チェーン
CA自身の公開鍵もまた、上位のCAが署名している。この連鎖を信頼チェーンと呼ぶ。

```
ルートCA --署名--> 中間CA --署名--> サーバ証明書
```

ルートCAの証明書は自己署名であり、チェーンの起点となる。

## トラストストアと信頼の起点
ではルートCAは何を根拠に信じるのか。

OSやブラウザには、信頼するルートCAの一覧があらかじめ組み込まれている。これをトラストストア（信頼の起点）と呼ぶ。

検証側は、受け取った証明書チェーンをたどり、最終的にトラストストア内のルートCAに行き着けば正当と判断する。

## 失効
有効期限内でも、秘密鍵の漏洩などで証明書を無効にしたい場合がある。

そのための仕組みが失効である。代表的にはCRL（失効リスト）とOCSP（オンライン照会）がある。

## ACMEとLet's Encrypt
かつて証明書の取得は手作業で煩雑だった。

これを自動化したのがACMEプロトコルであり、Let's Encryptが広く普及させた。

ドメインの管理権を自動で確認し、証明書を無料で発行・更新する。

# まとめ
本記事では鍵交換とPKIを扱った。

| 仕組み | 役割 | 一次ソース |
| --- | --- | --- |
| Diffie-Hellman | 共通鍵を安全に共有 | DH 1976 |
| ECDHE | 鍵交換＋前方秘匿性 | RFC 7919 |
| 証明書・CA | 公開鍵の真正性を保証 | RFC 5280 |
| ACME | 証明書の自動発行 | RFC 8555 |

鍵交換は「③鍵交換」、PKIは「①署名」の応用であり、いずれも第1回の基礎部品の組み合わせである。

次回（第3回）は、これらが実際のプロトコル（TLS・JWT・SSH）でどう組み合わさるかを見る。

# 参考
- Diffie, Hellman, "New Directions in Cryptography", IEEE Transactions on Information Theory 22(6), 1976. https://doi.org/10.1109/TIT.1976.1055638
- NIST SP 800-56A Rev.3: Recommendation for Pair-Wise Key-Establishment Schemes Using Discrete Logarithm Cryptography. https://csrc.nist.gov/pubs/sp/800/56/a/r3/final
- RFC 7919: Negotiated Finite Field Diffie-Hellman Ephemeral Parameters for TLS. https://www.rfc-editor.org/rfc/rfc7919
- RFC 5280: X.509 Public Key Infrastructure Certificate and CRL Profile. https://www.rfc-editor.org/rfc/rfc5280
- RFC 6960: X.509 PKI Online Certificate Status Protocol (OCSP). https://www.rfc-editor.org/rfc/rfc6960
- RFC 8555: Automatic Certificate Management Environment (ACME). https://www.rfc-editor.org/rfc/rfc8555

