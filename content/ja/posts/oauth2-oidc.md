---
title: "OAuth2、OIDCのキャッチアップのための資料"
slug: "oauth2-oidc"
date: 2021-10-05
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "FIDO"
  - "IAM"
  - "LDAP"
  - "OAuth"
  - "OpenIDConnect"
  - "SAML"
  - "SSO"
  - "認可"
  - "認証"
  - "リンク集"
draft: false
---

# 概要
OAuth2、OIDCのキャッチアップで読み漁った資料など。

## 仕様書
### OIDC
- [OpenID Connect Core 1.0 incorporating errata set 1](https://openid.net/specs/openid-connect-core-1_0.html)
- [OpenID Connect Dynamic Client Registration 1.0 incorporating errata set 1](https://openid.net/specs/openid-connect-registration-1_0.html)
- [OAuth 2.0 Multiple Response Type Encoding Practices](https://openid.net/specs/oauth-v2-multiple-response-types-1_0.html)
- [OAuth 2.0 Form Post Response Mode](https://openid.net/specs/oauth-v2-form-post-response-mode-1_0.html)
- [OpenID Connect RP-Initiated Logout 1.0 - draft 01](https://openid.net/specs/openid-connect-rpinitiated-1_0.html)
- [OpenID Connect Session Management 1.0 - draft 30](https://openid.net/specs/openid-connect-session-1_0.html)
- [OpenID Connect Front-Channel Logout 1.0 - draft 04](https://openid.net/specs/openid-connect-frontchannel-1_0.html)
- [OpenID Connect Back-Channel Logout 1.0 - draft 06](https://openid.net/specs/openid-connect-backchannel-1_0.html)
- [OpenID Connect Federation 1.0 - draft 17](https://openid.net/specs/openid-connect-federation-1_0.html)
- [OpenID Connect Basic Client Implementer's Guide 1.0 - draft 40](https://openid.net/specs/openid-connect-basic-1_0.html)
- [OpenID Connect Implicit Client Implementer's Guide 1.0 - draft 23](https://openid.net/specs/openid-connect-implicit-1_0.html)
- [OpenID 2.0 to OpenID Connect Migration 1.0](https://openid.net/specs/openid-connect-migration-1_0.html)

### OAuth 2.0
- [RFC 6749 The OAuth 2.0 Authorization Framework](https://datatracker.ietf.org/doc/html/rfc6749)
- [RFC 6750 The OAuth 2.0 Authorization Framework: Bearer Token Usage](https://datatracker.ietf.org/doc/html/rfc6750)
- [RFC 6819 OAuth 2.0 Threat Model and Security Considerations](https://datatracker.ietf.org/doc/html/rfc6819)

### JWx
- [RFC 7519 JSON Web Token (JWT)](https://datatracker.ietf.org/doc/html/rfc7519)
- [RFC 7515 JSON Web Signature (JWS)](https://datatracker.ietf.org/doc/html/rfc7515)
- [RFC 7516 JSON Web Encryption (JWE)](https://datatracker.ietf.org/doc/html/rfc7516)
- [Use Cases and Requirements for JSON Object Signing and Encryption (JOSE)](https://datatracker.ietf.org/doc/html/draft-ietf-jose-use-cases-03)
- [JSON Web Signature (JWS) draft-ietf-jose-json-web-signature-14](https://datatracker.ietf.org/doc/html/draft-ietf-jose-json-web-signature-14)

## 公認団体
- [openid.net](https://openid.net/connect/)

## 本
- [雰囲気で使わずきちんと理解する！整理してOAuth2.0を使うためのチュートリアルガイド](https://amzn.to/4dlYfXE)
- [【電子版】OAuth、OAuth認証、OpenID Connectの違いを整理して理解できる本](https://authya.booth.pm/items/1550861)
- [【電子版】OAuth・OIDCへの攻撃と対策を整理して理解できる本（リダイレクトへの攻撃編
](https://authya.booth.pm/items/1877818)
- [OAuth徹底入門 セキュアな認可システムを適用するための原則と実践](https://amzn.to/3Wj8F4o)

## Web
ブログ記事については投稿日時が古く、更新されていないものもあるので気をつけたい。

- [認証と認可の超サマリ　OAuth とか OpenID Connect とか SAML とかをまとめてざっと把握する本](https://zenn.dev/suzuki_hoge/books/2021-05-authentication-and-authorization-0259d3f/viewer/2-auth)
- [【レポート】Backend Engineer’s meetup \~マイクロサービスにおける認証認可基盤\~](https://dev.classmethod.jp/articles/merpay-microservice-auth/)
- [認証 【authentication】 certification](https://e-words.jp/w/%E8%AA%8D%E8%A8%BC.html)
- [よりよくわかる認証と認可](https://dev.classmethod.jp/articles/authentication-and-authorization-again/)
- [よくわかる認証と認可](https://dev.classmethod.jp/articles/authentication-and-authorization/)
- [OAuth 2.0 + OpenID Connect のフルスクラッチ実装者が知見を語る](https://qiita.com/TakahikoKawasaki/items/f2a0d25a4f05790b3baa)
- [Authlete を使って超高速で OAuth 2.0 & Web API サーバーを立てる](https://qiita.com/hidebike712/items/ede39abf0c860b3b96e5)
- [OAuth & OpenID Connect 関連仕様まとめ](https://qiita.com/TakahikoKawasaki/items/185d34814eb9f7ac7ef3)
- [OAuth & OpenID Connect の不適切実装まとめ](https://qiita.com/TakahikoKawasaki/items/efbbd2c5875577c911a3)
- [Authlete の OAuth 2.0 / OIDC 実装ナレッジ 完全に理解した](https://ritou.hatenablog.com/entry/2018/10/02/013902)
- [OAuth 2.0 / OIDC 実装の新アーキテクチャー](https://qiita.com/TakahikoKawasaki/items/b2a4fc39e0c1a1949aab)
- [ID連携の歴史とOpenID-Connect概要](https://speakerdeck.com/auth0japan/idlian-xi-falseli-shi-toopenid-connectgai-yao)
- [OpenID Connect 入門 〜コンシューマーにおけるID連携のトレンド〜](https://www.slideshare.net/kura_lab/openid-connect-id)
- [OpenID Connect 全フロー解説](https://qiita.com/TakahikoKawasaki/items/4ee9b55db9f7ef352b47)
- [一番分かりやすい OpenID Connect の説明](https://qiita.com/TakahikoKawasaki/items/498ca08bbfcc341691fe)
- [多分わかりやすいOpenID Connect](https://tech-lab.sios.jp/archives/8651)
- [「単なるOAUTH 2.0を認証に使うと、車が通れるほどのどでかいセキュリティー・ホールができる」について](https://tech-lab.sios.jp/archives/13002)
- [単なる OAuth 2.0 を認証に使うと、車が通れるほどのどでかいセキュリティー・ホールができる](https://www.sakimura.org/2012/02/1487/)
- [OAuth 2.0 クライアント認証](https://qiita.com/TakahikoKawasaki/items/63ed4a9d8d6e5109e401)
- [一番分かりやすい OAuth の説明](https://qiita.com/TakahikoKawasaki/items/e37caf50776e00e733be)
- [アプリ開発で知っておきたい認証技術 - OAuth 1.0 + OAuth 2.0 + OpenID Connect -](https://www.slideshare.net/ngzm/oauth-10-oauth-20-openid-connect)


# Working Group
- [openid.net/wg/](https://openid.net/wg/)

## 動画
- [「OAuth & OIDC 勉強会」録画公開と資料ダウンロード](https://www.authlete.com/ja/news/20200703_oauth_oidc_study/)

## 勉強会
- [conpass.com - Authlete, Inc.](https://authlete.connpass.com/)
- [conpass.com - Identity Dance School](https://idance.connpass.com/)

## Twitter
フォローさせてもらっているアカウント。

- [@authyasan](https://twitter.com/authyasan)
- [@authlete](https://twitter.com/authlete)
- [@authlete_jp](https://twitter.com/authlete_jp)
- [@darutk](https://twitter.com/darutk)

