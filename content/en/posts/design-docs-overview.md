---
title: About Design Docs
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

They do not have a fixed format, and the rule is to write them in a way that is meaningful for the project.

Design Docs have several benefits in the development process, such as:

- Identifying design issues and reducing rework
- Building consensus on design
- Organizing and confirming cross-cutting concerns
- Sharing knowledge from senior engineers

While there is no specific format, it is recommended to clarify the context, scope, goals, and non-goals of the design.

Regarding the length of the document, it is recommended to keep it short enough for busy people to read quickly.

Whether to write Design Docs should be based on whether the benefits of writing them outweigh the costs of maintaining them.

Design Docs have the following lifecycle:

- Creation and iteration (document re-editing)
- Review
- Implementation and iteration (document updates)
- Maintenance (document updates) and learning (assisting understanding of the system for those trying to interact with it)

Examples of Design Docs include:

- [docs.google.com - WebKit WebSocket design doc](https://docs.google.com/document/d/1s1ryja1V8dDotMK2WBGT2wnwchZ_x7Tag2L3OZfn5Po/preview)
- [www.chromium.org - Extensions](https://www.chromium.org/developers/design-documents/extensions/)

# Thoughts
- It seems necessary to define and understand the form of documents that are valuable to the team or project.
  - Why to write (Why), who will read (Who), how to write (How)
  - Especially if it differs in appearance between a design document and meeting minutes, it can be painful for both the writer and the reader.
- If you decide on a document type, there is a risk of thinking in a way that fits that type or finding it difficult to think outside of it.
  - I feel that awareness and understanding of the purpose are more important than the type.
- It is essential to carefully consider the operational costs of the document (maintenance such as review and updates) and whether it is worth paying those costs.
  - It might be acceptable not to update it like Mercari shop. Or perhaps just update it when you notice something.
  - Just like with systems, it is necessary to properly consider how to manage the operational aspects.
- The lifecycle of Design Docs should be well integrated into the agile development process.
- The form of operation may change depending on the viewing range, such as external publication (API documentation) or internal publication.
- If you want to conduct a thorough review process, managing the repository on GitHub might be an option.

# References
- [www.industrialempathy.com - Design Docs at Google](https://www.industrialempathy.com/posts/design-docs-at-google/)
- [messagepassing.github.io - The Inadequacy of Design Docs](https://messagepassing.github.io/011-designdocs/01-morrita/)
- [please-sleep.cou929.nu - About Google’s Design Doc](https://please-sleep.cou929.nu/20091116.html)
- [medium.com - Writing Technical Design Docs](https://medium.com/machine-words/writing-technical-design-docs-71f446e42f2e)
- [tkybpp.hatenablog.com - [Translation] The Document That Google Engineers Always Write When Developing Software: "Design Docs at Google"](https://tkybpp.hatenablog.com/entry/2020/08/03/090000)
  - Translation of [www.industrialempathy.com - Design Docs at Google](https://www.industrialempathy.com/posts/design-docs-at-google/)
- [myenigma.hatenablog.com - Introduction to Design Docs Used by Google and Others](https://myenigma.hatenablog.com/entry/2021/07/25/140308)
  - Design Docs as a place to discuss aspects not reflected in code (e.g., why a different method was not chosen)
- [atmarkit.itmedia.co.jp - Can Reduce Overtime!? A Beginner's Guide to Design Docs for Becoming a Senior Engineer](https://atmarkit.itmedia.co.jp/ait/articles/1606/21/news016.html)
  - Design Docs match the agile process.
- [www.flywheel.jp - Learning Design Docs from Design Docs](https://www.flywheel.jp/topics/design-doc-of-design-doc/)
  - Early feedback on design can be obtained.
- [engineering.mercari.com - About the Operation of Design Docs at Mercari Shops](https://engineering.mercari.com/blog/entry/20220225-design-docs-by-mercari-shops/)
  - Design Docs are generally not updated or maintained.
    - There is a benefit of being aware that it is not the latest design.
- [nhiroki.jp - Thoughts on Design Docs](https://nhiroki.jp/2021/03/31/design-docs)
  - The obsolescence of not updating Design Docs is not a significant issue; they have value as a snapshot.