---
title: "sqlcとは？Goで型安全なSQLを実現する実践入門"
slug: sqlc-introduction-golang
date: 2026-06-24
author: bmf-san
categories:
  - データベース
tags:
  - Golang
  - sqlc
  - sql
  - SQLite
  - コード生成
description: "sqlc とは何か、SQL から型安全な Go コードを生成する仕組みを、SQLite を使ったハンズオンで解説します。スキーマとクエリからコードを生成する流れ、メリットやハマりどころまで簡単に理解できます。"
translation_key: sqlc-introduction-golang
draft: false
---

# 概要
GoでSQLを扱うとき、`database/sql`を直接使うと定型コードが多くなり、ORMを使うと実行時まで誤りに気づきにくいという悩みがある。sqlcは、その中間を埋めるツールである。本記事では、SQLiteを使ったハンズオンを通して、sqlcの考え方を簡単に理解することを目的とする。

# sqlcとは
sqlcは、SQLから型安全なGoコードを自動生成するツールである。ORMとは逆の発想で、SQLそのものが主役になる点が特徴である。

スキーマ定義とクエリ定義を入力として与えると、`sqlc generate`コマンドが対応するGoコードを生成する。

```text
schema.sql（テーブル定義）┐
                         ├─→ sqlc generate ─→ 型安全なGoコード
query.sql（クエリ定義）   ┘
```

ORMのように実行時にクエリを組み立てるのではなく、ビルド前にコードが生成される。そのため、SQLの誤りを早期に発見でき、実行時のオーバーヘッドも小さい。

## メリット
- SQLをそのまま書けるため、SQLの表現力を活かせる
- 生成されるコードは型安全であり、カラムの型やNULL可否がGoの型に反映される
- `row.Scan`のような定型コードを手で書く必要がない
- 実行時リフレクションを使わないため、パフォーマンス面で有利である

## デメリット
- スキーマやクエリを変更するたびに、コードの再生成が必要である
- 条件によってWHERE句が変わるような動的クエリの組み立ては苦手である
- 対応するデータベースはPostgreSQL・MySQL・SQLiteに限られる

# ハンズオン：authorsテーブルを操作する
ここからは、実際に手を動かしてsqlcの流れを確認する。題材はシンプルな`authors`テーブルであり、データベースにはセットアップが容易なSQLiteを使う。

最終的なディレクトリ構成は次のとおりである。

```text
sqlc-handson/
├── schema.sql      # 入力：テーブル定義
├── query.sql       # 入力：クエリ定義
├── sqlc.yaml       # 設定
├── db/             # sqlc が生成（手で触らない）
│   ├── models.go
│   ├── db.go
│   └── query.sql.go
└── main.go         # アプリ
```

## 準備
sqlcはGoツールとしてインストールできる。

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

SQLiteドライバには、cgoが不要な`modernc.org/sqlite`を使う。

## スキーマ定義
まず、テーブルの設計図となる`schema.sql`を用意する。

```sql
CREATE TABLE authors (
    id   INTEGER PRIMARY KEY,
    name TEXT    NOT NULL,
    bio  TEXT
);
```

ここで重要なのは、`NOT NULL`の有無が生成されるGoの型を左右する点である。`name`は`NOT NULL`なので`string`になり、`bio`はNULLを許容するため`sql.NullString`になる。

## クエリ定義
次に、実行したいクエリを`query.sql`に書く。各クエリの直前のコメントが、sqlcへの指示になる。

```sql
-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = ? LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (name, bio)
VALUES (?, ?)
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = ?;
```

書式は`-- name: メソッド名 返し方`である。返し方には次の種類がある。

- `:one` … 1行を返す。戻り値は`(Author, error)`になる
- `:many` … 複数行を返す。戻り値は`([]Author, error)`になる
- `:exec` … 結果を返さない。戻り値は`error`のみになる

## 設定
生成方法を`sqlc.yaml`で指定する。

```yaml
version: "2"
sql:
  - engine: "sqlite"
    queries: "query.sql"
    schema: "schema.sql"
    gen:
      go:
        package: "db"
        out: "db"
        emit_json_tags: true
```

`engine`を変えるだけで、PostgreSQLやMySQL向けのコードにも切り替えられる。

## コード生成
入力が揃ったら、生成コマンドを実行する。

```bash
sqlc generate
```

`db/`ディレクトリに、次の3つのファイルが生成される。

- `models.go` … テーブルに対応する`Author`構造体
- `db.go` … `Queries`型などの実行基盤
- `query.sql.go` … 各クエリに対応するメソッド

たとえば`CreateAuthor`は、次のようなメソッドに展開される。

```go
func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, arg.Name, arg.Bio)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}
```

手で書くと退屈な`row.Scan`が、自動で生成されている点に注目してほしい。

## 生成コードを使う
生成された`Queries`を使い、データを操作する`main.go`を書く。

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"example.com/sqlc-handson/db"

	_ "modernc.org/sqlite"
)

func main() {
	ctx := context.Background()

	sqlDB, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	sqlDB.SetMaxOpenConns(1)

	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := sqlDB.ExecContext(ctx, string(schema)); err != nil {
		log.Fatal(err)
	}

	queries := db.New(sqlDB)

	author, err := queries.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language", Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created: #%d %s\n", author.ID, author.Name)
}
```

`db.CreateAuthorParams`や`sql.NullString`といった型は、すべて生成されたコードが提供する。NULLを表現したい場合は、`sql.NullString{Valid: false}`を渡せばよい。

# SQLを変えると即座に追従する
sqlcの真価は、SQLを変えるとGoコードが追従する点にある。たとえば、著者の総数を数えるクエリを`query.sql`に追加する。

```sql
-- name: CountAuthors :one
SELECT COUNT(*) AS count FROM authors;
```

再び`sqlc generate`を実行すると、次のメソッドが生成される。

```go
func (q *Queries) CountAuthors(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countAuthors)
	var count int64
	err := row.Scan(&count)
	return count, err
}
```

`SELECT COUNT(*)`は単一の数値を返すため、sqlcは構造体ではなく`int64`を返すコードを生成する。`main.go`を書き換えずに、`queries.CountAuthors(ctx)`をすぐ呼び出せる。

# ハマりどころ
ハンズオンの中で遭遇しやすい点をまとめる。

- クエリファイルには、各クエリ直前の`-- name:`コメント以外の説明コメントを書かない。`-- name:`やSQLのキーワードを含む自由なコメントは、パーサーを混乱させる原因になる
- `NOT NULL`の有無がGoの型に直結する。NULLを許容するカラムは`sql.NullString`などになり、`.Valid`で値の有無を判定する
- `db/`配下は生成物である。手で編集せず、SQLを変えて再生成する

# まとめ
sqlcは、「SQLを書く」「`sqlc generate`を実行する」「型安全なGoコードを使う」というシンプルなサイクルで動く。SQLの表現力をそのまま活かしつつ、型安全性とパフォーマンスを両立できる点が魅力である。

ORMの手軽さと生のSQLの透明性のあいだで悩んでいるなら、有力な選択肢の1つになるツールである。

# 参考
- [sqlc 公式ドキュメント](https://docs.sqlc.dev/)
- [sqlc GitHub リポジトリ](https://github.com/sqlc-dev/sqlc)
