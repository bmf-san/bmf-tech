---
title: Characteristics and Design of Nontemporal, Unitemporal, and Bitemporal
slug: nontemporarl-unitemporal-bitemporal-design
date: 2025-02-27T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - Bi-Temporal
  - Uni-Temporal
  - Non-Temporal
  - DB
translation_key: nontemporarl-unitemporal-bitemporal-design
---

Data models can be categorized based on how they manage time axes (such as history and validity periods).

1. **Nontemporal**
2. **Unitemporal**
3. **Bitemporal**

Each of these differs in terms of how finely and in what sense they manage time information.

This post explains the characteristics, design examples, advantages, and disadvantages of these data models.

## 1. Nontemporal Data Model

### Characteristics
- A data model that manages only the "current state" and **does not handle the time axis at all**, or only uses it for reference purposes.
- Records in the database always maintain the latest state and do not store past or future history or validity periods.

### Design Example
- Common OLTP system table structures or operational systems that do not require history management.
- Table definition example:
  ```sql
  CREATE TABLE product (
      product_id    INT PRIMARY KEY,
      product_name  VARCHAR(100),
      price         INT,
      update_date   DATETIME DEFAULT CURRENT_TIMESTAMP,
      -- update_date only indicates when it was updated, without managing validity periods
      ...
  );
  ```

### Advantages
- Simple table structure with low development and operational costs.
- Queries are straightforward, and performance is high (no need to consider history).

### Disadvantages
- Cannot track past or future states. Cannot accommodate cases where historical analysis is needed.
- Information is overwritten with each update, making it impossible to accurately track "when" and "how" changes occurred for auditing or compliance purposes.

## 2. Unitemporal Data Model

### Characteristics
- A data model that manages a **single time axis**. Generally manages either "system time (the period during which the record existed in the database)" or "business validity period (such as start and end dates for business purposes)".
- Implements date/time columns or version management mechanisms for history tracking.

### Pattern Examples
1. **Managing System Time (Transaction Time)**
   - Manages "when this record was INSERTed and when it was DELETEd (or invalidated)" in the database.
   - Useful mainly for auditing purposes or when referencing historical data in the DB.

2. **Managing Business Time (Validity Period)**
   - Represents "when this record is valid for business purposes".
   - Used in cases where business logic requires important "effective dates" such as pricing, contract periods, or organizational restructuring.

### Design Example (Business Time Management)
- Example of a unitemporal model for "contracts"
  ```sql
  CREATE TABLE contract (
      contract_id         INT PRIMARY KEY,
      contract_name       VARCHAR(100),
      valid_from          DATE,   -- Start date of this contract
      valid_to            DATE,   -- End date of this contract (business validity period)
      updated_at          DATETIME DEFAULT CURRENT_TIMESTAMP,
      ...
  );
  ```
- To reference and manage past contract details and future validity periods, operational rules for INSERT/UPDATE and application logic are necessary.

### Advantages
- **More flexible history tracking than Nontemporal**. Easier to reproduce the business state at a specified time.
- Models that retain system time are often used when combining version management and audit logs.

### Disadvantages
- Table management becomes more complex than Nontemporal. Logic is needed to maintain consistency during INSERT/UPDATE.
- As versions increase, data volume can grow, increasing query complexity and performance load.

## 3. Bitemporal Data Model

### Characteristics
- Manages both **system time** (the period during which the record is valid in the DB) and **business time** (the period during which it is valid for business).
- For example, it manages the following two time axes:
  1. **Validity Period (Business Time)**: Indicates the period during which this record is valid for business purposes.
  2. **Transaction Period (System Time)**: Indicates the period during which this record existed in the database.
- Can accommodate complex change management scenarios such as "the business state at a certain point in the past, but when was it registered and modified in the DB?", "want to retroactively correct future data", or "want to retroactively correct past validity periods".

### Design Example
Here is one example. In practice, types and constraints vary depending on the DB product and requirements, but the concept remains the same.

```sql
CREATE TABLE contract_bitemporal (
    contract_id          INT,
    contract_name        VARCHAR(100),
    business_from        DATE,   -- Business effective start date
    business_to          DATE,   -- Business effective end date
    system_from          TIMESTAMP, -- Time when this record became valid in the DB
    system_to            TIMESTAMP, -- Time when this record was invalidated (or logically deleted) in the DB
    ...
    PRIMARY KEY (contract_id, business_from, system_from)
    -- PK design method depends on requirements
);
```

- **business_from, business_to**
  Business contract validity period. For example, set to be valid from January 1, 2025, to December 31, 2025.
- **system_from, system_to**
  Insert the timestamp when the record was INSERTed into `system_from`, and when the record is updated and a new version is INSERTed, set the `system_to` of the old record.
- In bitemporal, it becomes possible to query history from both the "when (business perspective)" and "when (database perspective)".

### Advantages
- **Most flexible and complete history management possible**. Can reproduce the state of past business validity periods, along with "at what point in time from the DB perspective".
- Effective for advanced change history management required for auditing or legal compliance, such as "want to correct past data later and maintain a correct history again".

### Disadvantages
- Design and operation are the most complex. Complex controls and rules are required for all data Insert/Update/Delete operations.
- Data volume and version numbers can increase rapidly, necessitating careful performance and storage management.
- Application logic and queries must be carefully crafted to avoid confusing system time and business time.

---

## Summary and Selection Points
1. **Nontemporal**
   - Best for scenarios where "only the current state needs to be managed" and "there is no need to trace history".
   - Design, implementation, performance, and operation are the simplest.
2. **Unitemporal**
   - Effective when wanting to manage some history of past or future.
   - Choose when either system time or business time is critical.
3. **Bitemporal**
   - Used when needing to accurately manage both system time and business time and flexibly reproduce or correct past states.
   - Most feature-rich but complex in design, implementation, and operation, and data volume can increase easily.

### Considerations for Selection Criteria
- Bitemporal is often required in **businesses with strict auditing and legal compliance** or **where retroactive corrections frequently occur**.
- Assess trade-offs with implementation costs and performance against system load and complexity of requirements.
- Many prefer to set up Unitemporal or Bitemporal for future expansion, but this increases initial costs and operational burdens, making it important to balance with the priority of requirements.

**Note**
- Some RDBMS products natively manage system time history (so-called "temporal table" features) (e.g., SQL Server's System-Versioned Temporal Table, Oracle's Flashback, etc.).
- Whether a DB product natively supports bitemporal varies, but even without standard features, it can be achieved through table design and operational logic (though it becomes more complex).