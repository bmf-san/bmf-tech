---
title: "SLI・SLO・SLAを解説：エンジニアのための実践ガイド"
description: 'SLI・SLO・SLA の意味と違いを解説。エラーバジェットの考え方と、信頼性目標を運用に組み込むための実践的なガイドです。'
slug: sli-slo-sla-explained
date: 2022-09-10T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - SLI
  - SLA
  - SLO
translation_key: sli-slo-sla-explained
---


# SLI・SLO・SLAについて
SLI・SLO・SLAについて色々調べてみたことをまとめる。

# SLO・SLI・SLAとは何か
SLO、SLI、SLAとは、サービスレベル（Service Level）に関わる指標、目標、合意のことである。
サービスレベルとは一定の期間内で提供されたサービスを特定の方法で測定して表したものである。

- SLI（Service Level Indicator）
  - サービスレベル指標
	- サービスレベルを測定するための指標、メトリクス
	- ex. 可用性、レイテンシー、エラー率、スループット
- SLO（Serivce Level Objective）
  - サービスレベル目標
    - サービスレベルの目標とする定量または定性的な値
    - 外部依存を考慮する
      - 外部サービスとの通信部分、マネージドサービスのSLOなど外部連携している部分など
- SLA（Service Level Agreement）
  - サービスレベル合意
    - サービスの提供者と利用者の間で結ばれるサービスレベルに関する合意や保証のこと
    - SLOより緩めの目標値としたほうが良い

# SLI・SLOの設定方法
NewRelicが提唱しているベストプラクティスが取り組みやすくて良いと思う。

[newrelic.com - モダンなシステムにSLI/SLOを設定するときのベストプラクティス](https://newrelic.com/jp/blog/best-practices/best-practices-for-setting-slos-and-slis-for-modern-complex-systems)

システム境界を定義、境界ごとの機能定義、機能ごとの可用性の定義、可用性計測のためのSLI定義といった感じでSLI・SLOを策定する方法が紹介されている。

SLI・SLOの運用を始めるときは、なるべくシンプルに、緩めの値で運用を開始していくというのが推奨される。

cf. [sre.google - Chapter 4 - Service Level Objectives](https://sre.google/sre-book/service-level-objectives/#indicators-o8seIAcZ)

実際に自分が業務でSLI・SLOを策定したときは、このNewRelicのプラクティスに従ったが、機能単位のところは調整して余り細かくならないようにした。

機能の単位を最初から細かくしてしまうと運用が大変になってしまうので、運用していく中で適宜必要に応じて粒度を調整していくのが良いのではないかと思う。

# Tips
SLI・SLOに関連するキーワードについてのTips。

## 信頼性と可用性の違い
- 信頼性
  - システムが持っている性質で、故障への耐性の度合いのこと
- 可用性
  - システムが継続して稼働できる度合いのこと

## 稼働率と停止時間の一覧、可用性計算
|  稼働率  | 年間停止時間 | 月間停止時間 |
| -------- | ------------ | ------------ |
| 99.0%    | 87.6時間     | 7.6時間      |
| 99.5%    | 43.8時間     | 3.65時間     |
| 99.9%    | 8.76時間     | 43.8分       |
| 99.95%   | 4.38時間     | 21.9分       |
| 99.99%   | 52.56秒      | 4.38分       |
| 99.999%  | 5.256秒      | 26.28秒      |
| 99.9999% | 31.536秒     | 2.628秒      |

## エラーバジェットとは
エラーに対する予算で、SLOを基準として算出される許容可能な信頼性の指標のこと。
ex. SLO 99.99% → エラーバジェット 0.01%以下

# 所感
サービスレベルを測定可能にすることで、サービス利用者（ユーザーあるいはシステム）がサービスを満足に提供できるているかどうか観測可能になり、またサービス提供者にとってサービスレベルの改善が必要かどうかの指標になり得ると思った。

# 参考
- [newrelic.com - SLO、SLI、SLAとは何か？](https://newrelic.com/jp/topics/what-are-slos-slis-slas#:~:text=%E3%82%B5%E3%83%BC%E3%83%93%E3%82%B9%E3%83%AC%E3%83%99%E3%83%AB%E7%9B%AE%E6%A8%99%EF%BC%88SLO%3AService,%E6%B8%AC%E5%AE%9A%E5%80%A4%E3%81%8A%E3%82%88%E3%81%B3%E3%83%A1%E3%83%88%E3%83%AA%E3%82%AF%E3%82%B9%E3%81%A7%E3%81%99%E3%80%82)
- [newrelic.com - New Relic ハンズオン :SLI/SLO設計の基本](https://newrelic.com/sites/default/files/2022-02/NRU303_SLISLO_20220222.pdf)
- [cloud.google.com - SLO、SLI、SLA について考える : CRE が現場で学んだこと](https://cloud.google.com/blog/ja/products/gcp/availability-part-deux-cre-life-lessons)
- [cloud.google.com - SRE の基本（2021 年版）: SLI、SLA、SLO の比較](https://cloud.google.com/blog/ja/products/devops-sre/sre-fundamentals-sli-vs-slo-vs-sla)
- [cloud.google.com - SLOs, SLIs, SLAs, oh my—CRE life lessons](https://cloud.google.com/blog/products/devops-sre/availability-part-deux-cre-life-lessons)
- [cloud.google.com - 可用性とどう向き合うべきか、それが問題だ : CRE が現場で学んだこと](https://cloud.google.com/blog/ja/products/gcp/available-or-not-that-is-the-question-cre-life-lessons)
- [engineering.mericari.com - 2018/12/25 メルカリのWeb MicroservicesにおけるSLI/SLO](https://engineering.mercari.com/blog/entry/2018-12-25-150405/)
- [sre.google - sre-book](https://sre.google/sre-book/service-level-objectives/#indicators-o8seIAcZ)
- [qiita.com - SLI/SLOを策定するために考えたこと](https://qiita.com/t-okibayashi/items/9a5085803ac0b11554a0)
- [qiita.com - SREについて学ぶ - エラーバジェット編](https://qiita.com/katsulang/items/feb3070666607b7c924c#:~:text=%E3%82%A8%E3%83%A9%E3%83%BC%E3%83%90%E3%82%B8%E3%82%A7%E3%83%83%E3%83%88%E3%81%A8%E3%81%AF%E3%80%81%E3%82%B5%E3%83%BC%E3%83%93%E3%82%B9,%E6%8A%91%E3%81%88%E3%82%8B%E3%81%93%E3%81%A8%E3%81%AB%E3%81%AA%E3%82%8A%E3%81%BE%E3%81%99%E3%80%82)
- [bongineer.net - 信頼性と可用性の違い](https://bongineer.net/entry/rasis/)
- [mathwords.net - 稼働率（可用性）99％、99.9％などの停止時間はいくらか](https://mathwords.net/kadouritu)
- [wnkhs.net - 可用性の計算と想定値（代表的な数字にて）](https://wnkhs.net/availability-calculation/)

