---
title: "What is sqlc? A Practical Introduction to Type-Safe SQL in Go"
slug: sqlc-introduction-golang
date: 2026-06-24
author: bmf-san
categories:
  - Database
tags:
  - Golang
  - sqlc
  - SQL
  - SQLite
  - Code Generation
description: "Learn what sqlc is and how it generates type-safe Go code from SQL. A hands-on introduction using SQLite that covers the schema-to-code workflow, the benefits, and common pitfalls."
translation_key: sqlc-introduction-golang
---

# Overview
When you work with SQL in Go, the standard `database/sql` package forces you to write a lot of boilerplate, while an ORM hides mistakes until runtime. sqlc fills the gap between these two approaches. This article walks through a hands-on example with SQLite so you can grasp how sqlc works.

# What is sqlc
sqlc generates type-safe Go code from SQL. It takes the opposite approach to an ORM: SQL stays central.

You give sqlc a schema definition and a set of queries, and the `sqlc generate` command produces matching Go code.

```text
schema.sql (table definitions) ┐
                               ├─→ sqlc generate ─→ type-safe Go code
query.sql (query definitions)  ┘
```

Instead of building queries at runtime like an ORM, sqlc generates code before you build. This lets you catch SQL mistakes early and keeps runtime overhead low.

## Benefits
- You write plain SQL, so you keep the full expressive power of SQL
- The generated code stays type-safe: column types and nullability map directly to Go types
- You never hand-write boilerplate such as `row.Scan`
- It avoids runtime reflection, which helps performance

## Drawbacks
- You must regenerate code whenever you change the schema or a query
- It struggles with dynamic queries, such as a WHERE clause that changes by condition
- It supports only PostgreSQL, MySQL, and SQLite

# Hands-on: Working with an authors table
Let's work through the sqlc cycle step by step. We use a simple `authors` table and SQLite, which is easy to set up.

The final directory layout looks like this.

```text
sqlc-handson/
├── schema.sql      # input: table definition
├── query.sql       # input: query definitions
├── sqlc.yaml       # configuration
├── db/             # sqlc generates this (do not edit by hand)
│   ├── models.go
│   ├── db.go
│   └── query.sql.go
└── main.go         # application
```

## Setup
You can install sqlc as a Go tool.

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

For the SQLite driver, we use `modernc.org/sqlite`, which needs no cgo.

## Defining the schema
First, write `schema.sql`, the blueprint for your table.

```sql
CREATE TABLE authors (
    id   INTEGER PRIMARY KEY,
    name TEXT    NOT NULL,
    bio  TEXT
);
```

The key point: whether a column carries `NOT NULL` decides its Go type. Because `name` carries `NOT NULL`, it becomes `string`. Because `bio` allows NULL, it becomes `sql.NullString`.

## Defining queries
Next, write your queries in `query.sql`. The comment directly above each query tells sqlc what to do.

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

The format reads `-- name: MethodName :kind`. The kind takes one of these forms.

- `:one` returns one row, so the method returns `(Author, error)`
- `:many` returns several rows, so the method returns `([]Author, error)`
- `:exec` returns no rows, so the method returns only `error`

## Configuration
Specify how to generate code in `sqlc.yaml`.

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

Change `engine`, and sqlc switches to PostgreSQL or MySQL output.

## Generating code
Once the inputs line up, run the generate command.

```bash
sqlc generate
```

sqlc writes three files into the `db/` directory.

- `models.go` holds the `Author` struct that matches your table
- `db.go` holds the execution plumbing, such as the `Queries` type
- `query.sql.go` holds a method for each query

For example, sqlc expands `CreateAuthor` into this method.

```go
func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, arg.Name, arg.Bio)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}
```

Notice that sqlc writes the tedious `row.Scan` for you.

## Using the generated code
Now write `main.go` to drive the generated `Queries`.

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

The generated code provides every type you see here, such as `db.CreateAuthorParams` and `sql.NullString`. To store a NULL, pass `sql.NullString{Valid: false}`.

# SQL changes propagate instantly
The real strength of sqlc shows when you change SQL and the Go code follows. For example, add a query that counts authors to `query.sql`.

```sql
-- name: CountAuthors :one
SELECT COUNT(*) AS count FROM authors;
```

Run `sqlc generate` again, and sqlc produces this method.

```go
func (q *Queries) CountAuthors(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countAuthors)
	var count int64
	err := row.Scan(&count)
	return count, err
}
```

Because `SELECT COUNT(*)` returns a single value, sqlc returns `int64` directly instead of a struct. You can call `queries.CountAuthors(ctx)` right away, without touching `main.go`.

# Common pitfalls
These points trip people up most often during the hands-on.

- Keep every comment out of the query file except the `-- name:` line above each query. A free comment that contains `-- name:` or SQL keywords confuses the parser
- `NOT NULL` drives the Go type. A nullable column becomes a type such as `sql.NullString`, and you check `.Valid` to see whether a value exists
- Treat everything under `db/` as generated output. Change the SQL and regenerate instead of editing by hand

# Summary
sqlc runs on a simple cycle: write SQL, run `sqlc generate`, and use the type-safe Go code. It keeps the full power of SQL while giving you type safety and good performance.

If you waver between the convenience of an ORM and the transparency of raw SQL, sqlc makes a strong third option.

# References
- [sqlc official documentation](https://docs.sqlc.dev/)
- [sqlc GitHub repository](https://github.com/sqlc-dev/sqlc)
