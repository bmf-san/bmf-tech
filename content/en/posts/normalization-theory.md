---
title: About Normalization Theory
slug: normalization-theory
date: 2024-04-02T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - Normalization Theory
  - Normalization
translation_key: normalization-theory
---

# What is Normalization
Normalization is the design process to eliminate data duplication and prevent logical inconsistencies in data.

# Premises
## Keys
- Primary Key
  - An identifier that uniquely identifies a row
- Composite Key
  - A primary key composed of multiple attributes
- Foreign Key
  - A key used to reference another table
- Candidate Key
  - A set of attributes that can uniquely identify a row, irreducible and minimal (with the least number of attributes)
- Super Key
  - A combination of attributes that can uniquely identify a row along with extra attributes
- Irreducible
  - A state where there are no extra attributes (i.e., cannot reduce attributes further)

| Employee Number | ID   | Name       | Gender | Address           | Phone Number   |
|----------------|------|------------|--------|-------------------|----------------|
| 1              | 1001 | Yamada Taro| Male   | Chiyoda, Tokyo    | 03-1234-5678   |
| 2              | 1002 | Tanaka Hanako| Female | Shibuya, Tokyo    | 03-2345-6789   |
| 3              | 1003 | Suzuki Jiro| Male   | Shinjuku, Tokyo    | 03-3456-7890   |
| 4              | 1004 | Sato Saburo| Male   | Minato, Tokyo     | 03-4567-8901   |
| 5              | 1005 | Takahashi Shiro| Male | Meguro, Tokyo     | 03-5678-9012   |

Primary Key: {Employee Number, ID}

Candidate Keys: {Employee Number, ID}, {Phone Number}

Super Keys: {Employee Number, ID}, {Employee Number, ID, Name}, {Employee Number, ID, Name, Gender} etc...

## Functional Dependency
- Functional Dependency
  - The property that if A is known, B can also be determined
- Partial Functional Dependency
  - The property that an attribute is dependent on one of the candidate keys
- Full Functional Dependency
  - The property that all non-key attributes are functionally dependent on the primary key
- Transitive Functional Dependency
  - A functional dependency between non-key attributes, where if A is known, B can be determined, and if B is known, C can also be determined
- Join Dependency
  - The property that when decomposed relations are joined, they return to the original relation (i.e., lossless decomposition)
- Multivalued Dependency
  - A relationship where attributes A and C are independent, and B depends on A
  - A special case of join dependency
- Lossless Decomposition
  - Decomposing a relation in such a way that it can be reconstructed by joining the decomposed parts back together

# Normal Forms
## First Normal Form (1NF)
- Rows are unordered vertically, and columns are unordered horizontally
  - Although SQL specifications have order, avoid writing queries that depend on order
    - e.g., avoid SELECT *, use column positions in ORDER BY (ORDER BY 1) etc.
- No duplicate rows exist
- NULL values are not included
- Each column contains only one value that satisfies the domain (data type)

## Second Normal Form (2NF)
- Satisfies 1NF
- Removes partial functional dependencies, achieving full functional dependency
  - All non-key attributes are fully functionally dependent on the candidate key

## Third Normal Form (3NF)
- Satisfies 2NF
- Removes transitive functional dependencies
  - All non-key attributes are not transitively functionally dependent on the candidate key

## Boyce-Codd Normal Form (BCNF)
- Satisfies 3NF
- Removes partial and transitive functional dependencies of candidate keys
- All non-trivial functional dependencies are removed, and further lossless decomposition based on functional dependencies is not possible
  - Trivial: {Order Number, Product Number} → {Product Number}
  - Non-trivial: {Product Name} → {Material Name}

## Fourth Normal Form (4NF)
- Normalization based on multivalued dependencies

## Fifth Normal Form (5NF)
- A state where all non-trivial or implicit join dependencies are removed

# Normalization
## 1NF
Unnormalized
| Invoice Number | Product Number | Product Name |
|:--------------|:---------------|:-------------|
| 1            | A001<br>A002<br>A003 | Apple<br>Orange<br>Banana |
| 2            | A004<br>A005<br>A006 | Grape<br>Pear<br>Strawberry |

1NF
| Invoice Number | Product Number | Product Name |
| :------------ | :------------- | :----------- |
| 1            | A001          | Apple       |
| 1            | A002          | Orange      |
| 1            | A003          | Banana      |
| 2            | A004          | Grape       |
| 2            | A005          | Pear        |
| 2            | A006          | Strawberry  |

## 2NF
Before 2NF
| Invoice Number | Product Number | Product Name |
| :------------ | :------------- | :----------- |
| 1            | A001          | Apple       |
| 1            | A002          | Orange      |
| 1            | A003          | Banana      |
| 2            | A004          | Grape       |
| 2            | A005          | Pear        |
| 2            | A006          | Strawberry  |

The candidate key {Invoice Number, Product Number} and the non-key attribute Product Name are partially functionally dependent.

2NF
Sales Detail Table
| Invoice Number | Product Number |
| :------------ | :------------- |
| 1            | A001          |
| 1            | A002          |
| 1            | A003          |
| 2            | A004          |
| 2            | A005          |
| 2            | A006          |

Product Table
| Product Number | Product Name |
| :------------- | :----------- |
| A001          | Apple       |
| A002          | Orange      |
| A003          | Banana      |
| A004          | Grape       |
| A005          | Pear        |
| A006          | Strawberry  |

Partial functional dependency occurs only when the candidate key is a composite key, so it does not occur when the candidate key is a single attribute.

## 3NF
Before 3NF
| Invoice Number | Product Number | Customer Number | Customer Name |
| :------------ | :------------- | :-------------- | :------------- |
| 1            | A001          | B1              | Apple Corp     |
| 1            | A002          | B1              | Orange Corp    |
| 1            | A003          | B1              | Banana Corp    |
| 2            | A004          | C1              | Grape Corp     |
| 2            | A005          | C1              | Pear Corp      |
| 2            | A006          | C1              | Strawberry Corp |

When the primary key is {Invoice Number, Product Number}, there is a transitive dependency {Invoice Number, Customer Number} → {Customer Number} → {Customer Name}.

3NF
Sales Detail Table
| Invoice Number | Product Number |
| :------------ | :------------- |
| 1            | A001          |
| 1            | A002          |
| 1            | A003          |
| 2            | A004          |
| 2            | A005          |
| 2            | A006          |

Customer Table
| Customer Number | Customer Name |
| :-------------- | :------------- |
| B1              | Apple Corp    |
| B1              | Orange Corp   |
| B1              | Banana Corp   |
| C1              | Grape Corp    |
| C1              | Pear Corp     |
| C1              | Strawberry Corp|

## BCNF
Before BCNF
| Name  | Subject | Teacher |
| ------ | ------ | ---- |
| Bob   | Math  | Yamada |
| Tom   | Math  | Sato  |
| John  | Math  | Suzuki|
| John  | English| Ando  |

When the primary key is {Name, Subject}, there is a functional dependency {Teacher} → {Subject}, where the determinant (A in A → B) is not a super key.

BCNF
Enrollment Table
| Name  | Subject |
| ------ | ------ |
| Bob   | Math  |
| Tom   | Math  |
| John  | Math  |
| John  | English|

Teacher Table
| Teacher | Subject |
| ---- | ------ |
| Yamada | Math  |
| Sato  | Math  |
| Suzuki| Math  |
| Ando  | English|

The information {Name, Subject} → {Teacher} is lost, so we can no longer determine who John's math teacher is.

## 4NF
Before 4NF
| Name |   Hobby   |   Favorite Food   |
| ---- | -------- | -------- |
| Tanaka | Baseball | Ramen   |
| Suzuki | Soccer   | Sushi    |
| Sato   | Basketball| Curry   |

When the primary key is {Name, Hobby, Favorite Food}, multiple attributes are determined by {Name} → {Hobby} → {Favorite Food}.

4NF
Hobby Table
| Name |   Hobby   |
| ---- | -------- |
| Tanaka | Baseball |
| Suzuki | Soccer   |
| Sato   | Basketball|

Favorite Food Table
| Name |   Favorite Food   |
| ---- | -------- |
| Tanaka | Ramen   |
| Suzuki | Sushi    |
| Sato   | Curry    |

## 5NF
Before 5NF
| Store  | Stock Item | Manufacturer |
| ------ | -------- | ------ |
| Tokyo   | TV       | Company A |
| Tokyo   | TV       | Company B |
| Tokyo   | PC       | Company A |
| Kanagawa | TV       | Company A |

{Store} → {Stock Item}, {Store} → {Manufacturer}, {Stock Item} → {Manufacturer} can be decomposed into multiple parts.

5NF
Stock Table
| Store  | Stock Item |
| ------ | -------- |
| Tokyo   | TV       |
| Tokyo   | TV       |
| Tokyo   | PC       |
| Kanagawa | TV       |

Supplier Table
| Store  | Supplier |
| ------ | ------ |
| Tokyo   | Company A |
| Tokyo   | Company B |
| Tokyo   | Company A |
| Kanagawa | Company A |

Manufacturer Table
| Store  | Manufacturer |
| ------ | ------ |
| Tokyo   | Company A |
| Tokyo   | Company B |
| Tokyo   | Company A |
| Kanagawa | Company A |

# Thoughts
I am not very confident in my understanding beyond BCNF...

# References
- [amzn.to - Practical Introduction to Database Theory from Theory ~ Efficient SQL Using Relational Model (WEB+DB PRESS plus)](https://amzn.to/3TEltzx)
- [e-words.jp - Partial Functional Dependency](https://e-words.jp/w/%E9%83%A8%E5%88%86%E9%96%A2%E6%95%B0%E5%BE%93%E5%B1%9E.html)
- [datascience-lab.sakura.ne.jp - Understanding Various Types of Keys](https://datascience-lab.sakura.ne.jp/primarykey/)
- [poppingcarp.com - Various Keys such as Primary Key, Candidate Key, Foreign Key, Super Key | Database Basics](https://poppingcarp.com/various_key/)
- [koseki2580.github.io - Normalization](https://koseki2580.github.io/study-docs/docs/Database/normalization/#:~:text=%E8%87%AA%E6%98%8E%E3%81%AA%E9%96%A2%E6%95%B0%E5%BE%93%E5%B1%9E%E6%80%A7%20%E3%81%A8%E3%81%AF%20X%20%E2%86%92%20Y,%E3%82%AD%E3%83%BC%E3%81%AE%E3%81%93%E3%81%A8%E3%82%92%E3%81%84%E3%81%86%E3%80%82)
- [zenn.dev - Understanding Database Normalization through Illustrations](https://zenn.dev/keisuke90/articles/66ecb7956a6816#%E7%AC%AC%EF%BC%95%E6%AD%A3%E8%A6%8F%E5%BD%A2)
- [youtube.com - Database Normalization (1NF, 2NF, 3NF)](https://www.youtube.com/watch?app=desktop&v=zcLCZKOAOjE)
- [tabibou.com - Database Normalization from 1NF to 5NF, including Boyce-Codd Normalization Examples](https://tapibou.com/seikika)
- [poppingcarp.com - Various Keys such as Primary Key, Candidate Key, Foreign Key, Super Key | Database Basics](https://poppingcarp.com/various_key/)