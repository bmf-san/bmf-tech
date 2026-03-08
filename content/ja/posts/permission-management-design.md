---
title: 権限管理の設計について調べてみた
slug: permission-management-design
date: 2024-05-22T00:00:00Z
author: bmf-san
categories:
  - アーキテクチャ
tags:
  - ABAC
  - ACL
  - RBAC
  - 権限管理
  - システム設計
  - リンク集
translation_key: permission-management-design
---


# 概要
権限管理の設計について事例を調べてみたのでメモしておく。

# 調査メモ
調査した情報を整理してみたが、わかっていないこともあるのでちゃんと整理しきれていない。

## 権限を構成する要素
権限は次の要素で成り立つものと考えられそう。

- 誰が（Principal）
- 何に（Resource）
- 何を（Action）していいか（ALLOW）、いけないか（DENY）


## 権限設計の手法
設計手法としては次のようなものが一般的に思われる。

- ACL(Access Control List)
  - ユーザーごとに権限を設定する
  - 権限はリストで管理される
- RBAC(Role Based Access Control)
  - ユーザーにロールを割り当て、ロールに権限を割り当てる
- ABAC(Attribute Based Access Control)
  - ユーザー、リソース、環境などの属性に基づいて権限を設定する

権限の柔軟性、実装の複雑性は、ACL < RBAC < ABACとなる。

## 権限設計の観点
権限設計の観点としては、次のようなものがありそう。

- 権限の適用範囲
  - 権限がどこまで適用されるか？
  - 機能的範囲
    - 単一の機能を対象することもあれば、特定の機能群を対象にすることもある
    - ex. ユーザー情報取得APIが利用できるかどうか
  - データ的範囲（スコープ）
    - 最もプリミティブな権限の適用範囲
    - ex. ユーザー情報の閲覧権限→名前と年齢は見れるが、住所は見れない
  - 両方考える必要もあれば、片方だけの考慮で良い場合もあるが、柔軟な設計が求められる
- 権限の制御対象
  - 権限がどのような基準で適用されるか？
    - コンテンツ
      - 特定の機能、データなどを対象に権限を適用する
    - コンテキスト
      - 特定の条件を満たすかどうかで権限を適用する
        - ex. 決済金額が100万円以上のユーザーのみアクセス可能
      - そもそも権限として管理される対象なのか？権限として管理されるのであれば状態はどのように管理されるか？
    - 時間
      - 特定の時間帯にのみ権限を適用する
        - ex. 平日のみアクセス可能
      - アクセス権限も権限の範疇
- 権限の制約
  - 権限の優先度や権限間の依存関係などの制約
  - 相反するような権限
    - 特定のデータしか閲覧できない権限を持っているユーザーに、全データを閲覧できる権限を持たせた場合、どちらが優先されるか
    - 相反するような権限の関係性をどう表現するか？
      - 権限の新規追加時に人力で評価するのか、属性やロールに応じた優先度を定義するのか、集合論的なアプローチで評価するかなど実装面の考慮事項がある
      - ユーザー体験やセキュリティ（最小権限の原則）の視点で検討する必要があるかもしれない
- 権限適用のレイヤー
  - 権限適用されるシステムのレイヤーはどこなのか
    - アプリケーション、データベース、ネットワーク、OSなど
  - どこから権限適用が必要か？
  - 機能的範囲やデータ的範囲の設定によってレイヤーは定まりそう
- 管理者権限の取り扱い
  - 管理者権限をどのように設定するか
  - 管理者権限はセキュリティリスクを持つ
    - 最小権限の原則、権限分割、監査ログ、緊急時のオペレーション整備などリスク管理の観点を持つ必要がある
- 権限管理の運用フロー
  - 適切に権限が管理されるための運用フローを考える必要がある
  - 最小権限の原則
    - 必要最低限の権限のみを付与する
  - 定期的な監査とレビュー
    - 権限設定を定期的に確認し、不要な権限を取り除くなど
  - 一元管理
     - 権限の種類や権限の適用状況などが把握しやすいインタフェースの提供
     - 一貫性のある権限適用フロー
       - 権限管理が各々のシステムで独自実装されているような形だと一貫性が乱れやすい（と思う）
     - 権限の分離
       - 柔軟な権限設定が可能であること
       - 一つの権限の制御対象が広すぎないこと

## 要求されるシステム特性
権限管理を行うシステムにおいて求められるシステム特性について考えてみた。

- 拡張性
  - 任意の柔軟性で権限を追加できる
- スケーラビリティ
  - 柔軟性の高い権限設計であるとシステムの複雑性やデータ量の増加も想定されやすい
  - ユーザーや権限など線形に増加していった場合のキャパシティを想定する
- 信頼性
  - 権限の追加や変更が既存の権限に悪影響を及ぼさないようにする
- セキュリティ
  - 最小権限の原則を守る、爆風半径を小さくする、障害における緊急オペレーションの整備など

# 所感
業種や業界などサービスの事業ドメインにも依るので雑感ではあるが、権限管理は、BtoB向けサービスだと特に強く求められる機能の一つじゃないかなと思っている。

体系立てられた情報やベストプラクティスとして定まったものがあるわけではないような雰囲気を感じた。

本一冊書けてしまうような分野ではありそうだが、関連書籍はあんまりなさそう。

気にすべき観点は見えてきた気がするが、特に難しいのは権限の柔軟性をどこまで許容するかという部分かなと感じた。将来的なビジネス要件もある程度加味した上で設計を広げておく必要があると思った。

# 参考
- [ja.wikipedia.org - ロールベースアクセス制御](https://ja.wikipedia.org/wiki/%E3%83%AD%E3%83%BC%E3%83%AB%E3%83%99%E3%83%BC%E3%82%B9%E3%82%A2%E3%82%AF%E3%82%BB%E3%82%B9%E5%88%B6%E5%BE%A1)
- [ja.wikpedia.org - アクセス制御リスト](https://ja.wikipedia.org/wiki/%E3%82%A2%E3%82%AF%E3%82%BB%E3%82%B9%E5%88%B6%E5%BE%A1%E3%83%AA%E3%82%B9%E3%83%88)
- [kenfdev.hateblo.jp - アプリケーションにおける権限設計の課題](https://kenfdev.hateblo.jp/entry/2020/01/13/115032)
- [knooto.info - システムのアクセス制御 (操作権限管理) 設計](https://knooto.info/software-design-access-control/#top)
- [waterlow2013.hatenablog.com - これだけ抑えればOK!権限管理のDB設計デザインパターン](https://waterlow2013.hatenablog.com/entry/2017/01/27/233405)
- [www.lyricrime.com - システムの権限方式について](https://www.lyricrime.com/posts/access-control/)
- [zenn.dev/she_techblog - 認可のアーキテクチャに関する考察（Authorization Academy IIを読んで）](https://zenn.dev/she_techblog/articles/6eff1f28d107be?redirected=1)
- [www.osohq.com - Authorization Academy](https://www.osohq.com/academy)
- [dzone.com - Access Control Acronyms: ACL, RBAC, ABAC, PBAC, RAdAC, and a Dash of CBAC](https://dzone.com/articles/acl-rbac-abac-pbac-radac-and-a-dash-of-cbac)
- [www.onelogin.com - RBACとABAC: 正しい決断を行う](https://www.onelogin.com/jp-ja/learn/rbac-vs-abac)
- [csrc.nist.gov - The NIST Model for Role-Based Access Control: Towards A Unified Standard](https://csrc.nist.gov/CSRC/media/Publications/conference-paper/2000/07/26/the-nist-model-for-role-based-access-control-towards-a-unified-/documents/sandhu-ferraiolo-kuhn-00.pdf)
- [www.internetacademy.jp - ６種類のアクセス制御の手法](https://www.internetacademy.jp/it/management/security/six-types-of-access-control-method.html)
- [butterflymx.com - Effective Access Control Design & Access Control System Planning](https://butterflymx.com/blog/access-control-design/)
- [satoricyber.com - A Comprehensive Guide to Role-Based Access Control Design](https://satoricyber.com/data-access-control/a-comprehensive-guide-to-role-based-access-control-design/)
- [tsapps.nist.gov - Role Engineering: Methods and Standards](https://tsapps.nist.gov/publication/get_pdf.cfm?pub_id=909664#:~:text=Even%20after%20a%20complete%20picture,known%20as%20%22role%20engineering%22.)
- [medium.muz.li - How to design access control system for Saas application](https://medium.muz.li/how-to-design-access-control-system-for-saas-application-b6455c944186)
- [uxdesign.cc - Designing permissions for a SaaS app](https://uxdesign.cc/design-permissions-for-a-saas-app-db6c1825f20e)
- [applis.io - システムの権限管理を設計するときの考え方とは](https://applis.io/posts/how-to-manage-authorization)
- [link-and-motivation.hatenablog - 権限管理の苦い思い出を新規サービスで昇華した話](https://link-and-motivation.hatenablog.com/entry/20220401-authorization)
- [www.okta.com - RBAC vs. ABAC：定義と使用方法](https://www.okta.com/jp/identity-101/role-based-access-control-vs-attribute-based-access-control/)
- [www.onelogin.com - RBACとABAC: 正しい決断を行う](https://www.onelogin.com/jp-ja/learn/rbac-vs-abac)
- [zenn.dev - このユーザーはこの機能が使える、みたいなあれをどう実現するか](https://zenn.dev/dove/articles/bc6933dbb39509)
- [zenn.dev - 権限周りの設計](https://zenn.dev/dove/articles/8bed47a7a839ad)
- [qiita.com - 業務システムにおけるロールベースアクセス制御](https://qiita.com/kawasima/items/8dd7eda743f2fdcad78e)
- [note.com - ロールベースでSaaSの権限設計を考える](https://note.com/tumsat/n/nfbf88bfcbc29)
