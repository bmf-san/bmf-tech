---
title: Considerations and Solutions for Pagination Impact in OPA
description: 'Master pagination with Open Policy Agent, offset pagination, cursor pagination, and access control filtering.'
slug: pagination-impact-solutions-opa
date: 2025-06-14T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Open Policy Agent
  - Offset Pagination
  - Cursor Pagination
  - Access Control
translation_key: pagination-impact-solutions-opa
---



# Basics and Background of OPA
OPA (Open Policy Agent) is an engine that evaluates policies written in the Rego language based on input and external data to make decisions such as allow/deny.

For implementation examples, the [AWS Prescriptive Guidance on Multi-Tenant API Authorization Control](https://docs.aws.amazon.com/ja_jp/prescriptive-guidance/latest/saas-multitenant-api-access-authorization/introduction.html) is useful. It introduces strategies for implementing multi-tenant authorization in SaaS using OPA, which is also helpful for understanding the background of this article.

## Basic OPA Sequence
Below is the basic sequence of access control using OPA.

```mermaid
sequenceDiagram
    participant Client as クライアント
    participant API as アプリケーション
    participant OPA as OPA
    participant Policy as Regoポリシー
    participant Data as 外部データ

    Client->>API: リクエスト送信
    API->>OPA: input を含む評価要求
    OPA->>Data: ポリシー評価用データ取得
    Data-->>OPA: リソースデータ
    OPA->>Policy: Rego ポリシーで評価
    Policy-->>OPA: 結果（allow/deny）
    OPA-->>API: 判定結果
    API-->>Client: 応答
```

## Issues with Pagination
Below is the sequence when OPA is naively applied to pagination.

```mermaid
sequenceDiagram
    participant Client
    participant API
    participant DB
    participant OPA

    Client->>API: GET /resources?page=P&size=N
    Note over API: OFFSET/LIMIT で DB に問い合わせ【OPA前】
    API->>DB: SELECT ... OFFSET (P-1)*N LIMIT N
    DB-->>API: N件取得（未フィルタ）
    Note over API: 各リソースを OPA で評価
    API->>OPA: ポリシー評価（リソース単位）
    OPA-->>API: 許可 or 拒否
    alt 許可リソース数 >= N
        API-->>Client: N件返却
    else 許可数が不足
        loop 許可リソースがN件に達すまで繰り返し
            API->>DB: 次のOFFSETで追加取得
            DB-->>API: 追加データ
            API->>OPA: ポリシー評価
            OPA-->>API: 許可 or 拒否
        end
        API-->>Client: 許可された最大N件を返却
    end
```

## Offset Pagination
```mermaid
sequenceDiagram
    participant Client as クライアント
    participant App as アプリケーション
    participant DB as データベース

    Client->>App: GET /items?page=P&size=N
    Note over App: page, size から OFFSET, LIMIT 計算
    App->>DB: SELECT * FROM items WHERE <条件> ORDER BY ... LIMIT N OFFSET (P-1)*N
    DB-->>App: N件（または残り分だけ）
    App-->>Client: 結果返却
```

## Cursor Pagination
```mermaid
sequenceDiagram
    participant Client as クライアント
    participant App as アプリケーション
    participant DB as データベース

    Client->>App: GET /items?after=<cursor>&size=N
    Note over App: カーソル以降を WHERE で指定
    App->>DB: SELECT * FROM items WHERE id > <cursor> ORDER BY id ASC LIMIT N
    DB-->>App: N件の結果
    App-->>Client: 返却 (next_cursor 付)
```

## Challenges (Possible with SQL but Difficult with Naive OPA Evaluation)

| Aspect             | What was possible with SQL filtering             | Constraints with naive OPA evaluation                              |
| ------------------ | ------------------------------------------------ | ---------------------------------------------------------------- |
| Count Awareness    | Possible to get total count beforehand with WHERE clause | Cannot predict allowed count until after evaluation               |
| Order Consistency  | Achieve stable order and slice with ORDER BY + OFFSET/LIMIT | Mixed denied resources cause unstable display order and page boundaries |
| Page Number Consistency | Possible to display explicit pages like "21st to 40th" | Must construct pages based on allowed results, leading to inconsistency |
| Query Efficiency   | Possible to minimize retrieval and processing using indexes | Need to re-fetch and re-evaluate whenever allowed count is insufficient, increasing load |

## Consideration of Solutions
### 1. Naive Implementation (Offset or Cursor Pagination)
Retrieve all target resources beforehand and pass them to OPA for evaluation. OPA returns not only allow/deny decisions for each resource but also aggregated information like allowed list and denied count, enabling pagination and total hit count display on the app side.

- Advantages:
  - Relatively simple implementation
  - Stable pagination and hit count display based on OPA evaluation results
  - Can return an accurate slice to client requests
- Disadvantages:
  - Memory consumption and latency issues arise as data count increases due to full retrieval and evaluation
  - Re-evaluation may be needed upon re-request unless the evaluated list is retained or cached

### 2. Return Conditions for SQL Generation in OPA
Partially evaluate OPA policies as SQL WHERE clause equivalent conditions and apply them to SQL queries on the app side.

- Advantages: Efficient processing through SQL filtering
- Disadvantages: Requires input data for condition generation in OPA, making Rego policies act as condition generators (weakening policy design consistency)

### 3. Implement Pagination on the Frontend
The backend returns all allowed resources, and the client handles pagination.

- Advantages: Simple implementation, unaffected by OPA application order or count
- Disadvantages: Initial load and communication volume tend to be large due to full retrieval

## Conclusion
Applying OPA naively to list retrieval results in complex issues such as indeterminate return count, increased processing load, ambiguous page boundaries, and degraded user experience due to the inconsistency between pagination and policy evaluation.

There are trade-offs in applying OPA to pagination. It is necessary to clarify "what to prioritize (ease of implementation, performance, expressiveness, consistency)" and consider implementation accordingly.

## References
- [Write Policy in OPA, Enforce in SQL](https://blog.openpolicyagent.org/write-policy-in-opa-enforce-policy-in-sql-d9d24db93bf4)
- [GitHub Issue #1252: Pagination in OPA](https://github.com/open-policy-agent/opa/issues/1252)
- [AWS Prescriptive Guidance: Multi-Tenant API Access Authorization](https://docs.aws.amazon.com/ja_jp/prescriptive-guidance/latest/saas-multitenant-api-access-authorization/introduction.html)
