---
title: About NULL in Database Design
slug: null-in-db-design
date: 2024-03-30T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - DB
translation_key: null-in-db-design
---

# What is NULL
Indicates "no value exists" or "value is unknown."

Since it is not a value, it cannot be compared like a value.

It is considered a non-existent set, not an empty set (a set with zero elements).

# NULL Contradicting the Relational Model
The relational model is based on the closed world assumption (assuming everything that is not proven true is false), and NULL contradicts this.

Since the relational model is based on binary logic, the existence of ternary logic (TRUE, FALSE, Unknown) is unacceptable.

# Impact of NULL
- NULL remains NULL even when operated on or manipulated as a string.
- Possibility of unintended SELECT results
  - The interpretation of how NULL is handled can change the query.
- Introduces ternary logic (TRUE, FALSE, Unknown)
  - Makes SQL more complex.
- Negative impact on the optimizer
  - The existence of NULL affects the optimizer's calculations.
    - It negatively impacts the rewriting of queries for optimal performance and cost estimation of queries.

# Measures Against NULL
- Table normalization
  - By advancing normalization, columns can be made NOT NULL.
    - Avoid using default values that are synonymous with NULL while defining NOT NULL.
- Use of COALESCE function
  - A function that allows setting a default value when the specified column is NULL.
    - There is also IFNULL, but IFNULL is not part of the SQL standard.
  - Can be effectively used in scenarios where SQL evaluation results in NULL (e.g., results of aggregate functions like SUM or AVG, results of OUTER JOIN, evaluation results of NULLIF, etc.).

# Empty String vs NULL
An empty string has a length of 0 and is an existing string, while NULL is a non-existent set, and the two are distinct.

# Cases Where NULL is Acceptable
It may be acceptable to allow NULL when dealing with data that does not conform to the relational model.

# References
- [amzn.to - Practical Database Introduction from Theory ~ Efficient SQL with Relational Model (WEB+DB PRESS plus)](https://amzn.to/3TEltzx)