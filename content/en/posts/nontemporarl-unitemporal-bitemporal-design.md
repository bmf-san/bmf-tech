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
description: Exploring the features, design examples, advantages, and disadvantages of different temporal data models.
translation_key: nontemporarl-unitemporal-bitemporal-design
---

In data models, there are several patterns based on how the time axis (such as history or validity periods) is managed.

1. **Nontemporal**
2. **Unitemporal**
3. **Bitemporal**

Each differs in terms of "how finely and in what sense time information is managed."

This post explains the characteristics, design examples, advantages, and disadvantages of these data models.

## 1. Nontemporal Data Model

### Characteristics
- A data model that manages "only the present," either not handling the time axis at all or only to the extent of reference dates.
- Database records always maintain only the latest state, without storing past or future history or validity periods.

### Design Example
- Typical OLTP system table structures or operational systems that do not require history management.
- Table definition example:
  ```sql
  CREATE TABLE product (
      product_id    INT PRIMARY KEY,
      product_name  VARCHAR(100),
      price         INT,
      update_date   DATETIME DEFAULT CURRENT_TIMESTAMP,
      -- update_date merely indicates when it was updated, not managing validity periods
      ...
  );
  ```

### Advantages
- Simple table structure with low development and operational costs.
- Simple queries with high performance (no need to consider history).

### Disadvantages
- Cannot track past or future states. Not suitable for scenarios requiring historical analysis.
- Information is overwritten with each update, making it impossible to accurately track "when" and "how" changes occurred for audit or legal compliance.

## 2. Unitemporal Data Model

### Characteristics
- A data model managing a "**single time axis**." Typically manages either "system time (the period a record exists in the database)" or "business validity period (such as business start and end dates)."
- Implements date/time columns or version management mechanisms for history tracking.

### Pattern Examples
1. **Managing System Time (Transaction Time)**
   - Manages "when this record was INSERTed and when it was DELETEd (or invalidated)" in the database.
   - Useful mainly for audit purposes or when referencing historical data in the DB.

2. **Managing Business Time (Validity Period)**
   - Represents "when this record is valid from a business perspective."
   - Used in cases where business logic's "valid dates," such as prices, contract periods, or organizational restructuring periods, are important.

### Design Example (Business Time Management)
- Example of a "contract" unitemporal model
  ```sql
  CREATE TABLE contract (
      contract_id         INT PRIMARY KEY,
      contract_name       VARCHAR(100),
      valid_from          DATE,   -- The date this contract starts
      valid_to            DATE,   -- The date this contract ends (business validity period)
      updated_at          DATETIME DEFAULT CURRENT_TIMESTAMP,
      ...
  );
  ```
- Requires operational rules for INSERT/UPDATE and application logic to manage and reference past contract details and future validity periods.

### Advantages
- **More flexible history tracking than Nontemporal.** Easier to reproduce the business state at a specified time.
- When combining version management or audit logs, models retaining system time are often used.

### Disadvantages
- More complex table management than Nontemporal. Requires logic to maintain period consistency during INSERT/UPDATE.
- As versions increase, data volume swells, increasing query complexity and performance load.

## 3. Bitemporal Data Model

### Characteristics
- Manages both **system time** (the period a record is valid in the DB) and **business time** (the period valid from a business perspective).
- For example, manages the following two time axes:
  1. **Validity Period (Business Time)**: The period indicating "when this record is valid from a business perspective."
  2. **Transaction Period (System Time)**: The period indicating "when this record existed in the database."
- Capable of handling complex change management, such as "at a certain past point, the business state was this, but when was it registered and modified in the DB," "want to retroactively correct future data," or "want to retroactively correct past validity periods."

### Design Example
Below is an example. In practice, types and constraints vary depending on the DB product and requirements, but the concept remains the same.

```sql
CREATE TABLE contract_bitemporal (
    contract_id          INT,
    contract_name        VARCHAR(100),
    business_from        DATE,   -- Business validity start date
    business_to          DATE,   -- Business validity end date
    system_from          TIMESTAMP, -- Time this record became valid in the DB
    system_to            TIMESTAMP, -- Time this record was invalidated (or logically deleted) in the DB
    ...
    PRIMARY KEY (contract_id, business_from, system_from)
    -- PK design method depends on requirements
);
```

- **business_from, business_to**
  Business contract validity period. For example, set as valid from January 1, 2025, to December 31, 2025.
- **system_from, system_to**
  Insert the date/time the record was INSERTed into the DB into `system_from`, and when the record is UPDATEd and a new version is INSERTed, set the `system_to` of the old record.
- In bitemporal, it becomes possible to query history from both "when (business perspective)" and "when (database perspective)."

### Advantages
- **Most flexible and complete history management possible.** Can reproduce the past business validity period state and "when viewed from the DB perspective."
- Effective when high-level change history management is required for audit or legal requirements, such as "want to retroactively correct past data and maintain correct history again."

### Disadvantages
- Most complex design and operation. Requires complex controls and rules for all data Insert/Update/Delete.
- Data volume and version count can increase rapidly, necessitating performance and storage management ingenuity.
- Requires application logic and query development to avoid confusing system time and business time.

---

## Summary and Selection Points
1. **Nontemporal**
   - Optimal for scenarios where "only managing the current state is sufficient" and "no need to trace history."
   - Simplest design, implementation, performance, and operation.
2. **Unitemporal**
   - Effective when some degree of past or future history management is desired.
   - Choose when either system time or business time is critical.
3. **Bitemporal**
   - Used when both system time and business time need to be accurately managed, and past states need to be flexibly reproduced and corrected.
   - Most functional but challenging in design, implementation, and operation, with data volume prone to increase.

### Considerations for Selection
- **Strict audit/legal compliance requirements** or **frequent retroactive corrections** often necessitate bitemporal.
- Weigh implementation cost and performance against system load and requirement complexity.
- While many prefer to prepare for future expansion with Unitemporal or Bitemporal, initial costs and operations become heavier, so balancing requirement priorities is crucial.

**Supplementary**
- Some RDBMS products have built-in system time history management (so-called "temporal table" functionality) (e.g., SQL Server's System-Versioned Temporal Table, Oracle's Flashback, etc.).
- Whether bitemporal is natively supported varies by DB product, but even without standard features, it can be achieved through table design and operational logic (though complexity increases).