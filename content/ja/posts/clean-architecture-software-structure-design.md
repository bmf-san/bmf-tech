---
title: Clean Architecture 達人に学ぶソフトウェアの構造と設計
slug: clean-architecture-software-structure-design
date: 2018-08-01T00:00:00Z
author: bmf-san
categories:
  - アーキテクチャ
tags:
  - Clean Architecture
  - アーキテクチャ
  - 書評
translation_key: clean-architecture-software-structure-design
books:
  - asin: "4048930656"
    title: "Clean Architecture 達人に学ぶソフトウェアの構造と設計"
---


[Clean Architecture 達人に学ぶソフトウェアの構造と設計](https://amzn.to/4agMQ99)を読んだ。

Clean Architectureについて学びたいは本書と著者のブログをまず読むのが良いと思う。

- Clean Architectureについての明確な定義は語られていない
  - よくある同心円がClean Architectureであるとも主張されていないし、レイヤー構造を持つことが前提だとも書かれていない（ルールを満たしていくと必然的にレイヤー構造になっていくとは思うが）
- CleanなArchitectureだと考えられる特徴や規則について語られている
- ソフトウェアアーキテクチャの目的はシステムの構築、保守のための必要な人材を最小限にすること
- ポリモーフィズムを活用することで、依存関係逆転をすることができる
    - 独立デプロイ可能性、独立開発可能性を上げる
- 不変性の有用さ
  - デバッグがしやすい、スレッドセーフ、キャッシュ利用可能性が高い、テストがしやすいなど
- 構造化プログラミングは直接的な制御に、オブジェクト指向プログラミングは間接的な制御に、関数型プログラミングは代入に規律を課す
  - SOLID原則は変更に強く、理解しやすく、他のソフトウェアから利用しやすくソフトウェア構造を作るための原則
    - 単一責任の原則（SRP：Single Responsibility Principle）
      - モジュール変更の理由が一つになるように単一の責任を持つようにする
    - オープン・クローズドの原則（OCP：Open-Closed Principle）
      - 拡張にオープンで、変更にクローズドにする
    - リスコフの置換原則（LSP：Liskov Substitution Principle）
      - サブタイプの継承元であるスーパータイプをサブタイプに置換できるようにする
    - インターフェイス分離の原則（ISP：Interface Segregation Principle）
      - 利用しないものへの依存を回避する
    - 依存関係逆転の原則（DIP：Dependency Inversion Principle）
      - 上位レベルの方針の実装は下位レベルの詳細の実装に依存させず、下位レベルが上位レベルに依存するようにする
- コンポーネントはデプロイの単位
- コンポーネントの凝集性は開発の利便性と再利用性のトレードオフに関わる。
- コンポーネントの凝集性に関する原則
  - 再利用・リリース等価の原則（REP：Reuse-Release Equivalence Principle）
    - リリースされたものだけを再利用するようにする
  - 閉鎖性共通の原則（CCP：Common Closure Principle）
    -  同じ理由やタイミングで変更されるものは一つにまとめるようにする
  - 全再利用の原則（CRP：Common Reuse Principle）
    -  コンポーネントを使うときはコンポーネントの全てに依存するようにする
  - コンポーネントの結合は、開発の利便性と論理的な設計のトレードオフに関わる。
  - コンポーネントの結合に関する原則
    - 非循環依存関係の原則（ADP：Acyclic Dependencies Principle）
      - コンポーネントの依存関係に循環依存を含めないようにする
    - 安定依存の原則（SDP：Stable Dependencies Principle）
      - 安定度の高い（≒変更の頻度が少ない）方向に依存するようにすること
    - 安定度・抽象度等価の原則（SAP：Stable Abstractions Principle）
      - コンポーネントの抽象度は安定度と同程度になるようにする（安定度が高いものは抽象度が高く、安定度が低いものは抽象度が低くても良い）
  - ソフトウェアアーキテクチャの形状の目的は、開発・デプロイ・運用・保守を容易にすることであり、可能な限り長い期間できるだけ多くの選択肢を残すことが戦略となる
  - すぐに決めなくても良い詳細の決定についてはできるだけ遅延、またはいつでも変更できるようにしておくと良い
  - ソフトウェアは「振る舞いの価値」 と「構造の価値」持つが、後者のほうがソフトウェアをソフト（変更可能）にする価値であるゆえ価値が大きい
