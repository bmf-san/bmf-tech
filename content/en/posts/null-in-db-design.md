---
title: "NULL in Database Design: Common Pitfalls and Best Practices"
description: 'Learn how NULL works in relational databases, the three-valued logic problem, common query mistakes with NULL, and schema design strategies to avoid NULL-related bugs.'
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
Indicates "value does not exist" or "value is unknown."

Cannot be compared like a value because it is not a value.

It is not an empty set (a set with 0 elements) but is considered a non-existent set.

# NULL Contrary to Relational Model
The relational model is based on the closed world assumption (assuming that anything not proven true is false), and NULL contradicts this.

The relational model is based on binary logic, so three-valued logic, which includes something other than true and false, is unacceptable.

# Impact of NULL
- NULL remains NULL even after operations or string manipulations
- Possibility of unintended SELECT results
  - Queries change depending on how NULL is interpreted
- Introduces three-valued logic (TRUE, FALSE, Unknown)
  - SQL becomes more complex
- Adverse effects on the optimizer
  - The presence of NULL affects the optimizer's calculations
    - Negatively impacts rewriting queries for optimal performance and estimating query costs

# Measures Against NULL
- Table Normalization
  - By advancing normalization, columns can be set to NOT NULL
    - Avoid using default values that are synonymous with NULL while defining NOT NULL
- Use of COALESCE function
  - A function that allows setting a default value when a specified column is NULL
    - IFNULL also exists, but IFNULL is not SQL standard
  - Can be effectively used in patterns where SQL evaluation results in NULL (e.g., results of aggregate functions like SUM or AVG, results of OUTER JOIN, evaluation results of NULLIF, etc.)
  
# Empty String vs. NULL
An empty string has a length of 0 and is an existing string, whereas NULL is a non-existent set, and the two are distinguished.

# Cases Allowing NULL
It may be acceptable to allow NULL when dealing with data that does not fit the relational model.

# References
- [amzn.to - Learn Databases from Theory to Practice ~ Efficient SQL with Relational Model (WEB+DB PRESS plus)](https://amzn.to/3TEltzx)
