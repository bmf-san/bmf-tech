---
title: "認証・認可の基本"
slug: "post-246"
date: 2020-11-05
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
draft: false
---

# 概要
最近認証サービスの開発に携わっているので今一度基本的なことを再確認しておく意味でまとめる。

[Software Design 2020 11月号](https://gihyo.jp/magazine/SD/archive/2020/202011)の認証・認可の特集を参考にしている。

# IDと認証・認可の関係性
- ID
    - Identity（身元）
        - システム利用のための管理単位
            - ex. ユーザー、モノ、組織 etc...
    - Identifier（識別子）
        - データとして管理される単位
            - ex. SasS・SNS・他サービスなどで管理されるユーザーID
    - Attribute（属性）
        - Identifierを構成するそれぞれの情報
            - ex. Identityが人、Identifierがtaro_yamada、Attributeが性別、生年月日、住所、所属 etc...

# 識別・認証・認可のプロセス
- Identification・Authentication・Authorization
    - Identification（識別）
        - Identifierを一意に特定する
    - Authentication（認証）
        - Identifierの正当性（≒本人かどうか）の検証をクレデンシャルによって確認する
    - Authorization（認可）
        - ユーザーに対する権限の割り当てをどうするか決定する

# 認証・認可の組み込み
- IDとパーミッション
    - ロール
        - サービスまたはパーミッションの組み合わせのパターンを定義し、それにユーザーを割り当てる仕組み
    - RBAC（ロールベースアクセス制御）
        - 特定のリソースに対してアクセス制限をロールとして1つにまとめ、それに対してユーザーを割り当てる仕組み
    - ABAC（属性ベースアクセス制御）
        - 特定の属性に対してアクセス制限を行う仕組み
- 認証方式
    - クレデンシャルの特性
        - Something you know
            - ユーザーの記憶によるもの
        - Something you have
            - ユーザーが所持しているもの
        - Something you are
            - ユーザーの身体的な特徴に基づくもの
    - 2要素認証（多要素認証）
        - 上記の異なる特性をもつ2つ以上のクレデンシャルを組み合わせた認証
    - 2段階認証（多段階認証）
        - 認証プロセスを2回以上要求する認証

# Webサービスの認証・認可
- 自前のID管理データベース
- ソーシャルログイン
    - 認証回数軽減
    - 管理負荷軽減
- FIDO認証
    - First IDentity Online
    - 2012年にFIDOアライアンスという非営利団体が生み出した技術
    - 生体認証中心のオンライン認証
    - 公開鍵暗号方式を使用
    - 認証器（Authenticator）に認証情報を保管、認証を行う
- OAuth認可フレームワーク
    - Open Authorization
    - アプリ（OAuthクライアント）がユーザーの代理でAPIにアクセスすることを許可する仕組み
        - APIアクセスを認可する
- OAuth2.0
    - OAuth1.0は主にWebサービスが対象
    - OAuth2.0はスマホアプリも対象
    - HTTPSが必須

# WebサービスとAPIの認証・認可の違い
- 違い
    - Webサービス
        - 認証・認可後、HTTP Cookieにログイン状態を保持
    - API
        - トークンにより実行権限の有無判定
- OpenID Connectプロトコル
    - 認証結果を含むアイデンティティ情報をIDトークンに持たせることで、アイデンティティ情報を受け渡しできるようにOAuth2.0を拡張したもの
    - OAuth2.0は認可に特化しており、認証結果を含むアイデンティティ情報を受け渡す仕組みを実装しない
- OpenID Connect1.0の認証フローの種類
    - Authorization Code Flow
        - 認可コードとIDトークン（とアクセストークン）を交換する形で受け渡す
    - Implicit Flow
        - IDトークン受け渡し時に署名検証が必須
    - Hybrid Flow
        - 上記2つの融合
- OAuth2.0とOpenID Connect1.0の違い
    - OAuth2.0は認証について未定義。OpenID Connect1.0は定義されている
    - OAuth2.0ではアクセストークンの形式について定義されていないが、OpenID Connect1.0ではIDトークンの形式について定義されている
    - トークン発行までの流れは同じだが、OpenID Connect1.0はUserinfoエンドポイントというユーザー情報取得のためのAPIの実装が必要

# エンタープライズの認証・認可
- エンタープライズ対象のシステムでは、アクセス管理・アクセス制御が重要
- IAM（Identity and Access Management）
    - ユーザーやメンバーのID情報を管理し、認証、認可、アクセス権限を付与する考え方
    - コンシューマーのためのIAMではUXの向上が、エンタープライズでは、コーポレート・ガバナンスが重要課題
- ローカル認証
    - システムごとにID管理、認証、認可を行う認証
    - ユーザーやシステムが増えると管理が大変
- ディレクトリサービス
    - ネットワークに接続したリソース（システム、サーバー、アプリケーションなど）の所在・属性・設定情報などをまとめて記録・管理するサービス
    - LDAP（Lightweight Directory Access Protocol）
　　    - ディレクトリサービスにアクセスするための通信プロトコル
    - ケルベロス認証
        - サーバーとクライアントを相互認証し、身元確認するためのプロトコル。シングルサインオンを実現する技術の一つ
- SAML（Security Assertion Markup Language）
    - 異なるクラウドサービス間で認証を行うためのプロトコル

# 参考
[Software Design 2020 11月号](https://gihyo.jp/magazine/SD/archive/2020/202011)の認証・認可の特集が分かりやすかった。

