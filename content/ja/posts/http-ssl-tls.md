---
title: HTTPとSSL／TLS
slug: http-ssl-tls
date: 2018-04-18T00:00:00Z
author: bmf-san
categories:
  - ネットワーク
tags:
  - HTTP
  - HTTPS
  - SSL
  - TLS
translation_key: http-ssl-tls
---


# HTTPSとは

-   HTTP（Hyper Transfer Protocol） over TLS（Transport Layer Security）
-   TCPハンドシェイクの後にTLSハンドシェイクを行う
    -   完了後、暗号通信のままHTTPリクエストとレスポンスを交換

# TLSの機能

TLSというプロトコルが提供する機能

-   機密性
    -   ネットワーク経路上でデータの中身が見れない
-   完全性
    -   通信データの改ざん防止
        -   メッセージ認証（MAC: Message Authentication Code）により、完全性を確保
-   真正性
    -   なりすまし防止
    -   認証局が発行するサーバー証明書をブラウザがアクセス先が正当かの判断に使用

# 全面的なHTTPS化の背景

-   国家レベルのハッキングがインターネットの信頼性を損なう
-   HTTPS化の推奨

    -   暗号化前提の新技術開発
    -   ブラウザにおける平文HTTPの機能廃止
    -   HTTPSサイトのSEO優位性をブラウザベンダーが
    -   Let’s Encryptなどの無料サーバー証明書発行サービスの台頭

# SSLとTLSの違い

-   SSL（Secure Socket Layer）
-   TLS（Transport Layer Security）
-   TLSはSSLの次世代規格
    -   1995年にリリースされたSSL3.0の脆弱性（POODLE）が2014年に発見される。
        -   セキュリティ対策としてはSSLを無効化するしかない
    -   1999年TLS1.0がリリース
    -   SSLとはわずかな違いしかない
    -   2006年にTLS1.1がリリース

# 参考

-   [今なぜHTTPSかなのか？インターネットの陰雷性のために、技術者が知っておきたいTLSの歴史と技術背景](https://employment.en-japan.com/engineerhub/entry/2018/02/14/110000)
-   [SSLとTLSの違いとは](https://ssl.sakura.ad.jp/column/ssl_tls/)

