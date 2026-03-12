---
title: About Design Docs
description: An in-depth exploration of About Design Docs, covering design principles, trade-offs, and practical applications.
slug: design-docs-overview
date: 2022-10-07T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Design Docs
translation_key: design-docs-overview
---



# Overview
I researched about Design Docs.

# What are Design Docs
Design Docs are documents for software design.

They do not have a fixed format and are written in a way that is meaningful to the project.

Design Docs have several benefits in the development process, such as:

- Identifying design issues and reducing rework
- Building consensus on design
- Organizing and confirming cross-cutting concerns
- Sharing insights from senior engineers

Although there is no specific format, it is recommended to clarify the context, scope, goals, and non-goals of the design.

The length of the document should be short enough for busy people to read quickly.

Whether or not to write a Design Doc depends on whether the benefits of writing it outweigh the costs of maintaining it.

Design Docs have the following lifecycle:

- Creation and iteration (re-editing the document)
- Review
- Implementation and iteration (updating the document)
- Maintenance (updating the document) and learning (assisting system understanding for those who want to interact with the system)

Examples of Design Docs include:

- [docs.google.com - WebKit WebSocket design doc](https://docs.google.com/document/d/1s1ryja1V8dDotMK2WBGT2wnwchZ_x7Tag2L3OZfn5Po/preview)
- [www.chromium.org - Extensions](https://www.chromium.org/developers/design-documents/extensions/)

# Impressions
- It seems necessary to define and understand the form of documents that are valuable to the team or project
  - Why write it (Why), who reads it (Who), how to write it (How)
  - Especially if it's a design document or meeting minutes, it can be difficult for both the writer and reader if the document appears differently to different people
- If you decide on a document format, there is a risk of thinking within that format and finding it difficult to think outside the box
  - I feel that the purpose and understanding are more important than the format
- Consideration of the operational cost of documents (review, updates, etc.) and whether it is worth paying that cost
  - Like Mercari Shops, not updating might be an option. Or just updating when noticed
  - Just like the system, the operational aspect needs to be properly considered
- The lifecycle of Design Docs should be able to be integrated well into an agile development process
- The form of operation may change depending on the viewing range, such as external publication (API documents) or internal publication
- If you want to conduct a thorough review process, managing it as a repository on GitHub might be an option

# References
- [www.industrialempathy.com - Design Docs at Google](https://www.industrialempathy.com/posts/design-docs-at-google/)
- [messagepassing.github.io - The Uncoolness of Design Docs](https://messagepassing.github.io/011-designdocs/01-morrita/)
- [please-sleep.cou929.nu - About Google's Design Doc](https://please-sleep.cou929.nu/20091116.html)
- [medium.com - Writing Technical Design Docs](https://medium.com/machine-words/writing-technical-design-docs-71f446e42f2e)
- [tkybpp.hatenablog.com - [Translation] The Document "Design Docs at Google" That Google Engineers Always Write When Developing Software](https://tkybpp.hatenablog.com/entry/2020/08/03/090000)
  - Translation of [www.industrialempathy.com - Design Docs at Google](https://www.industrialempathy.com/posts/design-docs-at-google/)
- [myenigma.hatenablog.com - Introduction to Design Docs Used at Google, etc.](https://myenigma.hatenablog.com/entry/2021/07/25/140308)
  - Design Docs as a place to discuss parts not reflected in code (why another method was not chosen)
- [atmarkit.itmedia.co.jp - Can Overtime Be Reduced!? Introduction to Design Docs for Becoming a Senior Engineer](https://atmarkit.itmedia.co.jp/ait/articles/1606/21/news016.html)
  - Design Docs match the agile process
- [www.flywheel.jp - Learning Design Docs with Design Docs](https://www.flywheel.jp/topics/design-doc-of-design-doc/)
  - Early feedback on design can be obtained
- [engineering.mercari.com - Design Docs Operation at Mercari Shops](https://engineering.mercari.com/blog/entry/20220225-design-docs-by-mercari-shops/)
  - Design Docs are not updated or maintained in principle
    - The benefit of recognizing that it is not the latest design
- [nhiroki.jp - Thoughts on Design Docs](https://nhiroki.jp/2021/03/31/design-docs)
  - The obsolescence from not updating Design Docs is not much of an issue, as it has value as a snapshot
