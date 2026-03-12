---
title: Full-Text Search with MySQL
description: An in-depth look at Full-Text Search with MySQL, covering key concepts and practical insights.
slug: full-text-search-with-mysql
date: 2023-04-30T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - MySQL
translation_key: full-text-search-with-mysql
---

# Overview
MySQL has supported full-text search for quite some time, but I hadn't explored it until recently, so I decided to give it a try.

# Getting Started with Full-Text Search in MySQL
Using full-text search in MySQL is significantly less cumbersome than ElasticSearch.

You can easily perform full-text searches by adding a **FULLTEXT INDEX** to the columns you want to search and executing a query with **MATCH (col1,col2,...) AGAINST (expr [search_modifier])**.

ex. 
```sql
// Table with columns for FULLTEXT INDEX
CREATE TABLE `posts` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `title` varchar(255) DEFAULT NULL,
  `body` longtext DEFAULT NULL,
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

// Adding FULLTEXT INDEX
ALTER TABLE posts ADD FULLTEXT INDEX index_title_md_body (title, md_body) WITH PARSER ngram;

// Search query with MATCH ... AGAINST
SELECT
  *
FROM
  posts
WHERE MATCH (title, body)
AGAINST ("MySQLで全文検索" IN BOOLEAN MODE)
```

FULLTEXT INDEX can also be added using CREATE TABLE or CREATE INDEX.

# Full-Text Parsers
MySQL's full-text search supports ngram and MeCab parsers.

- [ngram Full-Text Parser](https://dev.mysql.com/doc/refman/8.0/en/fulltext-search-ngram.html)
- [MeCab Full-Text Parser Plugin](https://dev.mysql.com/doc/refman/8.0/en/fulltext-search-mecab.html)

By default, ngram is set.

If you want to use MeCab, you need to install the plugin.

# Full-Text Search Modes
There are three modes available, and you can specify which one to use.

The search results vary depending on the mode, so you can choose based on the search experience you want to provide.

- NATURAL LANGUAGE MODE
  - A mode that searches using natural language processing
  - cf. [Natural Language Full-Text Searches](https://dev.mysql.com/doc/refman/8.0/en/fulltext-natural-language.html)
- BOOLEAN MODE
  - A mode that searches using conditions like AND, OR, NOT
  - cf. [Boolean Full-Text Searches](https://dev.mysql.com/doc/refman/8.0/en/fulltext-boolean.html)
- QUERY EXPANSION
  - A mode that adds synonyms or related words to the search terms
  - cf. [Full-Text Searches with Query Expansion](https://dev.mysql.com/doc/refman/8.0/en/fulltext-query-expansion.html)

# Adjusting Search Precision
Apart from parsers and full-text search modes, there are other approaches to adjust the nature of the search:

- Changing parser settings
  - ex. Changing ngram token size
- Specifying character set
  - ex. utf8_general_ci, utf8_unicode_ci etc...
- Adjusting full-text stopwords
  - cf. [Full-Text Stopwords](https://dev.mysql.com/doc/refman/8.0/en/fulltext-stopwords.html)

Refer to [Fine-Tuning MySQL Full-Text Search](https://dev.mysql.com/doc/refman/8.0/en/fulltext-fine-tuning.html) as well.

# Impressions
I incorporated MySQL's full-text search feature into this blog.

[Search Article List](https://bmf-tech.com/posts/search?keyword=)

It seems to perform better than LIKE searches, but actual performance maintenance may vary by environment. However, if requirements are met, it is a sufficiently usable feature.