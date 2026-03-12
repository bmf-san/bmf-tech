---
title: なぜインターフェースの方が実装より変更頻度が低いのか
description: なぜインターフェースの方が実装より変更頻度が低いのかの手順と実践例を詳しく解説します。
slug: interface-change-frequency
date: 2025-10-18T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
  - インターフェース
translation_key: interface-change-frequency
---


[クリーンコードクックブック ―コードの設計と品質を改善するためのレシピ集](https://amzn.to/47uvc3g)を読んでいて、インターフェースの実装よりも変更頻度が低いという主張が気になったので、言語化してみた。

## インターフェースは「契約（contract）」であり「抽象」

インターフェース（interface）は、

> 「この機能はこう使える」という"契約"を表すものである。

一方で実装（implementation）は、

> 「実際にどうやって動かすか」という"具体的な手段"である。

両者は役割に違いがあり、変化に対する強度も異なる。

| 層                  | 役割    | 変わりやすさ  |
| ------------------ | ----- | ------- |
| 抽象（interface）      | 目的・約束 | 安定しやすい  |
| 具体（implementation） | 方法・手段 | 変わりやすい  |

## 契約は外部と共有されるため、むやみに変えられない

インターフェースを変更すると、それを使う**すべての呼び出し側のコードが壊れる**。

例：

```go
type UserRepository interface {
    Find(id int) (*User, error)
}
```

これを変えると：

```go
type UserRepository interface {
    FindByID(ctx context.Context, id int) (*User, error)
}
```

呼び出している箇所がすべて修正対象になる。

```go
// 修正前
user, err := repo.Find(123)

// 修正後
ctx := context.Background()
user, err := repo.FindByID(ctx, 123)
```

インターフェースの変更は波及範囲が広い（＝壊れやすい）ので慎重になる。

結果として、**めったに変えないように設計する**のである。

## 実装は裏側なので、自由に変えられる

実装は外部から直接呼ばれない。

内部のロジック、キャッシュ方法、アルゴリズム、ストレージなどは**利用者に影響を与えない範囲で変えられる**。

```go
type userRepository struct {
    db    *sql.DB
    cache map[int]*User // キャッシュを追加
    mu    sync.RWMutex  // 並行安全性のため
}

func (r *userRepository) Find(id int) (*User, error) {
    // インターフェースは変わらないが、内部実装は自由に変更可能

    // バージョン1: 直接DB検索
    // return r.findFromDB(id)

    // バージョン2: キャッシュ付き検索
    r.mu.RLock()
    user, exists := r.cache[id]
    r.mu.RUnlock()

    if exists {
        return user, nil
    }

    user, err := r.findFromDB(id)
    if err == nil {
        r.mu.Lock()
        r.cache[id] = user
        r.mu.Unlock()
    }
    return user, err
}

func (r *userRepository) findFromDB(id int) (*User, error) {
    // データベースアクセスロジック
    // PostgreSQL → MySQL に変更しても外部に影響なし
    var user User
    err := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)
    return &user, err
}
```

実装は**内部改善・最適化・リファクタリングの対象**になる。

つまり「**頻繁に変えても壊れにくい層**」である。

## 抽象のレベルが高いほど、変化に強くなる

抽象は「要件（なにをするか）」の表現である。

実装は「手段（どうやるか）」の表現である。

**手段は変わるが、目的は変わりにくい。**

```go
// 抽象レイヤー（安定）
type NotificationService interface {
    Send(message string, recipient string) error
}

// 実装レイヤー（変化しやすい）
type emailNotifier struct{}
func (e *emailNotifier) Send(message, recipient string) error {
    // SMTP → SendGrid → AWS SES など実装は変わる
}

type slackNotifier struct{}
func (s *slackNotifier) Send(message, recipient string) error {
    // Slack API実装
}

type smsNotifier struct{}
func (s *smsNotifier) Send(message, recipient string) error {
    // Twilio → AWS SNS など実装は変わる
}
```

- 「通知を送りたい」（抽象）は変わりにくい
- 「Email/Slack/SMSで送る」（実装）はよく変わる

よって、**抽象であるインターフェースの方が安定的**である。

## Goの設計哲学との関係

Goでは**インターフェースを小さく保ち、使う側（consumer）で定義する**のが慣習である。

これはつまり、

> 使われ方（＝契約）は安定しているが、
> 実装（＝内部ロジック）は自由に変えてよい。

ということである。

```go
// 小さなインターフェース（安定）
// 標準ライブラリのio.Readerは1メソッドのみ
type Reader interface {
    Read(p []byte) (n int, err error)
}

// 様々な実装（変化しやすい）
type fileReader struct { /* ファイル読み込み */ }
type networkReader struct { /* ネットワーク読み込み */ }
type compressedReader struct { /* 圧縮ファイル読み込み */ }
```

Goでは明示的な `implements` 宣言がないため（構造的部分型）、実装側は複数のインターフェースを意識せずに実装でき、使う側が必要な契約だけを定義できる。

この特性により、**依存方向が「安定 → 不安定」になるように設計する**のが自然である。

## まとめ

| 観点    | インターフェース    | 実装          |
| ----- | ----------- | ----------- |
| 役割    | 機能の約束（契約）   | 実際の動作（手段）   |
| 利用範囲  | 外部に公開される    | 内部のみ        |
| 変更の影響 | 大きい（壊れやすい）  | 小さい（内部完結）   |
| 結果    | 変更しにくい（＝安定） | 変更しやすい（＝頻繁） |
| 本質    | 「目的」は変わらない  | 「方法」は変わる    |

### 結論

インターフェースは「利用者との契約」であり、契約は一度決まると簡単に変えられない。

一方、実装は契約を守る範囲で自由に変えられる。

だから **インターフェースの方が変更頻度が低い**のである。
