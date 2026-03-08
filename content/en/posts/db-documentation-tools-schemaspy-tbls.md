---
title: Tools for Automatically Generating DB Documents (ER Diagrams, etc.) - schemaspy, tbls
slug: db-documentation-tools-schemaspy-tbls
date: 2020-07-09T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - Docker
  - ER
description: Exploring and reviewing open-source tools for automatically generating DB documents.
translation_key: db-documentation-tools-schemaspy-tbls
---

# Overview
When it comes to tools for automatically generating DB documents, MySQL Workbench is probably the first that comes to mind. However, I explored other good open-source software and tried them out, so here are my thoughts.

# Schemaspy
- [schemaspy.org](http://schemaspy.org/)
- [github.com - schemaspy/schemaspy](https://github.com/schemaspy/schemaspy)

A Java-based tool that connects to a database and generates DB documents in HTML format.

There is an image available on Dockerhub, so you can easily try it out.

For MySQL 5.7 (probably 5.8 as well), you should be able to run it like this:
`docker run -v "$PWD/schema:/output" --net="host" schemaspy/schemaspy:latest \
 -t mysql -host {DBHOST}:{DBPORT} -db {DBNAME} -u {DBUSER} -p {DBPASSWORD}`

In a MySQL 5.6 environment, you need to tweak the command a bit.
`docker run -v "$PWD/schema:/output" --net="host" schemaspy/schemaspy:latest -t mysql -host {DBHOST}:{DBPORT} -db {DBNAME} -u {DBUSER} -p {DBPASSWORD} -connprops useSSL\=false -s {DBNAME}`

Both can be tested with a one-liner, making it simple.

Of course, it's not limited to MySQL.

# tbls
- [github.com - k1Low/tbls](https://github.com/k1LoW/tbls)

A CI-friendly DB documentation tool that generates documents in markdown.

It can be installed via dep, rpm, brew, go, or docker.

Usage is simple, so refer to the GitHub README.

I adopted it for documenting my personal applications because I want to manage all documents in markdown.

# Impressions
- When comparing only ER diagrams, MySQL Workbench might be the easiest to view. However, it depends on the number of tables. If there are few tables, any tool seems fine.
- MySQL Workbench does not automatically exclude tables without relational connections when generating ER diagrams, but schemaspy seems to handle this properly.
- With schemaspy, if there are many tables, the relationship lines can be hard to follow, so it would be nice if the UI allowed for some adjustments. MySQL Workbench allows manual adjustments, offering flexibility (though there is an auto-arrange feature, it seems limited...)
- Regarding ER diagrams, there is a physical limit to outputting relationships for all tables, so users might need to be creative. I don't think there are many cases where you want to view relationships for all tables at once, so narrowing down the tables might be a good approach. I wonder if schemaspy allows narrowing down the tables to generate as ER diagrams... I haven't looked closely, but at a glance, it doesn't seem possible...
- tbls is CI-friendly, making it easy to integrate into CI, but schemaspy also seems relatively easy to integrate.
- [rarejob-tech-dept.hatenablog.com - Integrating ER Diagram Auto-generation into CI](https://rarejob-tech-dept.hatenablog.com/entry/2020/01/24/190000)

# References
- [ベジプロ - Unable to Output ER Diagrams with SchemaSpy](https://www.blog.v41.me/posts/749a3607-aa12-47d6-9441-8f7497602325)
- [sys-guard.com - Automatically Creating ER Diagrams with SchemaSpy on Docker](https://sys-guard.com/post-17119/)