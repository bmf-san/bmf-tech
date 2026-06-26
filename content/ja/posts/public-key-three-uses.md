---
title: "公開鍵の3つの使い道 ―署名・暗号化・鍵交換で実応用を読み解く"
slug: public-key-three-uses
date: 2026-06-25
author: bmf-san
categories:
  - アプリケーション
tags:
  - 公開鍵暗号
  - デジタル署名
  - 認証
  - セキュリティ
  - JWT
  - WebAuthn
  - OAuth
  - PKI
description: "公開鍵（鍵ペア）の用途は署名・暗号化・鍵交換の3つに集約できる。SSH公開鍵認証・WebAuthn・mTLS・private_key_jwt・コード署名・コンテナ署名などの実応用をこの3分類で整理し、HMACなど共通鍵との違いも一次ソース付きで解説する。"
translation_key: public-key-three-uses
draft: false
---

# はじめに
HTTPS、SSH、JWT、パスキーなど、現代の認証やセキュリティは公開鍵暗号に支えられている。

一見ばらばらに見えるこれらの技術も、公開鍵（鍵ペア）の用途という観点で見ると、たった3つに集約できる。

- ①署名: 秘密鍵で署名し、公開鍵で検証する
- ②暗号化: 公開鍵で暗号化し、秘密鍵で復号する
- ③鍵交換: 共通鍵を安全に共有する

世の中の応用は、ほぼこの3つの組み合わせである。本記事では、実際の技術をこの3分類で整理する。

本シリーズの構成は次のとおりである。本記事は実践編にあたる。

- 第1回: [暗号の基礎](https://bmf-tech.com/ja/posts/cryptography-fundamentals)（共通鍵/公開鍵・トラップドア関数・ハッシュ・署名）
- 第2回: [鍵交換とPKI](https://bmf-tech.com/ja/posts/cryptography-key-exchange-and-pki)
- 第3回: [暗号技術の応用](https://bmf-tech.com/ja/posts/cryptography-in-practice)（TLS・JWT・SSH）
- 実践編（本記事）: 公開鍵の3つの使い道

用途別に応用を俯瞰するのが本記事の位置づけである。

# 公開鍵の3つの使い道
まず全体像を示す。

| 用途 | 操作 | 守るもの |
| --- | --- | --- |
| ①署名 | 秘密鍵で署名し、公開鍵で検証 | 真正性・認証・改ざん検知 |
| ②暗号化 | 公開鍵で暗号化し、秘密鍵で復号 | 機密性 |
| ③鍵交換 | 公開してよい値を交換し共通鍵を導出 | 共通鍵の安全な共有 |

①と②は、鍵を使う向きがちょうど逆になっている。詳しくは基礎編で解説した。

以下、用途ごとに実応用を見ていく。

# ①署名 = 真正性・認証
最も応用が広いのは署名である。

秘密鍵を持つ本人だけが署名でき、公開鍵を持てば誰でも検証できる。この非対称性は、認証・真正性・改ざん検知に使える。

代表的な応用を、署名者（秘密鍵を持つ側）と検証者（公開鍵を持つ側）で整理する。

| パターン | 署名者（秘密鍵） | 検証者（公開鍵） | 主な用途 |
| --- | --- | --- | --- |
| SSH公開鍵認証 | クライアント | サーバー（`authorized_keys`） | リモートログイン |
| WebAuthn・パスキー | ユーザー端末 | サービス（RP） | パスワードレス認証 |
| mTLS | クライアントとサーバー双方 | 相手 | M2Mの相互認証 |
| private_key_jwt | クライアント | 認可サーバー | OAuthクライアント認証 |
| サービスアカウントJWT | ワークロード | IdP・トークン発行点 | M2Mのトークン取得 |
| SAML・OIDCのIdP署名 | IdP | SP・RP | フェデレーション |
| コード署名 | 配布元 | OS・ユーザー | 改ざん検知 |
| コンテナ署名（cosign） | 配布元 | デプロイ側 | サプライチェーン保護 |
| Gitコミット・タグ署名 | コミッタ | レビュアやGitHub | 真正性の担保 |
| 暗号通貨ウォレット | 所有者 | ネットワーク全体 | 取引への署名 |

いずれも「秘密鍵で署名し、公開鍵で検証する」という同じ構造である。

たとえばSSH公開鍵認証では、クライアントが秘密鍵で署名し、サーバーが`authorized_keys`の公開鍵で検証する。

OAuthのprivate_key_jwtも同型である。クライアントが秘密鍵で署名したJWTを、認可サーバーが登録済みの公開鍵で検証する。client_secretを送る方式より安全とされる。

# ②暗号化 = 機密性
公開鍵で暗号化し、秘密鍵で復号する用途である。①署名の鏡像にあたる。

公開鍵は誰でも使えるため、「特定の相手（秘密鍵の持ち主）だけが読める」メッセージを作れる。

代表的な応用は次のとおりである。

- PGP・GPG: メールやファイルの暗号化
- S/MIME: メールの暗号化と署名
- SOPS＋age: IaCの秘密情報（Secret）をファイルとして暗号化

ただし、通信の暗号化では②を直接は使わないことが多い。公開鍵暗号は遅いため、③鍵交換で共通鍵を作り、共通鍵暗号で本体を暗号化する。②が活きるのは、保存データや非同期メッセージなど、相手とリアルタイムに鍵交換できない場面である。

# ③鍵交換 = 共通鍵の安全な共有
共通鍵そのものを送らず、両者が公開してよい値を交換して、同じ共通鍵を手元で導出する用途である。

Diffie-Hellman（DH）やその楕円曲線版ECDHEが使われる。仕組みの詳細は[鍵交換とPKI](https://bmf-tech.com/ja/posts/cryptography-key-exchange-and-pki)で解説した。

代表的な応用は次のとおりである。

- TLS・HTTPS: サーバー認証（①署名）とECDHE（③鍵交換）の組み合わせ
- Signalなどのエンドツーエンド暗号: メッセージングのE2EE
- WireGuard・SSH: 通信路の暗号化

通信の暗号化に②でなく③を使う理由は、前方秘匿性にある。接続ごとに使い捨ての鍵を使うことで、後で秘密鍵が漏れても過去の通信は守られる。

# 信頼の土台: PKI
①②③のいずれも、「相手の公開鍵が本物か」という前提に立つ。

この前提を保証するのがPKIである。認証局（CA）が「この公開鍵は確かにこの主体のものだ」と署名した証明書を発行する。

PKI自体も①署名の応用である。詳細は[鍵交換とPKI](https://bmf-tech.com/ja/posts/cryptography-key-exchange-and-pki)で解説した。

# 注意: 公開鍵暗号ではないもの
最後に、公開鍵暗号と混同しやすい「共通鍵」の仕組みに触れておく。

次のものは公開鍵暗号ではなく、送受信者が同じ秘密の鍵を共有する共通鍵方式である。

- HMAC（`HS256`のJWT、AWS SigV4、多くのWebhook署名）
- client_secret方式のクライアント認証

これらは「署名」と呼ばれることもあるが、鍵を渡した相手も同じものを作れる点が公開鍵署名と決定的に違う。

JWTでは`alg`で見分けられる。

- `HS*`（`HS256`など）: HMAC。共通鍵
- `RS*`・`ES*`・`PS*`・`EdDSA`: 公開鍵による署名

# まとめ
公開鍵（鍵ペア）の用途は、突き詰めると3つしかない。

- ①署名: 認証・真正性・改ざん検知
- ②暗号化: 機密性
- ③鍵交換: 共通鍵の安全な共有

SSH、WebAuthn、mTLS、JWT、コード署名、暗号通貨まで、多くの応用はこの3つの組み合わせとして理解できる。新しい技術に出会ったときも、「秘密鍵で何をし、公開鍵で何を検証するのか」を問えば、構造が見えてくる。

# 参考
- W3C Web Authentication (WebAuthn) Level 2. https://www.w3.org/TR/webauthn-2/
- RFC 8705: OAuth 2.0 Mutual-TLS Client Authentication and Certificate-Bound Access Tokens. https://www.rfc-editor.org/rfc/rfc8705
- RFC 7523: JSON Web Token (JWT) Profile for OAuth 2.0 Client Authentication and Authorization Grants. https://www.rfc-editor.org/rfc/rfc7523
- OpenID Connect Core 1.0. https://openid.net/specs/openid-connect-core-1_0.html
- RFC 4880: OpenPGP Message Format. https://www.rfc-editor.org/rfc/rfc4880
- RFC 8551: S/MIME Version 4.0 Message Specification. https://www.rfc-editor.org/rfc/rfc8551
- RFC 2104: HMAC: Keyed-Hashing for Message Authentication. https://www.rfc-editor.org/rfc/rfc2104
- RFC 7518: JSON Web Algorithms (JWA). https://www.rfc-editor.org/rfc/rfc7518

