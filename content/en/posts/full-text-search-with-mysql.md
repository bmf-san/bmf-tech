---
title: Full-Text Search in MySQL
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
MySQL has supported full-text search for quite some time, but I hadn't touched it until recently, so I decided to give it a light try.

# Starting Full-Text Search in MySQL
Using full-text search in MySQL is significantly less cumbersome than using ElasticSearch.

You can easily perform full-text searches by adding a **FULLTEXT INDEX** to the columns you want to search and executing a search query with **MATCH (col1,col2,...) AGAINST (expr [search_modifier])**.

ex. 
```sql
// Table with columns to which FULLTEXT INDEX will be added
CREATE TABLE `posts` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `title` varchar(255) DEFAULT NULL,
  `body` longtext DEFAULT NULL,
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

// Adding FULLTEXT INDEX
ALTER TABLE posts ADD FULLTEXT INDEX index_title_md_body (title, md_body) WITH PARSER ngram;

// Search query using MATCH ... AGAINST
SELECT
  *
FROM
  posts
WHERE MATCH (title, body)
AGAINST ("Full-Text Search in MySQL" IN BOOLEAN MODE)
```

You can add a FULLTEXT INDEX not only with ALTER TABLE but also with CREATE TABLE or CREATE INDEX.

# Full-Text Parsers
MySQL's full-text search supports ngram and MeCab parsers.

- [ngram Full-Text Parser](https://dev.mysql.com/doc/refman/8.0/en/fulltext-search-ngram.html)
- [MeCab Full-Text Parser Plugin](https://dev.mysql.com/doc/refman/8.0/en/fulltext-search-mecab.html)

By default, ngram is set.

If you want to use MeCab, you need to install the plugin.
 
# Full-Text Search Modes
There are three modes available, and you can specify which mode you want to use.

The search results can vary depending on the mode, so you have options based on the search experience you want to create.

- NATURAL LANGUAGE MODE
  - A mode that searches using natural language processing.
  - cf. [Natural Language Full-Text Search](https://dev.mysql.com/doc/refman/8.0/en/fulltext-natural-language.html)
- BOOLEAN MODE
  - A mode that searches using conditions like AND, OR, NOT.
  - cf. [Boolean Full-Text Search](https://dev.mysql.com/doc/refman/8.0/en/fulltext-boolean.html)
- QUERY EXPANSION
  - A mode that searches by adding synonyms and related words to the search terms.
  - cf. [Full-Text Search Using Query Expansion](https://dev.mysql.com/doc/refman/8.0/en/fulltext-query-expansion.html)

# Adjusting Search Accuracy
As an approach to adjust the nature of searches beyond the parser and full-text search modes:

- Changing parser settings
  - ex. Changing the token size for ngram
- Specifying character sets
  - ex. utf8_general_ci, utf8_unicode_ci etc...
- Adjusting full-text stop words
  - cf. [Full-Text Stop Words](https://dev.mysql.com/doc/refman/8.0/en/fulltext-stopwords.html)

Refer to [Fine-Tuning MySQL Full-Text Search](https://dev.mysql.com/doc/refman/8.0/en/fulltext-fine-tuning.html) as well.

# Thoughts
I have incorporated MySQL's full-text search feature into this blog.

[Search Article List](https://bmf-tech.com/posts/search?keyword=)

It seems to perform better than LIKE searches, but actual performance maintenance needs to be verified in different environments. However, I found that if the requirements are met, it is a sufficiently usable feature.