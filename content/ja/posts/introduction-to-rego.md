---
title: Rego入門
description: Rego入門について調査・整理したメモ。基本概念と重要ポイントを解説します。
slug: introduction-to-rego
date: 2025-07-31T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Open Policy Agent
  - Rego
translation_key: introduction-to-rego
---


## 目次

- [第1章: Regoとは何か](#第1章-regoとは何か)
- [第2章: 基本構文とデータ構造](#第2章-基本構文とデータ構造)
- [第3章: ルールの種類と書き方](#第3章-ルールの種類と書き方)
- [第4章: 制御構造と演算子](#第4章-制御構造と演算子)
- [第5章: 内包表記とデータ操作](#第5章-内包表記とデータ操作)
- [第6章: 組み込み関数](#第6章-組み込み関数)
- [第7章: テストとデバッグ](#第7章-テストとデバッグ)

## 第1章: Regoとは何か

### 1.1 Regoの思想・特徴

Regoは**宣言型**のポリシー記述言語である。「どのように処理するか」ではなく、「何を満たすべきか」を記述する。

#### 命令型 vs 宣言型の比較

| 命令型（Go, JavaScript等）        | 宣言型（Rego）                   |
| --------------------------------- | -------------------------------- |
| for文でループして条件判定         | `some i; input.groups[i] == "admin"` |
| if/else による分岐処理            | 複数のルールによる条件記述       |
| 手続き的な実行フロー              | 条件の成立/不成立の判定          |

#### 入力（input）とデータ（data）の分離

Regoは2種類のデータを扱う：

| 種類    | 説明                           | 例                     |
| ------- | ------------------------------ | ---------------------- |
| `input` | 評価時に与えられる動的データ   | リクエスト内容         |
| `data`  | 参照用の静的・共有データ       | ユーザー情報、ロール定義 |

この分離により、ポリシーが汎用的で再利用可能になる。

### 1.2 ルールベース評価モデル

Regoの基本単位は**ルール**である：

- 複数のルールが評価され、条件を満たす場合は`true`
- 明示的な`default`がない場合、マッチしない場合は`undefined`（結果として評価されない）
- 意思決定の記述：「許可するか？」「どのフィールドが見えるか？」

## 第2章: 基本構文とデータ構造

### 2.1 基本構文要素

#### Package文
```rego
package authz  # 名前空間を定義
```

#### Import文
```rego
import data.roles      # データの参照
import input.user as u # エイリアス使用
```

#### コメント
```rego
# 一行コメント
allow {  # インラインコメント
    input.user.role == "admin"
}
```

### 2.2 データ型

| データ型    | 例                            | 説明           |
| ----------- | ----------------------------- | -------------- |
| 文字列      | `"admin"`, `"user"`           | ダブルクォート |
| 数値        | `42`, `3.14`                  | 整数・浮動小数点 |
| ブール値    | `true`, `false`               | 真偽値         |
| 配列        | `[1, 2, 3]`                   | 順序あり       |
| オブジェクト | `{"name": "Alice", "age": 30}` | キー・値のペア |
| セット      | `{"admin", "user"}`           | 重複なし       |

### 2.3 input と data

#### input - 動的データ
**[Playground で試す](https://play.openpolicyagent.org/)**

```rego
package input_example

import rego.v1

# リクエスト情報など
allow if {
    input.method == "GET"
    input.user.role == "admin"
}

# テスト用のサンプル実行
# Input: {"method": "GET", "user": {"role": "admin"}}
# 結果: allow = true

# Input: {"method": "POST", "user": {"role": "admin"}}
# 結果: allow = false
```

#### data - 静的データ
**[Playground で試す](https://play.openpolicyagent.org/)**

```rego
package data_example

import rego.v1

# 設定情報やマスターデータ
allow if {
    data.roles[input.user.role].permissions[_] == "read"
}

# テスト用のサンプル実行
# Input: {"user": {"role": "admin"}}
# 結果: allow = true

# Input: {"user": {"role": "guest"}}
# 結果: allow = false
```

**Dataセクションに貼り付けるJSON**:
```json
{
    "roles": {
        "admin": {
            "permissions": ["read", "write", "delete"]
        },
        "user": {
            "permissions": ["read"]
        }
    }
}
```

## 第3章: ルールの種類と書き方

### 3.1 ルール構文パターン

Regoのルールには複数のパターンがある：

| パターン                           | 構文例                               | 結果の型    | 説明             |
| ---------------------------------- | ------------------------------------ | ----------- | ---------------- |
| `<name> := <value>`                | `pi := 3.14`                        | 単一値      | 完全ルール       |
| `<name> if <body>`                 | `valid if input.age >= 18`          | boolean     | 条件付き完全ルール |
| `<name> contains <key> if <body>`  | `users contains name if ...`         | セット      | セットルール     |
| `<name>[<key>] := <value> if <body>` | `scores[user] := score if ...`     | オブジェクト | オブジェクトルール |

### 3.2 単一値ルール（Complete Rules）

単一の値を返すルール。

**[Playground で試す](https://play.openpolicyagent.org/)**

```rego
package quota

import rego.v1

# 定数値
default request_quota := 100

# 条件付き値
request_quota := 1000 if input.user.internal

request_quota := 50 if input.user.plan == "trial"

# テスト用のサンプル実行
# Input: {"user": {"plan": "basic"}}
# 結果: request_quota = 100

# Input: {"user": {"internal": true}}
# 結果: request_quota = 1000

# Input: {"user": {"plan": "trial"}}
# 結果: request_quota = 50
```

**評価順序**: 複数のルールがマッチした場合、最後に定義されたルールの値が採用される。

### 3.3 セットルール（Partial Set Rules）

セット（集合）を生成するルール。

**[Playground で試す](https://play.openpolicyagent.org/)**

```rego
package paths

import rego.v1

# 基本形
allowed_paths contains "/public"

# 条件付きセット生成
allowed_paths contains path if {
    some team in input.user.teams
    path := sprintf("/teams/%v/*", [team])
}

# テスト用のサンプル実行
# Input: {"user": {"teams": ["engineering", "design"]}}
# 結果: allowed_paths = {"/public", "/teams/engineering/*", "/teams/design/*"}

# Input: {"user": {}}
# 結果: allowed_paths = {"/public"}
```

### 3.4 オブジェクトルール（Partial Object Rules）

オブジェクト（マップ）を生成するルール。

**[Playground で試す](https://play.openpolicyagent.org/)**

```rego
package object_rules

import rego.v1

# プレフィックスごとのパスグループ化
paths_by_prefix[prefix] := paths if {
    some path in input.paths
    parts := split(path, "/")
    count(parts) > 0
    prefix := parts[0]  # prefixを具体的な値に束縛

    paths := [p |
        some p in input.paths
        p_parts := split(p, "/")
        count(p_parts) > 0
        p_parts[0] == prefix
    ]
}

# テスト用のサンプル実行
# Input: {"paths": ["admin/users", "admin/settings", "public/docs", "public/help"]}
# 結果: paths_by_prefix = {"admin": ["admin/users", "admin/settings"], "public": ["public/docs", "public/help"]}

# Input: {"paths": ["home/dashboard"]}
# 結果: paths_by_prefix = {"home": ["home/dashboard"]}
```

## 第4章: 制御構造と演算子

### 4.1 比較演算子

| 演算子 | 説明   | 例                        |
| ------ | ------ | ------------------------- |
| `==`   | 等しい | `input.role == "admin"`   |
| `!=`   | 異なる | `input.status != "banned"` |
| `<`    | 小なり | `input.age < 65`          |
| `<=`   | 以下   | `input.score <= 100`      |
| `>`    | 大なり | `input.salary > 50000`    |
| `>=`   | 以上   | `input.age >= 18`         |

### 4.2 論理演算子

#### AND（論理積）
**[Playground で試す](https://play.openpolicyagent.org/)**

```rego
package logical_and

import rego.v1

default allow := false

allow if {
    input.user.role == "admin"  # AND
    input.user.active == true   # AND
    input.method == "GET"       # AND
}

# テスト用のサンプル実行
# Input: {"user": {"role": "admin", "active": true}, "method": "GET"}
# 結果: allow = true

# Input: {"user": {"role": "admin", "active": false}, "method": "GET"}
# 結果: allow = false
```

#### OR（論理和）
**[Playground で試す](https://play.openpolicyagent.org/)**

```rego
package logical_or

import rego.v1

default allow := false

# 複数ルールでORを表現
allow if {
    input.user.role == "admin"
}

allow if {
    input.user.role == "manager"
    input.action == "read"
}

# テスト用のサンプル実行
# Input: {"user": {"role": "admin"}}
# 結果: allow = true

# Input: {"user": {"role": "manager"}, "action": "read"}
# 結果: allow = true

# Input: {"user": {"role": "user"}, "action": "read"}
# 結果: allow = false
```

#### NOT（否定）
**[Playground で試す](https://play.openpolicyagent.org/)**

```rego
package logical_not

import rego.v1

default allow := false

allow if {
    input.user.role == "user"
    not input.user.banned  # bannedでない場合
}

# テスト用のサンプル実行
# Input: {"user": {"role": "user", "banned": false}}
# 結果: allow = true

# Input: {"user": {"role": "user", "banned": true}}
# 結果: allow = false

# Input: {"user": {"role": "user"}}
# 結果: allow = true
```

### 4.3 量化（Quantification）

#### some - 存在量化
「少なくとも1つが条件を満たす」場合にtrue。

**[Playground で試す](https://play.openpolicyagent.org/)**

```rego
package quantification_some

import rego.v1

# 基本形
default has_admin_role := false

has_admin_role if {
    some i
    input.user.roles[i] == "admin"
}

# 簡略形
default has_admin_role_short := false

has_admin_role_short if {
    input.user.roles[_] == "admin"
}

# 変数束縛と組み合わせ
default allowed_action := false

allowed_action if {
    some role in input.user.roles
    role == "admin"  # 簡略化のため
}

# テスト用のサンプル実行
# Input: {"user": {"roles": ["user", "admin"]}}
# 結果: has_admin_role = true, has_admin_role_short = true, allowed_action = true

# Input: {"user": {"roles": ["user", "guest"]}}
# 結果: has_admin_role = false, has_admin_role_short = false, allowed_action = false
```

#### every - 全称量化
「すべてが条件を満たす」場合にtrue。

**[Playground で試す](https://play.openpolicyagent.org/)**

```rego
package quantification_every

import rego.v1

default all_files_owned := false

all_files_owned if {
    # 入力データの存在確認
    input.paths
    input.user.id
    is_array(input.paths)

    every path in input.paths {
        startswith(path, sprintf("/users/%v/", [input.user.id]))
    }
}

# テスト用のサンプル実行
# Input: {"user": {"id": "u123"}, "paths": ["/users/u123/file1.txt", "/users/u123/file2.txt"]}
# 結果: all_files_owned = true

# Input: {"user": {"id": "u123"}, "paths": ["/users/u123/file1.txt", "/public/file2.txt"]}
# 結果: all_files_owned = false
```

### 4.4 in演算子
**[Playground で試す](https://play.openpolicyagent.org/)**

```rego
package in_operator

import rego.v1

# セットのmembership
valid_method if {
    input.method in {"GET", "POST", "PUT"}
}

# 配列のmembership
valid_role if {
    input.user.role in ["admin", "manager", "user"]
}

# テスト用のサンプル実行
# Input: {"method": "GET", "user": {"role": "admin"}}
# 結果: valid_method = true, valid_role = true

# Input: {"method": "DELETE", "user": {"role": "guest"}}
# 結果: valid_method = false, valid_role = false
```

## 第5章: 内包表記とデータ操作

### 5.1 内包表記の種類

内包表記は既存のデータから新しいコレクションを生成する構文である。

| 種類           | 構文                    | 特徴                | 用途                     |
| -------------- | ----------------------- | ------------------- | ------------------------ |
| 配列内包       | `[term \| body]`        | 順序あり、重複可能  | リスト処理               |
| セット内包     | `{term \| body}`        | 順序なし、重複除去  | ユニークな値の集合       |
| オブジェクト内包 | `{key: value \| body}`  | キー・値のペア      | 構造化データの変換       |

### 5.2 配列内包

**[Playground で試す](https://play.openpolicyagent.org/)**

```rego
package array_comprehension

import rego.v1

# 基本形：数値を2倍にする
doubled := [result |
    some x in input.numbers
    result := x * 2
]

# 条件付き：偶数のみを2倍
doubled_evens := [result |
    some x in input.numbers
    x % 2 == 0
    result := x * 2
]

# 複雑な変換：ユーザー名の抽出
user_names := [name |
    some user in input.users
    user.active == true
    name := user.name
]

# テスト用のサンプル実行
# Input: {"numbers": [1, 2, 3, 4, 5], "users": [{"name": "Alice", "active": true}, {"name": "Bob", "active": false}, {"name": "Carol", "active": true}]}
# 結果:
# doubled = [2, 4, 6, 8, 10]
# doubled_evens = [4, 8]
# user_names = ["Alice", "Carol"]
```

### 5.3 セット内包

**[Playground で試す](https://play.openpolicyagent.org/)**

```rego
package set_comprehension

import rego.v1

# 重複を自動除去
unique_roles := {role |
    some user in input.users
    some role in user.roles
}

# 条件付きフィルタ
admin_users := {user.name |
    user := input.users[_]
    user.role == "admin"
}

# テスト用のサンプル実行
# Input: {"users": [{"name": "Alice", "role": "admin", "roles": ["admin", "user"]}, {"name": "Bob", "role": "user", "roles": ["user"]}, {"name": "Carol", "role": "admin", "roles": ["admin"]}]}
# 結果:
# unique_roles = {"admin", "user"}
# admin_users = {"Alice", "Carol"}
```

### 5.4 オブジェクト内包

**[Playground で試す](https://play.openpolicyagent.org/)**

```rego
package object_comprehension

import rego.v1

# ユーザーIDをキーとしたマップ作成
users_by_id := {user.id: user |
    user := input.users[_]
}

# 部署ごとの人数
dept_counts := {dept: count(members) |
    some dept in {user.department | user := input.users[_]}
    members := [user |
        user := input.users[_]
        user.department == dept
    ]
    count(members) > 0
}

# テスト用のサンプル実行
# Input: {"users": [{"id": "u1", "name": "Alice", "department": "eng"}, {"id": "u2", "name": "Bob", "department": "eng"}, {"id": "u3", "name": "Carol", "department": "sales"}]}
# 結果:
# users_by_id = {"u1": {"id": "u1", "name": "Alice", "department": "eng"}, "u2": {"id": "u2", "name": "Bob", "department": "eng"}, "u3": {"id": "u3", "name": "Carol", "department": "sales"}}
# dept_counts = {"eng": 2, "sales": 1}
```

### 5.5 ネストした内包表記

**[Playground で試す](https://play.openpolicyagent.org/)**

```rego
package nested_comprehension

import rego.v1

# 部署ごとのアクティブユーザー
active_by_dept := {dept: active_users |
    some dept in {user.department | user := input.users[_]}
    active_users := [user.name |
        user := input.users[_]
        user.department == dept
        user.active == true
    ]
}

# テスト用のサンプル実行
# Input: {"users": [{"name": "Alice", "department": "eng", "active": true}, {"name": "Bob", "department": "eng", "active": false}, {"name": "Carol", "department": "sales", "active": true}]}
# 結果:
# active_by_dept = {"eng": ["Alice"], "sales": ["Carol"]}
```

---

## 第6章: 組み込み関数

Regoには豊富な組み込み関数が用意されている。実行環境（Go SDK、WebAssembly等）によって利用可能な関数が異なるため、詳細は公式リファレンスを参照すること。

- **[Built-in Functions リファレンス](https://www.openpolicyagent.org/docs/policy-reference/builtins)**
  - 全ての組み込み関数の詳細仕様
  - 実行環境別の対応状況（Wasm / SDK-dependent）
  - 使用例とパラメータ説明

## 第7章: テストとデバッグ

### 7.1 テストの基本

Regoのテストは`test_`で始まる名前の関数として記述する。

#### 基本的なテスト構文
**[Playground で試す](https://play.openpolicyagent.org/)**

```rego
package example

import rego.v1

# テスト対象のポリシー
default allow := false

allow if {
    input.user.role == "admin"
}

# === テストコード ===

# 成功するテスト
test_admin_is_allowed if {
    allow with input as {"user": {"role": "admin"}}
}

# 失敗するテスト（notを使用）
test_guest_is_denied if {
    not allow with input as {"user": {"role": "guest"}}
}

# テスト結果
# allow = false
# test_admin_is_allowed = true
# test_guest_is_denied = true
```

### 7.2 with文による値の差し替え

`with`キーワードを使用して、テスト時に`input`や`data`を差し替えできる。

**[Playground で試す](https://play.openpolicyagent.org/)**

```rego
package with_example

import rego.v1

# テスト対象のポリシー
allow if {
    input.user.role == "admin"
    data.config.strict_mode == false
}

# === テストコード ===

# inputの差し替え
test_with_input if {
    allow with input as {"user": {"role": "admin"}}
          with data.config.strict_mode as false
}

# dataの差し替え
test_with_data if {
    allow with input as {"user": {"role": "admin"}}
          with data as {"config": {"strict_mode": false}}
}

# 複数の値を同時に差し替え
test_with_multiple if {
    allow with input as {"user": {"role": "admin"}}
          with data.config.strict_mode as false
}

# テスト結果
# allow = false
# test_with_input = true
# test_with_data = true
# test_with_multiple = true
```

### 7.3 テストケースの網羅

```rego
package authz_test

import data.authz

# 正常系のテスト
test_admin_user_allowed {
    authz.allow with input as {
        "user": {"role": "admin", "active": true},
        "action": "read"
    }
}

test_user_with_permission_allowed {
    authz.allow with input as {
        "user": {"role": "user", "active": true},
        "action": "read"
    } with data as {
        "permissions": {
            "user": ["read"]
        }
    }
}

# 異常系のテスト
test_inactive_user_denied {
    not authz.allow with input as {
        "user": {"role": "admin", "active": false},
        "action": "read"
    }
}

test_insufficient_permission_denied {
    not authz.allow with input as {
        "user": {"role": "user", "active": true},
        "action": "delete"
    } with data as {
        "permissions": {
            "user": ["read"]
        }
    }
}

# エッジケースのテスト
test_empty_input_denied {
    not authz.allow with input as {}
}

test_missing_role_denied {
    not authz.allow with input as {
        "user": {"active": true},
        "action": "read"
    }
}
```

### 7.4 デバッグ手法

#### print文によるデバッグ
```rego
allow if {
    user_role := input.user.role
    print("User role:", user_role)  # デバッグ出力

    permissions := data.roles[user_role].permissions
    print("Permissions:", permissions)  # デバッグ出力

    permissions[_] == input.action
}
```

#### trace関数（より詳細な出力）
```rego
allow if {
    trace(sprintf("Evaluating access for user: %v", [input.user.name]))
    input.user.role == "admin"
}
```

### 7.5 テストの実行

```bash
# 全テストの実行
opa test .

# 詳細出力
opa test . -v

# 特定のファイルのテスト
opa test policy_test.rego

# カバレッジ情報付き
opa test . --coverage
```

## 参考資料

- [OPA公式サイト](https://www.openpolicyagent.org/)
- [Rego Language Reference](https://www.openpolicyagent.org/docs/latest/policy-language/)
- [Rego Cheat Sheet](https://docs.styra.com/opa/rego-cheat-sheet)
- [OPA/Rego入門 (Zenn)](https://zenn.dev/mizutani/books/d2f1440cfbba94)
