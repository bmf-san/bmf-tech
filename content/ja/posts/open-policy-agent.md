---
title: "Open Policy Agentについて"
slug: "open-policy-agent"
date: 2025-05-13
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Open Policy Agent"
  - "権限管理"
draft: false
---

Open Policy Agentについて詳しく知りたかったので、公式ドキュメントを参照して調べてみた。

# Open Policy Agentとは
Open Policy Agent（OPA、「オーパ」と発音。）とは、ポリシー適用のためのOSSの汎用ポリシーエンジン。

宣言型言語（Rego、「レイゴ」と発音。）を用いてポリシーをコードとして記述することができる。

[Styra](https://www.styra.com/)によって開発されており、現在もStyraとOSSコミュニティが開発を継続している。

Cloud Native Computing Foundation（CNCF）のGraduatedプロジェクトとして認定されている。

# 概要
OPAはポリシー決定とポリシー適用を分離する。

任意の構造化データを入力として受け取り、ポリシーを適用して結果を出力する。

構造化データはJSON、YAML、CSV、XMLなど任意の形式である。

```text
                +---------------------+
                |      Service        |
                +---------------------+
                     ↑           ↑
      Request, Event |           | Decision (any JSON value)
                     |           |
                     v           |
                +---------------------+
                |         OPA         |
                +---------------------+
                     ↑           ↑
         Query (any  |           |
         JSON value) |           |
                     |           |
          +----------+-----------+----------+
          |                                 |
  +---------------+              +------------------+
  |   Policy       |              |      Data        |
  |   (Rego)       |              |    (JSON)        |
  +---------------+              +------------------+
```

OPAは入力としてのクエリとポリシーとデータを照合することで、ポリシー決定を行う。

OPAとポリシー（Rego）は特定のドメインに依存しないため、汎用的なポリシーエンジンとして様々な条件をポリシーに記述できる。

OPAはサイドカー、ホストレベルのデーモン、ライブラリとして利用できる。

# ポリシーデカップリング
OPAはポリシーをポリシー決定とポリシー適用に分離することで、ポリシーの再利用性やテスト容易性、デプロイ独立性を向上させることができる。

その結果、ビジネス要件への適応性が向上したり、ポリシーの違反や競合の検出能力を高めることができる。

# ポリシーと認可
OPAではポリシーと認可を次のように定義している。

- ポリシーとは、ソフトウェアの動作制御をするためのルールのこと
- 認可とは、ユーザーまたはシステムがどのリソースに対してアクションを実行できるかを定義するポリシーのこと

ポリシーという概念の中に認可が含まれている。認可をより一般化したものがポリシーという立ち位置になる。

# OPAのドキュメントモデル

OPAはポリシー評価のために3種類のデータを扱う：

## ベースドキュメント（Base Documents）

- 外部からOPAにロードされる構造化データ（JSON形式）。
- 主に`data`と`input`で参照される。
- ロード方式により分類される：

| モデル     | Regoでの参照 | ロード方法                    |
| ------- | -------- | ------------------------ |
| 非同期プッシュ | `data`   | OPAのAPI（例: PUT /v1/data） |
| 非同期プル   | `data`   | バンドル機能                   |
| 同期プッシュ  | `input`  | リクエスト時に直接渡す（例: POST）     |
| 同期プル    | ローカル変数   | `http.send`などの組み込み関数で取得  |

### 仮想ドキュメント（Virtual Documents）

- Regoポリシーで定義されたルールの評価結果。
- 外部からロードされるわけではなく、評価時に生成される。
- `data.foo.bar` のようにベースドキュメントと同じように参照できる。

### 参照と構文の統一

- `data` 以下にベースドキュメントと仮想ドキュメントが混在可能。
- どちらも数値、文字列、マップ、リストなど同じ型を持ち、同じ構文で扱える。
- `input` はリクエスト固有の一時的なベースドキュメントと捉える。

### パフォーマンスとキャッシュ

- 非同期にロードされた`data`やポリシーはOPA内部でキャッシュされる。
- `http.send`で取得した同期データもローカル変数として扱われ、必要に応じてキャッシュ可能。

# ポリシー言語
OPAではポリシーの記述にRegoという宣言型言語を使用する。

Regoはクエリの実行方法ではなく、クエリがどのような結果を返すかに集中してロジックを書くことができる。

ロールに応じたアクセス制御（RBAC）を行うポリシーを例に挙げる。

このポリシーでは次のようなことが実現できる。
- 役割（ロール）に応じたアクセス制御
- リソース単位でアクションを許可

以下は[Rego Playground](https://play.openpolicyagent.org/)で実行することができるサンプルコードになっている。

```rego
// Policy: ポリシーの定義
package example

default allow = false

allow if {
    perms := data.role_permissions[input.user.role][input.resource]
    perms[_] == input.action
}

// INPUT: ポリシー実行時に参照される入力データ
{
  "user": {
    "id": "u001",
    "role": "editor"
  },
  "resource": "projects",
  "action": "write"
}

// DATA: ポリシー実行時に参照されるデータ
{
  "role_permissions": {
    "viewer": {
      "projects": ["read"]
    },
    "editor": {
      "projects": ["read", "write"]
    },
    "admin": {
      "projects": ["read", "write", "delete"]
    }
  }
}

// OUTPUT: ポリシー実行時の結果
{
    "allow": true
}
```

OPAではテストのためのフレームワークも用意されており、ポリシー単体のでテストを行うことができる。

```rego
package example

test_allow_editor_write {
    input := {
        "user": {
            "id": "u001",
            "role": "editor"
        },
        "resource": "projects",
        "action": "write"
    }
    data := {
        "role_permissions": {
            "viewer": {
                "projects": ["read"]
            },
            "editor": {
                "projects": ["read", "write"]
            },
            "admin": {
                "projects": ["read", "write", "delete"]
            }
        }
    }
    allow with input as input with data as data
}
```

入力値と期待値を示すだけで簡単にテストを実装することができる。

パラメータ化テストやデータ駆動型テスト、モック、ベンチマーク、カバレッジなどにも対応している。

Regoの仕様については、[Policy Language](https://www.openpolicyagent.org/docs/latest/policy-language/)を参照。

ポリシーで定義できることについては、
[Policy Reference](https://www.openpolicyagent.org/docs/latest/policy-reference/)を参照。

[OPA/Rego入門](https://zenn.dev/mizutani/books/d2f1440cfbba94/viewer/chap-rego)はregoの基礎について分かりやすく書かれている。

# パフォーマンス
OPAは内部的にはTrieを活用したルールのインデックスにより最適化が図られている。

cf. [Optimizing OPA: Rule Indexing](https://blog.openpolicyagent.org/optimizing-opa-rule-indexing-59f03f17caf3)

Trieのデータ構造を意識してポリシーを記述することで、性能劣化の要因を防ぐことができる。（たぶん）

実際の注意点としては、[Policy Performance](https://www.openpolicyagent.org/docs/latest/policy-performance/)にいくつか記載されている。

OPAをライブラリとして利用するか、OPAサーバーを用意するかは並列化の対応が異なる。

>  If you are embedding OPA as a library, it is your responsibility to dispatch concurrent queries to different Goroutines/threads. If you are running the OPA server, it will parallelize concurrent requests and use as many cores as possible.

cf. [policy-performance/#resource-utilization](https://www.openpolicyagent.org/docs/latest/policy-performance/#resource-utilization)

ライブラリとして利用する場合は、並列化の対応が必要になる。OPAサーバーを利用する場合は、並列化の対応が自動で行われる。性能要件に応じて選択する必要がある。

`--optimize`という最適化レベルを指定するオプションが用意されており、`opa bundle`の最適化に費やす時間とリソースを調整することができる。ポリシーのサイズが最適化されることにより、ポリシーのロードや評価時のメモリ使用量を削減できる。大規模なポリシーやシビアな性能調整が必要な場合に活用できる。

# OPAの構築パターン
外部データを扱う構成のパターンについてまとめる。

| アプローチ                | パフォーマンス/可用性 | 更新頻度※1               | サイズ制限※2           | 推奨されるデータ       | セキュリティ                         | 備考                                                                                                 |
| ------------------------- | --------------------- | ---------------------- | -------------------- | ---------------------- | ------------------------------------ | ---------------------------------------------------------------------------------------------------- |
| **JWTトークン**           | 高い                  | ユーザーの再ログイン時 | 制限あり             | ユーザー属性           | トークン署名の検証、TTLの確認        | ユーザー認証時のみ更新。属性が頻繁に変わる場合には不向き。                                           |
| **`input`オーバーロード** | 高い                  | 更新頻度に依存         | サイズの制限は少ない | ローカルで動的なデータ | サービスとOPA間の接続セキュリティ    | 動的で頻繁に更新されるデータに適している。外部データの取り込みは開発者側の責任。                     |
| **バンドルAPI**           | 高い                  | 低い                   | サイズ制限あり       | 静的な中規模データ     | APIのアクセス制御                    | 外部データをバンドルとして一括で同期。データが静的であれば最適。                                     |
| **プッシュデータ**        | 高い                  | 高い                   | サイズ制限あり       | 動的で中規模データ     | APIのアクセス制御                    | 外部データの更新頻度を細かく調整可能。データが頻繁に更新される場合に適切。                           |
| **評価時のデータ取得**    | ネットワーク依存      | 常に最新               | サイズ制限なし       | 大規模または動的データ | 外部サービスへのアクセスセキュリティ | 外部データが非常に頻繁に更新される場合やデータ量が大きい場合に適している。ネットワークの可用性が鍵。 |

※1更新とは、外部データの変更がOPAに反映される頻度や容易さを指す。データが頻繁に変わる場合、リアルタイム性や反映の仕組みが重要になる。

※2サイズとは
サイズは、OPAが扱う外部データの量を指し、メモリ上に保持できるか、リクエスト毎に持ち回す必要があるかなどが選定に影響する。大規模データではオンデマンド取得が現実的。

[External Data](https://www.openpolicyagent.org/docs/latest/external-data/)を参照。

# 監視
Open TelemetryやPrometheusをサポートしている。

[Monitoring](https://www.openpolicyagent.org/docs/latest/monitoring/)を参照。

# 参考
- [www.openpolicyagent.org](https://www.openpolicyagent.org/docs/latest/)
