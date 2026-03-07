---
title: Tools for Automatically Generating DB Documentation (ER Diagrams, etc.) - schemaspy, tbls
slug: db-documentation-tools-schemaspy-tbls
date: 2020-07-09T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - Docker
  - ER
translation_key: db-documentation-tools-schemaspy-tbls
---

# Overview
When it comes to tools that can automatically generate DB documentation, MySQL Workbench is often mentioned as the leading option. However, I explored some other good OSS alternatives and wanted to share my impressions.

# Schemaspy
- [schemaspy.org](http://schemaspy.org/)
- [github.com - schemaspy/schemaspy](https://github.com/schemaspy/schemaspy)

This is a Java-based tool that connects to a database and generates DB documentation in HTML format.

There is an image available on Dockerhub, making it easy to try out.

For MySQL 5.7 (and probably 5.8 as well), you can use the following command:
```
docker run -v "$PWD/schema:/output" --net="host" schemaspy/schemaspy:latest \
 -t mysql -host {DBHOST}:{DBPORT} -db {DBNAME} -u {DBUSER} -p {DBPASSWORD}
```

For MySQL 5.6, you need to tweak the command slightly:
```
docker run -v "$PWD/schema:/output" --net="host" schemaspy/schemaspy:latest -t mysql -host {DBHOST}:{DBPORT} -db {DBNAME} -u {DBUSER} -p {DBPASSWORD} -connprops useSSL\=false -s {DBNAME}
```

Both can be tried out easily with a one-liner.

Of course, it works with databases other than MySQL as well.

# tbls
- [github.com - k1LoW/tbls](https://github.com/k1LoW/tbls)

This is a CI-friendly DB documentation tool that generates documentation in markdown format.

It can be installed via dep, rpm, brew, go, or docker.

The usage is simple, so refer to the README on GitHub.

I want to manage all documentation in markdown, so I have adopted it for my personal application documentation.

# Impressions
- If we compare only ER diagrams, MySQL Workbench might be the easiest to read. However, this depends on the number of tables. For a small number of tables, any tool seems fine.
- MySQL Workbench does not automatically drop tables that do not have relational connections, but schemaspy seems to handle that properly.
- The ER diagrams from schemaspy can be difficult to follow with many tables due to the relationship lines, so I think it would be good if the UI could be adjusted somehow. MySQL Workbench allows for manual adjustments, which provides flexibility (there is an auto-arrangement feature, but it seems to have limitations...).
- Regarding ER diagrams, I feel there are physical limits to outputting relationships for all tables, so some ingenuity from the user is necessary. I don't think there are many cases where one wants to see relationships for all tables at once, so narrowing down the tables would be a good approach. I wonder if schemaspy allows for filtering the tables to generate as ER diagrams... I haven't looked closely, but it doesn't seem obvious at first glance...
- tbls is CI-friendly, making it easy to integrate into CI, but schemaspy also seems relatively easy to incorporate.
- [rarejob-tech-dept.hatenablog.com - Integrating Automatic ER Diagram Creation into CI](https://rarejob-tech-dept.hatenablog.com/entry/2020/01/24/190000)

# References
- [ベジプロ - SchemaSpyでER図を出力できない](https://www.blog.v41.me/posts/749a3607-aa12-47d6-9441-8f7497602325)
- [sys-guard.com - SchemaSpyでER図を自動作成 on Docker 2019年7月22日 2019年7月22日 AWS, Linux/UNIX 129view](https://sys-guard.com/post-17119/)