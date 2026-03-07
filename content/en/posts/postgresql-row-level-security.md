---
title: Row Level Security (RLS) in PostgreSQL
slug: postgresql-row-level-security
date: 2025-05-23T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - PostgreSQL
  - Access Control
translation_key: postgresql-row-level-security
---

# Overview
In PostgreSQL, in addition to `GRANT`-based access control (at the table and column level), **Row Level Security (RLS)** is provided as a mechanism to control the visibility and updatability of specific rows on a per-user basis.

# Basic Concept of RLS
Row Level Security functions by defining "policies" for a table.

Policies define the rules that determine which users can access which rows.

# How to Enable RLS
By default, RLS is disabled, so it needs to be enabled.

```sql
-- Enable RLS on the table
ALTER TABLE table_name ENABLE ROW LEVEL SECURITY;
```

When enabled, rows that are not explicitly permitted become invisible/unupdatable.
If there are no policies defined on a table in an enabled state, it defaults to deny visibility/updatability.
Superusers and users with the BYPASSRLS attribute can ignore RLS, but the table owner can also enable RLS with `ALTER TABLE ... FORCE ROW LEVEL SECURITY`.
Referential integrity constraints such as uniqueness constraints and foreign key constraints ignore RLS.
To back up while RLS is enabled, `row_security` must be turned off to prevent data loss.

# Defining Policies
## Basic Syntax
Policies are defined using `CREATE POLICY`.

```sql
CREATE POLICY policy_name ON table_name
  | [ AS { PERMISSIVE | RESTRICTIVE } ] |              |              |                          |
  | [ FOR { ALL       | SELECT          | INSERT       | UPDATE       | DELETE } ]               |
  | [ TO { role_name  | PUBLIC          | CURRENT_ROLE | CURRENT_USER | SESSION_USER } [, ...] ] |
  [ USING (using_expression) ]
  [ WITH CHECK (check_expression) ]
```

- policy_name: The name of the policy (unique per table)
- PERMISSIVE | RESTRICTIVE: Type of policy (described later)
- FOR: Command to which the policy applies (ALL is default)
- TO: Target roles (default is PUBLIC, all users)
- USING: Conditions for existing row access permissions (boolean expression)
- WITH CHECK: Conditions for new insert/update row validation (boolean expression)

The `USING` expression has the following meanings:

- A filter that narrows down the rows a user can access during SELECT, UPDATE, or DELETE
- Only rows where the expression is true are "visible" and subject to operations by the user
- Rows are not visible and cannot be operated on if false or null (no error is raised)

The `WITH CHECK` expression has the following meanings:

- Validity check for "newly added or updated rows in the table" during INSERT or UPDATE
- If true, the operation succeeds; if false or null, the operation is denied with an error
- Executed before the insertion/update (after BEFORE triggers)

## Types of Policies
There are two types of policies.

- PERMISSIVE (Allow Policy)
  - Multiple permissive policies are combined using logical OR
  - Access is granted if any one permissive policy is satisfied
  - Default is permissive policy
- RESTRICTIVE (Restrict Policy)
  - Multiple restrictive policies are combined using logical AND
  - All restrictive policies must be satisfied

If there are no permissive policies, access is denied.

## Characteristics of Policy Application by Command

| Command |       Role of USING Expression        |   Role of WITH CHECK Expression    |             Notes             |
| -------- | -------------------------- | ----------------------- | ---------------------------- |
| SELECT   | Conditions for visible rows   | Not used                  | SELECT privilege required             |
| INSERT   | Not used                     | Conditions for inserted rows        | Used only for insertion checks       |
| UPDATE   | Conditions for rows to update       | Conditions for rows after update      | SELECT may also be required in many cases |
| DELETE   | Conditions for rows to delete       | Not used                  | SELECT privilege required             |
| ALL      | Applies to all SELECT/UPDATE/DELETE | Conditions for INSERT/UPDATE | Applies to all commands            |

The `ALL` policy is applied in combination with other command policies.
The `SELECT` policy cannot have a `WITH CHECK` expression since it is not applicable for insert/update targets.

# Practical Use Cases
## Use Cases
Examples of use cases for RLS include:

### 1. Multi-Tenant SaaS
Use Case:
Multiple customers (tenants) share the same table, but data from other customers should not be visible.

Example:
In a SaaS service, the users table contains users belonging to multiple companies (company_id). Based on the logged-in user's company_id, they should only be able to read and write their own data.

### 2. User-Specific Data Visibility
Use Case:
Data accessible varies by user (e.g., personal ToDo lists or note-taking apps).

Example:
In the todos table, users should only see ToDos they created themselves.

### 3. Data Access Control by Department or Position
Use Case:
In a performance evaluation system, department heads should see information about their own department members, but not information from other departments.

Example:
In the employees table, there is a department_id, and a list of viewable department IDs is set in a session variable for each user.

### 4. Compliance with Legal and Regulatory Requirements (Audit, Confidentiality)
Use Case:
In industries like finance or healthcare, access should only be granted to users who meet specific conditions.

Example:
In a medical database, doctors should only be able to view records of their own patients.

## Hands-On
Let's do a hands-on exercise with RLS based on the use cases.

### Environment Setup

```sql
# init/init.sql
-- 1. Multi-Tenant SaaS: users table
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

-- 2. User-Specific ToDo: todos table
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

-- 3. Department Control: employees table
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

-- 4. Medical Data Protection: medical_records table
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

Start with `docker compose up -d`.

## Functionality Verification
View users as the alice user (company_id=1).

```sql
docker exec -it rls_pg psql -U alice -d rls_demo
SET app.current_company_id = 1;
SELECT * FROM users;
-- Only records with company_id=1 are visible
| id  |  name   | company_id |
+-----+---------+------------+
| 1   | Alice   | 1          |
| 3   | Charlie | 1          |
```

View todos as the bob user (current_user = 'bob').

```sql
docker exec -it rls_pg psql -U bob -d rls_demo
SELECT * FROM todos;
-- Only records with user_id = 'bob' are visible
| id  |    task    | user_id |
+-----+------------+---------+
| 3   | Bob task 1 | bob     |
| 4   | Bob task 2 | bob     |
```

View employees as the hr_manager user (multiple departments allowed).

```sql
docker exec -it rls_pg psql -U hr_manager -d rls_demo
SET app.allowed_departments = '1,3';
SELECT * FROM employees;
-- Only records with department_id = 1 or 3 are visible
| id  | name  | department_id |
+-----+-------+---------------+
| 1   | Eve   | 1             |
| 3   | Grace | 3             |
| 4   | Heidi | 1             |
```

View medical_records as the doctor user (only their patients).
```sql
docker exec -it rls_pg psql -U doctor -d rls_demo
SET app.current_doctor_id = 100;
SELECT * FROM medical_records;
-- Only records with doctor_id = 100 are visible
| id  | patient_name | doctor_id |
+-----+--------------+-----------+
| 1   | Patient A    | 100       |
| 3   | Patient C    | 100       |
```

# Summary
PostgreSQL's Row Level Security (RLS) is a feature that enables fine-grained row-level access control in databases.

- RLS functions by defining **policies** for tables, controlling whether users can access specific rows.
- RLS is enabled with `ALTER TABLE table_name ENABLE ROW LEVEL SECURITY`.
- Policies are created with the `CREATE POLICY` statement, specifying conditions with the `USING` and `WITH CHECK` clauses.
- There are two types of policies: PERMISSIVE (OR combination) and RESTRICTIVE (AND combination).
- Flexible policies can be created in combination with context functions like `current_user` and `current_setting`.
- Applicable to various use cases such as multi-tenant SaaS, user-specific data separation, department-based access control, and regulatory compliance.

# References
- [www.postgresql.jp - ddl-rowsecurity](https://www.postgresql.jp/document/17/html/ddl-rowsecurity.html)
- [www.postgresql.jp - sql-createpolicy](https://www.postgresql.jp/document/17/html/sql-createpolicy.html)