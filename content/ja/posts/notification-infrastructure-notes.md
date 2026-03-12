---
title: 通知基盤構築についてのメモ書き
description: 通知基盤構築についてのメモ書きについて調査・整理したメモ。基本概念と重要ポイントを解説します。
slug: notification-infrastructure-notes
date: 2023-08-28T00:00:00Z
author: bmf-san
categories:
  - アーキテクチャ
tags:
  - 通知
translation_key: notification-infrastructure-notes
---


# 概要
通知基盤の構築に関してざっくりと考えたことや調べたことなどをまとめておく。

# 通知基盤とは
ユーザーに通知（メール・プッシュ・SMS・音声など）を行うためのシステム基盤。

クライアント（通知を依頼するシステム）からリクエストを受けて、送信先・送信内容など通知に関する処理を担うシステム。

# 通知基盤の設計・実装における観点
考えることが一杯ありそうだと思ったので、思いついた順で雑に書いた。整理できていない。

- 通知メッセージの性質
  - 性質
    - 重要なお知らせ
      - ex. 規約変更、決済関連など
    - マーケティング
      - ex. セール情報
  - 性質によって、アーキテクチャが重視する特性が変わるかもしれない
    - 例えばメールならお知らせ系はSES、マーケティング関連はAmazon Pinpointにとか、特性に応じて性能やコスト等で最適な形が変わりそう
- 通知のチャンネル
  - メール、プッシュ、SMS、音声など
- 送信パターン
  - 個別送信
  - グループ送信
  - キャンセル
    - 送信したものを取り消す仕組み
    - フェイルセーフも兼ねる
      - 1件送るはずが1000件に設定してしまった・・！みたいなヒューマンエラーとかを防ぐ仕組み、工夫があると良さそう。こういったミスはコストにダイレクトに響くと思われるため
  - 通知優先度
  - 条件指定
    - メールで通知できるならSMSは通知しない、など通知を条件分岐できる仕組み
    - この時間帯は送信しない時間指定なども
- コスト
  - コストを抑える仕組み、工夫
- スケーラビリティ
  - 通知基盤そのものはともかく、連携する内部システムにも気をつけるところがありそう
    - 例えば、ユーザーに関連する情報をリクエストする必要があるとそちらのパフォーマンスも考慮する必要が出てくる
      - 依存する外部システム（通知基盤外のサービス）は少ないほど良いはず
- メッセージング
  - キューイング
  - スケジューリング
- 送信先ユーザー情報の管理
  - 送信先情報
  - オプトイン・オプトアウト
- 認証
  - 個人情報に関わるデータにアクセスする際や認証関連APIとの連携が必要な場合とか？
- 通知メッセージのテンプレート
  - 静的・動的なデータを埋め込むためのテンプレート
  - 多言語対応は必要かどうか
- エラーハンドリングとリトライ
  - 送信失敗などエラーをどうキャッチして、再送信するか
  - メールの場合はバウンス処理など
- 監視
  - システムメトリクスだけでなく、通知関連のデータなども（ex. 受信率、クリック率等）
  - 通知のトラッキングができるように
    - トレースIDみたいなIDの発行が必要？
- 拡張性
  - 通知チャンネルの追加、テンプレートのカスタマイズ性など
- 外部サービスとの連携
  - 差し替え可能、依存し過ぎない設計にするなど
  - 分析基盤との統合・連携など
    - 通知の最適化が何かしらできると人の手で通知の運用をする部分が減って良さそう
    - 自動化できる部分は自動化したいという願い
- 運用
  - 当たり前だが運用（開発者視点、ビジネスサイド視点の両方）まで見据える
    - 今どういう運用がされているのか、これからどういう運用されるのか
- テスト
  - テストしやすさ、デバッグしやすさ
    - 複数のシステムが連携する構成になりうるので難易度が高い

# ソリューション
## SaaS
Salesforce、Braze、Airship、OneSignal、SendGridなど色々ある。

単に複数チャンネルの通知が送信できるだけではなく、顧客管理やマーケティングツールだったりなどと連携した通知ができる。

マルチチャンネルをサポートしているサービスは探すと結構色々出てくる。

## PaaS
複数チャンネルの通知をサポートしているプラットフォームとしてAWS Pinpointがある。

# 事例
国内外の事例を漁ってみた。

- [techblog.zozo.com - リアルタイムマーケティングシステムの紹介とそのリプレイス計画](https://techblog.zozo.com/entry/real-time-marketing-system)
  - AWS Pinpointのようにマルチチャンネルをサポートしつつ、MA要素も含んだ基盤を構築している事例
- [techblog.zozo.com - マーケティングオートメーションシステムを支えるリアルタイムデータ連携基盤をリプレイスした話](https://techblog.zozo.com/entry/ma-realtime-data-infrastructure-replacement)
  - 上記の続きで、その後リプレースした話
- [techblog.zozo.com - パーソナライズ配信におけるルールベースの最適化改善](https://techblog.zozo.com/entry/improving-optimization-for-personalized-marketing)
  - すごい。これが通知の真髄だ〜と思った。
  - ここまで到達するには通知がビジネスに与える影響が十分見積もられているだろうし、通知の工夫でビジネスインパクトを大きく出せるのだろうと思う
  - "真の目的である「ユーザーが本当にほしい通知だけの配信」の実現までには改善の余地が多く残されています。"とのこと
- [leandrofranchi.medium.com - How to design a Notification System](https://leandrofranchi.medium.com/how-to-design-a-notification-system-23f381cdeb00)
  - マルチチャンネルの通知基盤のシステムアーキテクチャ設計例
  - ユーザー関連のデータへのアクセスは特に考慮されていない、あくまで通知の配信のみスコープしている
  - 構成は代替こんな感じになるだろうという感想
  - 最初から通知の共通インタフェースを用意する必要は必ずしもなくて、通知チャンネルを作っていく度に落ち着いた段階でインタフェースを見出すというのも筋かなと思った
- [www.notificationapi.com - Notification Service Design - with diagrams](https://www.notificationapi.com/blog/notification-service-design-with-architectural-diagrams)
  - 上記と同じく、設計例
  - 大枠は上記と変わらない。こちらはユーザー設定周りが考慮されていて、上記はロギングが考慮されている設計になっている。
- [cloudificationzone.com - Notification System Design](https://cloudificationzone.com/2021/08/13/notification-system-design/)
  - こちらはもう少し具体化された設計例
  - Inboundが具体的にどういう通知なのかはちょっとわからなかった
- [atmarkit.itmedia.co.jp - プッシュ通知の基礎知識＆秒間1万を超えるプッシュ通知基盤のアーキテクチャと仕組みとは](https://atmarkit.itmedia.co.jp/ait/articles/1412/18/news022.html)
  - DynamoDBとNode.jsを活用したプッシュ通知基盤
- [zenn.dev - 分間10万リクエストを捌く、メール/プッシュ通知 大量配信AWSアーキテクチャ](https://zenn.dev/coconala/articles/a3a5e33cd1d984)
  - コストを抑えつつ、よりパフォーマンスの良いアーキテクチャに乗り換えた話
  - 一番時間が掛かる配信APIのリクエストを分散してリクエストできるようにすることでスケーラビリティが向上
    - 配信ワーカーの数が増えるとスケーラービリティが線形に上昇
- [www.slideshare.net - システム高速化フォーラム向け プッシュ通知基盤のアーキテクチャ](https://www.slideshare.net/recruitcojp/ss-42921628)

# AWS Pinpoint
個人的に気になっているAWS Pinpointについて個別に調べてみた。

- 複数のチャネルに対応したメッセージ（通知）のためのAWSサービス。2016年にリリースされている。
  - プッシュ通知、Eメール、SMS、音声メッセージに対応
  - 上記に加えて、[カスタムチャンネル](https://docs.aws.amazon.com/ja_jp/pinpoint/latest/userguide/channels-custom.html)という機能を使うと通知チャンネルを拡張できる
    - Facebook Messangerを追加するとか
- 従量課金
  - 思ったより高くなさそうな印象
  - メールは1万件あたり1.00USD、プッシュ通知は100万件無料、その後は100万件あたり1.00USD
  - ユーザー数数百万規模くらいで雑な試算してみるとそれなりの金額になりそうな予感
    - 当たり前だが費用対効果をちゃんと検証する必要がある。結構シビアに。
- 通知の分析もできる
  - マーケティングの施策と連携できそう
    - できることは限られそうではあるので、相性が良いかは要検討事項
- スケーラビリティ
  - 秒間に送信できる通知件数に上限がある
    - クォータの増加はリクエストできる
  - 通知の中で一番時間がネックになるであろう通知を送信する部分のスケールはAWS側にほぼ丸投げできる

## 参考
- [docs.aws.amazon.com - Amazon Pinpoint とは](https://docs.aws.amazon.com/ja_jp/pinpoint/latest/developerguide/welcome.html)
- [pages.awscloud.com - Amazon Pinpoint でユーザーを掴んで離すな
〜Amazon Digital User Engagement〜](https://pages.awscloud.com/rs/112-TZM-766/images/A3-01.pdf)
- [www.slideshare.net - Amazon Pinpoint × グロースハック活用事例集](https://www.slideshare.net/AmazonWebServicesJapan/amazon-pinpoint-x)
- [www.acrovision.jp - Amazon Pinpointとは？わかりやすく解説！低コストでプッシュ通知を実現しよう！](https://www.acrovision.jp/service/aws/?p=1421)
- [qiita.com - 初めてのAmazon Pinpoint①~概要編~](https://qiita.com/mottie/items/ebd3ed7a1a1d78ac0e76)
- [qiita.com - 初めてのAmazon Pinpoint②~実装編~](https://qiita.com/mottie/items/662f8c2938f5046471d9)
- [onetech.jp - AWS PINPOINTとは？機能や料金、メリットを徹底解説！](https://onetech.jp/blog/what-is-aws-pinpoint-15773)
- [tec.tecotec.co.jp - Amazon Pinpoint とは(プッシュ通知編)](https://tec.tecotec.co.jp/entry/2021/01/28/090000)
- [coffee-tech-blog.com - Amazon PinpointでのMA基盤構築の話](https://coffee-tech-blog.com/email-newsletter-automation-aws/)
- [www.ragate.co.jp - Amazon Pinpoint をAWS エキスパートが解説　AWSでマーケティングを簡単・手早く効率化させましょう
](https://www.ragate.co.jp/blog/articles/11830)

# 所感
誰が（運営、管理者、マーケティング担当者、開発者・・etc）、何を（メッセージ内容）、誰に、どの通知チャンネルで、いつ（いつまでに）通知したいのか、通知の総量はどれくらいなのか、ということあたりがまず整理されている必要があると思った。（そりゃそうだという感じだけど・・）

