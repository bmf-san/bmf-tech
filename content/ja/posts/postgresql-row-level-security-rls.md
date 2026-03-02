---
title: "PostgreSQLのRow Level Security（RLS）について"
slug: "postgresql-row-level-security-rls"
date: 2025-05-23
author: bmf-san
categories:
  - "データベース"
tags:
  - "PostgreSQL"
  - "権限管理"
draft: false
---

# 概要
PostgreSQLでは、`GRANT`ベースのアクセス制御（テーブルや列レベル）に加え、**ユーザー単位で特定の行の可視性や更新可否を制御する仕組み（行レベルのアクセス制御）**として`Row Level Security`（RLS）が提供されている。

# RLSの基本概念
Row Level Securityは、テーブルに対して「ポリシー」を定義することで機能する。

ポリシーは、どのユーザーがどの行にアクセスできるかを決定するルールを定義する。

# RLSを有効にする方法
デフォルトでは、RLSは無効であるため、有効にする必要がある。

```sql
-- テーブルに対してRLSを有効にする
ALTER TABLE テーブル名 ENABLE ROW LEVEL SECURITY;
```

有効にすると明示的に許可された行以外は不可視/更新不可になる。

有効な状態でテーブルにポリシーが存在しない場合は、デフォルト拒否で不可視/更新不可になる。

スーパーユーザーやBYPASSRLS属性を持つユーザーはRLSを無視するが、`ALTER TABLE ... FORCE ROW LEVEL SECURITY`でテーブルの所有者もRLSを有効にできる。

一意性制約・外部キー制約などの参照整合性制約はRLSを無視する。

RLSが有効な状態でバックアップするには、`row_security`をoffにしてデータ欠損を防止する必要がある。

# ポリシーの定義
## 基本構文
ポリシーは`CREATE POLICY`で定義する。

```sql
CREATE POLICY policy_name ON table_name
  | [ AS { PERMISSIVE | RESTRICTIVE } ] |              |              |                          |
  | [ FOR { ALL       | SELECT          | INSERT       | UPDATE       | DELETE } ]               |
  | [ TO { role_name  | PUBLIC          | CURRENT_ROLE | CURRENT_USER | SESSION_USER } [, ...] ] |
  [ USING (using_expression) ]
  [ WITH CHECK (check_expression) ]
```

- policy_name: ポリシーの名前（テーブルごとに一意）
- PERMISSIVE | RESTRICTIVE: ポリシーの種類（後述）
- FOR: ポリシーの適用コマンド（ALLがデフォルト）
- TO: 対象ロール（デフォルトは PUBLIC 全ユーザー）
- USING: 既存行アクセス許可条件（boolean式）
- WITH CHECK: 新規挿入・更新行の検査条件（boolean式）

`USING`式は次のような意味を持つ。

- SELECT・UPDATE・DELETE時に、ユーザがアクセスできる行を絞るフィルター
- 式が true の行だけがユーザに「見える」かつ操作対象となる
- false または null の場合は行は見えず操作不可（エラーは出ない）

`WITH CHECK`式は次のような意味を持つ。

- INSERT・UPDATE時に「新しくテーブルに追加・更新される行」の妥当性検査
- true なら操作成功、false または null ならエラーで操作拒否
- 挿入・更新前に実行される（BEFOREトリガの後）

## ポリシーの種類
ポリシーには2種類ある。

- PERMISSIVE（許容ポリシー）
  - 複数の許容ポリシーは論理和（OR）で結合される
  - いずれかの許容ポリシーを満たせばアクセス可能
  - デフォルトは許容ポリシー
- RESTRICTIVE（制限ポリシー）
  - 複数の制限ポリシーは論理積（AND）で結合される
  - すべての制限ポリシーを満たさなければならない

許容ポリシーが1つもなければアクセス不可になる。

## コマンド別ポリシー適用の特徴

| コマンド |       USING式の役割        |   WITH CHECK式の役割    |             備考             |
| -------- | -------------------------- | ----------------------- | ---------------------------- |
| SELECT   | 可視化される行の選択条件   | 不使用                  | SELECT権限が必要             |
| INSERT   | 不使用                     | 挿入行の検査条件        | 挿入時の検査にのみ使用       |
| UPDATE   | 更新対象行の選択条件       | 更新後の行検査条件      | SELECTも必要になる場合が多い |
| DELETE   | 削除対象行の選択条件       | 不使用                  | SELECT権限が必要             |
| ALL      | SELECT/UPDATE/DELETEすべて | INSERT/UPDATEの検査条件 | すべてのコマンドに適用される |

`ALL`ポリシーは他のコマンドポリシーと組み合わせて適用される。

`SELECT`ポリシーは挿入・更新対象ではないため、`WITH_CHECK`式を持てない。

# 実践的な使用例
## ユースケース
RLSを使ったユースケースには例えば以下のようなものがある。

### 1. マルチテナントSaaS
ユースケース：
複数の顧客（テナント）が同じテーブルを共有しているが、他の顧客のデータは見られないようにしたい。

事例：
SaaSサービスで users テーブルに複数の会社（company_id）に属するユーザーが存在。ログインユーザーの company_id に基づいて、自社データのみを読み書きできるようにする。


### 2. ユーザー別のデータ可視性
ユースケース：
ユーザーごとにアクセス可能なデータが異なる（例：個人用ToDoリストやメモアプリ）。

事例：
todos テーブルで、ユーザー自身が作成したToDoしか見られないようにする。

### 3. 部門や役職ごとのデータアクセス制御
ユースケース：
人事評価システムなどで、部門長は自部門のメンバー情報を見られるが、他部門の情報は非表示にしたい。

事例：
employees テーブルに department_id があり、ユーザーごとに閲覧可能な部門IDリストをセッション変数にセット。

### 4. 法的・規制要件への対応（監査、機密保持）
ユースケース：
金融や医療業界などで、特定の条件を満たしたユーザーのみにアクセスを許可する。

事例：
医療データベースで、医師は自分の担当患者のカルテのみを閲覧可能。

## ハンズオン
ユースケースに合わせてRLSのハンズオンをする。

### 環境構築

```sql
# init/init.sql
-- 1. マルチテナントSaaS: usersテーブル
CREATE TABLE users (
    id serial PRIMARY KEY,
    name text,
    company_id int
);

INSERT INTO users (name, company_id) VALUES
  ('Alice', 1),
  ('Bob', 2),
  ('Charlie', 1),
  ('David', 2);

ALTER TABLE users ENABLE ROW LEVEL SECURITY;

CREATE POLICY company_isolation_policy
  ON users
  USING (company_id = current_setting('app.current_company_id')::int);

-- 2. ユーザー別ToDo: todosテーブル
CREATE TABLE todos (
    id serial PRIMARY KEY,
    task text,
    user_id text
);

INSERT INTO todos (task, user_id) VALUES
  ('Alice task 1', 'alice'),
  ('Alice task 2', 'alice'),
  ('Bob task 1', 'bob'),
  ('Bob task 2', 'bob');

ALTER TABLE todos ENABLE ROW LEVEL SECURITY;

CREATE POLICY user_only_policy
  ON todos
  USING (user_id = current_user);

-- 3. 部門制御: employeesテーブル
CREATE TABLE employees (
    id serial PRIMARY KEY,
    name text,
    department_id int
);

INSERT INTO employees (name, department_id) VALUES
  ('Eve', 1),
  ('Frank', 2),
  ('Grace', 3),
  ('Heidi', 1);

ALTER TABLE employees ENABLE ROW LEVEL SECURITY;

CREATE POLICY department_view_policy
  ON employees
  USING (
    department_id = ANY (string_to_array(current_setting('app.allowed_departments'), ',')::int[])
);

-- 4. 医療データ保護: medical_recordsテーブル
CREATE TABLE medical_records (
    id serial PRIMARY KEY,
    patient_name text,
    doctor_id int
);

INSERT INTO medical_records (patient_name, doctor_id) VALUES
  ('Patient A', 100),
  ('Patient B', 200),
  ('Patient C', 100);

ALTER TABLE medical_records ENABLE ROW LEVEL SECURITY;

CREATE POLICY doctor_access_policy
  ON medical_records
  USING (
    doctor_id = current_setting('app.current_doctor_id')::int
);

CREATE USER alice PASSWORD 'alicepass';
CREATE USER bob PASSWORD 'bobpass';
CREATE USER hr_manager PASSWORD 'hrpass';
CREATE USER doctor PASSWORD 'docpass';

GRANT USAGE ON SCHEMA public TO alice, bob, hr_manager, doctor;

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE users TO alice, bob, hr_manager;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE todos TO alice, bob, hr_manager;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE employees TO alice, bob, hr_manager;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE medical_records TO doctor;
```

```yaml
# docker-compose.yaml
version: '3.8'
services:
  db:
    image: postgres:17.5
    container_name: rls_pg
    restart: always
    environment:
      POSTGRES_USER: admin
      POSTGRES_PASSWORD: secret
      POSTGRES_DB: rls_demo
    ports:
      - "5432:5432"
    volumes:
      - ./init:/docker-entrypoint-initdb.d
```

`docker compose up -d`で起動する。

## 動作検証
aliceユーザー（company_id=1）でusersを閲覧する。

```sql
docker exec -it rls_pg psql -U alice -d rls_demo
SET app.current_company_id = 1;
SELECT * FROM users;
-- company_id=1のレコードのみ見える
| id  |  name   | company_id |
+-----+---------+------------+
| 1   | Alice   | 1          |
| 3   | Charlie | 1          |
```

bobユーザーでtodosを閲覧（current_user = 'bob'）する。

```sql
docker exec -it rls_pg psql -U bob -d rls_demo
SELECT * FROM todos;
-- user_id = 'bob' のレコードのみ見える
| id  |    task    | user_id |
+-----+------------+---------+
| 3   | Bob task 1 | bob     |
| 4   | Bob task 2 | bob     |
```

hr_managerユーザーでemployeesを閲覧（複数部門を許可）する。

```sql
docker exec -it rls_pg psql -U hr_manager -d rls_demo
SET app.allowed_departments = '1,3';
SELECT * FROM employees;
-- department_id = 1 または 3 のレコードのみ見える
| id  | name  | department_id |
+-----+-------+---------------+
| 1   | Eve   | 1             |
| 3   | Grace | 3             |
| 4   | Heidi | 1             |
```

doctorユーザーでmedical_recordsを閲覧（担当患者のみ）
```sql
docker exec -it rls_pg psql -U doctor -d rls_demo
SET app.current_doctor_id = 100;
SELECT * FROM medical_records;
-- doctor_id = 100 のレコードのみ見える
| id  | patient_name | doctor_id |
+-----+--------------+-----------+
| 1   | Patient A    | 100       |
| 3   | Patient C    | 100       |
```

# まとめ
PostgreSQLのRow Level Security（RLS）は、データベースにおけるきめ細かな行レベルのアクセス制御を実現する機能である。

- RLSはテーブルに対して**ポリシー**を定義することで機能し、ユーザーが特定の行にアクセスできるかどうかを制御
- RLSは`ALTER TABLE テーブル名 ENABLE ROW LEVEL SECURITY`によって有効化
- ポリシーは`CREATE POLICY`ステートメントで作成し、`USING`句と`WITH CHECK`句で条件を指定
- PERMISSIVE（OR結合）とRESTRICTIVE（AND結合）の2種類のポリシータイプがある
- `current_user`や`current_setting`などのコンテキスト関数と組み合わせて柔軟なポリシーを作成可能
- マルチテナントSaaS、ユーザー別データ分離、部門別アクセス制御、規制コンプライアンスなど様々なユースケースに適用可能

# 参考
- [www.postgresql.jp - ddl-rowsecurity](https://www.postgresql.jp/document/17/html/ddl-rowsecurity.html)
- [www.postgresql.jp - sql-createpolicy](https://www.postgresql.jp/document/17/html/sql-createpolicy.html)
