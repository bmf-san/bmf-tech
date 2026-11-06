---
title: "{{.Title}}"
slug: {{.Slug}}
date: {{.Date}}
author: bmf-san
categories:
  -
tags:
  -
description: ""
translation_key: {{.Slug}}
recommended_books:
  - asin: ""
    title: ""
  - asin: ""
    title: ""
  - asin: ""
    title: ""
draft: true
---

<!--
Roundup (pillar) post scaffold. Targets "best books for X" search intent.
- List each book's ASIN + title under recommended_books: -> cover cards and ItemList structured data render at the end (books: is for single-book reviews and feeds the bookshelf widgets, so do not use it here).
- To also show covers at the top, add the "Book Review" tag (gated on $isReview).
- Link to individual review posts from the body where they exist (internal SEO / navigation).
- Keep description ~120 chars: who it is for + how the books were chosen.
- The picks are Japanese editions, so say so in the title and description (e.g. "8 Japanese Books for ...") and keep the note below; English "best books" searchers otherwise expect English editions.
- Create the JA version with the same SLUG via make new-ja-roundup so translation_key matches.
-->

(Two or three sentences: who this is for and how these books were chosen. Answer the search intent.)

The books in this list are Japanese editions, and the store links go to Amazon Japan.

## Why read about this topic

(Context and motivation. What readers gain from this area and these books.)

## The books

### 1. Title

(3-5 lines: who it suits, what you learn. Link the individual review if one exists: [review](/posts/slug/).)

### 2. Title

(Same.)

### 3. Title

(Same.)

## Wrap-up

(Next step for the reader: a suggested reading order, links to related roundups.)
